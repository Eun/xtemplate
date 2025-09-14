//go:build ignore

package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/Eun/xtemplate"
	"github.com/Eun/xtemplate/funcs"
)

func main() {
	dstFile, err := os.Create("README.md")
	if err != nil {
		panic(fmt.Errorf("failed to open destination file: %w", err))
	}
	defer dstFile.Close()

	tmpl := template.New("template")
	tmpl = tmpl.Funcs(xtemplate.FuncMap(tmpl, funcs.Safe, funcs.OSReadFile))
	tmpl, err = tmpl.ParseFiles("README.md.tmpl")
	if err != nil {
		panic(fmt.Errorf("failed to parse template: %w", err))
	}
	err = xtemplate.ExecuteTemplate(tmpl, dstFile, "README.md.tmpl", nil)
	if err != nil {
		panic(fmt.Errorf("failed to execute template: %w", err))
	}

	fmt.Println("Generated funcs/funcs.gen.go")
}
