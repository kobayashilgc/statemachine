# statemachine

go version state machine

## How to use

```go
package main

import (
	"statemachine/statemachine"
	"statemachine/statemachine_manager"
)

func foo(machine *statemachine.StateMachine, args ...interface{}) *statemachine.StateNode {
	toBeTrained := machine.GenerateStateNode("to be trained", "", nil)
	training := machine.GenerateStateNode("training", "", nil)
	machine.AddEdgeToNode("to be trained", []*statemachine.Edge{
		{
			Node: training,
			TransFunc: func(args ...interface{}) bool {
				a, b := 0, 0
				for i, arg := range args {
					if v, ok := arg.(int); ok && i == 0 {
						a = v
					}
					if v, ok := arg.(int); ok && i == 1 {
						b = v
					}
				}
				return a == b
			},
		},
	})
	return toBeTrained
}

func main() {
	// create a manager
	manager := statemachine_manager.StateMachineManagerFactory()
	// create a machine
	machine := statemachine.StateMachineFactory()
	// designer the machine
	machine.DesignStateMachine(foo, nil)
	// register the machine to the manager
	err := manager.RegisterStateMachine(1, machine)
	if err != nil {
		panic(err)
	}
	// get machine from manager
	m, err := manager.GetStateMachine(1)
	if err != nil {
		panic(err)
	}
	// update the machine
	m.Update(1, 1)
}

```