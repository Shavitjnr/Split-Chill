package openai

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"

	"github.com/invopop/jsonschema"
	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
	"github.com/Shavitjnr/split-chill-ai/pkg/llm/data"
	"github.com/Shavitjnr/split-chill-ai/pkg/llm/provider"
	"github.com/Shavitjnr/split-chill-ai/pkg/llm/provider/common"
	"github.com/Shavitjnr/split-chill-ai/pkg/log"
	"github.com/Shavitjnr/split-chill-ai/pkg/settings"
)


type OpenAIChatCompletionsAPIProvider interface {
	
	BuildChatCompletionsHttpRequest(c core.Context, uid int64) (*http.Request, error)

	
	GetModelID() string
}


type CommonOpenAIChatCompletionsAPILargeLanguageModelAdapter struct {
	common.HttpLargeLanguageModelAdapter
	apiProvider OpenAIChatCompletionsAPIProvider
}


type OpenAIMessageRole string


const (
	OpenAIMessageRoleSystem OpenAIMessageRole = "system"
	OpenAIMessageRoleUser   OpenAIMessageRole = "user"
)


type OpenAIChatCompletionsRequestResponseFormatType string


const (
	OpenAIChatCompletionsRequestResponseFormatTypeJsonObject OpenAIChatCompletionsRequestResponseFormatType = "json_object"
	OpenAIChatCompletionsRequestResponseFormatTypeJsonSchema OpenAIChatCompletionsRequestResponseFormatType = "json_schema"
)


type OpenAIChatCompletionsRequest struct {
	Model          string                                      `json:"model"`
	Stream         bool                                        `json:"stream"`
	Messages       []any                                       `json:"messages"`
	ResponseFormat *OpenAIChatCompletionsRequestResponseFormat `json:"response_format,omitempty"`
}


type OpenAIChatCompletionsRequestMessage[T string | []*OpenAIChatCompletionsRequestImageContent] struct {
	Role    OpenAIMessageRole `json:"role"`
	Content T                 `json:"content"`
}


type OpenAIChatCompletionsRequestImageContent struct {
	Type     string                                `json:"type"`
	ImageURL *OpenAIChatCompletionsRequestImageUrl `json:"image_url"`
}


type OpenAIChatCompletionsRequestResponseFormat struct {
	Type       OpenAIChatCompletionsRequestResponseFormatType `json:"type"`
	JsonSchema *jsonschema.Schema                             `json:"json_schema,omitempty"`
}


type OpenAIChatCompletionsRequestImageUrl struct {
	Url string `json:"url"`
}


type OpenAIChatCompletionsResponse struct {
	Choices []*OpenAIChatCompletionsResponseChoice `json:"choices"`
}


type OpenAIChatCompletionsResponseChoice struct {
	Message *OpenAIChatCompletionsResponseMessage `json:"message"`
}


type OpenAIChatCompletionsResponseMessage struct {
	Content *string `json:"content"`
}


func (p *CommonOpenAIChatCompletionsAPILargeLanguageModelAdapter) BuildTextualRequest(c core.Context, uid int64, request *data.LargeLanguageModelRequest, responseType data.LargeLanguageModelResponseFormat) (*http.Request, error) {
	requestBody, err := p.buildJsonRequestBody(c, uid, request, responseType)

	if err != nil {
		return nil, err
	}

	httpRequest, err := p.apiProvider.BuildChatCompletionsHttpRequest(c, uid)

	if err != nil {
		return nil, err
	}

	httpRequest.Body = io.NopCloser(bytes.NewReader(requestBody))
	httpRequest.Header.Set("Content-Type", "application/json")

	return httpRequest, nil
}


func (p *CommonOpenAIChatCompletionsAPILargeLanguageModelAdapter) ParseTextualResponse(c core.Context, uid int64, body []byte, responseType data.LargeLanguageModelResponseFormat) (*data.LargeLanguageModelTextualResponse, error) {
	chatCompletionsResponse := &OpenAIChatCompletionsResponse{}
	err := json.Unmarshal(body, &chatCompletionsResponse)

	if err != nil {
		log.Errorf(c, "[openai_common_compatible_large_language_model_adapter.ParseTextualResponse] failed to parse chat completions response for user \"uid:%d\", because %s", uid, err.Error())
		return nil, errs.ErrFailedToRequestRemoteApi
	}

	if chatCompletionsResponse == nil || chatCompletionsResponse.Choices == nil || len(chatCompletionsResponse.Choices) < 1 ||
		chatCompletionsResponse.Choices[0].Message == nil ||
		chatCompletionsResponse.Choices[0].Message.Content == nil {
		log.Errorf(c, "[openai_common_compatible_large_language_model_adapter.ParseTextualResponse] chat completions response is invalid for user \"uid:%d\"", uid)
		return nil, errs.ErrFailedToRequestRemoteApi
	}

	textualResponse := &data.LargeLanguageModelTextualResponse{
		Content: *chatCompletionsResponse.Choices[0].Message.Content,
	}

	return textualResponse, nil
}

