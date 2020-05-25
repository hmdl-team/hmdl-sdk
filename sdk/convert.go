package sdk

import (
	"strconv"
	"time"
)

func ToInt(val *int) int {
	return *val
}

func IntToBool(val int) bool {
	if val == 1 {
		return true
	}
	return false
}
func ToTime(val *time.Time) time.Time {
	return *val
}
func ToTimeWhere(val time.Time) string {
	snapshot := "2006-01-02 15:04:05"
	return val.Format(snapshot)
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