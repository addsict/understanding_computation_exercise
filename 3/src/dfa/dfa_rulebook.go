package dfa

type DFARulebook struct {
	rules []*FARule
}

func NewDFARulebook (rules []*FARule) *DFARulebook {
	return &DFARulebook{
		rules: rules,
	}
}

func (rb *DFARulebook) NextState(state int, character rune) int {
	return rb.RuleFor(state, character).NextState()
}

func (rb *DFARulebook) RuleFor(state int, character rune) *FARule {
	for _, rule := range rb.rules {
		if rule.CanApply(state, character) {
			return rule
		}
	}
	return nil
}
