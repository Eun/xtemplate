package xtemplate_test

import (
	"os"
	"text/template"

	"github.com/Eun/xtemplate"
	"github.com/Eun/xtemplate/funcs"
)

func Example() {
	// Create a template with restricted functions
	t := template.New("example")
	t = t.Funcs(xtemplate.FuncMap(t,
		funcs.Strings,     // Allow all string functions
		funcs.URLJoinPath, // Allow only url.JoinPath function
	))

	tmpl := `{{ strings.ToLower (url.JoinPath "https://example.com" "INDEX.HTML") }}`
	t, err := t.Parse(tmpl)
	if err != nil {
		panic(err)
	}

	err = xtemplate.Execute(t, os.Stdout, nil)
	if err != nil {
		panic(err)
	}
	// Output: https://example.com/index.html
}
