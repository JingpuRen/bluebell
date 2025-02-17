package models

// ParamSignUp 注册请求的参数结构体
type ParamSignUp struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}
