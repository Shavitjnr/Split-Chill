package qif


type qifTransactionClearedStatus string


const (
	qifClearedStatusUnreconciled qifTransactionClearedStatus = ""
	qifClearedStatusCleared      qifTransactionClearedStatus = "C"
	qifClearedStatusReconciled   qifTransactionClearedStatus = "R"
)


type qifTransactionType string


const (
	qifInvalidTransactionType         qifTransactionType = ""
	qifCheckTransactionType           qifTransactionType = "KC"
	qifDepositTransactionType         qifTransactionType = "KD"
	qifPaymentTransactionType         qifTransactionType = "KP"
	qifInvestmentTransactionType      qifTransactionType = "KI"
	qifElectronicPayeeTransactionType qifTransactionType = "KE"
)


type qifCategoryType string


const (
	qifIncomeTransaction  qifCategoryType = "I"
	qifExpenseTransaction qifCategoryType = "E"
)


type qifData struct {
	BankAccountTransactions       []*qifTransactionData
	CashAccountTransactions       []*qifTransactionData
	CreditCardAccountTransactions []*qifTransactionData
	AssetAccountTransactions      []*qifTransactionData
	LiabilityAccountTransactions  []*qifTransactionData
	MemorizedTransactions         []*qifMemorizedTransactionData
	InvestmentAccountTransactions []*qifInvestmentTransactionData
	Accounts                      []*qifAccountData
	Categories                    []*qifCategoryData
	Classes                       []*qifClassData
}


type qifTransactionData struct {
	Date                   string
	Amount                 string
	ClearedStatus          qifTransactionClearedStatus
	Num                    string
	Payee                  string
	Memo                   string
	Addresses              []string
	Category               string
	SubTransactionCategory []string
	SubTransactionMemo     []string
	SubTransactionAmount   []string
	Account                *qifAccountData
}


type qifInvestmentTransactionData struct {
	Date               string
	Action             string
	Security           string
	Price              string
	Quantity           string
	Amount             string
	ClearedStatus      qifTransactionClearedStatus
	Text               string
	Memo               string
	Commission         string
	AccountForTransfer string
	AmountTransferred  string
	Account            *qifAccountData
}


type qifMemorizedTransactionData struct {
	qifTransactionData
	TransactionType qifTransactionType
	Amortization    qifMemorizedTransactionAmortizationData
}


type qifMemorizedTransactionAmortizationData struct {
	FirstPaymentDate       string
	TotalYearsForLoan      string
	NumberOfPayments       string
	NumberOfPeriodsPerYear string
	InterestRate           string
	CurrentLoanBalance     string
	OriginalLoanAmount     string
}


type qifAccountData struct {
	Name                   string
	AccountType            string
	Description            string
	CreditLimit            string
	StatementBalanceDate   string
	StatementBalanceAmount string
}


type qifCategoryData struct {
	Name                   string
	Description            string
	TaxRelated             bool
	CategoryType           qifCategoryType
	BudgetAmount           string
	TaxScheduleInformation string
}


type qifClassData struct {
	Name        string
	Description string
}
