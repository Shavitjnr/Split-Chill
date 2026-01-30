package utils

import (
	"crypto/rand"
	"math/big"
	"regexp"
	"strings"
)

var (
	numberPattern = regexp.MustCompile("(-?\\d+)(\\.\\d+)?")
)


func IsStringOnlyContainsDigits(str string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] < '0' || str[i] > '9' {
			return false
		}
	}

	return true
}


func GetRandomInteger(max int) (int, error) {
	result, err := rand.Int(rand.Reader, big.NewInt(int64(max)))

	if err != nil {
		return 0, err
	}

	return int(result.Int64()), nil
}


func ParseFirstConsecutiveNumber(str string) (string, bool) {
	result := numberPattern.FindAllString(str, 1)

	if len(result) > 0 {
		return result[0], true
	} else {
		return "", false
	}
}


func TrimTrailingZerosInDecimal(num string) string {
	if len(num) < 1 {
		return num
	}

	dotPosition := strings.Index(num, ".")

	if dotPosition < 0 {
		return num
	}

	lastNonZeroPosition := len(num)

	for i := len(num) - 1; i > dotPosition+1; i-- {
		if num[i] == '0' {
			lastNonZeroPosition = i
		} else {
			break
		}
	}

	if lastNonZeroPosition >= len(num) {
		return num
	}

	return num[0:lastNonZeroPosition]
}
