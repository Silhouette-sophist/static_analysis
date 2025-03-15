package ast_visitor

import "go/ast"

type GoFunc struct {
	Id   int32
	Name string
	Pkg  string
}

type FuncVisitor struct {
	Funcs []GoFunc
}

func (v *FuncVisitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.FuncDecl:
		goFunc := GoFunc{
			Name: n.Name.Name,
		}
		v.Funcs = append(v.Funcs, goFunc)
	}
	return v
}
