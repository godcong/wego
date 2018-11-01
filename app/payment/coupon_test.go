package payment_test

import (
	"github.com/godcong/wego/app/payment"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
	"testing"
)

func TestCoupon_Send(t *testing.T) {
	coupon := payment.NewCoupon(core.C(util.Map{
		"sandbox": false,
		"mch_id":  "123123",
		"app_id":  "wx3c69535993f4651d",
		"secret":  "f8c7a2cf0c6ed44e2c719964bbe13b1e",
		"key":     "aTKnSUcTkbEnhwQNdutWkQxAjnhAz2jK",
		"aes_key": "DbWPitkfiWkhLwDPA48laxJojyiNqVwtK7R1ENPvEwC",
	}))
	resp := coupon.Send(util.Map{
		"coupon_stock_id":"1",
		"openid":"12341234",
	})

	t.Log(resp.Error())
	t.Log(resp.ToMap())

}
