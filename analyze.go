package prensentationcheck

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "presentationcheck",
	Doc:  "presentation層のルールチェック",
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	inspect.Preorder([]ast.Node{
		(*ast.FuncDecl)(nil),
	}, func(n ast.Node) {
		fn := n.(*ast.FuncDecl)

		ComplexityCheck(pass, fn)
		OccurrenceCheck(pass, fn)
	})

	return nil, nil
}
