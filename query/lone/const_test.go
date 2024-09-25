package lone_test

import (
	"testing"

	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/logop"

	"github.com/DiegoEmilio01/goexpfbdd/query/internal/test"
	"github.com/DiegoEmilio01/goexpfbdd/query/lone"
)

func runLoneConst(t *testing.T, id int, tc test.BTRecord, neg bool) {
	tree, _ := test.NewMockTree(tc.Dim, nil)
	ctx := query.BasicQContext(tree)

	var f test.Encodable = lone.Const{
		query.QConst{Val: tc.Val1},
	}
	if neg {
		f = logop.Not{Q: f}
	}

	test.EncodeAndRun(t, f, ctx, id, tc.ExpCode)
}

func runGuardedLoneConst(
	t *testing.T,
	id int,
	tc test.BTRecord,
	neg bool,
) {
	tree, _ := test.NewMockTree(tc.Dim, nil)
	ctx := query.BasicQContext(tree)

	x := query.QConst{ID: "x"}

	ctx.AddScope("x")
	_ = ctx.SetScope(1, tc.Val1)

	var f test.Encodable = lone.Const{x}
	if neg {
		f = logop.Not{Q: f}
	}

	test.EncodeAndRun(t, f, ctx, id, tc.ExpCode)
}

func TestConst_Encoding(t *testing.T) {
	for i, tc := range LonePTT {
		t.Run(tc.Name, func(t *testing.T) {
			runLoneConst(t, i, tc, false)
		})
	}
}

func TestConst_Encoding_Guarded(t *testing.T) {
	for i, tc := range LonePTT {
		t.Run(tc.Name, func(t *testing.T) {
			runGuardedLoneConst(t, i, tc, false)
		})
	}
}

func TestNotConst_Encoding(t *testing.T) {
	for i, tc := range LoneNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runLoneConst(t, i, tc, true)
		})
	}
}

func TestNotConst_Encoding_Guarded(t *testing.T) {
	for i, tc := range LoneNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runGuardedLoneConst(t, i, tc, true)
		})
	}
}

func TestConst_Encoding_WrongDim(t *testing.T) {
	tree, _ := test.NewMockTree(4, nil)
	ctx := query.BasicQContext(tree)

	x := query.QConst{Val: []query.FeatV{query.BOT, query.BOT, query.BOT}}

	f := lone.Const{x}
	_, err := f.Encoding(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConst_Encoding_NilCtx(t *testing.T) {
	x := query.QConst{Val: []query.FeatV{query.BOT}}

	f := lone.Const{x}
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
