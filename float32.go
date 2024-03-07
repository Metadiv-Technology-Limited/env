package env

import "github.com/Metadiv-Technology-Limited/env/internal/types"

/*
Float32 returns the value of the environment variable named by the key.
*/
func Float32(key string, defaultValue ...float32) float32 {
	m := types.NewManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToFloat32()
}

/*
Float32s returns the list of values of the environment variable named by the key.
*/
func Float32s(key string, defaultValue ...[]float32) []float32 {
	m := types.NewManager(key, defaultValue...)
	if m.EmptyValue() {
		return m.ReturnDefault()
	}
	return m.Value.ToFloat32s()
}

/*
MustFloat32 returns the value of the environment variable named by the key.
It panics if the value is empty.
*/
func MustFloat32(key string) float32 {
	m := types.NewMustManager[float32](key)
	return m.Value.ToFloat32()
}

/*
MustFloat32s returns the list of values of the environment variable named by the key.
It panics if the value is empty.
*/
func MustFloat32s(key string) []float32 {
	m := types.NewMustManager[float32](key)
	return m.Value.ToFloat32s()
}
