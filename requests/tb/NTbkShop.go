package tb

import "encoding/json"

type NTbkShop struct {
	// UserID 卖家ID
	UserID json.Number `json:"user_id,omitempty"`
	// ShopTitle 店铺名称
	ShopTitle string `json:"shop_title,omitempty"`
	// ShopType 店铺类型，B：天猫，C：淘宝
	ShopType string `json:"shop_type,omitempty"`
	// SellerNick 卖家昵称
	SellerNick string `json:"seller_nick,omitempty"`
	// PictURL 店标图片
	PictURL string `json:"pict_url,omitempty"`
	// ShopURL 店铺地址
	ShopURL string `json:"shop_url,omitempty"`
	// ClickURL 淘客地址
	ClickURL string `json:"click_url,omitempty"`
}
