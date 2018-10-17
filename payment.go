package wego

import (
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/net"
	"github.com/godcong/wego/util"
)

/*Security 安全*/
type Security interface {
	GetPublicKey() *net.Response
}

/*Order 订单*/
type Order interface {
	Unify(m util.Map) *net.Response
	Close(no string) *net.Response
	//Query(Map) Map
	QueryByTransactionID(id string) *net.Response
	QueryByOutTradeNumber(no string) *net.Response
}

/*JSSDK jssdk*/
type JSSDK interface {
	BridgeConfig(pid string) util.Map
	SdkConfig(pid string) util.Map
	AppConfig(pid string) util.Map
	ShareAddressConfig(accessToken interface{}) util.Map
}

/*Bill bill*/
type Bill interface {
	Get(string, string, util.Map) *net.Response
}

/*RedPack 红包*/
type RedPack interface {
	Info(util.Map) *net.Response
	SendNormal(util.Map) *net.Response
	SendGroup(util.Map) *net.Response
}

/*Refund 退款*/
type Refund interface {
	ByOutTradeNumber(tradeNum, num string, total, refund int, options util.Map) *net.Response
	ByTransactionID(tid, num string, total, refund int, options util.Map) *net.Response
	QueryByRefundID(id string) *net.Response
	QueryByOutRefundNumber(id string) *net.Response
	QueryByOutTradeNumber(id string) *net.Response
	QueryByTransactionID(id string) *net.Response

	//Refund(string, int, int, Map) Map
	//Query(Map) Map
}

/*Reverse reverse*/
type Reverse interface {
	ByOutTradeNumber(string) *net.Response
	ByTransactionId(string) *net.Response
}

/*Sandbox 沙箱*/
type Sandbox interface {
	GetKey() string
	GetCacheKey() string
}

/*Transfer 转账*/
type Transfer interface {
	QueryBalanceOrder(string) *net.Response
	ToBalance(util.Map) *net.Response
	QueryBankCardOrder(string) *net.Response
	ToBankCard(util.Map) *net.Response
}

/*Payment 支付*/
type Payment interface {
	//Sandbox() Sandbox

	Order() Order
	Refund() Refund
	Security() Security

	Pay(util.Map) util.Map
	Request(url string, m util.Map) *net.Response
	RequestRaw(url string, m util.Map) *net.Response
	SafeRequest(url string, m util.Map) *net.Response
	AuthCodeToOpenid(string) util.Map
}

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
func GetPayment() Payment {
	obj := GetApp().Get("payment").(Payment)
	log.Debug("GetPayment|obj:", obj)
	return obj
	//obj := new(payment.Payment)
	//GetApp().Get(obj)
	//log.Debug("GetPayment|obj:", obj)
	//return obj
}

/*GetSecurity 获取安全*/
func GetSecurity() Security {
	obj := GetPayment().Security()
	log.Debug("GetSecurity|obj:", obj)
	return obj
}

/*GetOrder 获取订单*/
func GetOrder() Order {
	obj := GetPayment().Order()
	log.Debug("GetOrder|obj:", obj)
	return obj
}

/*GetRefund 获取退款*/
func GetRefund() Refund {
	obj := GetPayment().Refund()
	log.Debug("GetRefund|obj:", obj)
	return obj
}
