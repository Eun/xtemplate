package xtemplate

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"

	"github.com/Eun/xtemplate/funcs"
)

// Conv provides functions to convert between types.
type Conv rootContext

func toBool(in any) bool {
	if b, ok := in.(bool); ok {
		return b
	}

	if str, ok := in.(string); ok {
		str = strings.ToLower(str)
		switch str {
		case "1", "t", "true", "yes":
			return true
		default:
			// ignore error here, as we'll just return false
			f, _ := strToFloat64(str)
			return f == 1
		}
	}

	val := reflect.Indirect(reflect.ValueOf(in))
	switch val.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return val.Int() == 1
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		return val.Uint() == 1
	case reflect.Float32, reflect.Float64:
		return val.Float() == 1
	default:
		return false
	}
}

func toBools(in []any) []bool {
	out := make([]bool, len(in))
	for i, v := range in {
		out[i] = toBool(v)
	}
	return out
}

// ToBool converts various types to bool.
//
// Example:
//
//	{{ conv.ToBool "true" }} // Output: true
func (ctx Conv) ToBool(in any) (bool, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.ConvToBool]; !ok {
		return false, &FuncNotAllowedError{Func: funcs.ConvToBool}
	}
	return toBool(in), nil
}

// ToBools converts a list of various types to bools.
//
// Example:
//
//	{{ $sl := slice.New "true" "false" 1 0 "yes" "no" }}
//	{{ conv.ToBools $sl }} // Output: [true false true false true false]
func (ctx Conv) ToBools(in []any) ([]bool, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.ConvToBools]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.ConvToBools}
	}
	return toBools(in), nil
}

func toString(in any) string {
	if in == nil {
		return ""
	}
	if s, ok := in.(string); ok {
		return s
	}
	if s, ok := in.(fmt.Stringer); ok {
		return s.String()
	}
	if s, ok := in.([]byte); ok {
		return string(s)
	}

	return fmt.Sprint(in)
}

func toStrings(in []any) []string {
	out := make([]string, len(in))
	for i, v := range in {
		out[i] = toString(v)
	}
	return out
}

// ToString converts various types to string.
//
// Example:
//
//	{{ conv.ToString 42 }} // Output: 42
func (ctx Conv) ToString(in any) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.ConvToString]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.ConvToString}
	}

	return toString(in), nil
}

// ToStrings converts a list of various types to strings.
//
// Example:
//
//	{{ $sl := slice.New 42 true 3.14 }}
//	{{ conv.ToStrings $sl }} // Output: [42 true 3.14]
func (ctx Conv) ToStrings(in []any) ([]string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.ConvToStrings]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.ConvToStrings}
	}
	return toStrings(in), nil
}

func strToFloat64(str string) (float64, error) {
	if strings.Contains(str, ",") {
		str = strings.ReplaceAll(str, ",", "")
	}

	// this is inefficient, but it's the only way I can think of to
	// properly convert octal integers to floats
	iv, err := strconv.ParseInt(str, 0, 64)
	if err != nil {
		// ok maybe it's a float?
		var fv float64
		fv, err = strconv.ParseFloat(str, 64)
		if err != nil {
			return 0, fmt.Errorf("could not convert %q to float64: %w", str, err)
		}

		return fv, nil
	}

	return float64(iv), nil
}

func toFloat64(v any) (float64, error) {
	if str, ok := v.(string); ok {
		return strToFloat64(str)
	}

	val := reflect.Indirect(reflect.ValueOf(v))
	switch val.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return float64(val.Int()), nil
	case reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return float64(val.Uint()), nil
	case reflect.Uint, reflect.Uint64:
		return float64(val.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return val.Float(), nil
	case reflect.Bool:
		if val.Bool() {
			return 1, nil
		}
		return 0, nil
	default:
		//nolint:err113 // allow dynamic error
		return 0, fmt.Errorf("could not convert %v to float64", v)
	}
}

func toFloat64s(in []any) ([]float64, error) {
	out := make([]float64, len(in))
	var err error
	for i, v := range in {
		out[i], err = toFloat64(v)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

// ToFloat64 converts various types to float64.
//
// Example:
//
//	{{ conv.ToFloat64 "3.14" }} // Output: 3.14
func (ctx Conv) ToFloat64(v any) (float64, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.ConvToFloat64]; !ok {
		return 0, &FuncNotAllowedError{Func: funcs.ConvToFloat64}
	}
	return toFloat64(v)
}

// ToFloat64s converts a list of various types to float64s.
//
// Example:
//
//	{{ $sl := slice.New "3.14" 42 "1e10" }}
//	{{ conv.ToFloat64s $sl }} // Output: [3.14 42 1e+10]
func (ctx Conv) ToFloat64s(in []any) ([]float64, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.ConvToFloat64s]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.ConvToFloat64s}
	}

	return toFloat64s(in)
}

