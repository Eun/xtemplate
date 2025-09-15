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

func Example_second() {
	result, err := xtemplate.QuickExecute(
		"{{ filepath.ToSlash ( filepath.Join ( os.Getenv \"HOME\" ) .file ) }}",
		map[string]any{"file": ".bashrc"},
		funcs.Safe,     // Safe functions
		funcs.OSGetenv, // Additional OS function
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(result) // Output: /root/.bashrc
}
