package core

import "context"

const nullContextId = "00000000-0000-0000-0000-00000000"


type NullContext struct {
	context.Context
}


func (c *NullContext) GetContextId() string {
	return nullContextId
}


func (c *NullContext) GetClientLocale() string {
	return ""
}


func NewNullContext() *NullContext {
	return &NullContext{
		Context: context.Background(),
	}
}
