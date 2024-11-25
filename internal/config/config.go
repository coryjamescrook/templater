package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	TemplatesPath       string
	TemplateDefFileName string
}

const (
	defaultTemplateDefFileName string = "template.yaml"
	TemplatesPathEnvVar        string = "TEMPLATES_PATH"
	TemplateDefFilenameEnvVar  string = "TEMPLATE_DEF_FILENAME"
)

var (
	wd, _                       = os.Getwd()
	defaultTemplatesPath string = filepath.Join(wd, "templates")
)

func loadTemplatesPath() string {
	envVal := os.Getenv(TemplatesPathEnvVar)
	if envVal == "" {
		return defaultTemplatesPath
	}

	return envVal
}

func loadTemplateDefFileName() string {
	envVal := os.Getenv(TemplateDefFilenameEnvVar)
	if envVal == "" {
		return defaultTemplateDefFileName
	}

	return envVal
}

func Load() *Config {
	c := Config{
		TemplatesPath:       loadTemplatesPath(),
		TemplateDefFileName: loadTemplateDefFileName(),
	}

	return &c
}
