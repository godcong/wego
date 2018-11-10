package payment

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

/*Notify 监听 */
type Notify struct {
	//*Payment
}

func newNotify(p *Payment) interface{} {
	return &Notify{
		//Payment: p,
	}
}

/*NewNotify 监听 */
func NewNotify(config *core.Config) *Notify {
	return newNotify(NewPayment(config)).(*Notify)
}

//RefundedNotify ...
func (n *Notify) RefundedNotify(p util.Map) {

}

//ScannedNotify ...
func (n *Notify) ScannedNotify(p util.Map) {

}

//PaidNotify ...
func (n *Notify) PaidNotify(p util.Map) {

}
