package static

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port int    `yaml:"port"`
	Key  string `yaml:"key"`
}

func NewConfig(path string) *Config {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicln(err.Error())
	}
	data := make(map[string]Config)
	if err := yaml.Unmarshal(file, &data); err != nil {
		log.Panicln(err.Error())
	}
	config := data["config"]
	return &config
}
