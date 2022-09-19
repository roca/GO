package server

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func ReadConfig(bytes []byte) Config {
	var config Config

	err := yaml.Unmarshal(bytes, &config)
	if err != nil {
		config.LoadError = err
	}

	b,_ := yaml.Marshal(config)

	fmt.Printf("Config loaded:\n%v\n-------------", string(b))

	return config
}
