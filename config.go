package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func readConfig(filename string) YamlConfig {
	yamlFile, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	var data YamlConfig

	err = yaml.Unmarshal(yamlFile, &data)

	if err != nil {
		log.Fatal(err)
	}

	return data
}
