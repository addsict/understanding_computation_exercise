package main

import (
	"dfa"
	"fmt"
)

func main() {
	rb := dfa.NewDFARulebook([]*dfa.FARule{
		dfa.NewFARule(1, 'b', 1),
		dfa.NewFARule(1, 'a', 2),
		dfa.NewFARule(2, 'a', 2),
		dfa.NewFARule(2, 'b', 2),
	})
	fmt.Printf("%d => (%c) => %d\n", 1, 'a', rb.NextState(1, 'a'))
	fmt.Printf("%d => (%c) => %d\n", 1, 'b', rb.NextState(1, 'b'))
	fmt.Printf("%d => (%c) => %d\n", 2, 'a', rb.NextState(2, 'a'))
	fmt.Printf("%d => (%c) => %d\n", 2, 'b', rb.NextState(2, 'b'))
}
