package main

import "text/template"

var (
	i18nTmpl = `package {{.Package}}

import (
	"github.com/meilihao/goi18n/v2"
)

var (
	{{- range .Structs}}
	{{.Name}} = struct {
		{{- range .Fileds}}
		{{.Name}} *goi18n.Elem
		{{- end}}
	}{
		{{ $parent := . }}
		{{- range .Fileds}}
		{{.Name}}: &goi18n.Elem{
			Key: "{{$parent.Name}}.{{.Name}}",
			Map: map[string]string{
				{{- range .Lists}}
				"{{.K}}":{{.V}},
				{{- end}}
			},
		},
		{{- end}}
	}

	{{- end}}
)
`

	i18nTmplImpl = template.Must(template.New("").Parse(i18nTmpl))
)

type TmplData struct {
	Package string

	Structs []*Struct
}

type Struct struct {
	Name   string
	Fileds []*Field
}

type Field struct {
	Name  string
	Lists []*KV
}
type KV struct {
	K, V string
}
