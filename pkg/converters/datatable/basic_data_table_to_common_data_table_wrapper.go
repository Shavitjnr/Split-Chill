package datatable


type basicDataTableToCommonDataTableWrapper struct {
	innerDataTable    BasicDataTable
	dataColumnIndexes map[string]int
}


type basicDataTableToCommonDataTableWrapperRow struct {
	rowData map[string]string
}


type basicDataTableToCommonDataTableWrapperRowIterator struct {
	commonDataTable *basicDataTableToCommonDataTableWrapper
	innerIterator   BasicDataTableRowIterator
}


func (t *basicDataTableToCommonDataTableWrapper) HeaderColumnCount() int {
	return len(t.innerDataTable.HeaderColumnNames())
}


func (t *basicDataTableToCommonDataTableWrapper) HasColumn(columnName string) bool {
	index, exists := t.dataColumnIndexes[columnName]
	return exists && index >= 0
}


func (t *basicDataTableToCommonDataTableWrapper) DataRowCount() int {
	return t.innerDataTable.DataRowCount()
}


func (t *basicDataTableToCommonDataTableWrapper) DataRowIterator() CommonDataTableRowIterator {
	return &basicDataTableToCommonDataTableWrapperRowIterator{
		commonDataTable: t,
		innerIterator:   t.innerDataTable.DataRowIterator(),
	}
}


func (r *basicDataTableToCommonDataTableWrapperRow) HasData(columnName string) bool {
	_, exists := r.rowData[columnName]
	return exists
}


func (r *basicDataTableToCommonDataTableWrapperRow) ColumnCount() int {
	return len(r.rowData)
}


func (r *basicDataTableToCommonDataTableWrapperRow) GetData(columnName string) string {
	return r.rowData[columnName]
}


func (t *basicDataTableToCommonDataTableWrapperRowIterator) HasNext() bool {
	return t.innerIterator.HasNext()
}


func (t *basicDataTableToCommonDataTableWrapperRowIterator) CurrentRowId() string {
	return t.innerIterator.CurrentRowId()
}


func (t *basicDataTableToCommonDataTableWrapperRowIterator) Next() CommonDataTableRow {
	basicDataRow := t.innerIterator.Next()

	if basicDataRow == nil {
		return nil
	}

	rowData := make(map[string]string, len(t.commonDataTable.dataColumnIndexes))

	for column, columnIndex := range t.commonDataTable.dataColumnIndexes {
		if columnIndex < 0 || columnIndex >= basicDataRow.ColumnCount() {
			continue
		}

		value := basicDataRow.GetData(columnIndex)
		rowData[column] = value
	}

	return &basicDataTableToCommonDataTableWrapperRow{
		rowData: rowData,
	}
}


func CreateNewCommonDataTableFromBasicDataTable(dataTable BasicDataTable) CommonDataTable {
	headerLineItems := dataTable.HeaderColumnNames()
	dataColumnIndexes := make(map[string]int, len(headerLineItems))

	for i := 0; i < len(headerLineItems); i++ {
		dataColumnIndexes[headerLineItems[i]] = i
	}

	return &basicDataTableToCommonDataTableWrapper{
		innerDataTable:    dataTable,
		dataColumnIndexes: dataColumnIndexes,
	}
}
