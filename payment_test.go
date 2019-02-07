package wego

import (
	"github.com/go-redis/redis"
	"github.com/godcong/wego/cache"
	"github.com/godcong/wego/util"
	"golang.org/x/text/encoding/simplifiedchinese"
	"testing"
)

var p = Property{
	Local: nil,
	AccessToken: &AccessTokenProperty{
		GrantType: GrantTypeClient,
		AppID:     "wx15810ca16324c225",
		Secret:    "4f2601f6567ac4b5741aa2dc2e5eadfd",
	},
	OAuth:           nil,
	OpenPlatform:    nil,
	OfficialAccount: nil,
	MiniProgram:     nil,
	Payment: &PaymentProperty{
		AppID:      "wx15810ca16324c225",
		MchID:      "1516796851",
		Key:        "aTKnSUcTkbEnhwQNdutWkQxAjnhAz2jK",
		CertPEM:    "",
		KeyPEM:     "",
		RootCaPEM:  "",
		PublicKey:  "",
		PrivateKey: "",
	},
	SafeCert: &SafeCertProperty{
		Cert: []byte(`-----BEGIN CERTIFICATE-----
MIIEajCCA9OgAwIBAgIEAkekmzANBgkqhkiG9w0BAQUFADCBijELMAkGA1UEBhMC
Q04xEjAQBgNVBAgTCUd1YW5nZG9uZzERMA8GA1UEBxMIU2hlbnpoZW4xEDAOBgNV
BAoTB1RlbmNlbnQxDDAKBgNVBAsTA1dYRzETMBEGA1UEAxMKTW1wYXltY2hDQTEf
MB0GCSqGSIb3DQEJARYQbW1wYXltY2hAdGVuY2VudDAeFw0xODEwMTcwMjM5NTRa
Fw0yODEwMTQwMjM5NTRaMIGZMQswCQYDVQQGEwJDTjESMBAGA1UECBMJR3Vhbmdk
b25nMREwDwYDVQQHEwhTaGVuemhlbjEQMA4GA1UEChMHVGVuY2VudDEOMAwGA1UE
CxMFTU1QYXkxLTArBgNVBAMUJOS4iua1t+WHr+Wlh+S/oeaBr+enkeaKgOaciemZ
kOWFrOWPuDESMBAGA1UEBBMJMTAyOTQyNjE2MIIBIjANBgkqhkiG9w0BAQEFAAOC
AQ8AMIIBCgKCAQEAvzZLiTs/Vg4dDhlsDmo8A38xEY2wkw5OEPABL905Nm3XvN4P
rJO2xYhEWNwPfbVHjO3t7dH5EEidFSRVPqSNbpDtolxteuCQvVJc3yT4TUY1MD0g
H1O4wCYHec7PzBi8Tl/mK//xewD5EEy981Nu53ozypm+Y7It1Ofrgm2U0YpcFwTP
pMS9gdz3a/oHJig6ZzhVTxZi7lbmMNcDNhGX/TJnzG8C+jt3EadeDEFqXxwcAlVJ
Ga6bb4mPtHJ34JjfrBAf7/BDnXb8Ifue3PN9L5GzztTk6dLRH8R0boCKOAMkkULc
3WgpQlP0fBuaPPyYOX56IIXrOOF3ad69W/z2kQIDAQABo4IBRjCCAUIwCQYDVR0T
BAIwADAsBglghkgBhvhCAQ0EHxYdIkNFUy1DQSBHZW5lcmF0ZSBDZXJ0aWZpY2F0
ZSIwHQYDVR0OBBYEFB4EPUbsvIXvl6JLcikaDPuIk/iXMIG/BgNVHSMEgbcwgbSA
FD4FJvYiYrQVW4jNZH6w1GKn5YZ0oYGQpIGNMIGKMQswCQYDVQQGEwJDTjESMBAG
A1UECBMJR3Vhbmdkb25nMREwDwYDVQQHEwhTaGVuemhlbjEQMA4GA1UEChMHVGVu
Y2VudDEMMAoGA1UECxMDV1hHMRMwEQYDVQQDEwpNbXBheW1jaENBMR8wHQYJKoZI
hvcNAQkBFhBtbXBheW1jaEB0ZW5jZW50ggkAu1SXK7wA6FcwDgYDVR0PAQH/BAQD
AgbAMBYGA1UdJQEB/wQMMAoGCCsGAQUFBwMCMA0GCSqGSIb3DQEBBQUAA4GBAGtM
TABzTFHfUQVUbW8J34z8DdLd/aIzBNGF8vU4B5jY+13ki/aWboTYjwQzFv2CC+gz
kpeun3dr2Rd3bqHM7d1xc0EfdWBwsXIuQML2tzFK0clq6NuhkeMPnVui8pQ5r+V0
c0ogMNcBHxUSwVr8AoiDhOnkvkOEUPAa8XaQ9mcM
-----END CERTIFICATE-----`),
		Key: []byte(`-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC/NkuJOz9WDh0O
GWwOajwDfzERjbCTDk4Q8AEv3Tk2bde83g+sk7bFiERY3A99tUeM7e3t0fkQSJ0V
JFU+pI1ukO2iXG164JC9UlzfJPhNRjUwPSAfU7jAJgd5zs/MGLxOX+Yr//F7APkQ
TL3zU27nejPKmb5jsi3U5+uCbZTRilwXBM+kxL2B3Pdr+gcmKDpnOFVPFmLuVuYw
1wM2EZf9MmfMbwL6O3cRp14MQWpfHBwCVUkZrptviY+0cnfgmN+sEB/v8EOddvwh
+57c830vkbPO1OTp0tEfxHRugIo4AySRQtzdaClCU/R8G5o8/Jg5fnoghes44Xdp
3r1b/PaRAgMBAAECggEAKFtFRgmDLJ5982PaMpCpafOGX7YP9nmmcFy3Y2rFkH9O
cUZM+JeUk8xW4IpRmq6QE2aEORWAA7fPr46jjz0bvKJKBWKMmMqcwdiIUXB0F5sh
zrSX+wSokaV1HnhP0zvFQqVq5w514QgriQ0pahnJ5oEcPT6FuHT6x1zLkBOgvTih
Inu60MRircuARUgMk5vfuwk08kOHEqOGtsGYzpyWGXY0a8VEn4ad2OXBer7oFdHL
JUWsjWMnUAzXIqiTVPW4pHaik7UxwILrRVLGpUl+fxVZfj5cmG+u6Uj2s+MbufCm
JFFb/VV5uj4JxRxXMIR64zli4b9e4yxjEAqZUWSXIQKBgQDd5PCH2IBvF35O29BO
ZPY9RjvOlaoBVBU5T2w1GR3YngE/PYZup4u6V4Z7Gl8YjjKixJusPyY+WIdajB/e
8Sk2t9VymKGneiNVLgYCisVqbMNQj1KsQowwC39pKHjRwQTagKLueDZBmD6G2b2b
vD4JAeKWFUs1/32vF2na1zL25wKBgQDcmhO8+uvZZ+Lp/gT3wvaztcWYbXVjacrB
kSlenkJIEdLhv3EV795s7i4FI1RsYTz2hO2MgpbJaYz1DkTdYsHhB32efc67ICGJ
e+q6hkudogNiJSf/q+9I0OdSGw77l3xv4IN/X2S1Lijl82Rt526K7ct0Mt5OICiX
FVmdvmyPxwKBgB3sb9uCAN/6ZRgxYIgR6Xsd65bSbBi9xBb0dP08I+BAfp6Op4Bo
q9k3UHVtewGOu5ljtjKmWXvy6F33o5ynHQ26ANjxaGtzdyxXhov1NuZBQJ2ullGQ
r+MSyd3ejVyMESF+U3qhT3/NDjGVT5Ke8ia40Fu5B2mCyOpB2A6hEshlAoGBAJHQ
ilF07iZDI9sVG8WFKjp5YQWy/E/c4rd1swmmxBsNJP4w45fjeHs6KFMfGjOVRo2Q
KdChlPfG+/QTHXPUTmDm6aK0d8x6nZBZxzrJf/vO1juuYbT+uejApNZBqnKknAz7
MxAjRtb3jvMIIY7/1bKhIcVJxmWPniTXaOf6sZmLAoGANh9hnndt9hFO/+riXzTW
aN0KNLzrmJvrBjJ7nlndbLpw1FD6GNnCEkXQwE9iEZ6QcZ8F2ORarKaorxYPXt99
6Vn+7YZAn/3KW/KRlpO7z+LuJjEqD4GcnhzdVXT7FjL6JFAnQasGoiut/9xFw1ar
+JZMm8XUy5L9RSw/xVOsErk=
-----END PRIVATE KEY-----`),
		RootCA: []byte(`-----BEGIN CERTIFICATE-----
MIIDIDCCAomgAwIBAgIENd70zzANBgkqhkiG9w0BAQUFADBOMQswCQYDVQQGEwJV
UzEQMA4GA1UEChMHRXF1aWZheDEtMCsGA1UECxMkRXF1aWZheCBTZWN1cmUgQ2Vy
dGlmaWNhdGUgQXV0aG9yaXR5MB4XDTk4MDgyMjE2NDE1MVoXDTE4MDgyMjE2NDE1
MVowTjELMAkGA1UEBhMCVVMxEDAOBgNVBAoTB0VxdWlmYXgxLTArBgNVBAsTJEVx
dWlmYXggU2VjdXJlIENlcnRpZmljYXRlIEF1dGhvcml0eTCBnzANBgkqhkiG9w0B
AQEFAAOBjQAwgYkCgYEAwV2xWGcIYu6gmi0fCG2RFGiYCh7+2gRvE4RiIcPRfM6f
BeC4AfBONOziipUEZKzxa1NfBbPLZ4C/QgKO/t0BCezhABRP/PvwDN1Dulsr4R+A
cJkVV5MW8Q+XarfCaCMczE1ZMKxRHjuvK9buY0V7xdlfUNLjUA86iOe/FP3gx7kC
AwEAAaOCAQkwggEFMHAGA1UdHwRpMGcwZaBjoGGkXzBdMQswCQYDVQQGEwJVUzEQ
MA4GA1UEChMHRXF1aWZheDEtMCsGA1UECxMkRXF1aWZheCBTZWN1cmUgQ2VydGlm
aWNhdGUgQXV0aG9yaXR5MQ0wCwYDVQQDEwRDUkwxMBoGA1UdEAQTMBGBDzIwMTgw
ODIyMTY0MTUxWjALBgNVHQ8EBAMCAQYwHwYDVR0jBBgwFoAUSOZo+SvSspXXR9gj
IBBPM5iQn9QwHQYDVR0OBBYEFEjmaPkr0rKV10fYIyAQTzOYkJ/UMAwGA1UdEwQF
MAMBAf8wGgYJKoZIhvZ9B0EABA0wCxsFVjMuMGMDAgbAMA0GCSqGSIb3DQEBBQUA
A4GBAFjOKer89961zgK5F7WF0bnj4JXMJTENAKaSbn+2kmOeUJXRmm/kEd5jhW6Y
7qj/WsjTVbJmcVfewCHrPSqnI0kBBIZCe/zuf6IWUrVnZ9NA2zsmWLIodz2uFHdh
1voqZiegDfqnc1zqcPGUIWVEX/r87yloqaKHee9570+sB3c4
-----END CERTIFICATE-----`),
	},
}