func toFloat32(v any) (float32, error) {
	f, err := toFloat64(v)
	if err != nil {
		return 0, err
	}
	if f > math.MaxFloat32 || f < -math.MaxFloat32 {
		//nolint:err113 // allow dynamic error
		return 0, fmt.Errorf("could not convert %f to float32, would overflow", f)
	}
	return float32(f), nil
}

func toFloat32s(in []any) ([]float32, error) {
	out := make([]float32, len(in))
	var err error
	for i, v := range in {
		out[i], err = toFloat32(v)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

// ToFloat32 converts various types to float32.
//
// Example:
//
//	{{ conv.ToFloat32 "3.14" }} // Output: 3.14
func (ctx Conv) ToFloat32(v any) (float32, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.ConvToFloat32]; !ok {
		return 0, &FuncNotAllowedError{Func: funcs.ConvToFloat32}
	}
	return toFloat32(v)
}

// ToFloat32s converts a list of various types to float32s.
//
// Example:
//
//	{{ $sl := slice.New "3.14" 42 "1e10" }}
//	{{ conv.ToFloat32s $sl }} // Output: [3.14 42 1e+10]
func (ctx Conv) ToFloat32s(in []any) ([]float32, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.ConvToFloat32s]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.ConvToFloat32s}
	}

	return toFloat32s(in)
}

func strToInt64(str string) (int64, error) {
	if strings.Contains(str, ",") {
		str = strings.ReplaceAll(str, ",", "")
	}

	iv, err := strconv.ParseInt(str, 0, 64)
	if err != nil {
		// maybe it's a float?
		var fv float64
		fv, err = strconv.ParseFloat(str, 64)
		if err != nil {
			return 0, fmt.Errorf("could not convert %q to int64: %w", str, err)
		}

		return int64(fv), nil
	}

	return iv, nil
}

func toInt64(v any) (int64, error) {
	if str, ok := v.(string); ok {
		return strToInt64(str)
	}

	val := reflect.Indirect(reflect.ValueOf(v))
	switch val.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return val.Int(), nil
	case reflect.Uint8, reflect.Uint16, reflect.Uint32:
		//nolint:gosec // G115 isn't applicable, this is a Uint32 at most
		return int64(val.Uint()), nil
	case reflect.Uint, reflect.Uint64:
		tv := val.Uint()

		if tv > math.MaxInt64 {
			//nolint:err113 // allow dynamic error
			return 0, fmt.Errorf("could not convert %d to int64, would overflow", tv)
		}

		return int64(tv), nil
	case reflect.Float32, reflect.Float64:
		return int64(val.Float()), nil
	case reflect.Bool:
		if val.Bool() {
			return 1, nil
		}

		return 0, nil
	default:
		//nolint:err113 // allow dynamic error
		return 0, fmt.Errorf("could not convert %v to int64", v)
	}
}

func toInt64s(in []any) ([]int64, error) {
	out := make([]int64, len(in))
	var err error
	for i, v := range in {
		out[i], err = toInt64(v)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

// ToInt64 converts various types to int64.
//
// Example:
//
//	{{ conv.ToInt64 "42" }} // Output: 42
func (ctx Conv) ToInt64(v any) (int64, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.ConvToInt64]; !ok {
		return 0, &FuncNotAllowedError{Func: funcs.ConvToInt64}
	}
	i, err := toInt64(v)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// ToInt64s converts a list of various types to int64s.
// // Example:
//
//	{{ $sl := slice.New "42" 7 "0x10" }}
//	{{ conv.ToInt64s $sl }} // Output: [42 7 16]
func (ctx Conv) ToInt64s(in []any) ([]int64, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.ConvToInt64s]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.ConvToInt64s}
	}

	return toInt64s(in)
}

func toInt[T int | int8 | int16 | int32](v any, minValue, maxValue int64) (T, error) {
	var zero T
	i64, err := toInt64(v)
	if err != nil {
		return zero, err
	}

	if i64 < minValue || i64 > maxValue {
		//nolint:err113 // allow dynamic error
		return zero, fmt.Errorf("could not convert %d, would overflow", i64)
	}

	return T(i64), nil
}

