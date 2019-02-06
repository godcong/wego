package _bak

import (
	"github.com/godcong/wego/app/payment"
)

//PaymentJSSDK result payment JSSDK
func PaymentJSSDK() *payment.JSSDK {
	return app.Payment("payment.default").JSSDK()
}

//PaymentRedPack result payment RedPack
func PaymentRedPack() *payment.RedPack {
	return app.Payment("payment.default").RedPack()
}

//PaymentBill result payment bill
func PaymentBill() *payment.Bill {
	return app.Payment("payment.default").Bill()
}

//PaymentSecurity result payment Security
func PaymentSecurity() *payment.Security {
	return app.Payment("payment.default").Security()
}

//PaymentOrder result payment Order
func PaymentOrder() *payment.Order {
	return app.Payment("payment.default").Order()
}

//PaymentRefund result payment Refund
func PaymentRefund() *payment.Refund {
	return app.Payment("payment.default").Refund()
}

//PaymentSandbox result payment Sandbox
func PaymentSandbox() *payment.Sandbox {
	return app.Payment("payment.default").Sandbox()
}

//PaymentReverse result payment Reverse
func PaymentReverse() *payment.Reverse {
	return app.Payment("payment.default").Reverse()
}

//PaymentTransfer result payment Transfer
func PaymentTransfer() *payment.Transfer {
	return app.Payment("payment.default").Transfer()
}

/*Payment result payment self*/
func Payment() *payment.Payment {
	return app.Payment("payment.default")
}
