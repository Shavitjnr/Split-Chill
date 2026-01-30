package feidee

import (
	"github.com/Shavitjnr/split-chill-ai/pkg/converters/datatable"
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
	"github.com/Shavitjnr/split-chill-ai/pkg/models"
	"github.com/Shavitjnr/split-chill-ai/pkg/utils"
)

var feideeMymoneyTransactionTypeNameMapping = map[models.TransactionType]string{
	models.TRANSACTION_TYPE_MODIFY_BALANCE: "余额变更",
	models.TRANSACTION_TYPE_INCOME:         "收入",
	models.TRANSACTION_TYPE_EXPENSE:        "支出",
	models.TRANSACTION_TYPE_TRANSFER:       "转账",
}

var feideeMymoneyTransactionTypeModifyOutstandingBalanceName = "负债变更"


type feideeMymoneyTransactionDataRowParser struct {
}


func (p *feideeMymoneyTransactionDataRowParser) GetAddedColumns() []datatable.TransactionDataTableColumn {
	return nil
}


func (p *feideeMymoneyTransactionDataRowParser) Parse(data map[datatable.TransactionDataTableColumn]string) (rowData map[datatable.TransactionDataTableColumn]string, rowDataValid bool, err error) {
	rowData = make(map[datatable.TransactionDataTableColumn]string, len(data))

	for column, value := range data {
		rowData[column] = value
	}

	if rowData[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TIME] != "" {
		rowData[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TIME] = p.getLongDateTime(rowData[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TIME])
	}

	if rowData[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TYPE] == feideeMymoneyTransactionTypeNameMapping[models.TRANSACTION_TYPE_MODIFY_BALANCE] {
		amount, err := utils.ParseAmount(rowData[datatable.TRANSACTION_DATA_TABLE_AMOUNT])

		if err != nil {
			return nil, false, errs.ErrAmountInvalid
		}

		
		if amount >= 0 {
			rowData[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TYPE] = feideeMymoneyTransactionTypeNameMapping[models.TRANSACTION_TYPE_INCOME]
		} else {
			rowData[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TYPE] = feideeMymoneyTransactionTypeNameMapping[models.TRANSACTION_TYPE_EXPENSE]
			rowData[datatable.TRANSACTION_DATA_TABLE_AMOUNT] = utils.FormatAmount(-amount)
		}
	} else if rowData[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TYPE] == feideeMymoneyTransactionTypeModifyOutstandingBalanceName {
		amount, err := utils.ParseAmount(rowData[datatable.TRANSACTION_DATA_TABLE_AMOUNT])

		if err != nil {
			return nil, false, errs.ErrAmountInvalid
		}

		
		if amount >= 0 {
			rowData[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TYPE] = feideeMymoneyTransactionTypeNameMapping[models.TRANSACTION_TYPE_EXPENSE]
		} else {
			rowData[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TYPE] = feideeMymoneyTransactionTypeNameMapping[models.TRANSACTION_TYPE_INCOME]
			rowData[datatable.TRANSACTION_DATA_TABLE_AMOUNT] = utils.FormatAmount(-amount)
		}
	}

	return rowData, true, nil
}


func (p *feideeMymoneyTransactionDataRowParser) getLongDateTime(str string) string {
	if utils.IsValidLongDateTimeFormat(str) {
		return str
	}

	if utils.IsValidLongDateTimeWithoutSecondFormat(str) {
		return str + ":00"
	}

	if utils.IsValidLongDateFormat(str) {
		return str + " 00:00:00"
	}

	return str
}


func createFeideeMymoneyTransactionDataRowParser() datatable.TransactionDataRowParser {
	return &feideeMymoneyTransactionDataRowParser{}
}
