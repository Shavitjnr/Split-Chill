package ofx

import (
	"encoding/xml"

	"github.com/Shavitjnr/split-chill-ai/pkg/models"
)


type oFXDeclarationVersion string

const (
	ofxVersion1 oFXDeclarationVersion = "100"
	ofxVersion2 oFXDeclarationVersion = "200"
)

const ofxDefaultTimezoneOffset = "+00:00"


type ofxAccountType string


const (
	ofxCheckingAccount             ofxAccountType = "CHECKING"
	ofxSavingsAccount              ofxAccountType = "SAVINGS"
	ofxMoneyMarketAccount          ofxAccountType = "MONEYMRKT"
	ofxLineOfCreditAccount         ofxAccountType = "CREDITLINE"
	ofxCertificateOfDepositAccount ofxAccountType = "CD"
)


type ofxTransactionType string


const (
	ofxGenericCreditTransaction          ofxTransactionType = "CREDIT"
	ofxGenericDebitTransaction           ofxTransactionType = "DEBIT"
	ofxInterestTransaction               ofxTransactionType = "INT"
	ofxDividendTransaction               ofxTransactionType = "DIV"
	ofxFIFeeTransaction                  ofxTransactionType = "FEE"
	ofxServiceChargeTransaction          ofxTransactionType = "SRVCHG"
	ofxDepositTransaction                ofxTransactionType = "DEP"
	ofxATMTransaction                    ofxTransactionType = "ATM"
	ofxPOSTransaction                    ofxTransactionType = "POS"
	ofxTransferTransaction               ofxTransactionType = "XFER"
	ofxCheckTransaction                  ofxTransactionType = "CHECK"
	ofxElectronicPaymentTransaction      ofxTransactionType = "PAYMENT"
	ofxCashWithdrawalTransaction         ofxTransactionType = "CASH"
	ofxDirectDepositTransaction          ofxTransactionType = "DIRECTDEP"
	ofxMerchantInitiatedDebitTransaction ofxTransactionType = "DIRECTDEBIT"
	ofxRepeatingPaymentTransaction       ofxTransactionType = "REPEATPMT"
	ofxHoldTransaction                   ofxTransactionType = "HOLD"
	ofxOtherTransaction                  ofxTransactionType = "OTHER"
)

var ofxTransactionTypeMapping = map[ofxTransactionType]models.TransactionType{
	ofxGenericCreditTransaction:          models.TRANSACTION_TYPE_EXPENSE,
	ofxGenericDebitTransaction:           models.TRANSACTION_TYPE_EXPENSE,
	ofxDividendTransaction:               models.TRANSACTION_TYPE_INCOME,
	ofxFIFeeTransaction:                  models.TRANSACTION_TYPE_EXPENSE,
	ofxServiceChargeTransaction:          models.TRANSACTION_TYPE_EXPENSE,
	ofxDepositTransaction:                models.TRANSACTION_TYPE_INCOME,
	ofxTransferTransaction:               models.TRANSACTION_TYPE_TRANSFER,
	ofxCheckTransaction:                  models.TRANSACTION_TYPE_EXPENSE,
	ofxElectronicPaymentTransaction:      models.TRANSACTION_TYPE_EXPENSE,
	ofxCashWithdrawalTransaction:         models.TRANSACTION_TYPE_EXPENSE,
	ofxDirectDepositTransaction:          models.TRANSACTION_TYPE_INCOME,
	ofxMerchantInitiatedDebitTransaction: models.TRANSACTION_TYPE_EXPENSE,
	ofxRepeatingPaymentTransaction:       models.TRANSACTION_TYPE_EXPENSE,
}


type ofxFile struct {
	XMLName                     xml.Name `xml:"OFX"`
	FileHeader                  *ofxFileHeader
	BankMessageResponseV1       *ofxBankMessageResponseV1       `xml:"BANKMSGSRSV1"`
	CreditCardMessageResponseV1 *ofxCreditCardMessageResponseV1 `xml:"CREDITCARDMSGSRSV1"`
}


