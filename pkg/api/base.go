package api

import (
	"fmt"
	"sort"
	"time"

	"github.com/Shavitjnr/split-chill-ai/pkg/avatars"
	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/duplicatechecker"
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
	"github.com/Shavitjnr/split-chill-ai/pkg/log"
	"github.com/Shavitjnr/split-chill-ai/pkg/models"
	"github.com/Shavitjnr/split-chill-ai/pkg/settings"
	"github.com/Shavitjnr/split-chill-ai/pkg/utils"
)

const internalTransactionPictureUrlFormat = "%spictures/%d.%s"


type ApiUsingConfig struct {
	container *settings.ConfigContainer
}


func (a *ApiUsingConfig) CurrentConfig() *settings.Config {
	return a.container.GetCurrentConfig()
}


func (a *ApiUsingConfig) GetTransactionPictureInfoResponse(pictureInfo *models.TransactionPictureInfo) *models.TransactionPictureInfoBasicResponse {
	originalUrl := fmt.Sprintf(internalTransactionPictureUrlFormat, a.CurrentConfig().RootUrl, pictureInfo.PictureId, pictureInfo.PictureExtension)
	return pictureInfo.ToTransactionPictureInfoBasicResponse(originalUrl)
}


func (a *ApiUsingConfig) GetTransactionPictureInfoResponseList(pictureInfos []*models.TransactionPictureInfo) models.TransactionPictureInfoBasicResponseSlice {
	pictureInfoResps := make(models.TransactionPictureInfoBasicResponseSlice, len(pictureInfos))

	for i := 0; i < len(pictureInfos); i++ {
		pictureInfoResps[i] = a.GetTransactionPictureInfoResponse(pictureInfos[i])
	}

	sort.Sort(pictureInfoResps)

	return pictureInfoResps
}


func (a *ApiUsingConfig) GetAfterRegisterNotificationContent(userLanguage string, clientLanguage string) string {
	language := userLanguage

	if language == "" {
		language = clientLanguage
	}

	if !a.CurrentConfig().AfterRegisterNotification.Enabled {
		return ""
	}

	if multiLanguageContent, exists := a.CurrentConfig().AfterRegisterNotification.MultiLanguageContent[language]; exists {
		return multiLanguageContent
	}

	return a.CurrentConfig().AfterRegisterNotification.DefaultContent
}


func (a *ApiUsingConfig) GetAfterLoginNotificationContent(userLanguage string, clientLanguage string) string {
	language := userLanguage

	if language == "" {
		language = clientLanguage
	}

	if !a.CurrentConfig().AfterLoginNotification.Enabled {
		return ""
	}

	if multiLanguageContent, exists := a.CurrentConfig().AfterLoginNotification.MultiLanguageContent[language]; exists {
		return multiLanguageContent
	}

	return a.CurrentConfig().AfterLoginNotification.DefaultContent
}


func (a *ApiUsingConfig) GetAfterOpenNotificationContent(userLanguage string, clientLanguage string) string {
	language := userLanguage

	if language == "" {
		language = clientLanguage
	}

	if !a.CurrentConfig().AfterOpenNotification.Enabled {
		return ""
	}

	if multiLanguageContent, exists := a.CurrentConfig().AfterOpenNotification.MultiLanguageContent[language]; exists {
		return multiLanguageContent
	}

	return a.CurrentConfig().AfterOpenNotification.DefaultContent
}


type ApiUsingDuplicateChecker struct {
	ApiUsingConfig
	container *duplicatechecker.DuplicateCheckerContainer
}


func (a *ApiUsingDuplicateChecker) GetSubmissionRemark(checkerType duplicatechecker.DuplicateCheckerType, uid int64, identification string) (bool, string) {
	return a.container.GetSubmissionRemark(checkerType, uid, identification)
}


func (a *ApiUsingDuplicateChecker) SetSubmissionRemarkWithCustomExpiration(checkerType duplicatechecker.DuplicateCheckerType, uid int64, identification string, remark string, expiration time.Duration) {
	a.container.SetSubmissionRemarkWithCustomExpiration(checkerType, uid, identification, remark, expiration)
}


func (a *ApiUsingDuplicateChecker) SetSubmissionRemarkIfEnable(checkerType duplicatechecker.DuplicateCheckerType, uid int64, identification string, remark string) {
	if a.CurrentConfig().EnableDuplicateSubmissionsCheck {
		a.container.SetSubmissionRemark(checkerType, uid, identification, remark)
	}
}


