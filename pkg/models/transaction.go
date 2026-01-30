package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
	"github.com/Shavitjnr/split-chill-ai/pkg/utils"
)

const MaximumTagsCountOfTransaction = 10
const MaximumPicturesCountOfTransaction = 10


type TransactionType byte


const (
	TRANSACTION_TYPE_MODIFY_BALANCE TransactionType = 1
	TRANSACTION_TYPE_INCOME         TransactionType = 2
	TRANSACTION_TYPE_EXPENSE        TransactionType = 3
	TRANSACTION_TYPE_TRANSFER       TransactionType = 4
)


func (t TransactionType) ToTransactionDbType() (TransactionDbType, error) {
	if t == TRANSACTION_TYPE_MODIFY_BALANCE {
		return TRANSACTION_DB_TYPE_MODIFY_BALANCE, nil
	} else if t == TRANSACTION_TYPE_EXPENSE {
		return TRANSACTION_DB_TYPE_EXPENSE, nil
	} else if t == TRANSACTION_TYPE_INCOME {
		return TRANSACTION_DB_TYPE_INCOME, nil
	} else if t == TRANSACTION_TYPE_TRANSFER {
		return TRANSACTION_DB_TYPE_TRANSFER_OUT, nil
	} else {
		return 0, errs.ErrTransactionTypeInvalid
	}
}


type TransactionRelatedAccountType byte


const (
	TRANSACTION_RELATED_ACCOUNT_TYPE_TRANSFER_FROM TransactionRelatedAccountType = 1
	TRANSACTION_RELATED_ACCOUNT_TYPE_TRANSFER_TO   TransactionRelatedAccountType = 2
)


type TransactionDbType byte


const (
	TRANSACTION_DB_TYPE_MODIFY_BALANCE TransactionDbType = 1
	TRANSACTION_DB_TYPE_INCOME         TransactionDbType = 2
	TRANSACTION_DB_TYPE_EXPENSE        TransactionDbType = 3
	TRANSACTION_DB_TYPE_TRANSFER_OUT   TransactionDbType = 4
	TRANSACTION_DB_TYPE_TRANSFER_IN    TransactionDbType = 5
)


func (t TransactionDbType) String() string {
	switch t {
	case TRANSACTION_DB_TYPE_MODIFY_BALANCE:
		return "Modify Balance"
	case TRANSACTION_DB_TYPE_INCOME:
		return "Income"
	case TRANSACTION_DB_TYPE_EXPENSE:
		return "Expense"
	case TRANSACTION_DB_TYPE_TRANSFER_OUT:
		return "Transfer Out"
	case TRANSACTION_DB_TYPE_TRANSFER_IN:
		return "Transfer In"
	default:
		return fmt.Sprintf("Invalid(%d)", int(t))
	}
}


func (t TransactionDbType) ToTransactionType() (TransactionType, error) {
	if t == TRANSACTION_DB_TYPE_MODIFY_BALANCE {
		return TRANSACTION_TYPE_MODIFY_BALANCE, nil
	} else if t == TRANSACTION_DB_TYPE_EXPENSE {
		return TRANSACTION_TYPE_EXPENSE, nil
	} else if t == TRANSACTION_DB_TYPE_INCOME {
		return TRANSACTION_TYPE_INCOME, nil
	} else if t == TRANSACTION_DB_TYPE_TRANSFER_OUT {
		return TRANSACTION_TYPE_TRANSFER, nil
	} else if t == TRANSACTION_DB_TYPE_TRANSFER_IN {
		return TRANSACTION_TYPE_TRANSFER, nil
	} else {
		return 0, errs.ErrTransactionTypeInvalid
	}
}


func (t TransactionDbType) ToTransactionRelatedAccountType() (TransactionRelatedAccountType, error) {
	if t == TRANSACTION_DB_TYPE_TRANSFER_OUT {
		return TRANSACTION_RELATED_ACCOUNT_TYPE_TRANSFER_TO, nil
	} else if t == TRANSACTION_DB_TYPE_TRANSFER_IN {
		return TRANSACTION_RELATED_ACCOUNT_TYPE_TRANSFER_FROM, nil
	} else {
		return 0, errs.ErrTransactionTypeInvalid
	}
}


