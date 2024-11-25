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
)

var (
	wd, _                       = os.Getwd()
	defaultTemplatesPath string = filepath.Join(wd, "templates")
)

func loadTemplatesPath() string {
	envVal := os.Getenv("TEMPLATES_PATH")
	if envVal == "" {
		return defaultTemplatesPath
	}

	return envVal
}

func loadTemplateDefFileName() string {
	envVal := os.Getenv("TEMPLATE_DEF_FILENAME")
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
