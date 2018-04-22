package crypt

import "testing"

var encodingAesKey = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFG"
var token = "pamtest"
var timeStamp = "1409304348"
var nonce = "xxxxxx"
var appId = "wxb11529c136998cb6"
var text = "<xml><ToUserName><![CDATA[oia2Tj我是中文jewbmiOUlr6X-1crbLOvLw]]></ToUserName><FromUserName><![CDATA[gh_7f083739789a]]></FromUserName><CreateTime>1407743423</CreateTime><MsgType><![CDATA[video]]></MsgType><Video><MediaId><![CDATA[eYJ1MbwPRJtOvIEabaxHs7TX2D-HV71s79GUxqdUkjm6Gs2Ed1KF3ulAOA9H1xG0]]></MediaId><Title><![CDATA[testCallBackReplyVideo]]></Title><Description><![CDATA[testCallBackReplyVideo]]></Description></Video></xml>"

func TestBizMsg_Encrypt(t *testing.T) {
	biz := NewBizMsg(token, encodingAesKey, appId)
	result, err := biz.Encrypt(text, timeStamp, nonce)
	t.Log(result, err)
	result, err = biz.Decrypt(result, timeStamp, nonce)
	t.Log(result, err)
}
