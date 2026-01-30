package datatable


type BasicDataTable interface {
	
	DataRowCount() int

	
	HeaderColumnNames() []string

	
	DataRowIterator() BasicDataTableRowIterator
}


type BasicDataTableRow interface {
	
	ColumnCount() int

	
	GetData(columnIndex int) string
}


type BasicDataTableRowIterator interface {
	
	HasNext() bool

	
	CurrentRowId() string

	
	Next() BasicDataTableRow
}
