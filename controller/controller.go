package controller

import (
	"graphyy/controller/auth"
	"graphyy/service"
)

// Controllers contains all the controllers
type Controllers struct {
	authController *auth.Controller
}

// InitControllers returns a new Controllers
func InitControllers(services *service.Services) *Controllers {
	return &Controllers{
		authController: auth.InitController(services.AuthService),
	}
}
