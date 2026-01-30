package ofx

import (
	"fmt"
	"strings"

	"github.com/Shavitjnr/split-chill-ai/pkg/converters/datatable"
	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
	"github.com/Shavitjnr/split-chill-ai/pkg/log"
	"github.com/Shavitjnr/split-chill-ai/pkg/models"
	"github.com/Shavitjnr/split-chill-ai/pkg/utils"
)

var ofxTransactionSupportedColumns = map[datatable.TransactionDataTableColumn]bool{
	datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TIME:         true,
	datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TIMEZONE:     true,
	datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TYPE:         true,
	datatable.TRANSACTION_DATA_TABLE_SUB_CATEGORY:             true,
	datatable.TRANSACTION_DATA_TABLE_ACCOUNT_NAME:             true,
	datatable.TRANSACTION_DATA_TABLE_ACCOUNT_CURRENCY:         true,
	datatable.TRANSACTION_DATA_TABLE_AMOUNT:                   true,
	datatable.TRANSACTION_DATA_TABLE_RELATED_ACCOUNT_NAME:     true,
	datatable.TRANSACTION_DATA_TABLE_RELATED_ACCOUNT_CURRENCY: true,
	datatable.TRANSACTION_DATA_TABLE_RELATED_AMOUNT:           true,
	datatable.TRANSACTION_DATA_TABLE_DESCRIPTION:              true,
}


type ofxTransactionData struct {
	ofxBaseStatementTransaction
	DefaultCurrency   string
	FromAccountId     string
	FromCreditAccount bool
	ToAccountId       string
}


type ofxTransactionDataTable struct {
	allData []*ofxTransactionData
}


type ofxTransactionDataRow struct {
	dataTable  *ofxTransactionDataTable
	data       *ofxTransactionData
	finalItems map[datatable.TransactionDataTableColumn]string
}


type ofxTransactionDataRowIterator struct {
	dataTable    *ofxTransactionDataTable
	currentIndex int
}


func (t *ofxTransactionDataTable) HasColumn(column datatable.TransactionDataTableColumn) bool {
	_, exists := ofxTransactionSupportedColumns[column]
	return exists
}


func (t *ofxTransactionDataTable) TransactionRowCount() int {
	return len(t.allData)
}


func (t *ofxTransactionDataTable) TransactionRowIterator() datatable.TransactionDataRowIterator {
	return &ofxTransactionDataRowIterator{
		dataTable:    t,
		currentIndex: -1,
	}
}


func (r *ofxTransactionDataRow) IsValid() bool {
	return true
}


func (r *ofxTransactionDataRow) GetData(column datatable.TransactionDataTableColumn) string {
	_, exists := ofxTransactionSupportedColumns[column]

	if exists {
		return r.finalItems[column]
	}

	return ""
}


func (t *ofxTransactionDataRowIterator) HasNext() bool {
	return t.currentIndex+1 < len(t.dataTable.allData)
}


func (t *ofxTransactionDataRowIterator) Next(ctx core.Context, user *models.User) (daraRow datatable.TransactionDataRow, err error) {
	if t.currentIndex+1 >= len(t.dataTable.allData) {
		return nil, nil
	}

	t.currentIndex++

	data := t.dataTable.allData[t.currentIndex]
	rowItems, err := t.parseTransaction(ctx, user, data)

	if err != nil {
		log.Errorf(ctx, "[ofx_transaction_table.Next] cannot parsing transaction in row#%d, because %s", t.currentIndex, err.Error())
		return nil, err
	}

	return &ofxTransactionDataRow{
		dataTable:  t.dataTable,
		data:       data,
		finalItems: rowItems,
	}, nil
}

