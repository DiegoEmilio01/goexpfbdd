package lup

import (
	"errors"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// VarVar is the variable-variable version of the LU formula.
type VarVar struct {
	I1 query.QVar
	I2 query.QVar
	// CountVarGen returns a variable generated from v that will be used to
	// encode the amount of features equal to bot in v.
	CountVarGen func(v query.QVar) query.QVar
}

// Encoding returns a CNF that is true if and only if the query variable l.I1
// has exactly one more BOT than query variable l.I2.
func (l VarVar) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}
	if l.CountVarGen == nil {
		return cnf.CNF{}, errors.New("Invalid nil var generation function")
	}

	sv1 := ctx.ScopeVar(l.I1)
	sv2 := ctx.ScopeVar(l.I2)
	svCount1 := l.CountVarGen(sv1)
	svCount2 := l.CountVarGen(sv2)

	ncnf := cnf.CNF{}

	ncnf = ncnf.AppendConsistency(varBotCountClauses(sv1, svCount1, ctx)...)
	ncnf = ncnf.AppendConsistency(varBotCountClauses(sv2, svCount2, ctx)...)

	// If we see a number of bots in I1 then we must see exactly one less on I2
	cl := cnf.Clause{-ctx.CNFVar(svCount1, ctx.Dim()-1, 0)} // when I1 is full, I2 cant be one level up
	ncnf = ncnf.AppendSemantics(cl)

	var i int
	for i = 1; i < ctx.Dim(); i++ {
		cl := cnf.Clause{
			-ctx.CNFVar(svCount1, ctx.Dim()-1, i),
			ctx.CNFVar(svCount2, ctx.Dim()-1, i-1),
		}
		ncnf = ncnf.AppendSemantics(cl)
	}

	return ncnf, nil
}