func (a *ApiUsingDuplicateChecker) RemoveSubmissionRemark(checkerType duplicatechecker.DuplicateCheckerType, uid int64, identification string) {
	a.container.RemoveSubmissionRemark(checkerType, uid, identification)
}


func (a *ApiUsingDuplicateChecker) RemoveSubmissionRemarkIfEnable(checkerType duplicatechecker.DuplicateCheckerType, uid int64, identification string) {
	if a.CurrentConfig().EnableDuplicateSubmissionsCheck {
		a.container.RemoveSubmissionRemark(checkerType, uid, identification)
	}
}


func (a *ApiUsingDuplicateChecker) CheckFailureCount(c *core.WebContext, uid int64) error {
	if a.CurrentConfig().MaxFailuresPerIpPerMinute > 0 {
		clientIp := c.ClientIP()
		ipFailureCount := a.container.GetFailureCount(clientIp)

		if ipFailureCount >= a.CurrentConfig().MaxFailuresPerIpPerMinute {
			log.Warnf(c, "[base.CheckFailureCount] operation failure via IP \"%s\", current failure count: %d reached the limit", clientIp, ipFailureCount)
			return errs.ErrFailureCountLimitReached
		}
	}

	if a.CurrentConfig().MaxFailuresPerUserPerMinute > 0 && uid > 0 {
		uidFailureCount := a.container.GetFailureCount(utils.Int64ToString(uid))

		if uidFailureCount >= a.CurrentConfig().MaxFailuresPerUserPerMinute {
			log.Warnf(c, "[base.CheckFailureCount] operation failure via uid \"%d\", current failure count: %d reached the limit", uid, uidFailureCount)
			return errs.ErrFailureCountLimitReached
		}
	}

	return nil
}


func (a *ApiUsingDuplicateChecker) CheckAndIncreaseFailureCount(c *core.WebContext, uid int64) error {
	clientIp := c.ClientIP()
	ipFailureCount := uint32(0)
	uidFailureCount := uint32(0)

	if a.CurrentConfig().MaxFailuresPerIpPerMinute > 0 {
		ipFailureCount = a.container.GetFailureCount(clientIp)
	}

	if a.CurrentConfig().MaxFailuresPerUserPerMinute > 0 && uid > 0 {
		uidFailureCount = a.container.GetFailureCount(utils.Int64ToString(uid))
	}

	if a.CurrentConfig().MaxFailuresPerIpPerMinute > 0 && ipFailureCount < a.CurrentConfig().MaxFailuresPerIpPerMinute {
		log.Warnf(c, "[base.CheckAndIncreaseFailureCount] operation failure via IP \"%s\", previous failure count: %d", clientIp, ipFailureCount)
		a.container.IncreaseFailureCount(clientIp)
	}

	if a.CurrentConfig().MaxFailuresPerUserPerMinute > 0 && uid > 0 && uidFailureCount < a.CurrentConfig().MaxFailuresPerUserPerMinute {
		log.Warnf(c, "[base.CheckAndIncreaseFailureCount] operation failure via uid \"%d\", previous failure count: %d", uid, uidFailureCount)
		a.container.IncreaseFailureCount(utils.Int64ToString(uid))
	}

	if a.CurrentConfig().MaxFailuresPerIpPerMinute > 0 && ipFailureCount >= a.CurrentConfig().MaxFailuresPerIpPerMinute {
		log.Warnf(c, "[base.CheckAndIncreaseFailureCount] operation failure via IP \"%s\", current failure count: %d reached the limit", clientIp, ipFailureCount)
		return errs.ErrFailureCountLimitReached
	}

	if a.CurrentConfig().MaxFailuresPerUserPerMinute > 0 && uid > 0 && uidFailureCount >= a.CurrentConfig().MaxFailuresPerUserPerMinute {
		log.Warnf(c, "[base.CheckAndIncreaseFailureCount] operation failure via uid \"%d\", current failure count: %d reached the limit", uid, uidFailureCount)
		return errs.ErrFailureCountLimitReached
	}

	return nil
}


type ApiUsingAvatarProvider struct {
	container *avatars.AvatarProviderContainer
}


func (a *ApiUsingAvatarProvider) GetAvatarUrl(user *models.User) string {
	return a.container.GetAvatarUrl(user)
}


type ApiWithUserInfo struct {
	ApiUsingConfig
	ApiUsingAvatarProvider
}


func (a *ApiWithUserInfo) GetUserBasicInfo(user *models.User) *models.UserBasicInfo {
	return user.ToUserBasicInfo(a.CurrentConfig().AvatarProvider, a.GetAvatarUrl(user))
}
