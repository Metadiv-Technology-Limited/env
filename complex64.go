package env

import "github.com/Metadiv-Technology-Limited/env/internal/types"

/*
Complex64 returns the value of the environment variable named by the key.
*/
func Complex64(key string, defaultValue ...complex64) complex64 {
	m := types.NewManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToComplex64()
}

/*
Complex64s returns the list of values of the environment variable named by the key.
*/
func Complex64s(key string, defaultValue ...[]complex64) []complex64 {
	m := types.NewManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToComplex64s()
}

/*
MustComplex64 returns the value of the environment variable named by the key.
It panics if the value is empty.
*/
func MustComplex64(key string) complex64 {
	m := types.NewMustManager[complex64](key)
	return m.Value.ToComplex64()
}

/*
MustComplex64s returns the list of values of the environment variable named by the key.
It panics if the value is empty.
*/
func MustComplex64s(key string) []complex64 {
	m := types.NewMustManager[complex64](key)
	return m.Value.ToComplex64s()
}
