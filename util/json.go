package util

import "encoding/json"
import "io/ioutil"
import "io"

func ReadJSONInto(r io.ReadCloser, data interface{}) error {
	defer r.Close()
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, data)
}
