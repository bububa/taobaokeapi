package tb

import (
	"net/url"
	"strconv"
	"strings"
)

type ScShopConvertRequest struct {
	// SiteID 备案的网站id, mm_xx_xx_xx pid三段式中的第二段
	SiteID uint64 `json:"site_id,omitempty"`
	// Fields 需返回的字段列表
	Fields string `json:"fields,omitempty"`
	// UserIDs 卖家ID串，用','分割，从taobao.tbk.shop.get接口获取user_id字段
	UserIDs []uint64 `json:"user_ids,omitempty"`
	// Platform 链接形式：1：PC，2：无线，默认：１
	Platform int `json:"platform,omitempty"`
	// AdzoneID 广告位ID，区分效果位置
	AdzoneID uint64 `json:"adzone_id,omitempty"`
	// Unid 自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道
	Unid string `json:"unid,omitempty"`
}

func (r ScShopConvertRequest) Method() string {
	return "taobao.tbk.sc.shop.convert"
}

func (r ScShopConvertRequest) Params() url.Values {
	values := url.Values{}
	values.Set("site_id", strconv.FormatUint(r.SiteID, 10))
	if r.Fields != "" {
		values.Set("fields", r.Fields)
	}
	totalUserIDs := len(r.UserIDs)
	if totalUserIDs > 0 {
		var builder strings.Builder
		for idx, uid := range r.UserIDs {
			builder.WriteString(strconv.FormatUint(uid, 10))
			if idx < totalUserIDs-1 {
				builder.WriteString(",")
			}
		}
		values.Set("user_ids", builder.String())
	}
	if r.Platform > 0 {
		values.Set("platform", strconv.Itoa(r.Platform))
	}
	values.Set("adzone_id", strconv.FormatUint(r.AdzoneID, 10))
	if r.Unid != "" {
		values.Set("unid", r.Unid)
	}
	return values
}

type ScShopConvertResponse struct {
	Results []NTbkShop `json:"results,omitempty"`
}
