package sdk

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenConfig interface {
	SecretKey() []byte
	SecretKeyString() string
	AccessTokenLifeTime() time.Duration
	RefreshTokenLifeTime() time.Duration
}

type User interface {
	GetId() string
	GetDisplayName() string
	GetRole() string
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func GenTokenPair(user User, cf TokenConfig) (*TokenResponse, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.GetId()
	claims["name"] = user.GetDisplayName()
	claims["role"] = user.GetRole()
	claims["exp"] = time.Now().Add(cf.AccessTokenLifeTime()).Unix()

	// Generate encoded token and send it as response.
	tokenString, err := token.SignedString(cf.SecretKey())
	if err != nil {
		return nil, err
	}

	// gen refresh_token
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rClaims := refreshToken.Claims.(jwt.MapClaims)
	rClaims["id"] = user.GetId()
	rClaims["exp"] = time.Now().Add(cf.RefreshTokenLifeTime()).Unix()

	refreshTokenString, err := refreshToken.SignedString(cf.SecretKey())
	if err != nil {
		return nil, err
	}

	res := TokenResponse{
		AccessToken:  tokenString,
		RefreshToken: refreshTokenString,
	}

	return &res, nil
}

func ParseToken(t string, cf TokenConfig) (*jwt.Token, error) {
	return jwt.Parse(t, func(jwtToken *jwt.Token) (i interface{}, err error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", jwtToken.Header["alg"])
		}

		return cf.SecretKey(), nil
	})

}
