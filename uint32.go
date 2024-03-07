package env

import "github.com/Metadiv-Technology-Limited/env/internal/types"

/*
Uint32 returns the value of the environment variable named by the key.
*/
func Uint32(key string, defaultValue ...uint32) uint32 {
	m := types.NewManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToUint32()
}

/*
Uint32s returns the list of values of the environment variable named by the key.
*/
func Uint32s(key string, defaultValue ...[]uint32) []uint32 {
	m := types.NewManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToUint32s()
}

/*
MustUint32 returns the value of the environment variable named by the key.
It panics if the value is empty.
*/
func MustUint32(key string) uint32 {
	m := types.NewMustManager[uint32](key)
	return m.Value.ToUint32()
}

/*
MustUint32s returns the list of values of the environment variable named by the key.
It panics if the value is empty.
*/
func MustUint32s(key string) []uint32 {
	m := types.NewMustManager[uint32](key)
	return m.Value.ToUint32s()
}
