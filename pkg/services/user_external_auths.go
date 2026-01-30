package services

import (
	"time"

	"xorm.io/xorm"

	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/datastore"
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
	"github.com/Shavitjnr/split-chill-ai/pkg/models"
)


type UserExternalAuthService struct {
	ServiceUsingDB
}


var (
	UserExternalAuths = &UserExternalAuthService{
		ServiceUsingDB: ServiceUsingDB{
			container: datastore.Container,
		},
	}
)


func (s *UserExternalAuthService) GetUserAllExternalAuthsByUid(c core.Context, uid int64) ([]*models.UserExternalAuth, error) {
	if uid <= 0 {
		return nil, errs.ErrUserIdInvalid
	}

	var userExternalAuths []*models.UserExternalAuth
	err := s.UserDB().NewSession(c).Where("uid=?", uid).Find(&userExternalAuths)

	return userExternalAuths, err
}


func (s *UserExternalAuthService) GetUserExternalAuthByUid(c core.Context, uid int64, externalAuthType core.UserExternalAuthType) (*models.UserExternalAuth, error) {
	if uid <= 0 {
		return nil, errs.ErrUserIdInvalid
	}

	userExternalAuth := &models.UserExternalAuth{}
	has, err := s.UserDB().NewSession(c).Where("uid=? AND external_auth_type=?", uid, externalAuthType).Get(userExternalAuth)

	if err != nil {
		return nil, err
	} else if !has {
		return nil, errs.ErrUserExternalAuthNotFound
	}

	return userExternalAuth, err
}


func (s *UserExternalAuthService) GetUserExternalAuthByExternalUserName(c core.Context, externalUserName string, externalAuthType core.UserExternalAuthType) (*models.UserExternalAuth, error) {
	userExternalAuth := &models.UserExternalAuth{}
	has, err := s.UserDB().NewSession(c).Where("external_auth_type=? AND external_username=?", externalAuthType, externalUserName).Get(userExternalAuth)

	if err != nil {
		return nil, err
	} else if !has {
		return nil, errs.ErrUserExternalAuthNotFound
	}

	return userExternalAuth, err
}


func (s *UserExternalAuthService) GetUserExternalAuthByExternalEmail(c core.Context, externalEmail string, externalAuthType core.UserExternalAuthType) (*models.UserExternalAuth, error) {
	userExternalAuth := &models.UserExternalAuth{}
	has, err := s.UserDB().NewSession(c).Where("external_auth_type=? AND external_email=?", externalAuthType, externalEmail).Get(userExternalAuth)

	if err != nil {
		return nil, err
	} else if !has {
		return nil, errs.ErrUserExternalAuthNotFound
	}

	return userExternalAuth, err
}


func (s *UserExternalAuthService) CreateUserExternalAuth(c core.Context, userExternalAuth *models.UserExternalAuth) error {
	if userExternalAuth.Uid <= 0 {
		return errs.ErrUserIdInvalid
	}

	userExternalAuth.CreatedUnixTime = time.Now().Unix()

	return s.UserDB().DoTransaction(c, func(sess *xorm.Session) error {
		exists, err := sess.Where("uid=? AND external_auth_type=?", userExternalAuth.Uid, userExternalAuth.ExternalAuthType).Limit(1).Exist(&models.UserExternalAuth{})

		if err != nil {
			return err
		} else if exists {
			return errs.ErrUserExternalAuthAlreadyExists
		}

		_, err = sess.Insert(userExternalAuth)
		return err
	})
}


func (s *UserExternalAuthService) DeleteUserExternalAuth(c core.Context, uid int64, externalAuthType core.UserExternalAuthType) error {
	if uid <= 0 {
		return errs.ErrUserIdInvalid
	}

	return s.UserDB().DoTransaction(c, func(sess *xorm.Session) error {
		deletedRows, err := sess.Where("uid=? AND external_auth_type=?", uid, externalAuthType).Delete(&models.UserExternalAuth{})

		if err != nil {
			return err
		} else if deletedRows < 1 {
			return errs.ErrUserExternalAuthNotFound
		}

		return nil
	})
}
