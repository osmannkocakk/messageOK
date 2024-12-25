package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	MySQL struct {
		Host     string `yaml:"localhost"`
		Port     string `yaml:"3306"`
		User     string `yaml:"root"`
		Password string `yaml:"Insider1GizliSifre*"`
		DBName   string `yaml:"insider_db"`
	} `yaml:"mysql"`
	Redis struct {
		Host string `yaml:"localhost"`
		Port string `yaml:"6379"`
	} `yaml:"redis"`
	Message struct {
		ApiUrl      string `yaml:"api_url"`
		HeaderKey   string `yaml:"header_key"`
		HeaderValue string `yaml:"header_value"`
	} `yaml:"message"`
}

func LoadConfig(path string) (*Config, error) {
	var config Config
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