const TransactionNoTagFilterValue = "none"


type TransactionTagFilterType byte


const (
	TRANSACTION_TAG_FILTER_HAS_ANY     TransactionTagFilterType = 0
	TRANSACTION_TAG_FILTER_HAS_ALL     TransactionTagFilterType = 1
	TRANSACTION_TAG_FILTER_NOT_HAS_ANY TransactionTagFilterType = 2
	TRANSACTION_TAG_FILTER_NOT_HAS_ALL TransactionTagFilterType = 3
)


type Transaction struct {
	TransactionId        int64             `xorm:"PK"`
	Uid                  int64             `xorm:"UNIQUE(UQE_transaction_uid_time) INDEX(IDX_transaction_uid_deleted_time) INDEX(IDX_transaction_uid_deleted_type_time) INDEX(IDX_transaction_uid_deleted_type_account_id_time) INDEX(IDX_transaction_uid_deleted_category_id_time) INDEX(IDX_transaction_uid_deleted_account_id_time) INDEX(IDX_transaction_uid_deleted_time_longitude_latitude) NOT NULL"`
	Deleted              bool              `xorm:"INDEX(IDX_transaction_uid_deleted_time) INDEX(IDX_transaction_uid_deleted_type_time) INDEX(IDX_transaction_uid_deleted_type_account_id_time) INDEX(IDX_transaction_uid_deleted_category_id_time) INDEX(IDX_transaction_uid_deleted_account_id_time) INDEX(IDX_transaction_uid_deleted_time_longitude_latitude) NOT NULL"`
	Type                 TransactionDbType `xorm:"INDEX(IDX_transaction_uid_deleted_type_time) INDEX(IDX_transaction_uid_deleted_type_account_id_time) NOT NULL"`
	CategoryId           int64             `xorm:"INDEX(IDX_transaction_uid_deleted_category_id_time) NOT NULL"`
	AccountId            int64             `xorm:"INDEX(IDX_transaction_uid_deleted_account_id_time) INDEX(IDX_transaction_uid_deleted_type_account_id_time) NOT NULL"`
	TransactionTime      int64             `xorm:"UNIQUE(UQE_transaction_uid_time) INDEX(IDX_transaction_uid_deleted_time) INDEX(IDX_transaction_uid_deleted_type_time) INDEX(IDX_transaction_uid_deleted_type_account_id_time) INDEX(IDX_transaction_uid_deleted_category_id_time) INDEX(IDX_transaction_uid_deleted_account_id_time) NOT NULL"`
	TimezoneUtcOffset    int16             `xorm:"NOT NULL"`
	Amount               int64             `xorm:"NOT NULL"`
	RelatedId            int64             `xorm:"NOT NULL"`
	RelatedAccountId     int64             `xorm:"NOT NULL"`
	RelatedAccountAmount int64             `xorm:"NOT NULL"`
	HideAmount           bool              `xorm:"NOT NULL"`
	Comment              string            `xorm:"VARCHAR(255) NOT NULL"`
	GeoLongitude         float64           `xorm:"INDEX(IDX_transaction_uid_deleted_time_longitude_latitude)"`
	GeoLatitude          float64           `xorm:"INDEX(IDX_transaction_uid_deleted_time_longitude_latitude)"`
	CreatedIp            string            `xorm:"VARCHAR(39)"`
	ScheduledCreated     bool
	CreatedUnixTime      int64
	UpdatedUnixTime      int64
	DeletedUnixTime      int64
}


type TransactionWithAccountBalance struct {
	*Transaction
	AccountOpeningBalance int64
	AccountClosingBalance int64
}


type TransactionGeoLocationRequest struct {
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}


