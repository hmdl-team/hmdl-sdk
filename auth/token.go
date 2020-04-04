package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"hmdl-user-service/helper"
	"hmdl-user-service/models/data_user"
	"os"
	"strconv"
	"time"
)

type JwtClaims struct {
	UserId string `json:"user_id"`
	Role   string `json:"role"`
	Name   string `json:"name"`
	jwt.StandardClaims
}

func (c JwtClaims) Valid() error {
	if err := c.StandardClaims.Valid(); err != nil {
		return err
	}

	if c.UserId == "" {
		return errors.New("Must provide a user ID")
	}

	return nil
}

func GenToken(user data_user.DM_TaiKhoan) (string, *time.Time, error) {
	var jwtKey = []byte(os.Getenv("JWK_KEY"))
	expirationTime := time.Now().Add(1 * time.Hour)

	//Định nghĩa
	claims := &JwtClaims{
		UserId: strconv.Itoa(user.DM_TaiKhoanId),
		Role:   helper.IntToString(user.DM_PhanQuyenID),
		Name:   user.TenTaiKhoan,
		StandardClaims: jwt.StandardClaims{
			Id:        "main_user_id",
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// sinh token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	stoken, err := token.SignedString(jwtKey)

	if err != nil {
		return "", nil, err
	}

	return stoken, &expirationTime, err
}

func GenTokenWithTime(user data_user.DM_TaiKhoan, hourNumber time.Duration) (string, *time.Time, error) {
	var jwtKey = []byte(os.Getenv("JWK_KEY"))
	expirationTime := time.Now().Add(hourNumber * time.Hour)

	//Định nghĩa
	claims := &JwtClaims{
		UserId: strconv.Itoa(user.DM_TaiKhoanId),
		Role:   helper.IntToString(user.DM_PhanQuyenID),
		Name:   user.TenTaiKhoan,
		StandardClaims: jwt.StandardClaims{
			Id:        "main_user_id",
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// sinh token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	stoken, err := token.SignedString(jwtKey)

	if err != nil {
		return "", nil, err
	}

	return stoken, &expirationTime, err
}

func DecodeToken(tokenString string) *JwtClaims {
	var jwtKey = []byte(os.Getenv("JWK_KEY"))
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if token != nil {
		if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
			return claims
		} else {
			fmt.Println(err)
		}
	}
	return nil
}
