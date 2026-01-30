package models

import (
	"encoding/json"
)

type UserApplicationCloudSettingType string

const (
	USER_APPLICATION_CLOUD_SETTING_TYPE_STRING             UserApplicationCloudSettingType = "string"
	USER_APPLICATION_CLOUD_SETTING_TYPE_NUMBER             UserApplicationCloudSettingType = "number"
	USER_APPLICATION_CLOUD_SETTING_TYPE_BOOLEAN            UserApplicationCloudSettingType = "boolean"
	USER_APPLICATION_CLOUD_SETTING_TYPE_STRING_BOOLEAN_MAP UserApplicationCloudSettingType = "string_boolean_map"
)

var ALL_ALLOWED_CLOUD_SYNC_APP_SETTING_KEY_TYPES = map[string]UserApplicationCloudSettingType{
	
	"showAccountBalance": USER_APPLICATION_CLOUD_SETTING_TYPE_BOOLEAN,
	
	"showAmountInHomePage":                        USER_APPLICATION_CLOUD_SETTING_TYPE_BOOLEAN,
	"timezoneUsedForStatisticsInHomePage":         USER_APPLICATION_CLOUD_SETTING_TYPE_NUMBER,
	"overviewAccountFilterInHomePage":             USER_APPLICATION_CLOUD_SETTING_TYPE_STRING_BOOLEAN_MAP,
	"overviewTransactionCategoryFilterInHomePage": USER_APPLICATION_CLOUD_SETTING_TYPE_STRING_BOOLEAN_MAP,
	
	"itemsCountInTransactionListPage":      USER_APPLICATION_CLOUD_SETTING_TYPE_NUMBER,
	"showTotalAmountInTransactionListPage": USER_APPLICATION_CLOUD_SETTING_TYPE_BOOLEAN,
	"showTagInTransactionListPage":         USER_APPLICATION_CLOUD_SETTING_TYPE_BOOLEAN,
	
	"autoSaveTransactionDraft":                                 USER_APPLICATION_CLOUD_SETTING_TYPE_STRING,
	"autoGetCurrentGeoLocation":                                USER_APPLICATION_CLOUD_SETTING_TYPE_BOOLEAN,
	"alwaysShowTransactionPicturesInMobileTransactionEditPage": USER_APPLICATION_CLOUD_SETTING_TYPE_BOOLEAN,
	
	"rememberLastSelectedFileTypeInImportTransactionDialog": USER_APPLICATION_CLOUD_SETTING_TYPE_BOOLEAN,
	"lastSelectedFileTypeInImportTransactionDialog":         USER_APPLICATION_CLOUD_SETTING_TYPE_STRING,
	
	"insightsExplorerDefaultDateRangeType": USER_APPLICATION_CLOUD_SETTING_TYPE_NUMBER,
	"showTagInInsightsExplorerPage":        USER_APPLICATION_CLOUD_SETTING_TYPE_BOOLEAN,
	
	"totalAmountExcludeAccountIds":  USER_APPLICATION_CLOUD_SETTING_TYPE_STRING_BOOLEAN_MAP,
	"accountCategoryOrders":         USER_APPLICATION_CLOUD_SETTING_TYPE_STRING,
	"hideCategoriesWithoutAccounts": USER_APPLICATION_CLOUD_SETTING_TYPE_BOOLEAN,
	
	"currencySortByInExchangeRatesPage": USER_APPLICATION_CLOUD_SETTING_TYPE_NUMBER,
	
	"statistics.defaultChartDataType":                 USER_APPLICATION_CLOUD_SETTING_TYPE_NUMBER,
	"statistics.defaultTimezoneType":                  USER_APPLICATION_CLOUD_SETTING_TYPE_NUMBER,
	"statistics.defaultAccountFilter":                 USER_APPLICATION_CLOUD_SETTING_TYPE_STRING_BOOLEAN_MAP,
	"statistics.defaultTransactionCategoryFilter":     USER_APPLICATION_CLOUD_SETTING_TYPE_STRING_BOOLEAN_MAP,
	"statistics.defaultSortingType":                   USER_APPLICATION_CLOUD_SETTING_TYPE_NUMBER,
	"statistics.defaultCategoricalChartType":          USER_APPLICATION_CLOUD_SETTING_TYPE_NUMBER,
	"statistics.defaultCategoricalChartDataRangeType": USER_APPLICATION_CLOUD_SETTING_TYPE_NUMBER,
	"statistics.defaultTrendChartType":                USER_APPLICATION_CLOUD_SETTING_TYPE_NUMBER,
	"statistics.defaultTrendChartDataRangeType":       USER_APPLICATION_CLOUD_SETTING_TYPE_NUMBER,
	"statistics.defaultAssetTrendsChartType":          USER_APPLICATION_CLOUD_SETTING_TYPE_NUMBER,
	"statistics.defaultAssetTrendsChartDataRangeType": USER_APPLICATION_CLOUD_SETTING_TYPE_NUMBER,
}


type UserApplicationCloudSetting struct {
	Uid             int64                        `xorm:"PK"`
	Settings        ApplicationCloudSettingSlice `xorm:"BLOB"`
	UpdatedUnixTime int64
}


type UserApplicationCloudSettingsUpdateRequest struct {
	Settings   ApplicationCloudSettingSlice `json:"settings"`
	FullUpdate bool                         `json:"fullUpdate"`
}


type ApplicationCloudSettingSlice []ApplicationCloudSetting


type ApplicationCloudSetting struct {
	SettingKey   string `json:"settingKey"`
	SettingValue string `json:"settingValue"`
}


func (s *ApplicationCloudSettingSlice) FromDB(data []byte) error {
	return json.Unmarshal(data, s)
}


func (s *ApplicationCloudSettingSlice) ToDB() ([]byte, error) {
	return json.Marshal(s)
}


func (s ApplicationCloudSettingSlice) Len() int {
	return len(s)
}


func (s ApplicationCloudSettingSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}


func (s ApplicationCloudSettingSlice) Less(i, j int) bool {
	return s[i].SettingKey < s[j].SettingKey
}
