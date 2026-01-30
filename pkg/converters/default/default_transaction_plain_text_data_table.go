package _default

import (
	"fmt"
	"strings"

	"github.com/Shavitjnr/split-chill-ai/pkg/converters/datatable"
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
)


type defaultPlainTextDataTable struct {
	columnSeparator       string
	lineSeparator         string
	allLines              []string
	headerLineColumnNames []string
}


type defaultPlainTextDataRow struct {
	allItems []string
}


type defaultPlainTextDataRowIterator struct {
	dataTable    *defaultPlainTextDataTable
	currentIndex int
}


type defaultTransactionPlainTextDataTableBuilder struct {
	columnSeparator       string
	lineSeparator         string
	columns               []datatable.TransactionDataTableColumn
	dataColumnNameMapping map[datatable.TransactionDataTableColumn]string
	dataLineFormat        string
	builder               *strings.Builder
}


func (t *defaultPlainTextDataTable) DataRowCount() int {
	if len(t.allLines) < 1 {
		return 0
	}

	return len(t.allLines) - 1
}


func (t *defaultPlainTextDataTable) HeaderColumnNames() []string {
	return t.headerLineColumnNames
}


func (t *defaultPlainTextDataTable) DataRowIterator() datatable.BasicDataTableRowIterator {
	return &defaultPlainTextDataRowIterator{
		dataTable:    t,
		currentIndex: 0,
	}
}


func (r *defaultPlainTextDataRow) ColumnCount() int {
	return len(r.allItems)
}


func (r *defaultPlainTextDataRow) GetData(columnIndex int) string {
	if columnIndex >= len(r.allItems) {
		return ""
	}

	return r.allItems[columnIndex]
}


func (t *defaultPlainTextDataRowIterator) HasNext() bool {
	return t.currentIndex+1 < len(t.dataTable.allLines)
}


func (t *defaultPlainTextDataRowIterator) CurrentRowId() string {
	return fmt.Sprintf("line#%d", t.currentIndex)
}


func (t *defaultPlainTextDataRowIterator) Next() datatable.BasicDataTableRow {
	if t.currentIndex+1 >= len(t.dataTable.allLines) {
		return nil
	}

	t.currentIndex++

	rowContent := t.dataTable.allLines[t.currentIndex]
	rowItems := strings.Split(rowContent, t.dataTable.columnSeparator)

	return &defaultPlainTextDataRow{
		allItems: rowItems,
	}
}


func (b *defaultTransactionPlainTextDataTableBuilder) AppendTransaction(data map[datatable.TransactionDataTableColumn]string) {
	dataRowParams := make([]any, len(b.columns))

	for i := 0; i < len(b.columns); i++ {
		dataRowParams[i] = data[b.columns[i]]
	}

	b.builder.WriteString(fmt.Sprintf(b.dataLineFormat, dataRowParams...))
}


func (b *defaultTransactionPlainTextDataTableBuilder) ReplaceDelimiters(text string) string {
	text = strings.Replace(text, "\r\n", " ", -1)
	text = strings.Replace(text, "\r", " ", -1)
	text = strings.Replace(text, "\n", " ", -1)
	text = strings.Replace(text, b.columnSeparator, " ", -1)
	text = strings.Replace(text, b.lineSeparator, " ", -1)

	return text
}


func (b *defaultTransactionPlainTextDataTableBuilder) String() string {
	return b.builder.String()
}

func (b *defaultTransactionPlainTextDataTableBuilder) generateHeaderLine() string {
	var ret strings.Builder

	for i := 0; i < len(b.columns); i++ {
		if ret.Len() > 0 {
			ret.WriteString(b.columnSeparator)
		}

		dataColumn := b.columns[i]
		columnName := b.dataColumnNameMapping[dataColumn]

		ret.WriteString(columnName)
	}

	ret.WriteString(b.lineSeparator)

	return ret.String()
}

func (b *defaultTransactionPlainTextDataTableBuilder) generateDataLineFormat() string {
	var ret strings.Builder

	for i := 0; i < len(b.columns); i++ {
		if ret.Len() > 0 {
			ret.WriteString(b.columnSeparator)
		}

		ret.WriteString("%s")
	}

	ret.WriteString(b.lineSeparator)

	return ret.String()
}

func createNewDefaultPlainTextDataTable(content string, columnSeparator string, lineSeparator string) (*defaultPlainTextDataTable, error) {
	allLines := strings.Split(content, lineSeparator)

	if len(allLines) < 2 {
		return nil, errs.ErrNotFoundTransactionDataInFile
	}

	headerLine := allLines[0]
	headerLine = strings.ReplaceAll(headerLine, "\r", "")
	headerLineItems := strings.Split(headerLine, columnSeparator)

	return &defaultPlainTextDataTable{
		columnSeparator:       columnSeparator,
		lineSeparator:         lineSeparator,
		allLines:              allLines,
		headerLineColumnNames: headerLineItems,
	}, nil
}

func createNewDefaultTransactionPlainTextDataTableBuilder(transactionCount int, columns []datatable.TransactionDataTableColumn, dataColumnNameMapping map[datatable.TransactionDataTableColumn]string, columnSeparator string, lineSeparator string) *defaultTransactionPlainTextDataTableBuilder {
	var builder strings.Builder
	builder.Grow(transactionCount * 100)

	dataTableBuilder := &defaultTransactionPlainTextDataTableBuilder{
		columnSeparator:       columnSeparator,
		lineSeparator:         lineSeparator,
		columns:               columns,
		dataColumnNameMapping: dataColumnNameMapping,
		builder:               &builder,
	}

	headerLine := dataTableBuilder.generateHeaderLine()
	dataLineFormat := dataTableBuilder.generateDataLineFormat()

	dataTableBuilder.builder.WriteString(headerLine)
	dataTableBuilder.dataLineFormat = dataLineFormat

	return dataTableBuilder
}
