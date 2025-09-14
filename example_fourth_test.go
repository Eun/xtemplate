package xtemplate_test

import (
	"fmt"

	"github.com/Eun/xtemplate"
	"github.com/Eun/xtemplate/funcs"
)

func Example_fourth() {
	tmpl := `
{{- define "getName" -}}
{{- if not .lastname -}}{{- return .firstname -}}{{- end -}}
{{- if eq .lastname "" -}}{{- return .firstname -}}{{- end -}}
{{ .lastname }}, {{ .firstname }}
{{- end -}}

Users:
{{ range .users -}}
* {{ tmpl.Exec "getName" . }}
{{ end -}}
`
	data := map[string]any{
		"users": []map[string]any{
			{"firstname": "Joe", "lastname": "Doe"},
			{"firstname": "Alice"},
		},
	}
	result, err := xtemplate.QuickExecute(tmpl, data, funcs.Safe)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	// Output: Users:
	// * Doe, Joe
	// * Alice
}
