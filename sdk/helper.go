package sdk

import (
	"github.com/dgrijalva/jwt-go"
	"strconv"
)

func GetUid(usr interface{}) int {
	user := usr.(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return int(claims["id"].(float64))
}

func StringToInt(value string) int {
	result, _ := strconv.Atoi(value)
	return result
}
