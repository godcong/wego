package payment

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
	"golang.org/x/text/encoding/simplifiedchinese"
	"testing"
	"time"
)

// TestBill_Download ...
func TestBill_Download(t *testing.T) {
	bill := NewBill(core.C(util.Map{
		"sandbox": true,
		"app_id":  "wx3c69535993f4651d",
		"secret":  "f8c7a2cf0c6ed44e2c719964bbe13b1e",
		"key":     "aTKnSUcTkbEnhwQNdutWkQxAjnhAz2jK",
		"aes_key": "DbWPitkfiWkhLwDPA48laxJojyiNqVwtK7R1ENPvEwC",
	}))
	resp := bill.Download(util.Map{"bill_date": "20181103"})
	_ = core.SaveEncodingTo(resp, "d:/test.csv", simplifiedchinese.GBK.NewEncoder())
	t.Log(resp.Error())
	t.Log(resp.ToMap())

}

func TestBill_BatchQueryComment(t *testing.T) {
	bill := NewBill(core.C(util.Map{
		"sandbox": true,
		"app_id":  "wx3c69535993f4651d",
		"secret":  "f8c7a2cf0c6ed44e2c719964bbe13b1e",
		"key":     "aTKnSUcTkbEnhwQNdutWkQxAjnhAz2jK",
		"aes_key": "DbWPitkfiWkhLwDPA48laxJojyiNqVwtK7R1ENPvEwC",
	}))
	resp := bill.BatchQueryComment(util.Map{
		"begin_time": time.Now().Unix(),
		"end_time":   time.Now().Unix(),
		"offset":     "0",
	})
	_ = core.SaveEncodingTo(resp, "d:/test.csv", simplifiedchinese.GBK.NewEncoder())
	t.Log(resp.Error())
	t.Log(resp.ToMap())
}
