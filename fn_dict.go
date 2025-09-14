package xtemplate

import (
	"github.com/Eun/xtemplate/funcs"
)

// Dict provides helper functions for dictionaries.
type Dict rootContext

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
func (ctx Dict) HasKey(m map[any]any, key any) (bool, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.DictHasKey]; !ok {
		return false, &FuncNotAllowedError{Func: funcs.DictHasKey}
	}
	_, exists := m[key]
	return exists, nil
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
func (ctx Dict) HasValue(m map[any]any, value any) (bool, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.DictHasValue]; !ok {
		return false, &FuncNotAllowedError{Func: funcs.DictHasValue}
	}
	for _, v := range m {
		if v == value {
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
func (ctx Dict) Keys(m map[any]any) ([]any, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.DictKeys]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.DictKeys}
	}
	keys := make([]any, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys, nil
}
