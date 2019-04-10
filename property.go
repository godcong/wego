package wego

import (
	"github.com/godcong/wego/util"
	"github.com/json-iterator/go"
	"golang.org/x/xerrors"
)

// NilPropertyProperty ...
const NilPropertyProperty = "%T point is null"

// SandboxOption ...
type SandboxOption struct {
	SubMchID string `xml:"sub_mch_id"`
	SubAppID string `xml:"sub_app_id"`
}

// SandboxProperty ...
type SandboxProperty struct {
	AppID     string
	AppSecret string
	MchID     string
	Key       string
}

// SafeCertProperty ...
type SafeCertProperty struct {
	Cert   []byte
	Key    []byte
	RootCA []byte
}

// PaymentProperty ...
type PaymentProperty struct {
	BodyType  BodyType          `xml:"body_type"`
	AppID     string            `xml:"app_id"`
	AppSecret string            `xml:"app_secret"`
	MchID     string            `xml:"mch_id"`
	Key       string            `xml:"key"`
	SafeCert  *SafeCertProperty `xml:"safe_cert"`
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
	AppID       string
	AppSecret   string
	Token       string
	AesKey      string
	AccessToken *AccessTokenProperty
	OAuth       *OAuthProperty
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
		"secret":     obj.AppSecret,
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

// JSSDKProperty ...
type JSSDKProperty struct {
	AppID       string
	MchID       string
	Key         string
	AccessToken *AccessTokenProperty
}

// JSSDKOption ...
type JSSDKOption struct {
	SubAppID string
	URL      string
}

// Property 属性配置，各个接口用到的参数
type Property struct {
	JSSDK                 *JSSDKProperty
	JSSDKOption           *JSSDKOption
	AccessToken           *AccessTokenProperty
	AccessTokenOption     *AccessTokenOption
	OAuth                 *OAuthProperty
	OpenPlatform          *OpenPlatformProperty
	OfficialAccount       *OfficialAccountProperty
	OfficialAccountOption *OfficialAccountOption
	MiniProgram           *MiniProgramProperty
	Payment               *PaymentProperty
	PaymentOption         *PaymentOption
	SafeCert              *SafeCertProperty
}

// ParseProperty ...
func ParseProperty(config *Config, v ...interface{}) (e error) {
	if config == nil {
		return xerrors.New("nil config")
	}

	for i := range v {
		switch t := v[i].(type) {
		case *JSSDKProperty:
			e = parseJSSDKProperty(config, t)
		case *AccessTokenProperty:
			e = parseAccessTokenProperty(config, t)
		case *PaymentProperty:
			e = parsePaymentProperty(config, t)
		case *OAuthProperty:
			e = parseOAuthProperty(config, t)
		case *OfficialAccountProperty:
			e = parseOfficialAccountProperty(config, t)
		default:
			e = xerrors.Errorf("wrong property point")
		}
		if e != nil {
			return e
		}
	}
	return nil
}

func parseAccessTokenProperty(config *Config, property *AccessTokenProperty) error {
	property = &AccessTokenProperty{
		GrantType: GrantTypeClient,
		AppID:     config.AppID,
		AppSecret: config.AppSecret,
	}
	return nil
}

func parseJSSDKProperty(config *Config, property *JSSDKProperty) error {
	var token AccessTokenProperty
	e := parseAccessTokenProperty(config, &token)
	if e != nil {
		return e
	}
	property = &JSSDKProperty{
		AppID:       config.AppID,
		MchID:       config.MchID,
		Key:         config.MchKey,
		AccessToken: &token,
	}

	return nil
}

func parsePaymentProperty(config *Config, property *PaymentProperty) error {
	var cert SafeCertProperty
	e := parseSafeCertProperty(config, &cert)
	if e != nil {
		return e
	}
	property = &PaymentProperty{
		AppID:     config.AppID,
		AppSecret: config.AppSecret,
		MchID:     config.MchID,
		Key:       config.MchKey,
		SafeCert:  &cert,
	}
	return nil
}

func parseSafeCertProperty(config *Config, property *SafeCertProperty) error {
	property = &SafeCertProperty{
		Cert:   config.PemCert,
		Key:    config.PemKEY,
		RootCA: config.RootCA,
	}
	return nil
}

func parseOfficialAccountProperty(config *Config, property *OfficialAccountProperty) (e error) {
	var token AccessTokenProperty
	e = parseAccessTokenProperty(config, &token)
	if e != nil {
		return e
	}

	var oauth OAuthProperty
	e = parseOAuthProperty(config, &oauth)
	if e != nil {
		return e
	}

	property = &OfficialAccountProperty{
		AppID:       config.AppID,
		AppSecret:   config.AppSecret,
		Token:       config.Token,
		AesKey:      config.AesKey,
		AccessToken: &token,
		OAuth:       &oauth,
	}
	return nil
}

func parseOAuthProperty(config *Config, property *OAuthProperty) error {
	property = &OAuthProperty{
		Scopes:      config.Scopes,
		RedirectURI: config.RedirectURI,
	}
	return nil
}
