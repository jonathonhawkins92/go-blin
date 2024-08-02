package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"go/types"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/go/packages"
)

// Helper function to get the text of an API call
func getAPICallText(callExpr *ast.CallExpr, fset *token.FileSet) string {
	var callText strings.Builder

	// Get the method name (e.g., "Get", "Post", etc.)
	if sel, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
		callText.WriteString(sel.Sel.Name)
	}

	callText.WriteString("(")

	// Convert arguments to string
	for i, arg := range callExpr.Args {
		if i > 0 {
			callText.WriteString(", ")
		}
		callText.WriteString(exprToString(arg, fset))
	}

	callText.WriteString(")")

	return callText.String()
}

type APICall struct {
	Method        string
	Path          string
	config        map[string]string
	HandlerParams string
	HandlerReturn string
}

func (c APICall) String() string {
	var config strings.Builder
	for key, value := range c.config {
		config.WriteString(key + ":" + value + "\n")
	}
	return config.String()
	// return fmt.Sprintf("method: %s, path: %s, handlerParmas: %s, handlerReturn: %s", c.Method, c.Path, c.HandlerParams, c.HandlerReturn)
}

func analyzeAPICall(callExpr *ast.CallExpr, fset *token.FileSet) APICall {
	call := APICall{
		config: make(map[string]string, 0),
	}

	// Get the method name (e.g., "Get", "Post", etc.)
	if sel, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
		call.Method = sel.Sel.Name
	}

	if len(callExpr.Args) >= 2 {
		// First argument should be PACKAGE_NAME.Config
		if configComp, ok := callExpr.Args[0].(*ast.CompositeLit); ok {
			for _, elt := range configComp.Elts {
				if kv, ok := elt.(*ast.KeyValueExpr); ok {
					if key, ok := kv.Key.(*ast.Ident); ok {
						if basic, ok := kv.Value.(*ast.BasicLit); ok {
							call.config[key.Name] = basic.Value
						} else {
							print("nah")
						}
					}
				}
			}
		}

		// Second argument should be the handler function
		if handlerFunc, ok := callExpr.Args[1].(*ast.FuncLit); ok {
			params := make([]string, len(handlerFunc.Type.Params.List))
			for i, param := range handlerFunc.Type.Params.List {
				params[i] = fmt.Sprintf("%s %s", param.Names[0], exprToString(param.Type, fset))
			}
			call.HandlerParams = strings.Join(params, ", ")

			if handlerFunc.Type.Results != nil {
				results := make([]string, len(handlerFunc.Type.Results.List))
				for i, result := range handlerFunc.Type.Results.List {
					results[i] = exprToString(result.Type, fset)
				}
				call.HandlerReturn = strings.Join(results, ", ")
			}
		}
	}

	return call
}

