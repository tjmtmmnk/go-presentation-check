package prensentationcheck

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

type occurrenceVisitor struct {
	repositoryCount int
	serviceCount    int
	pass            *analysis.Pass
}

func OccurrenceCheck(pass *analysis.Pass, fn ast.Node) {
	v := occurrenceVisitor{
		repositoryCount: 0,
		serviceCount:    0,
		pass:            pass,
	}
	ast.Walk(&v, fn)
}

func (v *occurrenceVisitor) Visit(n ast.Node) ast.Visitor {
	if call, ok := n.(*ast.CallExpr); ok {
		if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
			if id, ok := sel.X.(*ast.Ident); ok {
				if id.Name == "repository" {
					v.repositoryCount++
					if v.repositoryCount > 1 {
						v.pass.Reportf(call.Pos(), "repositoryの呼び出しは1回までです。usecaseを作ってください。")
					}
				}
				if id.Name == "service" {
					v.serviceCount++
					if v.serviceCount > 0 {
						v.pass.Reportf(call.Pos(), "serviceの呼び出しは禁止です。usecaseを作ってください。")
					}
				}
			}
		}
	}
	return v
}
