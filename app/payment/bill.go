package payment

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
	"strconv"
)

/*Bill 账单 */
type Bill struct {
	*Payment
}

func newBill(p *Payment) interface{} {
	return &Bill{
		Payment: p,
	}
}

/*NewBill 账单 */
func NewBill(config *core.Config) *Bill {
	return newBill(NewPayment(config)).(*Bill)
}

//Download 下载对账单
//接口链接
//https://api.mch.weixin.qq.com/pay/downloadbill
//是否需要证书
//不需要。
//请求参数
//字段名	变量名	必填	类型	示例值	描述
//对账单日期	bill_date	是	String(8)	20140603	下载对账单的日期，格式:20140603
func (b *Bill) Download(bd string, option ...util.Map) core.Responder {
	m := util.MapsToMap(util.Map{
		"appid":     b.Get("app_id"),
		"bill_date": bd,
	}, option)

	if !m.Has("bill_type") {
		m.Set("bill_type", "ALL")
	}

	return b.Request(payDownloadBill, m)
}

//DownloadFundFlow 下载资金账单
//资金账单日期	bill_date	是	String(8)	20140603	下载对账单的日期，格式：20140603
//资金账户类型	account_type	是	String(8)	Basic
//账单的资金来源账户：
//Basic  基本账户
//Operation 运营账户
//Fees 手续费账户
//压缩账单	tar_type	否	String(8)	GZIP	非必传参数，固定值：GZIP，返回格式为.gzip的压缩包账单。不传则默认为数据流形式。
func (b *Bill) DownloadFundFlow(bd string, at string, option ...util.Map) core.Responder {
	m := util.MapsToMap(util.Map{
		"appid":        b.Get("app_id"),
		"bill_date":    bd,
		"sign_type":    HMACSHA256,
		"account_type": at,
	}, option)

	return b.SafeRequest(payDownloadfundflow, m)
}

//BatchQueryComment 拉取订单评价数据
//接口链接
//https://api.mch.weixin.qq.com/billcommentsp/batchquerycomment
//开始时间 begin_time 是 String(19) 20170724000000 按用户评论时间批量拉取的起始时间，格式为yyyyMMddHHmmss
//结束时间 end_time 是 String(19) 20170725000000 按用户评论时间批量拉取的结束时间，格式为yyyyMMddHHmmss
//位移 offset 是 uint(64) 0 指定从某条记录的下一条开始返回记录。接口调用成功时，会返回本次查询最后一条数据的offset。商户需要翻页时，应该把本次调用返回的offset 作为下次调用的入参。注意offset是评论数据在微信支付后台保存的索引，未必是连续的
//条数 limit 否 uint(32) 100 一次拉取的条数, 最大值是200，默认是200
func (b *Bill) BatchQueryComment(beginTime, endTime string, offset int, option ...util.Map) core.Responder {
	m := util.MapsToMap(util.Map{
		"begin_time": beginTime,
		"end_time":   endTime,
		"offset":     strconv.Itoa(offset),
		"appid":      b.Get("app_id"),
		"sign_type":  HMACSHA256,
	}, option)

	return core.Request(b.Link(batchQueryComment), "post", util.Map{
		core.DataTypeXML:      b.initRequestWithIgnore(m, FieldSign, FieldSignType, FieldLimit),
		core.DataTypeSecurity: b.Config,
	})

}
