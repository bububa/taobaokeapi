package tb

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type GetPrivilegeLinkRequest struct {
	ItemID   uint64 `json:"item_id"`
	SiteID   uint64 `json:"site_id"`
	AdzoneID uint64 `json:"adzone_id"`
}

type GetPrivilegeLinkResponse struct {
	CategoryID          json.Number `json:"category_id,omitempty"`
	CouponClickUrl      string      `json:"coupon_click_url,omitempty"`
	CouponStartTime     string      `json:"coupon_start_time,omitempty"`
	CouponEndTime       string      `json:"coupon_end_time,omitempty"`
	CouponInfo          string      `json:"coupon_info,omitempty"`
	ItemID              json.Number `json:"item_id,omitempty"`
	MaxCommissionRate   json.Number `json:"max_commission_rate,omitempty"`
	MinCommissionRate   json.Number `json:"min_commission_rate,omitempty"`
	CouponTotalCount    json.Number `json:"coupon_total_count,omitempty"`
	CouponRemainCount   json.Number `json:"coupon_remain_count,omitempty"`
	MmCouponRemainCount json.Number `json:"mm_coupon_remain_count,omitempty"`
	MmCouponTotalCount  json.Number `json:"mm_coupon_total_count,omitempty"`
	MmCouponClickUrl    string      `json:"mm_coupon_click_url,omitempty"`
	MmCouponEndTime     string      `json:"mm_coupon_end_time,omitempty"`
	MmCouponStartTime   string      `json:"mm_coupon_start_time,omitempty"`
	MmCouponInfo        string      `json:"mm_coupon_info,omitempty"`
	CouponType          int         `json:"coupon_type,omitempty"`
	ItemUrl             string      `json:"item_url,omitempty"`
}

func (r GetPrivilegeLinkRequest) Method() string {
	return "taobao.tbk.privilege.get"
}

func (r GetPrivilegeLinkRequest) Params() url.Values {
	params := url.Values{}
	params.Set("item_id", strconv.FormatUint(r.ItemID, 10))
	params.Set("site_id", strconv.FormatUint(r.SiteID, 10))
	params.Set("adzone_id", strconv.FormatUint(r.AdzoneID, 10))
	return params
}
