package dfa

import (
	"fa"
)

type DFARulebook struct {
	rules []*fa.FARule
}

func NewDFARulebook (rules []*fa.FARule) *DFARulebook {
	return &DFARulebook{
		rules: rules,
	}
}

func (rb *DFARulebook) NextState(state int, character rune) int {
	return rb.ruleFor(state, character).NextState()
}

func (rb *DFARulebook) ruleFor(state int, character rune) *fa.FARule {
	for _, rule := range rb.rules {
		if rule.CanApply(state, character) {
			return rule
		}
	}
	return nil
}
