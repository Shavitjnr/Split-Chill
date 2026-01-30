package exchangerates

import (
	"bytes"
	"encoding/xml"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"

	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
	"github.com/Shavitjnr/split-chill-ai/pkg/log"
	"github.com/Shavitjnr/split-chill-ai/pkg/models"
	"github.com/Shavitjnr/split-chill-ai/pkg/utils"
	"github.com/Shavitjnr/split-chill-ai/pkg/validators"
)

const euroCentralBankExchangeRateUrl = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"
const euroCentralBankExchangeRateReferenceUrl = "https://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/index.en.html"
const euroCentralBankDataSource = "European Central Bank"
const euroCentralBankBaseCurrency = "EUR"

const euroCentralBankDataUpdateDateFormat = "2006-01-02 15"
const euroCentralBankDataUpdateDateTimezone = "Europe/Berlin"


type EuroCentralBankDataSource struct {
	HttpExchangeRatesDataSource
}


type EuroCentralBankExchangeRateData struct {
	XMLName          xml.Name                        `xml:"Envelope"`
	AllExchangeRates []*EuroCentralBankExchangeRates `xml:"Cube>Cube"`
}


type EuroCentralBankExchangeRates struct {
	Date          string                         `xml:"time,attr"`
	ExchangeRates []*EuroCentralBankExchangeRate `xml:"Cube"`
}


type EuroCentralBankExchangeRate struct {
	Currency string `xml:"currency,attr"`
	Rate     string `xml:"rate,attr"`
}


func (e *EuroCentralBankExchangeRateData) ToLatestExchangeRateResponse(c core.Context) *models.LatestExchangeRateResponse {
	if len(e.AllExchangeRates) < 1 {
		log.Errorf(c, "[euro_central_bank_datasource.ToLatestExchangeRateResponse] all exchange rates is empty")
		return nil
	}

	latestEuroCentralBankExchangeRate := e.AllExchangeRates[0]

	if len(latestEuroCentralBankExchangeRate.ExchangeRates) < 1 {
		log.Errorf(c, "[euro_central_bank_datasource.ToLatestExchangeRateResponse] exchange rates is empty")
		return nil
	}

	exchangeRates := make(models.LatestExchangeRateSlice, 0, len(latestEuroCentralBankExchangeRate.ExchangeRates))

	for i := 0; i < len(latestEuroCentralBankExchangeRate.ExchangeRates); i++ {
		exchangeRate := latestEuroCentralBankExchangeRate.ExchangeRates[i]

		if _, exists := validators.AllCurrencyNames[exchangeRate.Currency]; !exists {
			continue
		}

		if _, err := utils.StringToFloat64(exchangeRate.Rate); err != nil {
			continue
		}

		exchangeRates = append(exchangeRates, exchangeRate.ToLatestExchangeRate())
	}

	timezone, err := time.LoadLocation(euroCentralBankDataUpdateDateTimezone)

	if err != nil {
		log.Errorf(c, "[euro_central_bank_datasource.ToLatestExchangeRateResponse] failed to get timezone, timezone name is %s", euroCentralBankDataUpdateDateTimezone)
		return nil
	}

	updateDateTime := latestEuroCentralBankExchangeRate.Date + " 16" 
	updateTime, err := time.ParseInLocation(euroCentralBankDataUpdateDateFormat, updateDateTime, timezone)

	if err != nil {
		log.Errorf(c, "[euro_central_bank_datasource.ToLatestExchangeRateResponse] failed to parse update date, datetime is %s", updateDateTime)
		return nil
	}

	latestExchangeRateResp := &models.LatestExchangeRateResponse{
		DataSource:    euroCentralBankDataSource,
		ReferenceUrl:  euroCentralBankExchangeRateReferenceUrl,
		UpdateTime:    updateTime.Unix(),
		BaseCurrency:  euroCentralBankBaseCurrency,
		ExchangeRates: exchangeRates,
	}

	return latestExchangeRateResp
}


func (e *EuroCentralBankExchangeRate) ToLatestExchangeRate() *models.LatestExchangeRate {
	return &models.LatestExchangeRate{
		Currency: e.Currency,
		Rate:     e.Rate,
	}
}


func (e *EuroCentralBankDataSource) BuildRequests() ([]*http.Request, error) {
	req, err := http.NewRequest("GET", euroCentralBankExchangeRateUrl, nil)

	if err != nil {
		return nil, err
	}

	return []*http.Request{req}, nil
}


func (e *EuroCentralBankDataSource) Parse(c core.Context, content []byte) (*models.LatestExchangeRateResponse, error) {
	xmlDecoder := xml.NewDecoder(bytes.NewReader(content))
	xmlDecoder.CharsetReader = charset.NewReaderLabel

	euroCentralBankData := &EuroCentralBankExchangeRateData{}
	err := xmlDecoder.Decode(euroCentralBankData)

	if err != nil {
		log.Errorf(c, "[euro_central_bank_datasource.Parse] failed to parse xml data, content is %s, because %s", string(content), err.Error())
		return nil, errs.ErrFailedToRequestRemoteApi
	}

	latestExchangeRateResponse := euroCentralBankData.ToLatestExchangeRateResponse(c)

	if latestExchangeRateResponse == nil {
		log.Errorf(c, "[euro_central_bank_datasource.Parse] failed to parse latest exchange rate data, content is %s", string(content))
		return nil, errs.ErrFailedToRequestRemoteApi
	}

	return latestExchangeRateResponse, nil
}
