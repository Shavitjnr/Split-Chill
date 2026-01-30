package requestid

import (
	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/settings"
)


type RequestIdContainer struct {
	current RequestIdGenerator
}


var (
	Container = &RequestIdContainer{}
)


func InitializeRequestIdGenerator(c core.Context, config *settings.Config) error {
	generator, err := NewDefaultRequestIdGenerator(c, config)

	if err != nil {
		return err
	}

	Container.current = generator
	return nil
}


func (r *RequestIdContainer) GenerateRequestId(clientIpAddr string, clientPort uint16) string {
	if r.current == nil {
		return ""
	}

	return r.current.GenerateRequestId(clientIpAddr, clientPort)
}


func (r *RequestIdContainer) GetCurrentServerUniqId() uint16 {
	if r.current == nil {
		return 0
	}

	return r.current.GetCurrentServerUniqId()
}


func (r *RequestIdContainer) GetCurrentInstanceUniqId() uint16 {
	if r.current == nil {
		return 0
	}

	return r.current.GetCurrentInstanceUniqId()
}
