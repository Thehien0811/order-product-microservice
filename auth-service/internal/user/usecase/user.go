package usecase

import (
	"github.com/dgrijalva/jwt-go"
)

type SignUpInput struct {
	Username string
	Password string
}

type SignInInput struct {
	Username string
	Password string
}
type DetailUser struct {
	ID       string
	Username string
	Password string
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
