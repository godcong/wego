package payment

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

/*Merchant 账单 */
type Merchant struct {
	*Payment
}

func newMerchant(p *Payment) *Merchant {
	return &Merchant{
		Payment: p,
	}
}

/*NewMerchant 账单 */
func NewMerchant(config *core.Config) *Merchant {
	return newMerchant(NewPayment(config))
}

func (m *Merchant) AddSubMerchant() {

}
func (m *Merchant) manage(action string, params util.Map) core.Response {

	return m.SafeRequest(Link(mchSubmchmanage), params)
}
