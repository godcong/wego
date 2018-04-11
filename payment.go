package wego

import (
	"github.com/godcong/wego/core"
)

type Security interface {
	GetPublicKey() *core.Response
}
type Order interface {
	Unify(m core.Map) *core.Response
	Close(no string) *core.Response
	//Query(Map) Map
	QueryByTransactionId(id string) *core.Response
	QueryByOutTradeNumber(no string) *core.Response
}

type JSSDK interface {
	BridgeConfig(pid string) core.Map
	SdkConfig(pid string) core.Map
	AppConfig(pid string) core.Map
	ShareAddressConfig(accessToken interface{}) core.Map
}

type Bill interface {
	Get(string, string, core.Map) *core.Response
}

type RedPack interface {
	Info(core.Map) *core.Response
	SendNormal(core.Map) *core.Response
	SendGroup(core.Map) *core.Response
}

type Refund interface {
	ByOutTradeNumber(tradeNum, num string, total, refund int, options core.Map) *core.Response
	ByTransactionId(tid, num string, total, refund int, options core.Map) *core.Response
	QueryByRefundId(id string) *core.Response
	QueryByOutRefundNumber(id string) *core.Response
	QueryByOutTradeNumber(id string) *core.Response
	QueryByTransactionId(id string) *core.Response

	//Refund(string, int, int, Map) Map
	//Query(Map) Map
}

type Reverse interface {
	ByOutTradeNumber(string) *core.Response
	ByTransactionId(string) *core.Response
}

type Sandbox interface {
	GetKey() string
	GetCacheKey() string
}
type Transfer interface {
	QueryBalanceOrder(string) *core.Response
	ToBalance(core.Map) *core.Response
	QueryBankCardOrder(string) *core.Response
	ToBankCard(core.Map) *core.Response
}
type Payment interface {
	//Sandbox() Sandbox

	Order() Order
	Refund() Refund
	Security() Security

	Pay(core.Map) core.Map
	Request(url string, m core.Map) *core.Response
	RequestRaw(url string, m core.Map) *core.Response
	SafeRequest(url string, m core.Map) *core.Response
	AuthCodeToOpenid(string) core.Map
}

//
//func NewJSSDK(application Application, config core.Config) *payment.JSSDK {
//	return &payment.JSSDK{
//		Config: config,
//		//app:    application,
//	}
//}
//
//func NewRedPack(application core.Application, config core.Config) *payment.RedPack {
//	return &payment.RedPack{
//		Config: config,
//		//Payment: application.Payment(),
//	}
//}
//
//func NewBill(application Application, config core.Config) *payment.Bill {
//	return &payment.Bill{
//		Config: config,
//		//Payment: application.Payment(),
//	}
//}
//
//func NewSecurity(application Application, config core.Config) *payment.Security {
//	return &payment.Security{
//		Config: config,
//		//Payment: application.Payment(),
//	}
//}
//
//func NewOrder(application Application, config core.Config) *payment.Order {
//	return &payment.Order{
//		//app:    application,
//		//Payment: application.Payment(),
//		Config: config,
//	}
//}
//
//func NewSandbox(application Application, config core.Config) *core.Sandbox {
//	return &core.Sandbox{
//		Config: config,
//		//app:    application,
//	}
//}
//
//func NewReverse(application Application, config core.Config) *payment.Reverse {
//	return &payment.Reverse{
//		Config: config,
//		//Payment: application.Payment(),
//	}
//}
//
//func NewTransfer(application Application, config core.Config) *payment.Transfer {
//	return &payment.Transfer{
//		Config: config,
//		//Payment: application.Payment(),
//	}
//}

func GetPayment() Payment {
	obj := GetApp().Get("payment").(Payment)
	core.Debug("GetPayment|obj:", obj)
	return obj
	//obj := new(payment.Payment)
	//GetApp().Get(obj)
	//core.Debug("GetPayment|obj:", obj)
	//return obj
}

func GetSecurity() Security {
	obj := GetPayment().Security()
	core.Debug("GetSecurity|obj:", obj)
	return obj
}

func GetOrder() Order {
	obj := GetPayment().Order()
	core.Debug("GetOrder|obj:", obj)
	return obj
}

func GetRefund() Refund {
	obj := GetPayment().Refund()
	core.Debug("GetRefund|obj:", obj)
	return obj
}
