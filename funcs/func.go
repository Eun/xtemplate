// Package funcs package holds function identifiers and collections that can be used
// with the xtemplate package to specify allowed functions in templates.
package funcs

// Func represents a function with its namespace and name.
type Func struct {
	Namespace string
	Name      string
}

// Functions returns a slice containing the Func itself.
func (f Func) Functions() []Func {
	return []Func{f}
}

// Funcs is a collection of Func.
type Funcs []Func

// Functions returns the list of Func contained in the Funcs collection.
func (f Funcs) Functions() []Func {
	return f
}
