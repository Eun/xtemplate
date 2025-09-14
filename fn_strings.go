package xtemplate

import (
	"strings"

	"github.com/Eun/xtemplate/funcs"
)

// Strings provides access to functions in the strings package.
type Strings rootContext

// Compare compares two strings lexicographically and returns an integer comparing two strings.
// The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
//
// Example 1:
//
//	{{ strings.Compare "apple" "banana" }} // Output: -1
//
// Example 2:
//
//	{{ strings.Compare "banana" "apple" }} // Output: 1
//
// Example 3:
//
//	{{ strings.Compare "apple" "apple" }}  // Output: 0
func (ctx Strings) Compare(a, b string) (int, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsCompare]; !ok {
		return 0, &FuncNotAllowedError{Func: funcs.StringsCompare}
	}
	return strings.Compare(a, b), nil
}

// Contains reports whether substr is within s.
//
// Example 1:
//
//	{{ strings.Contains "hello world" "world" }} // Output: true
//
// Example 2:
//
//	{{ strings.Contains "hello world" "mars" }}  // Output: false
func (ctx Strings) Contains(s, substr string) (bool, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsContains]; !ok {
		return false, &FuncNotAllowedError{Func: funcs.StringsContains}
	}
	return strings.Contains(s, substr), nil
}

// ContainsAny reports whether any Unicode code points in chars are within s.
//
// Example 1:
//
//	{{ strings.ContainsAny "hello" "aeiou" }} // Output: true
//
// Example 2:
//
//	{{ strings.ContainsAny "rhythm" "aeiou" }} // Output: false
func (ctx Strings) ContainsAny(s, chars string) (bool, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsContainsAny]; !ok {
		return false, &FuncNotAllowedError{Func: funcs.StringsContainsAny}
	}
	return strings.ContainsAny(s, chars), nil
}

// ContainsRune reports whether the Unicode code point r is within s.
//
// Example 1:
//
//	{{ strings.ContainsRune "hello" 'e' }} // Output: true
//
// Example 2:
//
//	{{ strings.ContainsRune "hello" 'a' }} // Output: false
func (ctx Strings) ContainsRune(s string, r rune) (bool, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsContainsRune]; !ok {
		return false, &FuncNotAllowedError{Func: funcs.StringsContainsRune}
	}
	return strings.ContainsRune(s, r), nil
}

// Count counts the number of non-overlapping instances of substr in s.
// If substr is an empty string, Count returns 1 + the number of Unicode code points in s.
//
// Example 1:
//
//	{{ strings.Count "hello hello" "hello" }} // Output: 2
//
// Example 2:
//
//	{{ strings.Count "hello" "l" }}           // Output: 2
//
// Example 3:
//
//	{{ strings.Count "hello" "" }}            // Output: 6
func (ctx Strings) Count(s, substr string) (int, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsCount]; !ok {
		return 0, &FuncNotAllowedError{Func: funcs.StringsCount}
	}
	return strings.Count(s, substr), nil
}

// CutResult is the result type for the Cut function.
type CutResult struct {
	Before string
	After  string
	Found  bool
}

// Cut slices s around the first instance of sep, returning the text before and after sep.
// The found result reports whether sep appears in s.
//
// Example:
//
//	{{ strings.Cut "apple,banana" "," }} // Output: {apple banana true}
func (ctx Strings) Cut(s, sep string) (CutResult, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsCut]; !ok {
		return CutResult{}, &FuncNotAllowedError{Func: funcs.StringsCut}
	}
	before, after, found := strings.Cut(s, sep)
	return CutResult{
		Before: before,
		After:  after,
		Found:  found,
	}, nil
}

// CutPrefixResult is the result type for the CutPrefix function.
type CutPrefixResult struct {
	After string
	Found bool
}

// CutPrefix returns s without the provided leading prefix string and reports whether it found the prefix.
// If s doesn't start with prefix, CutPrefix returns s, false.
//
// Example:
//
//	{{ strings.CutPrefix "Hello, World!" "Hello, " }} // Output: {World! true}
func (ctx Strings) CutPrefix(s, sep string) (CutPrefixResult, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsCutPrefix]; !ok {
		return CutPrefixResult{}, &FuncNotAllowedError{Func: funcs.StringsCutPrefix}
	}
	after, found := strings.CutPrefix(s, sep)
	return CutPrefixResult{
		After: after,
		Found: found,
	}, nil
}

// CutSuffixResult is the result type for the CutSuffix function.
type CutSuffixResult struct {
	Before string
	Found  bool
}

