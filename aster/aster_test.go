package aster

import (
    "go/ast"
    "go/parser"
    "go/token"
    "log"
)

src := `
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
`

fset := token.NewFileSet()
file, err := parser.ParseFile(fset, "", src, parser.ParseComments)
if err != nil {
    log.Fatal(err)
}


type myVisitor struct{}

func (v *myVisitor) Visit(n ast.Node) ast.Visitor {
    if n != nil {
        // Process the node n
    }
    return v
}

// Usage
v := &myVisitor{}
ast.Walk(v, file)


ast.Inspect(file, func(n ast.Node) bool {
    switch x := n.(type) {
    case *ast.FuncDecl:
        // Function declaration
    case *ast.ImportSpec:
        // Import statement
    }
    return true
})

// collecting function names 
var funcs []string

ast.Inspect(file, func(n ast.Node) bool {
    if fn, ok := n.(*ast.FuncDecl); ok {
        funcs = append(funcs, fn.Name.Name)
    }
    return true
})

fmt.Println("Functions:", funcs)

// Printing all functions calls 

// ast.Inspect(file, func(n ast.Node) bool {
//     if call, ok := n.(*ast.CallExpr); ok {
//         switch fun := call.Fun.(type) {
//         case *ast.Ident:
//             fmt.Println("Function call:", fun.Name)
//         case *ast.SelectorExpr:
//             if x, ok := fun.X.(*ast.Ident); ok {
//                 fmt.Printf("Method call: %s.%s\n", x.Name, fun.Sel.Name)
//             }
//         }
//     }
//     return true
// })


// Change Name of functions: 

// ast.Inspect(file, func(n ast.Node) bool {
//     if ident, ok := n.(*ast.Ident); ok && ident.Name == "oldName" {
//         ident.Name = "newName"
//     }
//     return true
// })

