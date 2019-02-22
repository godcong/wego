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
	*JSSDKConfig
	ID string

	accessToken *AccessToken
	ticket      *Ticket
	subAppID    string
	url         string
	//CacheKey    func() string
}

/*NewJSSDK NewJSSDK */
func NewJSSDK(config *JSSDKConfig, options ...JSSDKConfigOption) *JSSDK {
	jssdk := &JSSDK{
		JSSDKConfig: config,
		accessToken: NewAccessToken(config.AccessToken),
	}
	jssdk.parse(options)
	return jssdk
}

func (obj *JSSDK) getURL() string {
	if obj.url != "" {
		return obj.url
	}
	return util.GetServerIP()
}

/*BridgeConfig bridge 设置 */
func (obj *JSSDK) BridgeConfig(pid string) util.Map {
	m := util.Map{
		"appId":     obj.getID(),
		"timeStamp": util.Time(),
		"nonceStr":  util.GenerateNonceStr(),
		"package":   strings.Join([]string{"prepay_id", pid}, "="),
		"signType":  "MD5",
	}

	m.Set("paySign", util.GenSign(m, obj.Key))
	return m
}

/*SDKConfig sdk 设置 */
func (obj *JSSDK) SDKConfig(pid string) util.Map {
	config := obj.BridgeConfig(pid)
	config.Set("timestamp", config.Get("timeStamp"))
	config.Delete("timeStamp")

	return config
}

/*AppConfig app 设置 */
func (obj *JSSDK) AppConfig(pid string) util.Map {
	m := util.Map{
		"appid":     obj.ID,
		"partnerid": obj.MchID,
		"prepayid":  pid,
		"noncestr":  util.GenerateNonceStr(),
		"timestamp": util.Time(),
		"package":   "Sign=WXPay",
	}

	m.Set("sign", util.GenSign(m, obj.Key))
	return m
}

// ShareAddressConfig ...
//	参数:token
//	类型:string或*core.AccessToken
func (obj *JSSDK) ShareAddressConfig(token interface{}) util.Map {

	m := util.Map{
		"appId":     obj.ID,
		"scope":     "jsapi_address",
		"timeStamp": util.Time(),
		"nonceStr":  util.GenerateNonceStr(),
		"signType":  "SHA1",
	}

	signMsg := util.Map{
		"appid":       m.Get("appId"),
		"url":         obj.getURL(),
		"timestamp":   m.Get("timeStamp"),
		"noncestr":    m.Get("nonceStr"),
		"accesstoken": parseAccessToken(token),
	}

	m.Set("addrSign", util.GenSHA1(signMsg.URLEncode()))

	return m
}

// BuildConfig ...
func (obj *JSSDK) BuildConfig(p util.Map) util.Map {
	ticket := obj.GetTicket("jsapi", false)
	nonce := util.GenerateNonceStr()
	ts := util.Time()
	url := p.GetStringD("url", obj.getURL())
	m := util.Map{
		"appID":     obj.ID,
		"nonceStr":  nonce,
		"timestamp": ts,
		"url":       url,
		"jsApiList": p.Get("jsApiList"),
		"signature": getTicketSignature(ticket, nonce, ts, url),
	}
	return m
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

func (obj *JSSDK) parse(options []JSSDKConfigOption) {
	if options == nil {
		return
	}
	obj.subAppID = options[0].SubAppID
	obj.ID = obj.getID()
	obj.url = options[0].URL
}

func getTicketSignature(ticket, nonce, ts, url string) string {
	return util.GenSHA1(fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s", ticket, nonce, ts, url))
}
