package xtemplate

import (
	"errors"
	"math"
	"slices"

	"github.com/Eun/xtemplate/funcs"
)

// Slice provides helper functions for slices.
type Slice rootContext

// ErrCannotSortAnySlice is returned when trying to sort a []any slice.
var ErrCannotSortAnySlice = errors.New("cannot sort []any slices")

// ErrArgNotSlice is returned when the argument provided is not a slice.
var ErrArgNotSlice = errors.New("argument must be a slice")

// ErrFirstArgumentMustBeSlice is returned when the first argument provided is not a slice.
var ErrFirstArgumentMustBeSlice = errors.New("first argument must be a slice")

// ErrCannotCompactAnySlice is returned when trying to compact a []any slice.
var ErrCannotCompactAnySlice = errors.New("cannot compact []any slices")

// New creates a slice from the provided values.
//
// Example:
//
//	{{ slice.New 1 "Hello" false }}
func (ctx Slice) New(vals ...any) ([]any, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.SliceNew]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.SliceNew}
	}

	return vals, nil
}

// NewStrings creates a string slice from the provided values.
//
// Example:
//
//	{{ slice.NewStrings "Hello" "World" }}
func (ctx Slice) NewStrings(vals ...any) ([]string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.SliceNewStrings]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.SliceNewStrings}
	}

	return toStrings(vals), nil
}

// NewInts creates an int slice from the provided values.
//
// Example:
//
//	{{ slice.NewInts 1 2 }}
func (ctx Slice) NewInts(vals ...any) ([]int, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.SliceNewInts]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.SliceNewInts}
	}

	return toInts[int](vals, math.MinInt, math.MaxInt)
}

// NewInt64s creates an int64 slice from the provided values.
//
// Example:
//
//	{{ slice.NewInt64s 1 2 }}
func (ctx Slice) NewInt64s(vals ...any) ([]int64, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.SliceNewInt64s]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.SliceNewInt64s}
	}

	return toInt64s(vals)
}

// NewFloat64s creates an int64 slice from the provided values.
//
// Example:
//
//	{{ slice.NewFloat64s 1.5 2.1 }}
func (ctx Slice) NewFloat64s(vals ...any) ([]float64, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.SliceNewFloat64s]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.SliceNewFloat64s}
	}

	return toFloat64s(vals)
}

// NewBools creates an int64 slice from the provided values.
//
// Example:
//
//	{{ slice.NewBools false true }}
func (ctx Slice) NewBools(vals ...any) ([]bool, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.SliceNewBools]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.SliceNewBools}
	}

	return toBools(vals), nil
}

// Contains checks if the slice contains the provided value.
//
// Example:
//
//	{{ $sl := slice.NewStrings "Hello" "World" }}
//	{{ slice.Contains $sl "World" }} // Output: true
//
//nolint:cyclop // cannot be simplified
func (ctx Slice) Contains(s any, v any) (bool, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.SliceContains]; !ok {
		return false, &FuncNotAllowedError{Func: funcs.SliceContains}
	}
	switch sl := s.(type) {
	case []any:
		for i := range sl {
			if v == sl[i] {
				return true, nil
			}
		}
		return false, nil
	case []bool:
		return ctx.Contains(toAnySlice(sl), v)
	case []float32:
		return ctx.Contains(toAnySlice(sl), v)
	case []float64:
		return ctx.Contains(toAnySlice(sl), v)
	case []string:
		return ctx.Contains(toAnySlice(sl), v)
	case []int:
		return ctx.Contains(toAnySlice(sl), v)
	case []int8:
		return ctx.Contains(toAnySlice(sl), v)
	case []int16:
		return ctx.Contains(toAnySlice(sl), v)
	case []int32:
		return ctx.Contains(toAnySlice(sl), v)
	case []int64:
		return ctx.Contains(toAnySlice(sl), v)
	case []uint8:
		return ctx.Contains(toAnySlice(sl), v)
	case []uint16:
		return ctx.Contains(toAnySlice(sl), v)
	case []uint32:
		return ctx.Contains(toAnySlice(sl), v)
	case []uint64:
		return ctx.Contains(toAnySlice(sl), v)
	}
	return false, ErrFirstArgumentMustBeSlice
}

