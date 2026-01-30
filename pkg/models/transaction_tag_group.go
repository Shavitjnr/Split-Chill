package models


type TransactionTagGroup struct {
	TagGroupId      int64  `xorm:"PK"`
	Uid             int64  `xorm:"INDEX(IDX_tag_group_uid_deleted_order) NOT NULL"`
	Deleted         bool   `xorm:"INDEX(IDX_tag_group_uid_deleted_order) NOT NULL"`
	Name            string `xorm:"VARCHAR(64) NOT NULL"`
	DisplayOrder    int32  `xorm:"INDEX(IDX_tag_group_uid_deleted_order) NOT NULL"`
	CreatedUnixTime int64
	UpdatedUnixTime int64
	DeletedUnixTime int64
}


type TransactionTagGroupGetRequest struct {
	Id int64 `form:"id,string" binding:"required,min=1"`
}


type TransactionTagGroupCreateRequest struct {
	Name string `json:"name" binding:"required,notBlank,max=64"`
}


type TransactionTagGroupModifyRequest struct {
	Id   int64  `json:"id,string" binding:"required,min=1"`
	Name string `json:"name" binding:"required,notBlank,max=64"`
}


type TransactionTagGroupMoveRequest struct {
	NewDisplayOrders []*TransactionTagGroupNewDisplayOrderRequest `json:"newDisplayOrders" binding:"required,min=1"`
}


type TransactionTagGroupNewDisplayOrderRequest struct {
	Id           int64 `json:"id,string" binding:"required,min=1"`
	DisplayOrder int32 `json:"displayOrder"`
}


type TransactionTagGroupDeleteRequest struct {
	Id int64 `json:"id,string" binding:"required,min=1"`
}


type TransactionTagGroupInfoResponse struct {
	Id           int64  `json:"id,string"`
	Name         string `json:"name"`
	DisplayOrder int32  `json:"displayOrder"`
}


func (t *TransactionTagGroup) ToTransactionTagGroupInfoResponse() *TransactionTagGroupInfoResponse {
	return &TransactionTagGroupInfoResponse{
		Id:           t.TagGroupId,
		Name:         t.Name,
		DisplayOrder: t.DisplayOrder,
	}
}


type TransactionTagGroupInfoResponseSlice []*TransactionTagGroupInfoResponse


func (s TransactionTagGroupInfoResponseSlice) Len() int {
	return len(s)
}


func (s TransactionTagGroupInfoResponseSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}


func (s TransactionTagGroupInfoResponseSlice) Less(i, j int) bool {
	return s[i].DisplayOrder < s[j].DisplayOrder
}
