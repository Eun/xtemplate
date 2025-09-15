package xtemplate_test

import (
	"errors"
	"fmt"

	"github.com/Eun/xtemplate"
	"github.com/Eun/xtemplate/funcs"
)

func Example_fifth() {
	tmpl := `
{{- if not .user -}}
{{ error "No user provided" }}
{{- end -}}
Welcome, {{ .user.name }}!
`

	result, err := xtemplate.QuickExecute(tmpl, map[string]any{}, funcs.Safe)
	if err != nil {
		var e xtemplate.CustomError
		if ok := errors.As(err, &e); ok {
			fmt.Println("Error:", e.Message) // Output: Error: No user provided
		} else {
			panic(err)
		}
	}
	fmt.Println(result)
}
