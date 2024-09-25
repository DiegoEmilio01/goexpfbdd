package lup

import (
	"errors"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// ConstConst is the constant-constant version of the LU formula.
type ConstConst struct {
	I1 query.QConst
	I2 query.QConst
}

// Encoding returns a CNF that is true if and only if the query constant l.I1
// has exactly one more BOT than query constant l.I2.
func (l ConstConst) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	sc1, _ := ctx.ScopeConst(l.I1)
	sc2, _ := ctx.ScopeConst(l.I2)

	if err := query.ValidateConstsDim(
		ctx.Dim(),
		sc1,
		sc2,
	); err != nil {
		return cnf.CNF{}, err
	}

	if sc1.BotCount() == sc2.BotCount()+1 {
		return cnf.TrueCNF, nil
	}

	return cnf.FalseCNF, nil
}