func analyzeStruct(pkg *packages.Package, typeSpec *ast.TypeSpec) {
	obj := pkg.TypesInfo.Defs[typeSpec.Name]
	if obj == nil {
		return
	}

	structType, ok := obj.Type().Underlying().(*types.Struct)
	if !ok {
		return
	}

	for i := 0; i < structType.NumFields(); i++ {
		field := structType.Field(i)
		fmt.Printf("  Field: %s, Type: %s\n", field.Name(), field.Type())
		// Analyze field tags, types, etc.
	}
}
func analyzeMethod(pkg *packages.Package, funcDecl *ast.FuncDecl) {
	obj := pkg.TypesInfo.Defs[funcDecl.Name]
	if obj == nil {
		return
	}

	fn, ok := obj.(*types.Func)
	if !ok {
		return
	}

	sig := fn.Type().(*types.Signature)
	fmt.Printf("  Receiver: %s\n", sig.Recv())
	fmt.Printf("  Params: %s\n", sig.Params())
	fmt.Printf("  Results: %s\n", sig.Results())

	fmt.Println("  Function content analysis:")

	// Create an ast.Inspector to traverse the function's AST
	inspector := inspector.New([]*ast.File{pkg.Syntax[0]})

	// Define a map to count occurrences of different statement types
	statementCounts := make(map[string]int)

	var apiVarName string

	// Use inspector to analyze the function's body
	inspector.Nodes([]ast.Node{&ast.FuncDecl{}}, func(n ast.Node, push bool) bool {
		if !push {
			return false
		}

		if fd, ok := n.(*ast.FuncDecl); ok && fd == funcDecl {
			ast.Inspect(fd.Body, func(n ast.Node) bool {
				switch node := n.(type) {
				case *ast.AssignStmt:
					assignText := getAssignmentStatementText(node, pkg.Fset)

					// Check if this is the api assignment
					if strings.Contains(assignText, PACKAGE_NAME+".API(") {
						// Extract the variable name
						if len(node.Lhs) > 0 {
							if ident, ok := node.Lhs[0].(*ast.Ident); ok {

								apiVarName = ident.Name
								// log.Println("apiVarName", apiVarName)
							}
						}
					}
					fmt.Println(assignText)
					statementCounts["assignments"]++
				case *ast.IfStmt:
					statementCounts["if statements"]++
				case *ast.ForStmt:
					statementCounts["for loops"]++
				case *ast.RangeStmt:
					statementCounts["range loops"]++
				case *ast.SwitchStmt:
					statementCounts["switch statements"]++
				case *ast.ReturnStmt:
					// fmt.Println(exprToString(node.Results[0], pkg.Fset))
					// fmt.Println(exprToString(node.Results[1], pkg.Fset))
					statementCounts["return statements"]++
				case *ast.CallExpr:
					if sel, ok := node.Fun.(*ast.SelectorExpr); ok {
						if ident, ok := sel.X.(*ast.Ident); ok && ident.Name == apiVarName {
							fmt.Println(analyzeAPICall(node, pkg.Fset).String())
						}
					}
					statementCounts["function calls"]++
				}
				return true
			})
			return false
		}
		return true
	})

	// Print the statement counts
	for stmtType, count := range statementCounts {
		fmt.Printf("    - %s: %d\n", stmtType, count)
	}

	// Analyze function complexity (you can adjust the threshold as needed)
	complexity := statementCounts["if statements"] + statementCounts["for loops"] +
		statementCounts["range loops"] + statementCounts["switch statements"]

	fmt.Printf("  Estimated complexity: %d\n", complexity)
	if complexity > 10 {
		fmt.Println("  Warning: This function may be too complex. Consider refactoring.")
	}

	// Check for error handling
	if statementCounts["if statements"] > 0 && statementCounts["return statements"] > 0 {
		fmt.Println("  Note: This function may contain error handling.")
	}

	fmt.Println() // Add a blank line for readability between function analyses
}

type FileInfo struct {
	Path        string
	PackageName string
}

func stringify(data any) (string, error) {
	prettyJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(prettyJSON), nil
}

func (fi *FileInfo) String() (string, error) {
	return stringify(fi)
}

func walkDirectory(rootDir string, rootFileName string) (map[string]File, error) {
	files := make(map[string]File, 0)

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %q: %v\n", path, err)
			return err
		}
		if filepath.Join(rootDir, rootFileName) == path {
			return nil
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			files[path] = File{
				Path:    path,
				Imports: make(map[string]Import),
			}
		}

		return nil
	})

	return files, err
}

func processGoFile(path string) (string, string, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return "", "", err
	}
	return file.Name.Name, file.GoVersion, nil
}

func findGoMod(dir string) (string, error) {
	var modPath string
	for {
		// Check if go.mod exists in the current directory
		modPath = filepath.Join(dir, "go.mod")
		if _, err := os.Stat(modPath); err == nil {
			return modPath, nil
		}

		// Move up to the parent directory
		parent := filepath.Dir(dir)

		// If we've reached the root directory, stop
		if parent == dir {
			return "", os.ErrNotExist
		}

		dir = parent
	}
}

