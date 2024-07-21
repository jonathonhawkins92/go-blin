package source

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"runtime"
	"strings"

	"golang.org/x/tools/go/packages"
)

// Node represents a node in our graph
type Node struct {
	ASTNode ast.Node
	Edges   []*Node
	File    string // Added to keep track of which file this node is from
}

// Graph represents our AST graph
type Graph struct {
	Root  *Node
	Files map[string]*ast.File // Keep track of all parsed files
}

// CreateGraph creates a graph from an AST
func CreateGraph(node ast.Node, filename string) *Graph {
	graph := &Graph{
		Root:  &Node{ASTNode: node, File: filename},
		Files: make(map[string]*ast.File),
	}
	buildGraph(graph.Root, node)
	return graph
}

// buildGraph recursively builds the graph
func buildGraph(graphNode *Node, astNode ast.Node) {
	ast.Inspect(astNode, func(n ast.Node) bool {
		if n == nil || n == astNode {
			return true
		}
		child := &Node{ASTNode: n, File: graphNode.File}
		graphNode.Edges = append(graphNode.Edges, child)
		buildGraph(child, n)
		return false
	})
}

// Search performs a depth-first search on the graph
func (g *Graph) Search(predicate func(ast.Node) bool) []ast.Node {
	var results []ast.Node
	searchDFS(g.Root, predicate, &results)
	return results
}

// searchDFS performs a depth-first search
func searchDFS(node *Node, predicate func(ast.Node) bool, results *[]ast.Node) {
	if predicate(node.ASTNode) {
		*results = append(*results, node.ASTNode)
	}
	for _, child := range node.Edges {
		searchDFS(child, predicate, results)
	}
}

// ProjectAnalyzer handles analyzing an entire Go project
type ProjectAnalyzer struct {
	Graph      *Graph
	Fset       *token.FileSet
	PackageMap map[string]*packages.Package
}

// NewProjectAnalyzer creates a new ProjectAnalyzer
func NewProjectAnalyzer() *ProjectAnalyzer {
	return &ProjectAnalyzer{
		Graph:      &Graph{Files: make(map[string]*ast.File)},
		Fset:       token.NewFileSet(),
		PackageMap: make(map[string]*packages.Package),
	}
}

// AnalyzeProject analyzes an entire Go project starting from the given directory
func (pa *ProjectAnalyzer) AnalyzeProject(dir string) error {
	cfg := &packages.Config{
		Mode: packages.NeedFiles | packages.NeedSyntax | packages.NeedTypes | packages.NeedImports,
		Dir:  dir,
		Fset: pa.Fset,
	}

	pkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		return fmt.Errorf("error loading packages: %v", err)
	}

	for _, pkg := range pkgs {
		pa.PackageMap[pkg.PkgPath] = pkg
		for _, file := range pkg.Syntax {
			filename := pa.Fset.File(file.Pos()).Name()
			pa.Graph.Files[filename] = file
			if pa.Graph.Root == nil {
				pa.Graph.Root = &Node{ASTNode: file, File: filename}
			}
			buildGraph(&Node{ASTNode: file, File: filename}, file)
		}
	}

	return nil
}

// AnalyzeFile analyzes a single Go file and its imports
func (pa *ProjectAnalyzer) AnalyzeFile(filePath string) error {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return fmt.Errorf("error getting absolute path: %v", err)
	}

	// Parse the file
	file, err := parser.ParseFile(pa.Fset, absPath, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("error parsing file: %v", err)
	}

	// Add the file to our graph
	pa.Graph.Files[absPath] = file
	if pa.Graph.Root == nil {
		pa.Graph.Root = &Node{ASTNode: file, File: absPath}
	}
	buildGraph(&Node{ASTNode: file, File: absPath}, file)

	// Analyze imports
	for _, imp := range file.Imports {
		if imp.Path != nil {
			importPath := strings.Trim(imp.Path.Value, "\"")
			err = pa.analyzeImport(importPath)
			if err != nil {
				return fmt.Errorf("error analyzing import %s: %v", importPath, err)
			}
		}
	}

	return nil
}

