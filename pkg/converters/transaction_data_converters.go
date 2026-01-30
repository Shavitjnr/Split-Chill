package converters

import (
	"github.com/Shavitjnr/split-chill-ai/pkg/converters/alipay"
	"github.com/Shavitjnr/split-chill-ai/pkg/converters/beancount"
	"github.com/Shavitjnr/split-chill-ai/pkg/converters/camt"
	"github.com/Shavitjnr/split-chill-ai/pkg/converters/converter"
	"github.com/Shavitjnr/split-chill-ai/pkg/converters/datatable"
	"github.com/Shavitjnr/split-chill-ai/pkg/converters/default"
	"github.com/Shavitjnr/split-chill-ai/pkg/converters/dsv"
	"github.com/Shavitjnr/split-chill-ai/pkg/converters/feidee"
	"github.com/Shavitjnr/split-chill-ai/pkg/converters/fireflyIII"
	"github.com/Shavitjnr/split-chill-ai/pkg/converters/gnucash"
	"github.com/Shavitjnr/split-chill-ai/pkg/converters/iif"
	"github.com/Shavitjnr/split-chill-ai/pkg/converters/jdcom"
	"github.com/Shavitjnr/split-chill-ai/pkg/converters/mt"
	"github.com/Shavitjnr/split-chill-ai/pkg/converters/ofx"
	"github.com/Shavitjnr/split-chill-ai/pkg/converters/qif"
	"github.com/Shavitjnr/split-chill-ai/pkg/converters/wechat"
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
	"github.com/Shavitjnr/split-chill-ai/pkg/models"
)


func GetTransactionDataExporter(fileType string) converter.TransactionDataExporter {
	if fileType == "csv" {
		return _default.DefaultTransactionDataCSVFileConverter
	} else if fileType == "tsv" {
		return _default.DefaultTransactionDataTSVFileConverter
	} else {
		return nil
	}
}


func GetTransactionDataImporter(fileType string) (converter.TransactionDataImporter, error) {
	if fileType == "Split Chill AI_csv" {
		return _default.DefaultTransactionDataCSVFileConverter, nil
	} else if fileType == "Split Chill AI_tsv" {
		return _default.DefaultTransactionDataTSVFileConverter, nil
	} else if fileType == "Split Chill AI_json" {
		return _default.DefaultTransactionDataJsonFileImporter, nil
	} else if fileType == "ofx" {
		return ofx.OFXTransactionDataImporter, nil
	} else if fileType == "qfx" {
		return ofx.OFXTransactionDataImporter, nil
	} else if fileType == "qif_ymd" {
		return qif.QifYearMonthDayTransactionDataImporter, nil
	} else if fileType == "qif_mdy" {
		return qif.QifMonthDayYearTransactionDataImporter, nil
	} else if fileType == "qif_dmy" {
		return qif.QifDayMonthYearTransactionDataImporter, nil
	} else if fileType == "iif" {
		return iif.IifTransactionDataFileImporter, nil
	} else if fileType == "camt052" {
		return camt.Camt052TransactionDataImporter, nil
	} else if fileType == "camt053" {
		return camt.Camt053TransactionDataImporter, nil
	} else if fileType == "mt940" {
		return mt.MT940TransactionDataFileImporter, nil
	} else if fileType == "gnucash" {
		return gnucash.GnuCashTransactionDataImporter, nil
	} else if fileType == "firefly_iii_csv" {
		return fireflyIII.FireflyIIITransactionDataCsvFileImporter, nil
	} else if fileType == "beancount" {
		return beancount.BeancountTransactionDataImporter, nil
	} else if fileType == "feidee_mymoney_csv" {
		return feidee.FeideeMymoneyAppTransactionDataCsvFileImporter, nil
	} else if fileType == "feidee_mymoney_xls" {
		return feidee.FeideeMymoneyWebTransactionDataXlsFileImporter, nil
	} else if fileType == "feidee_mymoney_elecloud_xlsx" {
		return feidee.FeideeMymoneyElecloudTransactionDataXlsxFileImporter, nil
	} else if fileType == "alipay_app_csv" {
		return alipay.AlipayAppTransactionDataCsvFileImporter, nil
	} else if fileType == "alipay_web_csv" {
		return alipay.AlipayWebTransactionDataCsvFileImporter, nil
	} else if fileType == "wechat_pay_app_xlsx" {
		return wechat.WeChatPayTransactionDataXlsxFileImporter, nil
	} else if fileType == "wechat_pay_app_csv" {
		return wechat.WeChatPayTransactionDataCsvFileImporter, nil
	} else if fileType == "jdcom_finance_app_csv" {
		return jdcom.JDComFinanceTransactionDataCsvFileImporter, nil
	} else {
		return nil, errs.ErrImportFileTypeNotSupported
	}
}


func IsCustomDelimiterSeparatedValuesFileType(fileType string) bool {
	return dsv.IsDelimiterSeparatedValuesFileType(fileType)
}


func CreateNewDelimiterSeparatedValuesDataParser(fileType string, fileEncoding string) (dsv.CustomTransactionDataDsvFileParser, error) {
	return dsv.CreateNewCustomTransactionDataDsvFileParser(fileType, fileEncoding)
}


func CreateNewDelimiterSeparatedValuesDataImporter(fileType string, fileEncoding string, columnIndexMapping map[datatable.TransactionDataTableColumn]int, transactionTypeNameMapping map[string]models.TransactionType, hasHeaderLine bool, timeFormat string, timezoneFormat string, amountDecimalSeparator string, amountDigitGroupingSymbol string, geoLocationSeparator string, geoLocationOrder string, transactionTagSeparator string) (converter.TransactionDataImporter, error) {
	return dsv.CreateNewCustomTransactionDataDsvFileImporter(fileType, fileEncoding, columnIndexMapping, transactionTypeNameMapping, hasHeaderLine, timeFormat, timezoneFormat, amountDecimalSeparator, amountDigitGroupingSymbol, geoLocationSeparator, geoLocationOrder, transactionTagSeparator)
}
