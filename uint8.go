package env

import "github.com/Metadiv-Technology-Limited/env/internal/types"

/*
Uint8 returns the value of the environment variable named by the key.
*/
func Uint8(key string, defaultValue ...uint8) uint8 {
	m := types.NewManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToUint8()
}

/*
Uint8s returns the list of values of the environment variable named by the key.
*/
func Uint8s(key string, defaultValue ...[]uint8) []uint8 {
	m := types.NewManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToUint8s()
}

/*
MustUint8 returns the value of the environment variable named by the key.
It panics if the value is empty.
*/
func MustUint8(key string) uint8 {
	m := types.NewMustManager[uint8](key)
	return m.Value.ToUint8()
}

/*
MustUint8s returns the list of values of the environment variable named by the key.
It panics if the value is empty.
*/
func MustUint8s(key string) []uint8 {
	m := types.NewMustManager[uint8](key)
	return m.Value.ToUint8s()
}
