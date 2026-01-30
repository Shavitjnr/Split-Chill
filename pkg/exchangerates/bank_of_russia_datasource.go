package exchangerates

import (
	"bytes"
	"encoding/xml"
	"math"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html/charset"

	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
	"github.com/Shavitjnr/split-chill-ai/pkg/log"
	"github.com/Shavitjnr/split-chill-ai/pkg/models"
	"github.com/Shavitjnr/split-chill-ai/pkg/utils"
	"github.com/Shavitjnr/split-chill-ai/pkg/validators"
)

const bankOfRussiaExchangeRateUrl = "https://cbr.ru/scripts/XML_daily_eng.asp"
const bankOfRussiaExchangeRateReferenceUrl = "https://www.cbr.ru/eng/currency_base/daily/"
const bankOfRussiaDataSource = "Банк России"
const bankOfRussiaBaseCurrency = "RUB"

const bankOfRussiaUpdateDateFormat = "02.01.2006 15:04"
const bankOfRussiaUpdateDateTimezone = "Europe/Moscow"


type BankOfRussiaDataSource struct {
	HttpExchangeRatesDataSource
}


type BankOfRussiaExchangeRateData struct {
	XMLName       xml.Name                    `xml:"ValCurs"`
	Date          string                      `xml:"Date,attr"`
	ExchangeRates []*BankOfRussiaExchangeRate `xml:"Valute"`
}


type BankOfRussiaExchangeRate struct {
	Currency string `xml:"CharCode"`
	Rate     string `xml:"VunitRate"`
}


func (e *BankOfRussiaExchangeRateData) ToLatestExchangeRateResponse(c core.Context) *models.LatestExchangeRateResponse {
	if len(e.ExchangeRates) < 1 {
		log.Errorf(c, "[bank_of_russia_datasource.ToLatestExchangeRateResponse] all exchange rates is empty")
		return nil
	}

	exchangeRates := make(models.LatestExchangeRateSlice, 0, len(e.ExchangeRates))

	for i := 0; i < len(e.ExchangeRates); i++ {
		exchangeRate := e.ExchangeRates[i]

		if _, exists := validators.AllCurrencyNames[exchangeRate.Currency]; !exists {
			continue
		}

		finalExchangeRate := exchangeRate.ToLatestExchangeRate(c)

		if finalExchangeRate == nil {
			continue
		}

		exchangeRates = append(exchangeRates, finalExchangeRate)
	}

	timezone, err := time.LoadLocation(bankOfRussiaUpdateDateTimezone)

	if err != nil {
		log.Errorf(c, "[bank_of_russia_datasource.ToLatestExchangeRateResponse] failed to get timezone, timezone name is %s", bankOfRussiaUpdateDateTimezone)
		return nil
	}

	updateDateTime := e.Date + " 15:30" 
	updateTime, err := time.ParseInLocation(bankOfRussiaUpdateDateFormat, updateDateTime, timezone)

	if err != nil {
		log.Errorf(c, "[bank_of_russia_datasource.ToLatestExchangeRateResponse] failed to parse update date, datetime is %s", updateDateTime)
		return nil
	}

	latestExchangeRateResp := &models.LatestExchangeRateResponse{
		DataSource:    bankOfRussiaDataSource,
		ReferenceUrl:  bankOfRussiaExchangeRateReferenceUrl,
		UpdateTime:    updateTime.Unix(),
		BaseCurrency:  bankOfRussiaBaseCurrency,
		ExchangeRates: exchangeRates,
	}

	return latestExchangeRateResp
}


func (e *BankOfRussiaExchangeRate) ToLatestExchangeRate(c core.Context) *models.LatestExchangeRate {
	rate, err := utils.StringToFloat64(strings.ReplaceAll(e.Rate, ",", "."))

	if err != nil {
		log.Warnf(c, "[bank_of_russia_datasource.ToLatestExchangeRate] failed to parse rate, currency is %s, rate is %s", e.Currency, e.Rate)
		return nil
	}

	if rate <= 0 {
		log.Warnf(c, "[bank_of_russia_datasource.ToLatestExchangeRate] rate is invalid, currency is %s, rate is %s", e.Currency, e.Rate)
		return nil
	}

	finalRate := 1 / rate

	if math.IsInf(finalRate, 0) {
		return nil
	}

	return &models.LatestExchangeRate{
		Currency: e.Currency,
		Rate:     utils.Float64ToString(finalRate),
	}
}


func (e *BankOfRussiaDataSource) BuildRequests() ([]*http.Request, error) {
	req, err := http.NewRequest("GET", bankOfRussiaExchangeRateUrl, nil)

	if err != nil {
		return nil, err
	}

	return []*http.Request{req}, nil
}


func (e *BankOfRussiaDataSource) Parse(c core.Context, content []byte) (*models.LatestExchangeRateResponse, error) {
	xmlDecoder := xml.NewDecoder(bytes.NewReader(content))
	xmlDecoder.CharsetReader = charset.NewReaderLabel

	bankOfRussiaData := &BankOfRussiaExchangeRateData{}
	err := xmlDecoder.Decode(bankOfRussiaData)

	if err != nil {
		log.Errorf(c, "[bank_of_russia_datasource.Parse] failed to parse xml data, content is %s, because %s", string(content), err.Error())
		return nil, errs.ErrFailedToRequestRemoteApi
	}

	latestExchangeRateResponse := bankOfRussiaData.ToLatestExchangeRateResponse(c)

	if latestExchangeRateResponse == nil {
		log.Errorf(c, "[bank_of_russia_datasource.Parse] failed to parse latest exchange rate data, content is %s", string(content))
		return nil, errs.ErrFailedToRequestRemoteApi
	}

	return latestExchangeRateResponse, nil
}
