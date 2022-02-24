package tb

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// ScOrderDetailsGetRequest 淘宝客-服务商-所有订单查询
type ScOrderDetailsGetRequest struct {
	// QueryType 查询时间类型，1：按照订单淘客创建时间查询，2:按照订单淘客付款时间查询，3:按照订单淘客结算时间查询，4:按照订单更新时间；
	QueryType int `json:"query_type,omitempty"`
	// PositionIndex 位点，除第一页之外，都需要传递；前端原样返回。
	PositionIndex string `json:"position_index,omitempty"`
	// PageSize 页大小，默认20，1~100
	PageSize int `json:"page_size,omitempty"`
	// MemberType 推广者角色类型,2:二方，3:三方，不传，表示所有角色
	MemberType int `json:"member_type,omitempty"`
	// TkStatus 淘客订单状态，12-付款，13-关闭，14-确认收货，3-结算成功;不传，表示所有状态
	TkStatus int `json:"tk_status,omitempty"`
	// EndTime 订单查询结束时间，订单开始时间至订单结束时间，中间时间段日常要求不超过3个小时，但如618、双11、年货节等大促期间预估时间段不可超过20分钟，超过会提示错误，调用时请务必注意时间段的选择，以保证亲能正常调用！
	EndTime string `json:"end_time,omitempty"`
	// StartTime 订单查询开始时间
	StartTime string `json:"start_time,omitempty"`
	// JumpType 跳转类型，当向前或者向后翻页必须提供,-1: 向前翻页,1：向后翻页
	JumpType int `json:"jump_type,omitempty"`
	// PageNo 第几页，默认1，1~100
	PageNo int `json:"page_no,omitempty"`
	// OrderScene 场景订单场景类型，1:常规订单，2:渠道订单，3:会员运营订单，默认为1
	OrderScene int `json:"order_scene,omitempty"`
}

func (r ScOrderDetailsGetRequest) Method() string {
	return "taobao.tbk.sc.order.details.get"
}

func (r ScOrderDetailsGetRequest) Params() url.Values {
	params := url.Values{}
	if r.QueryType > 0 {
		params.Set("query_type", strconv.Itoa(r.QueryType))
	}
	if r.PositionIndex != "" {
		params.Set("position_index", r.PositionIndex)
	}
	if r.PageSize > 0 {
		params.Set("page_size", strconv.Itoa(r.PageSize))
	}
	if r.MemberType > 0 {
		params.Set("member_type", strconv.Itoa(r.MemberType))
	}
	if r.TkStatus > 0 {
		params.Set("tk_status", strconv.Itoa(r.TkStatus))
	}
	params.Set("end_time", r.EndTime)
	params.Set("start_time", r.StartTime)
	if r.JumpType == 1 || r.JumpType == -1 {
		params.Set("jump_type", strconv.Itoa(r.JumpType))
	}
	if r.PageNo > 0 {
		params.Set("page_no", strconv.Itoa(r.PageNo))
	}
	if r.OrderScene > 0 {
		params.Set("order_scene", strconv.Itoa(r.OrderScene))
	}
	return params
}

type ScOrderDetailsGetResponse struct {
	// PublisherOrderDto
	Results *ScOrderDetailsGetResults `json:"results,omitempty"`
	// 位点字段，由调用方原样传递
	PositionIndex string `json:"position_index,omitempty"`
	// 页码
	PageNo json.Number `json:"page_no,omitempty"`
	// 页大小
	PageSize json.Number `json:"page_size,omitempty"`
	// 是否还有上一页
	HasPre string `json:"has_pre,omitempty"`
	// 是否还有下一页
	HasNext string `json:"has_next,omitempty"`
}

type ScOrderDetailsGetResults struct {
	Orders json.RawMessage `json:"publisher_order_dto,omitempty"`
}
