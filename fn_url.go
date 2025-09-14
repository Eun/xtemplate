package xtemplate

import (
	"net/url"

	"github.com/Eun/xtemplate/funcs"
)

// URL provides access to functions in the url package.
type URL rootContext

// JoinPath returns a URL string with the provided path elements joined to the existing path of base and
// the resulting path cleaned of any ./ or ../ elements.
// Any sequences of multiple slashes will be reduced to a single slash.
//
// Example:
//
//	{{ url.JoinPath "https://example.com/foo" "bar" "baz" }}
func (ctx URL) JoinPath(base string, elem ...string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.URLJoinPath]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.URLJoinPath}
	}
	return url.JoinPath(base, elem...)
}

// PathEscape escapes the string so it can be safely placed inside a URL path segment,
// replacing special characters (including /) with %XX sequences as needed.
//
// Example:
//
//	{{ url.PathEscape "hello world" }}
func (ctx URL) PathEscape(s string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.URLPathEscape]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.URLPathEscape}
	}
	return url.PathEscape(s), nil
}

// PathUnescape does the inverse transformation of PathEscape,
// converting each 3-byte encoded substring of the form "%AB" into the hex-decoded byte 0xAB.
//
// Example:
//
//	{{ url.PathUnescape "hello%20world" }}
func (ctx URL) PathUnescape(s string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.URLPathUnescape]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.URLPathUnescape}
	}
	return url.PathUnescape(s)
}

// QueryEscape escapes the string so it can be safely placed inside a URL query.
// It is identical to PathEscape except that it also escapes '?'.
//
// Example:
//
//	{{ url.QueryEscape "hello world?" }}
func (ctx URL) QueryEscape(s string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.URLQueryEscape]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.URLQueryEscape}
	}
	return url.QueryEscape(s), nil
}

// QueryUnescape does the inverse transformation of QueryEscape,
// converting each 3-byte encoded substring of the form "%AB" into the hex-decoded byte 0xAB.
//
// Example:
//
//	{{ url.QueryUnescape "hello%20world%3F" }}
func (ctx URL) QueryUnescape(s string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.URLQueryUnescape]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.URLQueryUnescape}
	}
	return url.QueryUnescape(s)
}
