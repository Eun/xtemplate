//go:build ignore

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Example struct {
	Template string
	Output   string
}

type Method struct {
	Context  string
	Name     string
	Examples []Example
}

func main() {
	// Parse the current package
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, ".", nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("Failed to parse package: %v", err)
	}

	var methods []Method

	// Iterate through all packages
	for _, pkg := range pkgs {
		// Skip the main package (this generator itself)
		if pkg.Name == "main" {
			continue
		}

		// Iterate through all files in the package
		for _, file := range pkg.Files {
			fileMethods := extractMethodsWithExamples(fset, file)
			methods = append(methods, fileMethods...)
		}
	}

	// Sort methods by context and name for consistent output
	sort.Slice(methods, func(i, j int) bool {
		if methods[i].Context != methods[j].Context {
			return methods[i].Context < methods[j].Context
		}
		return methods[i].Name < methods[j].Name
	})

	// Generate test files for each context
	if err := generateExampleTestFiles(methods); err != nil {
		log.Fatalf("Failed to generate example test files: %v", err)
	}

	fmt.Printf("Generated example test files for %d methods\n", len(methods))
}

func extractMethodsWithExamples(fset *token.FileSet, file *ast.File) []Method {
	var methods []Method

	// Find context types (those that extend rootContext)
	contextTypes := make(map[string]bool)
	ast.Inspect(file, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.TypeSpec:
			if ident, ok := node.Type.(*ast.Ident); ok && ident.Name == "rootContext" {
				contextTypes[node.Name.Name] = true
			}
		}
		return true
	})

	// Find methods for these context types and extract examples
	ast.Inspect(file, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.FuncDecl:
			// Check if this is a method (has a receiver)
			if node.Recv != nil && len(node.Recv.List) > 0 {
				recv := node.Recv.List[0]

				// Get receiver type name
				var typeName string
				switch recvType := recv.Type.(type) {
				case *ast.Ident:
					typeName = recvType.Name
				case *ast.StarExpr:
					if ident, ok := recvType.X.(*ast.Ident); ok {
						typeName = ident.Name
					}
				}

				// Check if this is an exported method on a context type
				if node.Name.IsExported() && contextTypes[typeName] {
					examples := extractExamplesFromDoc(node.Doc)
					if len(examples) > 0 {
						methods = append(methods, Method{
							Context:  typeName,
							Name:     node.Name.Name,
							Examples: examples,
						})
					}
				}
			}
		}
		return true
	})

	return methods
}

func extractExamplesFromDoc(doc *ast.CommentGroup) []Example {
	if doc == nil {
		return nil
	}

	var examples []Example

	// Regular expressions to match the new example patterns
	exampleHeaderRe := regexp.MustCompile(`^\s*//\s*Example\s*\d*:?\s*$`)
	codeLineRe := regexp.MustCompile(`^\s*//\s+(.+?)\s*(?://\s*Output:\s*(.+))?$`)
	outputLineRe := regexp.MustCompile(`^\s*//\s*Output:\s*(.+)$`)

	var currentExample *Example
	var codeLines []string

	for _, comment := range doc.List {
		line := comment.Text

		// Check if this line starts a new example
		if exampleHeaderRe.MatchString(line) {
			// Save previous example if it exists and has output
			if currentExample != nil && currentExample.Output != "" {
				examples = append(examples, *currentExample)
			}

			// Start new example
			currentExample = &Example{}
			codeLines = []string{}
			continue
		}

		// Skip if we're not in an example section
		if currentExample == nil {
			continue
		}

		// Skip empty comment lines
		trimmedLine := strings.TrimSpace(strings.TrimPrefix(line, "//"))
		if trimmedLine == "" {
			continue
		}

		// Check for code line (with optional output)
		if matches := codeLineRe.FindStringSubmatch(line); matches != nil {
			codePart := strings.TrimSpace(matches[1])
			if codePart != "" {
				codeLines = append(codeLines, codePart)

				// Check if this line has an output comment
				if len(matches) > 2 && matches[2] != "" {
					output := strings.TrimSpace(matches[2])
					if output != "" {
						// Build the complete template from all code lines
						fullTemplate := strings.Join(codeLines, "\n")
						// Wrap in template syntax if not already wrapped
						if !strings.HasPrefix(fullTemplate, "{{") {
							fullTemplate = "{{ " + fullTemplate + " }}"
						}

						currentExample.Template = fullTemplate
						currentExample.Output = output
						// Don't save yet - wait for next example or end
					}
				}
			}
			continue
		}

		// Check for standalone output line
		if matches := outputLineRe.FindStringSubmatch(line); matches != nil {
			output := strings.TrimSpace(matches[1])
			if output != "" && len(codeLines) > 0 {
				// Build the complete template from all code lines
				fullTemplate := strings.Join(codeLines, "\n")
				// Wrap in template syntax if not already wrapped
				if !strings.HasPrefix(fullTemplate, "{{") {
					fullTemplate = "{{ " + fullTemplate + " }}"
				}

				currentExample.Template = fullTemplate
				currentExample.Output = output
			}
			continue
		}

		// If we encounter a non-example line, stop processing
		if !strings.HasPrefix(trimmedLine, "Example") {
			break
		}
	}

	// Save the last example if it exists and has output
	if currentExample != nil && currentExample.Output != "" {
		examples = append(examples, *currentExample)
	}

	return examples
}

