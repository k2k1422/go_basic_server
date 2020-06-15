package SMTP

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

var Credential map[string]string

func init() {
	path, _ := os.Getwd()
	// Declaring a variable to hold the cache credential data
	// Opening the credential file present in the config folder
	jsonFile, err := os.Open(filepath.Join(path, "config", "smtpCredential.json"))
	if err != nil {
		// Failed to open the configuration file
		panic(err)
	}
	defer jsonFile.Close()
	// Reading data from the configuration json file
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// Decoding the credential data to previously declared credential variable
	if err = json.Unmarshal(byteValue, &Credential); err != nil {
		// Failed to decode the data present in the configuration file
		panic(err)
	}
}
