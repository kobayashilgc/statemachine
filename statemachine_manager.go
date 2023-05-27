package statemachine

import "errors"

// StateMachineFactory - factory function
func StateMachineFactory() *StateMachine {
	m := &StateMachine{}
	return m.designStateMachine()
}

func RegisterStateMachine(k int, m *StateMachine) error {
	if _, ok := machines[k]; !ok {
		logger.Error("key already exists, cannot register this machine")
		return errors.New("key already exists, cannot register this machine")
	}
	machines[k] = m
	return nil
}

func GetStateMachine(k int) (*StateMachine, error) {
	if _, ok := machines[k]; !ok {
		return nil, errors.New("no key")
	}
	return machines[k], nil
}
