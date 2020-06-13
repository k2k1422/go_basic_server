package User

import (
	"encoding/json"
	"net/http"
	"server/Crypt"
	"server/DataAccess"
	"server/DataModels"
	"server/Logging"
	"server/Response"

	"github.com/google/uuid"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	/*
		"createUser" will be used to  create a new user
	*/

	// Declaring a variable to hold the data present in the request body
	var registrationRequestData DataModels.UserDetails
	// Decoding the data present in the request body to the pre declared struct
	if err := json.NewDecoder(r.Body).Decode(&registrationRequestData); err != nil {
		Logging.ERROR.Println("Could not read the request body. ", err)
		Response.BadRequest(w, r, "100")
	} else {
		// Checking the presence of all the required fields for user creation
		if err := DataAccess.Validator.Struct(registrationRequestData); err != nil {
			// Failed to validate the request body
			Logging.ERROR.Println("Required data fields in the request body is missing ", err)
			Response.BadRequest(w, r, "101")
		} else {
			// Successfully validated the request body
			// Checking whether the user in the request is already present in the db
			if DataAccess.UserExistsByEmail(registrationRequestData.EmailID) {
				// The requested email id already exists in the database
				Logging.INFO.Println("User already exists with same email id.")
				Response.Conflict(w, r, "200")
			} else {
				// The requested user email id does not exist in the database
				// Generating a unique id for the new user
				registrationRequestData.ID = uuid.New().String()
				// Hashing the password of the user
				registrationRequestData.Password = Crypt.HashMD5(registrationRequestData.Password)
				// Register the new user
				if DataAccess.RegisterUser(registrationRequestData) {
					// Successfully registered the user for the application
					Logging.INFO.Println("Successfully created the user")
					Response.Created(w, r, "201")
				} else {
					// Failed to register the user
					Logging.ERROR.Println("Failed to create the user")
					Response.InternalServerError(w, r, "202")
				}
			}
		}
	}
}

func getUserList(w http.ResponseWriter, r *http.Request) {
	Response.Success(w, r, "204", DataAccess.GetUserList())

}
