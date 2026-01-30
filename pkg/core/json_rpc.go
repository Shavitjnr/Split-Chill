package core

import "encoding/json"


const JSONRPCVersion = "2.0"


type JSONRPCRequest struct {
	JSONRPC string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
	ID      any             `json:"id,omitempty"`
}


type JSONRPCResponse struct {
	JSONRPC string        `json:"jsonrpc"`
	Result  any           `json:"result,omitempty"`
	Error   *JSONRPCError `json:"error,omitempty"`
	ID      any           `json:"id,omitempty"`
}


type JSONRPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}


var JSONRPCParseError = &JSONRPCError{
	Code:    -32700,
	Message: "Parse error",
	Data:    nil,
}


var JSONRPCMethodNotFoundError = &JSONRPCError{
	Code:    -32601,
	Message: "Method not found",
	Data:    nil,
}


var JSONRPCInvalidParamsError = &JSONRPCError{
	Code:    -32602,
	Message: "Invalid params",
	Data:    nil,
}


var JSONRPCInternalError = &JSONRPCError{
	Code:    -32603,
	Message: "Internal error",
	Data:    nil,
}


func NewJSONRPCResponse(id any, result any) *JSONRPCResponse {
	return &JSONRPCResponse{
		JSONRPC: JSONRPCVersion,
		Result:  result,
		Error:   nil,
		ID:      id,
	}
}


func NewJSONRPCErrorResponse(id any, err *JSONRPCError) *JSONRPCResponse {
	return &JSONRPCResponse{
		JSONRPC: JSONRPCVersion,
		Result:  nil,
		Error: &JSONRPCError{
			Code:    err.Code,
			Message: err.Message,
			Data:    nil,
		},
		ID: id,
	}
}


func NewJSONRPCErrorResponseWithCause(id any, err *JSONRPCError, cause string) *JSONRPCResponse {
	return &JSONRPCResponse{
		JSONRPC: JSONRPCVersion,
		Result:  nil,
		Error: &JSONRPCError{
			Code:    err.Code,
			Message: err.Message,
			Data:    cause,
		},
		ID: id,
	}
}
