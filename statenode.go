package statemachine

type StateNode struct {
	name  string
	info  string
	edges []*Edge
}

type IStateNode interface {
	Set(name, info string, e []*Edge)
	GetName() string
	GetInfo() string
	Trans(args ...interface{}) *StateNode
}

func (n *StateNode) Set(name, info string, e []*Edge) {
	n.name, n.info, n.edges = name, info, e
}

func (n *StateNode) GetName() string {
	return n.name
}

func (n *StateNode) GetInfo() string {
	return n.name
}

// Trans - trans to next state, if no edge is ok, return current node itself
func (n *StateNode) Trans(args ...interface{}) *StateNode {
	var nextNode *StateNode = nil
	for _, e := range n.edges {
		if e.TransFunc(args...) {
			nextNode = e.Node
			break
		}
	}
	if nextNode == nil {
		return n
	}
	return nextNode
}
