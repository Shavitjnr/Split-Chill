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

const reserveBankOfAustraliaExchangeRateUrl = "https://www.rba.gov.au/rss/rss-cb-exchange-rates.xml"
const reserveBankOfAustraliaExchangeRateReferenceUrl = "https://www.rba.gov.au/statistics/frequency/exchange-rates.html"
const reserveBankOfAustraliaDataSource = "Reserve Bank of Australia"
const reserveBankOfAustraliaBaseCurrency = "AUD"

const reserveBankOfAustraliaDataUpdateDateFormat = "2006-01-02T15:04:05Z07:00"


type ReserveBankOfAustraliaDataSource struct {
	HttpExchangeRatesDataSource
}


type ReserveBankOfAustraliaData struct {
	XMLName xml.Name                          `xml:"RDF"`
	Channel *ReserveBankOfAustraliaRssChannel `xml:"channel"`
	Items   []*ReserveBankOfAustraliaRssItem  `xml:"item"`
}


type ReserveBankOfAustraliaRssChannel struct {
	Date string `xml:"date"`
}


type ReserveBankOfAustraliaRssItem struct {
	Statistics *ReserveBankOfAustraliaItemStatistics `xml:"statistics"`
}


type ReserveBankOfAustraliaItemStatistics struct {
	ExchangeRate *ReserveBankOfAustraliaExchangeRate `xml:"exchangeRate"`
}


type ReserveBankOfAustraliaExchangeRate struct {
	BaseCurrency   string                                         `xml:"baseCurrency"`
	TargetCurrency string                                         `xml:"targetCurrency"`
	Observation    *ReserveBankOfAustraliaExchangeRateObservation `xml:"observation"`
}


type ReserveBankOfAustraliaExchangeRateObservation struct {
	Value string `xml:"value"`
	Unit  string `xml:"unit"`
}


func (e *ReserveBankOfAustraliaData) ToLatestExchangeRateResponse(c core.Context) *models.LatestExchangeRateResponse {
	if e.Channel == nil {
		log.Errorf(c, "[reserve_bank_of_australia_datasource.ToLatestExchangeRateResponse] rss channel does not exist")
		return nil
	}

	if len(e.Items) < 1 {
		log.Errorf(c, "[reserve_bank_of_australia_datasource.ToLatestExchangeRateResponse] rss items is empty")
		return nil
	}

	exchangeRates := make(models.LatestExchangeRateSlice, 0, len(e.Items))

	for i := 0; i < len(e.Items); i++ {
		item := e.Items[i]

		if item.Statistics == nil || item.Statistics.ExchangeRate == nil || item.Statistics.ExchangeRate.Observation == nil {
			continue
		}

		if item.Statistics.ExchangeRate.BaseCurrency != reserveBankOfAustraliaBaseCurrency || item.Statistics.ExchangeRate.Observation.Unit != reserveBankOfAustraliaBaseCurrency {
			continue
		}

		if _, exists := validators.AllCurrencyNames[item.Statistics.ExchangeRate.TargetCurrency]; !exists {
			continue
		}

		if _, err := utils.StringToFloat64(item.Statistics.ExchangeRate.Observation.Value); err != nil {
			continue
		}

		exchangeRates = append(exchangeRates, item.Statistics.ExchangeRate.ToLatestExchangeRate())
	}

	updateDateTime := e.Channel.Date
	updateTime, err := time.Parse(reserveBankOfAustraliaDataUpdateDateFormat, updateDateTime)

	if err != nil {
		log.Errorf(c, "[reserve_bank_of_australia_datasource.ToLatestExchangeRateResponse] failed to parse update date, datetime is %s", updateDateTime)
		return nil
	}

	latestExchangeRateResp := &models.LatestExchangeRateResponse{
		DataSource:    reserveBankOfAustraliaDataSource,
		ReferenceUrl:  reserveBankOfAustraliaExchangeRateReferenceUrl,
		UpdateTime:    updateTime.Unix(),
		BaseCurrency:  reserveBankOfAustraliaBaseCurrency,
		ExchangeRates: exchangeRates,
	}

	return latestExchangeRateResp
}


func (e *ReserveBankOfAustraliaExchangeRate) ToLatestExchangeRate() *models.LatestExchangeRate {
	return &models.LatestExchangeRate{
		Currency: e.TargetCurrency,
		Rate:     e.Observation.Value,
	}
}


func (e *ReserveBankOfAustraliaDataSource) BuildRequests() ([]*http.Request, error) {
	req, err := http.NewRequest("GET", reserveBankOfAustraliaExchangeRateUrl, nil)

	if err != nil {
		return nil, err
	}

	return []*http.Request{req}, nil
}


func (e *ReserveBankOfAustraliaDataSource) Parse(c core.Context, content []byte) (*models.LatestExchangeRateResponse, error) {
	xmlDecoder := xml.NewDecoder(bytes.NewReader(content))
	xmlDecoder.CharsetReader = charset.NewReaderLabel

	reserveBankOfAustraliaData := &ReserveBankOfAustraliaData{}
	err := xmlDecoder.Decode(reserveBankOfAustraliaData)

	if err != nil {
		log.Errorf(c, "[reserve_bank_of_australia_datasource.Parse] failed to parse xml data, content is %s, because %s", string(content), err.Error())
		return nil, errs.ErrFailedToRequestRemoteApi
	}

	latestExchangeRateResponse := reserveBankOfAustraliaData.ToLatestExchangeRateResponse(c)

	if latestExchangeRateResponse == nil {
		log.Errorf(c, "[reserve_bank_of_australia_datasource.Parse] failed to parse latest exchange rate data, content is %s", string(content))
		return nil, errs.ErrFailedToRequestRemoteApi
	}

	return latestExchangeRateResponse, nil
}
