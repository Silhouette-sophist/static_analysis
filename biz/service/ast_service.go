package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.sophist.static_analysis/biz/ast_visitor"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func AstParseFile(ctx context.Context, path string) (any, error) {
	readBytes, err := os.ReadFile(path)
	if err != nil {
		hlog.CtxWarnf(ctx, "AstParseFile ReadFile err %v", err)
		return nil, err
	}
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, readBytes, parser.ParseComments)
	if err != nil {
		hlog.CtxWarnf(ctx, "AstParseFile ParseFile err %v", err)
		return nil, err
	}
	visitor := &ast_visitor.FuncVisitor{}
	ast.Walk(visitor, file)
	return visitor.Funcs, nil
}
