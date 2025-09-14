package xtemplate

import (
	"bytes"
	"errors"
	"fmt"
)

// Tmpl provides enhanced template execution capabilities.
type Tmpl rootContext

// Exec executes a named template with the provided data and returns the result as a string.
//
// Example 1:
//
//	{{ define "T1" }}Hello {{ . }}{{ end }}
//	{{ $result := tmpl.Exec "T1" "World" }}
//	Message: {{ $result }} // Output: Message: Hello World
func (ctx Tmpl) Exec(name string, data any) (string, error) {
	var buf bytes.Buffer
	err := ctx.template.ExecuteTemplate(&buf, name, data)
	if err != nil {
		var retErr ReturnError
		if errors.As(err, &retErr) {
			return fmt.Sprint(retErr.Value), nil
		}
		return "", fmt.Errorf("failed to execute partial template %q: %w", name, err)
	}
	return buf.String(), nil
}
