package avatars

import (
	"fmt"
	"strings"

	"github.com/Shavitjnr/split-chill-ai/pkg/models"
	"github.com/Shavitjnr/split-chill-ai/pkg/utils"
)


const gravatarUrlFormat = "https://www.gravatar.com/avatar/%s"


type GravatarAvatarProvider struct {
}


func NewGravatarAvatarProvider() *GravatarAvatarProvider {
	return &GravatarAvatarProvider{}
}


func (p *GravatarAvatarProvider) GetAvatarUrl(user *models.User) string {
	email := user.Email
	email = strings.TrimSpace(email)
	email = strings.ToLower(email)
	emailMd5 := utils.MD5EncodeToString([]byte(email))

	return fmt.Sprintf(gravatarUrlFormat, emailMd5)
}
