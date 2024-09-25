package lup_test

import (
	"testing"

	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/logop"
	"github.com/jtcaraball/goexpdt/query/predicates/subsumption"

	"github.com/DiegoEmilio01/goexpfbdd/query/internal/test"
	"github.com/DiegoEmilio01/goexpfbdd/query/lup"
)

func runLupVarConst(t *testing.T, id int, tc test.BTRecord, neg bool) {
	tree, _ := test.NewMockTree(tc.Dim, nil)
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	c1 := query.QConst{Val: tc.Val1}
	c2 := query.QConst{Val: tc.Val2}

	var f test.Encodable = lup.VarConst{x, c2, test.VarGenBotCount}
	if neg {
		f = logop.Not{Q: f}
	}

	f = logop.WithVar{
		I: x,
		Q: logop.And{
			Q1: logop.And{
				Q1: subsumption.VarConst{I1: x, I2: c1},
				Q2: subsumption.ConstVar{I1: c1, I2: x},
			},
			Q2: f,
		},
	}

	test.EncodeAndRun(t, f, ctx, id, tc.ExpCode)
}

func runGuardedLupVarConst(t *testing.T, id int, tc test.BTRecord, neg bool) {
	tree, _ := test.NewMockTree(tc.Dim, nil)
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	y := query.QConst{ID: "y"}
	c1 := query.QConst{Val: tc.Val1}

	ctx.AddScope("y")
	_ = ctx.SetScope(1, tc.Val2)

	var f test.Encodable = lup.VarConst{x, y, test.VarGenBotCount}
	if neg {
		f = logop.Not{Q: f}
	}

	f = logop.WithVar{
		I: x,
		Q: logop.And{
			Q1: logop.And{
				Q1: subsumption.VarConst{I1: x, I2: c1},
				Q2: subsumption.ConstVar{I1: c1, I2: x},
			},
			Q2: f,
		},
	}

	test.EncodeAndRun(t, f, ctx, id, tc.ExpCode)
}

func TestVarConst_Encoding(t *testing.T) {
	for i, tc := range LupPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runLupVarConst(t, i, tc, false)
		})
	}
}

func TestVarConst_Encoding_Guarded(t *testing.T) {
	for i, tc := range LupPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runGuardedLupVarConst(t, i, tc, false)
		})
	}
}

func TestNotVarConst_Encoding(t *testing.T) {
	for i, tc := range LupNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runLupVarConst(t, i, tc, true)
		})
	}
}

func TestNotVarConst_Encoding_Guarded(t *testing.T) {
	for i, tc := range LupNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runGuardedLupVarConst(t, i, tc, true)
		})
	}
}

func TestVarConst_Encoding_WrongDim(t *testing.T) {
	tree, _ := test.NewMockTree(4, nil)
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	y := query.QConst{Val: []query.FeatV{query.BOT, query.BOT, query.BOT}}

	f := lup.VarConst{x, y, test.VarGenBotCount}

	_, err := f.Encoding(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestVarConst_Encoding_NilCtx(t *testing.T) {
	x := query.QVar("x")
	y := query.QConst{Val: []query.FeatV{query.BOT}}

	f := lup.VarConst{x, y, test.VarGenBotCount}
	e := "Invalid encoding with nil ctx"

	_, err := f.Encoding(nil)
	if err == nil {
		t.Error("Nil context encoding error not caught.")
	} else if err.Error() != e {
		t.Errorf(
			"Incorrect error for nil context encoding. Expected %s but got %s",
			e,
			err.Error(),
		)
	}
}
