package auth

import "github.com/dgrijalva/jwt-go"

type UserClaims struct {
	UserData *UserData `json:"userdata"`
	jwt.StandardClaims
}
