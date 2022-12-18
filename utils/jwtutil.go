package utils

import (
	"casbinginxorm/models"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type Jwt struct {
}

var secret_key = []byte(os.Getenv("SECRET_KEY"))
var issuer = string(os.Getenv("issuer"))
var audience = string(os.Getenv("audience"))

func (j Jwt) CreateToken(user models.CreateTokenRequest) (models.CreateTokenResponse, error) {
	var err error
	claims := models.CustomClaims{
		Name: user.Username,
		StandardClaims: jwt.StandardClaims{
			Id:        user.Username,
			Subject:   user.Username,
			Issuer:    issuer,
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Audience:  audience,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwt := models.CreateTokenResponse{}

	jwt.AccessToken, err = token.SignedString(secret_key)
	if err != nil {
		return jwt, err
	}

	return j.createRefreshToken(jwt)
}

func (Jwt) ValidateToken(accessToken string) (bool, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret_key, nil
	})

	if err == nil && token.Valid {
		return true, nil
	}

	return false, err
}

func (Jwt) ValidateRefreshToken(model models.CreateTokenResponse) (bool, error) {
	sha1 := sha1.New()
	io.WriteString(sha1, os.Getenv("SECRET_KEY"))

	salt := string(sha1.Sum(nil))[0:16]
	block, err := aes.NewCipher([]byte(salt))
	if err != nil {
		return false, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return false, err
	}

	data, err := base64.URLEncoding.DecodeString(model.RefreshToken)
	if err != nil {
		return false, err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	plain, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return false, err
	}

	if string(plain) != model.AccessToken {
		return false, errors.New("invalid token")
	}

	claims := jwt.MapClaims{}
	parser := jwt.Parser{}
	token, _, err := parser.ParseUnverified(model.AccessToken, claims)

	if err == nil && token.Valid {
		return true, nil
	}

	return false, errors.New("invalid token")
}

func (Jwt) createRefreshToken(token models.CreateTokenResponse) (models.CreateTokenResponse, error) {
	sha1 := sha1.New()
	io.WriteString(sha1, os.Getenv("SECRET_KEY"))

	salt := string(sha1.Sum(nil))[0:16]
	block, err := aes.NewCipher([]byte(salt))
	if err != nil {
		fmt.Println(err.Error())

		return token, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return token, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return token, err
	}

	token.RefreshToken = base64.URLEncoding.EncodeToString(gcm.Seal(nonce, nonce, []byte(token.AccessToken), nil))

	return token, nil
}

func (Jwt) ParserToken(tokenString string) (*models.CustomClaims, error) {
	// https://gowalker.org/github.com/dgrijalva/jwt-go#ParseWithClaims
	// 输入用户自定义的Claims结构体对象,token,以及自定义函数来解析token字符串为jwt的Token结构体指针
	// Keyfunc是匿名函数类型: type Keyfunc func(*Token) (interface{}, error)
	// func ParseWithClaims(tokenString string, claims Claims, keyFunc Keyfunc) (*Token, error) {}
	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret_key, nil
	})

	if token != nil {
		if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
			return claims, nil
		}
	}

	return nil, err
}
