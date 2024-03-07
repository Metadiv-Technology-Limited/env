package env

import "github.com/Metadiv-Technology-Limited/env/internal/types"

/*
Complex128 returns the value of the environment variable named by the key.
*/
func Complex128(key string, defaultValue ...complex128) complex128 {
	m := types.NewManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToComplex128()
}

/*
Complex128s returns the list of values of the environment variable named by the key.
*/
func Complex128s(key string, defaultValue ...[]complex128) []complex128 {
	m := types.NewManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToComplex128s()
}

/*
MustComplex128 returns the value of the environment variable named by the key.
It panics if the value is empty.
*/
func MustComplex128(key string) complex128 {
	m := types.NewMustManager[complex128](key)
	return m.Value.ToComplex128()
}

/*
MustComplex128s returns the list of values of the environment variable named by the key.
It panics if the value is empty.
*/
func MustComplex128s(key string) []complex128 {
	m := types.NewMustManager[complex128](key)
	return m.Value.ToComplex128s()
}
