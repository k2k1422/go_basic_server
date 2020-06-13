package Cache

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"server/DataModels"
	"server/Logging"

	"github.com/go-redis/redis"
)

var AccessToken *redis.Client
var RefreshToken *redis.Client
var FilterStore *redis.Client
var ActionItemStore *redis.Client
var Connection []*redis.Client

func init() {
	path, _ := os.Getwd()
	// Declaring a variable to hold the cache credential data
	var credential DataModels.CacheCredential
	// Opening the credential file present in the config folder
	jsonFile, err := os.Open(filepath.Join(path, "config", "cacheCredential.json"))
	if err != nil {
		// Failed to open the configuration file
		panic(err)
	}
	defer jsonFile.Close()
	// Reading data from the configuration json file
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// Decoding the credential data to previously declared credential variable
	if err = json.Unmarshal(byteValue, &credential); err != nil {
		// Failed to decode the data present in the configuration file
		panic(err)
	}
	// Iterating the config data
	for index, _ := range credential.CacheType {
		// Creating redis connection based on the config details
		Connection = append(Connection,
			redis.NewClient(&redis.Options{
				Addr:     credential.CacheConnectionHost + ":" + credential.CacheConnectionPort,
				Password: credential.CacheConnectionPassword,
				DB:       index,
			}),
		)
	}
	// Iterating the connection array to check the connection status
	for index, _ := range Connection {
		if _, err := Connection[index].Ping().Result(); err != nil {
			// Dead peer
			Logging.ERROR.Println("Failed to establish connection. DB-", index)
			panic(err)
		} else {
			// Connection is active
			Logging.INFO.Println("Successfully established connection. Cache connection-", index)
		}
	}
	// Assigning the created connections to the pre declared redis client connection
	AccessToken = Connection[0]
	RefreshToken = Connection[1]
	FilterStore = Connection[2]
	ActionItemStore = Connection[3]

}
