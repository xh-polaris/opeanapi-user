package controller

type IAuthController interface {
	// 实现idl中定义的所有rpc方法
}

type AuthController struct {
	// 引入domain层的服务
}

func NewAuthController() *AuthController {
	// 需要注入依赖
	return &AuthController{}
}
