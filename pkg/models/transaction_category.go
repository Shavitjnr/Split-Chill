package models


const LevelOneTransactionCategoryParentId = 0


type TransactionCategoryType byte


const (
	CATEGORY_TYPE_INCOME   TransactionCategoryType = 1
	CATEGORY_TYPE_EXPENSE  TransactionCategoryType = 2
	CATEGORY_TYPE_TRANSFER TransactionCategoryType = 3
)


type TransactionCategory struct {
	CategoryId       int64                   `xorm:"PK"`
	Uid              int64                   `xorm:"INDEX(IDX_category_uid_deleted_type_parent_category_id_order) NOT NULL"`
	Deleted          bool                    `xorm:"INDEX(IDX_category_uid_deleted_type_parent_category_id_order) NOT NULL"`
	Type             TransactionCategoryType `xorm:"INDEX(IDX_category_uid_deleted_type_parent_category_id_order) NOT NULL"`
	ParentCategoryId int64                   `xorm:"INDEX(IDX_category_uid_deleted_type_parent_category_id_order) NOT NULL"`
	Name             string                  `xorm:"VARCHAR(64) NOT NULL"`
	DisplayOrder     int32                   `xorm:"INDEX(IDX_category_uid_deleted_type_parent_category_id_order) NOT NULL"`
	Icon             int64                   `xorm:"NOT NULL"`
	Color            string                  `xorm:"VARCHAR(6) NOT NULL"`
	Hidden           bool                    `xorm:"NOT NULL"`
	Comment          string                  `xorm:"VARCHAR(255) NOT NULL"`
	CreatedUnixTime  int64
	UpdatedUnixTime  int64
	DeletedUnixTime  int64
}


type TransactionCategoryListRequest struct {
	Type     TransactionCategoryType `form:"type" binding:"min=0"`
	ParentId int64                   `form:"parent_id,string,default=-1" binding:"min=-1"`
}


type TransactionCategoryGetRequest struct {
	Id int64 `form:"id,string" binding:"required,min=1"`
}


type TransactionCategoryCreateRequest struct {
	Name            string                  `json:"name" binding:"required,notBlank,max=64"`
	Type            TransactionCategoryType `json:"type" binding:"required"`
	ParentId        int64                   `json:"parentId,string" binding:"min=0"`
	Icon            int64                   `json:"icon,string" binding:"min=1"`
	Color           string                  `json:"color" binding:"required,len=6,validHexRGBColor"`
	Comment         string                  `json:"comment" binding:"max=255"`
	ClientSessionId string                  `json:"clientSessionId"`
}


type TransactionCategoryCreateBatchRequest struct {
	Categories []*TransactionCategoryCreateWithSubCategories `json:"categories" binding:"required"`
}


type TransactionCategoryCreateWithSubCategories struct {
	Name          string                              `json:"name" binding:"required,notBlank,max=64"`
	Type          TransactionCategoryType             `json:"type" binding:"required"`
	Icon          int64                               `json:"icon,string" binding:"min=1"`
	Color         string                              `json:"color" binding:"required,len=6,validHexRGBColor"`
	Comment       string                              `json:"comment" binding:"max=255"`
	SubCategories []*TransactionCategoryCreateRequest `json:"subCategories" binding:"required"`
}


type TransactionCategoryModifyRequest struct {
	Id       int64  `json:"id,string" binding:"required,min=1"`
	Name     string `json:"name" binding:"required,notBlank,max=64"`
	ParentId int64  `json:"parentId,string" binding:"min=0"`
	Icon     int64  `json:"icon,string" binding:"min=1"`
	Color    string `json:"color" binding:"required,len=6,validHexRGBColor"`
	Comment  string `json:"comment" binding:"max=255"`
	Hidden   bool   `json:"hidden"`
}


type TransactionCategoryHideRequest struct {
	Id     int64 `json:"id,string" binding:"required,min=1"`
	Hidden bool  `json:"hidden"`
}


type TransactionCategoryMoveRequest struct {
	NewDisplayOrders []*TransactionCategoryNewDisplayOrderRequest `json:"newDisplayOrders" binding:"required,min=1"`
}


type TransactionCategoryNewDisplayOrderRequest struct {
	Id           int64 `json:"id,string" binding:"required,min=1"`
	DisplayOrder int32 `json:"displayOrder"`
}


type TransactionCategoryDeleteRequest struct {
	Id int64 `json:"id,string" binding:"required,min=1"`
}


type TransactionCategoryInfoResponse struct {
	Id            int64                                `json:"id,string"`
	Name          string                               `json:"name"`
	ParentId      int64                                `json:"parentId,string"`
	Type          TransactionCategoryType              `json:"type"`
	Icon          int64                                `json:"icon,string"`
	Color         string                               `json:"color"`
	Comment       string                               `json:"comment"`
	DisplayOrder  int32                                `json:"displayOrder"`
	Hidden        bool                                 `json:"hidden"`
	SubCategories TransactionCategoryInfoResponseSlice `json:"subCategories,omitempty"`
}


func (c *TransactionCategory) ToTransactionCategoryInfoResponse() *TransactionCategoryInfoResponse {
	return &TransactionCategoryInfoResponse{
		Id:           c.CategoryId,
		Name:         c.Name,
		ParentId:     c.ParentCategoryId,
		Type:         c.Type,
		Icon:         c.Icon,
		Color:        c.Color,
		Comment:      c.Comment,
		DisplayOrder: c.DisplayOrder,
		Hidden:       c.Hidden,
	}
}


type TransactionCategoryInfoResponseSlice []*TransactionCategoryInfoResponse


func (s TransactionCategoryInfoResponseSlice) Len() int {
	return len(s)
}


func (s TransactionCategoryInfoResponseSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}


func (s TransactionCategoryInfoResponseSlice) Less(i, j int) bool {
	return s[i].DisplayOrder < s[j].DisplayOrder
}
