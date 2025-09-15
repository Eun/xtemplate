package xtemplate

import (
	"bytes"
	"errors"
	"fmt"
)

// OnlyOneArgumentIsAllowedError indicates that only one argument is allowed.
type OnlyOneArgumentIsAllowedError struct{}

func (e OnlyOneArgumentIsAllowedError) Error() string {
	return "only one argument is allowed"
}

// Tmpl provides enhanced template execution capabilities.
type Tmpl rootContext

// Exec executes a named template with the provided data and returns the result as a string.
//
// Example 1:
//
//	{{ define "T1" }}Hello {{ . }}{{ end }}
//	{{ $result := tmpl.Exec "T1" "World" }}
//	Message: {{ $result }} // Output: Message: Hello World
func (ctx Tmpl) Exec(name string, data ...any) (any, error) {
	var buf bytes.Buffer
	if len(data) > 1 {
		return nil, OnlyOneArgumentIsAllowedError{}
	} else if len(data) == 0 {
		data = nil
	}
	err := ctx.template.ExecuteTemplate(&buf, name, data)
	if err != nil {
		var retErr ReturnError
		if errors.As(err, &retErr) {
			return retErr.Value, nil
		}
		return "", fmt.Errorf("failed to execute partial template %q: %w", name, err)
	}
	return buf.String(), nil
}
