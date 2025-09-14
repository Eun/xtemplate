package funcs

import "slices"

// Safe is the set of functions that are considered safe for use in untrusted templates.
var Safe = slices.Concat(
	Cmp,
	Conv,
	Dict,
	FilePath,
	JSON,
	Path,
	Regexp,
	Slice,
	Strings,
	Tmpl,
	URL,
)
