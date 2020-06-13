package Crypt

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"server/Logging"
)

// Declaring the Key variable for global access
var Key string

func init() {
	/*
		Initializing the secret key from the  crypt.key file
	*/

	// Opening the file containing the key
	path, _ := os.Getwd()
	file, err := os.Open(filepath.Join(path, "config", "crypt.key"))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// Reading the data from the file pointer
	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	// Storing the key in previously declared variable for global access
	Key = string(data)
	Logging.INFO.Println("Crypt key loaded successfully")
}
