package models

import (
	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	Name string `json:"name"`
	// StandardClaims Implit Claims Interface(Valid() func)
	jwt.StandardClaims
}

type CreateTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
