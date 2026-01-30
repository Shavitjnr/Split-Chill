package services

import (
	"sort"
	"time"

	"xorm.io/xorm"

	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/datastore"
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
	"github.com/Shavitjnr/split-chill-ai/pkg/log"
	"github.com/Shavitjnr/split-chill-ai/pkg/models"
)


type UserApplicationCloudSettingsService struct {
	ServiceUsingDB
}


var (
	UserApplicationCloudSettings = &UserApplicationCloudSettingsService{
		ServiceUsingDB: ServiceUsingDB{
			container: datastore.Container,
		},
	}
)


func (s *UserApplicationCloudSettingsService) GetUserApplicationCloudSettingsByUid(c core.Context, uid int64) (*models.UserApplicationCloudSetting, error) {
	if uid <= 0 {
		return nil, errs.ErrUserIdInvalid
	}

	applicationCloudSetting := &models.UserApplicationCloudSetting{}
	has, err := s.UserDB().NewSession(c).ID(uid).Get(applicationCloudSetting)

	if err != nil {
		return nil, err
	} else if !has {
		return nil, nil
	}

	return applicationCloudSetting, nil
}


func (s *UserApplicationCloudSettingsService) UpdateUserApplicationCloudSettings(c core.Context, uid int64, settings models.ApplicationCloudSettingSlice, forceUpdate bool, lastUpdateTime int64) error {
	if uid <= 0 {
		return errs.ErrUserIdInvalid
	}

	sort.Sort(settings)

	userApplicationCloudSetting := &models.UserApplicationCloudSetting{
		Uid:             uid,
		Settings:        settings,
		UpdatedUnixTime: time.Now().Unix(),
	}

	return s.UserDB().DoTransaction(c, func(sess *xorm.Session) error {
		exists, err := sess.Cols("uid").Where("uid=?", uid).Exist(&models.UserApplicationCloudSetting{})

		if err != nil {
			return err
		}

		updatedRows := int64(0)

		if !exists {
			updatedRows, err = sess.Insert(userApplicationCloudSetting)
		} else if forceUpdate || lastUpdateTime <= 0 {
			updatedRows, err = sess.ID(uid).Cols("settings", "updated_unix_time").Update(userApplicationCloudSetting)
		} else {
			updatedRows, err = sess.ID(uid).Cols("settings", "updated_unix_time").Where("updated_unix_time=?", lastUpdateTime).Update(userApplicationCloudSetting)
		}

		if err != nil {
			return err
		} else if updatedRows < 1 {
			log.Errorf(c, "[user_app_cloud_settings.UpdateUserApplicationCloudSettings] failed to update user application cloud settings")
			return errs.ErrDatabaseOperationFailed
		}

		return nil
	})
}


func (s *UserApplicationCloudSettingsService) ClearUserApplicationCloudSettings(c core.Context, uid int64) error {
	if uid <= 0 {
		return errs.ErrUserIdInvalid
	}

	return s.UserDB().DoTransaction(c, func(sess *xorm.Session) error {
		_, err := sess.Where("uid=?", uid).Delete(&models.UserApplicationCloudSetting{})
		return err
	})
}
