package xtemplate

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/Eun/xtemplate/funcs"
)

// JSON provides access to functions in the encoding/json package.
type JSON rootContext

// Compact appends to dst the JSON-encoded src with insignificant space characters elided.
//
// Example:
//
//	{{ json.Compact .Buffer .JSONBytes }}
func (ctx JSON) Compact(dst *bytes.Buffer, src []byte) error {
	if _, ok := ctx.allowedFunctionSet[funcs.JSONCompact]; !ok {
		return &FuncNotAllowedError{Func: funcs.JSONCompact}
	}
	return json.Compact(dst, src)
}

// HTMLEscape appends to dst the JSON-encoded src with <, >, &, U+2028 and U+2029
// characters inside string literals changed to \u003c, \u003e, \u0026, \u2028, \u2029
// so that the JSON will be safe to embed inside HTML <script> tags.
// For historical reasons, web browsers don't honor the standard HTML
// escaping rules within <script> tags, but they do honor the JSON backslash
// escaping, and the JSON specification allows backslash-escaping of
// these characters, so this function enables JSON to be safely placed
// inside HTML <script> tags.
// HTMLEscape only affects the contents of string literals in the JSON.
// It has no effect on the structural characters of the JSON itself.
//
// Example:
//
//	{{ json.HTMLEscape .Buffer .JSONBytes }}
func (ctx JSON) HTMLEscape(dst *bytes.Buffer, src []byte) error {
	if _, ok := ctx.allowedFunctionSet[funcs.JSONHTMLEscape]; !ok {
		return &FuncNotAllowedError{Func: funcs.JSONHTMLEscape}
	}
	json.HTMLEscape(dst, src)
	return nil
}

// Indent appends to dst an indented form of the JSON-encoded src.
// Each element in a JSON object or array begins on a new line, indented according to the indentation nesting.
// The data appended to dst does not begin with the prefix nor any indentation,
// to make it easier to embed inside other formatted JSON data.
// Although leading space characters (space, tab, carriage return, newline)
// at the beginning of src are dropped, trailing space characters
// at the end of src are preserved and copied to dst.
// For example, if src has no trailing spaces, neither will dst;
// if src ends in a trailing newline, so will dst.
//
// Example:
//
//	{{ json.Indent .Buffer .JSONBytes "" "  " }}
func (ctx JSON) Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error {
	if _, ok := ctx.allowedFunctionSet[funcs.JSONIndent]; !ok {
		return &FuncNotAllowedError{Func: funcs.JSONIndent}
	}
	return json.Indent(dst, src, prefix, indent)
}

// Marshal returns the JSON encoding of v.
//
// Example:
//
//	{{ $dict := dict.New "foo" "bar" }}
//	{{ $buf := json.Marshal $dict }}
//	{{ conv.ToString $buf }} // Output: {"foo":"bar"}
func (ctx JSON) Marshal(v any) ([]byte, error) {
	if m, ok := v.(map[any]any); ok {
		// json.Marshal doesn't support map[any]any, so convert to map[string]any
		m2 := make(map[string]any, len(m))
		for k, v := range m {
			s, ok := k.(string)
			if !ok {
				//nolint:err113 // allow dynamic error
				return nil, fmt.Errorf("json.Marshal: map key %v is type %T is not a string", k, k)
			}
			m2[s] = v
		}
		v = m2
	}
	if _, ok := ctx.allowedFunctionSet[funcs.JSONMarshal]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.JSONMarshal}
	}
	return json.Marshal(v)
}

// MarshalIndent is like Marshal but applies Indent to format the output.
// Each JSON element in the output will begin on a new line beginning with prefix
// followed by one or more copies of indent according to the indentation nesting.
//
// Example:
//
//	{{ json.MarshalIndent .Data "" "  " }}
func (ctx JSON) MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.JSONMarshalIndent]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.JSONMarshalIndent}
	}
	return json.MarshalIndent(v, prefix, indent)
}

// Unmarshal parses the JSON-encoded data and stores the result
// in the value pointed to by v. If v is nil or not a pointer,
// Unmarshal returns an InvalidUnmarshalError.
//
// Example:
//
//	{{ json.Unmarshal .JSONBytes .Target }}
func (ctx JSON) Unmarshal(data []byte, v any) error {
	if _, ok := ctx.allowedFunctionSet[funcs.JSONUnmarshal]; !ok {
		return &FuncNotAllowedError{Func: funcs.JSONUnmarshal}
	}
	return json.Unmarshal(data, v)
}

// Valid reports whether data is a valid JSON encoding.
//
// Example:
//
//	{{ json.Valid .JSONBytes }}
func (ctx JSON) Valid(data []byte) (bool, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.JSONValid]; !ok {
		return false, &FuncNotAllowedError{Func: funcs.JSONValid}
	}
	return json.Valid(data), nil
}