type TransactionCreateRequest struct {
	Type                 TransactionType                `json:"type" binding:"required"`
	CategoryId           int64                          `json:"categoryId,string"`
	Time                 int64                          `json:"time" binding:"required,min=1"`
	UtcOffset            int16                          `json:"utcOffset" binding:"min=-720,max=840"`
	SourceAccountId      int64                          `json:"sourceAccountId,string" binding:"required,min=1"`
	DestinationAccountId int64                          `json:"destinationAccountId,string" binding:"min=0"`
	SourceAmount         int64                          `json:"sourceAmount" binding:"min=-99999999999,max=99999999999"`
	DestinationAmount    int64                          `json:"destinationAmount" binding:"min=-99999999999,max=99999999999"`
	HideAmount           bool                           `json:"hideAmount"`
	TagIds               []string                       `json:"tagIds"`
	PictureIds           []string                       `json:"pictureIds"`
	Comment              string                         `json:"comment" binding:"max=255"`
	GeoLocation          *TransactionGeoLocationRequest `json:"geoLocation" binding:"omitempty"`
	ClientSessionId      string                         `json:"clientSessionId"`
}


type TransactionModifyRequest struct {
	Id                   int64                          `json:"id,string" binding:"required,min=1"`
	CategoryId           int64                          `json:"categoryId,string"`
	Time                 int64                          `json:"time" binding:"required,min=1"`
	UtcOffset            int16                          `json:"utcOffset" binding:"min=-720,max=840"`
	SourceAccountId      int64                          `json:"sourceAccountId,string" binding:"required,min=1"`
	DestinationAccountId int64                          `json:"destinationAccountId,string" binding:"min=0"`
	SourceAmount         int64                          `json:"sourceAmount" binding:"min=-99999999999,max=99999999999"`
	DestinationAmount    int64                          `json:"destinationAmount" binding:"min=-99999999999,max=99999999999"`
	HideAmount           bool                           `json:"hideAmount"`
	TagIds               []string                       `json:"tagIds"`
	PictureIds           []string                       `json:"pictureIds"`
	Comment              string                         `json:"comment" binding:"max=255"`
	GeoLocation          *TransactionGeoLocationRequest `json:"geoLocation" binding:"omitempty"`
}


type TransactionImportRequest struct {
	Transactions    []*TransactionCreateRequest `json:"transactions"`
	ClientSessionId string                      `json:"clientSessionId"`
}


type TransactionImportProcessRequest struct {
	ClientSessionId string `form:"client_session_id"`
}

type TransactionTagFilter struct {
	TagIds []int64
	Type   TransactionTagFilterType
}


type TransactionCountRequest struct {
	Type         TransactionType `form:"type" binding:"min=0,max=4"`
	CategoryIds  string          `form:"category_ids"`
	AccountIds   string          `form:"account_ids"`
	TagFilter    string          `form:"tag_filter" binding:"validTagFilter"`
	AmountFilter string          `form:"amount_filter" binding:"validAmountFilter"`
	Keyword      string          `form:"keyword"`
	MaxTime      int64           `form:"max_time" binding:"min=0"` 
	MinTime      int64           `form:"min_time" binding:"min=0"` 
}


type TransactionListByMaxTimeRequest struct {
	Type         TransactionType `form:"type" binding:"min=0,max=4"`
	CategoryIds  string          `form:"category_ids"`
	AccountIds   string          `form:"account_ids"`
	TagFilter    string          `form:"tag_filter" binding:"validTagFilter"`
	AmountFilter string          `form:"amount_filter" binding:"validAmountFilter"`
	Keyword      string          `form:"keyword"`
	MaxTime      int64           `form:"max_time" binding:"min=0"` 
	MinTime      int64           `form:"min_time" binding:"min=0"` 
	Page         int32           `form:"page" binding:"min=0"`
	Count        int32           `form:"count" binding:"required,min=1,max=50"`
	WithCount    bool            `form:"with_count"`
	WithPictures bool            `form:"with_pictures"`
	TrimAccount  bool            `form:"trim_account"`
	TrimCategory bool            `form:"trim_category"`
	TrimTag      bool            `form:"trim_tag"`
}


