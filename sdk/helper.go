package sdk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
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

func MaHoa(data string) string {

	secret := "p@ssw0rdhmdl"

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))

	// Write Data to it
	h.Write([]byte(data))

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}