package auth

import (
	"errors"
	"fmt"
	"graphyy/internal"
	"graphyy/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func generateJWT(user model.User) model.AuthToken {
	secret := internal.AuthSecret
	expiresAt := time.Now().Add(time.Minute * 15).Unix()

	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = &model.AuthTokenClaim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		User: user,
	}

	tokenString, error := token.SignedString([]byte(secret))
	if error != nil {
		fmt.Println(error)
	}
	return model.AuthToken{
		Token:     tokenString,
		TokenType: "Bearer",
		ExpiresIn: expiresAt,
	}
}

func VerifyJWT(tknStr string) (model.User, error) {

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(internal.AuthSecret), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return model.User{}, errors.New("signature invalid")
		}
		return model.User{}, errors.New("could not parse the auth token")
	}
	if !token.Valid {
		return model.User{}, errors.New("invalid token")
	}

	var username string
	if keyExists(claims, "username") {
		username = claims["username"].(string)
	}

	var createdAt time.Time
	if keyExists(claims, "createdAt") {
		createdAt, err = time.Parse(time.RFC3339, claims["createdAt"].(string))
		if err != nil {
			fmt.Println(err)
		}
	}
	return model.User{Username: username, CreatedAt: createdAt}, nil
}

func keyExists(dict map[string]interface{}, key string) bool {
	val, ok := dict[key]
	return ok && val != nil
}
