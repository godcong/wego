package app

// SandboxConfig ...
type SandboxProperty struct {
	AppID  string
	Secret string
}

// Payment ...
type PaymentProperty struct {
	AppID      string `xorm:"app_id" json:"app_id"`
	MchID      string `xorm:"mch_id" json:"mch_id"`
	Key        string `xorm:"key" json:"key"`
	NotifyURL  string `xorm:"notify_url" json:"notify_url"`
	RefundURL  string `xorm:"refund_url" json:"refund_url"`
	CertPath   string `xorm:"cert_path" json:"cert_path"`
	KeyPath    string `xorm:"key_path" json:"key_path"`
	RootCaPath string `xorm:"root_ca_path" json:"root_ca_path"`
	PublicKey  string `xorm:"public_key" json:"public_key"`
	PrivateKey string `xorm:"private_key" json:"private_key"`
}

// OAuth ...
type OAuthProperty struct {
	Scopes      []string `xorm:"scopes"`
	RedirectURI string   `xorm:"redirect_uri"`
}

// OpenPlatform ...
type OpenPlatformProperty struct {
	AppID  string `xorm:"app_id"`
	Secret string `xorm:"secret"`
	Token  string `xorm:"token"`
	AesKey string `xorm:"aes_key"`
}

// OfficialAccount ...
type OfficialAccountProperty struct {
	AppID  string `xorm:"app_id"`
	Secret string `xorm:"secret"`
	Token  string `xorm:"token"`
	AesKey string `xorm:"aes_key"`
}

// MiniProgram ...
type MiniProgramProperty struct {
	AppID  string `xorm:"app_id"`
	Secret string `xorm:"secret"`
	Token  string `xorm:"token"`
	AesKey string `xorm:"aes_key"`
}

// WeConfig ...
type Property struct {
	SandboxProperty
	OAuthProperty
	OpenPlatformProperty
	OfficialAccountProperty
	MiniProgramProperty
	PaymentProperty
}