// type Response struct {
// 	Model       any
// 	Status      int
// 	ContentType string
// }

// type Method struct {
// 	Path        string
// 	Summary     string
// 	Description string
// 	Tags        []string
// }

// type API struct {
// 	Name   string
// 	Alias  string
// 	Path   string
// 	Method []Method
// }

//	type Package struct {
//		name  string
//		alias string
//		path  string
//	}
type Function[Params interface{}, ReturnType interface{}] struct {
	Name       string               `json:"name"`
	Params     Params               `json:"params"`
	ReturnType ReturnType           `json:"returnType"`
	Assignment []string             `json:"assignment"`
	Closure    []Function[any, any] `json:"closure"`
}

type Import struct {
	Path  string   `json:"path"`
	Name  string   `json:"name"`
	Alias []string `json:"alias"`
}

type File struct {
	Path      string             `json:"path"`
	Functions Function[any, any] `json:"functions"`
	Imports   map[string]Import  `json:"imports"`
}

type Package struct {
	ID        string              `json:"id"`
	Dir       string              `json:"dir"`
	Name      string              `json:"name"`
	GoVersion string              `json:"goVersion"`
	Files     map[string]File     `json:"files"`
	Alias     map[string]struct{} `json:"alias"`
	Imports   map[string]string   `json:"imports"`
	Functions Function[any, any]  `json:"omit"`
	Structs   *Structs            `json:"structs"`
}

type Project struct {
	Packages map[string]Package
	Dir      string
}

func (f *File) String() (string, error) {
	prettyJSON, err := json.MarshalIndent(f, "", "    ")
	if err != nil {
		return "", err
	}
	return string(prettyJSON), nil
}

func (s *Package) String() (string, error) {
	prettyJSON, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		return "", err
	}
	return string(prettyJSON), nil
}

func (s *Project) String() (string, error) {
	prettyJSON, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		return "", err
	}
	return string(prettyJSON), nil
}

const PACKAGE_NAME = "goblin"

// I want a lookup of files
// why? so I can traverse them to pick out all structs and calls to go-blin
// why? so that I can parse go-blin function calls and relfect the users structs as return types
// why? so that I can use them to make a swagger spec

// when the user does go generate
// we should parse the file and any of their files they import that are also theirs

// I will have to do two passes, once to get the calls to goblin and the data they return
// another to find all the types
// it's either that or I build a tree of the program as well and I don't want to... yet

// what should the tree look like?...
// well it should help me do the following:
// look up structs
// look up go-blin functions
// so do I need a tree?... no.
// how can I find go-blin calls?
// from the functoin signature of api.Get... but the user can change "api"
// so I need to find the definition of "api"
// what defines api?
// goblin.API... but the user can change "goblin"
// what defines goblin
// the import
// what if the api isn't in this file?
// we have to parse the users code building a doubly linked list for traversing

// how will I traverse?
// idea:
// go through each client file, create a list of ours on their files
// include path + aliases

// type Idea struct {
// 	files   map[string]File
// 	imports []File
// 	api     string
// 	calls   []string
// 	structs []string
// }

// Helper function to convert an expression to a string
func exprToString(expr ast.Expr, fset *token.FileSet) string {
	var buf bytes.Buffer
	err := printer.Fprint(&buf, fset, expr)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	return buf.String()
}

// Helper function to get the text of an assignment statement
func getAssignmentStatementText(assignStmt *ast.AssignStmt, fset *token.FileSet) string {
	var assignText strings.Builder

	// Convert left-hand side (LHS) to string
	for i, lhs := range assignStmt.Lhs {
		if i > 0 {
			assignText.WriteString(", ")
		}
		assignText.WriteString(exprToString(lhs, fset))
	}

	// Add assignment operator
	assignText.WriteString(" ")
	assignText.WriteString(assignStmt.Tok.String())
	assignText.WriteString(" ")

	// Convert right-hand side (RHS) to string
	for i, rhs := range assignStmt.Rhs {
		if i > 0 {
			assignText.WriteString(", ")
		}
		assignText.WriteString(exprToString(rhs, fset))
	}

	return assignText.String()
}

