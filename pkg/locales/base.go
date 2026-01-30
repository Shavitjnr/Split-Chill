package locales

import (
	"github.com/Shavitjnr/split-chill-ai/pkg/core"
)


type LocaleTextItems struct {
	GlobalTextItems             *GlobalTextItems
	DefaultTypes                *DefaultTypes
	DataConverterTextItems      *DataConverterTextItems
	VerifyEmailTextItems        *VerifyEmailTextItems
	ForgetPasswordMailTextItems *ForgetPasswordMailTextItems
}


type GlobalTextItems struct {
	AppName string
}


type DefaultTypes struct {
	DecimalSeparator    core.DecimalSeparator
	DigitGroupingSymbol core.DigitGroupingSymbol
}


type DataConverterTextItems struct {
	Alipay       string
	WeChatWallet string
}


type VerifyEmailTextItems struct {
	Title                     string
	SalutationFormat          string
	DescriptionAboveBtn       string
	VerifyEmail               string
	DescriptionBelowBtnFormat string
}


type ForgetPasswordMailTextItems struct {
	Title                     string
	SalutationFormat          string
	DescriptionAboveBtn       string
	ResetPassword             string
	DescriptionBelowBtnFormat string
}
