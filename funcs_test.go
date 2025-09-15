package xtemplate_test

import (
	"bytes"
	"errors"
	"testing"
	"text/template"

	"github.com/Eun/xtemplate"
	"github.com/Eun/xtemplate/funcs"
)

func TestCommonUseCases(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		tmpl    string
		want    string
		wantErr bool
	}{
		{
			name:    "simple return",
			tmpl:    `Hello {{ return "World" }} and Universe`,
			want:    "Hello World",
			wantErr: false,
		},
		{
			name: "panic",
			tmpl: `Hello {{ panic "oh no" }} and Universe`,
			// Note: the output is "Hello " because the panic happens during execution,
			// so the part before the panic is written to the buffer.
			want:    "Hello ",
			wantErr: true,
		},
		{
			name: "template call",
			tmpl: `
			{{- define "T1" }}Hello {{ . }}{{ end -}}
			{{- tmpl.Exec "T1" "World" -}}`,
			want:    "Hello World",
			wantErr: false,
		},
		{
			name: "template with inexplicit return",
			tmpl: `
			{{- define "T1" }}{{ strings.ToLower . }}{{ end -}}
			{{- $result := tmpl.Exec "T1" "World" -}}
			Hello {{ $result -}}`,
			want:    "Hello world",
			wantErr: false,
		},
		{
			name: "template with explicit return",
			tmpl: `
			{{- define "T1" }}{{ return ( strings.ToLower . ) }}{{ end -}}
			{{- $result := tmpl.Exec "T1" "World" -}}
			Hello {{ $result -}}
			`,
			want:    "Hello world",
			wantErr: false,
		},
		{
			name: "template with return dict",
			tmpl: `
			{{- define "T1" }}{{ return ( dict.New "name" "Joe" ) }}{{ end -}}
			{{- $result := tmpl.Exec "T1" -}}
			Hello {{ $result.name -}}
			`,
			want:    "Hello Joe",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var err error
			tmpl := template.New("template")
			tmpl = tmpl.Funcs(xtemplate.FuncMap(tmpl, funcs.All))
			tmpl = tmpl.Funcs(template.FuncMap{
				"panic": func(msg string) (any, error) {
					panic(msg)
				},
			})
			tmpl, err = tmpl.Parse(tt.tmpl)
			if err != nil {
				t.Errorf("Parse() error = %v", err)
				return
			}

			var buf bytes.Buffer
			err = xtemplate.Execute(tmpl, &buf, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != buf.String() {
				t.Errorf("Execute() got = %v, want %v", buf.String(), tt.want)
			}
		})
	}
}

func TestFuncNotAllowedError_Error(t *testing.T) {
	t.Parallel()

	tmpl := template.New("template")
	tmpl = tmpl.Funcs(xtemplate.FuncMap(tmpl, funcs.OSHostname))
	tmpl, err := tmpl.Parse(`{{ os.Getenv "PATH" }}`)
	if err != nil {
		t.Errorf("Parse() error = %v", err)
		return
	}
	var buf bytes.Buffer
	err = xtemplate.ExecuteTemplate(tmpl, &buf, "template", nil)
	if err == nil {
		t.Errorf("ExecuteTemplate() expected error, got nil")
		return
	}
	var funcNotAllowedError *xtemplate.FuncNotAllowedError
	if !errors.As(err, &funcNotAllowedError) {
		t.Errorf("ExecuteTemplate() error = %v, want FuncNotAllowedError", err)
		return
	}
	if funcNotAllowedError.Func != funcs.OSGetenv {
		t.Errorf("FuncNotAllowedError.Func = %v, want %v", funcNotAllowedError.Func, funcs.OSGetenv)
		return
	}
}
