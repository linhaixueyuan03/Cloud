// Code generated by goctl. DO NOT EDIT.
package types

type LoginDetailReply struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LoginDetailRequest struct {
	Idetity string `json:"idetity"`
}

type LoginReply struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
