package wego

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/core/log"
	"github.com/godcong/wego/core/util"
)

type Security interface {
	GetPublicKey() *core.Response
}
type Order interface {
	Unify(m util.Map) *core.Response
	Close(no string) *core.Response
	//Query(Map) Map
	QueryByTransactionId(id string) *core.Response
	QueryByOutTradeNumber(no string) *core.Response
}

type JSSDK interface {
	BridgeConfig(pid string) util.Map
	SdkConfig(pid string) util.Map
	AppConfig(pid string) util.Map
	ShareAddressConfig(accessToken interface{}) util.Map
}

type Bill interface {
	Get(string, string, util.Map) *core.Response
}

type RedPack interface {
	Info(util.Map) *core.Response
	SendNormal(util.Map) *core.Response
	SendGroup(util.Map) *core.Response
}

type Refund interface {
	ByOutTradeNumber(tradeNum, num string, total, refund int, options util.Map) *core.Response
	ByTransactionId(tid, num string, total, refund int, options util.Map) *core.Response
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
	ToBalance(util.Map) *core.Response
	QueryBankCardOrder(string) *core.Response
	ToBankCard(util.Map) *core.Response
}
type Payment interface {
	//Sandbox() Sandbox

	Order() Order
	Refund() Refund
	Security() Security

	Pay(util.Map) util.Map
	Request(url string, m util.Map) *core.Response
	RequestRaw(url string, m util.Map) *core.Response
	SafeRequest(url string, m util.Map) *core.Response
	AuthCodeToOpenid(string) util.Map
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
	log.Debug("GetPayment|obj:", obj)
	return obj
	//obj := new(payment.Payment)
	//GetApp().Get(obj)
	//log.Debug("GetPayment|obj:", obj)
	//return obj
}

func GetSecurity() Security {
	obj := GetPayment().Security()
	log.Debug("GetSecurity|obj:", obj)
	return obj
}

func GetOrder() Order {
	obj := GetPayment().Order()
	log.Debug("GetOrder|obj:", obj)
	return obj
}

func GetRefund() Refund {
	obj := GetPayment().Refund()
	log.Debug("GetRefund|obj:", obj)
	return obj
}
