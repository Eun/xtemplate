package xtemplate

import (
	"cmp"
	"errors"
	"math"

	"github.com/Eun/xtemplate/funcs"
)

// Cmp provides access to functions in the cmp package.
type Cmp rootContext

// ErrAtLeastOneArgumentIsRequired is returned when no arguments are provided to a function that requires
// at least one argument.
var ErrAtLeastOneArgumentIsRequired = errors.New("at least one argument is required")

// ErrOnlyOneSliceIsAllowed is returned when more than one slice argument is provided to a function that only
// allows one slice argument.
var ErrOnlyOneSliceIsAllowed = errors.New("only one slice argument is allowed")

// Or is a logical OR operator that returns the first non-zero value from the provided arguments.
//
// Example 1:
//
//	{{ cmp.Or "" "Hello" "World" }} // Output: Hello
//
// Example 2:
//
//	{{ cmp.Or 0 1 2 }} // Output: 1
//
// Example 3:
//
//	{{ cmp.Or ( slice.NewStrings "" "Hello" "World" ) }} // Output: Hello
//
//nolint:gocognit, gocyclo, cyclop, funlen // cannot be simplified
func (ctx Cmp) Or(s ...any) (any, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.CmpOr]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.CmpOr}
	}

	if len(s) == 0 {
		return nil, ErrAtLeastOneArgumentIsRequired
	}

	switch sl := s[0].(type) {
	case bool:
		return cmp.Or(toBools(s)...), nil
	case float32:
		v, err := toFloat32s(s)
		if err != nil {
			return nil, err
		}
		return cmp.Or(v...), nil
	case float64:
		v, err := toFloat64s(s)
		if err != nil {
			return nil, err
		}
		return cmp.Or(v...), nil
	case string:
		return cmp.Or(toStrings(s)...), nil
	case int:
		v, err := toInts[int](s, math.MinInt, math.MaxInt)
		if err != nil {
			return nil, err
		}
		return cmp.Or(v...), nil
	case int8:
		v, err := toInts[int8](s, math.MinInt8, math.MaxInt8)
		if err != nil {
			return nil, err
		}
		return cmp.Or(v...), nil
	case int16:
		v, err := toInts[int16](s, math.MinInt16, math.MaxInt16)
		if err != nil {
			return nil, err
		}
		return cmp.Or(v...), nil
	case int32:
		v, err := toInts[int32](s, math.MinInt32, math.MaxInt32)
		if err != nil {
			return nil, err
		}
		return cmp.Or(v...), nil
	case int64:
		v, err := toInt64s(s)
		if err != nil {
			return nil, err
		}
		return cmp.Or(v...), nil
	case uint8:
		v, err := toUints[uint8](s, math.MaxUint8)
		if err != nil {
			return nil, err
		}
		return cmp.Or(v...), nil
	case uint16:
		v, err := toUints[uint16](s, math.MaxUint16)
		if err != nil {
			return nil, err
		}
		return cmp.Or(v...), nil
	case uint32:
		v, err := toUints[uint32](s, math.MaxUint32)
		if err != nil {
			return nil, err
		}
		return cmp.Or(v...), nil
	case uint64:
		v, err := toUint64s(s)
		if err != nil {
			return nil, err
		}
		return cmp.Or(v...), nil
	case []bool:
		if len(s) != 1 {
			return nil, ErrOnlyOneSliceIsAllowed
		}
		return cmp.Or(sl...), nil
	case []float32:
		if len(s) != 1 {
			return nil, ErrOnlyOneSliceIsAllowed
		}
		return cmp.Or(sl...), nil
	case []float64:
		if len(s) != 1 {
			return nil, ErrOnlyOneSliceIsAllowed
		}
		return cmp.Or(sl...), nil
	case []string:
		if len(s) != 1 {
			return nil, ErrOnlyOneSliceIsAllowed
		}
		return cmp.Or(sl...), nil
	case []int:
		if len(s) != 1 {
			return nil, ErrOnlyOneSliceIsAllowed
		}
		return cmp.Or(sl...), nil
	case []int8:
		if len(s) != 1 {
			return nil, ErrOnlyOneSliceIsAllowed
		}
		return cmp.Or(sl...), nil
	case []int16:
		if len(s) != 1 {
			return nil, ErrOnlyOneSliceIsAllowed
		}
		return cmp.Or(sl...), nil
	case []int32:
		if len(s) != 1 {
			return nil, ErrOnlyOneSliceIsAllowed
		}
		return cmp.Or(sl...), nil
	case []int64:
		if len(s) != 1 {
			return nil, ErrOnlyOneSliceIsAllowed
		}
		return cmp.Or(sl...), nil
	case []uint8:
		if len(s) != 1 {
			return nil, ErrOnlyOneSliceIsAllowed
		}
		return cmp.Or(sl...), nil
	case []uint16:
		if len(s) != 1 {
			return nil, ErrOnlyOneSliceIsAllowed
		}
		return cmp.Or(sl...), nil
	case []uint32:
		if len(s) != 1 {
			return nil, ErrOnlyOneSliceIsAllowed
		}
		return cmp.Or(sl...), nil
	case []uint64:
		if len(s) != 1 {
			return nil, ErrOnlyOneSliceIsAllowed
		}
		return cmp.Or(sl...), nil
	default:
		return cmp.Or(s...), nil
	}
}
