package main

import (
	"fa"
	"nfa"
	"fmt"
)

func main() {
	rb := nfa.NewNFARulebook([]*fa.FARule{
		fa.NewFARule(1, 'a', 1), fa.NewFARule(1, 'b', 1), fa.NewFARule(1, 'b', 2),
		fa.NewFARule(2, 'a', 1), fa.NewFARule(2, 'a', 3), fa.NewFARule(2, 'b', 3),
		fa.NewFARule(3, 'a', 3), fa.NewFARule(3, 'b', 3), fa.NewFARule(3, 'a', 4), fa.NewFARule(3, 'b', 4),
	})

	nd := nfa.NewNFADesign(1, []int{4}, rb)
	fmt.Printf("accepts %s => %t\n", "abab", nd.CanAccept("abab"))
	fmt.Printf("accepts %s => %t\n", "aaaba", nd.CanAccept("aaaba"))
	fmt.Printf("accepts %s => %t\n", "aaabaa", nd.CanAccept("aaabaa"))
	fmt.Printf("accepts %s => %t\n", "aaabaaabab", nd.CanAccept("aaabaaabab"))
}
