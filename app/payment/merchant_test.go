package payment_test

import (
	"github.com/godcong/wego/app/payment"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
	"testing"
)

// TestMerchant_AddSubMerchant ...
func TestMerchant_AddSubMerchant(t *testing.T) {
	obj := payment.NewMerchant(core.C(util.Map{
		//"sandbox": true,
		"app_id":  "wx3c69535993f4651d",
		"secret":  "f8c7a2cf0c6ed44e2c719964bbe13b1e",
		"key":     "aTKnSUcTkbEnhwQNdutWkQxAjnhAz2jK",
		"aes_key": "DbWPitkfiWkhLwDPA48laxJojyiNqVwtK7R1ENPvEwC",
	}))
	resp := obj.AddSubMerchant(util.Map{
		"page_index": 1,
		"page_size":  10,
	})

	t.Log(resp.Error())
	t.Log(resp.ToMap())
}

// TestMerchant_QuerySubMerchantByMerchantId ...
func TestMerchant_QuerySubMerchantByMerchantId(t *testing.T) {
	obj := payment.NewMerchant(core.C(util.Map{
		"sandbox": true,
		"app_id":  "wx3c69535993f4651d",
		"secret":  "f8c7a2cf0c6ed44e2c719964bbe13b1e",
		"key":     "aTKnSUcTkbEnhwQNdutWkQxAjnhAz2jK",
		"aes_key": "DbWPitkfiWkhLwDPA48laxJojyiNqVwtK7R1ENPvEwC",
	}))
	resp := obj.QuerySubMerchantByMerchantID("123")

	t.Log(resp.Error())
	t.Log(resp.ToMap())
}

// TestMerchant_QuerySubMerchantByWeChatId ...
func TestMerchant_QuerySubMerchantByWeChatId(t *testing.T) {
	obj := payment.NewMerchant(core.C(util.Map{
		"sandbox": true,
		"app_id":  "wx3c69535993f4651d",
		"secret":  "f8c7a2cf0c6ed44e2c719964bbe13b1e",
		"key":     "aTKnSUcTkbEnhwQNdutWkQxAjnhAz2jK",
		"aes_key": "DbWPitkfiWkhLwDPA48laxJojyiNqVwtK7R1ENPvEwC",
	}))
	resp := obj.QuerySubMerchantByWeChatID("123")

	t.Log(resp.Error())
	t.Log(resp.ToMap())
}
