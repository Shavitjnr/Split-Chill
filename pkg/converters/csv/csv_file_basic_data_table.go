package csv

import (
	"encoding/csv"
	"fmt"
	"io"

	"github.com/Shavitjnr/split-chill-ai/pkg/converters/datatable"
	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
	"github.com/Shavitjnr/split-chill-ai/pkg/log"
)


type CsvFileBasicDataTable struct {
	allLines     [][]string
	hasTitleLine bool
}


type CsvFileBasicDataTableRow struct {
	dataTable *CsvFileBasicDataTable
	allItems  []string
}


type CsvFileBasicDataTableRowIterator struct {
	dataTable    *CsvFileBasicDataTable
	currentIndex int
}


func (t *CsvFileBasicDataTable) DataRowCount() int {
	if len(t.allLines) < 1 {
		return 0
	}

	if t.hasTitleLine {
		return len(t.allLines) - 1
	} else {
		return len(t.allLines)
	}
}


func (t *CsvFileBasicDataTable) HeaderColumnNames() []string {
	if len(t.allLines) < 1 {
		return nil
	}

	if t.hasTitleLine {
		return t.allLines[0]
	} else {
		return nil
	}
}


func (t *CsvFileBasicDataTable) DataRowIterator() datatable.BasicDataTableRowIterator {
	startIndex := -1

	if t.hasTitleLine {
		startIndex = 0
	}

	return &CsvFileBasicDataTableRowIterator{
		dataTable:    t,
		currentIndex: startIndex,
	}
}


func (r *CsvFileBasicDataTableRow) ColumnCount() int {
	return len(r.allItems)
}


func (r *CsvFileBasicDataTableRow) GetData(columnIndex int) string {
	if columnIndex >= len(r.allItems) {
		return ""
	}

	return r.allItems[columnIndex]
}


func (t *CsvFileBasicDataTableRowIterator) HasNext() bool {
	return t.currentIndex+1 < len(t.dataTable.allLines)
}


func (t *CsvFileBasicDataTableRowIterator) CurrentRowId() string {
	return fmt.Sprintf("line#%d", t.currentIndex)
}


func (t *CsvFileBasicDataTableRowIterator) Next() datatable.BasicDataTableRow {
	if t.currentIndex+1 >= len(t.dataTable.allLines) {
		return nil
	}

	t.currentIndex++

	rowItems := t.dataTable.allLines[t.currentIndex]

	return &CsvFileBasicDataTableRow{
		dataTable: t.dataTable,
		allItems:  rowItems,
	}
}


func CreateNewCsvBasicDataTable(ctx core.Context, reader io.Reader, hasTitleLine bool) (datatable.BasicDataTable, error) {
	return createNewCsvFileBasicDataTable(ctx, reader, ',', hasTitleLine)
}


func CreateNewCustomCsvBasicDataTable(allLines [][]string, hasTitleLine bool) datatable.BasicDataTable {
	return &CsvFileBasicDataTable{
		allLines:     allLines,
		hasTitleLine: hasTitleLine,
	}
}

func createNewCsvFileBasicDataTable(ctx core.Context, reader io.Reader, separator rune, hasTitleLine bool) (*CsvFileBasicDataTable, error) {
	csvReader := csv.NewReader(reader)
	csvReader.Comma = separator
	csvReader.FieldsPerRecord = -1

	allLines := make([][]string, 0)

	for {
		items, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Errorf(ctx, "[csv_file_basic_data_table.createNewCsvFileDataTable] cannot parse csv data, because %s", err.Error())
			return nil, errs.ErrInvalidCSVFile
		}

		if len(items) == 1 && items[0] == "" {
			continue
		}

		allLines = append(allLines, items)
	}

	return &CsvFileBasicDataTable{
		allLines:     allLines,
		hasTitleLine: hasTitleLine,
	}, nil
}
