package wego

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/godcong/wego/util"
	"github.com/json-iterator/go"
	"golang.org/x/xerrors"
)

// NilPropertyProperty ...
const NilPropertyProperty = "%T point is null"

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

// Config ...
func (property *SafeCertProperty) Config() (config *tls.Config, e error) {
	cert, e := tls.X509KeyPair(property.Cert, property.Key)
	if e != nil {
		return &tls.Config{}, e
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(property.RootCA)
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            certPool,
		InsecureSkipVerify: true, //client端略过对证书的校验
	}
	tlsConfig.BuildNameToCertificate()
	return tlsConfig, nil
}

// PaymentProperty ...
type PaymentProperty struct {
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
	AppID     string
	AppSecret string
	Token     string
	AesKey    string
	//OAuth     *OAuthProperty
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
	AppID string
	MchID string
	Key   string
}

// Property 属性配置，各个接口用到的参数
type Property struct {
	JSSDK           *JSSDKProperty
	AccessToken     *AccessTokenProperty
	OAuth           *OAuthProperty
	OpenPlatform    *OpenPlatformProperty
	OfficialAccount *OfficialAccountProperty
	MiniProgram     *MiniProgramProperty
	Payment         *PaymentProperty
	SafeCert        *SafeCertProperty
}

// ParseProperty ...
func ParseProperty(config *Config, v ...interface{}) (e error) {
	if config == nil {
		return xerrors.New("nil config")
	}

	for i := range v {
		switch t := v[i].(type) {
		case *Property:
		//TODO:
		case *SafeCertProperty:
			e = parseSafeCertProperty(config, t)
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
	*property = AccessTokenProperty{
		GrantType: GrantTypeClient,
		AppID:     config.AppID,
		AppSecret: config.AppSecret,
	}
	return nil
}

func parseJSSDKProperty(config *Config, property *JSSDKProperty) error {
	//var token AccessTokenProperty
	//e := parseAccessTokenProperty(config, &token)
	//if e != nil {
	//	return e
	//}
	*property = JSSDKProperty{
		AppID: config.AppID,
		MchID: config.MchID,
		Key:   config.MchKey,
	}

	return nil
}

func parsePaymentProperty(config *Config, property *PaymentProperty) error {
	*property = PaymentProperty{
		AppID:     config.AppID,
		AppSecret: config.AppSecret,
		MchID:     config.MchID,
		Key:       config.MchKey,
	}
	return nil
}

func parseSafeCertProperty(config *Config, property *SafeCertProperty) error {
	*property = SafeCertProperty{
		Cert:   config.PemCert,
		Key:    config.PemKEY,
		RootCA: config.RootCA,
	}
	return nil
}

func parseOfficialAccountProperty(config *Config, property *OfficialAccountProperty) (e error) {
	*property = OfficialAccountProperty{
		AppID:     config.AppID,
		AppSecret: config.AppSecret,
		Token:     config.Token,
		AesKey:    config.AesKey,
	}
	return nil
}

func parseOAuthProperty(config *Config, property *OAuthProperty) error {
	*property = OAuthProperty{
		Scopes:      config.Scopes,
		RedirectURI: config.RedirectURI,
	}
	return nil
}
