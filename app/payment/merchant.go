package payment

import "github.com/godcong/wego/core"

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
