package wego

import (
	"crypto/md5"
	"fmt"
	"github.com/godcong/wego/cache"
	"github.com/godcong/wego/util"
	log "github.com/sirupsen/logrus"
	"strings"
)

/*JSSDK JSSDK */
type JSSDK struct {
	*JSSDKProperty
	accessToken *AccessToken
	ticket      *Ticket
	subAppID    string
	url         string
	//CacheKey    func() string
}

/*NewJSSDK NewJSSDK */
func NewJSSDK(property *JSSDKProperty, options ...JSSDKOption) *JSSDK {
	jssdk := &JSSDK{
		JSSDKProperty: property,
		//accessToken:   NewAccessToken(property.AccessToken),
	}
	jssdk.parse(options...)
	return jssdk
}

func (obj *JSSDK) getURL() string {
	if obj.url != "" {
		return obj.url
	}
	return "http://localhost"
}

// BridgeConfig ...
type BridgeConfig struct {
	AppID     string `json:"appId"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	PaySign   string `json:"paySign"`
	SignType  string `json:"signType"`
	TimeStamp string `json:"timeStamp"`
}

/*BridgeConfig bridge 设置 */
func (obj *JSSDK) BridgeConfig(pid string) *BridgeConfig {
	config := &BridgeConfig{
		AppID:     obj.getID(),
		NonceStr:  util.GenerateNonceStr(),
		Package:   strings.Join([]string{"prepay_id", pid}, "="),
		SignType:  "MD5",
		TimeStamp: util.Time(),
	}
	config.PaySign = util.GenSign(util.Map{
		"appId":     config.AppID,
		"timeStamp": config.TimeStamp,
		"nonceStr":  config.NonceStr,
		"package":   config.Package,
		"signType":  config.SignType,
	}, obj.Key)
	return config
}

// SDKConfig ...
type SDKConfig struct {
	AppID     string `json:"appId"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	PaySign   string `json:"paySign"`
	SignType  string `json:"signType"`
	TimeStamp string `json:"timestamp"`
}

/*SDKConfig sdk 设置 */
func (obj *JSSDK) SDKConfig(pid string) *SDKConfig {
	config := obj.BridgeConfig(pid)
	return &SDKConfig{
		AppID:     config.AppID,
		NonceStr:  config.NonceStr,
		Package:   config.PaySign,
		PaySign:   config.PaySign,
		SignType:  config.SignType,
		TimeStamp: config.TimeStamp,
	}
}

// AppConfig ...
type AppConfig struct {
	AppID     string `json:"appid"`
	NonceStr  string `json:"noncestr"`
	Package   string `json:"package"`
	PartnerID string `json:"partnerid"`
	PrepayID  string `json:"prepayid"`
	Sign      string `json:"sign"`
	TimeStamp string `json:"timestamp"`
}

/*AppConfig app 设置 */
func (obj *JSSDK) AppConfig(pid string) *AppConfig {
	config := &AppConfig{
		AppID:     obj.getID(),
		NonceStr:  util.GenerateNonceStr(),
		Package:   "Sign=WXPay",
		PartnerID: obj.MchID,
		PrepayID:  pid,
		TimeStamp: util.Time(),
	}

	config.Sign = util.GenSign(util.Map{
		"appid":     config.AppID,
		"partnerid": config.PartnerID,
		"prepayid":  config.PrepayID,
		"noncestr":  config.NonceStr,
		"timestamp": config.TimeStamp,
		"package":   config.Package,
	}, obj.Key)
	return config
}

// ShareAddressConfig ...
type ShareAddressConfig struct {
	AddrSign  string `json:"addrSign"`
	AppID     string `json:"appId"`
	NonceStr  string `json:"nonceStr"`
	Scope     string `json:"scope"`
	SignType  string `json:"signType"`
	TimeStamp string `json:"timeStamp"`
}

// ShareAddressConfig ...
//	参数:token
//	类型:string或*core.AccessToken
func (obj *JSSDK) ShareAddressConfig(token interface{}) *ShareAddressConfig {
	config := &ShareAddressConfig{
		AppID:     obj.getID(),
		NonceStr:  util.GenerateNonceStr(),
		Scope:     "jsapi_address",
		SignType:  "SHA1",
		TimeStamp: util.Time(),
	}

	if token == nil {
		token = obj.accessToken.GetToken()
	}
	signMsg := util.Map{
		"appid":       config.AppID,
		"url":         obj.getURL(),
		"timestamp":   config.TimeStamp,
		"noncestr":    config.NonceStr,
		"accesstoken": parseAccessToken(token),
	}

	config.AddrSign = util.GenSHA1(signMsg.URLEncode())

	return config
}

// BuildConfig ...
type BuildConfig struct {
	AppID     string   `json:"appID"`
	JSAPIList []string `json:"jsApiList"`
	NonceStr  string   `json:"nonceStr"`
	Signature string   `json:"signature"`
	Timestamp string   `json:"timestamp"`
	URL       string   `json:"url"`
}

// BuildConfig ...
func (obj *JSSDK) BuildConfig(url string, list ...string) *BuildConfig {
	ticket := obj.GetTicket("jsapi", false)
	if ticket == "" {
		return nil
	}

	config := &BuildConfig{
		AppID:     obj.getID(),
		NonceStr:  util.GenerateNonceStr(),
		Timestamp: util.Time(),
		URL:       util.MustString(url, obj.getURL()),
		JSAPIList: list,
	}
	config.Signature = getTicketSignature(ticket, config.NonceStr, config.Timestamp, config.URL)
	return config
}

// GetTicket ...
func (obj *JSSDK) GetTicket(s string, refresh bool) string {
	key := obj.getCacheKey()
	log.Info("key:", key)
	if !refresh && cache.Has(key) {
		if v, b := cache.Get(key).(string); b {
			log.Infof("cached ticket:%+v", v)
			return v
		}
	}

	tr, e := NewTicket(obj.accessToken).GetTicketRes(s)
	if e != nil {
		log.Error(e)
		return ""
	}
	log.Infof("ticket:%+v", *tr)
	cache.SetWithTTL(obj.getCacheKey(), tr.Ticket, tr.ExpiresIn-500)
	return tr.Ticket

}

// getID ...
func (obj *JSSDK) getID() string {
	if obj.subAppID != "" {
		return obj.subAppID
	}
	return obj.AppID
}

func (obj *JSSDK) getCacheKey() string {
	c := md5.Sum([]byte("jssdk." + obj.getID()))
	return fmt.Sprintf("godcong.wego.jssdk.ticket.%x", c[:])
}

func (obj *JSSDK) parse(options ...JSSDKOption) {
	if options == nil {
		return
	}
	for _, o := range options {
		o(obj)
	}
}

func getTicketSignature(ticket, nonce, ts, url string) string {
	return util.GenSHA1(fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s", ticket, nonce, ts, url))
}
