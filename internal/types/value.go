package types

import (
	"strconv"
	"strings"
)

func NewValue(v string) *Value {
	return &Value{Value: v}
}

type Value struct {
	Value string
}

func (v *Value) ToString() string {
	return v.Value
}

func (v *Value) ToStrings() []string {
	return strings.Split(v.Value, ",")
}

func (v *Value) ToInt() int {
	output, err := strconv.Atoi(v.Value)
	if err != nil {
		panic("env: value is not an int.")
	}
	return output
}

func (v *Value) ToInts() []int {
	var output []int
	for _, v := range v.ToStrings() {
		output = append(output, NewValue(v).ToInt())
	}
	return output
}

func (v *Value) ToInt64() int64 {
	output, err := strconv.ParseInt(v.Value, 10, 64)
	if err != nil {
		panic("env: value is not an int64.")
	}
	return output
}

func (v *Value) ToInt64s() []int64 {
	var output []int64
	for _, v := range v.ToStrings() {
		output = append(output, NewValue(v).ToInt64())
	}
	return output
}

func (v *Value) ToInt32() int32 {
	output, err := strconv.ParseInt(v.Value, 10, 32)
	if err != nil {
		panic("env: value is not an int32.")
	}
	return int32(output)
}

func (v *Value) ToInt32s() []int32 {
	var output []int32
	for _, v := range v.ToStrings() {
		output = append(output, NewValue(v).ToInt32())
	}
	return output
}

func (v *Value) ToInt16() int16 {
	output, err := strconv.ParseInt(v.Value, 10, 16)
	if err != nil {
		panic("env: value is not an int16.")
	}
	return int16(output)
}

func (v *Value) ToInt16s() []int16 {
	var output []int16
	for _, v := range v.ToStrings() {
		output = append(output, NewValue(v).ToInt16())
	}
	return output
}

func (v *Value) ToInt8() int8 {
	output, err := strconv.ParseInt(v.Value, 10, 8)
	if err != nil {
		panic("env: value is not an int8.")
	}
	return int8(output)
}

func (v *Value) ToInt8s() []int8 {
	var output []int8
	for _, v := range v.ToStrings() {
		output = append(output, NewValue(v).ToInt8())
	}
	return output
}

func (v *Value) ToUint() uint {
	output, err := strconv.ParseUint(v.Value, 10, 64)
	if err != nil {
		panic("env: value is not an uint.")
	}
	return uint(output)
}

func (v *Value) ToUints() []uint {
	var output []uint
	for _, v := range v.ToStrings() {
		output = append(output, NewValue(v).ToUint())
	}
	return output
}

func (v *Value) ToUint64() uint64 {
	output, err := strconv.ParseUint(v.Value, 10, 64)
	if err != nil {
		panic("env: value is not an uint64.")
	}
	return output
}

func (v *Value) ToUint64s() []uint64 {
	var output []uint64
	for _, v := range v.ToStrings() {
		output = append(output, NewValue(v).ToUint64())
	}
	return output
}

func (v *Value) ToUint32() uint32 {
	output, err := strconv.ParseUint(v.Value, 10, 32)
	if err != nil {
		panic("env: value is not an uint32.")
	}
	return uint32(output)
}

func (v *Value) ToUint32s() []uint32 {
	var output []uint32
	for _, v := range v.ToStrings() {
		output = append(output, NewValue(v).ToUint32())
	}
	return output
}

func (v *Value) ToUint16() uint16 {
	output, err := strconv.ParseUint(v.Value, 10, 16)
	if err != nil {
		panic("env: value is not an uint16.")
	}
	return uint16(output)
}

func (v *Value) ToUint16s() []uint16 {
	var output []uint16
	for _, v := range v.ToStrings() {
		output = append(output, NewValue(v).ToUint16())
	}
	return output
}

func (v *Value) ToUint8() uint8 {
	output, err := strconv.ParseUint(v.Value, 10, 8)
	if err != nil {
		panic("env: value is not an uint8.")
	}
	return uint8(output)
}

func (v *Value) ToUint8s() []uint8 {
	var output []uint8
	for _, v := range v.ToStrings() {
		output = append(output, NewValue(v).ToUint8())
	}
	return output
}

func (v *Value) ToFloat32() float32 {
	output, err := strconv.ParseFloat(v.Value, 32)
	if err != nil {
		panic("env: value is not an float32.")
	}
	return float32(output)
}

func (v *Value) ToFloat32s() []float32 {
	var output []float32
	for _, v := range v.ToStrings() {
		output = append(output, NewValue(v).ToFloat32())
	}
	return output
}

func (v *Value) ToFloat64() float64 {
	output, err := strconv.ParseFloat(v.Value, 64)
	if err != nil {
		panic("env: value is not an float64.")
	}
	return output
}

func (v *Value) ToFloat64s() []float64 {
	var output []float64
	for _, v := range v.ToStrings() {
		output = append(output, NewValue(v).ToFloat64())
	}
	return output
}

func (v *Value) ToBool() bool {
	output, err := strconv.ParseBool(v.Value)
	if err != nil {
		panic("env: value is not an bool.")
	}
	return output
}

func (v *Value) ToBools() []bool {
	var output []bool
	for _, v := range v.ToStrings() {
		output = append(output, NewValue(v).ToBool())
	}
	return output
}

func (v *Value) ToByte() byte {
	output, err := strconv.ParseUint(v.Value, 10, 8)
	if err != nil {
		panic("env: value is not an byte.")
	}
	return byte(output)
}

func (v *Value) ToBytes() []byte {
	var output []byte
	for _, v := range v.ToStrings() {
		output = append(output, NewValue(v).ToByte())
	}
	return output
}

func (v *Value) ToRune() rune {
	output, err := strconv.ParseInt(v.Value, 10, 32)
	if err != nil {
		panic("env: value is not an rune.")
	}
	return rune(output)
}

func (v *Value) ToRunes() []rune {
	var output []rune
	for _, v := range v.ToStrings() {
		output = append(output, NewValue(v).ToRune())
	}
	return output
}

func (v *Value) ToComplex64() complex64 {
	output, err := strconv.ParseComplex(v.Value, 64)
	if err != nil {
		panic("env: value is not an complex64.")
	}
	return complex64(output)
}

func (v *Value) ToComplex64s() []complex64 {
	var output []complex64
	for _, v := range v.ToStrings() {
		output = append(output, NewValue(v).ToComplex64())
	}
	return output
}

func (v *Value) ToComplex128() complex128 {
	output, err := strconv.ParseComplex(v.Value, 128)
	if err != nil {
		panic("env: value is not an complex128.")
	}
	return complex128(output)
}

func (v *Value) ToComplex128s() []complex128 {
	var output []complex128
	for _, v := range v.ToStrings() {
		output = append(output, NewValue(v).ToComplex128())
	}
	return output
}
