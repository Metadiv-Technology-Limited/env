package env

import "github.com/Metadiv-Technology-Limited/env/internal/types"

/*
Int64 returns the value of the environment variable named by the key.
*/
func Int64(key string, defaultValue ...int64) int64 {
	m := types.NewManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToInt64()
}

/*
Int64s returns the list of values of the environment variable named by the key.
*/
func Int64s(key string, defaultValue ...[]int64) []int64 {
	m := types.NewManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToInt64s()
}

/*
MustInt64 returns the value of the environment variable named by the key.
It panics if the value is empty.
*/
func MustInt64(key string) int64 {
	m := types.NewMustManager[int64](key)
	return m.Value.ToInt64()
}

/*
MustInt64s returns the list of values of the environment variable named by the key.
It panics if the value is empty.
*/
func MustInt64s(key string) []int64 {
	m := types.NewMustManager[int64](key)
	return m.Value.ToInt64s()
}
