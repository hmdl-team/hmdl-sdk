package sdk

import (
	"strconv"
	"time"
)

func ToInt(val *int) int {
	return *val
}

func IntToString(val *int) string {
	return strconv.Itoa(*val)
}

func ToIntPointer(val int) *int {
	return &val
}

func CheckIntPar(par string) (int, error) {
	valParentId, err := strconv.Atoi(par)

	if err != nil || valParentId == 0 {
		return 0, err
	}

	return valParentId, nil
}

func BoolToInt(par bool) int {
	if par {
		return 1
	}
	return 0
}

func  TimeToDateString(ngay time.Time) string {
	return ngay.Format("2006-01-02")
}