package services

import (
	"fmt"
	"path/filepath"

	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/datastore"
	"github.com/Shavitjnr/split-chill-ai/pkg/mail"
	"github.com/Shavitjnr/split-chill-ai/pkg/settings"
	"github.com/Shavitjnr/split-chill-ai/pkg/storage"
	"github.com/Shavitjnr/split-chill-ai/pkg/utils"
	"github.com/Shavitjnr/split-chill-ai/pkg/uuid"
)


type ServiceUsingDB struct {
	container *datastore.DataStoreContainer
}


func (s *ServiceUsingDB) UserDB() *datastore.Database {
	return s.container.UserStore.Choose(0)
}


func (s *ServiceUsingDB) TokenDB(uid int64) *datastore.Database {
	return s.container.TokenStore.Choose(uid)
}


func (s *ServiceUsingDB) TokenDBByIndex(index int) *datastore.Database {
	return s.container.TokenStore.Get(index)
}


func (s *ServiceUsingDB) TokenDBCount() int {
	return s.container.TokenStore.Count()
}


func (s *ServiceUsingDB) UserDataDB(uid int64) *datastore.Database {
	return s.container.UserDataStore.Choose(uid)
}


func (s *ServiceUsingDB) UserDataDBByIndex(index int) *datastore.Database {
	return s.container.UserDataStore.Get(index)
}


func (s *ServiceUsingDB) UserDataDBCount() int {
	return s.container.UserDataStore.Count()
}


type ServiceUsingConfig struct {
	container *settings.ConfigContainer
}


func (s *ServiceUsingConfig) CurrentConfig() *settings.Config {
	return s.container.GetCurrentConfig()
}


type ServiceUsingMailer struct {
	container *mail.MailerContainer
}


func (s *ServiceUsingMailer) SendMail(message *mail.MailMessage) error {
	return s.container.SendMail(message)
}


type ServiceUsingUuid struct {
	container *uuid.UuidContainer
}


func (s *ServiceUsingUuid) GenerateUuid(uuidType uuid.UuidType) int64 {
	return s.container.GenerateUuid(uuidType)
}


func (s *ServiceUsingUuid) GenerateUuids(uuidType uuid.UuidType, count uint16) []int64 {
	return s.container.GenerateUuids(uuidType, count)
}


type ServiceUsingStorage struct {
	container *storage.StorageContainer
}


func (s *ServiceUsingStorage) ExistsAvatar(ctx core.Context, uid int64, fileExtension string) (bool, error) {
	return s.container.ExistsAvatar(ctx, s.getUserAvatarPath(uid, fileExtension))
}


func (s *ServiceUsingStorage) ReadAvatar(ctx core.Context, uid int64, fileExtension string) (storage.ObjectInStorage, error) {
	return s.container.ReadAvatar(ctx, s.getUserAvatarPath(uid, fileExtension))
}


func (s *ServiceUsingStorage) SaveAvatar(ctx core.Context, uid int64, object storage.ObjectInStorage, fileExtension string) error {
	return s.container.SaveAvatar(ctx, s.getUserAvatarPath(uid, fileExtension), object)
}


func (s *ServiceUsingStorage) DeleteAvatar(ctx core.Context, uid int64, fileExtension string) error {
	return s.container.DeleteAvatar(ctx, s.getUserAvatarPath(uid, fileExtension))
}


func (s *ServiceUsingStorage) ExistsTransactionPicture(ctx core.Context, uid int64, pictureId int64, fileExtension string) (bool, error) {
	return s.container.ExistsTransactionPicture(ctx, s.getTransactionPicturePath(uid, pictureId, fileExtension))
}


func (s *ServiceUsingStorage) ReadTransactionPicture(ctx core.Context, uid int64, pictureId int64, fileExtension string) (storage.ObjectInStorage, error) {
	return s.container.ReadTransactionPicture(ctx, s.getTransactionPicturePath(uid, pictureId, fileExtension))
}


func (s *ServiceUsingStorage) SaveTransactionPicture(ctx core.Context, uid int64, pictureId int64, object storage.ObjectInStorage, fileExtension string) error {
	return s.container.SaveTransactionPicture(ctx, s.getTransactionPicturePath(uid, pictureId, fileExtension), object)
}


func (s *ServiceUsingStorage) DeleteTransactionPicture(ctx core.Context, uid int64, pictureId int64, fileExtension string) error {
	return s.container.DeleteTransactionPicture(ctx, s.getTransactionPicturePath(uid, pictureId, fileExtension))
}

func (s *ServiceUsingStorage) getUserAvatarPath(uid int64, fileExtension string) string {
	return fmt.Sprintf("%d.%s", uid, fileExtension)
}

func (s *ServiceUsingStorage) getTransactionPicturePath(uid int64, pictureId int64, fileExtension string) string {
	return filepath.Join(utils.Int64ToString(uid), fmt.Sprintf("%d.%s", pictureId, fileExtension))
}