// CutSuffix returns s without the provided ending suffix string and reports whether it found the suffix.
// If s doesn't end with suffix, CutSuffix returns s, false.
//
// Example:
//
//	{{ strings.CutSuffix "Hello, World!" ", World!" }} // Output: {Hello true}
func (ctx Strings) CutSuffix(s, sep string) (CutSuffixResult, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsCutSuffix]; !ok {
		return CutSuffixResult{}, &FuncNotAllowedError{Func: funcs.StringsCutSuffix}
	}
	before, found := strings.CutSuffix(s, sep)
	return CutSuffixResult{
		Before: before,
		Found:  found,
	}, nil
}

// Equal reports whether s and t are the same string (case-sensitive).
//
// Example:
//
//	{{ strings.Equal "hello" "hello" }} // Output: true
func (ctx Strings) Equal(s, t string) (bool, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsEqual]; !ok {
		return false, &FuncNotAllowedError{Func: funcs.StringsEqual}
	}
	return s == t, nil
}

// EqualFold reports whether s and t are equal under Unicode case-folding, which is a more general
// form of case-insensitivity.
//
// Example:
//
//	{{ strings.EqualFold "Go" "go" }} // Output: true
func (ctx Strings) EqualFold(s, t string) (bool, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsEqualFold]; !ok {
		return false, &FuncNotAllowedError{Func: funcs.StringsEqualFold}
	}
	return strings.EqualFold(s, t), nil
}

// Fields splits the string s around each instance of one or more consecutive white space
// characters, returning a slice of substrings of s or an empty slice if s contains only white space.
//
// Example:
//
//	{{ strings.Fields "  hello   world  " }} // Output: [hello world]
func (ctx Strings) Fields(s string) ([]string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsFields]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.StringsFields}
	}
	return strings.Fields(s), nil
}

// HasPrefix tests whether the string s begins with prefix.
//
// Example:
//
//	{{ strings.HasPrefix "Hello, World!" "Hello" }} // Output: true
func (ctx Strings) HasPrefix(s, prefix string) (bool, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsHasPrefix]; !ok {
		return false, &FuncNotAllowedError{Func: funcs.StringsHasPrefix}
	}
	return strings.HasPrefix(s, prefix), nil
}

// HasSuffix tests whether the string s ends with suffix.
//
// Example:
//
//	{{ strings.HasSuffix "Hello, World!" "World!" }} // Output: true
func (ctx Strings) HasSuffix(s, suffix string) (bool, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsHasSuffix]; !ok {
		return false, &FuncNotAllowedError{Func: funcs.StringsHasSuffix}
	}
	return strings.HasSuffix(s, suffix), nil
}

// Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
//
// Example:
//
//	{{ strings.Index "hello world" "world" }} // Output: 6
func (ctx Strings) Index(s, substr string) (int, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsIndex]; !ok {
		return 0, &FuncNotAllowedError{Func: funcs.StringsIndex}
	}
	return strings.Index(s, substr), nil
}

// IndexAny returns the index of the first instance of any Unicode code point from chars in s,
// or -1 if no Unicode code point from chars is present in s.
//
// Example:
//
//	{{ strings.IndexAny "hello" "aeiou" }} // Output: 1
func (ctx Strings) IndexAny(s1, chars string) (int, error) {
	return strings.IndexAny(s1, chars), nil
}

// IndexByte returns the index of the first instance of the given byte in s, or -1 if c is not present in s.
//
// Example:
//
//	{{ strings.IndexByte "hello" 'l' }} // Output: 2
func (ctx Strings) IndexByte(s string, c byte) (int, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsIndexByte]; !ok {
		return 0, &FuncNotAllowedError{Func: funcs.StringsIndexByte}
	}
	return strings.IndexByte(s, c), nil
}

// IndexRune returns the index of the first instance of the Unicode code point r, or -1 if rune is not present in s.
//
// Example:
//
//	{{ strings.IndexRune "hello" 'e' }} // Output: 1
func (ctx Strings) IndexRune(s string, c rune) (int, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsIndexRune]; !ok {
		return 0, &FuncNotAllowedError{Func: funcs.StringsIndexRune}
	}
	return strings.IndexRune(s, c), nil
}

// Join concatenates the elements of a to create a single string. The separator string
// sep is placed between elements in the resulting string.
//
// Example:
//
//	{{ strings.Join ( slice.NewStrings "hello" "world" ) " " }} // Output: hello world
func (ctx Strings) Join(a []string, sep string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsJoin]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.StringsJoin}
	}
	return strings.Join(a, sep), nil
}

// LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
//
// Example:
//
//	{{ strings.LastIndex "hello hello" "hello" }} // Output: 6
func (ctx Strings) LastIndex(s, substr string) (int, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsLastIndex]; !ok {
		return 0, &FuncNotAllowedError{Func: funcs.StringsLastIndex}
	}
	return strings.LastIndex(s, substr), nil
}

// LastIndexAny returns the index of the last instance of any Unicode code point from chars in s,
// or -1 if no Unicode code point from chars is present in s.
//
// Example:
//
//	{{ strings.LastIndexAny "hello" "aeiou" }} // Output: 4
func (ctx Strings) LastIndexAny(s, substr string) (int, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsLastIndexAny]; !ok {
		return 0, &FuncNotAllowedError{Func: funcs.StringsLastIndexAny}
	}
	return strings.LastIndexAny(s, substr), nil
}

// LastIndexByte returns the index of the last instance of the given byte in s, or -1 if c is not present in s.
//
// Example:
//
//	{{ strings.LastIndexByte "hello" 'l' }} // Output: 3
func (ctx Strings) LastIndexByte(s string, c byte) (int, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsLastIndexByte]; !ok {
		return 0, &FuncNotAllowedError{Func: funcs.StringsLastIndexByte}
	}
	return strings.LastIndexByte(s, c), nil
}

// Repeat returns a new string consisting of count copies of the string s.
// It panics if count is negative or if the result of (len(s) * count) overflows.
//
// Example:
//
//	{{ strings.Repeat "ha" 3 }} // Output: hahaha
func (ctx Strings) Repeat(s string, count int) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsRepeat]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.StringsRepeat}
	}
	return strings.Repeat(s, count), nil
}

// Replace returns a copy of the string s with the first n non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string and after each UTF-8 sequence,
// yielding up to k+1 replacements for a k-rune string. If n < 0, there is no limit on the number of replacements.
//
// Example:
//
//	{{ strings.Replace "hello world hello" "hello" "hi" 1 }} // Output: hi world hello
func (ctx Strings) Replace(s, old, replacement string, n int) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsReplace]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.StringsReplace}
	}
	return strings.Replace(s, old, replacement, n), nil
}

// ReplaceAll returns a copy of the string s with all non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string and after each UTF-8 sequence,
// yielding up to k+1 replacements for a k-rune string.
//
// Example:
//
//	{{ strings.ReplaceAll "hello world hello" "hello" "hi" }} // Output: hi world hi
func (ctx Strings) ReplaceAll(s, old, replacement string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsReplaceAll]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.StringsReplaceAll}
	}
	return strings.ReplaceAll(s, old, replacement), nil
}

// Split slices s into all substrings separated by sep and returns a slice of the substrings between those separators.
// If s does not contain sep and sep is not empty, Split returns a slice of length 1 whose only element is s.
//
// Example:
//
//	{{ strings.Split "apple,banana,cherry" "," }} // Output: [apple banana cherry]
func (ctx Strings) Split(s, sep string) ([]string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsSplit]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.StringsSplit}
	}
	return strings.Split(s, sep), nil
}

// SplitAfter slices s into all substrings after each instance of sep and returns a slice of those substrings.
// If s does not contain sep and sep is not empty, SplitAfter returns a slice of length 1 whose only element is s.
//
// Example:
//
//	{{ strings.SplitAfter "apple,banana,cherry" "," }} // Output: [apple, banana, cherry]
func (ctx Strings) SplitAfter(s, sep string) ([]string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsSplitAfter]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.StringsSplitAfter}
	}
	return strings.SplitAfter(s, sep), nil
}

// SplitAfterN slices s into substrings after each instance of sep and returns a slice of those substrings.
// The count determines the number of substrings to return:
// n > 0: at most n substrings;
// n == 0: the result is nil (zero substrings);
// n < 0: all substrings.
//
// Example:
//
//	{{ strings.SplitAfterN "apple,banana,cherry" "," 2 }} // Output: [apple, banana,cherry]
func (ctx Strings) SplitAfterN(s, sep string, n int) ([]string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsSplitAfterN]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.StringsSplitAfterN}
	}
	return strings.SplitAfterN(s, sep, n), nil
}

