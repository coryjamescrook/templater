package template

import (
	"os"
	"path/filepath"

	"github.com/coryjamescrook/templater/internal/config"
	"gopkg.in/yaml.v2"
)

type Template struct {
	def  *TemplateDef
	path string
	data TemplateBuildData
}

func (t *Template) CollectData() {
	BuildDataForTemplate(t)
}

func (t Template) Build(buildDir string) {

	err := ExecuteTemplate(t, buildDir)

	if err != nil {
		panic(err)
	}
}

func CreateTemplate(templateName string) *Template {
	conf := config.Load()

	templatePath := filepath.Join(conf.TemplatesPath, templateName)
	templateDefFilePath := filepath.Join(templatePath, conf.TemplateDefFileName)

	templateDefFile, err := os.ReadFile(templateDefFilePath)
	if err != nil {
		panic(err)
	}

	tmplDef := &TemplateDef{}
	err = yaml.Unmarshal(templateDefFile, tmplDef)
	if err != nil {
		panic(err)
	}

	return &Template{
		path: templatePath,
		def:  tmplDef,
	}
}
