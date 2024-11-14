package utils

import (
	"strconv"
)

func IsNumber(a string) (int, bool) {
	max, err := strconv.Atoi(a);
	if err != nil {
		return max, false
	}
	return max,true
}