package payment

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

// Coupon ...
type Coupon struct {
	*Payment
}

func newCoupon(payment *Payment) interface{} {
	return &Coupon{
		Payment: payment,
	}
}

// NewCoupon ...
func NewCoupon(config *core.Config) *Coupon {
	return newCoupon(NewPayment(config)).(*Coupon)
}

// Send ...
func (c *Coupon) Send(maps util.Map) core.Responder {
	maps = util.MapNilMake(maps)
	maps.Set("appid", c.GetString("app_id"))
	maps.Set("openid_count", 1)
	return c.SafeRequest(mmpaymkttransfersSendCoupon, maps)
}

// QueryStock ...
func (c *Coupon) QueryStock(maps util.Map) core.Responder {
	maps = util.MapNilMake(maps)
	maps.Set("appid", c.GetString("app_id"))
	return c.SafeRequest(mmpaymkttransfersQueryCouponStock, maps)
}

// QueryInfo ...
func (c *Coupon) QueryInfo(maps util.Map) core.Responder {
	maps = util.MapNilMake(maps)
	maps.Set("appid", c.GetString("app_id"))
	return c.SafeRequest(mmpaymkttransfersQueryCouponsInfo, maps)
}
