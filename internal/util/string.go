package util

import "strings"

func IsEmpty(str string) bool {
	if str == "" || len(str) <= 0 {
		return true
	}
	return false
}

func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

func IsBlank(str string) bool {
	str = strings.Trim(str, " ")
	return IsEmpty(str)
}

func IsNotBlank(str string) bool {
	return !IsBlank(str)
}
