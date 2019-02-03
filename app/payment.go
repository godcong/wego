package app

//Payment ...
type Payment struct {
	option      *PaymentOption
	property    *PaymentProperty
	sandbox     *SandboxProperty
	accessToken *AccessToken
}

// PaymentOption ...
type PaymentOption struct {
}

// NewPayment ...
func NewPayment(property *Property, opts ...*PaymentOption) *Payment {
	var opt *PaymentOption
	if opts != nil {
		opt = opts[0]
	}
	return &Payment{
		option:      opt,
		property:    &property.PaymentProperty,
		sandbox:     &property.SandboxProperty,
		accessToken: nil,
	}

}
