package template

import (
	"bufio"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/coryjamescrook/templater/internal/config"
)

// loads all the files recursively in the template directory
// and creates new files for the template files
func ExecuteTemplate(t Template, buildDir string) error {
	conf := config.Load()

	fsys := os.DirFS(t.path)

	return fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		// return err if there is an error
		if err != nil {
			return err
		}

		// return `nil` if this is a dir
		if d.IsDir() {
			return nil
		}

		// skip template def files
		if d.Name() == conf.TemplateDefFileName {
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
		newFilePath, _ = filepath.Abs(newFilePath)

		newFileDir := strings.Replace(newFilePath, filepath.Base(newFilePath), "", 1)

		if _, err := os.Stat(newFileDir); os.IsNotExist(err) {
			os.MkdirAll(newFileDir, 0700)
		}

		f, err := os.Create(newFilePath)
		if err != nil {
			return err
		}
		defer f.Close()

		fileWriter := bufio.NewWriter(f)

		// do templating here
		tmpl := template.New(t.def.Name)
		templateFile, err := os.ReadFile(templateFilePath)
		if err != nil {
			return err
		}
		tmpl.Parse(string(templateFile))
		err = tmpl.Execute(fileWriter, t.data)
		if err != nil {
			return err
		}

		fileWriter.Flush()

		return nil
	})
}
