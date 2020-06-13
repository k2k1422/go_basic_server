package Logging

import (
	"log"
	"os"
	"path/filepath"
)

var (
	REQ   *log.Logger // For logging the api requests
	INFO  *log.Logger // For logging the application info messages
	WARN  *log.Logger // For logging the application warning messages
	ERROR *log.Logger // For logging the application error messages
	// DB    *log.Logger // For logging the database operations
)

func init() {
	/*
		This will initialize the log files
		application.log -> INFO, WARN and ERROR logs
		request.log -> REQ logs
		db.log -> DB logs
		operations.log -> application operation logs
	*/

	// Creating all the log files in append mode

	path, _ := os.Getwd()
	applicationLog, err := os.OpenFile(filepath.Join(path, "log", "application.log"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	requestLog, err := os.OpenFile(filepath.Join(path, "log", "request.log"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	// dbLog, err := os.OpenFile(filepath.Join(path, "log", "db.log"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	// if err != nil {
	// 	panic(err)
	// }

	// Creating new log instances with output to file pointer as per above documented division
	REQ = log.New(requestLog, "", log.LstdFlags|log.Lshortfile)
	INFO = log.New(applicationLog, "INFO ", log.LstdFlags|log.Lshortfile)
	WARN = log.New(applicationLog, "WARN ", log.LstdFlags|log.Lshortfile)
	ERROR = log.New(applicationLog, "ERROR ", log.LstdFlags|log.Lshortfile)
	// DB = log.New(dbLog, "", log.LstdFlags|log.Lshortfile)

}
