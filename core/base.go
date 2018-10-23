package core

import (
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
)

/*Base 基础 */
type Base struct {
	config *Config
	client *Client
	token  *AccessToken
}

/*URL URL */
type URL struct {
	config *Config
	token  *AccessToken
	client *Client
}

func newBase(config *Config) *Base {
	client := NewClient(config)
	return &Base{
		config: config,
		client: client,
		token:  NewAccessToken(config, client),
	}
}

//NewBase NewBase
func NewBase(config *Config) *Base {
	return newBase(config)
}

/*ClearQuota 公众号的所有api调用（包括第三方帮其调用）次数进行清零
公众号调用或第三方平台帮公众号调用对公众号的所有api调用（包括第三方帮其调用）次数进行清零：
HTTP请求：POST HTTP调用： https://api.weixin.qq.com/cgi-bin/clear_quota?access_token=ACCESS_TOKEN
参数说明：
参数	是否必须	说明
access_token	是	调用接口凭据
appid	是	公众号的APPID
正常情况下，会返回：
{ "errcode" :0, "errmsg" : "ok" }
如果调用超过限制次数，则返回：
{ "errcode" :48006, "errmsg" : "forbid to clear quota because of reaching the limit" }
*/
func (b *Base) ClearQuota() Response {
	token := b.token.GetToken()

	params := util.Map{
		"appid": b.config.Get("app_id"),
	}
	return b.client.PostJSON(Link(clearQuotaURLSuffix), params, util.Map{
		DataTypeQuery: token.KeyMap(),
	})

}

/*GetCallbackIP 请求微信的服务器IP列表
  接口调用请求说明
  http请求方式: GEThttps://api.weixin.qq.com/cgi-bin/getcallbackip?access_token=ACCESS_TOKEN
  参数说明
  参数	是否必须	说明
  access_token	是	公众号的access_token
  返回说明
  正常情况下，微信会返回下述JSON数据包给公众号：
  {    "ip_list": [        "127.0.0.1",         "127.0.0.2",         "101.226.103.0/25"    ]}
  参数	说明
  ip_list	微信服务器IP地址列表
  错误时微信会返回错误码等信息，JSON数据包示例如下（该示例为AppID无效错误）:
  {"errcode":40013,"errmsg":"invalid appid"}
  成功:
  {"ip_list":["101.226.62.77","101.226.62.78","101.226.62.79","101.226.62.80","101.226.62.81","101.226.62.82","101.226.62.83","101.226.62.84","101.226.62.85","101.226.62.86","101.226.103.59","101.226.103.60","101.226.103.61","101.226.103.62","101.226.103.63","101.226.103.69","101.226.103.70","101.226.103.71","101.226.103.72","101.226.103.73","140.207.54.73","140.207.54.74","140.207.54.75","140.207.54.76","140.207.54.77","140.207.54.78","140.207.54.79","140.207.54.80","182.254.11.203","182.254.11.202","182.254.11.201","182.254.11.200","182.254.11.199","182.254.11.198","59.37.97.100","59.37.97.101","59.37.97.102","59.37.97.103","59.37.97.104","59.37.97.105","59.37.97.106","59.37.97.107","59.37.97.108","59.37.97.109","59.37.97.110","59.37.97.111","59.37.97.112","59.37.97.113","59.37.97.114","59.37.97.115","59.37.97.116","59.37.97.117","59.37.97.118","112.90.78.158","112.90.78.159","112.90.78.160","112.90.78.161","112.90.78.162","112.90.78.163","112.90.78.164","112.90.78.165","112.90.78.166","112.90.78.167","140.207.54.19","140.207.54.76","140.207.54.77","140.207.54.78","140.207.54.79","140.207.54.80","180.163.15.149","180.163.15.151","180.163.15.152","180.163.15.153","180.163.15.154","180.163.15.155","180.163.15.156","180.163.15.157","180.163.15.158","180.163.15.159","180.163.15.160","180.163.15.161","180.163.15.162","180.163.15.163","180.163.15.164","180.163.15.165","180.163.15.166","180.163.15.167","180.163.15.168","180.163.15.169","180.163.15.170","101.226.103.0\/25","101.226.233.128\/25","58.247.206.128\/25","182.254.86.128\/25","103.7.30.21","103.7.30.64\/26","58.251.80.32\/27","183.3.234.32\/27","121.51.130.64\/27"]}
  失败:
  {"errcode":40013,"errmsg":"invalid appid"}
*/
func (b *Base) GetCallbackIP() Response {
	token := b.token.GetToken()
	b.client.SetRequestType(DataTypeJSON)
	return b.client.Get(APIWeixin+getCallbackIPURLSuffix, token.KeyMap())
}

/*ShortURL 转换短链接
https://apihk.mch.weixin.qq.com/tools/shorturl    （建议接入点：东南亚）
https://apius.mch.weixin.qq.com/tools/shorturl    （建议接入点：其它）
https://api.mch.weixin.qq.com/tools/shorturl        （建议接入点：中国国内）
请求参数
URL链接	long_url	是	String(512、	weixin：//wxpay/bizpayurl?sign=XXXXX&appid=XXXXX&mch_id=XXXXX&product_id=XXXXXX&time_stamp=XXXXXX&nonce_str=XXXXX	需要转换的URL，签名用原串，传输需URLencode
返回结果
返回状态码	return_code	是	String(16)	SUCCESS/FAIL
URL链接	short_url	是	String(64)	weixin：//wxpay/s/XXXXXX	转换后的URL
*/
func (u *URL) ShortURL(url string) Response {
	//TODO:need fix
	//token := u.token.GetToken()
	//ops := util.Map{
	//	DataTypeQuery: token.KeyMap(),
	//}
	m := util.Map{
		"action":   "long2short",
		"long_url": url,
	}
	resp := u.client.PostJSON(APIWeixin+shortURLSuffix, m, nil)
	log.Debug("URL|ShortURL", resp)
	return resp
}

/*NewURL NewURL*/
func NewURL(config *Config) *URL {
	client := NewClient(config)
	return &URL{
		config: config,
		token:  NewAccessToken(config, client),
		client: client,
	}
}
