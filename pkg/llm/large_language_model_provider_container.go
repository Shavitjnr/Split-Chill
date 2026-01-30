package llm

import (
	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
	"github.com/Shavitjnr/split-chill-ai/pkg/llm/data"
	"github.com/Shavitjnr/split-chill-ai/pkg/llm/provider"
	"github.com/Shavitjnr/split-chill-ai/pkg/llm/provider/googleai"
	"github.com/Shavitjnr/split-chill-ai/pkg/llm/provider/ollama"
	"github.com/Shavitjnr/split-chill-ai/pkg/llm/provider/openai"
	"github.com/Shavitjnr/split-chill-ai/pkg/settings"
)


type LargeLanguageModelProviderContainer struct {
	receiptImageRecognitionCurrentProvider provider.LargeLanguageModelProvider
}


var (
	Container = &LargeLanguageModelProviderContainer{}
)


func InitializeLargeLanguageModelProvider(config *settings.Config) error {
	var err error = nil

	if config.ReceiptImageRecognitionLLMConfig != nil {
		Container.receiptImageRecognitionCurrentProvider, err = initializeLargeLanguageModelProvider(config.ReceiptImageRecognitionLLMConfig, config.EnableDebugLog)

		if err != nil {
			return err
		}
	}

	return nil
}

func initializeLargeLanguageModelProvider(llmConfig *settings.LLMConfig, enableResponseLog bool) (provider.LargeLanguageModelProvider, error) {
	if llmConfig.LLMProvider == settings.OpenAILLMProvider {
		return openai.NewOpenAILargeLanguageModelProvider(llmConfig, enableResponseLog), nil
	} else if llmConfig.LLMProvider == settings.OpenAICompatibleLLMProvider {
		return openai.NewOpenAICompatibleLargeLanguageModelProvider(llmConfig, enableResponseLog), nil
	} else if llmConfig.LLMProvider == settings.OpenRouterLLMProvider {
		return openai.NewOpenRouterLargeLanguageModelProvider(llmConfig, enableResponseLog), nil
	} else if llmConfig.LLMProvider == settings.OllamaLLMProvider {
		return ollama.NewOllamaLargeLanguageModelProvider(llmConfig, enableResponseLog), nil
	} else if llmConfig.LLMProvider == settings.GoogleAILLMProvider {
		return googleai.NewGoogleAILargeLanguageModelProvider(llmConfig, enableResponseLog), nil
	} else if llmConfig.LLMProvider == "" {
		return nil, nil
	}

	return nil, errs.ErrInvalidLLMProvider
}


func (l *LargeLanguageModelProviderContainer) GetJsonResponseByReceiptImageRecognitionModel(c core.Context, uid int64, currentConfig *settings.Config, request *data.LargeLanguageModelRequest) (*data.LargeLanguageModelTextualResponse, error) {
	if currentConfig.ReceiptImageRecognitionLLMConfig == nil || Container.receiptImageRecognitionCurrentProvider == nil {
		return nil, errs.ErrInvalidLLMProvider
	}

	return l.receiptImageRecognitionCurrentProvider.GetJsonResponse(c, uid, currentConfig.ReceiptImageRecognitionLLMConfig, request)
}
