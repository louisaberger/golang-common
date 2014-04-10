package util

import (
	"github.com/go-yaml/yaml"

	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func ReadJsonFromFile(filename string, result interface{}) error {
	rawFileData, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(rawFileData, result); err != nil {
		return err
	}
	return nil
}

// Writes mapData to filename as "pretty" JSON.
func WriteJsonToFile(filename string, data interface{}, perm os.FileMode) error {
	rawData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// make json "pretty"
	var buf bytes.Buffer
	err = json.Indent(&buf, rawData, "", "    ")
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(filename, buf.Bytes(), perm); err != nil {
		return err
	}
	return nil
}

func ReadYamlFromFile(filename string, result interface{}) error {
	rawFileData, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(rawFileData, result); err != nil {
		return err
	}
	return nil
}

func WriteYamlToFile(filename string, data interface{}, perm os.FileMode) error {
	rawData, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(filename, rawData, perm); err != nil {
		return err
	}
	return nil
}

func FileExists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return true, err
	}
	return true, nil
}

func IsDir(path string) (bool, error) {
	fileinfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return fileinfo.IsDir(), nil
}
