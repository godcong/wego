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
func (b *Bill) Download(bd string, maps ...util.Map) core.Response {
	m := util.MapsToMap(maps)
	m.Set("appid", b.Get("app_id"))
	m.Set("bill_date", bd)
	if !m.Has("bill_type") {
		m.Set("bill_type", "ALL")
	}

	return b.Request(payDownloadBill, m)
}

// DownloadFundFlow ...
func (b *Bill) DownloadFundFlow(bd string, at string, maps ...util.Map) core.Response {
	m := util.MapsToMap(maps)
	m.Set("appid", b.Get("app_id"))
	m.Set("bill_date", bd)
	m.Set("sign_type", HMACSHA256)
	m.Set("account_type", at)

	return b.SafeRequest(payDownloadfundflow, m)
}

//BatchQueryComment 拉取订单评价数据
//接口链接
//https://api.mch.weixin.qq.com/billcommentsp/batchquerycomment
//开始时间 begin_time 是 String(19) 20170724000000 按用户评论时间批量拉取的起始时间，格式为yyyyMMddHHmmss
//结束时间 end_time 是 String(19) 20170725000000 按用户评论时间批量拉取的结束时间，格式为yyyyMMddHHmmss
//位移 offset 是 uint(64) 0 指定从某条记录的下一条开始返回记录。接口调用成功时，会返回本次查询最后一条数据的offset。商户需要翻页时，应该把本次调用返回的offset 作为下次调用的入参。注意offset是评论数据在微信支付后台保存的索引，未必是连续的
//条数 limit 否 uint(32) 100 一次拉取的条数, 最大值是200，默认是200
func (b *Bill) BatchQueryComment(beginTime, endTime string, offset int, maps ...util.Map) core.Response {
	m := util.MapsToMap(maps)
	m.Set("begin_time", beginTime)
	m.Set("end_time", endTime)
	m.Set("offset", strconv.Itoa(offset))
	m.Set("appid", b.Get("app_id"))
	m.Set("sign_type", HMACSHA256)
	//m.Set("sign_type", MD5)
	//if !m.Has("limit") {
	//	m.Set("limit", 200)
	return b.client.Request(b.Link(batchQueryComment), "post", util.Map{
		core.DataTypeXML:      b.initRequestWithIgnore(m, FieldSign, FieldSignType, FieldLimit),
		core.DataTypeSecurity: b.Config,
	})

}
