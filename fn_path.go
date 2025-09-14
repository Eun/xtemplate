package xtemplate

import (
	"path"

	"github.com/Eun/xtemplate/funcs"
)

// Path provides access to functions in the path package.
type Path rootContext

// Dir returns all but the last element of path, typically the path's directory.
// After dropping the final element, Dir calls Clean on the path and trailing slashes are removed.
// If the path is empty, Dir returns ".".
//
// Example:
//
//	{{ path.Dir "/foo/bar/baz" }} // Output: /foo/bar
func (ctx Path) Dir(s string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.PathDir]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.PathDir}
	}
	return path.Dir(s), nil
}

// Base returns the last element of path. Trailing slashes are removed before extracting the last element.
// If the path is empty, Base returns ".". If the path consists entirely of slashes, Base returns "/".
//
// Example:
//
//	{{ path.Base "/foo/bar/baz" }} // Output: baz
func (ctx Path) Base(s string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.PathBase]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.PathBase}
	}
	return path.Base(s), nil
}

// Join joins any number of path elements into a single path, separating them with slashes.
// Empty elements are ignored. The result is Cleaned.
//
// Example:
//
//	{{ path.Join "foo" "bar" "baz" }} // Output: foo/bar/baz
func (ctx Path) Join(s ...string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.PathJoin]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.PathJoin}
	}
	return path.Join(s...), nil
}

// Clean returns the shortest path name equivalent to path by purely lexical processing.
// It applies the following rules iteratively until no further processing can be done.
//
// Example:
//
//	{{ path.Clean "/foo//bar/../baz" }} // Output: /foo/baz
func (ctx Path) Clean(s string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.PathClean]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.PathClean}
	}
	return path.Clean(s), nil
}

// Ext returns the file name extension used by path. The extension is the suffix beginning at the final dot
// in the final slash-separated element of path; it is empty if there is no dot.
//
// Example:
//
//	{{ path.Ext "/foo/bar/baz.js" }} // Output: .js
func (ctx Path) Ext(s string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.PathExt]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.PathExt}
	}
	return path.Ext(s), nil
}
