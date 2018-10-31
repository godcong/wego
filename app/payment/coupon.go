package payment

// Coupon ...
type Coupon struct {
	*Payment
}

func newCoupon(payment *Payment) *Coupon {
	return &Coupon{
		Payment: payment,
	}
}
