package client

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	GATEWAY = "https://api.taobaokeapi.com"
)

type Client struct {
	usertoken string
	debug     bool
}

func NewClient(usertoken string) *Client {
	return &Client{
		usertoken: usertoken,
	}
}

func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

func (c *Client) Do(req Request, ret interface{}) (int64, error) {
	params := req.Params()
	params.Set("usertoken", c.usertoken)
	params.Set("method", req.Method())
	var builder strings.Builder
	builder.WriteString(GATEWAY)
	builder.WriteString("?")
	builder.WriteString(params.Encode())
	if c.debug {
		log.Printf("[taobaokeapi REQ] %s\n", builder.String())
	}
	httpResp, err := http.Get(builder.String())
	if err != nil {
		return 0, err
	}
	if err != nil {
		return 0, err
	}
	defer httpResp.Body.Close()
	var res Response
	err = decodeJSONHttpResponse(httpResp.Body, &res, c.debug)
	if err != nil {
		return 0, err
	}
	if res.IsError() {
		return 0, res
	}
	if ret == nil {
		return 0, nil
	}
	total, _ := res.TotalResults.Int64()
	if res.Result == nil && res.Data != nil {
		return total, json.Unmarshal(res.Data, &ret)
	} else if res.Result == nil && res.Results != nil {
		return total, json.Unmarshal(res.Results, &ret)
	} else if res.Result == nil && res.Data == nil {
		return total, nil
	}
	return total, json.Unmarshal(res.Result.Data, &ret)
}

func decodeJSONHttpResponse(r io.Reader, v interface{}, debug bool) error {
	if !debug {
		return json.NewDecoder(r).Decode(v)
	}
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	body2 := body
	buf := bytes.NewBuffer(make([]byte, 0, len(body2)+1024))
	if err := json.Indent(buf, body2, "", "    "); err == nil {
		body2 = buf.Bytes()
	}
	log.Printf("[taobaokeapi] [API] http response body:\n%s\n", body2)

	return json.Unmarshal(body, v)
}
