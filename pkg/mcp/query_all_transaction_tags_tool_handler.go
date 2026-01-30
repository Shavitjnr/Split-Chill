package mcp

import (
	"encoding/json"
	"reflect"

	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/log"
	"github.com/Shavitjnr/split-chill-ai/pkg/models"
	"github.com/Shavitjnr/split-chill-ai/pkg/settings"
)


type MCPAllQueryTransactionTagsResponse struct {
	Tags []string `json:"tags" jsonschema_description:"List of transaction tags"`
}

type mcpQueryAllTransactionTagsToolHandler struct{}

var MCPQueryAllTransactionTagsToolHandler = &mcpQueryAllTransactionTagsToolHandler{}


func (h *mcpQueryAllTransactionTagsToolHandler) Name() string {
	return "query_all_transaction_tags"
}


func (h *mcpQueryAllTransactionTagsToolHandler) Description() string {
	return "Query transaction tags for the current user in Split Chill AI."
}


func (h *mcpQueryAllTransactionTagsToolHandler) InputType() reflect.Type {
	return nil
}


func (h *mcpQueryAllTransactionTagsToolHandler) OutputType() reflect.Type {
	return reflect.TypeOf(&MCPAllQueryTransactionTagsResponse{})
}


func (h *mcpQueryAllTransactionTagsToolHandler) Handle(c *core.WebContext, callToolReq *MCPCallToolRequest, user *models.User, currentConfig *settings.Config, services MCPAvailableServices) (any, []*MCPTextContent, error) {
	uid := user.Uid
	tags, err := services.GetTransactionTagService().GetAllTagsByUid(c, uid)

	if err != nil {
		log.Errorf(c, "[query_all_transaction_tags.Handle] failed to get tags for user \"uid:%d\", because %s", uid, err.Error())
		return nil, nil, err
	}

	tagNames := make([]string, 0, len(tags))

	for i := 0; i < len(tags); i++ {
		if tags[i].Hidden {
			continue
		}

		tagNames = append(tagNames, tags[i].Name)
	}

	response := MCPAllQueryTransactionTagsResponse{
		Tags: tagNames,
	}

	content, err := json.Marshal(response)

	if err != nil {
		return nil, nil, err
	}

	return response, []*MCPTextContent{
		NewMCPTextContent(string(content)),
	}, nil
}
