package client

import (
	"encoding/json"
)

type Response struct {
	RequestID string          `json::"request_id,omitempty"`
	Code      string          `json:"code,omitempty"`
	Msg       string          `json:"msg,omitempty"`
	SubCode   string          `json:"sub_code,omitempty"`
	SubMsg    string          `json:"sub_msg,omitempty"`
	Result    *ResponseResult `json:"result,omitempty"`
}

func (this Response) IsError() bool {
	if this.Code != "" {
		return true
	}
	return false
}

func (this Response) Error() string {
	return this.SubMsg
}

type ResponseResult struct {
	Data json.RawMessage `json:"data,omitempty"`
}
