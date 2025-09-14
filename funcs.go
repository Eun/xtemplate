// Package xtemplate provides a way to create text/template with a restricted set of functions.
// It includes functions from various standard library packages such as path, filepath, strings, os, and encoding/json.
// Users can specify which functions are allowed to be used in the templates, enhancing security and control.
// The package also provides a special "return" function to short-circuit template execution and return a value.
// This is useful for scenarios where you want to exit a template early with a specific value
// without rendering the rest of the template.
package xtemplate

import (
	"fmt"
	"text/template"

	"github.com/Eun/xtemplate/funcs"
)

//go:generate go run generate_funcs.go
//go:generate go run generate_examples.go

type rootContext struct {
	template           *template.Template
	allowedFunctionSet map[funcs.Func]struct{}
}

// FuncNotAllowedError is returned when a function is called that is not in the allowed function set.
type FuncNotAllowedError struct {
	Func funcs.Func
}

func (e *FuncNotAllowedError) Error() string {
	return fmt.Sprintf("function %s.%s is not allowed", e.Func.Namespace, e.Func.Name)
}

// ReturnError is used to short-circuit template execution and return a value.
type ReturnError struct {
	Value any
}

func (r ReturnError) Error() string {
	return "return"
}

// AllowedFunctions is an interface that types can implement to provide a list of allowed functions.
type AllowedFunctions interface {
	Functions() []funcs.Func
}

// FuncMap returns a template.FuncMap containing only the functions specified in allowedFunctions.
//
//nolint:cyclop, funlen // cannot be simplified
func FuncMap(t *template.Template, allowedFunctions ...AllowedFunctions) template.FuncMap {
	allowedNamespaceSet, allowedFunctionSet := createAllowedFunctionSet(allowedFunctions)
	m := template.FuncMap{
		"return": func(value any) (any, error) {
			return nil, ReturnError{Value: value}
		},
	}

	rootCtx := rootContext{
		template:           t,
		allowedFunctionSet: allowedFunctionSet,
	}

	if _, ok := allowedNamespaceSet["conv"]; ok {
		m["conv"] = func(...any) (any, error) {
			return Conv(rootCtx), nil
		}
	}

	if _, ok := allowedNamespaceSet["cmp"]; ok {
		m["cmp"] = func(...any) (any, error) {
			return Cmp(rootCtx), nil
		}
	}

	if _, ok := allowedNamespaceSet["dict"]; ok {
		m["dict"] = func(...any) (any, error) {
			return Dict(rootCtx), nil
		}
	}

	if _, ok := allowedNamespaceSet["filepath"]; ok {
		m["filepath"] = func(...any) (any, error) {
			return FilePath(rootCtx), nil
		}
	}

	if _, ok := allowedNamespaceSet["json"]; ok {
		m["json"] = func(...any) (any, error) {
			return JSON(rootCtx), nil
		}
	}

	if _, ok := allowedNamespaceSet["os"]; ok {
		m["os"] = func(...any) (any, error) {
			return OS(rootCtx), nil
		}
	}

	if _, ok := allowedNamespaceSet["path"]; ok {
		m["path"] = func(...any) (any, error) {
			return Path(rootCtx), nil
		}
	}

	if _, ok := allowedNamespaceSet["regexp"]; ok {
		m["regexp"] = func(...any) (any, error) {
			return Regexp(rootCtx), nil
		}
	}

	if _, ok := allowedNamespaceSet["slice"]; ok {
		m["slice"] = func(...any) (any, error) {
			return Slice(rootCtx), nil
		}
	}

	if _, ok := allowedNamespaceSet["strings"]; ok {
		m["strings"] = func(...any) (any, error) {
			return Strings(rootCtx), nil
		}
	}

	if _, ok := allowedNamespaceSet["tmpl"]; ok {
		m["tmpl"] = func(...any) (any, error) {
			return Tmpl(rootCtx), nil
		}
	}

	if _, ok := allowedNamespaceSet["url"]; ok {
		m["url"] = func(...any) (any, error) {
			return URL(rootCtx), nil
		}
	}
	return m
}

func createAllowedFunctionSet(
	allowedFunctions []AllowedFunctions,
) (namespaceSet map[string]struct{}, functionSet map[funcs.Func]struct{}) {
	allFunctions := make([]funcs.Func, 0, len(allowedFunctions))
	for _, f := range allowedFunctions {
		allFunctions = append(allFunctions, f.Functions()...)
	}
	namespaceSet = make(map[string]struct{})
	functionSet = make(map[funcs.Func]struct{})
	for _, allowedFunction := range allFunctions {
		namespace := allowedFunction.Namespace
		functionName := allowedFunction.Name

		// lookup namespace
		nsSet, ok := funcs.NamespacesAndTheirFunctions[namespace]
		if !ok {
			continue
		}
		// lookup function in namespace
		_, ok = nsSet[functionName]
		if !ok {
			continue
		}

		namespaceSet[namespace] = struct{}{}

		f := funcs.Func{
			Namespace: allowedFunction.Namespace,
			Name:      allowedFunction.Name,
		}
		functionSet[f] = struct{}{}
	}
	return namespaceSet, functionSet
}
