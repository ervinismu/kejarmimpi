package models

//Response is for send template response
type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Token   string `json:"token"`
}
