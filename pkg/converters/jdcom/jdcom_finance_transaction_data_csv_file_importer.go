package jdcom

import (
	"bytes"
	"time"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"

	"github.com/Shavitjnr/split-chill-ai/pkg/converters/converter"
	"github.com/Shavitjnr/split-chill-ai/pkg/converters/csv"
	"github.com/Shavitjnr/split-chill-ai/pkg/converters/datatable"
	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
	"github.com/Shavitjnr/split-chill-ai/pkg/log"
	"github.com/Shavitjnr/split-chill-ai/pkg/models"
)


type jdComFinanceTransactionDataCsvFileImporter struct {
	fileHeaderLineBeginning         string
	dataHeaderStartContentBeginning string
}


var (
	JDComFinanceTransactionDataCsvFileImporter = &jdComFinanceTransactionDataCsvFileImporter{}
)


func (c *jdComFinanceTransactionDataCsvFileImporter) ParseImportedData(ctx core.Context, user *models.User, data []byte, defaultTimezone *time.Location, additionalOptions converter.TransactionDataImporterOptions, accountMap map[string]*models.Account, expenseCategoryMap map[string]map[string]*models.TransactionCategory, incomeCategoryMap map[string]map[string]*models.TransactionCategory, transferCategoryMap map[string]map[string]*models.TransactionCategory, tagMap map[string]*models.TransactionTag) (models.ImportedTransactionSlice, []*models.Account, []*models.TransactionCategory, []*models.TransactionCategory, []*models.TransactionCategory, []*models.TransactionTag, error) {
	fallback := unicode.UTF8.NewDecoder()
	reader := transform.NewReader(bytes.NewReader(data), unicode.BOMOverride(fallback))

	csvDataTable, err := csv.CreateNewCsvBasicDataTable(ctx, reader, false)

	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	dataTable, err := createNewJDComFinanceTransactionBasicDataTable(ctx, csvDataTable)

	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	commonDataTable := datatable.CreateNewCommonDataTableFromBasicDataTable(dataTable)

	if !commonDataTable.HasColumn(jdComFinanceTransactionTimeColumnName) ||
		!commonDataTable.HasColumn(jdComFinanceTransactionMerchantNameColumnName) ||
		!commonDataTable.HasColumn(jdComFinanceTransactionMemoColumnName) ||
		!commonDataTable.HasColumn(jdComFinanceTransactionAmountColumnName) ||
		!commonDataTable.HasColumn(jdComFinanceTransactionRelatedAccountColumnName) ||
		!commonDataTable.HasColumn(jdComFinanceTransactionStatusColumnName) ||
		!commonDataTable.HasColumn(jdComFinanceTransactionTypeColumnName) {
		log.Errorf(ctx, "[jdcom_finance_transaction_data_csv_file_importer.ParseImportedData] cannot parse jd.com finance csv data, because missing essential columns in header row")
		return nil, nil, nil, nil, nil, nil, errs.ErrMissingRequiredFieldInHeaderRow
	}

	transactionRowParser := createJDComFinanceTransactionDataRowParser(dataTable.HeaderColumnNames())
	transactionDataTable := datatable.CreateNewTransactionDataTableFromCommonDataTable(commonDataTable, jdComFinanceTransactionSupportedColumns, transactionRowParser)
	dataTableImporter := converter.CreateNewSimpleImporterWithTypeNameMapping(jdComFinanceTransactionTypeNameMapping)

	return dataTableImporter.ParseImportedData(ctx, user, transactionDataTable, defaultTimezone, additionalOptions, accountMap, expenseCategoryMap, incomeCategoryMap, transferCategoryMap, tagMap)
}
