package core

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
)

func TestNewFiscalYearStart_ValidMonthDay(t *testing.T) {
	testCases := []struct {
		month    uint8
		day      uint8
		expected FiscalYearStart
	}{
		{1, 1, 0x0101},   
		{4, 15, 0x040F},  
		{7, 1, 0x0701},   
		{12, 31, 0x0C1F}, 
	}

	for _, tc := range testCases {
		fiscal, err := NewFiscalYearStart(tc.month, tc.day)
		assert.Nil(t, err)
		assert.Equal(t, tc.expected, fiscal)
	}
}

func TestNewFiscalYearStart_InvalidMonthDay(t *testing.T) {
	testCases := []struct {
		month uint8
		day   uint8
	}{
		{0, 1},    
		{13, 1},   
		{1, 0},    
		{1, 32},   
		{2, 30},   
		{2, 29},   
		{4, 31},   
		{6, 31},   
		{9, 31},   
		{11, 32},  
		{255, 15}, 
		{5, 255},  
	}

	for _, tc := range testCases {
		fiscal, err := NewFiscalYearStart(tc.month, tc.day)
		assert.Equal(t, FiscalYearStart(0), fiscal)
		assert.Equal(t, errs.ErrFormatInvalid, err)
	}
}

func TestGetMonthDay_ValidFiscalYearStart(t *testing.T) {
	testCases := []struct {
		fiscalYear FiscalYearStart
		month      uint8
		day        uint8
	}{
		{0x0101, 1, 1},   
		{0x0C1F, 12, 31}, 
		{0x0701, 7, 1},   
		{0x040F, 4, 15},  
	}

	for _, tc := range testCases {
		month, day, err := tc.fiscalYear.GetMonthDay()
		assert.Nil(t, err)
		assert.Equal(t, tc.month, month)
		assert.Equal(t, tc.day, day)
	}
}

func TestGetMonthDay_InvalidFiscalYearStart(t *testing.T) {
	testCases := []struct {
		fiscalYear FiscalYearStart
	}{
		{0x0000}, 
		{0x0D01}, 
		{0x0100}, 
		{0x0120}, 
		{0x021D}, 
		{0x021E}, 
		{0x041F}, 
		{0x061F}, 
		{0x091F}, 
		{0x0B20}, 
		{0xFF01}, 
		{0x01FF}, 
		{0},      
	}

	for _, tc := range testCases {
		month, day, err := tc.fiscalYear.GetMonthDay()
		assert.Equal(t, uint8(0), month)
		assert.Equal(t, uint8(0), day)
		assert.Equal(t, errs.ErrFormatInvalid, err)
	}
}

func TestFiscalYearStart_String(t *testing.T) {
	testCases := []struct {
		fiscalYear FiscalYearStart
		expected   string
	}{
		{0x0101, "01-01"},   
		{0x0C1F, "12-31"},   
		{0x0701, "07-01"},   
		{0x040F, "04-15"},   
		{0x021D, "Invalid"}, 
		{0x0000, "Invalid"}, 
		{0x0D01, "Invalid"}, 
		{0x0120, "Invalid"}, 
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, tc.fiscalYear.String())
	}
}

func TestFiscalYearStartConstants(t *testing.T) {
	assert.Equal(t, FiscalYearStart(0xFFFF), FISCAL_YEAR_START_INVALID)
	assert.Equal(t, FiscalYearStart(0x0101), FISCAL_YEAR_START_DEFAULT)
	assert.Equal(t, FiscalYearStart(0x0101), FISCAL_YEAR_START_MIN)
	assert.Equal(t, FiscalYearStart(0x0C1F), FISCAL_YEAR_START_MAX)
}