func toInts[T int | int8 | int16 | int32](in []any, minValue, maxValue int64) ([]T, error) {
	out := make([]T, len(in))
	var err error
	for i, v := range in {
		out[i], err = toInt[T](v, minValue, maxValue)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

// ToInt8 converts various types to int8.
//
// Example:
//
//	{{ conv.ToInt8 "42" }} // Output: 42
func (ctx Conv) ToInt8(v any) (int8, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.ConvToInt8]; !ok {
		return 0, &FuncNotAllowedError{Func: funcs.ConvToInt8}
	}
	i, err := toInt[int8](v, math.MinInt8, math.MaxInt8)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// ToInt8s converts a list of various types to int8s.
//
// Example:
//
//	{{ $sl := slice.New "42" 7 "0x10" }}
//	{{ conv.ToInt8s $sl }} // Output: [42 7 16]
func (ctx Conv) ToInt8s(in []any) ([]int8, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.ConvToInt8s]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.ConvToInt8s}
	}
	return toInts[int8](in, math.MinInt8, math.MaxInt8)
}

// ToInt16 converts various types to int16.
//
// Example:
//
//	{{ conv.ToInt16 "42" }} // Output: 42
func (ctx Conv) ToInt16(v any) (int16, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.ConvToInt16]; !ok {
		return 0, &FuncNotAllowedError{Func: funcs.ConvToInt16}
	}
	i, err := toInt[int16](v, math.MinInt16, math.MaxInt16)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// ToInt16s converts a list of various types to int16s.
//
// Example:
//
//	{{ $sl := slice.New "42" 7 "0x10" }}
//	{{ conv.ToInt16s $sl }} // Output: [42 7 16]
func (ctx Conv) ToInt16s(in []any) ([]int16, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.ConvToInt16s]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.ConvToInt16s}
	}
	return toInts[int16](in, math.MinInt16, math.MaxInt16)
}

// ToInt32 converts various types to int32.
//
// Example:
//
//	{{ conv.ToInt32 "42" }} // Output: 42
func (ctx Conv) ToInt32(v any) (int32, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.ConvToInt32]; !ok {
		return 0, &FuncNotAllowedError{Func: funcs.ConvToInt32}
	}
	i, err := toInt[int32](v, math.MinInt32, math.MaxInt32)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// ToInt32s converts a list of various types to int32s.
//
// Example:
//
//	{{ $sl := slice.New "42" 7 "0x10" }}
//	{{ conv.ToInt32s $sl }} // Output: [42 7 16]
func (ctx Conv) ToInt32s(in []any) ([]int32, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.ConvToInt32s]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.ConvToInt32s}
	}
	return toInts[int32](in, math.MinInt32, math.MaxInt32)
}

// ToInt converts various types to int.
//
// Example:
//
//	{{ conv.ToInt "42" }} // Output: 42
func (ctx Conv) ToInt(v any) (int, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.ConvToInt]; !ok {
		return 0, &FuncNotAllowedError{Func: funcs.ConvToInt}
	}
	i, err := toInt[int](v, math.MinInt, math.MaxInt)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// ToInts converts a list of various types to ints.
//
// Example:
//
//	{{ $sl := slice.New "42" 7 "0x10" }}
//	{{ conv.ToInts $sl }} // Output: [42 7 16]
func (ctx Conv) ToInts(in []any) ([]int, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.ConvToInts]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.ConvToInts}
	}
	return toInts[int](in, math.MinInt, math.MaxInt)
}

func strToUint64(str string) (uint64, error) {
	if strings.Contains(str, ",") {
		str = strings.ReplaceAll(str, ",", "")
	}

	iv, err := strconv.ParseUint(str, 0, 64)
	if err != nil {
		// maybe it's a float?
		var fv float64
		fv, err = strconv.ParseFloat(str, 64)
		if err != nil {
			return 0, fmt.Errorf("could not convert %q to uint64: %w", str, err)
		}

		return uint64(fv), nil
	}

	return iv, nil
}

func toUint64(v any) (uint64, error) {
	if str, ok := v.(string); ok {
		return strToUint64(str)
	}

	val := reflect.Indirect(reflect.ValueOf(v))
	switch val.Kind() {
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		return val.Uint(), nil
	case reflect.Int8, reflect.Int16, reflect.Int32:
		//nolint:gosec // G115 isn't applicable, this is a Int32 at most
		return uint64(val.Int()), nil
	case reflect.Int, reflect.Int64:
		tv := val.Int()

		if tv < 0 {
			//nolint:err113 // allow dynamic error
			return 0, fmt.Errorf("could not convert %d to int64, would overflow", tv)
		}

		return uint64(tv), nil
	case reflect.Float32, reflect.Float64:
		return uint64(val.Float()), nil
	case reflect.Bool:
		if val.Bool() {
			return 1, nil
		}

		return 0, nil
	default:
		//nolint:err113 // allow dynamic error
		return 0, fmt.Errorf("could not convert %v to int64", v)
	}
}

