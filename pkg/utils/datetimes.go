package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
)

const (
	longDateFormat                        = "2006-01-02"
	longDateTimeFormat                    = "2006-01-02 15:04:05"
	longDateTimeWithTimezoneFormat        = "2006-01-02 15:04:05Z07:00"
	longDateTimeWithTimezoneFormat2       = "2006-01-02 15:04:05 Z0700"
	longDateTimeWithTimezoneRFC3339Format = "2006-01-02T15:04:05Z07:00"
	longDateTimeWithoutSecondFormat       = "2006-01-02 15:04"
	shortDateTimeFormat                   = "2006-1-2 15:4:5"
	yearMonthDateTimeFormat               = "2006-01"
	westernmostTimezoneUtcOffset          = -720 
	easternmostTimezoneUtcOffset          = 840  
)


func ParseNumericYearMonth(yearMonth string) (int32, int32, error) {
	yearMonthParts := strings.Split(yearMonth, "-")

	if len(yearMonthParts) != 2 {
		return 0, 0, errs.ErrParameterInvalid
	}

	year, err := StringToInt32(yearMonthParts[0])

	if err != nil {
		return 0, 0, err
	}

	month, err := StringToInt32(yearMonthParts[1])

	if err != nil {
		return 0, 0, err
	}

	return year, month, nil
}


func FormatUnixTimeToLongDate(unixTime int64, timezone *time.Location) string {
	t := parseFromUnixTime(unixTime)

	if timezone != nil {
		t = t.In(timezone)
	}

	return t.Format(longDateFormat)
}


func FormatUnixTimeToLongDateTime(unixTime int64, timezone *time.Location) string {
	t := parseFromUnixTime(unixTime)

	if timezone != nil {
		t = t.In(timezone)
	}

	return t.Format(longDateTimeFormat)
}


func FormatUnixTimeToLongDateTimeWithTimezone(unixTime int64, timezone *time.Location) string {
	t := parseFromUnixTime(unixTime)

	if timezone != nil {
		t = t.In(timezone)
	}

	return t.Format(longDateTimeWithTimezoneFormat)
}


func FormatUnixTimeToLongDateTimeWithTimezoneRFC3339Format(unixTime int64, timezone *time.Location) string {
	t := parseFromUnixTime(unixTime)

	if timezone != nil {
		t = t.In(timezone)
	}

	return t.Format(longDateTimeWithTimezoneRFC3339Format)
}

func FormatYearMonthDayToLongDateTime(year string, month string, day string) (string, error) {
	if len(year) == 2 {
		yearLast2Digits, err := StringToInt(year)

		if err != nil {
			return "", err
		}

		currentYear := time.Now().Year()
		currentYearLast2Digits := currentYear % 100

		if yearLast2Digits <= currentYearLast2Digits {
			year = IntToString(currentYear/100) + year
		} else {
			year = IntToString(currentYear/100-1) + year
		}
	}

	if len(month) < 2 {
		month = "0" + month
	}

	if len(day) < 2 {
		day = "0" + day
	}

	return fmt.Sprintf("%s-%s-%s 00:00:00", year, month, day), nil
}


func FormatUnixTimeToLongDateTimeInServerTimezone(unixTime int64) string {
	return parseFromUnixTime(unixTime).Format(longDateTimeFormat)
}


func FormatUnixTimeToLongDateTimeWithoutSecond(unixTime int64, timezone *time.Location) string {
	t := parseFromUnixTime(unixTime)

	if timezone != nil {
		t = t.In(timezone)
	}

	return t.Format(longDateTimeWithoutSecondFormat)
}


func FormatUnixTimeToYearMonth(unixTime int64, timezone *time.Location) string {
	t := parseFromUnixTime(unixTime)

	if timezone != nil {
		t = t.In(timezone)
	}

	return t.Format(yearMonthDateTimeFormat)
}


func FormatUnixTimeToNumericYearMonth(unixTime int64, timezone *time.Location) int32 {
	t := parseFromUnixTime(unixTime)

	if timezone != nil {
		t = t.In(timezone)
	}

	return int32(t.Year())*100 + int32(t.Month())
}


func FormatUnixTimeToNumericYearMonthDay(unixTime int64, timezone *time.Location) int32 {
	t := parseFromUnixTime(unixTime)

	if timezone != nil {
		t = t.In(timezone)
	}

	return int32(t.Year())*10000 + int32(t.Month())*100 + int32(t.Day())
}


