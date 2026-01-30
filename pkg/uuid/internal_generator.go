package uuid

import (
	"sync/atomic"
	"time"

	"github.com/Shavitjnr/split-chill-ai/pkg/settings"
)


const (
	internalUuidUnixTimeBits = 32
	internalUuidUnixTimeMask = (1 << internalUuidUnixTimeBits) - 1

	internalUuidTypeBits = 4
	internalUuidTypeMask = (1 << internalUuidTypeBits) - 1

	internalUuidServerIdBits = 8
	internalUuidServerIdMask = (1 << internalUuidServerIdBits) - 1

	internalUuidSeqIdBits = 19
	internalUuidSeqIdMask = (1 << internalUuidSeqIdBits) - 1

	seqNumberIdBits = 32
	seqNumberIdMask = (1 << seqNumberIdBits) - 1
)


type InternalUuidInfo struct {
	UnixTime     uint32
	UuidType     uint8
	UuidServerId uint8
	SequentialId uint32
}


type InternalUuidGenerator struct {
	uuidSeqNumbers [1 << internalUuidTypeBits]atomic.Uint64
	uuidServerId   uint8
}


func NewInternalUuidGenerator(config *settings.Config) (*InternalUuidGenerator, error) {
	generator := &InternalUuidGenerator{
		uuidServerId: config.UuidServerId,
	}

	return generator, nil
}


func (u *InternalUuidGenerator) GenerateUuid(idType UuidType) int64 {
	uuids := u.GenerateUuids(idType, 1)

	if len(uuids) < 1 {
		return 0
	}

	return uuids[0]
}


func (u *InternalUuidGenerator) GenerateUuids(idType UuidType, count uint16) []int64 {


	uuids := make([]int64, count)

	if count < 1 {
		return uuids
	}

	var unixTime uint64
	var newFirstSeqId uint64
	var newLastSeqId uint64
	uuidType := uint8(idType)

	for {
		unixTime = uint64(time.Now().Unix())
		newLastSeqId = u.uuidSeqNumbers[uuidType].Add(uint64(count))
		newSeqUnixTime := newLastSeqId >> seqNumberIdBits

		if unixTime == newSeqUnixTime {
			newFirstSeqId = newLastSeqId - uint64(count-1)
			break
		} else if unixTime < newSeqUnixTime {
			continue
		}

		currentSeqId := newLastSeqId
		newFirstSeqId = unixTime << seqNumberIdBits
		newLastSeqId = newFirstSeqId + uint64(count-1)

		if u.uuidSeqNumbers[uuidType].CompareAndSwap(currentSeqId, newLastSeqId) {
			break
		}
	}

	for i := 0; i < int(count); i++ {
		seqId := (newFirstSeqId + uint64(i)) & seqNumberIdMask

		
		if seqId > internalUuidSeqIdMask {
			return nil
		}

		uuids[i] = u.assembleUuid(unixTime, uuidType, seqId)
	}

	return uuids
}

func (u *InternalUuidGenerator) parseInternalUuidInfo(uuid int64) *InternalUuidInfo {
	seqId := uint32(uuid & internalUuidSeqIdMask)
	uuid = uuid >> internalUuidSeqIdBits

	uuidServerId := uint8(uuid & internalUuidServerIdMask)
	uuid = uuid >> internalUuidServerIdBits

	uuidType := uint8(uuid & internalUuidTypeMask)
	uuid = uuid >> internalUuidTypeBits

	unixTime := uint32(uuid & internalUuidUnixTimeMask)

	return &InternalUuidInfo{
		UnixTime:     unixTime,
		UuidType:     uuidType,
		UuidServerId: uuidServerId,
		SequentialId: seqId,
	}
}

func (u *InternalUuidGenerator) assembleUuid(unixTime uint64, uuidType uint8, seqId uint64) int64 {
	unixTimePart := (int64(unixTime) & internalUuidUnixTimeMask) << (internalUuidTypeBits + internalUuidServerIdBits + internalUuidSeqIdBits)
	uuidTypePart := (int64(uuidType) & internalUuidTypeMask) << (internalUuidServerIdBits + internalUuidSeqIdBits)
	uuidServerIdPart := (int64(u.uuidServerId) & internalUuidServerIdMask) << internalUuidSeqIdBits
	seqIdPart := int64(seqId) & internalUuidSeqIdMask

	return unixTimePart | uuidTypePart | uuidServerIdPart | seqIdPart
}
