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

func Escape(source string) string {
	var j int = 0
	if len(source) == 0 {
		return ""
	}
	tempStr := source[:]

	desc := make([]byte, len(tempStr)*2)
	for i := 0; i < len(tempStr); i++ {
		flag := false
		var escape byte
		switch tempStr[i] {
		case '\r':
			flag = true
			escape = '\r'
			break
		case '\n':
			flag = true
			escape = '\n'
			break
		case '\\':
			flag = true
			escape = '\\'
			break
		case '\'':
			flag = true
			escape = '\''
			break
		case '"':
			flag = true
			escape = '"'
			break
		case '\032':
			flag = true
			escape = 'Z'
			break
		default:
		}
		if flag {
			desc[j] = '\\'
			desc[j+1] = escape
			j = j + 2
		} else {
			desc[j] = tempStr[i]
			j = j + 1
		}
	}
	return string(desc[0:j])

}

func TaoChuoi(dsId []int) string {

	inCondition := ""
	for _, item := range dsId {
		if inCondition != "" {
			inCondition += ", "
		}
		if item > 0 {
			inCondition += strconv.Itoa(item)
		}
	}

	return inCondition

}
