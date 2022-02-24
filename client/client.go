package client

import (
	"encoding/json"
	"net/http"
)

const (
	GATEWAY = "https://api.taobaokeapi.com"
)

type Client struct {
	usertoken string
}

func NewClient(usertoken string) *Client {
	return &Client{
		usertoken: usertoken,
	}
}

func (c *Client) Do(req Request, ret interface{}) error {
	params := req.Params()
	params.Set("usertoken", c.usertoken)
	params.Set("method", req.Method())
	httpResp, err := http.PostForm(GATEWAY, params)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	var res Response
	err = json.NewDecoder(httpResp.Body).Decode(&res)
	if err != nil {
		return err
	}
	if res.IsError() {
		return res
	}
	if ret == nil {
		return nil
	}
	if res.Result == nil {
		return nil
	}
	return json.Unmarshal(res.Result.Data, &ret)
}
