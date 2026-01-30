package avatars

import (
	"github.com/Shavitjnr/split-chill-ai/pkg/models"
)


type NullAvatarProvider struct {
}


func NewNullAvatarProvider() *NullAvatarProvider {
	return &NullAvatarProvider{}
}


func (p *NullAvatarProvider) GetAvatarUrl(user *models.User) string {
	return ""
}
