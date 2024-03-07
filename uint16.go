package env

import "github.com/Metadiv-Technology-Limited/env/internal/types"

/*
Uint16 returns the value of the environment variable named by the key.
*/
func Uint16(key string, defaultValue ...uint16) uint16 {
	m := types.NewManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToUint16()
}

/*
Uint16s returns the list of values of the environment variable named by the key.
*/
func Uint16s(key string, defaultValue ...[]uint16) []uint16 {
	m := types.NewManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToUint16s()
}

/*
MustUint16 returns the value of the environment variable named by the key.
It panics if the value is empty.
*/
func MustUint16(key string) uint16 {
	m := types.NewMustManager[uint16](key)
	return m.Value.ToUint16()
}

/*
MustUint16s returns the list of values of the environment variable named by the key.
It panics if the value is empty.
*/
func MustUint16s(key string) []uint16 {
	m := types.NewMustManager[uint16](key)
	return m.Value.ToUint16s()
}
