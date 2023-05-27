package statemachine

import (
	"fmt"
)

type StateMachine struct {
	initNode *StateNode
	curNode  *StateNode
}

var logger *Logger

var machines map[int]*StateMachine

func init() {
	logger = &Logger{}
	logger.SetLevel(LevelInfo)
	machines = map[int]*StateMachine{}
}

type IStateMachine interface {
	design() *StateNode
	designStateMachine() *StateMachine
	Update(args ...interface{})
	Get() string
}

// design - design your state machine here
func (m *StateMachine) design() *StateNode {
	return nil
}

// desginStateMachine - design a state machine
func (m *StateMachine) designStateMachine() *StateMachine {
	initNode := m.design()
	m.initNode = initNode
	m.curNode = initNode
	return m
}

func (m *StateMachine) Update(args ...interface{}) {
	oldNodeName := m.curNode.GetName()
	nextNode := m.curNode.Trans(args...)
	m.curNode = nextNode
	newNodeName := m.curNode.GetName()
	logger.Info(fmt.Sprintf("machine has trans from <%s> to <%s>", oldNodeName, newNodeName))
}

func (m *StateMachine) Get() string {
	return m.curNode.GetName()
}
