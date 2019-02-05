package app

import (
	"github.com/godcong/wego/util"
	"github.com/json-iterator/go"
)

// SandboxProperty ...
type SandboxProperty struct {
	AppID  string
	Secret string
	MchID  string
	Key    string
}

// PaymentProperty ...
type PaymentProperty struct {
	AppID      string `json:"app_id"`
	MchID      string `json:"mch_id"`
	SubMchID   string `json:"sub_mch_id"`
	SubAppID   string `json:"sub_app_id"`
	Key        string `json:"key"`
	CertPEM    string `json:"cert_pem"`
	KeyPEM     string `json:"key_pem"`
	RootCaPEM  string `json:"root_ca_pem"`
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
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

// GrantTypeClient ...
const GrantTypeClient string = "client_credential"

// AccessTokenProperty ...
type AccessTokenProperty struct {
	GrantType string `toml:"grant_type"`
	AppID     string `toml:"app_id"`
	Secret    string `toml:"secret"`
}

// ToMap ...
func (obj *AccessTokenProperty) ToMap() util.Map {
	return util.Map{
		"grant_type": obj.GrantType,
		"appid":      obj.AppID,
		"secret":     obj.Secret,
	}
}

// ToJSON ...
func (obj *AccessTokenProperty) ToJSON() []byte {
	bytes, err := jsoniter.Marshal(obj)
	if err != nil {
		return nil
	}
	return bytes
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
	AccessToken     *AccessTokenProperty
	OAuth           *OAuthProperty
	OpenPlatform    *OpenPlatformProperty
	OfficialAccount *OfficialAccountProperty
	MiniProgram     *MiniProgramProperty
	Payment         *PaymentProperty
}