type TransactionListInMonthByPageRequest struct {
	Year         int32           `form:"year" binding:"required,min=1"`
	Month        int32           `form:"month" binding:"required,min=1"`
	Type         TransactionType `form:"type" binding:"min=0,max=4"`
	CategoryIds  string          `form:"category_ids"`
	AccountIds   string          `form:"account_ids"`
	TagFilter    string          `form:"tag_filter" binding:"validTagFilter"`
	AmountFilter string          `form:"amount_filter" binding:"validAmountFilter"`
	Keyword      string          `form:"keyword"`
	WithPictures bool            `form:"with_pictures"`
	TrimAccount  bool            `form:"trim_account"`
	TrimCategory bool            `form:"trim_category"`
	TrimTag      bool            `form:"trim_tag"`
}


type TransactionAllListRequest struct {
	Type         TransactionType `form:"type" binding:"min=0,max=4"`
	CategoryIds  string          `form:"category_ids"`
	AccountIds   string          `form:"account_ids"`
	TagFilter    string          `form:"tag_filter" binding:"validTagFilter"`
	AmountFilter string          `form:"amount_filter" binding:"validAmountFilter"`
	Keyword      string          `form:"keyword"`
	StartTime    int64           `form:"start_time" binding:"min=0"`
	EndTime      int64           `form:"end_time" binding:"min=0"`
	WithPictures bool            `form:"with_pictures"`
	TrimAccount  bool            `form:"trim_account"`
	TrimCategory bool            `form:"trim_category"`
	TrimTag      bool            `form:"trim_tag"`
}


type TransactionReconciliationStatementRequest struct {
	AccountId int64 `form:"account_id,string" binding:"required,min=1"`
	StartTime int64 `form:"start_time"`
	EndTime   int64 `form:"end_time"`
}


type TransactionStatisticRequest struct {
	StartTime              int64  `form:"start_time" binding:"min=0"`
	EndTime                int64  `form:"end_time" binding:"min=0"`
	TagFilter              string `form:"tag_filter" binding:"validTagFilter"`
	Keyword                string `form:"keyword"`
	UseTransactionTimezone bool   `form:"use_transaction_timezone"`
}


type TransactionStatisticTrendsRequest struct {
	YearMonthRangeRequest
	TagFilter              string `form:"tag_filter" binding:"validTagFilter"`
	Keyword                string `form:"keyword"`
	UseTransactionTimezone bool   `form:"use_transaction_timezone"`
}


type TransactionStatisticAssetTrendsRequest struct {
	StartTime int64 `form:"start_time"`
	EndTime   int64 `form:"end_time"`
}


type TransactionAmountsRequest struct {
	Query                  string `form:"query"`
	ExcludeAccountIds      string `form:"exclude_account_ids"`
	ExcludeCategoryIds     string `form:"exclude_category_ids"`
	UseTransactionTimezone bool   `form:"use_transaction_timezone"`
}


type TransactionAmountsRequestItem struct {
	Name      string
	StartTime int64
	EndTime   int64
}


type TransactionGetRequest struct {
	Id           int64 `form:"id,string" binding:"required,min=1"`
	WithPictures bool  `form:"with_pictures"`
	TrimAccount  bool  `form:"trim_account"`
	TrimCategory bool  `form:"trim_category"`
	TrimTag      bool  `form:"trim_tag"`
}


type TransactionMoveBetweenAccountsRequest struct {
	FromAccountId int64 `json:"fromAccountId,string" binding:"required,min=1"`
	ToAccountId   int64 `json:"toAccountId,string" binding:"required,min=1"`
}


type TransactionDeleteRequest struct {
	Id int64 `json:"id,string" binding:"required,min=1"`
}


type YearMonthRangeRequest struct {
	StartYearMonth string `form:"start_year_month"`
	EndYearMonth   string `form:"end_year_month"`
}


type TransactionGeoLocationResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}


