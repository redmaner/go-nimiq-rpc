package jsonrpc

import "encoding/json"

type Response struct {
	JSONRPC string          `json:"jsonrpc,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   Error           `json:"error,omitempty"`
	ID      string          `json:"id,omitempty"`
}

func ParseResponse(data []byte, resp *Response) error {
	if err := json.Unmarshal(data, resp); err != nil {
		return err
	}
	return nil
}
