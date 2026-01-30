package openai

import (
	"net/http"

	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/llm/provider"
	"github.com/Shavitjnr/split-chill-ai/pkg/settings"
)

const openAICompatibleChatCompletionsPath = "chat/completions"


type OpenAICompatibleChatCompletionsAPIProvider struct {
	OpenAIChatCompletionsAPIProvider
	OpenAICompatibleBaseURL string
	OpenAICompatibleAPIKey  string
	OpenAICompatibleModelID string
}


func (p *OpenAICompatibleChatCompletionsAPIProvider) BuildChatCompletionsHttpRequest(c core.Context, uid int64) (*http.Request, error) {
	req, err := http.NewRequest("POST", p.getFinalChatCompletionsRequestUrl(), nil)

	if err != nil {
		return nil, err
	}

	if p.OpenAICompatibleAPIKey != "" {
		req.Header.Set("Authorization", "Bearer "+p.OpenAICompatibleAPIKey)
	}

	return req, nil
}


func (p *OpenAICompatibleChatCompletionsAPIProvider) GetModelID() string {
	return p.OpenAICompatibleModelID
}

func (p *OpenAICompatibleChatCompletionsAPIProvider) getFinalChatCompletionsRequestUrl() string {
	url := p.OpenAICompatibleBaseURL

	if url[len(url)-1] != '/' {
		url += "/"
	}

	url += openAICompatibleChatCompletionsPath
	return url
}


func NewOpenAICompatibleLargeLanguageModelProvider(llmConfig *settings.LLMConfig, enableResponseLog bool) provider.LargeLanguageModelProvider {
	return newCommonOpenAIChatCompletionsAPILargeLanguageModelAdapter(llmConfig, enableResponseLog, &OpenAICompatibleChatCompletionsAPIProvider{
		OpenAICompatibleBaseURL: llmConfig.OpenAICompatibleBaseURL,
		OpenAICompatibleAPIKey:  llmConfig.OpenAICompatibleAPIKey,
		OpenAICompatibleModelID: llmConfig.OpenAICompatibleModelID,
	})
}
