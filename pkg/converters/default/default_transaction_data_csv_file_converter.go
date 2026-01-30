package _default


type defaultTransactionDataCSVFileConverter struct {
	defaultTransactionDataPlainTextConverter
}


var (
	DefaultTransactionDataCSVFileConverter = &defaultTransactionDataCSVFileConverter{
		defaultTransactionDataPlainTextConverter{
			columnSeparator: ",",
		},
	}
)