func FormatUnixTimeToNumericLocalDateTime(unixTime int64, timezone *time.Location) int64 {
	t := parseFromUnixTime(unixTime)

	if timezone != nil {
		t = t.In(timezone)
	}

	localDateTime := int64(t.Year())
	localDateTime = localDateTime*100 + int64(t.Month())
	localDateTime = localDateTime*100 + int64(t.Day())
	localDateTime = localDateTime*100 + int64(t.Hour())
	localDateTime = localDateTime*100 + int64(t.Minute())
	localDateTime = localDateTime*100 + int64(t.Second())

	return localDateTime
}


func GetMinUnixTimeWithSameLocalDateTime(unixTime int64, currentUtcOffset int16) int64 {
	return unixTime + int64(currentUtcOffset)*60 - easternmostTimezoneUtcOffset*60
}


func GetMaxUnixTimeWithSameLocalDateTime(unixTime int64, currentUtcOffset int16) int64 {
	return unixTime + int64(currentUtcOffset)*60 - westernmostTimezoneUtcOffset*60
}


func ParseFromLongDateFirstTime(t string, utcOffset int16) (time.Time, error) {
	timezone := time.FixedZone("Timezone", int(utcOffset)*60)
	return time.ParseInLocation(longDateFormat, t, timezone)
}


func ParseFromLongDateLastTime(t string, utcOffset int16) (time.Time, error) {
	timezone := time.FixedZone("Timezone", int(utcOffset)*60)
	lastTime, err := time.ParseInLocation(longDateFormat, t, timezone)

	if err != nil {
		return lastTime, err
	}

	return lastTime.Add(24 * time.Hour).Add(-1 * time.Nanosecond), nil
}


func ParseFromLongDateTimeToMinUnixTime(t string) (time.Time, error) {
	timezone := time.FixedZone("Timezone", easternmostTimezoneUtcOffset*60)
	return time.ParseInLocation(longDateTimeFormat, t, timezone)
}


func ParseFromLongDateTimeToMaxUnixTime(t string) (time.Time, error) {
	timezone := time.FixedZone("Timezone", westernmostTimezoneUtcOffset*60)
	return time.ParseInLocation(longDateTimeFormat, t, timezone)
}


func ParseFromLongDateTimeInFixedUtcOffset(t string, utcOffset int16) (time.Time, error) {
	timezone := time.FixedZone("Timezone", int(utcOffset)*60)
	return time.ParseInLocation(longDateTimeFormat, t, timezone)
}


func ParseFromLongDateTimeInTimeZone(t string, timezone *time.Location) (time.Time, error) {
	return time.ParseInLocation(longDateTimeFormat, t, timezone)
}


func ParseFromLongDateTimeWithTimezone(t string) (time.Time, error) {
	return time.Parse(longDateTimeWithTimezoneFormat, t)
}


func ParseFromLongDateTimeWithTimezone2(t string) (time.Time, error) {
	return time.Parse(longDateTimeWithTimezoneFormat2, t)
}


func ParseFromLongDateTimeWithTimezoneRFC3339Format(t string) (time.Time, error) {
	return time.Parse(longDateTimeWithTimezoneRFC3339Format, t)
}


func ParseFromLongDateTimeWithoutSecondInFixedUtcOffset(t string, utcOffset int16) (time.Time, error) {
	timezone := time.FixedZone("Timezone", int(utcOffset)*60)
	return time.ParseInLocation(longDateTimeWithoutSecondFormat, t, timezone)
}


func ParseFromShortDateTimeInFixedUtcOffset(t string, utcOffset int16) (time.Time, error) {
	timezone := time.FixedZone("Timezone", int(utcOffset)*60)
	return time.ParseInLocation(shortDateTimeFormat, t, timezone)
}

func ParseFromElapsedSeconds(elapsedSeconds int) (string, error) {
	if elapsedSeconds < 0 || elapsedSeconds >= 86400 {
		return "", errs.ErrFormatInvalid
	}

	second := elapsedSeconds % 60
	elapsedSeconds = elapsedSeconds - second
	elapsedSeconds = elapsedSeconds / 60

	minute := elapsedSeconds % 60
	elapsedSeconds = elapsedSeconds - minute
	elapsedSeconds = elapsedSeconds / 60

	hour := elapsedSeconds

	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second), nil
}


