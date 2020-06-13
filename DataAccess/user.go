package DataAccess

import (
	"server/DataModels"
	"server/Logging"
)

func UserExistsByEmail(email string) bool {
	/*
		"UserExistsByEmail" will be used to check whether a user exists in the database or not
	*/
	return !Connection.Table("sunpharma.ADL_USER_DATA").Where("email_id = ?", email).Find(&DataModels.UserDetails{}).RecordNotFound()
}

func GetUserByEmail(email string) DataModels.UserDetails {
	/*
		This will be used to get the user details by user email id
	*/
	var queryResult DataModels.UserDetails
	if err := Connection.Table("sunpharma.ADL_USER_DATA").First(&queryResult, "email_id = ?", email).Error; err != nil {
		Logging.ERROR.Println("Error while fetching user data based on the email id")
	}
	return queryResult
}
func RegisterUser(obj DataModels.UserDetails) bool {
	if err := Connection.Table("sunpharma.ADL_USER_DATA").Create(&obj).Error; err != nil {
		Logging.ERROR.Println("Failed to create the user")
		return false
	} else {
		Logging.INFO.Println("Successfully created the user")
		return true
	}
}

func GetUserNameByID(uid []string) []string {
	/*
		"GetAllUserByID" will be used get user details of all the users having the provided id as UID
	*/
	var queryResult []DataModels.UserDetails
	if err := Connection.Table("sunpharma.ADL_USER_DATA").Where("id IN (?)", uid).Find(&queryResult).Error; err != nil {
		Logging.ERROR.Println("Error while fetching user data based on uid")
	}
	var userName []string
	for _, queryResultInstance := range queryResult {
		userName = append(userName, queryResultInstance.FirstName+" "+queryResultInstance.LastName)
	}
	return userName
}

func GetUserList() []DataModels.UserDetails {
	var queryResult []DataModels.UserDetails
	if err := Connection.Table("sunpharma.ADL_USER_DATA").Select("id, emp_id, first_name, last_name, email_id").Find(&queryResult).Error; err != nil {
		Logging.ERROR.Println("Error while fetching user data based on uid")
	}
	return queryResult
}

func GetUserEmailByID(id []string) []string {
	var queryResult []DataModels.UserDetails
	var userEmail []string
	if err := Connection.Table("sunpharma.ADL_USER_DATA").Where("id IN (?)", id).Find(&queryResult).Error; err != nil {
		Logging.ERROR.Println("Error while fetching user data based on uid")
	}
	for _, user := range queryResult {
		userEmail = append(userEmail, user.EmailID)
	}
	return userEmail
}