type TransactionInfoResponse struct {
	Id                   int64                                    `json:"id,string"`
	TimeSequenceId       int64                                    `json:"timeSequenceId,string"`
	Type                 TransactionType                          `json:"type"`
	CategoryId           int64                                    `json:"categoryId,string"`
	Category             *TransactionCategoryInfoResponse         `json:"category,omitempty"`
	Time                 int64                                    `json:"time"`
	UtcOffset            int16                                    `json:"utcOffset"`
	SourceAccountId      int64                                    `json:"sourceAccountId,string"`
	SourceAccount        *AccountInfoResponse                     `json:"sourceAccount,omitempty"`
	DestinationAccountId int64                                    `json:"destinationAccountId,string,omitempty"`
	DestinationAccount   *AccountInfoResponse                     `json:"destinationAccount,omitempty"`
	SourceAmount         int64                                    `json:"sourceAmount"`
	DestinationAmount    int64                                    `json:"destinationAmount,omitempty"`
	HideAmount           bool                                     `json:"hideAmount"`
	TagIds               []string                                 `json:"tagIds"`
	Tags                 []*TransactionTagInfoResponse            `json:"tags,omitempty"`
	Pictures             TransactionPictureInfoBasicResponseSlice `json:"pictures,omitempty"`
	Comment              string                                   `json:"comment"`
	GeoLocation          *TransactionGeoLocationResponse          `json:"geoLocation,omitempty"`
	Editable             bool                                     `json:"editable"`
}


type TransactionCountResponse struct {
	TotalCount int64 `json:"totalCount"`
}


type TransactionInfoPageWrapperResponse struct {
	Items              TransactionInfoResponseSlice `json:"items"`
	NextTimeSequenceId *int64                       `json:"nextTimeSequenceId,string"`
	TotalCount         *int64                       `json:"totalCount,omitempty"`
}


type TransactionInfoPageWrapperResponse2 struct {
	Items      TransactionInfoResponseSlice `json:"items"`
	TotalCount int64                        `json:"totalCount"`
}


type TransactionReconciliationStatementResponseItem struct {
	*TransactionInfoResponse
	AccountOpeningBalance int64 `json:"accountOpeningBalance"`
	AccountClosingBalance int64 `json:"accountClosingBalance"`
}


type TransactionReconciliationStatementResponse struct {
	Transactions   []*TransactionReconciliationStatementResponseItem `json:"transactions"`
	TotalInflows   int64                                             `json:"totalInflows"`
	TotalOutflows  int64                                             `json:"totalOutflows"`
	OpeningBalance int64                                             `json:"openingBalance"`
	ClosingBalance int64                                             `json:"closingBalance"`
}


type TransactionStatisticResponse struct {
	StartTime int64                               `json:"startTime"`
	EndTime   int64                               `json:"endTime"`
	Items     []*TransactionStatisticResponseItem `json:"items"`
}


type TransactionStatisticResponseItem struct {
	CategoryId         int64                         `json:"categoryId,string"`
	AccountId          int64                         `json:"accountId,string"`
	RelatedAccountId   int64                         `json:"relatedAccountId,string,omitempty"`
	RelatedAccountType TransactionRelatedAccountType `json:"relatedAccountType,omitempty"`
	TotalAmount        int64                         `json:"amount"`
}


type TransactionStatisticTrendsResponseItem struct {
	Year  int32                               `json:"year"`
	Month int32                               `json:"month"`
	Items []*TransactionStatisticResponseItem `json:"items"`
}


type TransactionStatisticAssetTrendsResponseItem struct {
	Year  int32                                              `json:"year"`
	Month int32                                              `json:"month"`
	Day   int32                                              `json:"day"`
	Items []*TransactionStatisticAssetTrendsResponseDataItem `json:"items"`
}


type TransactionStatisticAssetTrendsResponseDataItem struct {
	AccountId             int64 `json:"accountId,string"`
	AccountOpeningBalance int64 `json:"accountOpeningBalance"`
	AccountClosingBalance int64 `json:"accountClosingBalance"`
}


type TransactionAmountsResponseItem struct {
	StartTime int64                                       `json:"startTime"`
	EndTime   int64                                       `json:"endTime"`
	Amounts   []*TransactionAmountsResponseItemAmountInfo `json:"amounts"`
}


