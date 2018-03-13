package payment

import (
	"strconv"

	"github.com/godcong/wego/core"
)

type Refund struct {
	core.Config
	*Payment
}

func (r *Refund) refund(num string, total, refund int, options core.Map) *core.Response {
	options.NilSet("out_refund_no", num)
	options.NilSet("total_fee", strconv.Itoa(total))
	options.NilSet("refund_fee", strconv.Itoa(refund))
	options.NilSet("appid", r.Config.Get("app_id"))

	return r.SafeRequest(REFUND_URL_SUFFIX, options)
}

//成功：
//<xml><return_code><![CDATA[SUCCESS]]></return_code>
//<return_msg><![CDATA[OK]]></return_msg>
//<appid><![CDATA[wxbafed7010e0f4531]]></appid>
//<mch_id><![CDATA[1497361732]]></mch_id>
//<nonce_str><![CDATA[RH3CCiZfUkzqCEwD]]></nonce_str>
//<sign><![CDATA[822690E5802AAFAB229B53FB7C36E625]]></sign>
//<result_code><![CDATA[SUCCESS]]></result_code>
//<transaction_id><![CDATA[4200000080201803137991413766]]></transaction_id>
//<out_trade_no><![CDATA[20180313160643671522177497]]></out_trade_no>
//<out_refund_no><![CDATA[1]]></out_refund_no>
//<refund_id><![CDATA[50000606082018031303804215880]]></refund_id>
//<refund_channel><![CDATA[]]></refund_channel>
//<refund_fee>30</refund_fee>
//<coupon_refund_fee>0</coupon_refund_fee>
//<total_fee>30</total_fee>
//<cash_fee>30</cash_fee>
//<coupon_refund_count>0</coupon_refund_count>
//<cash_refund_fee>30</cash_refund_fee>
//</xml>
func (r *Refund) ByOutTradeNumber(tradeNum, num string, total, refund int, options core.Map) core.Map {
	options = core.MapNilMake(options)
	options.NilSet("out_trade_no", tradeNum)
	return r.refund(num, total, refund, options).ToMap()
}

//成功：
//<xml><return_code><![CDATA[SUCCESS]]></return_code>
//<return_msg><![CDATA[OK]]></return_msg>
//<appid><![CDATA[wxbafed7010e0f4531]]></appid>
//<mch_id><![CDATA[1497361732]]></mch_id>
//<nonce_str><![CDATA[fv35lPYg52pIzMdQ]]></nonce_str>
//<sign><![CDATA[4DC97871C1CA1E1A152FD7FE79085039]]></sign>
//<result_code><![CDATA[SUCCESS]]></result_code>
//<transaction_id><![CDATA[4200000066201803138050731804]]></transaction_id>
//<out_trade_no><![CDATA[20180313155830338675328863]]></out_trade_no>
//<out_refund_no><![CDATA[2]]></out_refund_no>
//<refund_id><![CDATA[50000406362018031303826081322]]></refund_id>
//<refund_channel><![CDATA[]]></refund_channel>
//<refund_fee>3</refund_fee>
//<coupon_refund_fee>0</coupon_refund_fee>
//<total_fee>3</total_fee>
//<cash_fee>3</cash_fee>
//<coupon_refund_count>0</coupon_refund_count>
//<cash_refund_fee>3</cash_refund_fee>
//</xml>
func (r *Refund) ByTransactionId(tid, num string, total, refund int, options core.Map) core.Map {
	options = core.MapNilMake(options)
	options.NilSet("transaction_id", tid)
	return r.refund(num, total, refund, options).ToMap()
}

func (r *Refund) query(m core.Map) *core.Response {
	m.Set("appid", r.Config.Get("app_id"))
	return r.Request(REFUNDQUERY_URL_SUFFIX, m)
}

func (r *Refund) QueryByRefundId(id string) core.Map {
	return r.query(core.Map{"refund_id": id}).ToMap()
}

func (r *Refund) QueryByOutRefundNumber(id string) core.Map {
	return r.query(core.Map{"out_refund_no": id}).ToMap()
}

func (r *Refund) QueryByOutTradeNumber(id string) core.Map {
	return r.query(core.Map{"out_trade_no": id}).ToMap()
}

func (r *Refund) QueryByTransactionId(id string) core.Map {
	return r.query(core.Map{"transaction_id": id}).ToMap()
}
