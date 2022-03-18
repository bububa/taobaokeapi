package tb

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type ScTpwdConvertRequest struct {
	// PasswordContent 需要解析的淘口令
	PasswordContent string `json:"password_content,omitempty"`
	// AdzoneID 广告位ID
	AdzoneID uint64 `json:"adzone_id,omitempty"`
	// SiteID 备案的网站id, mm_xx_xx_xx pid三段式中的第二段
	SiteID uint64 `json:"site_id,omitempty"`
	// Dx 1表示商品转通用计划链接，其他值或不传表示优先转营销计划链接
	Dx int `json:"dx,omitempty"`
}

func (r ScTpwdConvertRequest) Method() string {
	return "taobao.tbk.sc.tpwd.convert"
}

func (r ScTpwdConvertRequest) Params() url.Values {
	values := url.Values{}
	values.Set("password_content", r.PasswordContent)
	values.Set("site_id", strconv.FormatUint(r.SiteID, 10))
	values.Set("adzone_id", strconv.FormatUint(r.AdzoneID, 10))
	if r.Dx == 1 {
		values.Set("dx", strconv.Itoa(r.Dx))
	}
	return values
}

type ScTpwdConvertResponse struct {
	// NumIid 商品Id
	NumIid json.Number `json:"num_iid,omitempty"`
	// ClickURL 商品淘客转链
	ClickURL string `json:"click_url,omitempty"`
}
