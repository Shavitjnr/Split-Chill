package models

import (
	"strings"

	"github.com/Shavitjnr/split-chill-ai/pkg/utils"
)

const UserCustomExchangeRateFactorInDatabase = int64(100000000)


type UserCustomExchangeRate struct {
	Uid             int64  `xorm:"PK NOT NULL"`
	DeletedUnixTime int64  `xorm:"PK NOT NULL"`
	Currency        string `xorm:"PK VARCHAR(3) NOT NULL"`
	Rate            int64  `xorm:"NOT NULL"`
	CreatedUnixTime int64
	UpdatedUnixTime int64
}


type UserCustomExchangeRateUpdateRequest struct {
	Currency string `json:"currency" binding:"required,len=3,validCurrency"`
	Rate     string `json:"rate"`
}


type UserCustomExchangeRateDeleteRequest struct {
	Currency string `json:"currency" binding:"required,len=3,validCurrency"`
}


type UserCustomExchangeRateUpdateResponse struct {
	LatestExchangeRate
	UpdateTime int64 `json:"updateTime"`
}


type LatestExchangeRateResponse struct {
	DataSource    string                  `json:"dataSource"`
	ReferenceUrl  string                  `json:"referenceUrl"`
	UpdateTime    int64                   `json:"updateTime"`
	BaseCurrency  string                  `json:"baseCurrency"`
	ExchangeRates LatestExchangeRateSlice `json:"exchangeRates"`
}


type LatestExchangeRate struct {
	Currency string `json:"currency"`
	Rate     string `json:"rate"`
}


func (r *UserCustomExchangeRate) ToLatestExchangeRate(baseCurrencyRate int64) *LatestExchangeRate {
	rate := float64(0)

	if baseCurrencyRate > 0 {
		rate = float64(r.Rate) / float64(baseCurrencyRate)
	}

	return &LatestExchangeRate{
		Currency: r.Currency,
		Rate:     utils.Float64ToString(rate),
	}
}


func (r *UserCustomExchangeRate) ToUserCustomExchangeRateUpdateResponse(baseCurrencyRate int64) *UserCustomExchangeRateUpdateResponse {
	return &UserCustomExchangeRateUpdateResponse{
		LatestExchangeRate: *r.ToLatestExchangeRate(baseCurrencyRate),
		UpdateTime:         r.UpdatedUnixTime,
	}
}


func CreateUserCustomExchangeRate(uid int64, currency string, exchangeRate string, baseCurrencyRate int64) (*UserCustomExchangeRate, error) {
	if baseCurrencyRate <= 0 {
		return &UserCustomExchangeRate{
			Uid:      uid,
			Currency: currency,
			Rate:     UserCustomExchangeRateFactorInDatabase,
		}, nil
	}

	rate, err := utils.StringToFloat64(exchangeRate)

	if err != nil {
		return nil, err
	}

	rate = rate * float64(baseCurrencyRate)

	return &UserCustomExchangeRate{
		Uid:      uid,
		Currency: currency,
		Rate:     int64(rate),
	}, nil
}


type LatestExchangeRateSlice []*LatestExchangeRate


func (s LatestExchangeRateSlice) Len() int {
	return len(s)
}


func (s LatestExchangeRateSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}


func (s LatestExchangeRateSlice) Less(i, j int) bool {
	return strings.Compare(s[i].Currency, s[j].Currency) < 0
}
