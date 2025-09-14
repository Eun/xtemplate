package xtemplate

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"text/template"
)

func finishExecute(err error, wr io.Writer) error {
	if err != nil {
		var retErr ReturnError
		if errors.As(err, &retErr) {
			_, _ = fmt.Fprint(wr, retErr.Value)
			return nil
		}
		return fmt.Errorf("failed to execute template: %w", err)
	}
	return nil
}

// ExecuteTemplate executes the named template within the given template with the provided data and writes
// the result to the given writer.
func ExecuteTemplate(t *template.Template, wr io.Writer, name string, data any) error {
	err := t.ExecuteTemplate(wr, name, data)
	return finishExecute(err, wr)
}

// Execute executes the given template with the provided data and writes the result to the given writer.
func Execute(t *template.Template, wr io.Writer, data any) error {
	err := t.Execute(wr, data)
	return finishExecute(err, wr)
}

// QuickExecute is a convenience function to parse and execute a template string with the given data and
// allowed functions and write the result to the given writer.
func QuickExecute(tmplStr string, data any, allowedFunctions ...AllowedFunctions) (string, error) {
	tmpl := template.New("template")
	tmpl = tmpl.Funcs(FuncMap(tmpl, allowedFunctions...))
	tmpl, err := tmpl.Parse(tmplStr)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}
	var buf bytes.Buffer
	err = Execute(tmpl, &buf, data)
	if err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}
	return buf.String(), nil
}