func init() {
	redis := cache.NewRedisCache(&redis.Options{
		Addr:     "localhost:6379",
		Password: "2rXfzaNKqX1b", // no password set
		DB:       1,              // use default DB
	})

	cache.RegisterCache(redis)
}

// TestPayment_SandboxSignKey ...
func TestPayment_SandboxSignKey(t *testing.T) {
	payment := NewPayment(&p, &PaymentOption{
		Sandbox: &SandboxProperty{},
	})
	key := payment.Sandbox().SignKey().ToMap()
	t.Log(key)
	if !key.Has("return_code") || !key.Has("return_msg") {
		t.Error(key)
	}
}

// TestAccessToken_GetToken ...
func TestAccessToken_GetToken(t *testing.T) {
	token := NewAccessToken(p.AccessToken)
	tk := token.GetToken()
	t.Log(tk)
}

// UnifyResult ...
type UnifyResult struct {
	Appid      string `xml:"appid"`
	CodeURL    string `xml:"code_url"`
	DeviceInfo string `xml:"device_info"`
	ErrCode    string `xml:"err_code"`
	ErrCodeDes string `xml:"err_code_des"`
	MchID      string `xml:"mch_id"`
	NonceStr   string `xml:"nonce_str"`
	PrepayID   string `xml:"prepay_id"`
	ResultCode string `xml:"result_code"`
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	Sign       string `xml:"sign"`
	TradeType  string `xml:"trade_type"`
}