type Owner string

const (
	Ours   Owner = "Ours"
	Theirs Owner = "Theirs"
)

type ChildStructInfo struct {
	name         string `json:"name"`
	isNamedField bool   `json:"isNamedField"`
	owner        Owner  `json:"owner"`
	fieldName    string `json:"fieldName"`
	fieldType    string `json:"fieldType"`
}

type field struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type StructInfo struct {
	Index            int          `json:"index"`
	Name             string       `json:"name"`
	Fields           []*types.Var `json:"fields"`
	Owner            Owner        `json:"owner"`
	IsContainingOurs bool         `json:"isContainingOurs"`
}

func NewStructInfo(name string, fields []*types.Var, owner Owner, isContainerOurs bool) *StructInfo {
	return &StructInfo{
		Index:            -1,
		Name:             name,
		Fields:           fields,
		Owner:            owner,
		IsContainingOurs: isContainerOurs,
	}
}

func (si *StructInfo) setIndex(value int) {
	si.Index = value
}

type Node struct {
	Parents  []int `json:"parents"`
	Children []int `json:"children"`
}

func NewNode(owner Owner) *Node {
	return &Node{
		Parents:  make([]int, 0),
		Children: make([]int, 0),
	}
}
func (n *Node) AddParent(index int) {
	n.Parents = append(n.Parents, index)
}
func (n *Node) AddChild(index int) {
	n.Children = append(n.Children, index)
}

type Structs struct {
	Network    map[int]*Node
	Data       []*StructInfo
	NameToData map[string]*StructInfo
}

func NewStructs() *Structs {
	return &Structs{
		Network:    make(map[int]*Node, 0),
		Data:       make([]*StructInfo, 0),
		NameToData: make(map[string]*StructInfo),
	}
}

func (s *Structs) SetParentAsIsContainingOurs(node *Node) {
	for _, parentIndex := range node.Parents {
		parentData := s.Data[parentIndex]
		if parentData == nil {
			continue
		}
		parentData.IsContainingOurs = true

		parentNode := s.Network[parentIndex]
		if parentNode == nil {
			continue
		}
		if len(parentNode.Parents) > 0 {
			s.SetParentAsIsContainingOurs(parentNode)
		}
	}
}

func (s *Structs) FinaliseOwnership() {
	for index, node := range s.Network {
		data := s.Data[index]
		if data == nil || data.Owner == Ours {
			continue
		}
		for _, childIndex := range node.Children {
			childData := s.Data[childIndex]
			if childData == nil {
				continue
			}
			if childData.Owner == Ours || childData.IsContainingOurs {
				data.IsContainingOurs = true
				break
			}
		}
		if data.IsContainingOurs {
			s.SetParentAsIsContainingOurs(node)
		}
	}
}

func (s *Structs) addNewNode(si *StructInfo) int {
	s.NameToData[si.Name] = si

	s.Data = append(s.Data, si)
	index := len(s.Data) - 1

	si.setIndex(index)

	s.Network[index] = NewNode(Theirs)

	return index
}

func (s *Structs) String() (string, error) {
	prettyJSON, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		return "", err
	}
	return string(prettyJSON), nil
}

