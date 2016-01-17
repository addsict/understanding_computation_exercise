package nfa

type NFADesign struct {
	startState int
	acceptStates []int
	rulebook *NFARulebook
}

func NewNFADesign(startState int, acceptStates []int, rulebook *NFARulebook) *NFADesign {
	return &NFADesign{
		startState: startState,
		acceptStates: acceptStates,
		rulebook: rulebook,
	}
}

func (nd *NFADesign) ToNFA() *NFA {
	return NewNFA([]int{ nd.startState }, nd.acceptStates, nd.rulebook)
}

func (nd *NFADesign) CanAccept(s string) bool {
	n := nd.ToNFA()
	n.ReadString(s)
	return n.IsAccepting()
}
