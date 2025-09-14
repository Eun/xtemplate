package xtemplate

import (
	"regexp"

	"github.com/Eun/xtemplate/funcs"
)

// Regexp provides access to functions in the regexp package.
type Regexp rootContext

// MatchString reports whether the string s contains any match of the regular expression pattern.
//
// Example 1:
//
//	{{ regexp.MatchString "p([a-z]+)ch" "peach" }} // Output: true
//
// Example 2:
//
//	{{ regexp.MatchString "p([a-z]+)ch" "apple" }} // Output: false
func (ctx Regexp) MatchString(pattern string, s string) (bool, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.RegexpMatchString]; !ok {
		return false, &FuncNotAllowedError{Func: funcs.RegexpMatchString}
	}
	return regexp.MatchString(pattern, s)
}

// QuoteMeta returns a string that escapes all regular expression metacharacters
// inside the argument text; the returned string is a regular expression matching
// the literal text.
//
// Example:
//
//	{{ regexp.QuoteMeta "Escaping $5.00?" }} // Output: Escaping \$5\.00\?
func (ctx Regexp) QuoteMeta(s string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.RegexpQuoteMeta]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.RegexpQuoteMeta}
	}
	return regexp.QuoteMeta(s), nil
}

// FindAllString returns a slice of all successive matches of the expression,
// as defined by the 'All' description in the package comment.
// A return value of nil indicates no match.
//
// Example:
//
//	{{ regexp.FindAllString "p([a-z]+)ch" "peach punch pinch" -1 }} // Output: [peach punch pinch]
func (ctx Regexp) FindAllString(pattern string, s string, n int) ([]string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.RegexpFindAllString]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.RegexpFindAllString}
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return re.FindAllString(s, n), nil
}

// FindAllStringIndex returns a slice of all successive matches of the expression,
// as defined by the 'All' description in the package comment.
// A return value of nil indicates no match.
//
// Example:
//
//	{{ regexp.FindAllStringIndex "p([a-z]+)ch" "peach punch" -1 }} // Output: [[0 5] [6 11]]
func (ctx Regexp) FindAllStringIndex(pattern string, s string, n int) ([][]int, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.RegexpFindAllString]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.RegexpFindAllString}
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return re.FindAllStringIndex(s, n), nil
}

// FindAllStringSubmatch returns a slice of all successive matches of the expression,
// as defined by the 'All' description in the package comment.
// A return value of nil indicates no match.
//
// Example 1:
//
//	{{ regexp.FindAllStringSubmatch "a(x*)b" "-ab-" -1 }} // Output: [[ab ]]
//
// Example 2:
//
//	{{ regexp.FindAllStringSubmatch "a(x*)b" "-axxb-" -1 }} // [[axxb xx]]
//
// Example 3:
//
//	{{ regexp.FindAllStringSubmatch "a(x*)b" "-ab-axb-" -1 }} // [[ab ] [axb x]]
//
// Example 4:
//
//	{{ regexp.FindAllStringSubmatch "a(x*)b" "-axxb-ab-" -1 }} // Output: [[axxb xx] [ab ]]
func (ctx Regexp) FindAllStringSubmatch(pattern string, s string, n int) ([][]string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.RegexpFindAllStringSubmatch]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.RegexpFindAllStringSubmatch}
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return re.FindAllStringSubmatch(s, n), nil
}

// FindAllStringSubmatchIndex returns a slice of all successive matches of the expression,
// as defined by the 'All' description in the package comment.
// A return value of nil indicates no match.
//
// Example 1:
//
//	{{ regexp.FindAllStringSubmatchIndex "a(x*)b" "-ab-" -1 }} // Output: [[1 3 2 2]]
//
// Example 2:
//
//	{{ regexp.FindAllStringSubmatchIndex "a(x*)b" "-axxb-" -1 }} // Output: [[1 5 2 4]]
//
// Example 3:
//
//	{{ regexp.FindAllStringSubmatchIndex "a(x*)b" "-ab-axb-" -1 }} // Output: [[1 3 2 2] [4 7 5 6]]
//
// Example 4:
//
//	{{ regexp.FindAllStringSubmatchIndex "a(x*)b" "-axxb-ab-" -1 }} // Output: [[1 5 2 4] [6 8 7 7]]
//
// Example 5:
//
//	{{ regexp.FindAllStringSubmatchIndex "a(x*)b" "-foo-" -1 }} // Output: []
func (ctx Regexp) FindAllStringSubmatchIndex(pattern string, s string, n int) ([][]int, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.RegexpFindAllStringSubmatchIndex]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.RegexpFindAllStringSubmatchIndex}
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return re.FindAllStringSubmatchIndex(s, n), nil
}