// SplitN slices s into substrings separated by sep and returns a slice of the substrings between those separators.
// The count determines the number of substrings to return
// n > 0: at most n substrings;
// n == 0: the result is nil (zero substrings);
// n < 0: all substrings.
//
// Example:
//
//	{{ strings.SplitN "apple,banana,cherry" "," 2 }} // Output: [apple banana,cherry]
func (ctx Strings) SplitN(s, sep string, n int) ([]string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsSplitN]; !ok {
		return nil, &FuncNotAllowedError{Func: funcs.StringsSplitN}
	}
	return strings.SplitN(s, sep, n), nil
}

// ToLower is a wrapper around strings.ToLower that lowercases the input string.
//
// Example:
//
//	{{ strings.ToLower "TEST" }} // Output: test
func (ctx Strings) ToLower(s string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsToLower]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.StringsToLower}
	}
	return strings.ToLower(s), nil
}

// ToTitle returns a copy of the string s with all Unicode letters mapped to their Unicode title case.
//
// Example:
//
//	{{ strings.ToTitle "hello world" }} // Output: HELLO WORLD
func (ctx Strings) ToTitle(s string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsToTitle]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.StringsToTitle}
	}
	return strings.ToTitle(s), nil
}

// ToUpper is a wrapper around strings.ToUpper that uppercases the input string.
//
// Example:
//
//	{{ strings.ToUpper "test" }} // Output: TEST
func (ctx Strings) ToUpper(s string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsToUpper]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.StringsToUpper}
	}
	return strings.ToUpper(s), nil
}

// ToValidUTF8 returns a copy of the string s with each run of invalid UTF-8 byte sequences
// replaced by the replacement string, which may be empty.
//
// Example:
//
//	{{ strings.ToValidUTF8 "Hello\xc5World" "?" }}
func (ctx Strings) ToValidUTF8(s, replacement string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsToValidUTF8]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.StringsToValidUTF8}
	}
	return strings.ToValidUTF8(s, replacement), nil
}

// Trim returns a slice of the string s with all leading and trailing Unicode code points
// contained in cutset removed.
//
// Example:
//
//	{{ strings.Trim "¡¡¡Hello, Gophers!!!" "!¡" }} // Output: Hello, Gophers
func (ctx Strings) Trim(s, cutset string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsTrim]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.StringsTrim}
	}
	return strings.Trim(s, cutset), nil
}

// TrimLeft returns a slice of the string s with all leading Unicode code points
// contained in cutset removed.
//
// Example:
//
//	{{ strings.TrimLeft "¡¡¡Hello, Gophers!!!" "!¡" }} // Output: Hello, Gophers!!!
func (ctx Strings) TrimLeft(s, cutset string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsTrimLeft]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.StringsTrimLeft}
	}
	return strings.TrimLeft(s, cutset), nil
}

// TrimPrefix returns s without the provided leading prefix string.
// If s doesn't start with prefix, s is returned unchanged.
//
// Example:
//
//	{{ strings.TrimPrefix "Hello, World!" "Hello, " }} // Output: World!
func (ctx Strings) TrimPrefix(s, prefix string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsTrimPrefix]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.StringsTrimPrefix}
	}
	return strings.TrimPrefix(s, prefix), nil
}

// TrimRight returns a slice of the string s, with all trailing Unicode code points
// contained in cutset removed.
//
// Example:
//
//	{{ strings.TrimRight "¡¡¡Hello, Gophers!!!" "!¡" }} // Output: ¡¡¡Hello, Gophers
func (ctx Strings) TrimRight(s, cutset string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsTrimRight]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.StringsTrimRight}
	}
	return strings.TrimRight(s, cutset), nil
}

// TrimSpace returns a slice of the string s, with all leading and trailing white space
// removed, as defined by Unicode.
//
// Example:
//
//	{{ strings.TrimSpace "  \t\n Hello, Gophers \n\t\r\n" }} // Output: Hello, Gophers
func (ctx Strings) TrimSpace(s string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsTrimSpace]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.StringsTrimSpace}
	}
	return strings.TrimSpace(s), nil
}

// TrimSuffix returns s without the provided trailing suffix string.
// If s doesn't end with suffix, s is returned unchanged.
//
// Example:
//
//	{{ strings.TrimSuffix "Hello, World!" ", World!" }} // Output: Hello
func (ctx Strings) TrimSuffix(s, prefix string) (string, error) {
	if _, ok := ctx.allowedFunctionSet[funcs.StringsTrimSuffix]; !ok {
		return "", &FuncNotAllowedError{Func: funcs.StringsTrimSuffix}
	}
	return strings.TrimSuffix(s, prefix), nil
}
