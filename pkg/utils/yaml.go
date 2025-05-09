package utils

import (
	"os"

	"gopkg.in/yaml.v3"
)

func ReadYaml(v interface{}, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(v)
	if err != nil {
		return err
	}
	return nil
}

func WriteYaml(v interface{}, path string) error {
	yamlData, err := yaml.Marshal(v)
	if err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(yamlData)
	return err
}
