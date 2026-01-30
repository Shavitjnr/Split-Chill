package core

import "context"


type Context interface {
	context.Context
	GetContextId() string
	GetClientLocale() string
}
