package service

import (
	"context"
	"testing"
)

func TestAstParseFile(t *testing.T) {
	path := "/Users/silhouette/code-analysis/static_analysis/biz/service/github_service.go"
	file, err := AstParseFile(context.Background(), path)
	if err != nil {
		t.Errorf("AstParseFile err %v", err)
		return
	}
	t.Logf("AstParseFile sucess %v", file)
}