func (t *ofxTransactionDataRowIterator) parseTransaction(ctx core.Context, user *models.User, ofxTransaction *ofxTransactionData) (map[datatable.TransactionDataTableColumn]string, error) {
	data := make(map[datatable.TransactionDataTableColumn]string, len(ofxTransactionSupportedColumns))

	if ofxTransaction.PostedDate == "" {
		return nil, errs.ErrMissingTransactionTime
	}

	datetime, timezone, err := t.parseTransactionTimeAndTimeZone(ctx, ofxTransaction.PostedDate)

	if err != nil {
		return nil, err
	}

	data[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TIME] = datetime
	data[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TIMEZONE] = timezone

	if ofxTransaction.TransactionType == "" {
		return nil, errs.ErrTransactionTypeInvalid
	}

	if ofxTransaction.FromAccountId == "" {
		return nil, errs.ErrMissingAccountData
	}

	data[datatable.TRANSACTION_DATA_TABLE_ACCOUNT_NAME] = ofxTransaction.FromAccountId

	if ofxTransaction.Currency != "" {
		data[datatable.TRANSACTION_DATA_TABLE_ACCOUNT_CURRENCY] = ofxTransaction.Currency
	} else {
		data[datatable.TRANSACTION_DATA_TABLE_ACCOUNT_CURRENCY] = ofxTransaction.DefaultCurrency
	}

	if data[datatable.TRANSACTION_DATA_TABLE_ACCOUNT_CURRENCY] == "" {
		return nil, errs.ErrAccountCurrencyInvalid
	}

	if ofxTransaction.Amount == "" {
		return nil, errs.ErrAmountInvalid
	}

	amount, err := utils.ParseAmount(utils.TrimTrailingZerosInDecimal(strings.ReplaceAll(ofxTransaction.Amount, ",", "."))) 

	if err != nil {
		log.Errorf(ctx, "[ofx_transaction_table.parseTransaction] cannot parsing transaction amount \"%s\", because %s", ofxTransaction.Amount, err.Error())
		return nil, errs.ErrAmountInvalid
	}

	if transactionType, exists := ofxTransactionTypeMapping[ofxTransaction.TransactionType]; exists { 
		data[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TYPE] = utils.IntToString(int(transactionType))

		if data[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TYPE] == ofxTransactionTypeNameMapping[models.TRANSACTION_TYPE_INCOME] { 
			data[datatable.TRANSACTION_DATA_TABLE_AMOUNT] = utils.FormatAmount(amount)
		} else if data[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TYPE] == ofxTransactionTypeNameMapping[models.TRANSACTION_TYPE_EXPENSE] { 
			data[datatable.TRANSACTION_DATA_TABLE_AMOUNT] = utils.FormatAmount(-amount)
		} else { 
			if amount >= 0 { 
				data[datatable.TRANSACTION_DATA_TABLE_AMOUNT] = utils.FormatAmount(amount)
				data[datatable.TRANSACTION_DATA_TABLE_RELATED_ACCOUNT_NAME] = data[datatable.TRANSACTION_DATA_TABLE_ACCOUNT_NAME]
				data[datatable.TRANSACTION_DATA_TABLE_RELATED_ACCOUNT_CURRENCY] = data[datatable.TRANSACTION_DATA_TABLE_ACCOUNT_CURRENCY]
				data[datatable.TRANSACTION_DATA_TABLE_RELATED_AMOUNT] = data[datatable.TRANSACTION_DATA_TABLE_AMOUNT]
				data[datatable.TRANSACTION_DATA_TABLE_ACCOUNT_NAME] = ""
			} else { 
				data[datatable.TRANSACTION_DATA_TABLE_AMOUNT] = utils.FormatAmount(-amount)
				data[datatable.TRANSACTION_DATA_TABLE_RELATED_ACCOUNT_NAME] = ofxTransaction.ToAccountId
				data[datatable.TRANSACTION_DATA_TABLE_RELATED_ACCOUNT_CURRENCY] = data[datatable.TRANSACTION_DATA_TABLE_ACCOUNT_CURRENCY]
				data[datatable.TRANSACTION_DATA_TABLE_RELATED_AMOUNT] = data[datatable.TRANSACTION_DATA_TABLE_AMOUNT]
			}
		}
	} else { 
		if amount >= 0 { 
			data[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TYPE] = ofxTransactionTypeNameMapping[models.TRANSACTION_TYPE_INCOME]
			data[datatable.TRANSACTION_DATA_TABLE_AMOUNT] = utils.FormatAmount(amount)
		} else { 
			data[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TYPE] = ofxTransactionTypeNameMapping[models.TRANSACTION_TYPE_EXPENSE]
			data[datatable.TRANSACTION_DATA_TABLE_AMOUNT] = utils.FormatAmount(-amount)
		}
	}

	if data[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TYPE] != ofxTransactionTypeNameMapping[models.TRANSACTION_TYPE_TRANSFER] {
		if ofxTransaction.FromCreditAccount || ofxTransaction.TransactionType == ofxGenericCreditTransaction {
			if amount >= 0 { 
				data[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TYPE] = ofxTransactionTypeNameMapping[models.TRANSACTION_TYPE_TRANSFER]
				data[datatable.TRANSACTION_DATA_TABLE_AMOUNT] = utils.FormatAmount(amount)
				data[datatable.TRANSACTION_DATA_TABLE_RELATED_ACCOUNT_NAME] = data[datatable.TRANSACTION_DATA_TABLE_ACCOUNT_NAME]
				data[datatable.TRANSACTION_DATA_TABLE_RELATED_ACCOUNT_CURRENCY] = data[datatable.TRANSACTION_DATA_TABLE_ACCOUNT_CURRENCY]
				data[datatable.TRANSACTION_DATA_TABLE_RELATED_AMOUNT] = data[datatable.TRANSACTION_DATA_TABLE_AMOUNT]
				data[datatable.TRANSACTION_DATA_TABLE_ACCOUNT_NAME] = ""
			} else { 
				data[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TYPE] = ofxTransactionTypeNameMapping[models.TRANSACTION_TYPE_EXPENSE]
				data[datatable.TRANSACTION_DATA_TABLE_AMOUNT] = utils.FormatAmount(-amount)
			}
		}
	}

	if ofxTransaction.Memo != "" {
		data[datatable.TRANSACTION_DATA_TABLE_DESCRIPTION] = ofxTransaction.Memo
	} else if ofxTransaction.Name != "" {
		data[datatable.TRANSACTION_DATA_TABLE_DESCRIPTION] = ofxTransaction.Name
	} else if ofxTransaction.Payee != nil {
		data[datatable.TRANSACTION_DATA_TABLE_DESCRIPTION] = ofxTransaction.Payee.Name
	} else {
		data[datatable.TRANSACTION_DATA_TABLE_DESCRIPTION] = ""
	}

	return data, nil
}

