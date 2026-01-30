package ofx

import (
	"time"

	"github.com/Shavitjnr/split-chill-ai/pkg/converters/converter"
	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/models"
	"github.com/Shavitjnr/split-chill-ai/pkg/utils"
)

var ofxTransactionTypeNameMapping = map[models.TransactionType]string{
	models.TRANSACTION_TYPE_INCOME:   utils.IntToString(int(models.TRANSACTION_TYPE_INCOME)),
	models.TRANSACTION_TYPE_EXPENSE:  utils.IntToString(int(models.TRANSACTION_TYPE_EXPENSE)),
	models.TRANSACTION_TYPE_TRANSFER: utils.IntToString(int(models.TRANSACTION_TYPE_TRANSFER)),
}


type ofxTransactionDataImporter struct {
}


var (
	OFXTransactionDataImporter = &ofxTransactionDataImporter{}
)


func (c *ofxTransactionDataImporter) ParseImportedData(ctx core.Context, user *models.User, data []byte, defaultTimezone *time.Location, additionalOptions converter.TransactionDataImporterOptions, accountMap map[string]*models.Account, expenseCategoryMap map[string]map[string]*models.TransactionCategory, incomeCategoryMap map[string]map[string]*models.TransactionCategory, transferCategoryMap map[string]map[string]*models.TransactionCategory, tagMap map[string]*models.TransactionTag) (models.ImportedTransactionSlice, []*models.Account, []*models.TransactionCategory, []*models.TransactionCategory, []*models.TransactionCategory, []*models.TransactionTag, error) {
	ofxDataReader, err := createNewOFXFileReader(ctx, data)

	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	ofxFile, err := ofxDataReader.read(ctx)

	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	transactionDataTable, err := createNewOFXTransactionDataTable(ofxFile)

	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	dataTableImporter := converter.CreateNewSimpleImporterWithTypeNameMapping(ofxTransactionTypeNameMapping)

	return dataTableImporter.ParseImportedData(ctx, user, transactionDataTable, defaultTimezone, additionalOptions, accountMap, expenseCategoryMap, incomeCategoryMap, transferCategoryMap, tagMap)
}
