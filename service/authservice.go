package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dakaii/superduperpotato/model"
	"github.com/dakaii/superduperpotato/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	storage storage.Storage
}

// NewAuthService constructs a AuthService
func NewAuthService(db storage.Storage) *AuthService {
	return &AuthService{
		db,
	}
}

// GetExistingUser fetches a user by the username from the db and returns it.
func (service *AuthService) GetExistingUser(username string) model.User {
	filter := bson.M{"username": username}
	var user model.User
	collection := service.storage.Mongo.Collection(storage.Auth)
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		fmt.Println(err.Error())
	}
	return model.User{Username: user.Username, Password: user.Password}
}

// SaveUser creates a new user in the db..
func (auth *AuthService) SaveUser(user model.User) (model.User, error) {
	hashedPass, err := hashPassword(user.Password)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		return model.User{Username: "", Password: ""}, err
	}
	user.ID = primitive.NewObjectID()
	user.Password = hashedPass
	user.CreatedAt = time.Now().UTC()

	fmt.Println("inserting a user with username:", user.Username)
	collection := auth.storage.Mongo.Collection(storage.Auth)
	insertResult, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		return model.User{Username: "", Password: ""}, err
	}
	fmt.Println("Inserted a user with ID:", insertResult.InsertedID)
	return user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
