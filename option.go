package wego

// PaymentOption ...
type PaymentOption func(obj *Payment)

// PaymentBodyType ...
func PaymentBodyType(s BodyType) PaymentOption {
	return func(obj *Payment) {
		obj.BodyType = s
	}
}

// PaymentSubID ...
func PaymentSubID(mch, app string) PaymentOption {
	return func(obj *Payment) {
		obj.subAppID = app
		obj.subMchID = mch
	}
}

// PaymentKey ...
func PaymentKey(public, privite string) PaymentOption {
	return func(obj *Payment) {
		obj.publicKey = public
		obj.privateKey = privite
	}
}

// PaymentRemote ...
func PaymentRemote(remote string) PaymentOption {
	return func(obj *Payment) {
		obj.remoteHost = remote
	}
}

// PaymentLocal ...
func PaymentLocal(local string) PaymentOption {
	return func(obj *Payment) {
		obj.localHost = local
	}
}

// PaymentSandbox ...
func PaymentSandbox(property *SandboxProperty) PaymentOption {
	return func(obj *Payment) {
		obj.useSandbox = true
		obj.sandbox = NewSandbox(property)
	}
}

// PaymentNotifyURL ...
func PaymentNotifyURL(s string) PaymentOption {
	return func(obj *Payment) {
		obj.notifyURL = s
	}
}

// PaymentRefundedURL ...
func PaymentRefundedURL(s string) PaymentOption {
	return func(obj *Payment) {
		obj.refundedURL = s
	}
}

// PaymentScannedURL ...
func PaymentScannedURL(s string) PaymentOption {
	return func(obj *Payment) {
		obj.scannedURL = s
	}
}

// PaymentOptions ...
type PaymentOptions struct {
	BodyType   BodyType         `xml:"body_type"`
	SubMchID   string           `xml:"sub_mch_id"`
	SubAppID   string           `xml:"sub_app_id"`
	PublicKey  string           `xml:"public_key"`
	PrivateKey string           `xml:"private_key"`
	RemoteHost string           `xml:"remote_host"`
	LocalHost  string           `xml:"local_host"`
	UseSandbox bool             `xml:"use_sandbox"`
	Sandbox    *SandboxProperty `xml:"sandbox"`

	NotifyURL   string `xml:"notify_url"`
	RefundedURL string `xml:"refunded_url"`
	ScannedURL  string `xml:"scanned_url"`
}
