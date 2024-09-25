package lup_test

import (
	"testing"

	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/logop"

	"github.com/DiegoEmilio01/goexpfbdd/query/internal/test"
	"github.com/DiegoEmilio01/goexpfbdd/query/lup"
)

func runLupConstConst(t *testing.T, id int, tc test.BTRecord, neg bool) {
	tree, _ := test.NewMockTree(tc.Dim, nil)
	ctx := query.BasicQContext(tree)

	var f test.Encodable = lup.ConstConst{
		query.QConst{Val: tc.Val1},
		query.QConst{Val: tc.Val2},
	}
	if neg {
		f = logop.Not{Q: f}
	}

	test.EncodeAndRun(t, f, ctx, id, tc.ExpCode)
}

func runGuardedLupConstConst(
	t *testing.T,
	id int,
	tc test.BTRecord,
	neg bool,
) {
	tree, _ := test.NewMockTree(tc.Dim, nil)
	ctx := query.BasicQContext(tree)

	x := query.QConst{ID: "x"}
	y := query.QConst{ID: "y"}

	ctx.AddScope("x")
	_ = ctx.SetScope(1, tc.Val1)
	ctx.AddScope("y")
	_ = ctx.SetScope(2, tc.Val2)

	var f test.Encodable = lup.ConstConst{x, y}
	if neg {
		f = logop.Not{Q: f}
	}

	test.EncodeAndRun(t, f, ctx, id, tc.ExpCode)
}

func TestConstConst_Encoding(t *testing.T) {
	for i, tc := range LupPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runLupConstConst(t, i, tc, false)
		})
	}
}

func TestConstConst_Encoding_Guarded(t *testing.T) {
	for i, tc := range LupPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runGuardedLupConstConst(t, i, tc, false)
		})
	}
}

func TestNotConstConst_Encoding(t *testing.T) {
	for i, tc := range LupNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runLupConstConst(t, i, tc, true)
		})
	}
}

func TestNotConstConst_Encoding_Guarded(t *testing.T) {
	for i, tc := range LupNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runGuardedLupConstConst(t, i, tc, true)
		})
	}
}

func TestConstConst_Encoding_WrongDim(t *testing.T) {
	tree, _ := test.NewMockTree(4, nil)
	ctx := query.BasicQContext(tree)

	x := query.QConst{Val: []query.FeatV{query.BOT, query.BOT, query.BOT}}
	y := query.QConst{Val: []query.FeatV{query.BOT, query.BOT, query.BOT}}

	f := lup.ConstConst{x, y}
	_, err := f.Encoding(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConstConst_Encoding_NilCtx(t *testing.T) {
	x := query.QConst{Val: []query.FeatV{query.BOT}}
	y := query.QConst{Val: []query.FeatV{query.BOT}}

	f := lup.ConstConst{x, y}
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
