package aster

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"

	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"go/format"

	"bytes"
)

func Fix(codePath string) (string, string) {

	code, _ := os.ReadFile(codePath)

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", code, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	// Ensure imports are correct
	ensureFmtImported(file)
	// Check for anything that shouldn't be there (advanced)

	// Change Function name
	file = change_name(file)
	// Change Return Value Name
	goodCode, tmpDir := print_file(file, fset)

	// defer os.RemoveAll(tmpDir)
	fmt.Println(goodCode)
	return goodCode, tmpDir
}

func ensureFmtImported(file *ast.File) {
	// Check if "fmt" is already imported
	for _, imp := range file.Imports {
		if imp.Path.Value == "\"fmt\"" {
			// "fmt" is already imported
			return
		}
	}

	// Create an ImportSpec for "fmt"
	fmtImport := &ast.ImportSpec{
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: "\"fmt\"",
		},
	}

	// Check if there is an existing import declaration
	var importDecl *ast.GenDecl
	for _, decl := range file.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.IMPORT {
			importDecl = genDecl
			break
		}
	}

	if importDecl != nil {
		// Append to existing import declaration
		importDecl.Specs = append(importDecl.Specs, fmtImport)
	} else {
		// Create a new import declaration and insert it into file.Decls
		importDecl = &ast.GenDecl{
			Tok:   token.IMPORT,
			Specs: []ast.Spec{fmtImport},
		}
		// Insert the import declaration after the package declaration
		file.Decls = append([]ast.Decl{importDecl}, file.Decls...)
	}

	// Update the file.Imports slice
	file.Imports = append(file.Imports, fmtImport)
}

func not() {
	src := `package main

    func add(a int, b int) int {
        return a + b
    }

    func main() {
        result := add(2, 3)
        fmt.Println("Res: ", result)
    }`

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	ensureFmtImported(file)

	// _ , tmpDir := print_file(file, fset)
	// os.RemoveAll(tmpDir)

	ast.Inspect(file, func(n ast.Node) bool {
		fn, ok := n.(*ast.FuncDecl)
		if !ok {
			return true
		}

		// Create the print statement
		printStmt := &ast.ExprStmt{
			X: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   ast.NewIdent("fmt"),
					Sel: ast.NewIdent("Println"),
				},
				Args: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.STRING,
						Value: fmt.Sprintf("\"Entering function %s\"", fn.Name.Name),
					},
				},
			},
		}
		// Insert at the beginning of the function body
		fn.Body.List = append([]ast.Stmt{printStmt}, fn.Body.List...)
		return true
	})

	fmt.Println("Executing Code")
	srcFile, tmpDir := print_file(file, fset)

	cmd := exec.Command("go", "run", srcFile)

	stdoutStderr, err := cmd.CombinedOutput()

	fmt.Printf("OUT: %s\n", stdoutStderr)
	fmt.Println("Finished Running")
	os.RemoveAll(tmpDir)

}

func print_file(file *ast.File, fset *token.FileSet) (string, string) {

	var buf bytes.Buffer
	err := format.Node(&buf, fset, file)
	if err != nil {
		panic(err)
	}
	generatedSrc := buf.String()

	tmpDir, _ := ioutil.TempDir("", "go-run")

	// defer os.RemoveAll(tmpDir)

	srcFile := filepath.Join(tmpDir, "main.go")
	ioutil.WriteFile(srcFile, []byte(generatedSrc), 0644)
	fmt.Println("SRC:", generatedSrc)
	return srcFile, tmpDir
}

func change_name_2(file *ast.File) *ast.File {
	// look for a function name.
	// replace with name of function that we want defined.
	// TODO: handle if there are multiple function called within submitted code

	ast.Inspect(file, func(n ast.Node) bool {
		if ident, ok := n.(*ast.Ident); ok && ident.Name == "Random" {
			ident.Name = "UsersFunction"
		}
		return true
	})
	return file
}

func change_name(file *ast.File) *ast.File {
	// look for a function name.
	// replace with name of function that we want defined.
	// TODO: handle if there are multiple function called within submitted code
    count := 0
	ast.Inspect(file, func(n ast.Node) bool {
        if count >= 2 {
            fmt.Println("More than Two Functions")
            return false
        }
		if fn, ok := n.(*ast.FuncDecl); ok && fn.Name.Name != "main" {
			fn.Name.Name = "UsersFunction"
            count +=1
		}
		return true
	})
	return file
}
