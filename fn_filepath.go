package xtemplate

import (
	"path/filepath"

	"github.com/Eun/xtemplate/funcs"
)

// FilePath provides access to functions in the path/filepath package.
type FilePath rootContext

// Dir returns all but the last element of path, typically the path's directory.
// After dropping the final element, Dir calls Clean on the path and trailing slashes are removed.
//
// Example:
//
//	{{ filepath.Dir "/foo/bar/baz.js" }} // Output: /foo/bar
func (ctx FilePath) Dir(s string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.FilePathDir]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.FilePathDir}
	}
	return filepath.Dir(s), nil
}

// Base returns the last element of path. Trailing path separators are removed before extracting the last element.
// If the path is empty, Base returns ".". If the path consists entirely of separators, Base returns a single separator.
//
// Example:
//
//	{{ filepath.Base "/foo/bar/baz.js" }} // Output: baz.js
func (ctx FilePath) Base(s string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.FilePathBase]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.FilePathBase}
	}
	return filepath.Base(s), nil
}

// Join joins any number of path elements into a single path, separating them with an OS specific Separator.
// Empty elements are ignored. The result is Cleaned.
//
// Example:
//
//	{{ filepath.Join "foo" "bar" "baz" }} // Output: foo/bar/baz
func (ctx FilePath) Join(s ...string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.FilePathJoin]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.FilePathJoin}
	}
	return filepath.Join(s...), nil
}

// Clean returns the shortest path name equivalent to path by purely lexical processing.
// It applies the following rules iteratively until no further processing can be done.
//
// Example:
//
//	{{ filepath.Clean "/foo//bar/../baz" }} // Output: /foo/baz
func (ctx FilePath) Clean(s string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.FilePathClean]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.FilePathClean}
	}
	return filepath.Clean(s), nil
}

// Ext returns the file name extension used by path. The extension is the suffix beginning at the final dot
// in the final element of path; it is empty if there is no dot.
//
// Example:
//
//	{{ filepath.Ext "/foo/bar/baz.js" }} // Output: .js
func (ctx FilePath) Ext(s string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.FilePathExt]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.FilePathExt}
	}
	return filepath.Ext(s), nil
}

// Abs returns an absolute representation of path. If the path is not absolute it will be joined with the current
// working directory to turn it into an absolute path.
//
// Example:
//
//	{{ filepath.Abs "foo/bar" }}
func (ctx FilePath) Abs(s string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.FilePathAbs]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.FilePathAbs}
	}
	return filepath.Abs(s)
}

// Rel returns a relative path that is lexically equivalent to targetpath when joined to basepath with an intervening
// separator. That is, Join(basepath, Rel(basepath, targetpath)) is equivalent to targetpath itself.
//
// Example:
//
//	{{ filepath.Rel "/a" "/a/b/c" }} // Output: b/c
func (ctx FilePath) Rel(basepath, targetpath string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.FilePathRel]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.FilePathRel}
	}
	return filepath.Rel(basepath, targetpath)
}
