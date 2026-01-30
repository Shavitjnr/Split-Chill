package _default


type defaultTransactionDataTSVFileConverter struct {
	defaultTransactionDataPlainTextConverter
}


var (
	DefaultTransactionDataTSVFileConverter = &defaultTransactionDataTSVFileConverter{
		defaultTransactionDataPlainTextConverter{
			columnSeparator: "\t",
		},
	}
)
