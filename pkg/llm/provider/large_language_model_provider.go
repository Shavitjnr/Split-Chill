package provider

import (
	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/llm/data"
	"github.com/Shavitjnr/split-chill-ai/pkg/settings"
)


type LargeLanguageModelProvider interface {
	
	GetJsonResponse(c core.Context, uid int64, currentLLMConfig *settings.LLMConfig, request *data.LargeLanguageModelRequest) (*data.LargeLanguageModelTextualResponse, error)
}
