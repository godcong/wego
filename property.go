package wego

import (
	"crypto/md5"
	"fmt"
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

func (obj *SandboxProperty) getCacheKey() string {
	name := obj.AppID + "." + obj.MchID
	return "godcong.wego.payment.sandbox." + fmt.Sprintf("%x", md5.Sum([]byte(name)))
}

// SignKey ...
func (obj *SandboxProperty) SignKey() Responder {
	m := make(util.Map)
	m.Set("mch_id", obj.MchID)
	m.Set("nonce_str", util.GenerateNonceStr())
	m.Set("sign", util.GenSign(m, obj.Key))
	resp := PostXML(util.URL(apiMCHWeixin, sandboxNew, getSignKey), nil, m)
	return resp
}

// SafeCertProperty ...
type SafeCertProperty struct {
	Cert   []byte
	Key    []byte
	RootCA []byte
}

// PaymentProperty ...
type PaymentProperty struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	MchID     string `json:"mch_id"`
	//CertPEM    string `json:"cert_pem"`
	//KeyPEM     string `json:"key_pem"`
	//RootCaPEM  string `json:"root_ca_pem"`
	//PublicKey  string `json:"public_key"`
	//PrivateKey string `json:"private_key"`
	SubMchID string `json:"sub_mch_id"`
	SubAppID string `json:"sub_app_id"`
}

// OAuthProperty ...
type OAuthProperty struct {
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

// AccessTokenProperty ...
type AccessTokenProperty struct {
	GrantType string `toml:"grant_type"`
	AppID     string `toml:"app_id"`
	AppSecret string `toml:"app_secret"`
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
	SafeCert        *SafeCertProperty
}
