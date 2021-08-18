package main

import "text/template"

// var (
// 	Node = struct {
// 		Name map[string]string
// 	}{
// 		Name: map[string]string{
// 			"a": "b",
// 		},
// 	}
// )

// func a() {
// 	fmt.Println(Node.Name)
// }

var (
	i18nTmpl = `package {{.Package}}

var (
	{{- range .Structs}}
	{{.Name}} = struct {
		{{- range .Fileds}}
		{{.Name}} map[string]string
		{{- end}}
	}{
		{{- range .Fileds}}
		{{.Name}}: map[string]string{
			{{- range .Lists}}
			"{{.K}}":{{.V}},
			{{- end}}
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