// analyzeImport analyzes an imported package
func (pa *ProjectAnalyzer) analyzeImport(importPath string) error {
	if _, ok := pa.PackageMap[importPath]; ok {
		// We've already analyzed this package
		return nil
	}

	cfg := &packages.Config{
		Mode: packages.NeedFiles | packages.NeedSyntax | packages.NeedTypes | packages.NeedImports,
		Fset: pa.Fset,
	}

	pkgs, err := packages.Load(cfg, importPath)
	if err != nil {
		return fmt.Errorf("error loading package %s: %v", importPath, err)
	}

	for _, pkg := range pkgs {
		pa.PackageMap[pkg.PkgPath] = pkg
		for _, file := range pkg.Syntax {
			filename := pa.Fset.File(file.Pos()).Name()
			pa.Graph.Files[filename] = file
			buildGraph(&Node{ASTNode: file, File: filename}, file)
		}
	}

	return nil
}

// FindImports finds all imports in the project
func (pa *ProjectAnalyzer) FindImports() map[string][]string {
	imports := make(map[string][]string)

	for filename, file := range pa.Graph.Files {
		fileImports := []string{}
		for _, imp := range file.Imports {
			if imp.Path != nil {
				fileImports = append(fileImports, strings.Trim(imp.Path.Value, "\""))
			}
		}
		imports[filename] = fileImports
	}

	return imports
}

// FindFunctions finds all function declarations in the project
func (pa *ProjectAnalyzer) FindFunctions() []*ast.FuncDecl {
	var functions []*ast.FuncDecl

	for _, file := range pa.Graph.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			if funcDecl, ok := n.(*ast.FuncDecl); ok {
				functions = append(functions, funcDecl)
			}
			return true
		})
	}

	return functions
}

// FindStructs finds all struct declarations in the project
func (pa *ProjectAnalyzer) FindStructs() []*ast.TypeSpec {
	var structs []*ast.TypeSpec

	for _, file := range pa.Graph.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			if typeSpec, ok := n.(*ast.TypeSpec); ok {
				if _, isStruct := typeSpec.Type.(*ast.StructType); isStruct {
					structs = append(structs, typeSpec)
				}
			}
			return true
		})
	}

	return structs
}

func Example() {
	// Get the caller's information
	_, callerFile, _, ok := runtime.Caller(1)
	if !ok {
		fmt.Printf("Unable to get caller information")
		return
	}

	// Get the absolute path of the caller's file
	absPath, err := filepath.Abs(callerFile)
	if err != nil {
		fmt.Printf("Error abs path: %v\n", err)
		return
	}

	analyzer := NewProjectAnalyzer()
	err = analyzer.AnalyzeFile(absPath)
	if err != nil {
		fmt.Printf("Error analyzing project: %v\n", err)
		return
	}

	// imports := analyzer.FindImports()
	// fmt.Println("Imports:")
	// for file, fileImports := range imports {
	// 	fmt.Printf("  %s:\n", file)
	// 	for _, imp := range fileImports {
	// 		fmt.Printf("    %s\n", imp)
	// 	}
	// }

	// functions := analyzer.FindFunctions()
	// fmt.Printf("\nFound %d function declarations:\n", len(functions))
	// for _, funcDecl := range functions {
	// 	fmt.Printf("- %s (in %s)\n", funcDecl.Name.Name, analyzer.Fset.Position(funcDecl.Pos()).Filename)
	// }

	structs := analyzer.FindStructs()
	fmt.Printf("\nFound %d struct declarations:\n", len(structs))
	for _, structDecl := range structs {
		fmt.Printf("- %s (in %s)\n", structDecl.Name.Name, analyzer.Fset.Position(structDecl.Pos()).Filename)
	}
}
