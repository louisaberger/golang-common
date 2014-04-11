package util

import (
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"os"
)

func ReadYamlFromFile(filename string, result interface{}) error {
	rawFileData, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(rawFileData, result)
}

func WriteYamlToFile(filename string, data interface{}, perm os.FileMode) error {
	rawData, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, rawData, perm)
}
