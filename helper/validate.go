package helper

import (
	"regexp"
)

func IsValidPhoneNumber(phone string) bool {
	if len(phone) != 10 {
		return false
	}
	valid, _ := regexp.MatchString(`[0-9]`, phone)
	return valid
}
