package suf_test

import (
	"github.com/jtcaraball/goexpdt/query"
	"github.com/DiegoEmilio01/goexpfbdd/query/internal/test"
)

var SufPTT = []test.BTRecord{
	{
		Dim:     3,
		Name:    "(1,_,_), (0,_,_)",
		Val1:    []query.FeatV{query.ONE, query.BOT, query.BOT},
		Val2:    []query.FeatV{query.ZERO, query.BOT, query.BOT},
		ExpCode: 10, // successful
	},
	{
		Dim:     3,
		Name:    "(1,_,_), (0,_,0)",
		Val1:    []query.FeatV{query.ONE, query.BOT, query.BOT},
		Val2:    []query.FeatV{query.ZERO, query.BOT, query.ZERO},
		ExpCode: 20, // unsuccessful
	},
	{
		Dim:     3,
		Name:    "(1,_,_), (1,1,0)",
		Val1:    []query.FeatV{query.ONE, query.BOT, query.BOT},
		Val2:    []query.FeatV{query.ONE, query.ONE, query.ZERO},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(1,_,1), (0,_,0)",
		Val1:    []query.FeatV{query.ONE, query.BOT, query.ONE},
		Val2:    []query.FeatV{query.ZERO, query.BOT, query.ZERO},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(1,1,_), (1,_,0)",
		Val1:    []query.FeatV{query.ONE, query.ONE, query.BOT},
		Val2:    []query.FeatV{query.ONE, query.BOT, query.ZERO},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(0,1,0), (0,_,1)",
		Val1:    []query.FeatV{query.ZERO, query.ONE, query.ZERO},
		Val2:    []query.FeatV{query.ZERO, query.BOT, query.ONE},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(1,1,0), (1,1,0)",
		Val1:    []query.FeatV{query.ONE, query.ONE, query.ZERO},
		Val2:    []query.FeatV{query.ONE, query.ONE, query.ZERO},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(0,1,0), (0,0,1)",
		Val1:    []query.FeatV{query.ZERO, query.ONE, query.ZERO},
		Val2:    []query.FeatV{query.ZERO, query.ZERO, query.ONE},
		ExpCode: 10,
	},
}

var SufNTT = []test.BTRecord{
	{
		Dim:     3,
		Name:    "(1,_,_), (0,_,_)",
		Val1:    []query.FeatV{query.ONE, query.BOT, query.BOT},
		Val2:    []query.FeatV{query.ZERO, query.BOT, query.BOT},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(1,_,_), (0,_,0)",
		Val1:    []query.FeatV{query.ONE, query.BOT, query.BOT},
		Val2:    []query.FeatV{query.ZERO, query.BOT, query.ZERO},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(1,_,_), (1,1,0)",
		Val1:    []query.FeatV{query.ONE, query.BOT, query.BOT},
		Val2:    []query.FeatV{query.ONE, query.ONE, query.ZERO},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(1,_,1), (0,_,0)",
		Val1:    []query.FeatV{query.ONE, query.BOT, query.ONE},
		Val2:    []query.FeatV{query.ZERO, query.BOT, query.ZERO},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(1,1,_), (1,_,0)",
		Val1:    []query.FeatV{query.ONE, query.ONE, query.BOT},
		Val2:    []query.FeatV{query.ONE, query.BOT, query.ZERO},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(0,1,0), (0,_,1)",
		Val1:    []query.FeatV{query.ZERO, query.ONE, query.ZERO},
		Val2:    []query.FeatV{query.ZERO, query.BOT, query.ONE},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(1,1,0), (1,1,0)",
		Val1:    []query.FeatV{query.ONE, query.ONE, query.ZERO},
		Val2:    []query.FeatV{query.ONE, query.ONE, query.ZERO},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(0,1,0), (0,0,1)",
		Val1:    []query.FeatV{query.ZERO, query.ONE, query.ZERO},
		Val2:    []query.FeatV{query.ZERO, query.ZERO, query.ONE},
		ExpCode: 20,
	},
}
