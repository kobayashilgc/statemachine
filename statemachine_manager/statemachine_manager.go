package statemachine_manager

import (
	"errors"
	"fmt"
	"statemachine/logger"
	sm "statemachine/statemachine"
	"sync"
)

var once sync.Once
var log *logger.Logger

type StateMachineManager struct {
	machines map[int]*sm.StateMachine
}

type IStateMachineManager interface {
	RegisterStateMachine(k int, m *sm.StateMachine) error
	GetStateMachine(k int) (*sm.StateMachine, error)
}

var manager *StateMachineManager

func StateMachineManagerFactory() *StateMachineManager {
	once.Do(func() {
		manager = &StateMachineManager{
			machines: map[int]*sm.StateMachine{},
		}
		log = &logger.Logger{}
		log.SetLevel(logger.LevelInfo)
	})
	log.Info("statemachine manager: create manager success")
	return manager
}

func (mgr *StateMachineManager) RegisterStateMachine(k int, m *sm.StateMachine) error {
	if !m.HasNode() {
		log.Error("statemachine manager: there is no node in machine, failed to register")
		return errors.New("statemachine manager: there is no node in machine, failed to register")
	}
	if _, ok := mgr.machines[k]; ok {
		log.Error(fmt.Sprintf("statemachine manager: machine of key:<%d> already exist, failed to register", k))
		return errors.New("key already exists, cannot register this machine")
	}
	mgr.machines[k] = m
	log.Info(fmt.Sprintf("statemachine manager: register machine of key:<%d> success", k))
	return nil
}

func (mgr *StateMachineManager) GetStateMachine(k int) (*sm.StateMachine, error) {
	if _, ok := mgr.machines[k]; !ok {
		log.Error(fmt.Sprintf("statemachine manager: machine of key:<%d> not exist!", k))
		return nil, errors.New("no key")
	}
	return mgr.machines[k], nil
}
