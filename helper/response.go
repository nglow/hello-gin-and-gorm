package helper

import "strings"

// Response is used for static json return
type Response struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	Errors interface{} `json:"errors"`
	Data interface{} `json:"data"`
}

// EmptyObj is used when data doesn't want to be null on json
type EmptyObj struct{}

// BuildResponse method is to inject data value to dynamic success response
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status: status,
		Message: message,
		Errors: nil,
		Data: data,
	}
	return res
}

// BuildErrorResponse Method is to inject data value to dynamic failed response
func BuildErrorResponse(message string, err string, data interface{}) Response {
	splitError := strings.Split(err, "\n")
	res := Response {
		Status:  false,
		Message: message,
		Errors:  splitError,
		Data:    data,
	}
	return res
}
