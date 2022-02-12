package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id, omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt, omitempty"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt, omitempty"`
	DeletedAt time.Time          `json:"deletedAt" bson:"deletedAt, omitempty"`
	Username  string             `json:"username" bson:"username, omitempty"`
	Password  string             `json:"password" bson:"password, omitempty"`
}

//AuthToken struct
type AuthToken struct {
	TokenType string `json:"tokenType"`
	Token     string `json:"accessToken"`
	ExpiresIn int64  `json:"expiresIn"`
}

//AuthTokenClaim struct
type AuthTokenClaim struct {
	jwt.StandardClaims
	User
}
