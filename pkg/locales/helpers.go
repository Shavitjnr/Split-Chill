package locales

import "github.com/Shavitjnr/split-chill-ai/pkg/core"


func GetLocaleTextItems(locale string) *LocaleTextItems {
	localeInfo, exists := AllLanguages[locale]

	if exists {
		return localeInfo.Content
	}

	return DefaultLanguage
}


func IsDecimalSeparatorEqualsDigitGroupingSymbol(decimalSeparator core.DecimalSeparator, digitGroupingSymbol core.DigitGroupingSymbol, locale string) bool {
	if decimalSeparator == core.DECIMAL_SEPARATOR_DEFAULT && digitGroupingSymbol == core.DIGIT_GROUPING_SYMBOL_DEFAULT {
		return false
	}

	if byte(decimalSeparator) == byte(digitGroupingSymbol) {
		return true
	}

	localeTextItems := GetLocaleTextItems(locale)

	if decimalSeparator == core.DECIMAL_SEPARATOR_DEFAULT {
		decimalSeparator = localeTextItems.DefaultTypes.DecimalSeparator
	}

	if digitGroupingSymbol == core.DIGIT_GROUPING_SYMBOL_DEFAULT {
		digitGroupingSymbol = localeTextItems.DefaultTypes.DigitGroupingSymbol
	}

	return byte(decimalSeparator) == byte(digitGroupingSymbol)
}
