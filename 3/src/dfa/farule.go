package dfa

type FARule struct {
	state int
	character rune
	nextState int
}

func NewFARule(state int, character rune, nextState int) *FARule {
	return &FARule{
		state: state,
		character: character,
		nextState: nextState,
	}
}

func (r *FARule) CanApply(state int, character rune) bool {
	return r.state == state && r.character == character
}

func (r *FARule) NextState() int {
	return r.nextState
}
