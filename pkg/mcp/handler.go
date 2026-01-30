package mcp

import (
	"reflect"

	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/models"
	"github.com/Shavitjnr/split-chill-ai/pkg/services"
	"github.com/Shavitjnr/split-chill-ai/pkg/settings"
)


type MCPAvailableServices interface {
	GetTransactionService() *services.TransactionService
	GetTransactionCategoryService() *services.TransactionCategoryService
	GetTransactionTagService() *services.TransactionTagService
	GetAccountService() *services.AccountService
	GetUserService() *services.UserService
}


type MCPToolHandler[T MCPTextContent | MCPImageContent | MCPAudioContent | MCPResourceLink | MCPEmbeddedResource] interface {
	
	Name() string

	
	Description() string

	
	InputType() reflect.Type

	
	OutputType() reflect.Type

	
	Handle(*core.WebContext, *MCPCallToolRequest, *models.User, *settings.Config, MCPAvailableServices) (any, []*T, error)
}
