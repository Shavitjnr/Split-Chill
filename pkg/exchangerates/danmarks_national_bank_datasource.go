package exchangerates

import (
	"bytes"
	"encoding/xml"
	"math"
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

const danmarksNationalbankExchangeRateUrl = "https://www.nationalbanken.dk/api/currencyratesxml?lang=en"
const danmarksNationalbankExchangeRateReferenceUrl = "https://www.nationalbanken.dk/en/what-we-do/stable-prices-monetary-policy-and-the-danish-economy/exchange-rates"
const danmarksNationalbankDataSource = "Danmarks Nationalbank"

const danmarksNationalbankDataUpdateDateFormat = "2006-01-02 15"
const danmarksNationalbankDataUpdateDateTimezone = "Europe/Copenhagen"


type DanmarksNationalbankDataSource struct {
	HttpExchangeRatesDataSource
}


type DanmarksNationalbankExchangeRateData struct {
	XMLName            xml.Name                                  `xml:"exchangerates"`
	DailyExchangeRates []*DanmarksNationalbankDailyExchangeRates `xml:"dailyrates"`
	BaseCurrency       string                                    `xml:"refcur,attr"`
}


type DanmarksNationalbankDailyExchangeRates struct {
	Date          string                              `xml:"id,attr"`
	ExchangeRates []*DanmarksNationalbankExchangeRate `xml:"currency"`
}


type DanmarksNationalbankExchangeRate struct {
	Currency string `xml:"code,attr"`
	Rate     string `xml:"rate,attr"`
}


func (e *DanmarksNationalbankExchangeRateData) ToLatestExchangeRateResponse(c core.Context) *models.LatestExchangeRateResponse {
	if len(e.DailyExchangeRates) < 1 {
		log.Errorf(c, "[danmarks_national_bank_datasource.ToLatestExchangeRateResponse] daily exchange rates is empty")
		return nil
	}

	latestDanmarksNationalbankExchangeRate := e.DailyExchangeRates[0]

	if len(latestDanmarksNationalbankExchangeRate.ExchangeRates) < 1 {
		log.Errorf(c, "[danmarks_national_bank_datasource.ToLatestExchangeRateResponse] exchange rates is empty")
		return nil
	}

	exchangeRates := make(models.LatestExchangeRateSlice, 0, len(latestDanmarksNationalbankExchangeRate.ExchangeRates))

	for i := 0; i < len(latestDanmarksNationalbankExchangeRate.ExchangeRates); i++ {
		exchangeRate := latestDanmarksNationalbankExchangeRate.ExchangeRates[i]

		if _, exists := validators.AllCurrencyNames[exchangeRate.Currency]; !exists {
			continue
		}

		finalExchangeRate := exchangeRate.ToLatestExchangeRate(c)

		if finalExchangeRate == nil {
			continue
		}

		exchangeRates = append(exchangeRates, finalExchangeRate)
	}

	timezone, err := time.LoadLocation(danmarksNationalbankDataUpdateDateTimezone)

	if err != nil {
		log.Errorf(c, "[danmarks_national_bank_datasource.ToLatestExchangeRateResponse] failed to get timezone, timezone name is %s", danmarksNationalbankDataUpdateDateTimezone)
		return nil
	}

	updateDateTime := latestDanmarksNationalbankExchangeRate.Date + " 16" 
	updateTime, err := time.ParseInLocation(danmarksNationalbankDataUpdateDateFormat, updateDateTime, timezone)

	if err != nil {
		log.Errorf(c, "[danmarks_national_bank_datasource.ToLatestExchangeRateResponse] failed to parse update date, datetime is %s", updateDateTime)
		return nil
	}

	latestExchangeRateResp := &models.LatestExchangeRateResponse{
		DataSource:    danmarksNationalbankDataSource,
		ReferenceUrl:  danmarksNationalbankExchangeRateReferenceUrl,
		UpdateTime:    updateTime.Unix(),
		BaseCurrency:  e.BaseCurrency,
		ExchangeRates: exchangeRates,
	}

	return latestExchangeRateResp
}


func (e *DanmarksNationalbankExchangeRate) ToLatestExchangeRate(c core.Context) *models.LatestExchangeRate {
	rate, err := utils.StringToFloat64(e.Rate)

	if err != nil {
		log.Warnf(c, "[danmarks_national_bank_datasource.ToLatestExchangeRate] failed to parse rate, currency is %s, rate is %s", e.Currency, e.Rate)
		return nil
	}

	if rate <= 0 {
		log.Warnf(c, "[danmarks_national_bank_datasource.ToLatestExchangeRate] rate is invalid, currency is %s, rate is %s", e.Currency, e.Rate)
		return nil
	}

	finalRate := 100 / rate 

	if math.IsInf(finalRate, 0) {
		return nil
	}

	return &models.LatestExchangeRate{
		Currency: e.Currency,
		Rate:     utils.Float64ToString(finalRate),
	}
}


func (e *DanmarksNationalbankDataSource) BuildRequests() ([]*http.Request, error) {
	req, err := http.NewRequest("GET", danmarksNationalbankExchangeRateUrl, nil)

	if err != nil {
		return nil, err
	}

	return []*http.Request{req}, nil
}


func (e *DanmarksNationalbankDataSource) Parse(c core.Context, content []byte) (*models.LatestExchangeRateResponse, error) {
	xmlDecoder := xml.NewDecoder(bytes.NewReader(content))
	xmlDecoder.CharsetReader = charset.NewReaderLabel

	danmarksNationalbankData := &DanmarksNationalbankExchangeRateData{}
	err := xmlDecoder.Decode(danmarksNationalbankData)

	if err != nil {
		log.Errorf(c, "[danmarks_national_bank_datasource.Parse] failed to parse xml data, content is %s, because %s", string(content), err.Error())
		return nil, errs.ErrFailedToRequestRemoteApi
	}

	latestExchangeRateResponse := danmarksNationalbankData.ToLatestExchangeRateResponse(c)

	if latestExchangeRateResponse == nil {
		log.Errorf(c, "[danmarks_national_bank_datasource.Parse] failed to parse latest exchange rate data, content is %s", string(content))
		return nil, errs.ErrFailedToRequestRemoteApi
	}

	return latestExchangeRateResponse, nil
}