func analyzePackage(project *Project, dir string, config *packages.Config) {
	pkgs, err := packages.Load(config, dir)
	if err != nil {
		slog.Error(
			"unable to load packages",
			"error", err,
		)
		return
	}

	if len(pkgs) == 0 {
		slog.Error("No packages found")
		return
	}

	for _, pkg := range pkgs {
		currentPackage, found := project.Packages[pkg.ID]
		if !found {
			project.Packages[pkg.ID] = Package{
				ID:        pkg.ID,
				Name:      pkg.Name,
				Dir:       dir,
				GoVersion: "",
				Files:     make(map[string]File),
				Alias:     map[string]struct{}{},
				Imports:   make(map[string]string),
				Structs:   NewStructs(),
			}
			currentPackage = project.Packages[pkg.ID]
		}

		for _, file := range pkg.Syntax {
			if currentPackage.GoVersion == "" {
				currentPackage.GoVersion = file.GoVersion
			}

			currentFilePos := pkg.Fset.Position(file.Pos())
			currentFileAbsolutePath, err := filepath.Abs(currentFilePos.Filename)
			if err != nil {
				slog.Error(
					"unable to getting absolute path",
					"filename", currentFilePos.Filename,
					"error", err,
				)
				continue
			}

			currentFile, ok := currentPackage.Files[currentFileAbsolutePath]
			if !ok {
				currentPackage.Files[currentFileAbsolutePath] = File{
					Path:    currentFileAbsolutePath,
					Imports: make(map[string]Import),
				}
				currentFile = currentPackage.Files[currentFileAbsolutePath]
			}

			for _, currentImport := range file.Imports {
				// We only care about our packages, don't be nosey!
				if !strings.Contains(currentImport.Path.Value, PACKAGE_NAME) {
					continue
				}

				if currentImport.Path == nil ||
					currentImport.Path.ValuePos <= file.FileStart ||
					currentImport.Path.ValuePos >= file.FileEnd {
					slog.Error(
						"package out of bounds",
					)
					continue
				}

				currentImportPath := strings.Trim(currentImport.Path.Value, "\"")
				currentImportAbsolutePath, err := filepath.Abs(currentImportPath)
				if err != nil {
					slog.Error(
						"unable to parse absolute import path",
						"error", err,
					)
					continue
				}

				parts := strings.Split(currentImportPath, "/")
				name := parts[len(parts)-1]
				entry, ok := currentFile.Imports[currentImportAbsolutePath]
				if !ok {
					currentFile.Imports[currentImportAbsolutePath] = Import{
						Path:  currentImportAbsolutePath,
						Name:  name,
						Alias: make([]string, 0),
					}
					entry = currentFile.Imports[currentImportAbsolutePath]
					currentPackage.Imports[name] = name
				}

				if currentImport.Name != nil && currentImport.Name.Name != "" {
					entry.Alias = append(entry.Alias, currentImport.Name.Name)
					currentPackage.Imports[currentImport.Name.Name] = name
					currentFile.Imports[currentImportAbsolutePath] = entry
				}
			}

			// INFO: must parse structs before functions as functions maybe be on structs that extend our structs
			// INFO: becuase struct can extend structs we have to track all structs and the do a second pass :c
			ast.Inspect(file, func(n ast.Node) bool {
				typeSpec, isTypeSpec := n.(*ast.TypeSpec)
				if !isTypeSpec {
					return true
				}

				obj := pkg.TypesInfo.Defs[typeSpec.Name]
				if obj == nil {
					return true
				}

				currentStructType, isStruct := obj.Type().Underlying().(*types.Struct)
				if !isStruct {
					return true
				}

				currentStructInfo := currentPackage.Structs.NameToData[typeSpec.Name.Name]
				if currentStructInfo != nil && currentStructInfo.Fields != nil {
					slog.Error("This shouldn't be possible, but we are reparsing", "struct name:", typeSpec.Name.Name)
					return true
				}

				children := make([]ChildStructInfo, 0)
				fields := make([]*types.Var, 0)
				isContainingOurs := false
				for i := 0; i < currentStructType.NumFields(); i++ {
					field := currentStructType.Field(i)
					fields = append(fields, field)
					fieldType := field.Type()
					if fieldType == nil {
						continue
					}

					if _, isStruct := field.Type().Underlying().(*types.Struct); !isStruct {
						continue
					}

					owner := Theirs
					if strings.Contains(fieldType.String(), PACKAGE_NAME+".") ||
						strings.Contains(fieldType.String(), PACKAGE_NAME+"/") {
						owner = Ours
						isContainingOurs = true
					}

					parts := strings.Split(field.Type().String(), ".")
					// namespace
					childStructName := parts[1]

					name := field.Name()
					if childStructName != name && !isContainingOurs || owner == Ours {
						name = childStructName
					}

					children = append(children, ChildStructInfo{
						name:         name,
						isNamedField: childStructName != field.Name(),
						fieldName:    field.Name(),
						fieldType:    field.Type().String(),
						owner:        owner,
					})
				}

				if currentStructInfo == nil {
					// it's new!
					currentStructInfo = NewStructInfo(typeSpec.Name.Name, fields, Theirs, isContainingOurs)
					currentPackage.Structs.addNewNode(currentStructInfo)
				} else if currentStructInfo.Fields == nil {
					// it's someones child
					currentStructInfo.Fields = fields
					currentStructInfo.IsContainingOurs = isContainingOurs
				}

				currentNode := currentPackage.Structs.Network[currentStructInfo.Index]
				if currentNode == nil {
					return true
				}

				if isContainingOurs {
					currentStructInfo.IsContainingOurs = isContainingOurs
				}

				for _, child := range children {
					name := child.name
					if child.owner == Ours {
						name = PACKAGE_NAME + "." + child.name
					}
					childStructInfo := currentPackage.Structs.NameToData[name]
					if childStructInfo == nil {
						childStructInfo = NewStructInfo(name, nil, child.owner, false)
						currentPackage.Structs.addNewNode(childStructInfo)
					}
					childNode := currentPackage.Structs.Network[childStructInfo.Index]
					if childNode == nil {
						continue
					}
					childNode.AddParent(currentStructInfo.Index)
					currentNode.AddChild(childStructInfo.Index)
				}

				return true
			})
			currentPackage.Structs.FinaliseOwnership()

			continue

			// Create an ast.Inspector to traverse the function's AST
			inspector := inspector.New([]*ast.File{file})

			// Define a map to count occurrences of different statement types

			log.Println("Imports", currentPackage.Imports)
			// var apiVarName string
			inspector.Nodes([]ast.Node{&ast.FuncDecl{}}, func(node ast.Node, push bool) bool {
				if !push {
					return false
				}

				if currentFunction, ok := node.(*ast.FuncDecl); ok {
					obj := pkg.TypesInfo.Defs[currentFunction.Name]
					if obj == nil {
						return false
					}

					fn, ok := obj.(*types.Func)
					if !ok {
						return false
					}

					sig, ok := fn.Type().(*types.Signature)
					if !ok {
						return false
					}

					log.Println("Name", currentFunction.Name.Name)
					baseStruct := sig.Recv()
					if baseStruct != nil {
						fmt.Printf("  Receiver: %s\n", baseStruct)
						if strings.Contains(baseStruct.String(), PACKAGE_NAME+".") ||
							strings.Contains(baseStruct.String(), PACKAGE_NAME+"/") {
							varName := baseStruct.Type().Underlying().String()
							fmt.Printf("    %s: %s\n", baseStruct.Name(), varName)
						}
					}
					// input := sig.Recv()
					input := sig.Params()
					if input != nil {
						fmt.Printf("  Params: %s\n", input)
						if strings.Contains(input.String(), PACKAGE_NAME+".") ||
							strings.Contains(input.String(), PACKAGE_NAME+"/") {
							for i := 0; i < input.Len(); i++ {
								result := input.At(i)
								if result == nil {
									continue
								}
								varName := result.Type().Underlying().String()
								fmt.Printf("    %s: %s\n", result.Name(), varName)
							}
						}
					}
					fmt.Printf("  Results: %s\n", sig.Results())
					output := sig.Results()
					if output != nil {
						fmt.Printf("  Results: %s\n", output)
						if strings.Contains(output.String(), PACKAGE_NAME+".") ||
							strings.Contains(output.String(), PACKAGE_NAME+"/") {
							for i := 0; i < output.Len(); i++ {
								result := output.At(i)
								if result == nil {
									continue
								}
								varName := result.Type().Underlying().String()
								fmt.Printf("    %d: %s\n", i, varName)
							}
						}
					}

					fmt.Println("  Function content analysis:")
					ast.Inspect(currentFunction.Body, func(n ast.Node) bool {
						switch node := n.(type) {
						case *ast.AssignStmt:
							assignText := getAssignmentStatementText(node, pkg.Fset)
							log.Println("assignText", assignText)

							for importName := range currentFile.Imports {
								if strings.Contains(assignText, importName) {
									log.Println("PackageName", importName)
								}
							}
							// // Check if this is the api assignment
							// if strings.Contains(assignText, PACKAGE_NAME+".API(") {
							// 	// Extract the variable name
							// 	if len(node.Lhs) > 0 {
							// 		if ident, ok := node.Lhs[0].(*ast.Ident); ok {

							// 			// apiVarName = ident.Name
							// 			// log.Println("apiVarName", apiVarName)
							// 		}
							// 	}
							// }
							// fmt.Println(assignText)
						case *ast.IfStmt:
						case *ast.ForStmt:
						case *ast.RangeStmt:
						case *ast.SwitchStmt:
						case *ast.ReturnStmt:
							// fmt.Println(exprToString(node.Results[0], pkg.Fset))
							// fmt.Println(exprToString(node.Results[1], pkg.Fset))
						case *ast.CallExpr:
							// if sel, ok := node.Fun.(*ast.SelectorExpr); ok {
							// 	if ident, ok := sel.X.(*ast.Ident); ok && ident.Name == apiVarName {
							// 		fmt.Println(analyzeAPICall(node, pkg.Fset).String())
							// 	}
							// }
						}
						return true
					})
					return false
				}
				return true
			})

			// ast.Inspect(file, func(n ast.Node) bool {
			// 	if fn, ok := n.(*ast.FuncDecl); ok {
			// 		funcs = append(funcs, fn)
			// 	}
			// 	return true
			// })
			// for _, currentImport := range currentFile.Imports {
			// 	for _, alias := range currentImport.AliasIdent {
			// 		analyzeMethod(pkg, alias)
			// 		obj := pkg.TypesInfo.Defs[alias]
			// 		if obj == nil {
			// 			return
			// 		}

			// 	}

			// }
			// slog.Info("look what we found:", "currentFile", currentFile)
		}

		// slog.Info("look what we found:", "defs", pkg.TypesInfo.Defs)

	}
}

