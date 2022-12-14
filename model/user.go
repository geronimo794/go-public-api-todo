package model

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username" form:"username" validate:"required,len=6,alphanum"`
	Password string `json:"password" form:"password" validate:"required,len=6,alphanum"`
}
