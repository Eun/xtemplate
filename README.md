# xtemplate
[![Actions Status](https://github.com/Eun/xtemplate/workflows/push/badge.svg)](https://github.com/Eun/xtemplate/actions)
[![Coverage Status](https://coveralls.io/repos/github/Eun/xtemplate/badge.svg?branch=main)](https://coveralls.io/github/Eun/xtemplate?branch=main)
[![PkgGoDev](https://img.shields.io/badge/pkg.go.dev-reference-blue)](https://pkg.go.dev/github.com/Eun/xtemplate)
[![go-report](https://goreportcard.com/badge/github.com/Eun/xtemplate)](https://goreportcard.com/report/github.com/Eun/xtemplate)
---

A secure Go template engine with a restricted set of functions for safe template execution. **xtemplate** provides enhanced control over template functionality by allowing you to specify which functions are permitted, making it ideal for scenarios where you need to execute untrusted templates safely.

## Features

- 🔒 **Security-focused**: Restrict which functions can be used in templates
- 🚀 **Rich function library**: Includes functions for strings, paths, JSON, conversions, and more
- 🎯 **Fine-grained control**: Allow entire namespaces or individual functions
- ⚡ **Quick execution**: Simple API for common use cases
- 🔄 **Early return**: Special `return` function to short-circuit template execution
- ⚠️ **Custom error**: Special `error` function to stop template execution and return an error
- 📦 **Zero dependencies**: Built on Go's standard library

## Installation

```bash
go get github.com/Eun/xtemplate
```

## Quick Start

### Basic Usage

```go
package main

import (
	"os"
	"text/template"

	"github.com/Eun/xtemplate"
	"github.com/Eun/xtemplate/funcs"
)

func main()
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
```

### Quick Execute (One-liner)

For simple cases, use `QuickExecute`:

```go
package main

import (
	"fmt"
	"os"

	"github.com/Eun/xtemplate"
	"github.com/Eun/xtemplate/funcs"
)

func init() {
	os.Setenv("HOME", "/root")
}

func main()
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
```

## Available Function Namespaces

| Namespace  | Description               | Example Functions                                |
|------------|---------------------------|--------------------------------------------------|
| [strings](https://pkg.go.dev/github.com/Eun/xtemplate#Strings)  | String manipulation       | `ToLower`, `ToUpper`, `Replace`, `Split`, `Join` |
| [conv](https://pkg.go.dev/github.com/Eun/xtemplate#Conv)     | Type conversions          | `ToString`, `ToInt`, `ToBool`, `ToFloat64`       |
| [json](https://pkg.go.dev/github.com/Eun/xtemplate#Json)     | JSON operations           | `Marshal`, `Unmarshal`, `Valid`                  |
| [filepath](https://pkg.go.dev/github.com/Eun/xtemplate#Filepath) | File path operations      | `Join`, `Dir`, `Base`, `Ext`, `Clean`            |
| [path](https://pkg.go.dev/github.com/Eun/xtemplate#Path)     | URL path operations       | `Join`, `Dir`, `Base`, `Ext`, `Clean`            |
| [dict](https://pkg.go.dev/github.com/Eun/xtemplate#Dict)     | Dictionary/map operations | `New`, `HasKey`, `HasValue`, `Keys`              |
| [slice](https://pkg.go.dev/github.com/Eun/xtemplate#Slice)    | Slice operations          | `New`, `Sort`, `Reverse`, `Contains`             |
| [url](https://pkg.go.dev/github.com/Eun/xtemplate#Url)      | URL operations            | `JoinPath`, `QueryEscape`                        |
| [os](https://pkg.go.dev/github.com/Eun/xtemplate#Os)       | OS operations             | `Getenv`, `Hostname`, `UserHomeDir`              |
| [tmpl](https://pkg.go.dev/github.com/Eun/xtemplate#Tmpl)     | Template operations       | `Exec`                                           |
| [regexp](https://pkg.go.dev/github.com/Eun/xtemplate#Regexp)   | Regular expressions       | `Match`, `ReplaceAll`, `Split`                   |
| [cmp](https://pkg.go.dev/github.com/Eun/xtemplate#Cmp)      | Comparison operations     | `Or`                                             |

## Function Collections

Use predefined collections for common scenarios:

```go
import "github.com/Eun/xtemplate/funcs"

// Safe functions - recommended for untrusted templates
funcs.Safe

// All functions - use with caution
funcs.All

// Individual namespaces
funcs.Strings
funcs.Conv
funcs.JSON
// ... etc

// Individual functions
funcs.StringsToLower
funcs.URLJoinPath
funcs.OSGetenv
// ... etc
```

## Advanced Examples

### Data Processing Template

```go
package main

import (
	"fmt"

	"github.com/Eun/xtemplate"
	"github.com/Eun/xtemplate/funcs"
)

func main()
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
```

### Early Return Example

Use the special `return` function to exit template execution early:

```go
package main

import (
	"fmt"

	"github.com/Eun/xtemplate"
	"github.com/Eun/xtemplate/funcs"
)

func main()
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
```

### Custom Error Example

Use the special `error` function to exit template execution early with an error:

```go
package main

import (
	"errors"
	"fmt"

	"github.com/Eun/xtemplate"
	"github.com/Eun/xtemplate/funcs"
)

func main()
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
```

### Template Inclusion

Define and include sub-templates:

```go
package main

import (
	"fmt"

	"github.com/Eun/xtemplate"
	"github.com/Eun/xtemplate/funcs"
)

func main()
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
```

## Security Considerations

**xtemplate** is designed for secure template execution:

- ✅ **Use `funcs.Safe`** for untrusted templates
- ✅ **Whitelist specific functions** when you need more control
- ⚠️ **Be cautious with `funcs.All`** - includes potentially dangerous functions
- ⚠️ **OS functions** like `os.Getenv` can expose sensitive information
- ⚠️ **Template functions** like `tmpl.Execute` can lead to infinite recursion

### Safe vs All Functions

```go
// Safe for untrusted templates
funcs.Safe // Includes: strings, conv, json, filepath, path, dict, slice, url, tmpl, cmp

// Use with caution - includes OS functions and more
funcs.All // Includes everything, including os.Getenv, os.Hostname, etc.
```

## Error Handling

```go
result, err := xtemplate.QuickExecute(template, data, funcs.Safe)
if err != nil {
    // Handle parsing or execution errors
    var funcErr *xtemplate.FuncNotAllowedError
    if errors.As(err, &funcErr) {
        fmt.Printf("Function not allowed: %s.%s\n", funcErr.Func.Namespace, funcErr.Func.Name)
    }
}
```

## Complete Function List
See the [GoDoc](https://pkg.go.dev/github.com/Eun/xtemplate) for a complete list of available functions and their descriptions.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
