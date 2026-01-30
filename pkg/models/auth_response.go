package models


type AuthResponse struct {
	Token                    string                        `json:"token"`
	Need2FA                  bool                          `json:"need2FA"`
	User                     *UserBasicInfo                `json:"user"`
	ApplicationCloudSettings *ApplicationCloudSettingSlice `json:"applicationCloudSettings,omitempty"`
	NotificationContent      string                        `json:"notificationContent,omitempty"`
}


type RegisterResponse struct {
	AuthResponse
	NeedVerifyEmail       bool `json:"needVerifyEmail"`
	PresetCategoriesSaved bool `json:"presetCategoriesSaved"`
}
