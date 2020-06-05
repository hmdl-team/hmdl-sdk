package sdk

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func SetJwtMiddlewares(g *echo.Group) {
	jwtKey := os.Getenv("JWK_KEY")

	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: jwt.SigningMethodHS256.Name,
		ContextKey:    "user",
		Claims:        &JwtClaims{},
		SigningKey:    []byte(jwtKey),
	}))
}


func SetJwtNhanVienId(g *echo.Group) {
	jwtKey := os.Getenv("JWK_KEY")

	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: jwt.SigningMethodHS256.Name,
		ContextKey:    "user",
		Claims:        &JwNhanVien{},
		SigningKey:    []byte(jwtKey),
	}))
}