// TestOrder_Unify ...
func TestOrder_Unify(t *testing.T) {
	m := make(util.Map)
	m.Set("body", "腾讯充值中心-QQ会员充值")
	m.Set("out_trade_no", "otababababq1234"+"6")
	//m.Set("device_info", "")
	////m.Set("fee_type", "CNY")
	m.Set("total_fee", "551")
	////m.Set("spbill_create_ip", "123.12.12.123")
	//m.Set("notify_url", "https://test.letiantian.me/wxpay/notify")
	m.Set("trade_type", "NATIVE")
	//m.Set("openid", "oLyBi0hSYhggnD-kOIms0IzZFqrc")
	//m.Set("openid", "oE_gl0Yr54fUjBhU5nBlP4hS2efo")

	////m.Set("product_id", "12")
	newPayment := NewPayment(&p, &PaymentOption{
		RemoteAddress: "",
		LocalHost:     "",
		UseSandbox:    true,
		Sandbox: &SandboxProperty{
			AppID:  "wx15810ca16324c225",
			Secret: "4f2601f6567ac4b5741aa2dc2e5eadfd",
			MchID:  "1516796851",
			Key:    "aTKnSUcTkbEnhwQNdutWkQxAjnhAz2jK",
		},
		NotifyURL: "",
		RefundURL: "",
	}).Unify(m)
	if newPayment.Error() != nil {
		t.Log(newPayment)
	}

	t.Log(string(newPayment.Bytes()))
	t.Log(newPayment.ToMap())
	t.Log(newPayment.Error())
	var ur UnifyResult
	e := newPayment.Unmarshal(&ur)
	t.Log(ur, e)
	//order := payment.NewOrder()
	//resp := order.Unify(m)
	//log.Println(resp.ToMap())
	//{"appid":"wx426b3015555a46be","code_url":"weixin://wxpay/bizpayurl?pr=D3sNT8y","mch_id":"1900009851","nonce_str":"FRFByNNdrzRuEGkp","prepay_id":"wx20180220113507842dff20340601189342","result_code":"SUCCESS","return_code":"SUCCESS","return_msg":"OK","sign":"D398DA0644A14D0BC00A8B82D8D4ECDC","trade_type":"NATIVE"}
}

