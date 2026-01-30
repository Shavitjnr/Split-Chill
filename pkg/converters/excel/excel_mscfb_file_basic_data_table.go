package excel

import (
	"bytes"
	"fmt"

	"github.com/extrame/xls"

	"github.com/Shavitjnr/split-chill-ai/pkg/converters/datatable"
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
)


type ExcelMSCFBFileBasicDataTable struct {
	workbook              *xls.WorkBook
	headerLineColumnNames []string
	hasTitleLine          bool
}


type ExcelMSCFBFileBasicDataTableRow struct {
	sheet    *xls.WorkSheet
	rowIndex int
}


type ExcelMSCFBFileBasicDataTableRowIterator struct {
	dataTable              *ExcelMSCFBFileBasicDataTable
	currentSheetIndex      int
	currentRowIndexInSheet int
}


func (t *ExcelMSCFBFileBasicDataTable) DataRowCount() int {
	totalDataRowCount := 0

	for i := 0; i < t.workbook.NumSheets(); i++ {
		sheet := t.workbook.GetSheet(i)

		if sheet == nil {
			continue
		}

		if t.hasTitleLine {
			if sheet.MaxRow < 1 {
				continue
			}

			totalDataRowCount += int(sheet.MaxRow)
		} else {
			if sheet.MaxRow <= 0 && sheet.Row(0) == nil {
				continue
			}

			totalDataRowCount += int(sheet.MaxRow) + 1
		}
	}

	return totalDataRowCount
}


func (t *ExcelMSCFBFileBasicDataTable) HeaderColumnNames() []string {
	if !t.hasTitleLine {
		return nil
	}

	return t.headerLineColumnNames
}


func (t *ExcelMSCFBFileBasicDataTable) DataRowIterator() datatable.BasicDataTableRowIterator {
	startIndex := -1

	if t.hasTitleLine {
		startIndex = 0
	}

	return &ExcelMSCFBFileBasicDataTableRowIterator{
		dataTable:              t,
		currentSheetIndex:      0,
		currentRowIndexInSheet: startIndex,
	}
}


func (r *ExcelMSCFBFileBasicDataTableRow) ColumnCount() int {
	row := r.sheet.Row(r.rowIndex)
	return row.LastCol() + 1
}


func (r *ExcelMSCFBFileBasicDataTableRow) GetData(columnIndex int) string {
	row := r.sheet.Row(r.rowIndex)
	return row.Col(columnIndex)
}


func (t *ExcelMSCFBFileBasicDataTableRowIterator) HasNext() bool {
	workbook := t.dataTable.workbook

	if t.currentSheetIndex >= workbook.NumSheets() {
		return false
	}

	currentSheet := workbook.GetSheet(t.currentSheetIndex)

	if t.currentRowIndexInSheet+1 <= int(currentSheet.MaxRow) && currentSheet.Row(t.currentRowIndexInSheet+1) != nil {
		return true
	}

	for i := t.currentSheetIndex + 1; i < workbook.NumSheets(); i++ {
		sheet := workbook.GetSheet(i)

		if t.dataTable.hasTitleLine {
			if sheet.MaxRow < 1 {
				continue
			}
		} else {
			if sheet.MaxRow <= 0 && sheet.Row(0) == nil {
				continue
			}
		}

		return true
	}

	return false
}


func (t *ExcelMSCFBFileBasicDataTableRowIterator) CurrentRowId() string {
	return fmt.Sprintf("sheet#%d-row#%d", t.currentSheetIndex, t.currentRowIndexInSheet)
}


func (t *ExcelMSCFBFileBasicDataTableRowIterator) Next() datatable.BasicDataTableRow {
	workbook := t.dataTable.workbook

	for i := t.currentSheetIndex; i < workbook.NumSheets(); i++ {
		sheet := workbook.GetSheet(i)

		if t.currentRowIndexInSheet+1 <= int(sheet.MaxRow) && sheet.Row(t.currentRowIndexInSheet+1) != nil {
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

	if t.currentSheetIndex >= workbook.NumSheets() {
		return nil
	}

	currentSheet := workbook.GetSheet(t.currentSheetIndex)

	if t.currentRowIndexInSheet > int(currentSheet.MaxRow) || currentSheet.Row(t.currentRowIndexInSheet) == nil {
		return nil
	}

	return &ExcelMSCFBFileBasicDataTableRow{
		sheet:    currentSheet,
		rowIndex: int(t.currentRowIndexInSheet),
	}
}


func CreateNewExcelMSCFBFileBasicDataTable(data []byte, hasTitleLine bool) (datatable.BasicDataTable, error) {
	reader := bytes.NewReader(data)
	workbook, err := xls.OpenReader(reader, "")

	if err != nil {
		return nil, err
	}

	var firstRowItems []string

	for i := 0; i < workbook.NumSheets(); i++ {
		sheet := workbook.GetSheet(i)

		if sheet.MaxRow <= 0 && sheet.Row(0) == nil {
			continue
		}

		row := sheet.Row(0)

		if row == nil {
			continue
		}

		if i == 0 {
			for j := 0; j <= row.LastCol(); j++ {
				headerItem := row.Col(j)

				if headerItem == "" {
					break
				}

				firstRowItems = append(firstRowItems, headerItem)
			}
		} else {
			for j := 0; j <= min(row.LastCol(), len(firstRowItems)-1); j++ {
				headerItem := row.Col(j)

				if headerItem != firstRowItems[j] {
					return nil, errs.ErrFieldsInMultiTableAreDifferent
				}
			}
		}
	}

	var headerLineColumnNames []string = nil

	if hasTitleLine {
		headerLineColumnNames = firstRowItems
	}

	return &ExcelMSCFBFileBasicDataTable{
		workbook:              workbook,
		headerLineColumnNames: headerLineColumnNames,
		hasTitleLine:          hasTitleLine,
	}, nil
}
