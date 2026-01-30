package fireflyIII

import (
	"github.com/Shavitjnr/split-chill-ai/pkg/converters/datatable"
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
	"github.com/Shavitjnr/split-chill-ai/pkg/models"
	"github.com/Shavitjnr/split-chill-ai/pkg/utils"
)


type fireflyIIITransactionDataRowParser struct {
}


func (p *fireflyIIITransactionDataRowParser) GetAddedColumns() []datatable.TransactionDataTableColumn {
	return []datatable.TransactionDataTableColumn{
		datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TIMEZONE,
	}
}


func (p *fireflyIIITransactionDataRowParser) Parse(data map[datatable.TransactionDataTableColumn]string) (rowData map[datatable.TransactionDataTableColumn]string, rowDataValid bool, err error) {
	rowData = make(map[datatable.TransactionDataTableColumn]string, len(data))

	for column, value := range data {
		rowData[column] = value
	}

	
	if rowData[datatable.TRANSACTION_DATA_TABLE_SUB_CATEGORY] == "" {
		if rowData[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TYPE] == fireflyIIITransactionTypeNameMapping[models.TRANSACTION_TYPE_INCOME] {
			rowData[datatable.TRANSACTION_DATA_TABLE_SUB_CATEGORY] = rowData[datatable.TRANSACTION_DATA_TABLE_ACCOUNT_NAME]
		} else if rowData[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TYPE] == fireflyIIITransactionTypeNameMapping[models.TRANSACTION_TYPE_EXPENSE] {
			rowData[datatable.TRANSACTION_DATA_TABLE_SUB_CATEGORY] = rowData[datatable.TRANSACTION_DATA_TABLE_RELATED_ACCOUNT_NAME]
		}
	}

	
	if rowData[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TIME] != "" {
		dateTime, err := utils.ParseFromLongDateTimeWithTimezoneRFC3339Format(rowData[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TIME])

		if err != nil {
			return nil, false, errs.ErrTransactionTimeInvalid
		}

		rowData[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TIME] = utils.FormatUnixTimeToLongDateTime(dateTime.Unix(), dateTime.Location())
		rowData[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TIMEZONE] = utils.FormatTimezoneOffset(dateTime.Unix(), dateTime.Location())
	}

	
	if rowData[datatable.TRANSACTION_DATA_TABLE_AMOUNT] != "" {
		rowData[datatable.TRANSACTION_DATA_TABLE_AMOUNT] = utils.TrimTrailingZerosInDecimal(rowData[datatable.TRANSACTION_DATA_TABLE_AMOUNT])
		amount, err := utils.ParseAmount(rowData[datatable.TRANSACTION_DATA_TABLE_AMOUNT])

		if err != nil {
			return nil, false, errs.ErrAmountInvalid
		}

		if rowData[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TYPE] == fireflyIIITransactionTypeNameMapping[models.TRANSACTION_TYPE_EXPENSE] {
			rowData[datatable.TRANSACTION_DATA_TABLE_AMOUNT] = utils.FormatAmount(-amount)
		} else {
			rowData[datatable.TRANSACTION_DATA_TABLE_AMOUNT] = utils.FormatAmount(amount)
		}
	}

	if rowData[datatable.TRANSACTION_DATA_TABLE_RELATED_AMOUNT] != "" {
		rowData[datatable.TRANSACTION_DATA_TABLE_RELATED_AMOUNT] = utils.TrimTrailingZerosInDecimal(rowData[datatable.TRANSACTION_DATA_TABLE_RELATED_AMOUNT])
		amount, err := utils.ParseAmount(rowData[datatable.TRANSACTION_DATA_TABLE_RELATED_AMOUNT])

		if err != nil {
			return nil, false, errs.ErrAmountInvalid
		}

		if rowData[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TYPE] == fireflyIIITransactionTypeNameMapping[models.TRANSACTION_TYPE_EXPENSE] {
			rowData[datatable.TRANSACTION_DATA_TABLE_RELATED_AMOUNT] = utils.FormatAmount(-amount)
		} else {
			rowData[datatable.TRANSACTION_DATA_TABLE_RELATED_AMOUNT] = utils.FormatAmount(amount)
		}
	} else {
		rowData[datatable.TRANSACTION_DATA_TABLE_RELATED_AMOUNT] = rowData[datatable.TRANSACTION_DATA_TABLE_AMOUNT]
	}

	
	if rowData[datatable.TRANSACTION_DATA_TABLE_RELATED_ACCOUNT_CURRENCY] == "" {
		rowData[datatable.TRANSACTION_DATA_TABLE_RELATED_ACCOUNT_CURRENCY] = rowData[datatable.TRANSACTION_DATA_TABLE_ACCOUNT_CURRENCY]
	}

	
	if rowData[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TYPE] == fireflyIIITransactionTypeNameMapping[models.TRANSACTION_TYPE_MODIFY_BALANCE] {
		rowData[datatable.TRANSACTION_DATA_TABLE_ACCOUNT_NAME] = rowData[datatable.TRANSACTION_DATA_TABLE_RELATED_ACCOUNT_NAME]
	}

	
	if rowData[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TYPE] == fireflyIIITransactionTypeNameMapping[models.TRANSACTION_TYPE_INCOME] {
		rowData[datatable.TRANSACTION_DATA_TABLE_ACCOUNT_NAME] = rowData[datatable.TRANSACTION_DATA_TABLE_RELATED_ACCOUNT_NAME]
	}

	return rowData, true, nil
}


func createFireflyIIITransactionDataRowParser() datatable.TransactionDataRowParser {
	return &fireflyIIITransactionDataRowParser{}
}
