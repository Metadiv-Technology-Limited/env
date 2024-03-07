package env

import "github.com/Metadiv-Technology-Limited/env/internal/types"

/*
Int32 returns the value of the environment variable named by the key.
*/
func Int32(key string, defaultValue ...int32) int32 {
	m := types.NewManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToInt32()
}

/*
Int32s returns the value of the environment variable named by the key.
*/
func Int32s(key string, defaultValue ...[]int32) []int32 {
	m := types.NewManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToInt32s()
}

/*
MustInt32 returns the value of the environment variable named by the key.
If the value is empty, it will panic.
*/
func MustInt32(key string) int32 {
	m := types.NewMustManager[int16](key)
	return m.Value.ToInt32()
}

/*
MustInt32s returns the value of the environment variable named by the key.
If the value is empty, it will panic.
*/
func MustInt32s(key string) []int32 {
	m := types.NewMustManager[int16](key)
	return m.Value.ToInt32s()
}
