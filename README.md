# readme

## run
```bash
cd cmd
go build -o goi18n main.go tmpl.go
mv goi18n $GOPATH/bin
```

## test
```bash
cd examples
goi18n --from etc --nav=etc/nav.yaml --to internal/i18n
```