package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"hmdl-user-service/auth"
	"os"
)

func SetJwtMiddlewares(g *echo.Group) {
	jwtKey := os.Getenv("JWK_KEY")

	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: jwt.SigningMethodHS256.Name,
		ContextKey:    "user",
		Claims:        &auth.JwtClaims{},
		SigningKey:    []byte(jwtKey),
	}))
}
