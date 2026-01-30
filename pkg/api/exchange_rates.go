package api

import (
	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
	"github.com/Shavitjnr/split-chill-ai/pkg/exchangerates"
	"github.com/Shavitjnr/split-chill-ai/pkg/log"
	"github.com/Shavitjnr/split-chill-ai/pkg/models"
	"github.com/Shavitjnr/split-chill-ai/pkg/services"
	"github.com/Shavitjnr/split-chill-ai/pkg/settings"
)


type ExchangeRatesApi struct {
	ApiUsingConfig
	users                   *services.UserService
	userCustomExchangeRates *services.UserCustomExchangeRatesService
}


var (
	ExchangeRates = &ExchangeRatesApi{
		ApiUsingConfig: ApiUsingConfig{
			container: settings.Container,
		},
		users:                   services.Users,
		userCustomExchangeRates: services.UserCustomExchangeRates,
	}
)


func (a *ExchangeRatesApi) LatestExchangeRateHandler(c *core.WebContext) (any, *errs.Error) {
	exchangeRateResponse, err := exchangerates.Container.GetLatestExchangeRates(c, c.GetCurrentUid(), a.CurrentConfig())

	if err != nil {
		return nil, errs.Or(err, errs.ErrOperationFailed)
	}

	return exchangeRateResponse, nil
}


func (a *ExchangeRatesApi) UserCustomExchangeRateUpdateHandler(c *core.WebContext) (any, *errs.Error) {
	var customExchangeRateUpdateReq models.UserCustomExchangeRateUpdateRequest
	err := c.ShouldBindJSON(&customExchangeRateUpdateReq)

	if err != nil {
		log.Warnf(c, "[exchange_rates.UserCustomExchangeRateUpdateHandler] parse request failed, because %s", err.Error())
		return nil, errs.NewIncompleteOrIncorrectSubmissionError(err)
	}

	uid := c.GetCurrentUid()
	user, err := a.users.GetUserById(c, uid)

	if err != nil {
		log.Errorf(c, "[exchange_rates.UserCustomExchangeRateUpdateHandler] failed to get user \"uid:%d\", because %s", uid, err.Error())
		return nil, errs.Or(err, errs.ErrOperationFailed)
	}

	if customExchangeRateUpdateReq.Currency == user.DefaultCurrency {
		return nil, errs.ErrCannotUpdateExchangeRateForDefaultCurrency
	}

	newCustomExchangeRate, defaultCurrencyExchangeRate, err := a.userCustomExchangeRates.UpdateCustomExchangeRate(c, uid, customExchangeRateUpdateReq.Currency, customExchangeRateUpdateReq.Rate, user.DefaultCurrency)

	if err != nil {
		log.Errorf(c, "[exchange_rates.UserCustomExchangeRateUpdateHandler] failed to update user custom exchange rate \"currency:%s\" for user \"uid:%d\", because %s", customExchangeRateUpdateReq.Currency, uid, err.Error())
		return nil, errs.Or(err, errs.ErrOperationFailed)
	}

	log.Infof(c, "[exchange_rates.UserCustomExchangeRateUpdateHandler] user \"uid:%d\" has updated user custom exchange rate \"currency:%s\" successfully", uid, customExchangeRateUpdateReq.Currency)
	return newCustomExchangeRate.ToUserCustomExchangeRateUpdateResponse(defaultCurrencyExchangeRate.Rate), nil
}


func (a *ExchangeRatesApi) UserCustomExchangeRateDeleteHandler(c *core.WebContext) (any, *errs.Error) {
	var customExchangeRateDeleteReq models.UserCustomExchangeRateDeleteRequest
	err := c.ShouldBindJSON(&customExchangeRateDeleteReq)

	if err != nil {
		log.Warnf(c, "[exchange_rates.UserCustomExchangeRateDeleteHandler] parse request failed, because %s", err.Error())
		return nil, errs.NewIncompleteOrIncorrectSubmissionError(err)
	}

	uid := c.GetCurrentUid()
	user, err := a.users.GetUserById(c, uid)

	if err != nil {
		log.Errorf(c, "[exchange_rates.UserCustomExchangeRateDeleteHandler] failed to get user \"uid:%d\", because %s", uid, err.Error())
		return nil, errs.Or(err, errs.ErrOperationFailed)
	}

	if customExchangeRateDeleteReq.Currency == user.DefaultCurrency {
		return nil, errs.ErrCannotDeleteExchangeRateForDefaultCurrency
	}

	err = a.userCustomExchangeRates.DeleteCustomExchangeRate(c, uid, customExchangeRateDeleteReq.Currency)

	if err != nil {
		log.Errorf(c, "[exchange_rates.UserCustomExchangeRateDeleteHandler] failed to delete user custom exchange rate \"currency:%s\" for user \"uid:%d\", because %s", customExchangeRateDeleteReq.Currency, uid, err.Error())
		return nil, errs.Or(err, errs.ErrOperationFailed)
	}

	log.Infof(c, "[exchange_rates.UserCustomExchangeRateDeleteHandler] user \"uid:%d\" has deleted user custom exchange rate \"currency:%s\"", uid, customExchangeRateDeleteReq.Currency)
	return true, nil
}
