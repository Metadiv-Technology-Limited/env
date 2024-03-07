package env

import "github.com/Metadiv-Technology-Limited/env/internal/types"

/*
Int16 returns the value of the environment variable named by the key.
*/
func Int16(key string, defaultValue ...int16) int16 {
	m := types.NewManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToInt16()
}

/*
Int16s returns the list of values of the environment variable named by the key.
*/
func Int16s(key string, defaultValue ...[]int16) []int16 {
	m := types.NewManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToInt16s()
}

/*
MustInt16 returns the value of the environment variable named by the key.
It panics if the value is empty.
*/
func MustInt16(key string) int16 {
	m := types.NewMustManager[int16](key)
	return m.Value.ToInt16()
}

/*
MustInt16s returns the list of values of the environment variable named by the key.
It panics if the value is empty.
*/
func MustInt16s(key string) []int16 {
	m := types.NewMustManager[int16](key)
	return m.Value.ToInt16s()
}
