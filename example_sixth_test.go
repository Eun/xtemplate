package xtemplate_test

import (
	"fmt"

	"github.com/Eun/xtemplate"
	"github.com/Eun/xtemplate/funcs"
)

func Example_sixth() {
	tmpl := `
{{- define "getName" -}}
	{{- if not .user -}}
		{{ return "Anonymous" }}
	{{- end -}}
	{{- return .user.name -}}
{{- end -}}
Welcome, {{ tmpl.Exec "getName" . }}!
`

	result, err := xtemplate.QuickExecute(tmpl, map[string]any{}, funcs.Safe)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	// Output: Welcome, Anonymous!
}
