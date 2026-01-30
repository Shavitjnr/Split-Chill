package openai

import (
	"net/http"

	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/llm/provider"
	"github.com/Shavitjnr/split-chill-ai/pkg/settings"
)


type OpenAIOfficialChatCompletionsAPIProvider struct {
	OpenAIChatCompletionsAPIProvider
	OpenAIAPIKey  string
	OpenAIModelID string
}

const openAIChatCompletionsUrl = "https://api.openai.com/v1/chat/completions"


func (p *OpenAIOfficialChatCompletionsAPIProvider) BuildChatCompletionsHttpRequest(c core.Context, uid int64) (*http.Request, error) {
	req, err := http.NewRequest("POST", openAIChatCompletionsUrl, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+p.OpenAIAPIKey)

	return req, nil
}


func (p *OpenAIOfficialChatCompletionsAPIProvider) GetModelID() string {
	return p.OpenAIModelID
}


func NewOpenAILargeLanguageModelProvider(llmConfig *settings.LLMConfig, enableResponseLog bool) provider.LargeLanguageModelProvider {
	return newCommonOpenAIChatCompletionsAPILargeLanguageModelAdapter(llmConfig, enableResponseLog, &OpenAIOfficialChatCompletionsAPIProvider{
		OpenAIAPIKey:  llmConfig.OpenAIAPIKey,
		OpenAIModelID: llmConfig.OpenAIModelID,
	})
}
