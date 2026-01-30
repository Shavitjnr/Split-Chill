package models

const TransactionPictureNewPictureTransactionId = int64(0)


type TransactionPictureInfo struct {
	Uid              int64  `xorm:"INDEX(IDX_transaction_picture_uid_deleted_transaction_id_picture_id) INDEX(IDX_transaction_picture_uid_deleted_picture_id) NOT NULL"`
	Deleted          bool   `xorm:"INDEX(IDX_transaction_picture_uid_deleted_transaction_id_picture_id) INDEX(IDX_transaction_picture_uid_deleted_picture_id) NOT NULL"`
	TransactionId    int64  `xorm:"INDEX(IDX_transaction_picture_uid_deleted_transaction_id_picture_id) NOT NULL"`
	PictureId        int64  `xorm:"PK INDEX(IDX_transaction_picture_uid_deleted_transaction_id_picture_id) INDEX(IDX_transaction_picture_uid_deleted_picture_id)"`
	PictureExtension string `xorm:"VARCHAR(10) NOT NULL"`
	CreatedIp        string `xorm:"VARCHAR(39)"`
	CreatedUnixTime  int64
	UpdatedUnixTime  int64
	DeletedUnixTime  int64
}


type TransactionPictureUnusedDeleteRequest struct {
	Id int64 `json:"id,string" binding:"required,min=1"`
}


type TransactionPictureInfoBasicResponse struct {
	PictureId   int64  `json:"pictureId,string"`
	OriginalUrl string `json:"originalUrl"`
}


func (p *TransactionPictureInfo) ToTransactionPictureInfoBasicResponse(originalUrl string) *TransactionPictureInfoBasicResponse {
	return &TransactionPictureInfoBasicResponse{
		PictureId:   p.PictureId,
		OriginalUrl: originalUrl,
	}
}


type TransactionPictureInfoBasicResponseSlice []*TransactionPictureInfoBasicResponse


func (s TransactionPictureInfoBasicResponseSlice) Len() int {
	return len(s)
}


func (s TransactionPictureInfoBasicResponseSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}


func (s TransactionPictureInfoBasicResponseSlice) Less(i, j int) bool {
	return s[i].PictureId < s[j].PictureId
}
