package statemachine

import "testing"

func TestLogger(t *testing.T) {
	l := &Logger{}
	l.SetLevel(LevelNo)
	l.Info("info")
	l.Warning("warning")
	l.Error("error")
	l.SetLevel(LevelNo)
	l.Warning("warning")
}
