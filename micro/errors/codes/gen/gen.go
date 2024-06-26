package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v3"
)

var (
	cfgFile, outFile, pkgName string
)

func main() {
	flag.StringVar(&cfgFile, "c", "errors.yml", "The error configuration file used to generate go file.")
	flag.StringVar(&outFile, "o", "error.gen.go", "The output go filepath.")
	flag.StringVar(&pkgName, "p", "example", "The go package name.")
	flag.Parse()

	// 读取 error 配置文件
	data, _ := os.ReadFile(cfgFile)
	doc := struct {
		PkgName      string
		ErrConstList []ErrConst
	}{
		PkgName:      pkgName,
		ErrConstList: nil,
	}
	err := yaml.Unmarshal(data, &(doc.ErrConstList))
	die(err)

	// 写入目标 go 文件
	f, err := os.Create(outFile)
	die(err)
	defer f.Close()

	err = errorTemplate.Execute(f, doc)
	die(err)
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type ErrConst struct {
	Code    int
	Key     string
	Msg     string
	Comment string
}

func (c ErrConst) CommentLine() string {
	return fmt.Sprintf("// %s %s", c.Key, c.Comment)
}

var errorTemplate = template.Must(template.New("").Parse(`// Code generated by go generate. DO NOT EDIT.

package {{.PkgName}}

import (
    "fmt"
    "errors"
)

const (
    {{- range .ErrConstList }}
    {{.CommentLine}}
    {{.Key}} Code = {{.Code}}
    {{""}}
    {{- end }}
)

var codeToMsg = map[Code]string{
    {{- range .ErrConstList }}
    {{.Key}} : {{.Msg | printf "%q" }},
    {{- end }} 
}
`))
