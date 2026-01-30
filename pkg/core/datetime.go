package core

import "fmt"


type WeekDay byte


const (
	WEEKDAY_SUNDAY    WeekDay = 0
	WEEKDAY_MONDAY    WeekDay = 1
	WEEKDAY_TUESDAY   WeekDay = 2
	WEEKDAY_WEDNESDAY WeekDay = 3
	WEEKDAY_THURSDAY  WeekDay = 4
	WEEKDAY_FRIDAY    WeekDay = 5
	WEEKDAY_SATURDAY  WeekDay = 6
	WEEKDAY_INVALID   WeekDay = 255
)


func (d WeekDay) String() string {
	switch d {
	case WEEKDAY_SUNDAY:
		return "Sunday"
	case WEEKDAY_MONDAY:
		return "Monday"
	case WEEKDAY_TUESDAY:
		return "Tuesday"
	case WEEKDAY_WEDNESDAY:
		return "Wednesday"
	case WEEKDAY_THURSDAY:
		return "Thursday"
	case WEEKDAY_FRIDAY:
		return "Friday"
	case WEEKDAY_SATURDAY:
		return "Saturday"
	case WEEKDAY_INVALID:
		return "Invalid"
	default:
		return fmt.Sprintf("Invalid(%d)", int(d))
	}
}


type LongDateFormat byte


const (
	LONG_DATE_FORMAT_DEFAULT  LongDateFormat = 0
	LONG_DATE_FORMAT_YYYY_M_D LongDateFormat = 1
	LONG_DATE_FORMAT_M_D_YYYY LongDateFormat = 2
	LONG_DATE_FORMAT_D_M_YYYY LongDateFormat = 3
	LONG_DATE_FORMAT_INVALID  LongDateFormat = 255
)


func (f LongDateFormat) String() string {
	switch f {
	case LONG_DATE_FORMAT_DEFAULT:
		return "Default"
	case LONG_DATE_FORMAT_YYYY_M_D:
		return "YYYY_MM_D"
	case LONG_DATE_FORMAT_M_D_YYYY:
		return "M_D_YYYY"
	case LONG_DATE_FORMAT_D_M_YYYY:
		return "D_M_YYYY"
	case LONG_DATE_FORMAT_INVALID:
		return "Invalid"
	default:
		return fmt.Sprintf("Invalid(%d)", int(f))
	}
}


type ShortDateFormat byte


const (
	SHORT_DATE_FORMAT_DEFAULT  ShortDateFormat = 0
	SHORT_DATE_FORMAT_YYYY_M_D ShortDateFormat = 1
	SHORT_DATE_FORMAT_M_D_YYYY ShortDateFormat = 2
	SHORT_DATE_FORMAT_D_M_YYYY ShortDateFormat = 3
	SHORT_DATE_FORMAT_INVALID  ShortDateFormat = 255
)


func (f ShortDateFormat) String() string {
	switch f {
	case SHORT_DATE_FORMAT_DEFAULT:
		return "Default"
	case SHORT_DATE_FORMAT_YYYY_M_D:
		return "YYYY_MM_D"
	case SHORT_DATE_FORMAT_M_D_YYYY:
		return "M_D_YYYY"
	case SHORT_DATE_FORMAT_D_M_YYYY:
		return "D_M_YYYY"
	case SHORT_DATE_FORMAT_INVALID:
		return "Invalid"
	default:
		return fmt.Sprintf("Invalid(%d)", int(f))
	}
}


type LongTimeFormat byte


const (
	LONG_TIME_FORMAT_DEFAULT    LongTimeFormat = 0
	LONG_TIME_FORMAT_HH_MM_SS   LongTimeFormat = 1
	LONG_TIME_FORMAT_A_HH_MM_SS LongTimeFormat = 2
	LONG_TIME_FORMAT_HH_MM_SS_A LongTimeFormat = 3
	LONG_TIME_FORMAT_INVALID    LongTimeFormat = 255
)


func (f LongTimeFormat) String() string {
	switch f {
	case LONG_TIME_FORMAT_DEFAULT:
		return "Default"
	case LONG_TIME_FORMAT_HH_MM_SS:
		return "HH_MM_SS"
	case LONG_TIME_FORMAT_A_HH_MM_SS:
		return "A_HH_MM_SS"
	case LONG_TIME_FORMAT_HH_MM_SS_A:
		return "HH_MM_SS_A"
	case LONG_TIME_FORMAT_INVALID:
		return "Invalid"
	default:
		return fmt.Sprintf("Invalid(%d)", int(f))
	}
}


type ShortTimeFormat byte


const (
	SHORT_TIME_FORMAT_DEFAULT ShortTimeFormat = 0
	SHORT_TIME_FORMAT_HH_MM   ShortTimeFormat = 1
	SHORT_TIME_FORMAT_A_HH_MM ShortTimeFormat = 2
	SHORT_TIME_FORMAT_HH_MM_A ShortTimeFormat = 3
	SHORT_TIME_FORMAT_INVALID ShortTimeFormat = 255
)


func (f ShortTimeFormat) String() string {
	switch f {
	case SHORT_TIME_FORMAT_DEFAULT:
		return "Default"
	case SHORT_TIME_FORMAT_HH_MM:
		return "HH_MM"
	case SHORT_TIME_FORMAT_A_HH_MM:
		return "A_HH_MM"
	case SHORT_TIME_FORMAT_HH_MM_A:
		return "HH_MM_A"
	case SHORT_TIME_FORMAT_INVALID:
		return "Invalid"
	default:
		return fmt.Sprintf("Invalid(%d)", int(f))
	}
}
