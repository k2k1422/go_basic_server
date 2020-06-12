package DataModels

type Response struct {
	ResponseCode    string      `json:"response_code"`
	ResponseMessage string      `json:"response_message"`
	Data            interface{} `json:"data"`
}