type ofxFileHeader struct {
	OFXDeclarationVersion oFXDeclarationVersion
	OFXDataVersion        string
	Security              string
	OldFileUid            string
	NewFileUid            string
}


type ofxBankMessageResponseV1 struct {
	StatementTransactionResponse *ofxBankStatementTransactionResponse `xml:"STMTTRNRS"`
}


type ofxCreditCardMessageResponseV1 struct {
	StatementTransactionResponse *ofxCreditCardStatementTransactionResponse `xml:"CCSTMTTRNRS"`
}


type ofxBankStatementTransactionResponse struct {
	StatementResponse *ofxBankStatementResponse `xml:"STMTRS"`
}


type ofxCreditCardStatementTransactionResponse struct {
	StatementResponse *ofxCreditCardStatementResponse `xml:"CCSTMTRS"`
}


type ofxBankStatementResponse struct {
	DefaultCurrency string                  `xml:"CURDEF"`
	AccountFrom     *ofxBankAccount         `xml:"BANKACCTFROM"`
	TransactionList *ofxBankTransactionList `xml:"BANKTRANLIST"`
}


type ofxCreditCardStatementResponse struct {
	DefaultCurrency string                        `xml:"CURDEF"`
	AccountFrom     *ofxCreditCardAccount         `xml:"CCACCTFROM"`
	TransactionList *ofxCreditCardTransactionList `xml:"BANKTRANLIST"`
}


type ofxBankAccount struct {
	BankId      string         `xml:"BANKID"`
	BranchId    string         `xml:"BRANCHID"`
	AccountId   string         `xml:"ACCTID"`
	AccountType ofxAccountType `xml:"ACCTTYPE"`
	AccountKey  string         `xml:"ACCTKEY"`
}


type ofxCreditCardAccount struct {
	AccountId  string `xml:"ACCTID"`
	AccountKey string `xml:"ACCTKEY"`
}


type ofxBankTransactionList struct {
	StartDate             string                         `xml:"DTSTART"`
	EndDate               string                         `xml:"DTEND"`
	StatementTransactions []*ofxBankStatementTransaction `xml:"STMTTRN"`
}


type ofxCreditCardTransactionList struct {
	StartDate             string                               `xml:"DTSTART"`
	EndDate               string                               `xml:"DTEND"`
	StatementTransactions []*ofxCreditCardStatementTransaction `xml:"STMTTRN"`
}


type ofxBaseStatementTransaction struct {
	TransactionId    string             `xml:"FITID"`
	TransactionType  ofxTransactionType `xml:"TRNTYPE"`
	PostedDate       string             `xml:"DTPOSTED"`
	Amount           string             `xml:"TRNAMT"`
	Name             string             `xml:"NAME"`
	Payee            *ofxPayee          `xml:"PAYEE"`
	Memo             string             `xml:"MEMO"`
	Currency         string             `xml:"CURRENCY"`
	OriginalCurrency string             `xml:"ORIGCURRENCY"`
}


type ofxBankStatementTransaction struct {
	ofxBaseStatementTransaction
	AccountTo *ofxBankAccount `xml:"BANKACCTTO"`
}


type ofxCreditCardStatementTransaction struct {
	ofxBaseStatementTransaction
	AccountTo *ofxCreditCardAccount `xml:"CCACCTTO"`
}


type ofxPayee struct {
	Name       string `xml:"NAME"`
	Address1   string `xml:"ADDR1"`
	Address2   string `xml:"ADDR2"`
	Address3   string `xml:"ADDR3"`
	City       string `xml:"CITY"`
	State      string `xml:"STATE"`
	PostalCode string `xml:"POSTALCODE"`
	Country    string `xml:"COUNTRY"`
	Phone      string `xml:"PHONE"`
}
