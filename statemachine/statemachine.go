package statemachine

import (
	"errors"
	"fmt"

	"statemachine/logger"

	"github.com/google/uuid"
)

type StateMachine struct {
	nodeMap  map[string]*StateNode
	initNode *StateNode
	curNode  *StateNode
}

var log *logger.Logger

func init() {
	log = &logger.Logger{}
	log.SetLevel(logger.LevelInfo)
}

// StateMachineFactory - factory function
func StateMachineFactory() *StateMachine {
	m := &StateMachine{
		nodeMap: map[string]*StateNode{},
	}
	return m
}

type IStateMachine interface {
	GenerateStateNode(name, info string, edges []*Edge) *StateNode
	AddEdgeToNode(name string, edges []*Edge)
	DesignStateMachine(designFunc func(machine *StateMachine, args ...interface{}) *StateNode, args ...interface{}) *StateMachine
	Update(args ...interface{})
	Get() string
	HasNode() bool
}

func (m *StateMachine) HasNode() bool {
	b := len(m.nodeMap) > 0
	if !b {
		log.Error("statemachine: machine need to be designed first!")
	}
	return b
}

// generateStateNode - check and generate a state node
func (m *StateMachine) GenerateStateNode(name, info string, edges []*Edge) *StateNode {
	if _, ok := m.nodeMap[name]; ok {
		log.Warning(fmt.Sprintf("statemachine: duplicated node name: %s", name))
	}
	node := &StateNode{
		uuid:  uuid.New(),
		name:  name,
		info:  info,
		edges: edges,
	}
	m.nodeMap[name] = node
	return node
}

func (m *StateMachine) AddEdgeToNode(name string, edge []*Edge) {
	if _, ok := m.nodeMap[name]; !ok {
		log.Warning(fmt.Sprintf("statemachine: node:<%s> does not exist, cannot add edges to it", name))
		return
	}
	m.nodeMap[name].edges = edge
}

// desginStateMachine - design a state machine
func (m *StateMachine) DesignStateMachine(designFunc func(machine *StateMachine, args ...interface{}) *StateNode, args ...interface{}) *StateMachine {
	if designFunc == nil {
		log.Warning("statemachine: designer func is nil!")
	}
	initNode := designFunc(m, args...)
	m.initNode = initNode
	m.curNode = initNode
	return m
}

// Update - update machine, may trans to another state
func (m *StateMachine) Update(args ...interface{}) error {
	if !m.HasNode() {
		log.Error("statemachine: no node, cannot update")
		return errors.New("no node")
	}
	oldNodeName := m.curNode.GetName()
	oldNodeUUID := m.curNode.GetUUID()
	nextNode := m.curNode.Trans(args...)
	m.curNode = nextNode
	newNodeName := m.curNode.GetName()
	newNodeUUID := m.curNode.GetUUID()
	if oldNodeUUID == newNodeUUID {
		log.Info(fmt.Sprintf("statemachine: machine has trans from <%s> to itself", oldNodeName))
	} else {
		if oldNodeName == newNodeName {
			log.Warning(fmt.Sprintf("statemachine: machine has trans from <%s> to <%s>, where the two nodes has the same name", oldNodeName, newNodeName))
		} else {
			log.Info(fmt.Sprintf("statemachine: machine has trans from <%s> to <%s>", oldNodeName, newNodeName))
		}
	}
	return nil
}

func (m *StateMachine) Get() string {
	if !m.HasNode() {
		log.Warning("statemachine: no node, return an unknown state")
		return "unknown state"
	}
	return m.curNode.GetName()
}
