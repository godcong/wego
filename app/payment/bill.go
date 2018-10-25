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
对账单日期	bill_date	是	String(8)	20140603	下载对账单的日期，格式：20140603
账单类型	bill_type	是	String(8)	ALL
ALL，返回当日所有订单信息，默认值
SUCCESS，返回当日成功支付的订单
REFUND，返回当日退款订单
RECHARGE_REFUND，返回当日充值退款订单
压缩账单	tar_type	否	String(8)	GZIP	非必传参数，固定值：GZIP，返回格式为.gzip的压缩包账单。不传则默认为数据流形式。
*/
func (b *Bill) Get(bd string, op util.Map) core.Response {
	m := make(util.Map)
	m.Set("appid", b.config.Get("app_id"))
	m.Set("bill_date", bd)
	if op == nil || !op.Has("bill_type") {
		m.Set("bill_type", "ALL")
	}

	m.Join(op)
	return b.Request(downloadBillURLSuffix, m)
}
