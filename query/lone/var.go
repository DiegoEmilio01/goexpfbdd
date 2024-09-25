package lone

import (
	"errors"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// Var is the variable-variable version of the LU formula.
type Var struct {
	I query.QVar
	// CountVarGen returns a variable generated from v that will be used to
	// encode the amount of features equal to bot in v.
	CountVarGen func(v query.QVar) query.QVar
}

// Encoding returns a CNF that is true if and only if the query variable l.I
// has only 1 defined feature.
func (l Var) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}
	if l.CountVarGen == nil {
		return cnf.CNF{}, errors.New("Invalid nil var generation function")
	}

	sv := ctx.ScopeVar(l.I)
	svCount := l.CountVarGen(sv)

	ncnf := cnf.CNF{}
	ncnf = ncnf.AppendConsistency(varBotCountClauses(sv, svCount, ctx)...)

	for i := 0; i < ctx.Dim()+1; i++ {
		if ctx.Dim()-1 != i {
			ncnf = ncnf.AppendSemantics(
				cnf.Clause{-ctx.CNFVar(svCount, ctx.Dim()-1, i)},
			)
		}
	}

	return ncnf, nil
}
