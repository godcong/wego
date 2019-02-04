package app

// Configure ...
type Configure struct {
	LocalHost   string
	PaidURL     string
	RefundedURL string
	ScannedURL  string
}

// Config ...
func Config() *Configure {
	return &Configure{
		LocalHost:   "http://localhost",
		PaidURL:     "paid_cb",
		RefundedURL: "refunded_cb",
		ScannedURL:  "scanned_cb",
	}
}
