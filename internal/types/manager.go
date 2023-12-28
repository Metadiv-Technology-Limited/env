package types

import (
	"os"
)

func NewManager[T any](key string, defaultValue ...T) *Manager[T] {
	var d T
	var has bool
	if len(defaultValue) > 0 {
		d = defaultValue[0]
		has = true
	}
	return &Manager[T]{
		Key:          key,
		Value:        NewValue(os.Getenv(key)),
		DefaultValue: d,
		HasDefault:   has,
		Must:         false,
	}
}

func NewMustManager[T any](key string) *Manager[T] {
	m := NewManager[T](key)
	m.Must = true
	if os.Getenv("GIN_MODE") == "release" || os.Getenv("GIN_MODE") == "debug" {
		if m.Value.ToString() == "" {
			panic("env: key " + key + " not found")
		}
	}
	return m
}

type Manager[T any] struct {
	Key          string
	Value        *Value
	DefaultValue T
	HasDefault   bool
	Must         bool
}

func (m *Manager[T]) EmptyValue() bool {
	return m.Value.Value == ""
}

func (m *Manager[T]) ReturnDefault() T {
	if m.HasDefault {
		return m.DefaultValue
	}
	return *new(T)
}
