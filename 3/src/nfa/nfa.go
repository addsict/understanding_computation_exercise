package nfa

type NFA struct {
	currStates []int
	acceptStates []int
	rulebook *NFARulebook
}

func NewNFA(currStates []int, acceptStates []int, rulebook *NFARulebook) *NFA {
	return &NFA{
		currStates: currStates,
		acceptStates: acceptStates,
		rulebook: rulebook,
	}
}

func (nfa *NFA) IsAccepting() bool {
	for _, as := range nfa.acceptStates {
		for _, cs := range nfa.currStates {
			if cs == as {
				return true
			}
		}
	}
	return false
}

func (nfa *NFA) ReadChar(char rune) {
	nfa.currStates = nfa.rulebook.NextStates(nfa.currStates, char)
}

func (nfa *NFA) ReadString(s string) {
	for _, c := range s {
		nfa.ReadChar(c)
	}
}
