package app

// SandboxProperty ...
type SandboxProperty struct {
	UseSandbox bool
	AppID      string
	Secret     string
	MchID      string
	Key        string
}

// PaymentProperty ...
type PaymentProperty struct {
	AppID      string
	MchID      string
	Key        string
	CertPEM    string
	KeyPEM     string
	RootCaPEM  string
	PublicKey  string
	PrivateKey string
}

// OAuthProperty ...
type OAuthProperty struct {
	Scopes      []string
	RedirectURI string
}

// OpenPlatformProperty ...
type OpenPlatformProperty struct {
	AppID  string
	Secret string
	Token  string
	AesKey string
}

// OfficialAccountProperty ...
type OfficialAccountProperty struct {
	AppID  string
	Secret string
	Token  string
	AesKey string
}

// MiniProgramProperty ...
type MiniProgramProperty struct {
	AppID  string
	Secret string
	Token  string
	AesKey string
}

// LocalProperty ...
type LocalProperty struct {
	Address     string `toml:"address"`
	NotifyURL   string `toml:"paid_url"`
	RefundedURL string `toml:"refunded_url"`
	ScannedURL  string `toml:"scanned_url"`
}

// Property ...
type Property struct {
	Local           *LocalProperty
	OAuth           *OAuthProperty
	OpenPlatform    *OpenPlatformProperty
	OfficialAccount *OfficialAccountProperty
	MiniProgram     *MiniProgramProperty
	Payment         *PaymentProperty
}
