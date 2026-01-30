package validators

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"

	"github.com/Shavitjnr/split-chill-ai/pkg/core"
)

type fiscalYearStartContainer struct {
	FiscalYearStart core.FiscalYearStart `validate:"validFiscalYearStart"`
}

func TestValidateFiscalYearStart_ValidValues(t *testing.T) {
	validate := validator.New()
	err := validate.RegisterValidation("validFiscalYearStart", ValidateFiscalYearStart)
	assert.Nil(t, err)

	testCases := []struct {
		name  string
		value core.FiscalYearStart
	}{
		{"January 1", 0x0101},   
		{"December 31", 0x0C1F}, 
		{"July 1", 0x0701},      
		{"April 15", 0x040F},    
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			container := fiscalYearStartContainer{FiscalYearStart: tc.value}
			err := validate.Struct(container)
			assert.Nil(t, err)
		})
	}
}

func TestValidateFiscalYearStart_InvalidValues(t *testing.T) {
	validate := validator.New()
	err := validate.RegisterValidation("validFiscalYearStart", ValidateFiscalYearStart)
	assert.Nil(t, err)

	testCases := []struct {
		name  string
		value core.FiscalYearStart
	}{
		{"Zero value", 0},             
		{"Month 0", 0x0001},           
		{"Month 13", 0x0D01},          
		{"Day 0", 0x0100},             
		{"January 32", 0x0120},        
		{"February 29", 0x021D},       
		{"February 30", 0x021E},       
		{"April 31", 0x041F},          
		{"June 31", 0x061F},           
		{"September 31", 0x091F},      
		{"November 32", 0x0B20},       
		{"Invalid month 255", 0xFF01}, 
		{"Invalid day 255", 0x01FF},   
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			container := fiscalYearStartContainer{FiscalYearStart: tc.value}
			err := validate.Struct(container)
			assert.NotNil(t, err)
		})
	}
}
