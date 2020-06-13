package Cache

import (
	"server/Logging"
	"time"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

func SetAccessToken(uid string, token string) bool {
	/*
		SetAccessToken will be used to store the access_token of a user session in the cache
		The key will be the uid of the user and the value will be the access_token
		Parameters - uid of the user (string) and generated access_token (string)
		Return - True	If the uid and access_token is stored in the cache successfully.
				 False	If the uid and access_token could not be stored in the cache.
	*/

	// Getting the access_token of the user from the cache
	_, err := AccessToken.Get(uid).Result()
	if err == redis.Nil {
		// No key value pair exists in the cache having the provided uid as the key
		Logging.INFO.Println("No access token found for the uid")
	} else {
		// Old access_token is found in the cache
		Logging.INFO.Println("Access token found for the user. Dropping the access token")
		// Dropping the access_token from the cache
		AccessToken.Del(uid)
	}
	// Keeping the uid and access_token as key value pair in the cache
	// The expiry time of the key value pair is 24 hours
	if _, err := AccessToken.SetNX(uid, token, 24*time.Hour).Result(); err != nil {
		// Failed to store the uid and access_token in the cache
		Logging.ERROR.Println("could not set the access token")
		return false
	} else {
		// Successfully stored the uid and access_token key value pair in the cache
		Logging.INFO.Println("Successfully set the access token")
		return true
	}
}

func SetRefreshToken(uid string, token string) bool {
	/*
		SetRefreshToken will be used to store the refresh_token of a user session in the cache
		The key will be the uid of the user and the value will be the refresh_token
		Parameters - uid of the user (string) and generated refresh_token (string)
		Return - True	If the uid and refresh_token is stored in the cache successfully.
				 False	If the uid and refresh_token could not be stored in the cache.
	*/

	// Getting the access_token of the user from the cache
	_, err := RefreshToken.Get(uid).Result()
	if err == redis.Nil {
		// No key value pair exists in the cache having the provided uid as the key
		Logging.INFO.Println("No refresh token found for the uid")
	} else {
		// Old refresh_toke found in the cache
		Logging.INFO.Println("Refresh token found for the user. Dropping the refresh token")
		// Dropping the old refresh_token
		RefreshToken.Del(uid)
	}
	// Keeping the uid and refresh_token as key value pair in the cache
	// The expiry time of the key value pair is 1 week
	if _, err := RefreshToken.SetNX(uid, token, 7*24*time.Hour).Result(); err != nil {
		// Failed to store the uid and refresh_token in the cache
		Logging.ERROR.Println("could not set the refresh token")
		return false
	} else {
		// Successfully stored the uid and refresh_token key value pair in the cache
		Logging.INFO.Println("Successfully set the refresh token")
		return true
	}
}

func VerifyRefreshToken(uid string, refreshToken string) bool {
	/*
		VerifyRefreshToken will be used to validate a refresh_token
		Parameters - uid of the user (string) and refresh_token (string)
		Return - True	If the stored refresh_token matched wit the provided refresh_token
				 False	If the stored refresh_token did not match with the provided refresh_token
	*/

	// Extracting the refresh_token present in the cache
	SessionRefreshToken, err := RefreshToken.Get(uid).Result()
	if err == redis.Nil {
		// No refresh_token found having the key as the provided uid
		Logging.INFO.Println("Refresh token is not present in the cache")
		return false
	} else {
		// refresh_token found in cache having the key as the provided uid
		if SessionRefreshToken == refreshToken {
			// The provided refresh_toke matched with the stored refresh_token
			Logging.INFO.Println("The provided refresh_token matched with the stored refresh_token")
			return true
		} else {
			// The provided refresh_token did not match with the stored refresh_token
			Logging.INFO.Println("The provided refresh_token did not match with the stored refresh_token")
			return false
		}
	}
}

func VerifyAccessToken(uid string, accessToken string) bool {
	/*
		VerifyAccessToken will be used to validate a access_token
		Parameters - uid of the user (string) and access_token (string)
		Return - True	If the stored access_token matched wit the provided access_token
				 False	If the stored access_token did not match with the provided access_token
	*/

	// Extracting the access_token present in the cache
	SessionAccessToken, err := AccessToken.Get(uid).Result()
	if err == redis.Nil {
		// No access_token found in the cache having the key as the provided uid
		Logging.INFO.Println("Access token is not present in the cache")
		return false
	} else {
		// access_token found in the cache having the key as the provided uid
		// Checking the access_token
		if SessionAccessToken == accessToken {
			// The provided access_token matched with the stored access_token
			Logging.INFO.Println("The provided access_token matched with the stored access_token")
			return true
		} else {
			// The provided access_token did not match with the stored access_token
			Logging.INFO.Println("The provided access_token did not match with the stored access_token")
			return false
		}
	}
}

func GetAccessTokenUsingRefreshToken(uid string) (bool, string) {
	/*
		GetAccessTokenUsingRefreshToken will be used to generate a new access_token based on the provided refresh_token
		Parameters - uid of the user (string)
		Return - True, <access_token>	If the operation is successful
				 False, "" 				If the operation is unsuccessful
	*/
	// Getting the access_token of the provided uid from the cache
	_, err := AccessToken.Get(uid).Result()
	if err == redis.Nil {
		// No access_token found having the key as the provided uid
		Logging.WARN.Println("No access token found for the request refresh token ")
	} else {
		// access_token found having the key as the provided uid
		Logging.WARN.Println("Access token found. Dropping the access token")
		// Dropping the access_token
		AccessToken.Del(uid)
	}
	// Generating a new access_token
	newAccessToken := uuid.New().String() + "-" + uuid.New().String()
	// Keeping the access_token in the cache
	if _, err := AccessToken.SetNX(uid, newAccessToken, 24*time.Hour).Result(); err != nil {
		// Failed to keep the uid and access_token as key value pair in the cache
		Logging.ERROR.Println("could not set the access token for the user.")
		return false, ""
	} else {
		// Successfully kept the uid and access_token as key value in the cahce
		Logging.INFO.Println("Successfully set the access token for the user")
		return true, newAccessToken
	}
}
