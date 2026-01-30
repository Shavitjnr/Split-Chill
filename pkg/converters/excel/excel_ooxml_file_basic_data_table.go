package excel

import (
	"bytes"
	"fmt"

	"github.com/xuri/excelize/v2"

	"github.com/Shavitjnr/split-chill-ai/pkg/converters/datatable"
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
)


type excelOOXMLSheet struct {
	sheetName string
	allData   [][]string
}


type ExcelOOXMLFileBasicDataTable struct {
	sheets                []*excelOOXMLSheet
	headerLineColumnNames []string
	hasTitleLine          bool
}


type ExcelOOXMLFileBasicDataTableRow struct {
	sheet    *excelOOXMLSheet
	rowData  []string
	rowIndex int
}


type ExcelOOXMLFileBasicDataTableRowIterator struct {
	dataTable              *ExcelOOXMLFileBasicDataTable
	currentSheetIndex      int
	currentRowIndexInSheet int
}


func (t *ExcelOOXMLFileBasicDataTable) DataRowCount() int {
	totalDataRowCount := 0

	for i := 0; i < len(t.sheets); i++ {
		sheet := t.sheets[i]

		if len(sheet.allData) < 1 {
			continue
		}

		if t.hasTitleLine {
			totalDataRowCount += len(sheet.allData) - 1
		} else {
			totalDataRowCount += len(sheet.allData)
		}
	}

	return totalDataRowCount
}


func (t *ExcelOOXMLFileBasicDataTable) HeaderColumnNames() []string {
	if !t.hasTitleLine {
		return nil
	}

	return t.headerLineColumnNames
}


func (t *ExcelOOXMLFileBasicDataTable) DataRowIterator() datatable.BasicDataTableRowIterator {
	startIndex := -1

	if t.hasTitleLine {
		startIndex = 0
	}

	return &ExcelOOXMLFileBasicDataTableRowIterator{
		dataTable:              t,
		currentSheetIndex:      0,
		currentRowIndexInSheet: startIndex,
	}
}


func (r *ExcelOOXMLFileBasicDataTableRow) ColumnCount() int {
	return len(r.rowData)
}


func (r *ExcelOOXMLFileBasicDataTableRow) GetData(columnIndex int) string {
	if columnIndex < 0 || columnIndex >= len(r.rowData) {
		return ""
	}

	return r.rowData[columnIndex]
}


func (t *ExcelOOXMLFileBasicDataTableRowIterator) HasNext() bool {
	sheets := t.dataTable.sheets

	if t.currentSheetIndex >= len(sheets) {
		return false
	}

	currentSheet := sheets[t.currentSheetIndex]

	if t.currentRowIndexInSheet+1 < len(currentSheet.allData) {
		return true
	}

	for i := t.currentSheetIndex + 1; i < len(sheets); i++ {
		sheet := sheets[i]

		if t.dataTable.hasTitleLine {
			if len(sheet.allData) <= 1 {
				continue
			}
		} else {
			if len(sheet.allData) <= 0 {
				continue
			}
		}

		return true
	}

	return false
}


func (t *ExcelOOXMLFileBasicDataTableRowIterator) CurrentRowId() string {
	return fmt.Sprintf("sheet#%d-row#%d", t.currentSheetIndex, t.currentRowIndexInSheet)
}


func (t *ExcelOOXMLFileBasicDataTableRowIterator) Next() datatable.BasicDataTableRow {
	sheets := t.dataTable.sheets

	for i := t.currentSheetIndex; i < len(sheets); i++ {
		sheet := sheets[i]

		if t.currentRowIndexInSheet+1 < len(sheet.allData) {
			t.currentRowIndexInSheet++
			break
		}

		t.currentSheetIndex++

		if t.dataTable.hasTitleLine {
			t.currentRowIndexInSheet = 0
		} else {
			t.currentRowIndexInSheet = -1
		}
	}

	if t.currentSheetIndex >= len(sheets) {
		return nil
	}

	currentSheet := sheets[t.currentSheetIndex]

	if t.currentRowIndexInSheet >= len(currentSheet.allData) {
		return nil
	}

	return &ExcelOOXMLFileBasicDataTableRow{
		sheet:    currentSheet,
		rowData:  currentSheet.allData[t.currentRowIndexInSheet],
		rowIndex: t.currentRowIndexInSheet,
	}
}


func CreateNewExcelOOXMLFileBasicDataTable(data []byte, hasTitleLine bool) (datatable.BasicDataTable, error) {
	reader := bytes.NewReader(data)
	file, err := excelize.OpenReader(reader)

	defer file.Close()

	if err != nil {
		return nil, err
	}

	sheetNames := file.GetSheetList()
	var firstRowItems []string
	var sheets []*excelOOXMLSheet

	for i := 0; i < len(sheetNames); i++ {
		sheetName := sheetNames[i]
		allData, err := file.GetRows(sheetName)

		if err != nil {
			return nil, err
		}

		if allData == nil || len(allData) < 1 {
			continue
		}

		row := allData[0]

		if i == 0 {
			for j := 0; j < len(row); j++ {
				headerItem := row[j]

				if headerItem == "" {
					break
				}

				firstRowItems = append(firstRowItems, headerItem)
			}
		} else {
			for j := 0; j < min(len(row), len(firstRowItems)); j++ {
				headerItem := row[j]

				if headerItem != firstRowItems[j] {
					return nil, errs.ErrFieldsInMultiTableAreDifferent
				}
			}
		}

		sheets = append(sheets, &excelOOXMLSheet{
			sheetName: sheetName,
			allData:   allData,
		})
	}

	var headerLineColumnNames []string = nil

	if hasTitleLine {
		headerLineColumnNames = firstRowItems
	}

	return &ExcelOOXMLFileBasicDataTable{
		sheets:                sheets,
		headerLineColumnNames: headerLineColumnNames,
		hasTitleLine:          hasTitleLine,
	}, nil
}
