package gnucash

import "encoding/xml"

const gnucashCommodityCurrencySpace = "CURRENCY"
const gnucashRootAccountType = "ROOT"
const gnucashEquityAccountType = "EQUITY"
const gnucashIncomeAccountType = "INCOME"
const gnucashExpenseAccountType = "EXPENSE"

const gnucashSlotEquityType = "equity-type"
const gnucashSlotEquityTypeOpeningBalance = "opening-balance"

var gnucashAssetOrLiabilityAccountTypes = map[string]bool{
	"ASSET":      true,
	"BANK":       true,
	"CASH":       true,
	"CREDIT":     true,
	"LIABILITY":  true,
	"MUTUAL":     true,
	"PAYABLE":    true,
	"RECEIVABLE": true,
	"STOCK":      true,
}


type gnucashDatabase struct {
	XMLName xml.Name            `xml:"gnc-v2"`
	Counts  []*gnucashCountData `xml:"count-data"`
	Books   []*gnucashBookData  `xml:"book"`
}


type gnucashCountData struct {
	Key   string `xml:"type,attr"`
	Value string `xml:",chardata"`
}


type gnucashBookData struct {
	Id           string                    `xml:"id"`
	Counts       []*gnucashCountData       `xml:"count-data"`
	Accounts     []*gnucashAccountData     `xml:"account"`
	Transactions []*gnucashTransactionData `xml:"transaction"`
}


type gnucashCommodityData struct {
	Space string `xml:"space"`
	Id    string `xml:"id"`
}


type gnucashSlotData struct {
	Key   string `xml:"key"`
	Value string `xml:"value"`
}


type gnucashAccountData struct {
	Name        string                `xml:"name"`
	Id          string                `xml:"id"`
	AccountType string                `xml:"type"`
	Description string                `xml:"description"`
	ParentId    string                `xml:"parent"`
	Commodity   *gnucashCommodityData `xml:"commodity"`
	Slots       []*gnucashSlotData    `xml:"slots>slot"`
}


type gnucashTransactionData struct {
	Id          string                         `xml:"id"`
	Currency    *gnucashCommodityData          `xml:"currency"`
	PostedDate  string                         `xml:"date-posted>date"`
	EnteredDate string                         `xml:"date-entered>date"`
	Description string                         `xml:"description"`
	Splits      []*gnucashTransactionSplitData `xml:"splits>split"`
}


type gnucashTransactionSplitData struct {
	Id              string `xml:"id"`
	ReconciledState string `xml:"reconciled-state"`
	Value           string `xml:"value"`
	Quantity        string `xml:"quantity"`
	Account         string `xml:"account"`
}