func toAnySlice[T any](in []T) []any {
	n := make([]any, len(in))
	for i := range in {
		n[i] = in[i]
	}
	return n
}

// Reverse reverses the order of elements in the provided slice.
//
// Example:
//
//	{{ slice.Reverse ( slice.NewStrings "Hello" "World" ) }} // Output: [World Hello]
//
//nolint:cyclop // cannot be simplified
func (ctx Slice) Reverse(s any) (any, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.SliceReverse]; !ok {
		return false, &FuncNotAllowedError{Func: funcs.SliceReverse}
	}
	switch sl := s.(type) {
	case []any:
		sl = slices.Clone(sl)
		slices.Reverse(sl)
		return sl, nil
	case []bool:
		return ctx.Reverse(toAnySlice(sl))
	case []float32:
		return ctx.Reverse(toAnySlice(sl))
	case []float64:
		return ctx.Reverse(toAnySlice(sl))
	case []string:
		return ctx.Reverse(toAnySlice(sl))
	case []int:
		return ctx.Reverse(toAnySlice(sl))
	case []int8:
		return ctx.Reverse(toAnySlice(sl))
	case []int16:
		return ctx.Reverse(toAnySlice(sl))
	case []int32:
		return ctx.Reverse(toAnySlice(sl))
	case []int64:
		return ctx.Reverse(toAnySlice(sl))
	case []uint8:
		return ctx.Reverse(toAnySlice(sl))
	case []uint16:
		return ctx.Reverse(toAnySlice(sl))
	case []uint32:
		return ctx.Reverse(toAnySlice(sl))
	case []uint64:
		return ctx.Reverse(toAnySlice(sl))
	}
	return false, ErrArgNotSlice
}

// Sort sorts the provided slice.
//
// Example:
//
//	{{ slice.Sort ( slice.NewStrings "World" "Hello" ) }} // Output: [Hello World]
//
//nolint:cyclop, funlen // cannot be simplified
func (ctx Slice) Sort(s any) (any, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.SliceSort]; !ok {
		return false, &FuncNotAllowedError{Func: funcs.SliceSort}
	}
	switch sl := s.(type) {
	case []any:
		return nil, ErrCannotSortAnySlice
	case []bool:
		sl = slices.Clone(sl)
		sortBool(sl)
		return sl, nil
	case []float32:
		sl = slices.Clone(sl)
		slices.Sort(sl)
		return sl, nil
	case []float64:
		sl = slices.Clone(sl)
		slices.Sort(sl)
		return sl, nil
	case []string:
		sl = slices.Clone(sl)
		slices.Sort(sl)
		return sl, nil
	case []int:
		sl = slices.Clone(sl)
		slices.Sort(sl)
		return sl, nil
	case []int8:
		sl = slices.Clone(sl)
		slices.Sort(sl)
		return sl, nil
	case []int16:
		sl = slices.Clone(sl)
		slices.Sort(sl)
		return sl, nil
	case []int32:
		sl = slices.Clone(sl)
		slices.Sort(sl)
		return sl, nil
	case []int64:
		sl = slices.Clone(sl)
		slices.Sort(sl)
		return sl, nil
	case []uint8:
		sl = slices.Clone(sl)
		slices.Sort(sl)
		return sl, nil
	case []uint16:
		sl = slices.Clone(sl)
		slices.Sort(sl)
		return sl, nil
	case []uint32:
		sl = slices.Clone(sl)
		slices.Sort(sl)
		return sl, nil
	case []uint64:
		sl = slices.Clone(sl)
		slices.Sort(sl)
		return sl, nil
	}
	return false, ErrArgNotSlice
}

