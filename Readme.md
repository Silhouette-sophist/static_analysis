## 说明

-  1.拉取仓库代码，并执行go mod tidy拉取所有依赖包
-  2.执行go list获取主包信息，再通过go list -test=false -deps=true -f '{{.Dir}}'获取入口main所依赖的包
-  3.ssa工具集进行程序分析调用
  - 加载所有包信息，packages.Load(cfg)，得到[]*packages.Package
  - 判断依赖是否有问题，packages.Visit(pkgs, nil, func(pkg *packages.Package){if len(pkg.Errors) > 0})
  - 基于关注包构建ssa程序， ssautil.AllPackages(packages, 0)获取ssa.Program
  - 基于ssa构建ssaProgram[耗时]，ssaProgram.Build()
  - 获取包含main函数的包，ssautil.MainPackages(ssaPackages)
  - 判断主仓文件中使用到的函数标识符，基于ast visitor编译导出的Ident
    - pta分析，pointer.Config{Mains: mainPackages, BuildCallGraph: true}，pointer.Analyze(pointerCfg)，返回result.CallGraph即是
    - rta分析，rta.Analyze(roots, true)，注意roots是所有main包中的init和main方法做根节点，返回result.CallGraph即是
    - vta_pkg分析，
    - vta分析，ssautil.AllFunctions(prog)， vta.CallGraph(allFunctions, cha.CallGraph(prog))即是CallGraph