func toUint64s(in []any) ([]uint64, error) {
	out := make([]uint64, len(in))
	var err error
	for i, v := range in {
		out[i], err = toUint64(v)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

// ToUint64 converts various types to uint64.
//
// Example:
//
//	{{ conv.ToUint64 "42" }} // Output: 42
func (ctx Conv) ToUint64(v any) (uint64, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.ConvToUint64]; !ok {
		return 0, &FuncNotAllowedError{Func: funcs.ConvToUint64}
	}
	i, err := toUint64(v)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// ToUint64s converts a list of various types to uint64s.
// // Example:
//
//	{{ $sl := slice.New "42" 7 "0x10" }}
//	{{ conv.ToUint64s $sl }} // Output: [42 7 16]
func (ctx Conv) ToUint64s(in []any) ([]uint64, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.ConvToUint64s]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.ConvToUint64s}
	}

	return toUint64s(in)
}

func toUint[T uint | uint8 | uint16 | uint32](v any, maxValue uint64) (T, error) {
	var zero T
	i64, err := toUint64(v)
	if err != nil {
		return zero, err
	}

	if i64 > maxValue {
		//nolint:err113 // allow dynamic error
		return zero, fmt.Errorf("could not convert %d, would overflow", i64)
	}

	return T(i64), nil
}

func toUints[T uint | uint8 | uint16 | uint32](in []any, maxValue uint64) ([]T, error) {
	out := make([]T, len(in))
	var err error
	for i, v := range in {
		out[i], err = toUint[T](v, maxValue)
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

// ToUint8 converts various types to uint8.
//
// Example:
//
//	{{ conv.ToUint8 "42" }} // Output: 42
func (ctx Conv) ToUint8(v any) (uint8, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.ConvToUint8]; !ok {
		return 0, &FuncNotAllowedError{Func: funcs.ConvToUint8}
	}
	i, err := toUint[uint8](v, math.MaxUint8)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// ToUint8s converts a list of various types to uint8s.
//
// Example:
//
//	{{ $sl := slice.New "42" 7 "0x10" }}
//	{{ conv.ToUint8s $sl }} // Output: [42 7 16]
func (ctx Conv) ToUint8s(in []any) ([]uint8, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.ConvToUint8s]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.ConvToUint8s}
	}
	return toUints[uint8](in, math.MaxUint8)
}

// ToUint16 converts various types to uint16.
//
// Example:
//
//	{{ conv.ToUint16 "42" }} // Output: 42
func (ctx Conv) ToUint16(v any) (uint16, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.ConvToUint16]; !ok {
		return 0, &FuncNotAllowedError{Func: funcs.ConvToUint16}
	}
	i, err := toUint[uint16](v, math.MaxUint16)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// ToUint16s converts a list of various types to uint16s.
//
// Example:
//
//	{{ $sl := slice.New "42" 7 "0x10" }}
//	{{ conv.ToUint16s $sl }} // Output: [42 7 16]
func (ctx Conv) ToUint16s(in []any) ([]uint16, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.ConvToUint16s]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.ConvToUint16s}
	}
	return toUints[uint16](in, math.MaxUint16)
}

// ToUint32 converts various types to uint32.
//
// Example:
//
//	{{ conv.ToUint32 "42" }} // Output: 42
func (ctx Conv) ToUint32(v any) (uint32, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.ConvToUint32]; !ok {
		return 0, &FuncNotAllowedError{Func: funcs.ConvToUint32}
	}
	i, err := toUint[uint32](v, math.MaxUint32)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// ToUint32s converts a list of various types to uint32s.
//
// Example:
//
//	{{ $sl := slice.New "42" 7 "0x10" }}
//	{{ conv.ToUint32s $sl }} // Output: [42 7 16]
func (ctx Conv) ToUint32s(in []any) ([]uint32, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.ConvToUint32s]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.ConvToUint32s}
	}
	return toUints[uint32](in, math.MaxUint32)
}

// ToUint converts various types to uint.
//
// Example:
//
//	{{ conv.ToUint "42" }} // Output: 42
func (ctx Conv) ToUint(v any) (uint, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.ConvToUint]; !ok {
		return 0, &FuncNotAllowedError{Func: funcs.ConvToUint}
	}
	i, err := toUint[uint](v, math.MaxUint)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// ToUints converts a list of various types to uints.
//
// Example:
//
//	{{ $sl := slice.New "42" 7 "0x10" }}
//	{{ conv.ToUints $sl }} // Output: [42 7 16]
func (ctx Conv) ToUints(in []any) ([]uint, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.ConvToUints]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.ConvToUints}
	}
	return toUints[uint](in, math.MaxUint)
}
