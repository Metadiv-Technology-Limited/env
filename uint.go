package env

import "github.com/Metadiv-Technology-Limited/env/internal/types"

/*
Uint returns the value of the environment variable named by the key.
*/
func Uint(key string, defaultValue ...uint) uint {
	m := types.NewManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToUint()
}

/*
Uints returns the list of values of the environment variable named by the key.
*/
func Uints(key string, defaultValue ...[]uint) []uint {
	m := types.NewManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToUints()
}

/*
MustUint returns the value of the environment variable named by the key.
It panics if the value is empty.
*/
func MustUint(key string) uint {
	m := types.NewMustManager[uint](key)
	return m.Value.ToUint()
}

/*
MustUints returns the list of values of the environment variable named by the key.
It panics if the value is empty.
*/
func MustUints(key string) []uint {
	m := types.NewMustManager[uint](key)
	return m.Value.ToUints()
}
