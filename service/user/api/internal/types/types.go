// Code generated by goctl. DO NOT EDIT.
package types

type RegisterRequest struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Country  string `json:"country"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
}

type LoginRequest struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type Response struct {
	Message string `json:"message"`
}
