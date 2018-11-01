package payment

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
	"golang.org/x/text/encoding/simplifiedchinese"
	"testing"
)

func TestBill_Download(t *testing.T) {
	bill := NewBill(core.C(util.Map{
		"sandbox": true,
		"app_id": "wx3c69535993f4651d",
		"secret": "f8c7a2cf0c6ed44e2c719964bbe13b1e",
		"key":     "aTKnSUcTkbEnhwQNdutWkQxAjnhAz2jK",
		"aes_key": "DbWPitkfiWkhLwDPA48laxJojyiNqVwtK7R1ENPvEwC",
	}))
	resp := bill.Download("20181103")
	_ = core.SaveEncodingTo(resp, "d:/test.csv", simplifiedchinese.GBK.NewEncoder())
	t.Log(resp.Error())
	t.Log(resp.ToMap())

}
