package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ContentDir  string `yaml:"content_dir"`
	OutputDir   string `yaml:"output_dir"`
	TemplateDir string `yaml:"template_dir"`
}

func Load(path string) (Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	cf := Config{}

	err = yaml.Unmarshal(file, &cf)
	if err != nil {
		return Config{}, err
	}

	return cf, nil
}
