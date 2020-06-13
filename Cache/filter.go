package Cache

import (
	"encoding/json"
	"server/Logging"

	"github.com/go-redis/redis"
)

func SetFilter(uid string, filter map[string][]interface{}) bool {
	if _, err := FilterStore.Get(uid).Result(); err == redis.Nil {
		Logging.WARN.Println("Filters does not exist in the cache. Storing the filters")
	} else {
		Logging.WARN.Println("Filters found in the cache, Flushing the data")
		FilterStore.Del(uid)
	}
	if filterByteData, err := json.Marshal(filter); err != nil {
		Logging.ERROR.Println("Failed to encode the filters")
		return false
	} else {
		if _, err := FilterStore.SetNX(uid, string(filterByteData), 0).Result(); err != nil {
			Logging.ERROR.Println("Failed to store the filters in the cache")
			return false
		} else {
			Logging.INFO.Println("Successfully stored the filters in the cache")
			return true
		}
	}
}

func GetFilter(uid string) map[string][]interface{} {
	var filter = make(map[string][]interface{})
	if previousFilters, err := FilterStore.Get(uid).Result(); err == redis.Nil {
		Logging.WARN.Println("No previous filter found for the user")
	} else {
		if err := json.Unmarshal([]byte(previousFilters), &filter); err != nil {
			Logging.ERROR.Println("Failed to decode the previous filters of the requested user")
		} else {
			Logging.INFO.Println("Successfully decoded the previous filter of the requested user")
		}
	}
	return filter
}
