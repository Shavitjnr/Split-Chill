package core

import (
	"fmt"

	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
)


type FiscalYearStart uint16


const (
	FISCAL_YEAR_START_DEFAULT FiscalYearStart = 0x0101 
	FISCAL_YEAR_START_MIN     FiscalYearStart = 0x0101 
	FISCAL_YEAR_START_MAX     FiscalYearStart = 0x0C1F 
	FISCAL_YEAR_START_INVALID FiscalYearStart = 0xFFFF 
)

var MONTH_MAX_DAYS = []uint8{
	uint8(31), 
	uint8(28), 
	uint8(31), 
	uint8(30), 
	uint8(31), 
	uint8(30), 
	uint8(31), 
	uint8(31), 
	uint8(30), 
	uint8(31), 
	uint8(30), 
	uint8(31), 
}


func NewFiscalYearStart(month uint8, day uint8) (FiscalYearStart, error) {
	if !isValidFiscalYearMonthDay(month, day) {
		return 0, errs.ErrFormatInvalid
	}

	return FiscalYearStart(uint16(month)<<8 | uint16(day)), nil
}


func (f FiscalYearStart) GetMonthDay() (uint8, uint8, error) {
	if f < FISCAL_YEAR_START_MIN || f > FISCAL_YEAR_START_MAX {
		return 0, 0, errs.ErrFormatInvalid
	}

	
	month := uint8(f >> 8)
	day := uint8(f & 0xFF)

	if !isValidFiscalYearMonthDay(month, day) {
		return 0, 0, errs.ErrFormatInvalid
	}

	return month, day, nil
}


func (f FiscalYearStart) String() string {
	month, day, err := f.GetMonthDay()

	if err != nil {
		return "Invalid"
	}

	return fmt.Sprintf("%02d-%02d", month, day)
}


func isValidFiscalYearMonthDay(month uint8, day uint8) bool {
	return uint8(1) <= month && month <= uint8(12) && uint8(1) <= day && day <= MONTH_MAX_DAYS[int(month)-1]
}


type FiscalYearFormat uint8


const (
	FISCAL_YEAR_FORMAT_DEFAULT           FiscalYearFormat = 0
	FISCAL_YEAR_FORMAT_STARTYYYY_ENDYYYY FiscalYearFormat = 1
	FISCAL_YEAR_FORMAT_STARTYYYY_ENDYY   FiscalYearFormat = 2
	FISCAL_YEAR_FORMAT_STARTYY_ENDYY     FiscalYearFormat = 3
	FISCAL_YEAR_FORMAT_ENDYYYY           FiscalYearFormat = 4
	FISCAL_YEAR_FORMAT_ENDYY             FiscalYearFormat = 5
	FISCAL_YEAR_FORMAT_INVALID           FiscalYearFormat = 255 
)


func (f FiscalYearFormat) String() string {
	switch f {
	case FISCAL_YEAR_FORMAT_DEFAULT:
		return "Default"
	case FISCAL_YEAR_FORMAT_STARTYYYY_ENDYYYY:
		return "StartYYYY-EndYYYY"
	case FISCAL_YEAR_FORMAT_STARTYYYY_ENDYY:
		return "StartYYYY-EndYY"
	case FISCAL_YEAR_FORMAT_STARTYY_ENDYY:
		return "StartYY-EndYY"
	case FISCAL_YEAR_FORMAT_ENDYYYY:
		return "EndYYYY"
	case FISCAL_YEAR_FORMAT_ENDYY:
		return "EndYY"
	case FISCAL_YEAR_FORMAT_INVALID:
		return "Invalid"
	default:
		return fmt.Sprintf("Invalid(%d)", int(f))
	}
}