func (t *ofxTransactionDataRowIterator) parseTransactionTimeAndTimeZone(ctx core.Context, datetime string) (string, string, error) {
	if len(datetime) < 8 {
		return "", "", errs.ErrTransactionTimeInvalid
	}

	var err error
	var year, month, day string
	hour := "00"
	minute := "00"
	second := "00"
	tzOffset := ofxDefaultTimezoneOffset

	if len(datetime) >= 8 { 
		if !utils.IsStringOnlyContainsDigits(datetime[0:8]) {
			log.Errorf(ctx, "[ofx_transaction_table.parseTransactionTimeAndTimeZone] cannot parse time \"%s\", because contains non-digit character", datetime)
			return "", "", errs.ErrTransactionTimeInvalid
		}

		year = datetime[0:4]
		month = datetime[4:6]
		day = datetime[6:8]
	}

	if len(datetime) >= 14 { 
		if !utils.IsStringOnlyContainsDigits(datetime[8:14]) {
			log.Errorf(ctx, "[ofx_transaction_table.parseTransactionTimeAndTimeZone] cannot parse time \"%s\", because contains non-digit character", datetime)
			return "", "", errs.ErrTransactionTimeInvalid
		}

		hour = datetime[8:10]
		minute = datetime[10:12]
		second = datetime[12:14]
	}

	squareBracketStartIndex := strings.Index(datetime, "[")

	if squareBracketStartIndex > 0 { 
		timezoneInfo := datetime[squareBracketStartIndex+1 : len(datetime)-1]
		timezoneItems := strings.Split(timezoneInfo, ":")
		tzOffset, err = utils.FormatTimezoneOffsetFromHoursOffset(timezoneItems[0])

		if err != nil {
			log.Errorf(ctx, "[ofx_transaction_table.parseTransactionTimeAndTimeZone] cannot parse timezone offset \"%s\", because %s", timezoneInfo, err.Error())
			return "", "", errs.ErrTransactionTimeZoneInvalid
		}
	}

	return fmt.Sprintf("%s-%s-%s %s:%s:%s", year, month, day, hour, minute, second), tzOffset, nil
}

