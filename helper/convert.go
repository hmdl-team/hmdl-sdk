package helper

import (
	"errors"
	"strconv"
)

func ToInt(val *int) int {
	return *val
}

func CheckIntPar(par string) (int, error) {
	valParentId, err := strconv.Atoi(par)

	if err != nil || valParentId == 0 {
		return 0, errors.New("")
	}

	return valParentId, nil
}
