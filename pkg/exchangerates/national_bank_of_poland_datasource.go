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

const nationalBankOfPolandDailyExchangeRateUrl = "https://api.nbp.pl/api/exchangerates/tables/A?format=xml"
const nationalBankOfPolandInconvertibleCurrencyExchangeRateUrl = "https://api.nbp.pl/api/exchangerates/tables/B?format=xml"
const nationalBankOfPolandExchangeRateReferenceUrl = "https://nbp.pl/en/statistic-and-financial-reporting/rates/"
const nationalBankOfPolandDataSource = "Narodowy Bank Polski"
const nationalBankOfPolandBaseCurrency = "PLN"

const nationalBankOfPolandDataUpdateDateFormat = "2006-01-02 15:04"
const nationalBankOfPolandDataUpdateDateTimezone = "Europe/Warsaw"


type NationalBankOfPolandDataSource struct {
	HttpExchangeRatesDataSource
}


type NationalBankOfPolandExchangeRateData struct {
	XMLName          xml.Name                            `xml:"ArrayOfExchangeRatesTable"`
	Date             string                              `xml:"ExchangeRatesTable>EffectiveDate"`
	AllExchangeRates []*NationalBankOfPolandExchangeRate `xml:"ExchangeRatesTable>Rates>Rate"`
}


type NationalBankOfPolandExchangeRate struct {
	Currency string `xml:"Code"`
	Rate     string `xml:"Mid"`
}


func (e *NationalBankOfPolandExchangeRateData) ToLatestExchangeRateResponse(c core.Context) *models.LatestExchangeRateResponse {
	if len(e.AllExchangeRates) < 1 {
		log.Errorf(c, "[national_bank_of_poland_datasource.ToLatestExchangeRateResponse] all exchange rates is empty")
		return nil
	}

	exchangeRates := make(models.LatestExchangeRateSlice, 0, len(e.AllExchangeRates))

	for i := 0; i < len(e.AllExchangeRates); i++ {
		exchangeRate := e.AllExchangeRates[i]

		if _, exists := validators.AllCurrencyNames[exchangeRate.Currency]; !exists {
			continue
		}

		finalExchangeRate := exchangeRate.ToLatestExchangeRate(c)

		if finalExchangeRate == nil {
			continue
		}

		exchangeRates = append(exchangeRates, finalExchangeRate)
	}

	timezone, err := time.LoadLocation(nationalBankOfPolandDataUpdateDateTimezone)

	if err != nil {
		log.Errorf(c, "[national_bank_of_poland_datasource.ToLatestExchangeRateResponse] failed to get timezone, timezone name is %s", nationalBankOfPolandDataUpdateDateTimezone)
		return nil
	}

	updateDateTime := e.Date + " 12:15" 
	updateTime, err := time.ParseInLocation(nationalBankOfPolandDataUpdateDateFormat, updateDateTime, timezone)

	if err != nil {
		log.Errorf(c, "[national_bank_of_poland_datasource.ToLatestExchangeRateResponse] failed to parse update date, datetime is %s", updateDateTime)
		return nil
	}

	latestExchangeRateResp := &models.LatestExchangeRateResponse{
		DataSource:    nationalBankOfPolandDataSource,
		ReferenceUrl:  nationalBankOfPolandExchangeRateReferenceUrl,
		UpdateTime:    updateTime.Unix(),
		BaseCurrency:  nationalBankOfPolandBaseCurrency,
		ExchangeRates: exchangeRates,
	}

	return latestExchangeRateResp
}


func (e *NationalBankOfPolandExchangeRate) ToLatestExchangeRate(c core.Context) *models.LatestExchangeRate {
	rate, err := utils.StringToFloat64(e.Rate)

	if err != nil {
		log.Warnf(c, "[national_bank_of_poland_datasource.ToLatestExchangeRate] failed to parse rate, currency is %s, rate is %s", e.Currency, e.Rate)
		return nil
	}

	if rate <= 0 {
		log.Warnf(c, "[national_bank_of_poland_datasource.ToLatestExchangeRate] rate is invalid, currency is %s, rate is %s", e.Currency, e.Rate)
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


func (e *NationalBankOfPolandDataSource) GetRequestUrls() []string {
	return []string{nationalBankOfPolandInconvertibleCurrencyExchangeRateUrl, nationalBankOfPolandDailyExchangeRateUrl}
}


func (e *NationalBankOfPolandDataSource) BuildRequests() ([]*http.Request, error) {
	inconvertibleCurrencyReq, err := http.NewRequest("GET", nationalBankOfPolandInconvertibleCurrencyExchangeRateUrl, nil)

	if err != nil {
		return nil, err
	}

	dailyReq, err := http.NewRequest("GET", nationalBankOfPolandDailyExchangeRateUrl, nil)

	if err != nil {
		return nil, err
	}

	return []*http.Request{inconvertibleCurrencyReq, dailyReq}, nil
}


func (e *NationalBankOfPolandDataSource) Parse(c core.Context, content []byte) (*models.LatestExchangeRateResponse, error) {
	xmlDecoder := xml.NewDecoder(bytes.NewReader(content))
	xmlDecoder.CharsetReader = charset.NewReaderLabel

	nationalBankOfPolandData := &NationalBankOfPolandExchangeRateData{}
	err := xmlDecoder.Decode(nationalBankOfPolandData)

	if err != nil {
		log.Errorf(c, "[national_bank_of_poland_datasource.Parse] failed to parse xml data, content is %s, because %s", string(content), err.Error())
		return nil, errs.ErrFailedToRequestRemoteApi
	}

	latestExchangeRateResponse := nationalBankOfPolandData.ToLatestExchangeRateResponse(c)

	if latestExchangeRateResponse == nil {
		log.Errorf(c, "[national_bank_of_poland_datasource.Parse] failed to parse latest exchange rate data, content is %s", string(content))
		return nil, errs.ErrFailedToRequestRemoteApi
	}

	return latestExchangeRateResponse, nil
}
