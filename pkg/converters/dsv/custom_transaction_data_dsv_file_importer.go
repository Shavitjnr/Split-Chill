package dsv

import (
	"bytes"
	"encoding/csv"
	"io"
	"strings"
	"time"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"

	"github.com/Shavitjnr/split-chill-ai/pkg/converters/converter"
	csvconverter "github.com/Shavitjnr/split-chill-ai/pkg/converters/csv"
	"github.com/Shavitjnr/split-chill-ai/pkg/converters/datatable"
	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
	"github.com/Shavitjnr/split-chill-ai/pkg/log"
	"github.com/Shavitjnr/split-chill-ai/pkg/models"
	"github.com/Shavitjnr/split-chill-ai/pkg/utils"
)

var supportedFileTypeSeparators = map[string]rune{
	"custom_csv": ',',
	"custom_tsv": '\t',
	"custom_ssv": ';',
}

var supportedFileEncodings = map[string]encoding.Encoding{
	"utf-8":        unicode.UTF8,                                           
	"utf-8-bom":    unicode.UTF8BOM,                                        
	"utf-16le":     unicode.UTF16(unicode.LittleEndian, unicode.UseBOM),    
	"utf-16be":     unicode.UTF16(unicode.BigEndian, unicode.UseBOM),       
	"utf-16le-bom": unicode.UTF16(unicode.LittleEndian, unicode.ExpectBOM), 
	"utf-16be-bom": unicode.UTF16(unicode.BigEndian, unicode.ExpectBOM),    
	"cp437":        charmap.CodePage437,                                    
	"cp863":        charmap.CodePage863,                                    
	"cp037":        charmap.CodePage037,                                    
	"cp1047":       charmap.CodePage1047,                                   
	"cp1140":       charmap.CodePage1140,                                   
	"iso-8859-1":   charmap.ISO8859_1,                                      
	"cp850":        charmap.CodePage850,                                    
	"cp858":        charmap.CodePage858,                                    
	"windows-1252": charmap.Windows1252,                                    
	"iso-8859-15":  charmap.ISO8859_15,                                     
	"iso-8859-4":   charmap.ISO8859_4,                                      
	"iso-8859-10":  charmap.ISO8859_10,                                     
	"cp865":        charmap.CodePage865,                                    
	"iso-8859-2":   charmap.ISO8859_2,                                      
	"cp852":        charmap.CodePage852,                                    
	"windows-1250": charmap.Windows1250,                                    
	"iso-8859-14":  charmap.ISO8859_14,                                     
	"iso-8859-3":   charmap.ISO8859_3,                                      
	"cp860":        charmap.CodePage860,                                    
	"iso-8859-7":   charmap.ISO8859_7,                                      
	"windows-1253": charmap.Windows1253,                                    
	"iso-8859-9":   charmap.ISO8859_9,                                      
	"windows-1254": charmap.Windows1254,                                    
	"iso-8859-13":  charmap.ISO8859_13,                                     
	"windows-1257": charmap.Windows1257,                                    
	"iso-8859-16":  charmap.ISO8859_16,                                     
	"iso-8859-5":   charmap.ISO8859_5,                                      
	"cp855":        charmap.CodePage855,                                    
	"cp866":        charmap.CodePage866,                                    
	"windows-1251": charmap.Windows1251,                                    
	"koi8r":        charmap.KOI8R,                                          
	"koi8u":        charmap.KOI8U,                                          
	"iso-8859-6":   charmap.ISO8859_6,                                      
	"windows-1256": charmap.Windows1256,                                    
	"iso-8859-8":   charmap.ISO8859_8,                                      
	"cp862":        charmap.CodePage862,                                    
	"windows-1255": charmap.Windows1255,                                    
	"windows-874":  charmap.Windows874,                                     
	"windows-1258": charmap.Windows1258,                                    
	"gb18030":      simplifiedchinese.GB18030,                              
	"gbk":          simplifiedchinese.GBK,                                  
	"big5":         traditionalchinese.Big5,                                
	"euc-kr":       korean.EUCKR,                                           
	"euc-jp":       japanese.EUCJP,                                         
	"iso-2022-jp":  japanese.ISO2022JP,                                     
	"shift_jis":    japanese.ShiftJIS,                                      
}

var customTransactionTypeNameMapping = map[models.TransactionType]string{
	models.TRANSACTION_TYPE_MODIFY_BALANCE: utils.IntToString(int(models.TRANSACTION_TYPE_MODIFY_BALANCE)),
	models.TRANSACTION_TYPE_INCOME:         utils.IntToString(int(models.TRANSACTION_TYPE_INCOME)),
	models.TRANSACTION_TYPE_EXPENSE:        utils.IntToString(int(models.TRANSACTION_TYPE_EXPENSE)),
	models.TRANSACTION_TYPE_TRANSFER:       utils.IntToString(int(models.TRANSACTION_TYPE_TRANSFER)),
}

type CustomTransactionDataDsvFileParser interface {
	ParseDsvFileLines(ctx core.Context, data []byte) ([][]string, error)
}


type customTransactionDataDsvFileImporter struct {
	fileEncoding               encoding.Encoding
	separator                  rune
	columnIndexMapping         map[datatable.TransactionDataTableColumn]int
	transactionTypeNameMapping map[string]models.TransactionType
	hasHeaderLine              bool
	timeFormat                 string
	timezoneFormat             string
	amountDecimalSeparator     string
	amountDigitGroupingSymbol  string
	geoLocationSeparator       string
	geoLocationOrder           converter.TransactionGeoLocationOrder
	transactionTagSeparator    string
}


