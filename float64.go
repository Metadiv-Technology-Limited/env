package env

import "github.com/Metadiv-Technology-Limited/env/internal/types"

/*
Float64 returns the value of the environment variable named by the key.
*/
func Float64(key string, defaultValue ...float64) float64 {
	m := types.NewManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToFloat64()
}

/*
Float64s returns the list of values of the environment variable named by the key.
*/
func Float64s(key string, defaultValue ...[]float64) []float64 {
	m := types.NewManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToFloat64s()
}

/*
MustFloat64 returns the value of the environment variable named by the key.
It panics if the value is empty.
*/
func MustFloat64(key string) float64 {
	m := types.NewMustManager[float64](key)
	return m.Value.ToFloat64()
}

/*
MustFloat64s returns the list of values of the environment variable named by the key.
It panics if the value is empty.
*/
func MustFloat64s(key string) []float64 {
	m := types.NewMustManager[float64](key)
	return m.Value.ToFloat64s()
}
