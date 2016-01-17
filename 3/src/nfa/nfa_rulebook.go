package nfa

import (
	"fa"
	"fmt"
)

type NFARulebook struct {
	rules []*fa.FARule
}

func NewNFARulebook (rules []*fa.FARule) *NFARulebook {
	return &NFARulebook{
		rules: rules,
	}
}

func (rb *NFARulebook) NextStates(states []int, character rune) []int {
	var nextStates []int
	for _, s := range states {
		rules := rb.rulesFor(s, character)
		for _, r := range rules {
			nextStates = append(nextStates, r.NextState())
		}
	}

	var uniqStates []int
	for _, ns := range nextStates {
		found := false
		for _, us := range uniqStates {
			if us == ns {
				found = true
			}
		}
		if !found {
			uniqStates = append(uniqStates, ns)
		}
	}
	fmt.Printf("ns: %v, us: %v\n", nextStates, uniqStates)
	return uniqStates
}

func (rb *NFARulebook) rulesFor(state int, character rune) []*fa.FARule {
	var rules []*fa.FARule
	for _, r := range rb.rules {
		if r.CanApply(state, character) {
			rules = append(rules, r)
		}
	}
	return rules
}
