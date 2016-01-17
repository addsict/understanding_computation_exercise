package dfa

type DFADesign struct {
	startState int
	acceptStates []int
	rulebook *DFARulebook
}

func NewDFADesign(startState int, acceptStates []int, rulebook *DFARulebook) *DFADesign {
	return &DFADesign{
		startState: startState,
		acceptStates: acceptStates,
		rulebook: rulebook,
	}
}

func (dd *DFADesign) ToDFA() *DFA {
	return NewDFA(dd.startState, dd.acceptStates, dd.rulebook)
}

func (dd *DFADesign) CanAccept(s string) bool {
	d := dd.ToDFA()
	d.ReadString(s)
	return d.IsAccepting()
}
