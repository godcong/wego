package payment_test

import (
	"github.com/godcong/wego/app/payment"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
	"testing"
)

func TestBase_Pay(t *testing.T) {
	base := payment.NewBase(core.C(util.Map{
		"sandbox": true,
		"app_id": "wx3c69535993f4651d",
		"secret": "f8c7a2cf0c6ed44e2c719964bbe13b1e",
		"key":     "aTKnSUcTkbEnhwQNdutWkQxAjnhAz2jK",
		"aes_key": "DbWPitkfiWkhLwDPA48laxJojyiNqVwtK7R1ENPvEwC",
	}))
	resp := base.Pay(util.Map{
		"body":         "image形象店-深圳腾大- QQ公仔",
		"out_trade_no": "1217752501201407033233368018",
		"total_fee":    "888",
		"auth_code":    "120061098828009406",
	})
	t.Log(resp.Error())
	t.Log(resp.ToMap())
}

func TestBase_AuthCodeToOpenid(t *testing.T) {
	base := payment.NewBase(core.C(util.Map{
		"app_id":  "wx3c69535993f4651d",
		"secret":  "f8c7a2cf0c6ed44e2c719964bbe13b1e",
		"mch_id":  "1516796851",
	}))
	resp := base.AuthCodeToOpenid("1212121")
	t.Log(resp.Error())
	t.Log(resp.ToMap())
}