type TransactionMonthAmountsResponseItem struct {
	Year    int32                                       `json:"year"`
	Month   int32                                       `json:"month"`
	Amounts []*TransactionAmountsResponseItemAmountInfo `json:"amounts"`
}


type TransactionAmountsResponseItemAmountInfo struct {
	Currency      string `json:"currency"`
	IncomeAmount  int64  `json:"incomeAmount"`
	ExpenseAmount int64  `json:"expenseAmount"`
}


func ParseTransactionTagFilter(tagFilterStr string) ([]*TransactionTagFilter, error) {
	if tagFilterStr == "" || tagFilterStr == TransactionNoTagFilterValue {
		return []*TransactionTagFilter{}, nil
	}

	filters := strings.Split(tagFilterStr, ";")
	transactionTagFilters := make([]*TransactionTagFilter, 0, len(filters))

	for _, filter := range filters {
		tagFilterItem := strings.Split(filter, ":")

		if len(tagFilterItem) != 2 {
			return nil, errs.ErrFormatInvalid
		}

		tagFilterType, err := utils.StringToInt(tagFilterItem[0])

		if err != nil || (tagFilterType < int(TRANSACTION_TAG_FILTER_HAS_ANY) || tagFilterType > int(TRANSACTION_TAG_FILTER_NOT_HAS_ALL)) {
			return nil, errs.ErrFormatInvalid
		}

		textualTagIds := strings.Split(tagFilterItem[1], ",")
		tagIds := make([]int64, 0, len(textualTagIds))

		for _, tagIdStr := range textualTagIds {
			tagId, err := utils.StringToInt64(tagIdStr)

			if err != nil {
				return nil, errs.ErrTransactionTagIdInvalid
			}

			tagIds = append(tagIds, tagId)
		}

		transactionTagFilter := &TransactionTagFilter{
			TagIds: tagIds,
			Type:   TransactionTagFilterType(tagFilterType),
		}

		transactionTagFilters = append(transactionTagFilters, transactionTagFilter)
	}

	return transactionTagFilters, nil
}


func (t *Transaction) IsEditable(currentUser *User, clientTimezone *time.Location, account *Account, relatedAccount *Account) bool {
	if currentUser == nil || !currentUser.CanEditTransactionByTransactionTime(t.TransactionTime, clientTimezone) {
		return false
	}

	if account == nil || account.Hidden {
		return false
	}

	if t.Type == TRANSACTION_DB_TYPE_TRANSFER_OUT {
		if relatedAccount == nil || relatedAccount.Hidden {
			return false
		}
	}

	return true
}


func (t *Transaction) ToTransactionInfoResponse(tagIds []int64, editable bool) *TransactionInfoResponse {
	transactionType, err := t.Type.ToTransactionType()

	if err != nil {
		return nil
	}

	sourceAccountId := t.AccountId
	sourceAmount := t.Amount

	destinationAccountId := int64(0)
	destinationAmount := int64(0)

	if t.Type == TRANSACTION_DB_TYPE_TRANSFER_OUT {
		destinationAccountId = t.RelatedAccountId
		destinationAmount = t.RelatedAccountAmount
	} else if t.Type == TRANSACTION_DB_TYPE_TRANSFER_IN {
		sourceAccountId = t.RelatedAccountId
		sourceAmount = t.RelatedAccountAmount

		destinationAccountId = t.AccountId
		destinationAmount = t.Amount
	}

	geoLocation := &TransactionGeoLocationResponse{}

	if t.GeoLongitude != 0 || t.GeoLatitude != 0 {
		geoLocation.Longitude = t.GeoLongitude
		geoLocation.Latitude = t.GeoLatitude
	} else {
		geoLocation = nil
	}

	return &TransactionInfoResponse{
		Id:                   t.TransactionId,
		TimeSequenceId:       t.TransactionTime,
		Type:                 transactionType,
		CategoryId:           t.CategoryId,
		Time:                 utils.GetUnixTimeFromTransactionTime(t.TransactionTime),
		UtcOffset:            t.TimezoneUtcOffset,
		SourceAccountId:      sourceAccountId,
		DestinationAccountId: destinationAccountId,
		SourceAmount:         sourceAmount,
		DestinationAmount:    destinationAmount,
		HideAmount:           t.HideAmount,
		TagIds:               utils.Int64ArrayToStringArray(tagIds),
		Comment:              t.Comment,
		GeoLocation:          geoLocation,
		Editable:             editable,
	}
}


