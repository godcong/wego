package wego

import (
	"github.com/godcong/wego/app/payment"
)

//JSSDK result payment JSSDK
func JSSDK() *payment.JSSDK {
	return app.Payment("payment.default").JSSDK()
}

//RedPack result payment RedPack
func RedPack() *payment.RedPack {
	return app.Payment("payment.default").RedPack()
}

//Bill result payment bill
func Bill() *payment.Bill {
	return app.Payment("payment.default").Bill()
}

//Security result payment security
func Security() *payment.Security {
	return app.Payment("payment.default").Security()
}

//Order result payment order
func Order() *payment.Order {
	return app.Payment("payment.default").Order()
}

//Refund result payment refund
func Refund() *payment.Refund {
	return app.Payment("payment.default").Refund()
}

//func Sandbox() *core.Sandbox {
//	return app.Payment("payment.default")
//}

//Reverse result payment reverse
func Reverse() *payment.Reverse {
	return app.Payment("payment.default").Reverse()
}

//Transfer result payment Transfer
func Transfer() *payment.Transfer {
	return app.Payment("payment.default").Transfer()
}

/*Payment result payment*/
func Payment() *payment.Payment {
	return app.Payment("payment.default")
}
