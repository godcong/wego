package app

// SandboxProperty ...
type SandboxProperty struct {
	AppID  string
	Secret string
	MchID  string
	Key    string
}

// PaymentProperty ...
type PaymentProperty struct {
	AppID      string
	MchID      string
	Key        string
	NotifyURL  string
	RefundURL  string
	CertPath   string
	KeyPath    string
	RootCaPath string
	PublicKey  string
	PrivateKey string
}

// OAuthProperty ...
type OAuthProperty struct {
	Scopes      []string `xorm:"scopes"`
	RedirectURI string   `xorm:"redirect_uri"`
}

// OpenPlatformProperty ...
type OpenPlatformProperty struct {
	AppID  string `xorm:"app_id"`
	Secret string `xorm:"secret"`
	Token  string `xorm:"token"`
	AesKey string `xorm:"aes_key"`
}

// OfficialAccountProperty ...
type OfficialAccountProperty struct {
	AppID  string `xorm:"app_id"`
	Secret string `xorm:"secret"`
	Token  string `xorm:"token"`
	AesKey string `xorm:"aes_key"`
}

// MiniProgramProperty ...
type MiniProgramProperty struct {
	AppID  string `xorm:"app_id"`
	Secret string `xorm:"secret"`
	Token  string `xorm:"token"`
	AesKey string `xorm:"aes_key"`
}

// Property ...
type Property struct {
	Sandbox         *SandboxProperty
	OAuth           *OAuthProperty
	OpenPlatform    *OpenPlatformProperty
	OfficialAccount *OfficialAccountProperty
	MiniProgram     *MiniProgramProperty
	Payment         *PaymentProperty
}