func IsUnixTimeEqualsYearAndMonth(unixTime int64, timezone *time.Location, year int32, month int32) bool {
	date := parseFromUnixTime(unixTime).In(timezone)
	return date.Year() == int(year) && int(date.Month()) == int(month)
}


func GetTimezoneOffsetMinutes(unixTime int64, timezone *time.Location) int16 {
	_, tzOffset := parseFromUnixTime(unixTime).In(timezone).Zone()
	tzMinuteOffset := int16(tzOffset / 60)

	return tzMinuteOffset
}


func GetServerTimezoneOffsetMinutes() int16 {
	_, tzOffset := time.Now().Zone()
	tzMinuteOffset := int16(tzOffset / 60)

	return tzMinuteOffset
}


func FormatTimezoneOffset(unixTime int64, timezone *time.Location) string {
	tzMinutesOffset := GetTimezoneOffsetMinutes(unixTime, timezone)

	sign := "+"
	hourAbsOffset := tzMinutesOffset / 60
	minuteAbsOffset := tzMinutesOffset % 60

	if hourAbsOffset < 0 {
		sign = "-"
		hourAbsOffset = -hourAbsOffset
		minuteAbsOffset = -minuteAbsOffset
	}

	return fmt.Sprintf("%s%02d:%02d", sign, hourAbsOffset, minuteAbsOffset)
}


func FormatTimezoneOffsetFromHoursOffset(hoursOffset string) (string, error) {
	hoursOffsetValue, err := StringToFloat64(hoursOffset)

	if err != nil {
		return "", errs.ErrFormatInvalid
	}

	tzMinutesOffset := int16(hoursOffsetValue * 60)

	sign := "+"
	hourAbsOffset := tzMinutesOffset / 60
	minuteAbsOffset := tzMinutesOffset % 60

	if hourAbsOffset < 0 {
		sign = "-"
		hourAbsOffset = -hourAbsOffset
		minuteAbsOffset = -minuteAbsOffset
	}

	return fmt.Sprintf("%s%02d:%02d", sign, hourAbsOffset, minuteAbsOffset), nil
}


func ParseFromTimezoneOffset(tzOffset string) (*time.Location, error) {
	if len(tzOffset) != 6 { 
		return nil, errs.ErrFormatInvalid
	}

	sign := tzOffset[0]

	if sign != '-' && sign != '+' {
		return nil, errs.ErrFormatInvalid
	}

	offsets := strings.Split(tzOffset[1:], ":")

	if len(offsets) != 2 {
		return nil, errs.ErrFormatInvalid
	}

	hourAbsOffset, err := StringToInt(offsets[0])

	if err != nil {
		return nil, err
	}

	minuteAbsOffset, err := StringToInt(offsets[1])

	if err != nil {
		return nil, err
	}

	totalMinuteOffset := hourAbsOffset*60 + minuteAbsOffset

	if sign == '-' {
		totalMinuteOffset = -totalMinuteOffset
	}

	totalOffset := totalMinuteOffset * 60
	return time.FixedZone("Timezone", totalOffset), nil
}


func GetMinTransactionTimeFromUnixTime(unixTime int64) int64 {
	return unixTime * 1000
}


func GetMaxTransactionTimeFromUnixTime(unixTime int64) int64 {
	return unixTime*1000 + 999
}


func GetUnixTimeFromTransactionTime(transactionTime int64) int64 {
	return transactionTime / 1000
}


func GetTransactionTimeRangeByYearMonth(year int32, month int32) (int64, int64, error) {
	startMinUnixTime, err := ParseFromLongDateTimeToMinUnixTime(fmt.Sprintf("%d-%02d-01 00:00:00", year, month))
	startMaxUnixTime, err := ParseFromLongDateTimeToMaxUnixTime(fmt.Sprintf("%d-%02d-01 00:00:00", year, month))

	if err != nil {
		return 0, 0, err
	}

	endMaxUnixTime := startMaxUnixTime.AddDate(0, 1, 0)

	minTransactionTime := GetMinTransactionTimeFromUnixTime(startMinUnixTime.Unix())
	maxTransactionTime := GetMinTransactionTimeFromUnixTime(endMaxUnixTime.Unix()) - 1

	return minTransactionTime, maxTransactionTime, nil
}


func GetStartOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}


func parseFromUnixTime(unixTime int64) time.Time {
	return time.Unix(unixTime, 0)
}
