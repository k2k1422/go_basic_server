package DataAccess

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"server/DataModels"
	"server/Logging"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/go-playground/validator.v9"
)

// Declaring the Database client connection
var Connection *gorm.DB

// Declaring the Validator object
var Validator *validator.Validate

func init() {
	/*
		Initializing one db client connection.
		Assigning the connections to previously declared db-client connection object for global access
	*/

	// Declaring a variable to hold the data present in the databaseCredential.json file
	var credential DataModels.DBCredential

	path, _ := os.Getwd()
	// Opening the credential file present in the config folder
	jsonFile, err := os.Open(filepath.Join(path, "config", "databaseCredential.json"))
	if err != nil {
		Logging.ERROR.Println("Failed to open the database configuration file")
		panic(err)
	}
	defer jsonFile.Close()
	// Reading data from the configuration json file
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// Decoding the credential data to previously declared credential variable
	if err = json.Unmarshal(byteValue, &credential); err != nil {
		Logging.ERROR.Println("Failed to decode the database connection credentials")
		panic(err)
	}
	// Constructing the database URI based on the credential
	URI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		credential.DBHost, credential.DBPort, credential.DBUser, credential.DBName, credential.DBPassword)

	// Creating a db client connection
	//Connection, err = gorm.Open("postgres", "postgres://aptus:aptus@localhost/ficosa?sslmode=disable")
	Connection, err = gorm.Open("postgres", URI)
	if err != nil {
		Logging.ERROR.Println("Failed to connect to the database")
		panic(err)
	} else {
		Logging.INFO.Println("Successfully connected to the database")
	}

	// Enabling the Logging mode for the database operations
	Connection.LogMode(true)
	// Providing a logging pipeline to keep the db logs
	Connection.SetLogger(Logging.DB)
	// Setting number of allowed connection to the database
	Connection.DB().SetMaxOpenConns(5)
	// Setting number of allowed idle connection to the database
	Connection.DB().SetMaxIdleConns(2)

	// Creating and assigning new struct validator
	Validator = validator.New()

	ValidateSchema()
}
