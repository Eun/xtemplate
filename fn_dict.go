package xtemplate

import (
	"errors"
	"reflect"

	"github.com/Eun/xtemplate/funcs"
)

// Dict provides helper functions for dictionaries.
type Dict rootContext

// ErrIsNotMap is returned when a provided value is not a map.
var ErrIsNotMap = errors.New("the provided value is not a map")

// New creates a map from a list of key/value pairs.
//
// Example:
//
//	{{ dict.New "name" "Frank" "age" 42 }} // Output: map[age:42 name:Frank]
func (ctx Dict) New(vals ...any) (map[any]any, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.DictNew]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.DictNew}
	}

	result := make(map[any]any)
	if len(vals)%2 != 0 {
		vals = append(vals, nil)
	}
	for i := 0; i < len(vals); i += 2 {
		key := vals[i]
		value := vals[i+1]
		result[key] = value
	}
	return result, nil
}

// HasKey checks if a map contains a given key.
//
// Example 1:
//
//	{{ dict.HasKey (dict.New "name" "Frank" "age" 42) "name" }} // Output: true
//
// Example 2:
//
//	{{ dict.HasKey (dict.New "name" "Frank" "age" 42) "email" }} // Output: false
func (ctx Dict) HasKey(m any, key any) (bool, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.DictHasKey]; !ok {
		return false, &FuncNotAllowedError{Func: funcs.DictHasKey}
	}
	v := reflect.ValueOf(m)
	if !v.IsValid() {
		return false, ErrIsNotMap
	}
	if v.Kind() != reflect.Map {
		return false, ErrIsNotMap
	}
	k := reflect.ValueOf(key)
	if !k.IsValid() {
		return false, nil
	}

	if k.Type() != v.Type().Key() {
		// convert the key to the correct type if possible
		if k.Type().ConvertibleTo(v.Type().Key()) {
			k = k.Convert(v.Type().Key())
		} else {
			return false, nil
		}
	}
	val := v.MapIndex(k)
	return val.IsValid(), nil
}

// HasValue checks if a map contains a given value.
//
// Example 1:
//
//	{{ dict.HasValue (dict.New "name" "Frank" "age" 42) 42 }} // Output: true
//
// Example 2:
//
//	{{ dict.HasValue (dict.New "name" "Frank" "age" 42) "Joe" }} // Output: false
func (ctx Dict) HasValue(m any, value any) (bool, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.DictHasValue]; !ok {
		return false, &FuncNotAllowedError{Func: funcs.DictHasValue}
	}

	v := reflect.ValueOf(m)
	if !v.IsValid() {
		return false, ErrIsNotMap
	}
	if v.Kind() != reflect.Map {
		return false, ErrIsNotMap
	}
	for _, key := range v.MapKeys() {
		val := v.MapIndex(key)
		if val.Interface() == value {
			return true, nil
		}
	}
	return false, nil
}

// Keys returns the keys of a map as a slice.
//
// Example:
//
//	{{ $dict := dict.New "name" "Frank" "age" 42 }}
//	{{ $keys := conv.ToStrings ( dict.Keys $dict ) }}
//	{{ slice.Sort $keys }} // Output: [age name]
func (ctx Dict) Keys(m any) ([]any, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.DictKeys]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.DictKeys}
	}
	v := reflect.ValueOf(m)
	if !v.IsValid() {
		return nil, ErrIsNotMap
	}
	if v.Kind() != reflect.Map {
		return nil, ErrIsNotMap
	}
	keys := make([]any, v.Len())
	for i, key := range v.MapKeys() {
		keys[i] = key.Interface()
	}
	return keys, nil
}

// IsEmpty checks if a map is empty.
//
// Example:
//
//	{{ dict.IsEmpty (dict.New) }} // Output: true
func (ctx Dict) IsEmpty(m any) (bool, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.DictIsEmpty]; !ok {
		return false, &FuncNotAllowedError{Func: funcs.DictIsEmpty}
	}
	v := reflect.ValueOf(m)
	if !v.IsValid() {
		return false, ErrIsNotMap
	}
	if v.Kind() != reflect.Map {
		return false, ErrIsNotMap
	}
	return v.Len() == 0, nil
}

// Len returns the length of a map.
//
// Example:
//
//	{{ dict.Len (dict.New) }} // Output: 0
func (ctx Dict) Len(m any) (int, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.DictLen]; !ok {
		return 0, &FuncNotAllowedError{Func: funcs.DictLen}
	}
	v := reflect.ValueOf(m)
	if !v.IsValid() {
		return 0, ErrIsNotMap
	}
	if v.Kind() != reflect.Map {
		return 0, ErrIsNotMap
	}
	return v.Len(), nil
}
