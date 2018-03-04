package wego_test

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/godcong/wego"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/core/message"
)

func TestOfficialAccount(t *testing.T) {
	//o := wego.GetOfficialAccount().Base()
	//log.Println(o.GetCallbackIp())

}

var msgText = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>1348831860</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[thisisatest]]></Content><MsgId>1234567890123456</MsgId></xml>`)
var msgImage = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>1348831860</CreateTime><MsgType><![CDATA[image]]></MsgType><PicUrl><![CDATA[thisisaurl]]></PicUrl><MediaId><![CDATA[media_id]]></MediaId><MsgId>1234567890123456</MsgId></xml>`)
var msgVoice = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>1357290913</CreateTime><MsgType><![CDATA[voice]]></MsgType><MediaId><![CDATA[media_id]]></MediaId><Format><![CDATA[Format]]></Format><MsgId>1234567890123456</MsgId></xml>`)
var msgVoice2 = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>1357290913</CreateTime><MsgType><![CDATA[voice]]></MsgType><MediaId><![CDATA[media_id]]></MediaId><Format><![CDATA[Format]]></Format><Recognition><![CDATA[腾讯微信团队]]></Recognition><MsgId>1234567890123456</MsgId></xml>`)
var msgVideo = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>1357290913</CreateTime><MsgType><![CDATA[video]]></MsgType><MediaId><![CDATA[media_id]]></MediaId><ThumbMediaId><![CDATA[thumb_media_id]]></ThumbMediaId><MsgId>1234567890123456</MsgId></xml>`)
var msgShortVideo = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>1357290913</CreateTime><MsgType><![CDATA[shortvideo]]></MsgType><MediaId><![CDATA[media_id]]></MediaId><ThumbMediaId><![CDATA[thumb_media_id]]></ThumbMediaId><MsgId>1234567890123456</MsgId></xml>`)
var msgLocaltion = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>1351776360</CreateTime><MsgType><![CDATA[location]]></MsgType><Location_X>23.134521</Location_X><Location_Y>113.358803</Location_Y><Scale>20</Scale><Label><![CDATA[位置信息]]></Label><MsgId>1234567890123456</MsgId></xml>`)
var msgLink = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>1351776360</CreateTime><MsgType><![CDATA[link]]></MsgType><Title><![CDATA[公众平台官网链接]]></Title><Description><![CDATA[公众平台官网链接]]></Description><Url><![CDATA[url]]></Url><MsgId>1234567890123456</MsgId></xml>`)

var evtSubscribe = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[FromUser]]></FromUserName><CreateTime>123456789</CreateTime><MsgType><![CDATA[event]]></MsgType><Event><![CDATA[subscribe]]></Event></xml>`)
var evtQRScene = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[FromUser]]></FromUserName><CreateTime>123456789</CreateTime><MsgType><![CDATA[event]]></MsgType><Event><![CDATA[subscribe]]></Event><EventKey><![CDATA[qrscene_123123]]></EventKey><Ticket><![CDATA[TICKET]]></Ticket></xml>`)
var evtScan = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[FromUser]]></FromUserName><CreateTime>123456789</CreateTime><MsgType><![CDATA[event]]></MsgType><Event><![CDATA[SCAN]]></Event><EventKey><![CDATA[SCENE_VALUE]]></EventKey><Ticket><![CDATA[TICKET]]></Ticket></xml>`)
var evtLocation = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>123456789</CreateTime><MsgType><![CDATA[event]]></MsgType><Event><![CDATA[LOCATION]]></Event><Latitude>23.137466</Latitude><Longitude>113.352425</Longitude><Precision>119.385040</Precision></xml>`)
var evtClick = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[FromUser]]></FromUserName><CreateTime>123456789</CreateTime><MsgType><![CDATA[event]]></MsgType><Event><![CDATA[CLICK]]></Event><EventKey><![CDATA[EVENTKEY]]></EventKey></xml>`)
var evtView = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[FromUser]]></FromUserName><CreateTime>123456789</CreateTime><MsgType><![CDATA[event]]></MsgType><Event><![CDATA[VIEW]]></Event><EventKey><![CDATA[www.qq.com]]></EventKey></xml>`)

var rltText = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>12345678</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[你好]]></Content></xml>`)
var rltImage = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>12345678</CreateTime><MsgType><![CDATA[image]]></MsgType><Image><MediaId><![CDATA[media_id]]></MediaId></Image></xml>`)
var rltVoice = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>12345678</CreateTime><MsgType><![CDATA[voice]]></MsgType><Voice><MediaId><![CDATA[media_id]]></MediaId></Voice></xml>`)
var rltVideo = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>12345678</CreateTime><MsgType><![CDATA[video]]></MsgType><Video><MediaId><![CDATA[media_id]]></MediaId><Title><![CDATA[title]]></Title><Description><![CDATA[description]]></Description></Video></xml>`)
var rltMusic = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>12345678</CreateTime><MsgType><![CDATA[music]]></MsgType><Music><Title><![CDATA[TITLE]]></Title><Description><![CDATA[DESCRIPTION]]></Description><MusicUrl><![CDATA[MUSIC_Url]]></MusicUrl><HQMusicUrl><![CDATA[HQ_MUSIC_Url]]></HQMusicUrl><ThumbMediaId><![CDATA[media_id]]></ThumbMediaId></Music></xml>`)
var rltNews = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>12345678</CreateTime><MsgType><![CDATA[news]]></MsgType><ArticleCount>2</ArticleCount><Articles><item><Title><![CDATA[title1]]></Title><Description><![CDATA[description1]]></Description><PicUrl><![CDATA[picurl]]></PicUrl><Url><![CDATA[url]]></Url></item><item><Title><![CDATA[title]]></Title><Description><![CDATA[description]]></Description><PicUrl><![CDATA[picurl]]></PicUrl><Url><![CDATA[url]]></Url></item></Articles></xml>`)

func TestGetApp(t *testing.T) {

	server := wego.GetOfficialAccount().Server()
	//s := official_account.NewServer()
	server.RegisterCallback(func(mess *core.Message) []byte {
		v, _ := json.Marshal(message.Text{
			Message: message.Message{},
			Content: message.CDATA{},
		})
		core.Info(*mess)
		if mess.Event.Compare(message.EventView) == 0 {
			core.Info(message.EventView)
		}
		return v
	})

	rlist := [][]byte{
		rltText,
		rltImage,
		rltVoice,
		rltVideo,
		rltMusic,
		rltNews,
	}
	count := 0

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if count >= len(rlist) {
			count = 0
		}

		server.Monitor(w, r)
		//t.Log(string(body))
		//var msg core.Message
		//var msg core.Message
		//e = xml.Unmarshal(body, &msg)

		////xml.NewDecoder(bytes.NewReader(body)).
		//if msg.GetType() == message.TypeText {
		//	b, e := xml.Marshal(msg)
		//	t.Log(string(b), e)
		//	t.Log(m)
		//}
		count++
	}))
	defer ts.Close()

	list := [][]byte{msgText,
		msgImage,
		msgVoice,
		msgVoice2,
		msgVideo,
		msgShortVideo,
		msgLocaltion,
		msgLink,
		evtSubscribe,
		evtQRScene,
		evtScan,
		evtLocation,
		evtClick,
		evtView,
	}

	for _, v := range list {
		resp, e := http.Post(ts.URL+"/callback", "Content-Type:application/xml", bytes.NewReader(v))
		msg := new(core.Message)
		b, e := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(b, msg)
		core.Info(*msg, e)
	}

}
