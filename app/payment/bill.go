package payment

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

/*Bill 账单 */
type Bill struct {
	*Payment
}

func newBill(p *Payment) *Bill {
	return &Bill{
		Payment: p,
	}
}

/*NewBill 账单 */
func NewBill(config *core.Config) *Bill {
	return newBill(NewPayment(config))
}

/*Get 下载对账单
接口链接
https://api.mch.weixin.qq.com/pay/downloadbill
是否需要证书
不需要。
请求参数
字段名	变量名	必填	类型	示例值	描述
对账单日期	bill_date	是	String(8)	20140603	下载对账单的日期，格式:20140603
*/
func (b *Bill) Get(bd string, op ...util.Map) core.Response {
	m := util.MapsToMap(op)
	m.Set("appid", b.Get("app_id"))
	m.Set("bill_date", bd)
	if !m.Has("bill_type") {
		m.Set("bill_type", "ALL")
	}

	return b.Request(downloadBillURLSuffix, m)
}
