package datatable


type CommonDataTable interface {
	
	HeaderColumnCount() int

	
	HasColumn(columnName string) bool

	
	DataRowCount() int

	
	DataRowIterator() CommonDataTableRowIterator
}


type CommonDataTableRow interface {
	
	ColumnCount() int

	
	HasData(columnName string) bool

	
	GetData(columnName string) string
}


type CommonDataTableRowIterator interface {
	
	HasNext() bool

	
	CurrentRowId() string

	
	Next() CommonDataTableRow
}
