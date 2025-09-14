package xtemplate_test

import (
	"fmt"

	"github.com/Eun/xtemplate"
	"github.com/Eun/xtemplate/funcs"
)

func Example_fifth() {
	tmpl := `
{{- if not .user -}}
{{ return "Error: No user provided" }}
{{- end -}}
Welcome, {{ .user.name }}!
`

	result, err := xtemplate.QuickExecute(tmpl, map[string]any{}, funcs.Safe)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	// Output: Error: No user provided
}
