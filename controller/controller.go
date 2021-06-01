package controller

import (
	"github.com/dakaii/superduperpotato/controller/auth"
	"github.com/dakaii/superduperpotato/service"
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
