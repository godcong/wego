package wego

import (
	"crypto/md5"
	"fmt"
	"github.com/godcong/wego/util"
	"github.com/json-iterator/go"
	"github.com/pelletier/go-toml"
	"log"
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

// PaymentConfig ...
type PaymentConfig struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	MchID     string `json:"mch_id"`
	Key       string `json:"key"`
	//CertPEM    string `json:"cert_pem"`
	//KeyPEM     string `json:"key_pem"`
	//RootCaPEM  string `json:"root_ca_pem"`
	//PublicKey  string `json:"public_key"`
	//PrivateKey string `json:"private_key"`
	SubMchID    string `json:"sub_mch_id"`
	SubAppID    string `json:"sub_app_id"`
	NotifyURL   string `json:"notify_url"`
	RefundedURL string `json:"refunded_url"`
	ScannedURL  string `json:"scanned_url"`
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
	SafeCert        *SafeCertProperty
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
