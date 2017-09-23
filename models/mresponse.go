package models

//Response is for send template response
type Response struct {
	Code    string `form:"code" json:"code"`
	Message string `form:"message" json:"message"`
	Token   string `form:"token" json:"token"`
	Status  bool   `form:"status" json:"status"`
}
