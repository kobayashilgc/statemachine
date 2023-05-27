package statemachine

import (
	"github.com/sirupsen/logrus"
)

type LoggerLevel uint8

const (
	LevelInfo LoggerLevel = iota
	LevelWarning
	LevelError
	LevelNo
)

type Logger struct {
	level LoggerLevel
}

type LoggerInterface interface {
	SetLevel(x int)
	Info(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
}

func (l *Logger) SetLevel(x LoggerLevel) {
	l.level = x
}

func (l *Logger) Info(args ...interface{}) {
	if l.level <= LevelInfo {
		logrus.Info(args...)
	}
}

func (l *Logger) Warning(args ...interface{}) {
	if l.level <= LevelWarning {
		logrus.Warning(args...)
	}
}

func (l *Logger) Error(args ...interface{}) {
	if l.level <= LevelError {
		logrus.Error(args...)
	}
}
