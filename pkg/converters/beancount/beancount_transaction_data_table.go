package beancount

import (
	"strings"

	"github.com/Shavitjnr/split-chill-ai/pkg/converters/datatable"
	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
	"github.com/Shavitjnr/split-chill-ai/pkg/log"
	"github.com/Shavitjnr/split-chill-ai/pkg/models"
	"github.com/Shavitjnr/split-chill-ai/pkg/utils"
)

var beancountTransactionSupportedColumns = map[datatable.TransactionDataTableColumn]bool{
	datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TIME:         true,
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

var BEANCOUNT_TRANSACTION_TAG_SEPARATOR = "#"


type beancountTransactionDataTable struct {
	allData    []*beancountTransactionEntry
	accountMap map[string]*beancountAccount
}


type beancountTransactionDataRow struct {
	dataTable  *beancountTransactionDataTable
	data       *beancountTransactionEntry
	finalItems map[datatable.TransactionDataTableColumn]string
}


type beancountTransactionDataRowIterator struct {
	dataTable    *beancountTransactionDataTable
	currentIndex int
}


func (t *beancountTransactionDataTable) HasColumn(column datatable.TransactionDataTableColumn) bool {
	_, exists := beancountTransactionSupportedColumns[column]
	return exists
}


func (t *beancountTransactionDataTable) TransactionRowCount() int {
	return len(t.allData)
}


func (t *beancountTransactionDataTable) TransactionRowIterator() datatable.TransactionDataRowIterator {
	return &beancountTransactionDataRowIterator{
		dataTable:    t,
		currentIndex: -1,
	}
}


func (r *beancountTransactionDataRow) IsValid() bool {
	return true
}


func (r *beancountTransactionDataRow) GetData(column datatable.TransactionDataTableColumn) string {
	_, exists := beancountTransactionSupportedColumns[column]

	if exists {
		return r.finalItems[column]
	}

	return ""
}


func (t *beancountTransactionDataRowIterator) HasNext() bool {
	return t.currentIndex+1 < len(t.dataTable.allData)
}


func (t *beancountTransactionDataRowIterator) Next(ctx core.Context, user *models.User) (daraRow datatable.TransactionDataRow, err error) {
	if t.currentIndex+1 >= len(t.dataTable.allData) {
		return nil, nil
	}

	t.currentIndex++

	data := t.dataTable.allData[t.currentIndex]
	rowItems, err := t.parseTransaction(ctx, user, data)

	if err != nil {
		return nil, err
	}

	return &beancountTransactionDataRow{
		dataTable:  t.dataTable,
		data:       data,
		finalItems: rowItems,
	}, nil
}

func (t *beancountTransactionDataRowIterator) parseTransaction(ctx core.Context, user *models.User, beancountEntry *beancountTransactionEntry) (map[datatable.TransactionDataTableColumn]string, error) {
	data := make(map[datatable.TransactionDataTableColumn]string, len(beancountTransactionSupportedColumns))

	if beancountEntry.Date == "" {
		return nil, errs.ErrMissingTransactionTime
	}

	
	data[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TIME] = strings.ReplaceAll(beancountEntry.Date, "/", "-") + " 00:00:00"

	if len(beancountEntry.Postings) == 2 {
		splitData1 := beancountEntry.Postings[0]
		splitData2 := beancountEntry.Postings[1]

		account1 := t.dataTable.accountMap[splitData1.Account]
		account2 := t.dataTable.accountMap[splitData2.Account]

		if account1 == nil || account2 == nil {
			return nil, errs.ErrMissingAccountData
		}

		amount1, err := utils.ParseAmount(splitData1.Amount)

		if err != nil {
			log.Errorf(ctx, "[beancount_transaction_data_table.parseTransaction] cannot parse amount \"%s\", because %s", splitData1.Amount, err.Error())
			return nil, errs.ErrAmountInvalid
		}

		amount2, err := utils.ParseAmount(splitData2.Amount)

		if err != nil {
			log.Errorf(ctx, "[beancount_transaction_data_table.parseTransaction] cannot parse amount \"%s\", because %s", splitData2.Amount, err.Error())
			return nil, errs.ErrAmountInvalid
		}

		if ((account1.AccountType == beancountEquityAccountType || account1.AccountType == beancountIncomeAccountType) && (account2.AccountType == beancountAssetsAccountType || account2.AccountType == beancountLiabilitiesAccountType)) ||
			((account2.AccountType == beancountEquityAccountType || account2.AccountType == beancountIncomeAccountType) && (account1.AccountType == beancountAssetsAccountType || account1.AccountType == beancountLiabilitiesAccountType)) { 
			fromAccount := account1
			toAccount := account2
			toCurrency := splitData2.Commodity
			toAmount := amount2

			if (account2.AccountType == beancountEquityAccountType || account2.AccountType == beancountIncomeAccountType) && (account1.AccountType == beancountAssetsAccountType || account1.AccountType == beancountLiabilitiesAccountType) {
				fromAccount = account2
				toAccount = account1
				toCurrency = splitData1.Commodity
				toAmount = amount1
			}

			if fromAccount.isOpeningBalanceEquityAccount() {
				data[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TYPE] = utils.IntToString(int(models.TRANSACTION_TYPE_MODIFY_BALANCE))
			} else {
				data[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TYPE] = utils.IntToString(int(models.TRANSACTION_TYPE_INCOME))
			}

			data[datatable.TRANSACTION_DATA_TABLE_SUB_CATEGORY] = fromAccount.Name
			data[datatable.TRANSACTION_DATA_TABLE_ACCOUNT_NAME] = toAccount.Name
			data[datatable.TRANSACTION_DATA_TABLE_ACCOUNT_CURRENCY] = toCurrency
			data[datatable.TRANSACTION_DATA_TABLE_AMOUNT] = utils.FormatAmount(toAmount)
		} else if account1.AccountType == beancountExpensesAccountType && (account2.AccountType == beancountAssetsAccountType || account2.AccountType == beancountLiabilitiesAccountType) ||
			(account2.AccountType == beancountExpensesAccountType && (account1.AccountType == beancountAssetsAccountType || account1.AccountType == beancountLiabilitiesAccountType)) { 
			fromAccount := account1
			fromCurrency := splitData1.Commodity
			fromAmount := amount1
			toAccount := account2

			if account1.AccountType == beancountExpensesAccountType && (account2.AccountType == beancountAssetsAccountType || account2.AccountType == beancountLiabilitiesAccountType) {
				fromAccount = account2
				fromCurrency = splitData2.Commodity
				fromAmount = amount2
				toAccount = account1
			}

			data[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TYPE] = utils.IntToString(int(models.TRANSACTION_TYPE_EXPENSE))
			data[datatable.TRANSACTION_DATA_TABLE_SUB_CATEGORY] = toAccount.Name
			data[datatable.TRANSACTION_DATA_TABLE_ACCOUNT_NAME] = fromAccount.Name
			data[datatable.TRANSACTION_DATA_TABLE_ACCOUNT_CURRENCY] = fromCurrency
			data[datatable.TRANSACTION_DATA_TABLE_AMOUNT] = utils.FormatAmount(-fromAmount)
		} else if (account1.AccountType == beancountAssetsAccountType || account1.AccountType == beancountLiabilitiesAccountType) &&
			(account2.AccountType == beancountAssetsAccountType || account2.AccountType == beancountLiabilitiesAccountType) {
			var fromAccount, toAccount *beancountAccount
			var fromAmount, toAmount int64
			var fromCurrency, toCurrency string

			if amount1 < 0 {
				fromAccount = account1
				fromCurrency = splitData1.Commodity
				fromAmount = -amount1
				toAccount = account2
				toCurrency = splitData2.Commodity
				toAmount = amount2
			} else if amount2 < 0 {
				fromAccount = account2
				fromCurrency = splitData2.Commodity
				fromAmount = -amount2
				toAccount = account1
				toCurrency = splitData1.Commodity
				toAmount = amount1
			} else {
				log.Errorf(ctx, "[beancount_transaction_data_table.parseTransaction] cannot parse transfer transaction, because unexcepted account amounts \"%d\" and \"%d\"", amount1, amount2)
				return nil, errs.ErrInvalidBeancountFile
			}

			data[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TYPE] = utils.IntToString(int(models.TRANSACTION_TYPE_TRANSFER))
			data[datatable.TRANSACTION_DATA_TABLE_SUB_CATEGORY] = ""
			data[datatable.TRANSACTION_DATA_TABLE_ACCOUNT_NAME] = fromAccount.Name
			data[datatable.TRANSACTION_DATA_TABLE_ACCOUNT_CURRENCY] = fromCurrency
			data[datatable.TRANSACTION_DATA_TABLE_AMOUNT] = utils.FormatAmount(fromAmount)
			data[datatable.TRANSACTION_DATA_TABLE_RELATED_ACCOUNT_NAME] = toAccount.Name
			data[datatable.TRANSACTION_DATA_TABLE_RELATED_ACCOUNT_CURRENCY] = toCurrency
			data[datatable.TRANSACTION_DATA_TABLE_RELATED_AMOUNT] = utils.FormatAmount(toAmount)
		} else {
			log.Errorf(ctx, "[beancount_transaction_data_table.parseTransaction] cannot parse transaction, because unexcepted account types \"%d\" and \"%d\"", account1.AccountType, account2.AccountType)
			return nil, errs.ErrThereAreNotSupportedTransactionType
		}
	} else if len(beancountEntry.Postings) <= 1 {
		log.Errorf(ctx, "[beancount_transaction_data_table.parseTransaction] cannot parse transaction, because postings count is %d", len(beancountEntry.Postings))
		return nil, errs.ErrInvalidBeancountFile
	} else {
		log.Errorf(ctx, "[beancount_transaction_data_table.parseTransaction] cannot parse split transaction, because postings count is %d", len(beancountEntry.Postings))
		return nil, errs.ErrNotSupportedSplitTransactions
	}

	data[datatable.TRANSACTION_DATA_TABLE_TAGS] = strings.Join(beancountEntry.Tags, BEANCOUNT_TRANSACTION_TAG_SEPARATOR)
	data[datatable.TRANSACTION_DATA_TABLE_DESCRIPTION] = beancountEntry.Narration

	return data, nil
}

func createNewBeancountTransactionDataTable(beancountData *beancountData) (*beancountTransactionDataTable, error) {
	if beancountData == nil {
		return nil, errs.ErrNotFoundTransactionDataInFile
	}

	return &beancountTransactionDataTable{
		allData:    beancountData.Transactions,
		accountMap: beancountData.Accounts,
	}, nil
}
