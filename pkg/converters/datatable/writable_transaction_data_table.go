package datatable

import (
	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/log"
	"github.com/Shavitjnr/split-chill-ai/pkg/models"
)


type WritableTransactionDataTable struct {
	allData          []map[TransactionDataTableColumn]string
	supportedColumns map[TransactionDataTableColumn]bool
	rowParser        TransactionDataRowParser
	addedColumns     map[TransactionDataTableColumn]bool
}


type WritableTransactionDataRow struct {
	dataTable    *WritableTransactionDataTable
	rowData      map[TransactionDataTableColumn]string
	rowDataValid bool
}


type WritableTransactionDataRowIterator struct {
	dataTable *WritableTransactionDataTable
	nextIndex int
}


func (t *WritableTransactionDataTable) Add(data map[TransactionDataTableColumn]string) {
	finalData := make(map[TransactionDataTableColumn]string, len(data))

	for column, value := range data {
		_, exists := t.supportedColumns[column]

		if exists {
			finalData[column] = value
		}
	}

	t.allData = append(t.allData, finalData)
}


func (t *WritableTransactionDataTable) Get(index int) (*WritableTransactionDataRow, error) {
	if index >= len(t.allData) {
		return nil, nil
	}

	rowData := t.allData[index]
	rowDataValid := true

	if t.rowParser != nil {
		var err error
		rowData, rowDataValid, err = t.rowParser.Parse(rowData)

		if err != nil {
			return nil, err
		}
	}

	return &WritableTransactionDataRow{
		dataTable:    t,
		rowData:      rowData,
		rowDataValid: rowDataValid,
	}, nil
}


func (t *WritableTransactionDataTable) HasColumn(column TransactionDataTableColumn) bool {
	_, exists := t.supportedColumns[column]

	if exists {
		return exists
	}

	if t.addedColumns != nil {
		_, exists = t.addedColumns[column]

		if exists {
			return exists
		}
	}

	return false
}


func (t *WritableTransactionDataTable) TransactionRowCount() int {
	return len(t.allData)
}


func (t *WritableTransactionDataTable) TransactionRowIterator() TransactionDataRowIterator {
	return &WritableTransactionDataRowIterator{
		dataTable: t,
		nextIndex: 0,
	}
}


func (r *WritableTransactionDataRow) ColumnCount() int {
	if !r.rowDataValid {
		return 0
	}

	columnCount := 0

	for column := range r.rowData {
		if r.dataTable.supportedColumns[column] || r.dataTable.addedColumns[column] {
			columnCount++
		}
	}

	return columnCount
}


func (r *WritableTransactionDataRow) IsValid() bool {
	return r.rowDataValid
}


func (r *WritableTransactionDataRow) GetData(column TransactionDataTableColumn) string {
	if !r.rowDataValid {
		return ""
	}

	_, exists := r.dataTable.supportedColumns[column]

	if exists {
		return r.rowData[column]
	}

	if r.dataTable.addedColumns != nil {
		_, exists = r.dataTable.addedColumns[column]

		if exists {
			return r.rowData[column]
		}
	}

	return ""
}


func (t *WritableTransactionDataRowIterator) HasNext() bool {
	return t.nextIndex < len(t.dataTable.allData)
}


func (t *WritableTransactionDataRowIterator) Next(ctx core.Context, user *models.User) (daraRow TransactionDataRow, err error) {
	if t.nextIndex >= len(t.dataTable.allData) {
		return nil, nil
	}

	rowData := t.dataTable.allData[t.nextIndex]
	rowDataValid := true

	if t.dataTable.rowParser != nil {
		rowData, rowDataValid, err = t.dataTable.rowParser.Parse(rowData)

		if err != nil {
			log.Errorf(ctx, "[writable_transaction_data_table.Next] cannot parse data row, because %s", err.Error())
			return nil, err
		}
	}

	t.nextIndex++

	return &WritableTransactionDataRow{
		dataTable:    t.dataTable,
		rowData:      rowData,
		rowDataValid: rowDataValid,
	}, nil
}


func CreateNewWritableTransactionDataTable(columns []TransactionDataTableColumn) *WritableTransactionDataTable {
	return CreateNewWritableTransactionDataTableWithRowParser(columns, nil)
}


func CreateNewWritableTransactionDataTableWithRowParser(columns []TransactionDataTableColumn, rowParser TransactionDataRowParser) *WritableTransactionDataTable {
	supportedColumns := make(map[TransactionDataTableColumn]bool, len(columns))

	for i := 0; i < len(columns); i++ {
		column := columns[i]
		supportedColumns[column] = true
	}

	var addedColumns map[TransactionDataTableColumn]bool

	if rowParser != nil {
		addedColumnsByParser := rowParser.GetAddedColumns()
		addedColumns = make(map[TransactionDataTableColumn]bool, len(addedColumnsByParser))

		for i := 0; i < len(addedColumnsByParser); i++ {
			addedColumns[addedColumnsByParser[i]] = true
		}
	}

	return &WritableTransactionDataTable{
		allData:          make([]map[TransactionDataTableColumn]string, 0),
		supportedColumns: supportedColumns,
		rowParser:        rowParser,
		addedColumns:     addedColumns,
	}
}
