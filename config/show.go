package config

import (
	"fmt"

	"github.com/vivek-yadav/mango/utils"
	"gopkg.in/yaml.v2"
)

func (conf *Config) Show() string {
	yamlData, err := yaml.Marshal(conf)

	utils.CheckFatal(err, "Error while Marshaling.")

	yamlStr := fmt.Sprintln(" --- YAML with maps and arrays ---")
	yamlStr += fmt.Sprintln(string(yamlData))
	return yamlStr
}