func main() {
	callerFile := os.Getenv("GOFILE")
	if callerFile == "" {
		slog.Error("GOFILE environment variable not set")
		return
	}

	// Get the absolute path of the caller
	path, err := filepath.Abs(callerFile)
	if err != nil {
		slog.Error(
			"Failed to get absolute path:",
			"error", err,
		)
		return
	}

	// Get the dir of the caller
	dir := filepath.Dir(path)

	// Traverse upwards from the call to find the go mod path
	goModPath, err := findGoMod(dir)
	if err != nil {
		if os.IsNotExist(err) {
			slog.Error("No go.mod file found in parent directories")
		} else {
			slog.Error(
				"Error:",
				"error", err,
			)
		}
		return
	}

	project := Project{
		Dir:      filepath.Dir(goModPath),
		Packages: make(map[string]Package),
	}

	analyzePackage(&project, dir, &packages.Config{
		Mode: packages.NeedFiles | packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedImports,
		Dir:  dir,
	})

	for _, pkg := range project.Packages {
		log.Println(pkg.Structs.String())
		// for _, s := range pkg.Structs.Data {
		// 	log.Println()
		// 	log.Println(s.Name)
		// 	for _, field := range s.Fields {
		// 		log.Println("name", field.Name(), "type", field.Type().String())
		// 	}
		// }
	}
}
