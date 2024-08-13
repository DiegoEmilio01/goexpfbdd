// The flevel package defines the logical operator
// ForAllGuardedL which is defined to handle
// a particular variable declarations.
package flevel

import (
	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// A LogOpQ allows for the encoding of its meaning into a CNF formula.
type LogOpQ interface {
	// Encoding returns takes in a ctx representing the state of a query or
	// sub-query and returns its CNF encoding along side an error.
	Encoding(ctx query.QContext) (cnf.CNF, error)
}
