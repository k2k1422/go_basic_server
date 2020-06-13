package Auth

import (
	"encoding/json"
	"net/http"
	"server/Cache"
	"server/Crypt"
	"server/DataAccess"
	"server/DataModels"
	"server/Logging"
	"server/Response"
	"time"

	"github.com/google/uuid"
)

func login(w http.ResponseWriter, r *http.Request) {

	// Declaring a variable to hold the data coming in the request body
	var loginRequestData DataModels.UserLoginRequest
	// Decoding the request body
	if err := json.NewDecoder(r.Body).Decode(&loginRequestData); err != nil {
		// Failed to decode the request body
		Logging.ERROR.Println("Failed to decode the request body", err)
		Response.BadRequest(w, r, "100")
	} else {
		// Successfully decoded the request body
		if err := DataAccess.Validator.Struct(loginRequestData); err != nil {
			// All the required fields are not present in the request body
			Logging.ERROR.Println("Could not validate the request body. ", err)
			Response.BadRequest(w, r, "101")
		} else {
			// All the required fields are present in the request body
			// Checking a user of the provided email id exists or not
			if DataAccess.UserExistsByEmail(loginRequestData.EmailID) {
				// Getting the data of the user by the email id
				tempUser := DataAccess.GetUserByEmail(loginRequestData.EmailID)
				// Comparing password with the provided  password
				if tempUser.Password == Crypt.HashMD5(loginRequestData.Password) {
					// Provided credentials are correct
					if tempUser.Status {
						// Generating an access_token
						accessToken := uuid.New().String() + "-" + uuid.New().String()
						// Generating an refresh_token
						refreshToken := uuid.New().String() + "-" + uuid.New().String()
						// Storing the access_token and refresh_token in the cache
						// for session management as the value of the key uid
						if Cache.SetAccessToken(tempUser.ID, accessToken) &&
							Cache.SetRefreshToken(tempUser.ID, refreshToken) {
							// Declaring and assigning a variable to hold the response data
							// (uid, access_token, refresh_token, expiry time and the activity access of the user)
							loginResponse := DataModels.AuthToken{
								ID:                     tempUser.ID,
								UserName:               DataAccess.GetUserNameByID([]string{tempUser.ID})[0],
								AccessToken:            accessToken,
								AccessTokenExpiryTime:  time.Now().Add(24 * time.Hour).Unix(),
								RefreshToken:           refreshToken,
								RefreshTokenExpiryTime: time.Now().Add(7 * 24 * time.Hour).Unix(),
							}
							Response.Success(w, r, "306", loginResponse)

						} else {
							// Could not set the access_token or refresh_token in the cache
							Logging.ERROR.Println(r.RequestURI, "Could not create a session for the user-", tempUser.ID)
							Response.InternalServerError(w, r, "303")
						}
					} else {
						// The user id is temporarily deactivated
						Logging.ERROR.Println(r.RequestURI, "User temporarily disabled", loginRequestData.EmailID)
						Response.Unauthorized(w, r, "301")
					}
				} else {
					// User email id and password did not match with the data present in the database
					Logging.ERROR.Println(r.RequestURI, "Unauthorized credential", loginRequestData.EmailID)
					Response.Unauthorized(w, r, "302")
				}
			} else {
				// User does not exist in the database
				Logging.ERROR.Println(r.RequestURI, "User does not exist")
				Response.Unauthorized(w, r, "300")
			}
		}
	}
}

func getAccessToken(w http.ResponseWriter, r *http.Request) {

	// Declaring a variable to hold the data coming in the request body
	var getAccessTokenRequest DataModels.AuthToken
	// Decoding the request data on previously declared variable
	if err := json.NewDecoder(r.Body).Decode(&getAccessTokenRequest); err != nil {
		// Failed to decode the request body into the declared variable
		Logging.ERROR.Println("Could not read the request body. ", err)
		Response.BadRequest(w, r, "100")
	} else {
		//  Validating the request body for required fields
		if err := DataAccess.Validator.Struct(getAccessTokenRequest); err != nil {
			Logging.ERROR.Println("Could not validate the request body. ", err)
			Response.BadRequest(w, r, "101")
		} else {
			// Validating the refresh_token
			if Cache.VerifyRefreshToken(getAccessTokenRequest.ID, getAccessTokenRequest.RefreshToken) {
				// Getting a access_token based on refresh_token and uid
				if status, token := Cache.GetAccessTokenUsingRefreshToken(getAccessTokenRequest.ID); status != true {
					Logging.ERROR.Println("	Could not get the access_token", err)
					Response.InternalServerError(w, r, "308")
				} else {
					// Populating a response body based on the generated access_token and access_token expiry time
					response := struct {
						AccessToken           string `json:"access_token"`
						AccessTokenExpiryTime int64  `json:"access_token_expiry_time"`
					}{
						AccessToken:           token,
						AccessTokenExpiryTime: time.Now().Add(7 * time.Hour).Unix(),
					}
					Response.Success(w, r, "310", response)

				}
			} else {
				// Refresh token is not valid
				Logging.ERROR.Println("	refresh_token is not valid", err)
				Response.Unauthorized(w, r, "307")
			}
		}
	}
}
