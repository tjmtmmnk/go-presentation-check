// Copyright 2020 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package prensentationcheck

import (
	"go/ast"
	"go/types"
)

var (
	errTyp = types.Universe.Lookup("error").Type().Underlying().(*types.Interface)
)

// Complexity calculates the cyclomatic complexity of a function.
// The 'fn' node is either a *ast.FuncDecl or a *ast.FuncLit.
func Complexity(fn ast.Node, typeInfo *types.Info) int {
	v := complexityVisitor{
		complexity: 1,
		typeInfo:   typeInfo,
	}
	ast.Walk(&v, fn)
	return v.complexity
}

type complexityVisitor struct {
	// complexity is the cyclomatic complexity
	complexity int
	typeInfo   *types.Info
}

// Visit implements the ast.Visitor interface.
func (v *complexityVisitor) Visit(n ast.Node) ast.Visitor {
	// エラーハンドリングなら対象外にする
	if ifs, ok := n.(*ast.IfStmt); ok {
		if cond, ok := ifs.Cond.(*ast.BinaryExpr); ok {
			t := v.typeInfo.TypeOf(cond.X)
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
	return v
}

func isErr(t types.Type) bool {
	return types.Implements(t, errTyp)
}
