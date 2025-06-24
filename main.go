package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	src, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	code := prepareCode(string(src))

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", code, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	ast.Inspect(node, func(n ast.Node) bool {
		typeSpec, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}

		structType, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			return true
		}

		for _, field := range structType.Fields.List {
			if len(field.Names) == 0 {
				continue
			}

			fieldName := field.Names[0].Name
			jsonTag := toLowerCase(fieldName) // Просто нижний регистр

			if field.Tag == nil {
				field.Tag = &ast.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf("`json:\"%s\"`", jsonTag),
				}
			} else if !strings.Contains(field.Tag.Value, "json:") {
				tagValue := strings.Trim(field.Tag.Value, "`")
				field.Tag.Value = fmt.Sprintf("`%s json:\"%s\"`", tagValue, jsonTag)
			}
		}
		return true
	})

	var buf bytes.Buffer
	if err := format.Node(&buf, fset, node); err != nil {
		log.Fatal(err)
	}

	fmt.Print(extractStructDefinition(buf.String()))
}

func prepareCode(src string) string {
	if strings.Contains(src, "type ") && strings.Contains(src, "struct") {
		return "package main\n\n" + src
	}
	return "package main\n\ntype TempStruct struct " + src
}

func extractStructDefinition(code string) string {
	start := strings.Index(code, "type ")
	if start == -1 {
		return code
	}
	end := strings.Index(code[start:], "\n\n")
	if end == -1 {
		return code[start:]
	}
	return code[start : start+end]
}

// Просто преобразует в нижний регистр без разделителей
func toLowerCase(s string) string {
	return strings.ToLower(s)
}
