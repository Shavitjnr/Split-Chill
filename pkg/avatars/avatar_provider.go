package avatars

import "github.com/Shavitjnr/split-chill-ai/pkg/models"


type AvatarProvider interface {
	GetAvatarUrl(user *models.User) string
}
