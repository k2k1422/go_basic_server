package Response

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

var Status map[string]string

func init() {

	path, _ := os.Getwd()
	jsonFile, err := os.Open(filepath.Join(path, "config", "statusMessageMap.json"))
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &Status)
	if err != nil {
		panic(err)
	}
}
