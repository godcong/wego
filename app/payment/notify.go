package payment

import "github.com/godcong/wego/core"

/*Notify 账单 */
type Notify struct {
	*Payment
}

func newNotify(p *Payment) interface{} {
	return &Notify{
		Payment: p,
	}
}

/*NewNotify 账单 */
func NewNotify(config *core.Config) *Notify {
	return newNotify(NewPayment(config)).(*Notify)
}
