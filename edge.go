package statemachine

type Edge struct {
	TransFunc func(args ...interface{}) bool
	Node      *StateNode
}
