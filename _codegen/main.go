package main

import (
	"flag"
	"go/doc"
	"go/types"
	"log"
	"os"
	"text/template"

	"github.com/stretchr/testify/_codegen/internal/imports"
)

var (
	pkg       = flag.String("assert-path", "github.com/stretchr/testify/assert", "Path to the assert package")
	includeF  = flag.Bool("include-format-funcs", false, "include format functions such as Errorf and Equalf")
	outputPkg = flag.String("output-package", "", "package for the resulting code")
	tmplFile  = flag.String("template", "", "What file to load the function template from")
	out       = flag.String("out", "", "What file to write the source code to")
)

func main() {
	flag.Parse()

	scope, docs, err := parsePackageSource(*pkg)
	if err != nil {
		log.Fatal(err)
	}

	importer, funcs, err := analyzeCode(scope, docs)
	if err != nil {
		log.Fatal(err)
	}

	if err := generateCode(importer, funcs); err != nil {
		log.Fatal(err)
	}
}

func generateCode(importer imports.Importer, funcs []testFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func parseTemplates() (*template.Template, *template.Template, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func outputFile() (*os.File, error) { _ = "STUB: not implemented"; return nil, nil }

func analyzeCode(scope *types.Scope, docs *doc.Package) (imports.Importer, []testFunc, error) {
	_ = "STUB: not implemented"
	return *new(imports.Importer), nil, nil
}

func parsePackageSource(pkg string) (*types.Scope, *doc.Package, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type testFunc struct {
	CurrentPkg string
	DocInfo    *doc.Func
	TypeInfo   *types.Func
}

func (f *testFunc) Qualifier(p *types.Package) string { _ = "STUB: not implemented"; return "" }

func (f *testFunc) Params() string { _ = "STUB: not implemented"; return "" }

func (f *testFunc) ForwardedParams() string { _ = "STUB: not implemented"; return "" }

func (f *testFunc) ParamsFormat() string { _ = "STUB: not implemented"; return "" }

func (f *testFunc) ForwardedParamsFormat() string { _ = "STUB: not implemented"; return "" }

func (f *testFunc) Comment() string { _ = "STUB: not implemented"; return "" }

func (f *testFunc) CommentFormat() string { _ = "STUB: not implemented"; return "" }

func (f *testFunc) CommentWithoutT(receiver string) string { _ = "STUB: not implemented"; return "" }

func requireComment(comment string) string { _ = "STUB: not implemented"; return "" }

func (f *testFunc) CommentRequire() string { _ = "STUB: not implemented"; return "" }

func (f *testFunc) CommentRequireWithoutT(receiver string) string {
	_ = "STUB: not implemented"
	return ""
}

var headerTemplate = `// Code generated with github.com/stretchr/testify/_codegen; DO NOT EDIT.

package {{.Name}}

import (
{{range $path, $name := .Imports}}
	{{$name}} "{{$path}}"{{end}}
)
`

var funcTemplate = `{{.Comment}}
func (fwd *AssertionsForwarder) {{.DocInfo.Name}}({{.Params}}) bool {
	return assert.{{.DocInfo.Name}}({{.ForwardedParams}})
}`
