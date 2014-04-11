package util

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

func ReadJSONInto(r io.ReadCloser, result interface{}) error {
	defer r.Close()
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, result)
}

func ReadJsonFromFile(filename string, result interface{}) error {
	rawFileData, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(rawFileData, result)
}

// Writes data to filename as "pretty" JSON.
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

	return ioutil.WriteFile(filename, buf.Bytes(), perm)
}