func (p *CommonOpenAIChatCompletionsAPILargeLanguageModelAdapter) buildJsonRequestBody(c core.Context, uid int64, request *data.LargeLanguageModelRequest, responseType data.LargeLanguageModelResponseFormat) ([]byte, error) {
	if p.apiProvider.GetModelID() == "" {
		return nil, errs.ErrInvalidLLMModelId
	}

	chatCompletionsRequest := &OpenAIChatCompletionsRequest{
		Model:    p.apiProvider.GetModelID(),
		Stream:   request.Stream,
		Messages: make([]any, 0, 2),
	}

	if request.SystemPrompt != "" {
		chatCompletionsRequest.Messages = append(chatCompletionsRequest.Messages, &OpenAIChatCompletionsRequestMessage[string]{
			Role:    OpenAIMessageRoleSystem,
			Content: request.SystemPrompt,
		})
	}

	if len(request.UserPrompt) > 0 {
		if request.UserPromptType == data.LARGE_LANGUAGE_MODEL_REQUEST_PROMPT_TYPE_IMAGE_URL {
			imageBase64Data := "data:" + request.UserPromptContentType + ";base64," + base64.StdEncoding.EncodeToString(request.UserPrompt)
			chatCompletionsRequest.Messages = append(chatCompletionsRequest.Messages, &OpenAIChatCompletionsRequestMessage[[]*OpenAIChatCompletionsRequestImageContent]{
				Role: OpenAIMessageRoleUser,
				Content: []*OpenAIChatCompletionsRequestImageContent{
					{
						Type: "image_url",
						ImageURL: &OpenAIChatCompletionsRequestImageUrl{
							Url: imageBase64Data,
						},
					},
				},
			})
		} else {
			chatCompletionsRequest.Messages = append(chatCompletionsRequest.Messages, &OpenAIChatCompletionsRequestMessage[string]{
				Role:    OpenAIMessageRoleUser,
				Content: string(request.UserPrompt),
			})
		}
	}

	if responseType == data.LARGE_LANGUAGE_MODEL_RESPONSE_FORMAT_JSON {
		if request.ResponseJsonObjectType != nil {
			schemeGenerator := jsonschema.Reflector{
				Anonymous:      true,
				DoNotReference: true,
				ExpandedStruct: true,
			}

			schema := schemeGenerator.ReflectFromType(request.ResponseJsonObjectType)
			schema.Version = ""

			chatCompletionsRequest.ResponseFormat = &OpenAIChatCompletionsRequestResponseFormat{
				Type:       OpenAIChatCompletionsRequestResponseFormatTypeJsonSchema,
				JsonSchema: schema,
			}
		} else {
			chatCompletionsRequest.ResponseFormat = &OpenAIChatCompletionsRequestResponseFormat{
				Type: OpenAIChatCompletionsRequestResponseFormatTypeJsonObject,
			}
		}
	}

	requestBodyBytes, err := json.Marshal(chatCompletionsRequest)

	if err != nil {
		log.Errorf(c, "[openai_common_compatible_large_language_model_adapter.buildJsonRequestBody] failed to marshal request body for user \"uid:%d\", because %s", uid, err.Error())
		return nil, errs.ErrOperationFailed
	}

	log.Debugf(c, "[openai_common_compatible_large_language_model_adapter.buildJsonRequestBody] request body is %s", requestBodyBytes)
	return requestBodyBytes, nil
}

func newCommonOpenAIChatCompletionsAPILargeLanguageModelAdapter(llmConfig *settings.LLMConfig, enableResponseLog bool, apiProvider OpenAIChatCompletionsAPIProvider) provider.LargeLanguageModelProvider {
	return common.NewCommonHttpLargeLanguageModelProvider(llmConfig, enableResponseLog, &CommonOpenAIChatCompletionsAPILargeLanguageModelAdapter{
		apiProvider: apiProvider,
	})
}
