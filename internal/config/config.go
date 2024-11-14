package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	TemplatesPath string
	WD            string
}

var (
	wd, _                       = os.Getwd()
	defaultTemplatesPath string = filepath.Join(wd, "templates")
)

func templatesPath() string {
	envVal := os.Getenv("TEMPLATES_PATH")
	if envVal == "" {
		return defaultTemplatesPath
	}

	return envVal
}

func Load() *Config {
	c := Config{
		TemplatesPath: templatesPath(),
		WD:            wd,
	}

	return &c
}
