package models


type TransactionTag struct {
	TagId           int64  `xorm:"PK"`
	Uid             int64  `xorm:"INDEX(IDX_tag_uid_deleted_group_order) NOT NULL"`
	Deleted         bool   `xorm:"INDEX(IDX_tag_uid_deleted_group_order) NOT NULL"`
	TagGroupId      int64  `xorm:"INDEX(IDX_tag_uid_deleted_group_order) NOT NULL DEFAULT 0"`
	Name            string `xorm:"VARCHAR(64) NOT NULL"`
	DisplayOrder    int32  `xorm:"INDEX(IDX_tag_uid_deleted_group_order) NOT NULL"`
	Hidden          bool   `xorm:"NOT NULL"`
	CreatedUnixTime int64
	UpdatedUnixTime int64
	DeletedUnixTime int64
}


type TransactionTagGetRequest struct {
	Id int64 `form:"id,string" binding:"required,min=1"`
}


type TransactionTagCreateRequest struct {
	GroupId int64  `json:"groupId,string"`
	Name    string `json:"name" binding:"required,notBlank,max=64"`
}


type TransactionTagCreateBatchRequest struct {
	Tags       []*TransactionTagCreateRequest `json:"tags" binding:"required"`
	GroupId    int64                          `json:"groupId,string"`
	SkipExists bool                           `json:"skipExists"`
}


type TransactionTagModifyRequest struct {
	Id      int64  `json:"id,string" binding:"required,min=1"`
	GroupId int64  `json:"groupId,string"`
	Name    string `json:"name" binding:"required,notBlank,max=64"`
}


type TransactionTagHideRequest struct {
	Id     int64 `json:"id,string" binding:"required,min=1"`
	Hidden bool  `json:"hidden"`
}


type TransactionTagMoveRequest struct {
	NewDisplayOrders []*TransactionTagNewDisplayOrderRequest `json:"newDisplayOrders" binding:"required,min=1"`
}


type TransactionTagNewDisplayOrderRequest struct {
	Id           int64 `json:"id,string" binding:"required,min=1"`
	DisplayOrder int32 `json:"displayOrder"`
}


type TransactionTagDeleteRequest struct {
	Id int64 `json:"id,string" binding:"required,min=1"`
}


type TransactionTagInfoResponse struct {
	Id           int64  `json:"id,string"`
	Name         string `json:"name"`
	TagGroupId   int64  `json:"groupId,string"`
	DisplayOrder int32  `json:"displayOrder"`
	Hidden       bool   `json:"hidden"`
}


func (t *TransactionTag) FillFromOtherTag(tag *TransactionTag) {
	t.TagId = tag.TagId
	t.Uid = tag.Uid
	t.Deleted = tag.Deleted
	t.Name = tag.Name
	t.TagGroupId = tag.TagGroupId
	t.DisplayOrder = tag.DisplayOrder
	t.Hidden = tag.Hidden
	t.CreatedUnixTime = tag.CreatedUnixTime
	t.UpdatedUnixTime = tag.UpdatedUnixTime
	t.DeletedUnixTime = tag.DeletedUnixTime
}


func (t *TransactionTag) ToTransactionTagInfoResponse() *TransactionTagInfoResponse {
	return &TransactionTagInfoResponse{
		Id:           t.TagId,
		Name:         t.Name,
		TagGroupId:   t.TagGroupId,
		DisplayOrder: t.DisplayOrder,
		Hidden:       t.Hidden,
	}
}


type TransactionTagInfoResponseSlice []*TransactionTagInfoResponse


func (s TransactionTagInfoResponseSlice) Len() int {
	return len(s)
}


func (s TransactionTagInfoResponseSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}


func (s TransactionTagInfoResponseSlice) Less(i, j int) bool {
	if s[i].TagGroupId != s[j].TagGroupId {
		return s[i].TagGroupId < s[j].TagGroupId
	}

	return s[i].DisplayOrder < s[j].DisplayOrder
}
