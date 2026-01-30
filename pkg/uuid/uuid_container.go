package uuid

import (
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
	"github.com/Shavitjnr/split-chill-ai/pkg/settings"
)


type UuidContainer struct {
	current UuidGenerator
}


var (
	Container = &UuidContainer{}
)


func InitializeUuidGenerator(config *settings.Config) error {
	if config.UuidGeneratorType == settings.InternalUuidGeneratorType {
		generator, err := NewInternalUuidGenerator(config)
		Container.current = generator

		return err
	}

	return errs.ErrInvalidUuidMode
}


func (u *UuidContainer) GenerateUuid(uuidType UuidType) int64 {
	if u.current == nil {
		return 0
	}

	return u.current.GenerateUuid(uuidType)
}


func (u *UuidContainer) GenerateUuids(uuidType UuidType, count uint16) []int64 {
	if u.current == nil {
		return nil
	}

	return u.current.GenerateUuids(uuidType, count)
}