func (t *TransactionAmountsRequest) GetTransactionAmountsRequestItems() ([]*TransactionAmountsRequestItem, error) {
	items := strings.Split(t.Query, "|")
	requestItems := make([]*TransactionAmountsRequestItem, 0, len(items))

	for i := 0; i < len(items); i++ {
		itemValues := strings.Split(items[i], "_")

		if len(itemValues) != 3 {
			return nil, errs.ErrQueryItemsInvalid
		}

		startTime, err := utils.StringToInt64(itemValues[1])

		if err != nil {
			return nil, err
		}

		endTime, err := utils.StringToInt64(itemValues[2])

		if err != nil {
			return nil, err
		}

		requestItem := &TransactionAmountsRequestItem{
			Name:      itemValues[0],
			StartTime: startTime,
			EndTime:   endTime,
		}

		requestItems = append(requestItems, requestItem)
	}

	return requestItems, nil
}


func (t *YearMonthRangeRequest) GetNumericYearMonthRange() (int32, int32, int32, int32, error) {
	var startYear, startMonth, endYear, endMonth int32
	var err error

	if t.StartYearMonth != "" {
		startYear, startMonth, err = utils.ParseNumericYearMonth(t.StartYearMonth)

		if err != nil {
			return 0, 0, 0, 0, err
		}
	}

	if t.EndYearMonth != "" {
		endYear, endMonth, err = utils.ParseNumericYearMonth(t.EndYearMonth)

		if err != nil {
			return 0, 0, 0, 0, err
		}
	}

	return startYear, startMonth, endYear, endMonth, nil
}


type TransactionInfoResponseSlice []*TransactionInfoResponse


func (s TransactionInfoResponseSlice) Len() int {
	return len(s)
}


func (s TransactionInfoResponseSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}


func (s TransactionInfoResponseSlice) Less(i, j int) bool {
	if s[i].Time != s[j].Time {
		return s[i].Time > s[j].Time
	}

	return s[i].Id > s[j].Id
}


type TransactionStatisticTrendsResponseItemSlice []*TransactionStatisticTrendsResponseItem


func (s TransactionStatisticTrendsResponseItemSlice) Len() int {
	return len(s)
}


func (s TransactionStatisticTrendsResponseItemSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}


func (s TransactionStatisticTrendsResponseItemSlice) Less(i, j int) bool {
	if s[i].Year != s[j].Year {
		return s[i].Year < s[j].Year
	}

	return s[i].Month < s[j].Month
}


type TransactionStatisticAssetTrendsResponseItemSlice []*TransactionStatisticAssetTrendsResponseItem


func (s TransactionStatisticAssetTrendsResponseItemSlice) Len() int {
	return len(s)
}


func (s TransactionStatisticAssetTrendsResponseItemSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}


func (s TransactionStatisticAssetTrendsResponseItemSlice) Less(i, j int) bool {
	if s[i].Year != s[j].Year {
		return s[i].Year < s[j].Year
	}

	if s[i].Month != s[j].Month {
		return s[i].Month < s[j].Month
	}

	return s[i].Day < s[j].Day
}


type TransactionAmountsResponseItemAmountInfoSlice []*TransactionAmountsResponseItemAmountInfo


func (s TransactionAmountsResponseItemAmountInfoSlice) Len() int {
	return len(s)
}


func (s TransactionAmountsResponseItemAmountInfoSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}


func (s TransactionAmountsResponseItemAmountInfoSlice) Less(i, j int) bool {
	return strings.Compare(s[i].Currency, s[j].Currency) < 0
}
