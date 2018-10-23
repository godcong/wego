package wego

import (
	"github.com/godcong/wego/app/payment"
)

//
///*Security 安全*/
//type Security interface {
//	GetPublicKey() core.Response
//}
//
///*Order 订单*/
//type Order interface {
//	Unify(m util.Map) core.Response
//	Close(no string) core.Response
//	//Query(Map) Map
//	QueryByTransactionID(id string) core.Response
//	QueryByOutTradeNumber(no string) core.Response
//}
//
///*JSSDK jssdk*/
//type JSSDK interface {
//	BridgeConfig(pid string) util.Map
//	SdkConfig(pid string) util.Map
//	AppConfig(pid string) util.Map
//	ShareAddressConfig(accessToken interface{}) util.Map
//}
//
///*Bill bill*/
//type Bill interface {
//	Get(string, string, util.Map) core.Response
//}
//
///*RedPack 红包*/
//type RedPack interface {
//	Info(util.Map) core.Response
//	SendNormal(util.Map) core.Response
//	SendGroup(util.Map) core.Response
//}
//
///*Refund 退款*/
//type Refund interface {
//	ByOutTradeNumber(tradeNum, num string, total, refund int, options util.Map) core.Response
//	ByTransactionID(tid, num string, total, refund int, options util.Map) core.Response
//	QueryByRefundID(id string) core.Response
//	QueryByOutRefundNumber(id string) core.Response
//	QueryByOutTradeNumber(id string) core.Response
//	QueryByTransactionID(id string) core.Response
//
//	//Refund(string, int, int, Map) Map
//	//Query(Map) Map
//}
//
///*Reverse reverse*/
//type Reverse interface {
//	ByOutTradeNumber(string) core.Response
//	ByTransactionId(string) core.Response
//}
//
///*Sandbox 沙箱*/
//type Sandbox interface {
//	GetKey() string
//	GetCacheKey() string
//}
//
///*Transfer 转账*/
//type Transfer interface {
//	QueryBalanceOrder(string) core.Response
//	ToBalance(util.Map) core.Response
//	QueryBankCardOrder(string) core.Response
//	ToBankCard(util.Map) core.Response
//}
//
///*Payment 支付*/
//type Payment interface {
//	//Sandbox() Sandbox
//
//	Order() Order
//	Refund() Refund
//	Security() Security
//
//	Pay(util.Map) util.Map
//	Request(url string, m util.Map) core.Response
//	RequestRaw(url string, m util.Map) core.Response
//	SafeRequest(url string, m util.Map) core.Response
//	AuthCodeToOpenid(string) util.Map
//}

//
//func NewJSSDK(application Application, config Config) *payment.JSSDK {
//	return &payment.JSSDK{
//		Config: config,
//		//app:    application,
//	}
//}
//
//func NewRedPack(application core.Application, config Config) *payment.RedPack {
//	return &payment.RedPack{
//		Config: config,
//		//Payment: application.Payment(),
//	}
//}
//
//func NewBill(application Application, config Config) *payment.Bill {
//	return &payment.Bill{
//		Config: config,
//		//Payment: application.Payment(),
//	}
//}
//
//func NewSecurity(application Application, config Config) *payment.Security {
//	return &payment.Security{
//		Config: config,
//		//Payment: application.Payment(),
//	}
//}
//
//func NewOrder(application Application, config Config) *payment.Order {
//	return &payment.Order{
//		//app:    application,
//		//Payment: application.Payment(),
//		Config: config,
//	}
//}
//
//func NewSandbox(application Application, config Config) *core.Sandbox {
//	return &core.Sandbox{
//		Config: config,
//		//app:    application,
//	}
//}
//
//func NewReverse(application Application, config Config) *payment.Reverse {
//	return &payment.Reverse{
//		Config: config,
//		//Payment: application.Payment(),
//	}
//}
//
//func NewTransfer(application Application, config Config) *payment.Transfer {
//	return &payment.Transfer{
//		Config: config,
//		//Payment: application.Payment(),
//	}
//}

/*GetPayment 获取支付*/
func Payment() *payment.Payment {
	var p *payment.Payment
	b := App().Get("payment", &p)
	if b {
		return p
	}
	p = payment.NewPayment(App().Config().GetSubConfig("payment.default"), App().Client(), App().AccessToken())
	App().Register("payment", p)
	return p
}
