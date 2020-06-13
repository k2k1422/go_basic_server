package DataModels

// Model to hold the data of a login request of a client
type UserLoginRequest struct {
	EmailID  string `json:"email_id" validate:"required"` // Email Id of the user
	Password string `json:"password" validate:"required"` // Password of the user
}

// Model required to hold the data which will be sent to client on successful login
type AuthToken struct {
	ID                     string `json:"uid" validate:"required"` // Email Id of the user
	UserName               string `json:"user_name"`
	AccessToken            string `json:"access_token"` // access token for the session
	AccessTokenExpiryTime  int64  `json:"access_token_expiry_time"`
	RefreshToken           string `json:"refresh_token" validate:"required"`
	RefreshTokenExpiryTime int64  `json:"refresh_token_expiry_time"`
}
