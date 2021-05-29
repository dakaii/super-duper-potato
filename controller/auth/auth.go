package auth

import (
	"errors"
	"graphyy/model"
	"graphyy/service"

	"golang.org/x/crypto/bcrypt"
)

// declaring the repository interface in the controller package allows us to easily swap out the actual implementation, enforcing loose coupling.
type authService interface {
	GetExistingUser(username string) model.User
	SaveUser(user model.User) (model.User, error)
}

// Controller contains the service, which contains database-related logic, as an injectable dependency, allowing us to decouple business logic from db logic.
type Controller struct {
	authService
}

// InitController initializes the user controller.
func InitController(authService *service.AuthService) *Controller {
	return &Controller{
		authService,
	}
}

// Signup lets users sign up for this application and returns a jwt.
func (c *Controller) Signup(user model.User) (model.AuthToken, error) {
	if !isValidUsername(user.Username) {
		return model.AuthToken{}, errors.New("invalid username")
	}
	existingUser := c.authService.GetExistingUser(user.Username)
	if existingUser.Username != "" {
		return model.AuthToken{}, errors.New("this username is already in use")
	}
	user, err := c.authService.SaveUser(user)
	if err != nil {
		return model.AuthToken{}, err
	}

	token := generateJWT(user)
	return token, nil
}

// Login returns a jwt.
func (c *Controller) Login(user model.User) (model.AuthToken, error) {
	existingUser := c.authService.GetExistingUser(user.Username)
	if existingUser.Username == "" {
		return model.AuthToken{}, errors.New("no user found with the inputted username")
	}
	isValid := checkPasswordHash(user.Password, existingUser.Password)
	if !isValid {
		return model.AuthToken{}, errors.New("invalid credentials")
	}

	token := generateJWT(user)
	return token, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func isValidUsername(username string) bool {
	return len(username) > 6
}
