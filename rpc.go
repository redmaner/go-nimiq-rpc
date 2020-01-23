// Copyright 2019 Jake "redmaner" van der Putten.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	ErrRPCParse = errors.New("JSON-RPC: Invalid JSON was received by the server. An error occurred on the server while parsing the JSON text")

	// ErrRPCInvalidReq is returned when the JSON sent was not a valid request object.
	ErrRPCInvalidReq = errors.New("JSON-RPC: The JSON sent is not a valid Request object")

	// ErrRPCMethodNotFound is returned when the requested method does not exist or is not available on the server.
	ErrRPCMethodNotFound = errors.New("JSON-RPC: The method does not exist / is not available")

	// ErrRPCInvalidParams is returned when the parameters for the requested method are invalid
	ErrRPCInvalidParams = errors.New("JSON-RPC: Invalid method parameter(s)")

	// ErrRPCInternal is returned when a internal error occurred on the RPC server
	ErrRPCInternal = errors.New("JSON-RPC: Internal JSON-RPC error")

	// ErrRPCServer is returned when a server error occurred.
	ErrRPCServer = errors.New("JSON-RPC: Server error")
)

// RPCRequest represents a JSON-RPC 2.0 request object
type RPCRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
	ID      int64       `json:"id"`
}

// NewRPCRequest returns a new RPCRequest using the method and params as arguments
func NewRPCRequest(method string, params interface{}) *RPCRequest {
	return &RPCRequest{
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
	}
}

// RPCResponse represents a JSON-RPC 2.0 response object
type RPCResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   RPCError        `json:"error,omitempty"`
	ID      int64           `json:"id"`
}

// RPCError represents a JSON-RPC error object
type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

// Parse can be used to parse a Go error from an JSON-RPC error object
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
