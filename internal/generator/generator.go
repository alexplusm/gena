package generator

import (
	"fmt"
	"log"
	"os"
)

type generator struct {
	basePath  string
	templates interface{}
}

type IGenerator interface {
	Gen() (err error)
}

func New(basePath string) IGenerator {
	return &generator{basePath: basePath}
}

func (g *generator) Gen() (err error) {

	tt := new(templates)
	if err = tt.load(); err != nil {
		log.Fatalln(err)
	}

	err = os.MkdirAll(g.basePath+"/cmd", os.ModePerm)

	f, err := os.Create(g.basePath + "/cmd/main.go")
	if _, err = f.Write(tt.mainGoTemplate); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(err, f)

	return
}

type templates struct {
	mainGoTemplate []byte
}

func (t *templates) load() (err error) {
	mainGoTemplatePath := "templates/main.go.tpl"
	tmp, err := os.ReadFile(mainGoTemplatePath)
	if err != nil {
		return err
	}

	t.mainGoTemplate = tmp
	return
}
