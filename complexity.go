package prensentationcheck

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

var (
	errTyp = types.Universe.Lookup("error").Type().Underlying().(*types.Interface)
)

func ComplexityCheck(pass *analysis.Pass, fn ast.Node) {
	v := complexityVisitor{
		complexity: 1,
		pass:       pass,
	}
	ast.Walk(&v, fn)
}

type complexityVisitor struct {
	complexity int
	pass       *analysis.Pass
}

// Visit implements the ast.Visitor interface.
func (v *complexityVisitor) Visit(n ast.Node) ast.Visitor {
	// エラーハンドリングなら対象外にする
	if ifs, ok := n.(*ast.IfStmt); ok {
		if cond, ok := ifs.Cond.(*ast.BinaryExpr); ok {
			t := v.pass.TypesInfo.TypeOf(cond.X)
			if isErr(t) {
				return v
			}
		}
	}

	switch n := n.(type) {
	case *ast.IfStmt:
		v.complexity++
	case *ast.CaseClause:
		if n.List != nil { // ignore default case
			v.complexity++
		}
	case *ast.CommClause:
		if n.Comm != nil { // ignore default case
			v.complexity++
		}
	}

	if isLogic(n) {
		if v.complexity > 1 {
			v.pass.Reportf(n.Pos(), "ロジックを書いてはいけません")
		}
	}

	return v
}

func isLogic(n ast.Node) bool {
	switch n.(type) {
	case *ast.IfStmt, *ast.CaseClause, *ast.CommClause:
		return true
	}
	return false
}

func isErr(t types.Type) bool {
	return types.Implements(t, errTyp)
}
