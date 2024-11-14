package templaters

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"text/template"

	"github.com/coryjamescrook/templater/internal/config"
	"gopkg.in/yaml.v3"
)

const (
	templateDefFileName string = "template.yaml"
)

var conf = config.Load()

type Templater struct{}

type TemplateBuildData map[string]interface{}

type Template struct {
	def  *TemplateDef
	path string
	data TemplateBuildData
}

func transformInput(inputType string, input string) interface{} {
	var transformed interface{}

	switch inputType {
	case "string":
		transformed = input
	case "boolean":
		t, err := strconv.ParseBool(input)
		if err != nil {
			panic(err)
		}
		transformed = t
	case "integer":
		i, err := strconv.ParseInt(input, 10, 0)
		if err != nil {
			panic(err)
		}
		transformed = i
	default:
		log.Fatalf("%s does not have transform logic configured", inputType)
	}

	return transformed
}

func (t *Template) CollectData() {
	d := TemplateBuildData{}

	// cli prompt to collect data properties
	for propName, propDef := range t.def.DataSchema.Properties {
		// collect input for value
		var input string
		fmt.Printf("Enter %s value for %s: ", propDef.Type, propName)
		fmt.Scanln(&input)

		// set this value for the template data
		d[propName] = transformInput(propDef.Type, input)
	}
	// once data is collected, set the `data` value on the `Template` instance
	t.data = d
}

func (t Template) Build(buildDir string) {
	fsys := os.DirFS(t.path)
	log.Printf("fsys: %s\n", fsys)
	log.Printf("conf: %s", conf)
	log.Printf("buildDir: %s", buildDir)

	// load all the files recursively in the template directory
	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		log.Printf("WALK DIR inter:\npath: %s\nfilename: %s\n", path, d.Name())
		// return err if there is an error
		if err != nil {
			return err
		}

		// return `nil` if this is a dir
		if d.IsDir() {
			return nil
		}

		// skip template def files
		if d.Name() == templateDefFileName {
			return nil
		}

		// skip if the file doesn't have a .template extension
		if m, _ := regexp.Match(".template", []byte(d.Name())); !m {
			return nil
		}

		// setup files
		templateFilePath := filepath.Join(t.path, d.Name())

		newFilePath := strings.Replace(templateFilePath, ".template", "", 1)
		newFilePath = strings.Replace(newFilePath, t.path, buildDir, 1)
		log.Printf("transformed new file path: %s\n", newFilePath)

		if _, err := os.Stat(newFilePath); os.IsNotExist(err) {
			os.MkdirAll(newFilePath, os.ModePerm)
		}

		f, err := os.Create(newFilePath)
		if err != nil {
			return err
		}
		defer f.Close()

		fileWriter := bufio.NewWriter(f)

		// do templating here
		tmpl := template.New(t.def.Name)
		log.Printf("opening template file path: %s\n", templateFilePath)
		templateFile, err := os.ReadFile(templateFilePath)
		if err != nil {
			return err
		}
		tmpl.Parse(string(templateFile))
		err = tmpl.Execute(fileWriter, t.data)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
}

type TemplateDefDataSchemaProperty struct {
	Type     string `yaml:"type"`
	Required bool   `yaml:"required"`
}

type TemplateDefDataSchema struct {
	Properties map[string]TemplateDefDataSchemaProperty `yaml:"properties"`
}

type TemplateDef struct {
	Name       string                `yaml:"name"`
	DataSchema TemplateDefDataSchema `yaml:"data_schema"`
}

func (tr Templater) CreateTemplate(templateName string) *Template {
	templatePath := filepath.Join(conf.TemplatesPath, templateName)
	templateDefFilePath := filepath.Join(templatePath, templateDefFileName)

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