func (c *customTransactionDataDsvFileImporter) ParseDsvFileLines(ctx core.Context, data []byte) ([][]string, error) {
	reader := transform.NewReader(bytes.NewReader(data), c.fileEncoding.NewDecoder())
	csvReader := csv.NewReader(reader)
	csvReader.Comma = c.separator
	csvReader.FieldsPerRecord = -1

	allLines := make([][]string, 0)

	for {
		items, err := csvReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Errorf(ctx, "[custom_transaction_data_dsv_file_importer.ParseDsvFileLines] cannot parse dsv data, because %s", err.Error())
			return nil, errs.ErrInvalidCSVFile
		}

		if len(items) == 1 && items[0] == "" {
			continue
		}

		for index := range items {
			items[index] = strings.Trim(items[index], " ")
		}

		allLines = append(allLines, items)
	}

	return allLines, nil
}


func (c *customTransactionDataDsvFileImporter) ParseImportedData(ctx core.Context, user *models.User, data []byte, defaultTimezone *time.Location, additionalOptions converter.TransactionDataImporterOptions, accountMap map[string]*models.Account, expenseCategoryMap map[string]map[string]*models.TransactionCategory, incomeCategoryMap map[string]map[string]*models.TransactionCategory, transferCategoryMap map[string]map[string]*models.TransactionCategory, tagMap map[string]*models.TransactionTag) (models.ImportedTransactionSlice, []*models.Account, []*models.TransactionCategory, []*models.TransactionCategory, []*models.TransactionCategory, []*models.TransactionTag, error) {
	allLines, err := c.ParseDsvFileLines(ctx, data)

	if err != nil {
		return nil, nil, nil, nil, nil, nil, err
	}

	dataTable := csvconverter.CreateNewCustomCsvBasicDataTable(allLines, c.hasHeaderLine)
	transactionDataTable := CreateNewCustomPlainTextDataTable(dataTable, c.columnIndexMapping, c.transactionTypeNameMapping, c.timeFormat, c.timezoneFormat, c.amountDecimalSeparator, c.amountDigitGroupingSymbol)
	dataTableImporter := converter.CreateNewImporterWithTypeNameMapping(customTransactionTypeNameMapping, c.geoLocationSeparator, c.geoLocationOrder, c.transactionTagSeparator)

	return dataTableImporter.ParseImportedData(ctx, user, transactionDataTable, defaultTimezone, additionalOptions, accountMap, expenseCategoryMap, incomeCategoryMap, transferCategoryMap, tagMap)
}


func IsDelimiterSeparatedValuesFileType(fileType string) bool {
	_, exists := supportedFileTypeSeparators[fileType]
	return exists
}


func CreateNewCustomTransactionDataDsvFileParser(fileType string, fileEncoding string) (CustomTransactionDataDsvFileParser, error) {
	separator, exists := supportedFileTypeSeparators[fileType]

	if !exists {
		return nil, errs.ErrImportFileTypeNotSupported
	}

	enc, exists := supportedFileEncodings[fileEncoding]

	if !exists {
		return nil, errs.ErrImportFileEncodingNotSupported
	}

	return &customTransactionDataDsvFileImporter{
		fileEncoding: enc,
		separator:    separator,
	}, nil
}


func CreateNewCustomTransactionDataDsvFileImporter(fileType string, fileEncoding string, columnIndexMapping map[datatable.TransactionDataTableColumn]int, transactionTypeNameMapping map[string]models.TransactionType, hasHeaderLine bool, timeFormat string, timezoneFormat string, amountDecimalSeparator string, amountDigitGroupingSymbol string, geoLocationSeparator string, geoLocationOrder string, transactionTagSeparator string) (converter.TransactionDataImporter, error) {
	separator, exists := supportedFileTypeSeparators[fileType]

	if !exists {
		return nil, errs.ErrImportFileTypeNotSupported
	}

	enc, exists := supportedFileEncodings[fileEncoding]

	if !exists {
		return nil, errs.ErrImportFileEncodingNotSupported
	}

	if geoLocationOrder == "" {
		geoLocationOrder = string(converter.TRANSACTION_GEO_LOCATION_ORDER_LONGITUDE_LATITUDE)
	} else if geoLocationOrder != string(converter.TRANSACTION_GEO_LOCATION_ORDER_LONGITUDE_LATITUDE) &&
		geoLocationOrder != string(converter.TRANSACTION_GEO_LOCATION_ORDER_LATITUDE_LONGITUDE) {
		return nil, errs.ErrImportFileTypeNotSupported
	}

	if _, exists = columnIndexMapping[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TIME]; !exists {
		return nil, errs.ErrMissingRequiredFieldInHeaderRow
	}

	if _, exists = columnIndexMapping[datatable.TRANSACTION_DATA_TABLE_TRANSACTION_TYPE]; !exists {
		return nil, errs.ErrMissingRequiredFieldInHeaderRow
	}

	if _, exists = columnIndexMapping[datatable.TRANSACTION_DATA_TABLE_AMOUNT]; !exists {
		return nil, errs.ErrMissingRequiredFieldInHeaderRow
	}

	return &customTransactionDataDsvFileImporter{
		fileEncoding:               enc,
		separator:                  separator,
		columnIndexMapping:         columnIndexMapping,
		transactionTypeNameMapping: transactionTypeNameMapping,
		hasHeaderLine:              hasHeaderLine,
		timeFormat:                 timeFormat,
		timezoneFormat:             timezoneFormat,
		amountDecimalSeparator:     amountDecimalSeparator,
		amountDigitGroupingSymbol:  amountDigitGroupingSymbol,
		geoLocationSeparator:       geoLocationSeparator,
		geoLocationOrder:           converter.TransactionGeoLocationOrder(geoLocationOrder),
		transactionTagSeparator:    transactionTagSeparator,
	}, nil
}
