package core

import (
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
)

/*URL URL */
type URL struct {
	config *Config
	token  *AccessToken
	client *Client
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
	token := u.token.GetToken()
	m := util.Map{
		"action":   "long2short",
		"long_url": url,
	}
	resp := u.client.PostJSON(APIWeixin+shortURLSuffix, token.KeyMap(), m)
	log.Debug("URL|ShortURL", resp)
	return resp
}

/*NewURL NewURL*/
func NewURL(config *Config) *URL {
	client := NewClient(config)
	return &URL{
		config: config,
		token:  NewAccessToken(config),
		client: client,
	}
}
