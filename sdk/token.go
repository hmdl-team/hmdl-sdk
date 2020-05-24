package sdk

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"

	"os"
	"time"
)

type JwtClaims struct {
	UserId int `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

type JwNhanVien struct {
	NhanVienId int `json:"nhan_vien_id"`
	jwt.StandardClaims
}

func GenToken(userId int) (string, *time.Time, error) {
	var jwtKey = []byte(os.Getenv("JWK_KEY"))
	expirationTime := time.Now().Add(1 * time.Hour)

	//Định nghĩa
	claims := &JwtClaims{
		UserId: userId,

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

func GenTokenWithTime(userId int, hourNumber time.Duration) (string, *time.Time, error) {
	var jwtKey = []byte(os.Getenv("JWK_KEY"))
	expirationTime := time.Now().Add(hourNumber * time.Hour)

	claims := &JwtClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			Id:        "main_user_id",
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	stoken, err := token.SignedString(jwtKey)

	if err != nil {
		return "", nil, err
	}

	return stoken, &expirationTime, err
}

func GenTokenNhanVienId(nhanVienId int, hourNumber time.Duration) (string, *time.Time, error) {
	var jwtKey = []byte(os.Getenv("JWK_KEY"))
	expirationTime := time.Now().Add(hourNumber * time.Hour)

	//Định nghĩa
	claims := &JwNhanVien{
		NhanVienId: nhanVienId,
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
