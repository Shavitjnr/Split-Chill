package feidee

import (
	"time"

	"github.com/Shavitjnr/split-chill-ai/pkg/converters/converter"
	"github.com/Shavitjnr/split-chill-ai/pkg/converters/datatable"
	"github.com/Shavitjnr/split-chill-ai/pkg/converters/excel"
	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/models"
)

var feideeMymoneyElecloudDataColumnNameMapping = map[datatable.TransactionDataTableColumn]string{
	datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TIME:     "日期",
	datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TYPE:     "交易类型",
	datatable.TRANSACTION_DATA_TABLE_CATEGORY:             "分类",
	datatable.TRANSACTION_DATA_TABLE_SUB_CATEGORY:         "子分类",
	datatable.TRANSACTION_DATA_TABLE_ACCOUNT_NAME:         "账户1",
	datatable.TRANSACTION_DATA_TABLE_ACCOUNT_CURRENCY:     "账户币种",
	datatable.TRANSACTION_DATA_TABLE_AMOUNT:               "金额",
	datatable.TRANSACTION_DATA_TABLE_RELATED_ACCOUNT_NAME: "账户2",
	datatable.TRANSACTION_DATA_TABLE_DESCRIPTION:          "备注",
	datatable.TRANSACTION_DATA_TABLE_MEMBER:               "成员",
	datatable.TRANSACTION_DATA_TABLE_PROJECT:              "项目",
	datatable.TRANSACTION_DATA_TABLE_MERCHANT:             "商家",
}


type feideeMymoneyElecloudTransactionDataXlsxFileImporter struct {
	converter.DataTableTransactionDataImporter
}


var (
	FeideeMymoneyElecloudTransactionDataXlsxFileImporter = &feideeMymoneyElecloudTransactionDataXlsxFileImporter{}
)


func (c *feideeMymoneyElecloudTransactionDataXlsxFileImporter) ParseImportedData(ctx core.Context, user *models.User, data []byte, defaultTimezone *time.Location, additionalOptions converter.TransactionDataImporterOptions, accountMap map[string]*models.Account, expenseCategoryMap map[string]map[string]*models.TransactionCategory, incomeCategoryMap map[string]map[string]*models.TransactionCategory, transferCategoryMap map[string]map[string]*models.TransactionCategory, tagMap map[string]*models.TransactionTag) (models.ImportedTransactionSlice, []*models.Account, []*models.TransactionCategory, []*models.TransactionCategory, []*models.TransactionCategory, []*models.TransactionTag, error) {
	dataTable, err := excel.CreateNewExcelOOXMLFileBasicDataTable(data, true)

	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	transactionRowParser := createFeideeMymoneyElecloudTransactionDataRowParser()
	transactionDataTable := datatable.CreateNewTransactionDataTableFromBasicDataTableWithRowParser(dataTable, feideeMymoneyElecloudDataColumnNameMapping, transactionRowParser)
	dataTableImporter := converter.CreateNewSimpleImporter(feideeMymoneyElecloudTransactionTypeNameMapping)

	return dataTableImporter.ParseImportedData(ctx, user, transactionDataTable, defaultTimezone, additionalOptions, accountMap, expenseCategoryMap, incomeCategoryMap, transferCategoryMap, tagMap)
}
