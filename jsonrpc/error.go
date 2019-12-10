package jsonrpc

import "errors"

var (
	ErrParse          = errors.New("JSON-RPC: Invalid JSON was received by the server. An error occurred on the server while parsing the JSON text.")
	ErrInvalidReq     = errors.New("JSON-RPC: The JSON sent is not a valid Request object.")
	ErrMethodNotFound = errors.New("JSON-RPC: The method does not exist / is not available.")
	ErrInvalidParams  = errors.New("JSON-RPC: Invalid method parameter(s).")
	ErrInternal       = errors.New("JSON-RPC: Internal JSON-RPC error.")
	ErrServer         = errors.New("JSON-RPC: Server error")
)

type Error struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message, omitempty"`
	Data    string `json:"data,omitempty"`
}

func (re *Error) Parse() error {
	switch {
	case re.Code == -32700:
		return ErrParse
	case re.Code == -32600:
		return ErrInvalidReq
	case re.Code == -32601:
		return ErrMethodNotFound
	case re.Code == -32602:
		return ErrInvalidParams
	case re.Code == -32603:
		return ErrInternal
	case re.Code >= -32000 && re.Code <= -32099:
		return ErrServer
	default:
		return nil
	}
}
