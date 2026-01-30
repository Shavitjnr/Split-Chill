package openai

import (
	"net/http"

	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/llm/provider"
	"github.com/Shavitjnr/split-chill-ai/pkg/settings"
)


type OpenRouterChatCompletionsAPIProvider struct {
	OpenAIChatCompletionsAPIProvider
	OpenRouterAPIKey  string
	OpenRouterModelID string
}

const openRouterChatCompletionsUrl = "https://openrouter.ai/api/v1/chat/completions"


func (p *OpenRouterChatCompletionsAPIProvider) BuildChatCompletionsHttpRequest(c core.Context, uid int64) (*http.Request, error) {
	req, err := http.NewRequest("POST", openRouterChatCompletionsUrl, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+p.OpenRouterAPIKey)
	req.Header.Set("HTTP-Referer", "https://Split Chill AI.shavitjnr.net/")
	req.Header.Set("X-Title", core.ApplicationName)

	return req, nil
}


func (p *OpenRouterChatCompletionsAPIProvider) GetModelID() string {
	return p.OpenRouterModelID
}


func NewOpenRouterLargeLanguageModelProvider(llmConfig *settings.LLMConfig, enableResponseLog bool) provider.LargeLanguageModelProvider {
	return newCommonOpenAIChatCompletionsAPILargeLanguageModelAdapter(llmConfig, enableResponseLog, &OpenRouterChatCompletionsAPIProvider{
		OpenRouterAPIKey:  llmConfig.OpenRouterAPIKey,
		OpenRouterModelID: llmConfig.OpenRouterModelID,
	})
}
