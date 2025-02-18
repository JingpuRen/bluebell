package models

// tip : 这里都是和请求参数相关的结构

// ParamSignUp 注册请求的参数结构体
type ParamSignUp struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}

type ParamSignIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
