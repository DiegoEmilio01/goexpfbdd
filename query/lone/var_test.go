package lone_test

import (
	"testing"

	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/logop"
	"github.com/jtcaraball/goexpdt/query/predicates/subsumption"

	"github.com/DiegoEmilio01/goexpfbdd/query/internal/test"
	"github.com/DiegoEmilio01/goexpfbdd/query/lone"
)

func runLoneVar(t *testing.T, id int, tc test.BTRecord, neg bool) {
	tree, _ := test.NewMockTree(tc.Dim, nil)
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	c1 := query.QConst{Val: tc.Val1}

	var f test.Encodable = lone.Var{x, test.VarGenBotCount}
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

func runGuardedLoneVar(t *testing.T, id int, tc test.BTRecord, neg bool) {
	tree, _ := test.NewMockTree(tc.Dim, nil)
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	c1 := query.QConst{Val: tc.Val1}

	//ctx.AddScope("x")
	//_ = ctx.SetScope(1, tc.Val1)

	var f test.Encodable = lone.Var{x, test.VarGenBotCount}
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

func TestVar_Encoding(t *testing.T) {
	for i, tc := range LonePTT {
		t.Run(tc.Name, func(t *testing.T) {
			runLoneVar(t, i, tc, false)
		})
	}
}

func TestVar_Encoding_Guarded(t *testing.T) {
	for i, tc := range LonePTT {
		t.Run(tc.Name, func(t *testing.T) {
			runGuardedLoneVar(t, i, tc, false)
		})
	}
}

func TestNotVar_Encoding(t *testing.T) {
	for i, tc := range LoneNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runLoneVar(t, i, tc, true)
		})
	}
}

func TestNotVar_Encoding_Guarded(t *testing.T) {
	for i, tc := range LoneNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runGuardedLoneVar(t, i, tc, true)
		})
	}
}

func TestVar_Encoding_NilCtx(t *testing.T) {
	x := query.QVar("x")

	f := lone.Var{x, test.VarGenBotCount}
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
