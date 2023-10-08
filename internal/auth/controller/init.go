package controller

type AuthControllers interface {
}

type AuthController struct {
	service AuthControllers
}

func NewAuthController(service AuthControllers) *AuthController {
	return &AuthController{service: service}
}
