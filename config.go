package wego

import (
	"github.com/godcong/wego/util"
	"github.com/json-iterator/go"
	"github.com/pelletier/go-toml"
	"log"
)

// SandboxOption ...
type SandboxOption struct {
	SubMchID string `xml:"sub_mch_id"`
	SubAppID string `xml:"sub_app_id"`
}

// SandboxConfig ...
type SandboxConfig struct {
	AppID  string
	Secret string
	MchID  string
	Key    string
	//option *SandboxOption
}

// SafeCertConfig ...
type SafeCertConfig struct {
	Cert   []byte
	Key    []byte
	RootCA []byte
}

// PaymentOption ...
type PaymentOption struct {
	BodyType    *BodyType       `xml:"body_type"`
	SubMchID    string          `xml:"sub_mch_id"`
	SubAppID    string          `xml:"sub_app_id"`
	PublicKey   string          `xml:"public_key"`
	PrivateKey  string          `xml:"private_key"`
	RemoteHost  string          `xml:"remote_host"`
	LocalHost   string          `xml:"local_host"`
	UseSandbox  bool            `xml:"use_sandbox"`
	Sandbox     *SandboxConfig  `xml:"sandbox"`
	SafeCert    *SafeCertConfig `xml:"safe_cert"`
	NotifyURL   string          `xml:"notify_url"`
	RefundedURL string          `xml:"refunded_url"`
	ScannedURL  string          `xml:"scanned_url"`
}

// PaymentConfig ...
type PaymentConfig struct {
	AppID     string `xml:"app_id"`
	AppSecret string `xml:"app_secret"`
	MchID     string `xml:"mch_id"`
	Key       string `xml:"key"`
	CertPEM   string `xml:"cert_pem"`
	KeyPEM    string `xml:"key_pem"`
	RootCaPEM string `xml:"root_ca_pem"`
}

// OAuthConfig ...
type OAuthConfig struct {
	Scopes      []string
	RedirectURI string
}

// OpenPlatformProperty ...
type OpenPlatformProperty struct {
	AppID     string
	AppSecret string
	Token     string
	AesKey    string
}

// OfficialAccountProperty ...
type OfficialAccountProperty struct {
	AppID     string
	AppSecret string
	Token     string
	AesKey    string
}

// MiniProgramProperty ...
type MiniProgramProperty struct {
	AppID     string
	AppSecret string
	Token     string
	AesKey    string
}

// GrantTypeClient ...
const GrantTypeClient string = "client_credential"

// AccessTokenConfig ...
type AccessTokenConfig struct {
	GrantType string `toml:"grant_type"`
	AppID     string `toml:"app_id"`
	AppSecret string `toml:"app_secret"`
}

// ToMap ...
func (obj *AccessTokenConfig) ToMap() util.Map {
	return util.Map{
		"grant_type": obj.GrantType,
		"appid":      obj.AppID,
		"secret":     obj.AppSecret,
	}
}

// ToJSON ...
func (obj *AccessTokenConfig) ToJSON() []byte {
	bytes, err := jsoniter.Marshal(obj)
	if err != nil {
		return nil
	}
	return bytes
}

// Config ...
type Config struct {
	AccessToken     *AccessTokenConfig
	OAuth           *OAuthConfig
	OpenPlatform    *OpenPlatformProperty
	OfficialAccount *OfficialAccountProperty
	MiniProgram     *MiniProgramProperty
	Payment         *PaymentConfig
	PaymentOption   *PaymentOption
	SafeCert        *SafeCertConfig
}

// LocalConfig ...
type LocalConfig struct {
}

// DefaultConfig ...
func DefaultConfig() *Config {
	return &Config{}
}

// LoadConfig ...
func LoadConfig(path string) *Config {
	cfg := DefaultConfig()
	t, e := toml.LoadFile(path)
	if e != nil {
		log.Println("filepath: " + path)
		log.Println(e.Error())
		return DefaultConfig()
	}

	e = t.Unmarshal(cfg)
	if e != nil {
		log.Println("filepath: " + path)
		log.Println(e.Error())
		return DefaultConfig()
	}

	return cfg

}
