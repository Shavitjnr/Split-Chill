package requestid


type RequestIdGenerator interface {
	GenerateRequestId(clientIpAddr string, clientPort uint16) string
	GetCurrentServerUniqId() uint16
	GetCurrentInstanceUniqId() uint16
}
