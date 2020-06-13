package Response

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"server/Logging"
)

var Status map[string]string

func init() {

	path, _ := os.Getwd()
	jsonFile, err := os.Open(filepath.Join(path, "config", "statusMessageMap.json"))
	if err != nil {
		Logging.ERROR.Println("Could not open the statusMessageMap.json file. ", err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &Status)
	if err != nil {
		Logging.ERROR.Println("Could not decode the statusMessageMap.json data")
	}
	Logging.INFO.Println("Loaded the data of statusMessageMap.json successfully")
}
