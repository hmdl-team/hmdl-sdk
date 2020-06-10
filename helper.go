package sdk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	"sort"
	"strconv"
	"time"
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

//Get total number month in 2 time
func GetMonthFromTwoTime(a, b time.Time) float32 {

	var year, month, day, hour, min, sec int

	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = y2 - y1
	month = int(M2 - M1)
	day = d2 - d1
	hour = h2 - h1
	min = m2 - m1
	sec = s2 - s1

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return float32((year * 12) + month + (day / 31))
}

// Separate objects into several size
func splitObjects(objArr []interface{}, size int) [][]interface{} {
	var chunkSet [][]interface{}
	var chunk []interface{}

	for len(objArr) > size {
		chunk, objArr = objArr[:size], objArr[size:]
		chunkSet = append(chunkSet, chunk)
	}
	if len(objArr) > 0 {
		chunkSet = append(chunkSet, objArr[:])
	}

	return chunkSet
}

// Enable map keys to be retrieved in same order when iterating
func sortedKeys(val map[string]interface{}) []string {
	var keys []string
	for key := range val {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

// Check if string value is contained in slice
func containString(s []string, value string) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}
	return false
}
