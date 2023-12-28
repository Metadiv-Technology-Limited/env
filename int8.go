package env

import "github.com/Metadiv-Technology-Limited/env/internal/types"

/*
Int8 returns the value of the environment variable named by the key.
*/
func Int8(key string, defaultValue ...int8) int8 {
	m := types.NewManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToInt8()
}

/*
Int8s returns the list of values of the environment variable named by the key.
*/
func Int8s(key string, defaultValue ...[]int8) []int8 {
	m := types.NewManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToInt8s()
}

/*
MustInt8 returns the value of the environment variable named by the key.
It panics if the value is empty.
*/
func MustInt8(key string) int8 {
	m := types.NewMustManager[int8](key)
	return m.Value.ToInt8()
}

/*
MustInt8s returns the list of values of the environment variable named by the key.
It panics if the value is empty.
*/
func MustInt8s(key string) []int8 {
	m := types.NewMustManager[int8](key)
	return m.Value.ToInt8s()
}
