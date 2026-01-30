package googleai

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
	"github.com/Shavitjnr/split-chill-ai/pkg/llm/data"
	"github.com/Shavitjnr/split-chill-ai/pkg/llm/provider"
	"github.com/Shavitjnr/split-chill-ai/pkg/llm/provider/common"
	"github.com/Shavitjnr/split-chill-ai/pkg/log"
	"github.com/Shavitjnr/split-chill-ai/pkg/settings"
)

const googleAIGenerateContentAPIFormat = "https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent"


type GoogleAILargeLanguageModelAdapter struct {
	common.HttpLargeLanguageModelAdapter
	GoogleAIAPIKey  string
	GoogleAIModelID string
}


type GoogleAIGenerateContentRequest struct {
	Contents []*GoogleAIGenerateContentRequestContent `json:"contents"`
}


type GoogleAIGenerateContentRequestContent struct {
	Parts []*GoogleAIGenerateContentRequestContentPart `json:"parts"`
}


type GoogleAIGenerateContentRequestContentPart struct {
	Text       string                                    `json:"text,omitempty"`
	InlineData *GoogleAIGenerateContentRequestInlineData `json:"inlineData,omitempty"`
}


type GoogleAIGenerateContentRequestInlineData struct {
	MimeType string `json:"mimeType"`
	Data     string `json:"data"`
}


type GoogleAIGenerateContentResponse struct {
	Candidates []*GoogleAIGenerateContentResponseCandidate `json:"candidates"`
}


type GoogleAIGenerateContentResponseCandidate struct {
	Content *GoogleAIGenerateContentResponseContent `json:"content"`
}


type GoogleAIGenerateContentResponseContent struct {
	Part []*GoogleAIGenerateContentResponseContentPart `json:"parts"`
}


type GoogleAIGenerateContentResponseContentPart struct {
	Text *string `json:"text"`
}


func (p *GoogleAILargeLanguageModelAdapter) BuildTextualRequest(c core.Context, uid int64, request *data.LargeLanguageModelRequest, responseType data.LargeLanguageModelResponseFormat) (*http.Request, error) {
	requestBody, err := p.buildJsonRequestBody(c, uid, request, responseType)

	if err != nil {
		return nil, err
	}

	requestUrl := fmt.Sprintf(googleAIGenerateContentAPIFormat, p.GoogleAIModelID)
	httpRequest, err := http.NewRequest("POST", requestUrl, bytes.NewReader(requestBody))

	if err != nil {
		return nil, err
	}

	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("X-goog-api-key", p.GoogleAIAPIKey)

	return httpRequest, nil
}


func (p *GoogleAILargeLanguageModelAdapter) ParseTextualResponse(c core.Context, uid int64, body []byte, responseType data.LargeLanguageModelResponseFormat) (*data.LargeLanguageModelTextualResponse, error) {
	generateContentResponse := &GoogleAIGenerateContentResponse{}
	err := json.Unmarshal(body, &generateContentResponse)

	if err != nil {
		log.Errorf(c, "[google_ai_large_language_model_adapter.ParseTextualResponse] failed to parse generate content response for user \"uid:%d\", because %s", uid, err.Error())
		return nil, errs.ErrFailedToRequestRemoteApi
	}

	if generateContentResponse == nil || generateContentResponse.Candidates == nil || len(generateContentResponse.Candidates) < 1 ||
		generateContentResponse.Candidates[0].Content == nil || len(generateContentResponse.Candidates[0].Content.Part) < 1 ||
		generateContentResponse.Candidates[0].Content.Part[0].Text == nil {
		log.Errorf(c, "[google_ai_large_language_model_adapter.ParseTextualResponse] generate content response is invalid for user \"uid:%d\"", uid)
		return nil, errs.ErrFailedToRequestRemoteApi
	}

	textualResponse := &data.LargeLanguageModelTextualResponse{
		Content: *generateContentResponse.Candidates[0].Content.Part[0].Text,
	}

	return textualResponse, nil
}

func (p *GoogleAILargeLanguageModelAdapter) buildJsonRequestBody(c core.Context, uid int64, request *data.LargeLanguageModelRequest, responseType data.LargeLanguageModelResponseFormat) ([]byte, error) {
	if p.GoogleAIModelID == "" {
		return nil, errs.ErrInvalidLLMModelId
	}

	generateContentRequest := &GoogleAIGenerateContentRequest{
		Contents: []*GoogleAIGenerateContentRequestContent{
			{
				Parts: make([]*GoogleAIGenerateContentRequestContentPart, 0, 2),
			},
		},
	}

	if request.SystemPrompt != "" {
		generateContentRequest.Contents[0].Parts = append(generateContentRequest.Contents[0].Parts, &GoogleAIGenerateContentRequestContentPart{
			Text: request.SystemPrompt,
		})
	}

	if len(request.UserPrompt) > 0 {
		if request.UserPromptType == data.LARGE_LANGUAGE_MODEL_REQUEST_PROMPT_TYPE_IMAGE_URL {
			imageBase64Data := base64.StdEncoding.EncodeToString(request.UserPrompt)
			generateContentRequest.Contents[0].Parts = append(generateContentRequest.Contents[0].Parts, &GoogleAIGenerateContentRequestContentPart{
				InlineData: &GoogleAIGenerateContentRequestInlineData{
					MimeType: request.UserPromptContentType,
					Data:     imageBase64Data,
				},
			})
		} else {
			generateContentRequest.Contents[0].Parts = append(generateContentRequest.Contents[0].Parts, &GoogleAIGenerateContentRequestContentPart{
				Text: string(request.UserPrompt),
			})
		}
	}

	requestBodyBytes, err := json.Marshal(generateContentRequest)

	if err != nil {
		log.Errorf(c, "[google_ai_large_language_model_adapter.buildJsonRequestBody] failed to marshal request body for user \"uid:%d\", because %s", uid, err.Error())
		return nil, errs.ErrOperationFailed
	}

	log.Debugf(c, "[google_ai_large_language_model_adapter.buildJsonRequestBody] request body is %s", requestBodyBytes)
	return requestBodyBytes, nil
}


func NewGoogleAILargeLanguageModelProvider(llmConfig *settings.LLMConfig, enableResponseLog bool) provider.LargeLanguageModelProvider {
	return common.NewCommonHttpLargeLanguageModelProvider(llmConfig, enableResponseLog, &GoogleAILargeLanguageModelAdapter{
		GoogleAIAPIKey:  llmConfig.GoogleAIAPIKey,
		GoogleAIModelID: llmConfig.GoogleAIModelID,
	})
}
