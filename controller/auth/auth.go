package auth

import (
	"encoding/json"
	"net/http"

	"github.com/dakaii/superduperpotato/model"
	"github.com/dakaii/superduperpotato/service"

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
func (c *Controller) Signup(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var newUser model.User
	err := decoder.Decode(&newUser)
	if err != nil {
		panic(err)
	}
	if !isValidUsername(newUser.Username) {
		http.Error(w, "invalid username", http.StatusBadRequest)
		return
	}
	existingUser := c.authService.GetExistingUser(newUser.Username)
	if existingUser.Username != "" {
		http.Error(w, "this username is already in use", http.StatusBadRequest)
		return
	}
	createdUser, err := c.authService.SaveUser(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token := generateJWT(createdUser)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}

// Login returns a jwt.
func (c *Controller) Login(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var user model.User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	existingUser := c.authService.GetExistingUser(user.Username)
	if existingUser.Username == "" {
		w.Write([]byte("no user found with the inputted username"))
		return
	}
	isValid := checkPasswordHash(user.Password, existingUser.Password)
	if !isValid {
		w.Write([]byte("Invalid credentials"))
		return
	}

	token := generateJWT(user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func isValidUsername(username string) bool {
	return len(username) > 6
}
