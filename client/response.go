package client

import (
	"encoding/json"
)

type Response struct {
	RequestID    string          `json:"request_id,omitempty"`
	Code         string          `json:"code,omitempty"`
	Msg          string          `json:"msg,omitempty"`
	SubCode      string          `json:"sub_code,omitempty"`
	SubMsg       string          `json:"sub_msg,omitempty"`
	Result       *ResponseResult `json:"result,omitempty"`
	Data         json.RawMessage `json:"data,omitempty"`
	Results      json.RawMessage `json:"results,omitempty"`
	TotalResults json.Number     `json:"total_results,omitempty"`
}

func (r Response) IsError() bool {
	return r.Code != ""
}

func (r Response) Error() string {
	return r.SubMsg
}

type ResponseResult struct {
	Data json.RawMessage `json:"data,omitempty"`
}
