package xtemplate_test

import (
	"fmt"
	"os"

	"github.com/Eun/xtemplate"
	"github.com/Eun/xtemplate/funcs"
)

func init() {
	os.Setenv("HOME", "/root")
}

func ExampleQuickExecute() {
	s, _ := xtemplate.QuickExecute(
		"Hello {{strings.ToLower .name}}",
		map[string]any{"name": "Joe"},
		funcs.Safe,
	)
	fmt.Println(s) // Output: Hello joe
}

func ExampleQuickExecute_second() {
	s, _ := xtemplate.QuickExecute(
		"{{filepath.ToSlash (filepath.Join (os.Getenv \"HOME\") .file)}}",
		map[string]any{"file": ".bashrc"},
		funcs.Safe,     // Allow all safe functions
		funcs.OSGetenv, // Allow only the os.Getenv function
	)
	fmt.Println(s) // Output: /root/.bashrc
}
