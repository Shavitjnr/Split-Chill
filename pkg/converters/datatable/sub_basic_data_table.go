package datatable


type SubBasicDataTable struct {
	baseTable BasicDataTable
	fromIndex int
	toIndex   int
}


type SubBasicDataTableRowIterator struct {
	dataTable     *SubBasicDataTable
	innerIterator BasicDataTableRowIterator
	currentIndex  int
}


func (t *SubBasicDataTable) DataRowCount() int {
	return t.toIndex - t.fromIndex
}


func (t *SubBasicDataTable) HeaderColumnNames() []string {
	return t.baseTable.HeaderColumnNames()
}


func (t *SubBasicDataTable) DataRowIterator() BasicDataTableRowIterator {
	innerIterator := t.baseTable.DataRowIterator()
	currentIndex := -1

	
	for currentIndex = -1; currentIndex < t.fromIndex-1 && innerIterator.HasNext(); currentIndex++ {
		innerIterator.Next()
	}

	return &SubBasicDataTableRowIterator{
		dataTable:     t,
		innerIterator: innerIterator,
		currentIndex:  currentIndex,
	}
}


func (t *SubBasicDataTableRowIterator) HasNext() bool {
	return t.currentIndex+1 < t.dataTable.toIndex && t.innerIterator.HasNext()
}


func (t *SubBasicDataTableRowIterator) CurrentRowId() string {
	return t.innerIterator.CurrentRowId()
}


func (t *SubBasicDataTableRowIterator) Next() BasicDataTableRow {
	if t.currentIndex+1 >= t.dataTable.toIndex {
		return nil
	}

	t.currentIndex++
	return t.innerIterator.Next()
}


func CreateSubBasicTable(dataTable BasicDataTable, fromIndex, toIndex int) *SubBasicDataTable {
	if fromIndex < 0 {
		fromIndex = 0
	}

	if fromIndex > dataTable.DataRowCount() {
		fromIndex = dataTable.DataRowCount()
	}

	if toIndex > dataTable.DataRowCount() {
		toIndex = dataTable.DataRowCount()
	}

	if toIndex < fromIndex {
		toIndex = fromIndex
	}

	return &SubBasicDataTable{
		baseTable: dataTable,
		fromIndex: fromIndex,
		toIndex:   toIndex,
	}
}
