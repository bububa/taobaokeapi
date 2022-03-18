package tb

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type ShopGetRequest struct {
	// Fields 需返回的字段列表
	Fields string `json:"fields,omitempty"`
	// Q 查询词
	Q string `json:"q,omitempty"`
	// Sort 排序_des（降序），排序_asc（升序），佣金比率（commission_rate）， 商品数量（auction_count），销售总数量（total_auction）
	Sort string `json:"sort,omitempty"`
	// IsTmall 是否商城的店铺，设置为true表示该是属于淘宝商城的店铺，设置为false或不设置表示不判断这个属性
	IsTmall bool `json:"is_tmall,omitempty"`
	// StartCredit 信用等级下限，1~20
	StartCredit int `json:"start_credit,omitempty"`
	// EndCredit 信用等级上限，1~20
	EndCredit int `json:"end_credit,omitempty"`
	// StartCommissionRate 淘客佣金比率下限，1~10000
	StartCommissionRate int `json:"start_commission_rate,omitempty"`
	// EndCommissionRate 淘客佣金比率上限，1~10000
	EndCommissionRate int `json:"end_commission_rate,omitempty"`
	// StartTotalAction 店铺商品总数下限
	StartTotalAction int `json:"start_total_action,omitempty"`
	// EndTotalAction 店铺商品总数上限
	EndTotalAction int `json:"end_total_action,omitempty"`
	// StartAuctionCount 累计推广商品下限
	StartAuctionCount int64 `json:"start_auction_count,omitempty"`
	// EndAuctionCount 累计推广商品上限
	EndAuctionCount int64 `json:"end_auction_count,omitempty"`
	// Platform 链接形式：1：PC，2：无线，默认：１
	Platform int `json:"platform,omitempty"`
	// PageNo 第几页，默认1，1~100
	PageNo int `json:"page_no,omitempty"`
	// PageSize 页大小，默认20，1~100
	PageSize int `json:"page_size,omitempty"`
}

func (r ShopGetRequest) Method() string {
	return "taobao.tbk.shop.get"
}

func (r ShopGetRequest) Params() url.Values {
	values := url.Values{}
	values.Set("q", r.Q)
	if r.Fields != "" {
		values.Set("fields", r.Fields)
	}
	if r.Sort != "" {
		values.Set("sort", r.Sort)
	}
	if r.IsTmall {
		values.Set("is_tmall", "true")
	}
	if r.StartCredit > 0 {
		values.Set("start_credit", strconv.Itoa(r.StartCredit))
	}
	if r.EndCredit > 0 {
		values.Set("end_credit", strconv.Itoa(r.EndCredit))
	}
	if r.StartCommissionRate > 0 {
		values.Set("start_commission_rate", strconv.Itoa(r.StartCommissionRate))
	}
	if r.EndCommissionRate > 0 {
		values.Set("end_commission_rate", strconv.Itoa(r.EndCommissionRate))
	}
	if r.StartTotalAction > 0 {
		values.Set("start_total_acton", strconv.Itoa(r.StartTotalAction))
	}
	if r.EndTotalAction > 0 {
		values.Set("end_total_acton", strconv.Itoa(r.EndTotalAction))
	}
	if r.StartAuctionCount > 0 {
		values.Set("start_auction_count", strconv.FormatInt(r.StartAuctionCount, 10))
	}
	if r.EndAuctionCount > 0 {
		values.Set("end_auction_count", strconv.FormatInt(r.EndAuctionCount, 10))
	}
	if r.Platform > 0 {
		values.Set("platform", strconv.Itoa(r.Platform))
	}
	if r.PageNo > 0 {
		values.Set("page_no", strconv.Itoa(r.PageNo))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	return values
}

type ShopGetResults struct {
	Shops json.RawMessage `json:"n_tbk_shop,omitempty"`
}