// Append appends the provided values to the slice.
//
// Example:
//
//	{{ $sl := slice.NewStrings "Joe" }}
//	{{ slice.Append $sl "Alice" "Bob" }} // Output: [Joe Alice Bob]
//
//nolint:cyclop, funlen // cannot be simplified
func (ctx Slice) Append(s any, vals ...any) (any, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.SliceAppend]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.SliceAppend}
	}
	switch sl := s.(type) {
	case []any:
		return append(sl, vals...), nil
	case []bool:
		return append(sl, toBools(vals)...), nil
	case []float32:
		v, err := toFloat32s(vals)
		if err != nil {
			return nil, err
		}
		return append(sl, v...), nil
	case []float64:
		v, err := toFloat64s(vals)
		if err != nil {
			return nil, err
		}
		return append(sl, v...), nil
	case []string:
		return append(sl, toStrings(vals)...), nil
	case []int:
		v, err := toInts[int](vals, math.MinInt, math.MaxInt)
		if err != nil {
			return nil, err
		}
		return append(sl, v...), nil
	case []int8:
		v, err := toInts[int8](vals, math.MinInt8, math.MaxInt8)
		if err != nil {
			return nil, err
		}
		return append(sl, v...), nil
	case []int16:
		v, err := toInts[int16](vals, math.MinInt16, math.MaxInt16)
		if err != nil {
			return nil, err
		}
		return append(sl, v...), nil
	case []int32:
		v, err := toInts[int32](vals, math.MinInt32, math.MaxInt32)
		if err != nil {
			return nil, err
		}
		return append(sl, v...), nil
	case []int64:
		v, err := toInt64s(vals)
		if err != nil {
			return nil, err
		}
		return append(sl, v...), nil
	case []uint8:
		v, err := toUints[uint8](vals, math.MaxUint8)
		if err != nil {
			return nil, err
		}
		return append(sl, v...), nil
	case []uint16:
		v, err := toUints[uint16](vals, math.MaxUint16)
		if err != nil {
			return nil, err
		}
		return append(sl, v...), nil
	case []uint32:
		v, err := toUints[uint32](vals, math.MaxUint32)
		if err != nil {
			return nil, err
		}
		return append(sl, v...), nil
	case []uint64:
		v, err := toUint64s(vals)
		if err != nil {
			return nil, err
		}
		return append(sl, v...), nil
	}
	return nil, ErrFirstArgumentMustBeSlice
}

// Prepend appends the provided values to the slice.
//
// Example:
//
//	{{ $sl := slice.NewStrings "Joe" }}
//	{{ slice.Prepend $sl "Alice" "Bob" }} // Output: [Alice Bob Joe]
//
//nolint:cyclop, funlen // cannot be simplified
func (ctx Slice) Prepend(s any, vals ...any) (any, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.SlicePrepend]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.SlicePrepend}
	}
	switch sl := s.(type) {
	case []any:
		return append(vals, sl...), nil
	case []bool:
		return append(toBools(vals), sl...), nil
	case []float32:
		v, err := toFloat32s(vals)
		if err != nil {
			return nil, err
		}
		return append(v, sl...), nil
	case []float64:
		v, err := toFloat64s(vals)
		if err != nil {
			return nil, err
		}
		return append(v, sl...), nil
	case []string:
		return append(toStrings(vals), sl...), nil
	case []int:
		v, err := toInts[int](vals, math.MinInt, math.MaxInt)
		if err != nil {
			return nil, err
		}
		return append(v, sl...), nil
	case []int8:
		v, err := toInts[int8](vals, math.MinInt8, math.MaxInt8)
		if err != nil {
			return nil, err
		}
		return append(v, sl...), nil
	case []int16:
		v, err := toInts[int16](vals, math.MinInt16, math.MaxInt16)
		if err != nil {
			return nil, err
		}
		return append(v, sl...), nil
	case []int32:
		v, err := toInts[int32](vals, math.MinInt32, math.MaxInt32)
		if err != nil {
			return nil, err
		}
		return append(v, sl...), nil
	case []int64:
		v, err := toInt64s(vals)
		if err != nil {
			return nil, err
		}
		return append(v, sl...), nil
	case []uint8:
		v, err := toUints[uint8](vals, math.MaxUint8)
		if err != nil {
			return nil, err
		}
		return append(v, sl...), nil
	case []uint16:
		v, err := toUints[uint16](vals, math.MaxUint16)
		if err != nil {
			return nil, err
		}
		return append(v, sl...), nil
	case []uint32:
		v, err := toUints[uint32](vals, math.MaxUint32)
		if err != nil {
			return nil, err
		}
		return append(v, sl...), nil
	case []uint64:
		v, err := toUint64s(vals)
		if err != nil {
			return nil, err
		}
		return append(v, sl...), nil
	}
	return nil, ErrFirstArgumentMustBeSlice
}

