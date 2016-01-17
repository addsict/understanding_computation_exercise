package dfa

type DFA struct {
	currState int
	acceptStates []int
	rulebook *DFARulebook
}

func NewDFA(currState int, acceptStates []int, rulebook *DFARulebook) *DFA {
	return &DFA{
		currState: currState,
		acceptStates: acceptStates,
		rulebook: rulebook,
	}
}

func (dfa *DFA) IsAccepting() bool {
	for _, state := range dfa.acceptStates {
		if state == dfa.currState {
			return true
		}
	}
	return false
}

func (dfa *DFA) ReadChar(char rune) {
	dfa.currState = dfa.rulebook.NextState(dfa.currState, char)
}

func (dfa *DFA) ReadString(s string) {
	for _, c := range s {
		dfa.ReadChar(c)
	}
}
