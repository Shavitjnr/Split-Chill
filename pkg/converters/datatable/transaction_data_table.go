package datatable

import (
	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/models"
)


type TransactionDataTable interface {
	
	HasColumn(column TransactionDataTableColumn) bool

	
	TransactionRowCount() int

	
	TransactionRowIterator() TransactionDataRowIterator
}


type TransactionDataRow interface {
	
	IsValid() bool

	
	GetData(column TransactionDataTableColumn) string
}


type TransactionDataRowIterator interface {
	
	HasNext() bool

	
	Next(ctx core.Context, user *models.User) (daraRow TransactionDataRow, err error)
}


type TransactionDataRowParser interface {
	
	GetAddedColumns() []TransactionDataTableColumn

	
	Parse(data map[TransactionDataTableColumn]string) (rowData map[TransactionDataTableColumn]string, rowDataValid bool, err error)
}


type TransactionDataTableBuilder interface {
	
	AppendTransaction(data map[TransactionDataTableColumn]string)

	
	ReplaceDelimiters(text string) string
}


type TransactionDataTableColumn byte


const (
	TRANSACTION_DATA_TABLE_TRANSACTION_TIME         TransactionDataTableColumn = 1
	TRANSACTION_DATA_TABLE_TRANSACTION_TIMEZONE     TransactionDataTableColumn = 2
	TRANSACTION_DATA_TABLE_TRANSACTION_TYPE         TransactionDataTableColumn = 3
	TRANSACTION_DATA_TABLE_CATEGORY                 TransactionDataTableColumn = 4
	TRANSACTION_DATA_TABLE_SUB_CATEGORY             TransactionDataTableColumn = 5
	TRANSACTION_DATA_TABLE_ACCOUNT_NAME             TransactionDataTableColumn = 6
	TRANSACTION_DATA_TABLE_ACCOUNT_CURRENCY         TransactionDataTableColumn = 7
	TRANSACTION_DATA_TABLE_AMOUNT                   TransactionDataTableColumn = 8
	TRANSACTION_DATA_TABLE_RELATED_ACCOUNT_NAME     TransactionDataTableColumn = 9
	TRANSACTION_DATA_TABLE_RELATED_ACCOUNT_CURRENCY TransactionDataTableColumn = 10
	TRANSACTION_DATA_TABLE_RELATED_AMOUNT           TransactionDataTableColumn = 11
	TRANSACTION_DATA_TABLE_GEOGRAPHIC_LOCATION      TransactionDataTableColumn = 12
	TRANSACTION_DATA_TABLE_TAGS                     TransactionDataTableColumn = 13
	TRANSACTION_DATA_TABLE_DESCRIPTION              TransactionDataTableColumn = 14
	TRANSACTION_DATA_TABLE_PAYEE                    TransactionDataTableColumn = 101
	TRANSACTION_DATA_TABLE_MEMBER                   TransactionDataTableColumn = 102
	TRANSACTION_DATA_TABLE_PROJECT                  TransactionDataTableColumn = 103
	TRANSACTION_DATA_TABLE_MERCHANT                 TransactionDataTableColumn = 104
)


const TRANSACTION_DATA_TABLE_TIMEZONE_NOT_AVAILABLE = "TIMEZONE_NOT_AVAILABLE"
