package xtemplate_test

import (
	"fmt"

	"github.com/Eun/xtemplate"
	"github.com/Eun/xtemplate/funcs"
)

func Example_third() {
	data := map[string]any{
		"users": []map[string]any{
			{"firstname": "John", "lastname": "Doe", "age": "25", "active": "true"},
			{"firstname": "Jane", "lastname": "Smith", "age": "30", "active": "false"},
		},
	}

	tmpl := `
{{- range .users -}}
Name: {{ strings.Join ( slice.NewStrings .firstname .lastname ) " " }}
Age: {{ conv.ToInt .age }}
Active: {{ conv.ToBool .active }}
---
{{ end -}}
`

	result, err := xtemplate.QuickExecute(tmpl, data, funcs.Safe)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	// Output: Name: John Doe
	// Age: 25
	// Active: true
	// ---
	// Name: Jane Smith
	// Age: 30
	// Active: false
	// ---
}
