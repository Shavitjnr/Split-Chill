package datatable

import (
	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/log"
	"github.com/Shavitjnr/split-chill-ai/pkg/models"
)


type CommonTransactionDataRowParser interface {
	
	Parse(ctx core.Context, user *models.User, dataRow CommonDataTableRow, rowId string) (rowData map[TransactionDataTableColumn]string, rowDataValid bool, err error)
}


type commonDataTableToTransactionDataTableWrapper struct {
	innerDataTable       CommonDataTable
	supportedDataColumns map[TransactionDataTableColumn]bool
	rowParser            CommonTransactionDataRowParser
}


type commonDataTableToTransactionDataTableWrapperRow struct {
	transactionDataTable *commonDataTableToTransactionDataTableWrapper
	rowData              map[TransactionDataTableColumn]string
	rowDataValid         bool
}


type commonDataTableToTransactionDataTableWrapperRowIterator struct {
	transactionDataTable *commonDataTableToTransactionDataTableWrapper
	innerIterator        CommonDataTableRowIterator
}


func (t *commonDataTableToTransactionDataTableWrapper) HasColumn(column TransactionDataTableColumn) bool {
	_, exists := t.supportedDataColumns[column]
	return exists
}


func (t *commonDataTableToTransactionDataTableWrapper) TransactionRowCount() int {
	return t.innerDataTable.DataRowCount()
}


func (t *commonDataTableToTransactionDataTableWrapper) TransactionRowIterator() TransactionDataRowIterator {
	return &commonDataTableToTransactionDataTableWrapperRowIterator{
		transactionDataTable: t,
		innerIterator:        t.innerDataTable.DataRowIterator(),
	}
}


func (r *commonDataTableToTransactionDataTableWrapperRow) IsValid() bool {
	return r.rowDataValid
}


func (r *commonDataTableToTransactionDataTableWrapperRow) GetData(column TransactionDataTableColumn) string {
	if !r.rowDataValid {
		return ""
	}

	_, exists := r.transactionDataTable.supportedDataColumns[column]

	if !exists {
		return ""
	}

	return r.rowData[column]
}


func (t *commonDataTableToTransactionDataTableWrapperRowIterator) HasNext() bool {
	return t.innerIterator.HasNext()
}


func (t *commonDataTableToTransactionDataTableWrapperRowIterator) Next(ctx core.Context, user *models.User) (daraRow TransactionDataRow, err error) {
	commonDataRow := t.innerIterator.Next()

	if commonDataRow == nil {
		return nil, nil
	}

	rowId := t.innerIterator.CurrentRowId()
	rowData, rowDataValid, err := t.transactionDataTable.rowParser.Parse(ctx, user, commonDataRow, rowId)

	if err != nil {
		log.Errorf(ctx, "[common_data_table_to_transaction_data_table_wrapper.Next] cannot parse data row, because %s", err.Error())
		return nil, err
	}

	return &commonDataTableToTransactionDataTableWrapperRow{
		transactionDataTable: t.transactionDataTable,
		rowData:              rowData,
		rowDataValid:         rowDataValid,
	}, nil
}


func CreateNewTransactionDataTableFromCommonDataTable(dataTable CommonDataTable, supportedDataColumns map[TransactionDataTableColumn]bool, rowParser CommonTransactionDataRowParser) TransactionDataTable {
	return &commonDataTableToTransactionDataTableWrapper{
		innerDataTable:       dataTable,
		supportedDataColumns: supportedDataColumns,
		rowParser:            rowParser,
	}
}
