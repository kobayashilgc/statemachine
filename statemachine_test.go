package statemachine

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/sirupsen/logrus"
)

func TestStateMachine(t *testing.T) {
	m := StateMachineFactory()
	RegisterStateMachine(1, m)
	mm, err := GetStateMachine(1)
	assert.Equal(t, err, nil)
	mm.Update(nil)
	logrus.Info(m.Get())
}
