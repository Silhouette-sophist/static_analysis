package ssa_analysis

import (
	"golang.org/x/tools/go/ssa"
)

type BuildCustomGraphParam struct {
	CodeDir   string
	Algo      string
	RootPkg   string
	ModeDiffs []ModeDiff
}

type ModeDiff struct {
	Pkg        string
	NewVersion string
	OldVersion string
	DiffType   string
}

// 递归函数，用于打印调用链
func printCallChain(funcValue *ssa.Function, depth int) {
	indent := ""
	for i := 0; i < depth; i++ {
		indent += "  "
	}
	for _, block := range funcValue.Blocks {
		for _, instr := range block.Instrs {
			if call, ok := instr.(*ssa.Call); ok {
				if callee := call.Call.StaticCallee(); callee != nil {
					printCallChain(callee, depth+1)
				}
			}
		}
	}
}

func SsaAnalysis(param BuildCustomGraphParam) {

}
