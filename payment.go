package wego

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/payment"
)

type Security interface {
	GetPublicKey() core.Map
}
type Order interface {
	Unify(m core.Map) core.Map
	Close(no string) core.Map
	//Query(Map) Map
	QueryByTransactionId(id string) core.Map
	QueryByOutTradeNumber(no string) core.Map
}

type JSSDK interface {
	BridgeConfig(pid string) core.Map
	SdkConfig(pid string) core.Map
	AppConfig(pid string) core.Map
	ShareAddressConfig(accessToken interface{}) core.Map
}

type Bill interface {
	Get(string, string, core.Map) core.Map
}

type RedPack interface {
	Info(core.Map) core.Map
	SendNormal(core.Map) core.Map
	SendGroup(core.Map) core.Map
}

type Refund interface {
	ByOutTradeNumber(tradeNum, num string, total, refund int, options core.Map) core.Map
	ByTransactionId(tid, num string, total, refund int, options core.Map) core.Map
	QueryByRefundId(id string) core.Map
	QueryByOutRefundNumber(id string) core.Map
	QueryByOutTradeNumber(id string) core.Map
	QueryByTransactionId(id string) core.Map

	//Refund(string, int, int, Map) Map
	//Query(Map) Map
}

type Reverse interface {
	ByOutTradeNumber(string) core.Map
	ByTransactionId(string) core.Map
}

type Sandbox interface {
	GetKey() string
	GetCacheKey() string
}
type Transfer interface {
	QueryBalanceOrder(string) core.Map
	ToBalance(core.Map) core.Map
	QueryBankCardOrder(string) core.Map
	ToBankCard(core.Map) core.Map
}
type Payment interface {
	//Sandbox() Sandbox
	Order() *payment.Order
	Refund() *payment.Refund
	Security() *payment.Security
	Pay(core.Map) core.Map
	Request(url string, m core.Map) core.Map
	RequestRaw(url string, m core.Map) []byte
	SafeRequest(url string, m core.Map) core.Map
	AuthCodeToOpenid(string) core.Map
}

func NewJSSDK(application Application, config core.Config) *payment.JSSDK {
	return &payment.JSSDK{
		Config: config,
		//app:    application,
	}
}

func NewRedPack(application core.Application, config core.Config) *payment.RedPack {
	return &payment.RedPack{
		Config: config,
		//Payment: application.Payment(),
	}
}

func NewBill(application Application, config core.Config) *payment.Bill {
	return &payment.Bill{
		Config: config,
		//Payment: application.Payment(),
	}
}

func NewSecurity(application Application, config core.Config) *payment.Security {
	return &payment.Security{
		Config: config,
		//Payment: application.Payment(),
	}
}

func NewOrder(application Application, config core.Config) *payment.Order {
	return &payment.Order{
		//app:    application,
		//Payment: application.Payment(),
		Config: config,
	}
}

func NewSandbox(application Application, config core.Config) *core.Sandbox {
	return &core.Sandbox{
		Config: config,
		//app:    application,
	}
}

func NewReverse(application Application, config core.Config) *payment.Reverse {
	return &payment.Reverse{
		Config: config,
		//Payment: application.Payment(),
	}
}

func NewTransfer(application Application, config core.Config) *payment.Transfer {
	return &payment.Transfer{
		Config: config,
		//Payment: application.Payment(),
	}
}

func GetPayment() Payment {
	obj := GetApp().Get("payment").(Payment)
	return obj
}

func GetSecurity() Security {
	payment := GetApp().Get("payment").(Payment)
	return payment.Security()
}

func GetOrder() Order {
	payment := GetApp().Get("payment").(Payment)
	return payment.Order()
}

func GetRefund() Refund {
	payment := GetApp().Get("payment").(Payment)
	return payment.Refund()
}
