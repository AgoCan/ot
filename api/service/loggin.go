package service

import "ot/pkg/response"

type User struct{}

func (u *User) Login() *response.ResponseError {
	
	return response.NewResponse(200, 200, "ok")
}
