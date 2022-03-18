package tb

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type ItemInfoGetRequest struct {
	// NumIIds 商品ID串，用,分割，最大40个
	NumIids string `json:"num_iids,omitempty"`
	// Platform 链接形式：1：PC，2：无线，默认：１
	Platform int `json:"platform,omitempty"`
	// Ip ip地址，影响邮费获取，如果不传或者传入不准确，邮费无法精准提供
	Ip string `json:"ip,omitempty"`
}

func (r ItemInfoGetRequest) Method() string {
	return "taobao.tbk.item.info.get"
}

func (r ItemInfoGetRequest) Params() url.Values {
	values := url.Values{}
	values.Set("num_iids", r.NumIids)
	if r.Platform > 0 {
		values.Set("platform", strconv.Itoa(r.Platform))
	}
	if r.Ip != "" {
		values.Set("ip", r.Ip)
	}
	return values
}

type ItemInfoGetResults struct {
	Items json.RawMessage `json:"n_tbk_item,omitempty"`
}
