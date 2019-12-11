package nimiqrpc

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (

	// ErrRPCIDMismatch is returned when the request ID and response ID don't match.
	// According to the JSON-RPC 2.0 spec, the request and response ID should always match.
	ErrRPCIDMismatch = errors.New("JSON-RPC: The request ID and the response ID didn't match")

	// ErrRPCParse is returned when an invalid JSON was received by the server.
	// This is a standard error from the JSON-RPC 2.0 spec.
	ErrRPCParse = errors.New("JSON-RPC: Invalid JSON was received by the server. An error occurred on the server while parsing the JSON text")

	// ErrRPCInvalidReq is returned when the JSON sent was not a valid request object.
	// This is a standard error from the JSON-RPC 2.0 spec.
	ErrRPCInvalidReq = errors.New("JSON-RPC: The JSON sent is not a valid Request object")

	// ErrRPCMethodNotFound is returned when the requested method does not exist or is not available on the server.
	// This is a standard error from the JSON-RPC 2.0 spec.
	ErrRPCMethodNotFound = errors.New("JSON-RPC: The method does not exist / is not available")

	// ErrRPCInvalidParams is returned when the parameters for the requested method are invalid
	// This is a standard error from the JSON-RPC 2.0 spec.
	ErrRPCInvalidParams = errors.New("JSON-RPC: Invalid method parameter(s)")

	// ErrRPCInternal is returned when a internal error occured on the RPC server
	// This is a standard error from the JSON-RPC 2.0 spec.
	ErrRPCInternal = errors.New("JSON-RPC: Internal JSON-RPC error")

	// ErrRPCServer is returned when a server error occured.
	// This is a standard error from the JSON-RPC 2.0 spec.
	ErrRPCServer = errors.New("JSON-RPC: Server error")
)

// RPCError represents a JSON-RPC error object
type RPCError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    string `json:"data,omitempty"`
}

// Parse can be used to parse a GO error from an JSON-RPC error object
func (re *RPCError) Parse() error {
	switch {
	case re.Code == -32700:
		return ErrRPCParse
	case re.Code == -32600:
		return ErrRPCInvalidReq
	case re.Code == -32601:
		return ErrRPCMethodNotFound
	case re.Code == -32602:
		return ErrRPCInvalidParams
	case re.Code == -32603:
		return ErrRPCInternal
	case re.Code >= -32000 && re.Code <= -32099:
		return ErrRPCServer
	case re.Code > 0:
		return fmt.Errorf("Nimiq RPC error: %s", re.Message)
	default:
		return nil
	}
}

// RPCRequest represents a JSON-RPC 2.0 request object
type RPCRequest struct {
	JSONRPC string      `json:"jsonrpc,omitempty"`
	Method  string      `json:"method,omitempty"`
	Params  interface{} `json:"params,omitempty"`
	ID      int         `json:"id,omitempty"`
}

// NewRPCRequest returns a new RPCRequest using the method, params and id as parameters
func NewRPCRequest(method string, params interface{}, id int) *RPCRequest {
	return &RPCRequest{
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
		ID:      id,
	}
}

// RPCResponse represents a JSON-RPC 2.0 response object
type RPCResponse struct {
	JSONRPC string          `json:"jsonrpc,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   RPCError        `json:"error,omitempty"`
	ID      int             `json:"id,omitempty"`
}

// ParseRPCResponse can be used to parse a JSON-RPC 2.0 response object from a
// raw slice of bytes.
func ParseRPCResponse(data []byte, resp *RPCResponse) error {
	if err := json.Unmarshal(data, resp); err != nil {
		return err
	}
	return nil
}
