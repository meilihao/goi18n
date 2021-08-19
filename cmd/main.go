package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	conf = struct {
		Nav  string
		From string
		To   string
	}{}

	rootCmd = &cobra.Command{
		Use: "i18n",
		Run: rootRun,
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&conf.From, "from", "etc", "store lan files")
	rootCmd.PersistentFlags().StringVar(&conf.Nav, "nav", "nav.yaml", "nav to load lan files")
	rootCmd.PersistentFlags().StringVar(&conf.To, "to", "internal/i18n", "store generated go files")
}

func rootRun(cmd *cobra.Command, args []string) {
	var langs []string

	data, _ := os.ReadFile(filepath.Join(conf.From, conf.Nav))
	if err := yaml.Unmarshal([]byte(data), &langs); err != nil {
		log.Fatal(err)
	}
	if len(langs) == 0 {
		log.Fatal("no lang to generate")
	}
	log.Printf("found lang(%d): %v\n", len(langs), langs)

	rLangs := make(map[string][]*LangSection, len(langs))
	for _, lang := range langs {
		rLang, err := LoadLang(lang)
		if err != nil {
			log.Fatalf("failed to load lang(%s): %v", lang, err)
		}
		rLangs[lang] = rLang
	}

	tData := &TmplData{
		Package: filepath.Base(conf.To),
		Structs: make([]*Struct, 0),
	}

	for _, ls := range rLangs[langs[0]] {
		stru := &Struct{
			Name:   ls.Section,
			Fileds: make([]*Field, 0),
		}
		tData.Structs = append(tData.Structs, stru)

		for i, n := 0, len(ls.KVList); i < n; i += 2 {
			f := &Field{
				Name:  ls.KVList[i],
				Lists: make([]*KV, 0),
			}

			for _, lang := range langs {
				f.Lists = append(f.Lists, &KV{
					K: lang,
					V: GetLangValue(lang, ls.Section, f.Name, rLangs[lang]),
				})
			}

			// fmt.Printf("%+v, %+v\n", f.Lists[0], f.Lists[1])

			stru.Fileds = append(stru.Fileds, f)
		}
	}

	buf := bytes.NewBuffer(nil)
	if err := i18nTmplImpl.Execute(buf, tData); err != nil {
		log.Fatal(err)
	}

	os.MkdirAll(conf.To, 0755)
	if err := os.WriteFile(filepath.Join(conf.To, "i18n.go"), buf.Bytes(), 0644); err != nil {
		log.Fatal(err)
	}

	mycmd := exec.Command("gofmt", "-w", filepath.Join(conf.To, "i18n.go"))
	output, err := mycmd.CombinedOutput()
	if err != nil {
		log.Fatalf("gofmt failed:%s, %s", err.Error(), string(output))
	}
}

func GetLangValue(lang, section, key string, rLang []*LangSection) string {
	for _, sec := range rLang {
		if sec.Section != section {
			continue
		}

		for i, n := 0, len(sec.KVList); i < n; i += 2 {
			if sec.KVList[i] == key {
				return sec.KVList[i+1]
			}
		}

		log.Fatalf("not found key(%s) in section(%s) in lang(%s)", key, section, lang)
	}

	log.Fatalf("not found section(%s) in lang(%s)", section, lang)
	return ""
}

func main() {
	rootCmd.Execute()
}

// Strings2Interfaces []string -> []interface{}
func Strings2Interfaces(ss []string) []interface{} {
	ls := make([]interface{}, len(ss))

	for i := range ss {
		ls[i] = ss[i]
	}

	return ls
}

type LangSection struct {
	Section string
	KVList  []string
}

func LoadLang(lang string) ([]*LangSection, error) {
	fp := filepath.Join(conf.From, lang+".yaml")
	data, _ := os.ReadFile(fp)
	if len(data) == 0 {
		return nil, fmt.Errorf("no data in %s", fp)
	}

	m := make(map[string]map[string]string, 0)
	if err := yaml.Unmarshal([]byte(data), &m); err != nil {
		log.Fatalf("parse yaml(%s) error: %v", fp, err)
	}

	sKeys := make([]string, 0)
	for k := range m {
		sKeys = append(sKeys, k)
	}
	if len(sKeys) == 0 {
		return nil, fmt.Errorf("no setction in %s", lang)
	}
	sort.Strings(sKeys)

	rLang := make([]*LangSection, 0)
	for _, sk := range sKeys {
		rSec := &LangSection{
			Section: sk,
		}

		vKeys := make([]string, 0)

		for sk := range m[sk] {
			vKeys = append(vKeys, sk)
		}
		sort.Strings(vKeys)

		for _, key := range vKeys {
			rSec.KVList = append(rSec.KVList, key, "`"+m[sk][key]+"`")
		}

		rLang = append(rLang, rSec)
	}

	// fmt.Printf("%s %+v\n", lang, rLang[0])

	return rLang, nil
}
