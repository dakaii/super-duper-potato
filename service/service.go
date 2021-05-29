package service

import "graphyy/storage"

// Services contains all the repo structs
type Services struct {
	*AuthService
}

// InitServices should be called in main.go
func InitServices(db storage.Storage) *Services {
	authService := NewAuthService(db)
	return &Services{AuthService: authService}
}
