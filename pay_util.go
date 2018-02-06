package wxpay

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/godcong/wopay/util"
	"github.com/silenceper/wechat/oauth"
)

var (
	ErrorSignType  = errors.New("sign type error")
	ErrorParameter = errors.New("JsonApiParameters() check error")
	ErrorToken     = errors.New("EditAddressParameters() token is nil")
)

//IsSignatureValid check sign
func IsSignatureValid(xml, key string) bool {
	data := util.XmlToMap(xml)

	if !data.IsExist(FIELD_SIGN) {
		return false
	}
	sign1 := data.Get(FIELD_SIGN)
	sign2, _ := GenerateSignature(data, key, SIGN_TYPE_MD5)
	return sign1 == sign2
}

//GenerateSignature make sign from map data
func GenerateSignature(reqData PayData, key string, signType SignType) (string, error) {
	keys := reqData.SortKeys()
	var sign []string

	for _, k := range keys {
		if k == FIELD_SIGN {
			continue
		}
		v := strings.TrimSpace(reqData[k])
		if len(v) > 0 {
			sign = append(sign, strings.Join([]string{k, v}, "="))
		}
	}
	sign = append(sign, strings.Join([]string{"key", key}, "="))
	sb := strings.Join(sign, "&")
	if signType == SIGN_TYPE_MD5 {
		sb = util.MakeSignMD5(sb)
		return sb, nil
	} else if signType == SIGN_TYPE_HMACSHA256 {
		sb = util.MakeSignHMACSHA256(sb, key)
		return sb, nil
	} else {
		return "", ErrorSignType
	}
}

//SandboxSignKey get wechat sandbox sign key
func SandboxSignKey() (string, error) {
	config := PayConfigInstance()
	data := make(PayData)
	data.Set("mch_id", config.MchID())
	data.Set("nonce_str", util.GenerateNonceStr())
	sign, _ := GenerateSignature(data, config.Key(), SIGN_TYPE_MD5)
	data.Set("sign", sign)
	pay := NewPay(config)
	return pay.RequestWithoutCert(SANDBOX_SIGNKEY_URL_SUFFIX, data)
}

func JsonApiParameters(data PayData) (string, error) {
	if !data.IsExist("appid") ||
		!data.IsExist("prepay_id") ||
		data.Get("prepay_id") == "" {
		return "", ErrorParameter
	}

	pay := make(PayData)
	pay.Set("appid", data.Get("appid"))
	pay.Set("timeStamp", util.CurrentTimeStampString())
	pay.Set("nonceStr", util.GenerateNonceStr())
	pay.Set("package", "prepay_id="+data.Get("prepay_id"))
	pay.Set("signType", SIGN_TYPE_MD5.ToString())
	s, e := GenerateSignature(pay, PayConfigInstance().Key(), SIGN_TYPE_MD5)
	if e != nil {
		return "", e
	}
	pay.Set("paySign", s)
	b, err := json.Marshal(pay)
	return string(b), err
}

func EditAddressParameters(url string, token *oauth.ResAccessToken) (string, error) {
	if token == nil {
		return "", ErrorToken
	}
	pay := make(PayData)
	pay.Set("appid", PayConfigInstance().AppID())
	pay.Set("url", url)
	pay.Set("timestamp", util.CurrentTimeStampString())
	pay.Set("noncestr", util.GenerateNonceStr())
	pay.Set("accesstoken", token.AccessToken)
	param := util.ToUrlParams(pay)
	addrSign := util.SHA1(param)
	afterData := PayData{
		"addrSign":  addrSign,
		"signType":  "sha1",
		"scope":     "jsapi_address",
		"appId":     pay.Get("appid"),
		"timeStamp": pay.Get("timestamp"),
		"nonceStr":  pay.Get("noncestr"),
	}
	return afterData.ToJson(), nil
}
