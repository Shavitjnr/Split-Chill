package avatars

import (
	"fmt"

	"github.com/Shavitjnr/split-chill-ai/pkg/models"
	"github.com/Shavitjnr/split-chill-ai/pkg/settings"
)

const internalAvatarUrlFormat = "%savatar/%d.%s"


type InternalStorageAvatarProvider struct {
	webRootUrl string
}


func NewInternalStorageAvatarProvider(config *settings.Config) *InternalStorageAvatarProvider {
	return &InternalStorageAvatarProvider{
		webRootUrl: config.RootUrl,
	}
}


func (p *InternalStorageAvatarProvider) GetAvatarUrl(user *models.User) string {
	if user.CustomAvatarType == "" {
		return ""
	}

	return fmt.Sprintf(internalAvatarUrlFormat, p.webRootUrl, user.Uid, user.CustomAvatarType)
}
