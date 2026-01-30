package avatars

import (
	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
	"github.com/Shavitjnr/split-chill-ai/pkg/models"
	"github.com/Shavitjnr/split-chill-ai/pkg/settings"
)


type AvatarProviderContainer struct {
	current AvatarProvider
}


var (
	Container = &AvatarProviderContainer{}
)


func InitializeAvatarProvider(config *settings.Config) error {
	if config.AvatarProvider == core.USER_AVATAR_PROVIDER_INTERNAL {
		Container.current = NewInternalStorageAvatarProvider(config)
		return nil
	} else if config.AvatarProvider == core.USER_AVATAR_PROVIDER_GRAVATAR {
		Container.current = NewGravatarAvatarProvider()
		return nil
	} else if config.AvatarProvider == "" {
		Container.current = NewNullAvatarProvider()
		return nil
	}

	return errs.ErrInvalidAvatarProvider
}


func (p *AvatarProviderContainer) GetAvatarUrl(user *models.User) string {
	if p.current == nil {
		return ""
	}

	return p.current.GetAvatarUrl(user)
}
