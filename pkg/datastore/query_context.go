package datastore

import (
	"fmt"
	"time"

	"xorm.io/xorm/log"

	"github.com/Shavitjnr/split-chill-ai/pkg/core"
)


type XOrmContextAdapter struct {
	requestId string
}


func (c *XOrmContextAdapter) Deadline() (deadline time.Time, ok bool) {
	return
}


func (c *XOrmContextAdapter) Done() <-chan struct{} {
	return nil
}


func (c *XOrmContextAdapter) Err() error {
	return nil
}


func (c *XOrmContextAdapter) Value(key any) any {
	if key == log.SessionIDKey && c.requestId != "" {
		return fmt.Sprintf("%s", c.requestId)
	}

	return nil
}

func NewXOrmContextAdapter(c core.Context) *XOrmContextAdapter {
	if c != nil {
		return &XOrmContextAdapter{
			requestId: c.GetContextId(),
		}
	}

	return &XOrmContextAdapter{}
}