// TestPayment_Pay ...
func TestPayment_Pay(t *testing.T) {
	newPayment := NewPayment(&p, &PaymentOption{
		RemoteAddress: "",
		LocalHost:     "",
		UseSandbox:    true,
		Sandbox: &SandboxProperty{
			AppID:  "wx15810ca16324c225",
			Secret: "4f2601f6567ac4b5741aa2dc2e5eadfd",
			MchID:  "1516796851",
			Key:    "aTKnSUcTkbEnhwQNdutWkQxAjnhAz2jK",
		},
		NotifyURL: "",
		RefundURL: "",
	})
	resp := newPayment.Pay(util.Map{
		"body":         "image形象店-深圳腾大- QQ公仔",
		"out_trade_no": "1217752501201407033233368018",
		"total_fee":    "1",
		"auth_code":    "120061098828009406",
	})
	t.Log(resp.ToMap())
	t.Log(resp.Error())
}

// TestPayment_DownloadFundFlow ...
func TestPayment_DownloadFundFlow(t *testing.T) {
	newPayment := NewPayment(&p, &PaymentOption{
		RemoteAddress: "",
		LocalHost:     "",
		UseSandbox:    true,
		Sandbox: &SandboxProperty{
			AppID:  "wx15810ca16324c225",
			Secret: "4f2601f6567ac4b5741aa2dc2e5eadfd",
			MchID:  "1516796851",
			Key:    "aTKnSUcTkbEnhwQNdutWkQxAjnhAz2jK",
		},
		NotifyURL: "",
		RefundURL: "",
	})
	resp := newPayment.DownloadFundFlow("20181109", "Basic")
	_ = SaveEncodingTo(resp, "d:/test.csv", simplifiedchinese.GBK.NewEncoder())
	t.Log(resp.Error())
	t.Log(resp.ToMap())
}

// TestPayment_ByOutTradeNumber ...
func TestPayment_ByOutTradeNumber(t *testing.T) {
	newPayment := NewPayment(&p, &PaymentOption{
		RemoteAddress: "",
		LocalHost:     "",
		UseSandbox:    true,
		Sandbox: &SandboxProperty{
			AppID:  "wx15810ca16324c225",
			Secret: "4f2601f6567ac4b5741aa2dc2e5eadfd",
			MchID:  "1516796851",
			Key:    "aTKnSUcTkbEnhwQNdutWkQxAjnhAz2jK",
		},
		NotifyURL: "",
		RefundURL: "",
	})
	resp := newPayment.ReverseByOutTradeNumber("1217752501201407033233368018")
	_ = SaveEncodingTo(resp, "d:/test.csv", simplifiedchinese.GBK.NewEncoder())
	t.Log(resp.Error())
	t.Log(resp.ToMap())
}