func generateExampleTestFiles(methods []Method) error {
	// Group methods by context
	contextMethods := make(map[string][]Method)
	for _, method := range methods {
		contextMethods[method.Context] = append(contextMethods[method.Context], method)
	}

	// Generate a test file for each context
	for context, ctxMethods := range contextMethods {
		filename := fmt.Sprintf("fn_%s_example_test.go", strings.ToLower(context))

		if err := generateTestFile(filename, context, ctxMethods); err != nil {
			return fmt.Errorf("failed to generate %s: %w", filename, err)
		}

		fmt.Printf("Generated %s\n", filename)
	}

	return nil
}

// numberToWord converts a number (2, 3, 4, etc.) to its word representation
func numberToWord(n int) string {
	words := map[int]string{
		2: "second", 3: "third", 4: "fourth", 5: "fifth",
		6: "sixth", 7: "seventh", 8: "eighth", 9: "ninth", 10: "tenth",
		11: "eleventh", 12: "twelfth", 13: "thirteenth", 14: "fourteenth", 15: "fifteenth",
		16: "sixteenth", 17: "seventeenth", 18: "eighteenth", 19: "nineteenth", 20: "twentieth",
	}

	if word, exists := words[n]; exists {
		return word
	}

	// For numbers beyond 20, fall back to numeric representation
	return fmt.Sprintf("%d", n)
}

func generateTestFile(filename, context string, methods []Method) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Write package declaration and imports
	fmt.Fprintln(file, "// Code generated by generate_examples.go; DO NOT EDIT.")
	fmt.Fprintln(file, "package xtemplate_test")
	fmt.Fprintln(file)
	fmt.Fprintln(file, "import (")
	fmt.Fprintln(file, "\t\"fmt\"\n")
	fmt.Fprintln(file, "\t\"github.com/Eun/xtemplate\"")
	fmt.Fprintln(file, "\t\"github.com/Eun/xtemplate/funcs\"")
	fmt.Fprintln(file, ")")
	fmt.Fprintln(file)

	// Generate example functions
	for _, method := range methods {
		for i, example := range method.Examples {
			var funcName string
			if i == 0 {
				funcName = fmt.Sprintf("Example%s_%s", method.Context, method.Name)
			} else {
				suffix := numberToWord(i + 1)
				funcName = fmt.Sprintf("Example%s_%s_%s", method.Context, method.Name, suffix)
			}

			fmt.Fprintf(file, "func %s() {\n", funcName)
			fmt.Fprintln(file, "\ts, _ := xtemplate.QuickExecute(")
			fmt.Fprintf(file, "\t\t`%s`,\n", example.Template)
			fmt.Fprintln(file, "\t\tnil,")
			fmt.Fprintln(file, "\t\tfuncs.All,")
			fmt.Fprintln(file, "\t)")
			fmt.Fprintf(file, "\tfmt.Println(s) // Output: %s\n", example.Output)

			fmt.Fprintln(file, "}")
			fmt.Fprintln(file)
		}
	}

	return nil
}
