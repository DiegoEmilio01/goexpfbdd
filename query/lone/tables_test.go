package lone_test

import (
	"github.com/DiegoEmilio01/goexpfbdd/query/internal/test"
	"github.com/jtcaraball/goexpdt/query"
)

var LonePTT = []test.BTRecord{
	{
		Dim:     3,
		Name:    "(1,_,_)",
		Val1:    []query.FeatV{query.ONE, query.BOT, query.BOT},
		ExpCode: 10, // successful
	},
	{
		Dim:     3,
		Name:    "(1,1,_)",
		Val1:    []query.FeatV{query.ONE, query.ONE, query.BOT},
		ExpCode: 20, // unsuccessful
	},
	{
		Dim:     3,
		Name:    "(1,1,0)",
		Val1:    []query.FeatV{query.ONE, query.ONE, query.ZERO},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(1,_,1)",
		Val1:    []query.FeatV{query.ONE, query.BOT, query.ONE},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(_,1,_)",
		Val1:    []query.FeatV{query.BOT, query.ONE, query.BOT},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(_,_,_)",
		Val1:    []query.FeatV{query.BOT, query.BOT, query.BOT},
		ExpCode: 20,
	},
}

var LoneNTT = []test.BTRecord{
	{
		Dim:     3,
		Name:    "(1,_,_)",
		Val1:    []query.FeatV{query.ONE, query.BOT, query.BOT},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(1,1,_)",
		Val1:    []query.FeatV{query.ONE, query.ONE, query.BOT},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(1,1,0)",
		Val1:    []query.FeatV{query.ONE, query.ONE, query.ZERO},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(1,_,1)",
		Val1:    []query.FeatV{query.ONE, query.BOT, query.ONE},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(_,1,_)",
		Val1:    []query.FeatV{query.BOT, query.ONE, query.BOT},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(_,_,_)",
		Val1:    []query.FeatV{query.BOT, query.BOT, query.BOT},
		ExpCode: 10,
	},
}
