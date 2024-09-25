package lone

import (
	"errors"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// Const is the constant version of the L1 formula.
type Const struct {
	I query.QConst
}

// Encoding returns a CNF that is true if and only if the query constant l.I
// has only 1 defined feature.
func (l Const) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	sc, _ := ctx.ScopeConst(l.I)

	if err := query.ValidateConstsDim(
		ctx.Dim(),
		sc,
	); err != nil {
		return cnf.CNF{}, err
	}

	if sc.BotCount() == ctx.Dim()-1 {
		return cnf.TrueCNF, nil
	}

	return cnf.FalseCNF, nil
}
