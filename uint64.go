package env

import "github.com/Metadiv-Technology-Limited/env/internal/types"

/*
Uint64 returns the value of the environment variable named by the key.
*/
func Uint64(key string, defaultValue ...uint64) uint64 {
	m := types.NewManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToUint64()
}

/*
Uint64s returns the list of values of the environment variable named by the key.
*/
func Uint64s(key string, defaultValue ...[]uint64) []uint64 {
	m := types.NewManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToUint64s()
}

/*
MustUint64 returns the value of the environment variable named by the key.
It panics if the value is empty.
*/
func MustUint64(key string) uint64 {
	m := types.NewMustManager[uint64](key)
	return m.Value.ToUint64()
}

/*
MustUint64s returns the list of values of the environment variable named by the key.
It panics if the value is empty.
*/
func MustUint64s(key string) []uint64 {
	m := types.NewMustManager[uint64](key)
	return m.Value.ToUint64s()
}