// Len returns the length of the provided slice.
//
// Example:
//
//	{{ $sl := slice.NewStrings "Hello" "World" }}
//	{{ slice.Len $sl }} // Output: 2
//
//nolint:cyclop // cannot be simplified
func (ctx Slice) Len(s any) (int, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.SliceLen]; !ok {
		return 0, &FuncNotAllowedError{Func: funcs.SliceLen}
	}
	switch sl := s.(type) {
	case []any:
		return len(sl), nil
	case []bool:
		return len(sl), nil
	case []float32:
		return len(sl), nil
	case []float64:
		return len(sl), nil
	case []string:
		return len(sl), nil
	case []int:
		return len(sl), nil
	case []int8:
		return len(sl), nil
	case []int16:
		return len(sl), nil
	case []int32:
		return len(sl), nil
	case []int64:
		return len(sl), nil
	case []uint8:
		return len(sl), nil
	case []uint16:
		return len(sl), nil
	case []uint32:
		return len(sl), nil
	case []uint64:
		return len(sl), nil
	}
	return 0, ErrArgNotSlice
}

func uniqSlice[T comparable](sl []T) []T {
	seen := make(map[T]struct{}, len(sl))
	result := make([]T, 0, len(sl))
	for _, v := range sl {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

// Unique removes duplicate elements from the provided slice.
// Example:
//
//	{{ slice.Unique ( slice.NewStrings "Hello" "World" "Hello" ) }} // Output: [Hello World]
//
//nolint:cyclop // cannot be simplified
func (ctx Slice) Unique(s any) (any, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.SliceUnique]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.SliceUnique}
	}
	switch sl := s.(type) {
	case []any:
		return uniqSlice(sl), nil
	case []bool:
		return uniqSlice(sl), nil
	case []float32:
		return uniqSlice(sl), nil
	case []float64:
		return uniqSlice(sl), nil
	case []string:
		return uniqSlice(sl), nil
	case []int:
		return uniqSlice(sl), nil
	case []int8:
		return uniqSlice(sl), nil
	case []int16:
		return uniqSlice(sl), nil
	case []int32:
		return uniqSlice(sl), nil
	case []int64:
		return uniqSlice(sl), nil
	case []uint8:
		return uniqSlice(sl), nil
	case []uint16:
		return uniqSlice(sl), nil
	case []uint32:
		return uniqSlice(sl), nil
	case []uint64:
		return uniqSlice(sl), nil
	}
	return nil, ErrArgNotSlice
}

// Compact replaces consecutive runs of equal elements with a single copy.
// This is like the uniq command found on Unix.
//
// Example:
//
//	{{ slice.Compact ( slice.NewStrings "Hello" "Hello" "World" "World" ) }} // Output: [Hello World]
//
//nolint:cyclop // cannot be simplified
func (ctx Slice) Compact(s any) (any, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.SliceCompact]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.SliceCompact}
	}
	switch sl := s.(type) {
	case []any:
		return nil, ErrCannotCompactAnySlice
	case []bool:
		return slices.Compact(slices.Clone(sl)), nil
	case []float32:
		return slices.Compact(slices.Clone(sl)), nil
	case []float64:
		return slices.Compact(slices.Clone(sl)), nil
	case []string:
		return slices.Compact(slices.Clone(sl)), nil
	case []int:
		return slices.Compact(slices.Clone(sl)), nil
	case []int8:
		return slices.Compact(slices.Clone(sl)), nil
	case []int16:
		return slices.Compact(slices.Clone(sl)), nil
	case []int32:
		return slices.Compact(slices.Clone(sl)), nil
	case []int64:
		return slices.Compact(slices.Clone(sl)), nil
	case []uint8:
		return slices.Compact(slices.Clone(sl)), nil
	case []uint16:
		return slices.Compact(slices.Clone(sl)), nil
	case []uint32:
		return slices.Compact(slices.Clone(sl)), nil
	case []uint64:
		return slices.Compact(slices.Clone(sl)), nil
	}
	return nil, ErrArgNotSlice
}

func sortBool(sl []bool) {
	slices.SortFunc(sl, func(a, b bool) int {
		if a == b {
			return 0
		}
		if !a && b {
			return -1
		}
		return 1
	})
}
