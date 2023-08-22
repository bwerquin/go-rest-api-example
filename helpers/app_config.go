package helpers

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Add URL config file
const configPath = "config.yml"

// Specific structure for config file
type Cfg struct {
	LOG string `yaml:"logFile"`
}

var AppConfig Cfg

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadConfig() {

	f, err := os.Open(configPath)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&AppConfig)

	if err != nil {
		fmt.Println(err)
	}
}
