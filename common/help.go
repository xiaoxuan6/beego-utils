package common

import (
	"regexp"
)

var emailPattern = regexp.MustCompile("[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[a-zA-Z0-9](?:[\\w-]*[\\w])?")

func IsEmail(b []byte) bool {
	return emailPattern.Match(b)
}

func IFTHEN(expr bool, a, b interface{}) interface{} {
	if expr {
		return a
	}

	return b
}

/*************** Map ******************/

func MapOperatorIf(data map[string]interface{}, key string, a interface{}) interface{} {
	if ok := MapExists(data, key); !ok {
		return a
	}
	return data[key]
}

func MapExists(data map[string]interface{}, key string) bool {
	if err := MapVal(data, key); err != nil {
		return true
	}

	return false
}

func MapVal(data map[string]interface{}, key string) interface{} {
	if _, ok := data[key]; ok {
		return data[key]
	}

	return nil
}