func createNewOFXTransactionDataTable(file *ofxFile) (*ofxTransactionDataTable, error) {
	if file == nil {
		return nil, errs.ErrNotFoundTransactionDataInFile
	}

	allData := make([]*ofxTransactionData, 0)

	if file.BankMessageResponseV1 != nil &&
		file.BankMessageResponseV1.StatementTransactionResponse != nil &&
		file.BankMessageResponseV1.StatementTransactionResponse.StatementResponse != nil &&
		file.BankMessageResponseV1.StatementTransactionResponse.StatementResponse.TransactionList != nil {
		statement := file.BankMessageResponseV1.StatementTransactionResponse.StatementResponse
		bankTransactions := statement.TransactionList.StatementTransactions
		fromAccountId := ""
		fromCreditAccount := false

		if statement.AccountFrom != nil {
			fromAccountId = statement.AccountFrom.AccountId

			if statement.AccountFrom.AccountType == ofxLineOfCreditAccount {
				fromCreditAccount = true
			}
		}

		for i := 0; i < len(bankTransactions); i++ {
			toAccountId := ""

			if bankTransactions[i].AccountTo != nil {
				toAccountId = bankTransactions[i].AccountTo.AccountId
			}

			allData = append(allData, &ofxTransactionData{
				ofxBaseStatementTransaction: bankTransactions[i].ofxBaseStatementTransaction,
				DefaultCurrency:             statement.DefaultCurrency,
				FromAccountId:               fromAccountId,
				FromCreditAccount:           fromCreditAccount,
				ToAccountId:                 toAccountId,
			})
		}
	}

	if file.CreditCardMessageResponseV1 != nil &&
		file.CreditCardMessageResponseV1.StatementTransactionResponse != nil &&
		file.CreditCardMessageResponseV1.StatementTransactionResponse.StatementResponse != nil &&
		file.CreditCardMessageResponseV1.StatementTransactionResponse.StatementResponse.TransactionList != nil {
		statement := file.CreditCardMessageResponseV1.StatementTransactionResponse.StatementResponse
		bankTransactions := statement.TransactionList.StatementTransactions
		fromAccountId := ""

		if statement.AccountFrom != nil {
			fromAccountId = statement.AccountFrom.AccountId
		}

		for i := 0; i < len(bankTransactions); i++ {
			toAccountId := ""

			if bankTransactions[i].AccountTo != nil {
				toAccountId = bankTransactions[i].AccountTo.AccountId
			}

			allData = append(allData, &ofxTransactionData{
				ofxBaseStatementTransaction: bankTransactions[i].ofxBaseStatementTransaction,
				DefaultCurrency:             statement.DefaultCurrency,
				FromAccountId:               fromAccountId,
				FromCreditAccount:           true,
				ToAccountId:                 toAccountId,
			})
		}
	}

	return &ofxTransactionDataTable{
		allData: allData,
	}, nil
}
