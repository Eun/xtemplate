package xtemplate_test

import (
	"os"
	"text/template"

	"github.com/Eun/xtemplate"
	"github.com/Eun/xtemplate/funcs"
)

func ExampleFuncMap() {
	t := template.New("template")
	t = t.Funcs(xtemplate.FuncMap(t,
		funcs.Strings,     // allow all functions in the strings namespace
		funcs.URLJoinPath, // allow only the url.JoinPath function
	))
	t, err := t.Parse(`{{ strings.ToLower ( url.JoinPath "https://example.com" "INDEX.HTML" ) }}`)
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, nil) // Output: https://example.com/index.html
	if err != nil {
		panic(err)
	}
}
