package main

import (
	"fa"
	"dfa"
	"fmt"
)

func main() {
	rb := dfa.NewDFARulebook([]*fa.FARule{
		fa.NewFARule(1, 'a', 2),
		fa.NewFARule(1, 'b', 2),
		fa.NewFARule(2, 'a', 3),
		fa.NewFARule(2, 'b', 3),
		fa.NewFARule(3, 'a', 4),
		fa.NewFARule(3, 'b', 5),
		fa.NewFARule(4, 'a', 4),
		fa.NewFARule(4, 'b', 4),
		fa.NewFARule(5, 'a', 5),
		fa.NewFARule(5, 'b', 5),
	})

	// DFA that "b" must be at 3rd position of input sequence
	dd := dfa.NewDFADesign(1, []int{5}, rb)
	fmt.Printf("accepts %s => %t\n", "abab", dd.CanAccept("abab"))
	fmt.Printf("accepts %s => %t\n", "aabab", dd.CanAccept("aabab"))
	fmt.Printf("accepts %s => %t\n", "bbaba", dd.CanAccept("bbaba"))
}
