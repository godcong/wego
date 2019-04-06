package wego

import (
	"github.com/godcong/wego/util"
	"github.com/json-iterator/go"
)

// SandboxOption ...
type SandboxOption struct {
	SubMchID string `xml:"sub_mch_id"`
	SubAppID string `xml:"sub_app_id"`
}

// SandboxConfig ...
type SandboxConfig struct {
	AppID     string
	AppSecret string
	MchID     string
	Key       string
}

// SafeCertConfig ...
type SafeCertConfig struct {
	Cert   []byte
	Key    []byte
	RootCA []byte
}

// PaymentOption ...
type PaymentOption struct {
	BodyType   *BodyType      `xml:"body_type"`
	SubMchID   string         `xml:"sub_mch_id"`
	SubAppID   string         `xml:"sub_app_id"`
	PublicKey  string         `xml:"public_key"`
	PrivateKey string         `xml:"private_key"`
	RemoteHost string         `xml:"remote_host"`
	LocalHost  string         `xml:"local_host"`
	UseSandbox bool           `xml:"use_sandbox"`
	Sandbox    *SandboxConfig `xml:"sandbox"`

	NotifyURL   string `xml:"notify_url"`
	RefundedURL string `xml:"refunded_url"`
	ScannedURL  string `xml:"scanned_url"`
}

// PaymentConfig ...
type PaymentConfig struct {
	AppID     string          `xml:"app_id"`
	AppSecret string          `xml:"app_secret"`
	MchID     string          `xml:"mch_id"`
	Key       string          `xml:"key"`
	SafeCert  *SafeCertConfig `xml:"safe_cert"`
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

// OfficialAccountConfig ...
type OfficialAccountConfig struct {
	AppID       string
	AppSecret   string
	Token       string
	AesKey      string
	AccessToken *AccessTokenConfig
	OAuth       *OAuthConfig
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

// JSSDKConfig ...
type JSSDKConfig struct {
	AppID       string
	MchID       string
	Key         string
	AccessToken *AccessTokenConfig
}

// JSSDKOption ...
type JSSDKOption struct {
	SubAppID string
	URL      string
}

// Property 属性配置，各个接口用到的参数
type Property struct {
	JSSDK                 *JSSDKConfig
	JSSDKOption           *JSSDKOption
	AccessToken           *AccessTokenConfig
	AccessTokenOption     *AccessTokenOption
	OAuth                 *OAuthConfig
	OpenPlatform          *OpenPlatformProperty
	OfficialAccount       *OfficialAccountConfig
	OfficialAccountOption *OfficialAccountOption
	MiniProgram           *MiniProgramProperty
	Payment               *PaymentConfig
	PaymentOption         *PaymentOption
	SafeCert              *SafeCertConfig
}
