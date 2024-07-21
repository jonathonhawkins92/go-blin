package myparser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

// Goal:
// Generate openapi return model from all possible custom responses

func getRuntimeFunctionInfo(fn interface{}) (name, file string, line int) {
	funcValue := reflect.ValueOf(fn)
	funcPtr := funcValue.Pointer()
	funcForPC := runtime.FuncForPC(funcPtr)
	if funcForPC != nil {
		name := funcForPC.Name()
		file, line := funcForPC.FileLine(funcPtr)
		return name, file, line
	}
	return
}

func analyzeFunction(fn interface{}, filepath string) {
	// Get the function name
	// funcValue := reflect.ValueOf(fn)
	// funcName := runtime.FuncForPC(funcValue.Pointer()).Name()

	funcName, fileName, lineNumber := getRuntimeFunctionInfo(fn)

	fmt.Printf("Analyzing function: %s (File: %s:%d)\n", funcName, fileName, lineNumber)
	// Parse the file containing the function
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filepath, nil, 0)
	if err != nil {
		fmt.Printf("Error parsing file: %v\n", err)
		return
	}

	// Find the function in the AST
	var funcNode ast.Node
	ast.Inspect(file, func(n ast.Node) bool {
		if _, ok := n.(*ast.FuncLit); ok {
			if fset.Position(n.Pos()).Line == lineNumber {
				funcNode = n
				return false
			}
		}
		return true
	})

	if funcNode == nil {
		fmt.Println("Function not found in AST")
		return
	}

	// Analyze return statements
	returnTypes := analyzeReturns(funcNode)

	fmt.Println("Possible return types:")
	for returnType := range returnTypes {
		fmt.Printf("  - %s\n", returnType)
	}
}

func analyzeReturns(node ast.Node) map[string]bool {
	returnTypes := make(map[string]bool)
	ast.Inspect(node, func(n ast.Node) bool {
		if ret, ok := n.(*ast.ReturnStmt); ok {
			for _, expr := range ret.Results {
				if t := getExpressionType(expr); t != "" {
					returnTypes[t] = true
				}
			}
		}
		return true
	})
	return returnTypes
}

func getExpressionType(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.Ident:
		return e.Name
	case *ast.CallExpr:
		if ident, ok := e.Fun.(*ast.Ident); ok {
			return ident.Name
		}
	case *ast.UnaryExpr:
		return getExpressionType(e.X)
	case *ast.CompositeLit:
		if t, ok := e.Type.(*ast.Ident); ok {
			return t.Name
		}
	case *ast.FuncLit:
		return "func"
	case *ast.BasicLit:
		return strings.ToLower(e.Kind.String())
	}
	return fmt.Sprintf("unknown(%T)", expr)
}

func extractFunctionSource(fn interface{}) (string, error) {
	// Get function name and file
	_, fileName, lineNumber := getRuntimeFunctionInfo(fn)

	// Read the entire source file
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		return "", fmt.Errorf("error parsing file: %v", err)
	}

	// Find the function in the AST
	var funcNode ast.Node
	ast.Inspect(file, func(n ast.Node) bool {
		if n == nil {
			return true
		}
		if fset.Position(n.Pos()).Line == lineNumber {
			switch n.(type) {
			case *ast.FuncDecl, *ast.FuncLit:
				funcNode = n
				return false
			}
		}
		return true
	})

	if funcNode == nil {
		return "", fmt.Errorf("function not found in AST")
	}

	// Extract the function source
	var buf strings.Builder
	err = printer.Fprint(&buf, fset, funcNode)
	if err != nil {
		return "", fmt.Errorf("error printing function: %v", err)
	}

	return buf.String(), nil
}

func PrintFunction(fn interface{}) {

	// Get the caller's information
	_, callerFile, _, ok := runtime.Caller(1)
	if !ok {
		fmt.Printf("Unable to get caller information")
		return
	}

	// Get the absolute path of the caller's file
	absPath, err := filepath.Abs(callerFile)
	if err != nil {
		fmt.Printf("Error getting absolute path: " + err.Error())
		return
	}
	fmt.Println(absPath)

	res, err := extractFunctionSource(fn)
	if err != nil {
		fmt.Printf("Error extractFunctionSource: " + err.Error())
		return
	}
	fmt.Println(res)
}

func AnalyzeFunction(fn interface{}) {

	// Get the caller's information
	_, callerFile, _, ok := runtime.Caller(1)
	if !ok {
		fmt.Printf("Unable to get caller information")
		return
	}

	// Get the absolute path of the caller's file
	absPath, err := filepath.Abs(callerFile)
	if err != nil {
		fmt.Printf("Error getting absolute path: " + err.Error())
		return
	}
	analyzeFunction(fn, absPath)

	res, err := extractFunctionSource(fn)
	if err != nil {
		fmt.Printf("Error extractFunctionSource: " + err.Error())
		return
	}
	fmt.Println(res)
	// // Get the function value and type
	// funcValue := reflect.ValueOf(fn)
	// funcType := funcValue.Type()

	// // Get the function name and package path
	// funcPtr := funcValue.Pointer()
	// funcName := runtime.FuncForPC(funcPtr).Name()
	// lastDot := strings.LastIndexByte(funcName, '.')
	// packagePath, shortFuncName := funcName[:lastDot], funcName[lastDot+1:]

	// fmt.Printf("Analyzing function: %s/%s\n", funcName, shortFuncName)
	// fmt.Printf("Package: %s\n", packagePath)
	// fmt.Printf("Number of parameters: %d\n", funcType.NumIn())
	// fmt.Printf("Number of return values: %d\n", funcType.NumOut())

	// // Get function body
	// funcBody := extractFunctionBody(fn)

	// // Parse the function body
	// // fset := token.NewFileSet()
	// expr, err := parser.ParseExpr(funcBody)
	// if err != nil {
	// 	fmt.Printf("Error parsing function body: %v\n", err)
	// 	return
	// }

	// // Analyze return statements
	// returnTypes := make(map[string]bool)
	// ast.Inspect(expr, func(n ast.Node) bool {
	// 	if ret, ok := n.(*ast.ReturnStmt); ok {
	// 		for _, result := range ret.Results {
	// 			returnTypes[fmt.Sprintf("%T", result)] = true
	// 		}
	// 	}
	// 	return true
	// })

	// fmt.Println("Return statement types found:")
	// for returnType := range returnTypes {
	// 	fmt.Printf("  - %s\n", returnType)
	// }
}
