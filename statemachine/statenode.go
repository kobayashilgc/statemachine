package statemachine

import "github.com/google/uuid"

type StateNode struct {
	uuid  uuid.UUID
	name  string
	info  string
	edges []*Edge
}

type IStateNode interface {
	GetUUID() uuid.UUID
	GetName() string
	GetInfo() string
	Trans(args ...interface{}) *StateNode
}

func (n *StateNode) GetUUID() uuid.UUID {
	return n.uuid
}

func (n *StateNode) GetName() string {
	return n.name
}

func (n *StateNode) GetInfo() string {
	return n.info
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
