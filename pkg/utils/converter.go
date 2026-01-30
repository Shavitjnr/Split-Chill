package utils

import (
	"strconv"
	"strings"

	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
)


func IntToString(num int) string {
	return strconv.Itoa(num)
}


func StringToInt(str string) (int, error) {
	return strconv.Atoi(str)
}


func StringTryToInt(str string, defaultValue int) int {
	num, err := StringToInt(str)

	if err != nil {
		return defaultValue
	}

	return num
}


func StringToInt32(str string) (int32, error) {
	val, err := strconv.ParseInt(str, 10, 32)

	if err != nil {
		return 0, err
	}

	return int32(val), nil
}


func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}


func Int64ArrayToStringArray(num []int64) []string {
	ret := make([]string, 0, len(num))

	for i := 0; i < len(num); i++ {
		ret = append(ret, Int64ToString(num[i]))
	}

	return ret
}


func StringToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}


func StringArrayToInt64Array(strs []string) ([]int64, error) {
	ret := make([]int64, 0, len(strs))

	for i := 0; i < len(strs); i++ {
		val, err := StringToInt64(strs[i])

		if err != nil {
			return nil, err
		}

		ret = append(ret, val)
	}

	return ret, nil
}


func StringTryToInt64(str string, defaultValue int64) int64 {
	num, err := StringToInt64(str)

	if err != nil {
		return defaultValue
	}

	return num
}


func Float64ToString(num float64) string {
	return strconv.FormatFloat(num, 'f', -1, 64)
}


func StringToFloat64(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}


func FormatAmount(value int64) string {
	displayAmount := Int64ToString(value)
	negative := displayAmount[0] == '-'

	if negative {
		displayAmount = displayAmount[1:]
	}

	integer := SubString(displayAmount, 0, len(displayAmount)-2)
	decimals := SubString(displayAmount, -2, 2)

	if integer == "" {
		integer = "0"
	}

	if len(decimals) == 0 {
		decimals = "00"
	} else if len(decimals) == 1 {
		decimals = "0" + decimals
	}

	if negative {
		return "-" + integer + "." + decimals
	}

	return integer + "." + decimals
}


func ParseAmount(amount string) (int64, error) {
	if len(amount) < 1 {
		return 0, nil
	}

	sign := int64(1)

	if amount[0] == '-' {
		amount = amount[1:]
		sign = -1
	} else if amount[0] == '+' {
		amount = amount[1:]
		sign = 1
	}

	if len(amount) < 1 {
		return 0, errs.ErrNumberInvalid
	}

	items := strings.Split(amount, ".")

	if len(items) > 2 {
		return 0, errs.ErrNumberInvalid
	}

	var err error
	integer := int64(0)
	decimals := int64(0)

	if len(items[0]) > 0 {
		integer, err = StringToInt64(items[0])

		if err != nil {
			return 0, err
		}

		if integer < 0 {
			return 0, errs.ErrNumberInvalid
		}
	}

	if len(items) == 2 {
		if len(items[1]) > 2 {
			return 0, errs.ErrNumberInvalid
		}

		decimals, err = StringToInt64(items[1])

		if err != nil {
			return 0, err
		}

		if decimals < 0 {
			return 0, errs.ErrNumberInvalid
		}

		if len(items[1]) == 1 {
			decimals = decimals * 10
		}
	}

	return sign*integer*100 + sign*decimals, nil
}
