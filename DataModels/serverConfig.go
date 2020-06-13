package DataModels

// Model to hold the data extracted from the databaseCredential config file
type DBCredential struct {
	DBUser           string `json:"db_user"`            // User of the database connection
	DBPassword       string `json:"db_password"`        // Password for the database connection
	DBConnectionType string `json:"db_connection_type"` // Type of database connection (TCP)
	DBHost           string `json:"db_host"`            // Host of the database
	DBPort           string `json:"db_port"`            // Port in which the database server is running
	DBName           string `json:"db_name"`
}

// Model to hold the data extracted from the cacheCredential config file
type CacheCredential struct {
	CacheType               []string `json:"cache_type"` // Type of cache
	CacheConnectionType     string   `json:"cache_connection_type"`
	CacheConnectionHost     string   `json:"cache_connection_host"`
	CacheConnectionPort     string   `json:"cache_connection_port"`
	CacheConnectionPassword string   `json:"cache_connection_password"`
}

type Response struct {
	ResponseCode    string      `json:"response_code"`
	ResponseMessage string      `json:"response_message"`
	Data            interface{} `json:"data"`
}

type WSGI map[string]string
