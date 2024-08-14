package flevel

import (
	"errors"
	"fmt"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// ForAllGuardedL represents a FOR ALL guarded quantifier.
type ForAllGuardedL struct {
	// I corresponds to the constant that will be used to materialize the
	// instances. Its ID will be used for scope setting.
	I query.QConst
	// Q corresponds to a sub-query that implements the LogOpQ interface and
	// that is expected to make use of I.
	Q LogOpQ
}

// instanceIter allows to iterate over first level instances representing them
// as a slice of features stored in t.
type instanceIter struct {
	t    *[]query.FeatV
	fd   int
	cp   int
	jump bool
}

// newInstanceIter initializes a instanceIter and returns it alongside
// a pointer to the slice where it will write values to.
func newInstanceIter(dim int) (*[]query.FeatV, instanceIter, error) {
	iter := instanceIter{}

	t := make([]query.FeatV, dim)
	iter.t = &t
	iter.fd = dim
	iter.cp = 0
	iter.jump = true

	return &t, iter, nil
}

// Next attempts to moves the iterator to the next instance, updating the value of
// t. Returns true if there is a next instance and false otherwise.
func (i *instanceIter) Next() bool {
	if i.cp < i.fd {

		if i.jump {
			if i.cp >= 1 {
				(*i.t)[i.cp-1] = query.BOT
			}
			(*i.t)[i.cp] = query.ZERO
			i.jump = false
			return true
		}

		(*i.t)[i.cp] = query.ONE
		i.cp += 1
		i.jump = true
		return true
	}

	return false
}

// Encoding returns the CNF formula equivalent to the conjunction all the
// possible CNF formulas of f.Q resulting from instantiating every value of f.I
// in the ctx's model.
func (f ForAllGuardedL) Encoding(ctx query.QContext) (ncnf cnf.CNF, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("ForAllGuardedL: %w", err)
		}
	}()

	if f.Q == nil {
		return cnf.CNF{}, errors.New("Invalid encoding of nil child")
	}
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	ncnf, err = f.buildEncoding(ctx)

	return ncnf, err
}

func (f ForAllGuardedL) buildEncoding(
	ctx query.QContext,
) (ncnf cnf.CNF, err error) {
	var cv *[]query.FeatV
	var iter instanceIter
	var icnf cnf.CNF

	ctx.AddScope(f.I.ID)
	defer func() {
		if perr := ctx.PopScope(); perr != nil {
			err = errors.Join(err, perr)
		}
	}()

	i := 0
	cv, iter, err = newInstanceIter(ctx.Dim())
	if err != nil {
		return ncnf, err
	}

	for iter.Next() {
		if err = ctx.SetScope(i, *cv); err != nil {
			return cnf.CNF{}, err
		}

		icnf, err = f.Q.Encoding(ctx)
		if err != nil {
			return cnf.CNF{}, err
		}

		ncnf = ncnf.Conjunction(icnf)

		i += 1
	}

	return ncnf, err
}
