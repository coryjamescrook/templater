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
	t.Parallel()
	subject := Load

	tests := []struct {
		name               string
		setupFunc          func()
		expectTemplatePath func(conf *Config) bool
		expectDefFileName  func(conf *Config) bool
	}{
		{
			name:      "With defaults",
			setupFunc: nil,
			expectTemplatePath: func(conf *Config) bool {
				rx := regexp.MustCompile("templates$")
				return rx.Match([]byte(conf.TemplatesPath))
			},
			expectDefFileName: func(conf *Config) bool {
				return conf.TemplateDefFileName == "template.yaml"
			},
		},
		{
			name: "With ENV Variable overrides",
			setupFunc: func() {
				os.Setenv(TemplatesPathEnvVar, "some/path/here/templates")
				os.Setenv(TemplateDefFilenameEnvVar, "manifest.yaml")
			},
			expectTemplatePath: func(conf *Config) bool {
				return conf.TemplatesPath == "some/path/here/templates"
			},
			expectDefFileName: func(conf *Config) bool {
				return conf.TemplateDefFileName == "manifest.yaml"
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// run setup if necessary
			if test.setupFunc != nil {
				test.setupFunc()
			}

			conf := subject()
			if test.expectTemplatePath == nil {
				t.Log("expected a `expectDefFileName` expectation function")
				t.FailNow()
			}
			if test.expectDefFileName == nil {
				t.Log("expected a `expectDefFileName` expectation function")
				t.FailNow()
			}

			if !test.expectTemplatePath(conf) {
				t.Fail()
			}

			if !test.expectDefFileName(conf) {
				t.Fail()
			}
		})
	}
}
