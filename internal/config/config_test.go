package config

import (
	"os"
	"regexp"
	"testing"
)

func TestCleanup(t *testing.T) {
	t.Cleanup(func() {
		os.Setenv(TemplatesPathEnvVar, "")
		os.Setenv(TemplateDefFilenameEnvVar, "")
	})
}

func TestLoad(t *testing.T) {
	subject := Load

	t.Run("With defaults", func(t *testing.T) {
		conf := subject()

		rx := regexp.MustCompile("templates$")
		templatePathEndsCorrectly := rx.Match([]byte(conf.TemplatesPath))

		if !templatePathEndsCorrectly {
			t.Fail()
		}

		if conf.TemplateDefFileName != "template.yaml" {
			t.Fail()
		}
	})

	t.Run("With ENV Variable overrides", func(t *testing.T) {
		tp := "some/path/here/templates"
		os.Setenv(TemplatesPathEnvVar, tp)

		td := "manifest.yaml"
		os.Setenv(TemplateDefFilenameEnvVar, td)

		conf := subject()

		if conf.TemplatesPath != tp {
			t.Fail()
		}

		if conf.TemplateDefFileName != td {
			t.Fail()
		}
	})
}
