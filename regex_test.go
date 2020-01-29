package main

import (
	"github.com/KazumaTakata/shunting-yard"
	"testing"
)

func TestRegex(t *testing.T) {

	operators := []shunting.Operator{}
	operators = append(operators, shunting.Operator{Value: '|', Precedence: 0, IsLeftAssociative: true})
	operators = append(operators, shunting.Operator{Value: '.', Precedence: 1, IsLeftAssociative: true})
	operators = append(operators, shunting.Operator{Value: '+', Precedence: 2, IsLeftAssociative: true})
	operators = append(operators, shunting.Operator{Value: '*', Precedence: 2, IsLeftAssociative: true})

	i2p := shunting.NewIn2Post(operators)

	run_regex(i2p, "a+.b", "aaaab", t)
	run_regex(i2p, "(1|2|3)", "1", t)
	run_regex(i2p, "a+.(1|b|c)", "aaaab", t)

}

func run_regex(i2p shunting.In2Post, regex, input string, t *testing.T) {

	postfix := i2p.Parse(regex)
	postfix = []byte(postfix)

	start_state := ConstructNFA(postfix)
	state_list := SimulateNFA(start_state, input)

	//	fmt.Printf("%+v", state_list)

	if !isMatch(state_list) {
		t.Errorf("regex failed")

	}

}