// FindString returns a string holding the text of the leftmost match in s of the regular expression.
// If there is no match, the return value is an empty string.
//
// Example:
//
//	{{ regexp.FindString "p([a-z]+)ch" "peach punch pinch" }} // Output: peach
func (ctx Regexp) FindString(pattern string, s string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.RegexpFindString]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.RegexpFindString}
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}
	return re.FindString(s), nil
}

// FindStringIndex returns a two-element slice of integers defining the location of the leftmost match
// in s of the regular expression. The match itself is at s[loc[0]:loc[1]].
// A return value of nil indicates no match.
//
// Example:
//
//	{{ regexp.FindStringIndex "p([a-z]+)ch" "peach punch" }} // Output: [0 5]
func (ctx Regexp) FindStringIndex(pattern string, s string) ([]int, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.RegexpFindStringIndex]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.RegexpFindStringIndex}
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return re.FindStringIndex(s), nil
}

// FindStringSubmatch returns a slice of strings holding the text of the leftmost match of the regular expression
// in s and the matches, if any, of its subexpressions.
// A return value of nil indicates no match.
//
// Example1:
//
//	{{ regexp.FindStringSubmatch "a(x*)b(y|z)c" "-axxxbyc-" }} // Output: [axxxbyc xxx y]
//
// Example 2:
//
//	{{ regexp.FindStringSubmatch "a(x*)b(y|z)c" "-abzc-" }} // Output: [abzc  z]
func (ctx Regexp) FindStringSubmatch(pattern string, s string) ([]string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.RegexpFindStringSubmatch]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.RegexpFindStringSubmatch}
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return re.FindStringSubmatch(s), nil
}

// FindStringSubmatchIndex returns a slice of integers holding the text of the leftmost match of the regular expression
// in s and the matches, if any, of its subexpressions.
// A return value of nil indicates no match.
//
// Example:
//
//	{{ regexp.FindStringSubmatchIndex "p([a-z]+)ch" "peach" }} // Output: [0 5 1 3]
func (ctx Regexp) FindStringSubmatchIndex(pattern string, s string) ([]int, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.RegexpFindStringSubmatchIndex]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.RegexpFindStringSubmatchIndex}
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return re.FindStringSubmatchIndex(s), nil
}

// ReplaceAllLiteralString returns a copy of s, replacing matches of the Regexp
// with the replacement string repl. The replacement repl is substituted directly,
// without using Expand.
//
// Example:
//
//	{{ regexp.ReplaceAllLiteralString "a(x*)b" "-ab-axxb-" "T" }} // Output: -T-T-
func (ctx Regexp) ReplaceAllLiteralString(pattern string, s string, repl string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.RegexpReplaceAllLiteralString]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.RegexpReplaceAllLiteralString}
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}
	return re.ReplaceAllLiteralString(s, repl), nil
}

// ReplaceAllString returns a copy of s, replacing matches of the Regexp
// with the replacement string repl. Inside repl, $ signs are interpreted as in Expand,
// so for instance $1 represents the text of the first submatch.
//
// Example:
//
//	{{ regexp.ReplaceAllString "a(x*)b" "-ab-axxb-" "${1}W" }} // Output: -W-xxW-
func (ctx Regexp) ReplaceAllString(pattern string, s string, repl string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.RegexpReplaceAllString]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.RegexpReplaceAllString}
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}
	return re.ReplaceAllString(s, repl), nil
}

// Split slices s into substrings separated by the expression and returns a slice of
// the substrings between those expression matches.
//
// Example 1:
//
//	{{ regexp.Split "a" "banana" -1 }} // Output: [b n n ]
//
// Example 2:
//
//	{{ regexp.Split "a" "apple" 0 }} // Output: []
//
// Example 3:
//
//	{{ regexp.Split "a" "grape" 1 }} // Output: [grape]
//
// Example 4:
//
//	{{ regexp.Split "z+" "pizza" 2 }} // Output: [pi a]
//
//nolint:dupword // false positive in the example
func (ctx Regexp) Split(pattern string, s string, n int) ([]string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.RegexpSplit]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.RegexpSplit}
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return re.Split(s, n), nil
}
