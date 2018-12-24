package app

// SystemConfig ...
type SystemConfig struct {
	Debug bool
	Cache string
}

// LogConfig ...
type LogConfig struct {
	Path string
}

// SandboxConfig ...
type SandboxConfig struct {
	AppID  string
	Secret string
}

// BaseConfig ...
type BaseConfig struct {
	AppID  string
	Secret string
	Token  string
	AesKey string
}

// HTTPConfig ...
type HTTPConfig struct {
	TimeOut   int
	KeepAlive int
}

// OfficialAccountConfig ...
type OfficialAccountConfig struct {
	Sandbox SandboxConfig
	def     BaseConfig
}

// PaymentConfig ...
type PaymentConfig struct {
	Sandbox        bool
	AppID          string
	MerchantID     string
	Key            string
	NotifyURL      string
	RefundURL      string
	CertPath       string
	KeyPath        string
	RootCAPath     string
	PublicKeyPath  string
	PrivateKeyPath string
}

// WeConfig ...
type WeConfig struct {
}
