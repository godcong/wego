package wego_test

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/godcong/wego"
	_ "github.com/godcong/wego/app/official"
	"github.com/godcong/wego/config"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/core/message"
	"github.com/godcong/wego/log"
)

func TestOfficialAccount(t *testing.T) {
	//wego.GetApp().Get()
	//log.Println(o.GetCallbackIP())

}

var msgText = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>1348831860</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[thisisatest]]></Content><MsgID>1234567890123456</MsgID></xml>`)
var msgImage = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>1348831860</CreateTime><MsgType><![CDATA[image]]></MsgType><PicURL><![CDATA[thisisaurl]]></PicURL><MediaID><![CDATA[media_id]]></MediaID><MsgID>1234567890123456</MsgID></xml>`)
var msgVoice = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>1357290913</CreateTime><MsgType><![CDATA[voice]]></MsgType><MediaID><![CDATA[media_id]]></MediaID><Format><![CDATA[Format]]></Format><MsgID>1234567890123456</MsgID></xml>`)
var msgVoice2 = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>1357290913</CreateTime><MsgType><![CDATA[voice]]></MsgType><MediaID><![CDATA[media_id]]></MediaID><Format><![CDATA[Format]]></Format><Recognition><![CDATA[腾讯微信团队]]></Recognition><MsgID>1234567890123456</MsgID></xml>`)
var msgVideo = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>1357290913</CreateTime><MsgType><![CDATA[video]]></MsgType><MediaID><![CDATA[media_id]]></MediaID><ThumbMediaID><![CDATA[thumb_media_id]]></ThumbMediaID><MsgID>1234567890123456</MsgID></xml>`)
var msgShortVideo = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>1357290913</CreateTime><MsgType><![CDATA[shortvideo]]></MsgType><MediaID><![CDATA[media_id]]></MediaID><ThumbMediaID><![CDATA[thumb_media_id]]></ThumbMediaID><MsgID>1234567890123456</MsgID></xml>`)
var msgLocaltion = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>1351776360</CreateTime><MsgType><![CDATA[location]]></MsgType><LocationX>23.134521</LocationX><LocationY>115.358803</LocationY><Scale>20</Scale><Label><![CDATA[位置信息]]></Label><MsgID>1234567890123456</MsgID></xml>`)
var msgLink = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>1351776360</CreateTime><MsgType><![CDATA[link]]></MsgType><Title><![CDATA[公众平台官网链接]]></Title><Description><![CDATA[公众平台官网链接]]></Description><URL><![CDATA[url]]></URL><MsgID>1234567890123456</MsgID></xml>`)

var evtSubscribe = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[FromUser]]></FromUserName><CreateTime>123456789</CreateTime><MsgType><![CDATA[event]]></MsgType><Event><![CDATA[subscribe]]></Event></xml>`)
var evtQRScene = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[FromUser]]></FromUserName><CreateTime>123456789</CreateTime><MsgType><![CDATA[event]]></MsgType><Event><![CDATA[subscribe]]></Event><EventKey><![CDATA[qrscene_123123]]></EventKey><Ticket><![CDATA[TICKET]]></Ticket></xml>`)
var evtScan = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[FromUser]]></FromUserName><CreateTime>123456789</CreateTime><MsgType><![CDATA[event]]></MsgType><Event><![CDATA[SCAN]]></Event><EventKey><![CDATA[SCENE_VALUE]]></EventKey><Ticket><![CDATA[TICKET]]></Ticket></xml>`)
var evtLocation = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>123456789</CreateTime><MsgType><![CDATA[event]]></MsgType><Event><![CDATA[LOCATION]]></Event><Latitude>23.137466</Latitude><Longitude>113.352425</Longitude><Precision>119.385040</Precision></xml>`)
var evtClick = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[FromUser]]></FromUserName><CreateTime>123456789</CreateTime><MsgType><![CDATA[event]]></MsgType><Event><![CDATA[CLICK]]></Event><EventKey><![CDATA[EVENTKEY]]></EventKey></xml>`)
var evtView = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[FromUser]]></FromUserName><CreateTime>123456789</CreateTime><MsgType><![CDATA[event]]></MsgType><Event><![CDATA[VIEW]]></Event><EventKey><![CDATA[www.qq.com]]></EventKey></xml>`)
var evtView2 = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[FromUser]]></FromUserName><CreateTime>123456789</CreateTime><MsgType><![CDATA[event]]></MsgType><Event><![CDATA[VIEW]]></Event><EventKey><![CDATA[www.qq.com]]></EventKey><MenuId>MENUID</MenuId></xml>`)
var evtScancodePush = []byte(`<xml><ToUserName><![CDATA[gh_e136c6e50636]]></ToUserName><FromUserName><![CDATA[oMgHVjngRipVsoxg6TuX3vz6glDg]]></FromUserName><CreateTime>1408090502</CreateTime><MsgType><![CDATA[event]]></MsgType><Event><![CDATA[scancode_push]]></Event><EventKey><![CDATA[6]]></EventKey><ScanCodeInfo><ScanType><![CDATA[qrcode]]></ScanType><ScanResult><![CDATA[1]]></ScanResult></ScanCodeInfo></xml>`)
var evtScancodeWaitmsg = []byte(`<xml><ToUserName><![CDATA[gh_e136c6e50636]]></ToUserName><FromUserName><![CDATA[oMgHVjngRipVsoxg6TuX3vz6glDg]]></FromUserName><CreateTime>1408090606</CreateTime><MsgType><![CDATA[event]]></MsgType><Event><![CDATA[scancode_waitmsg]]></Event><EventKey><![CDATA[6]]></EventKey><ScanCodeInfo><ScanType><![CDATA[qrcode]]></ScanType><ScanResult><![CDATA[2]]></ScanResult></ScanCodeInfo></xml>`)
var evtPicSysphoto = []byte(`<xml><ToUserName><![CDATA[gh_e136c6e50636]]></ToUserName><FromUserName><![CDATA[oMgHVjngRipVsoxg6TuX3vz6glDg]]></FromUserName><CreateTime>1408090651</CreateTime><MsgType><![CDATA[event]]></MsgType><Event><![CDATA[pic_sysphoto]]></Event><EventKey><![CDATA[6]]></EventKey><SendPicsInfo><Count>1</Count><PicList><item><PicMd5Sum><![CDATA[1b5f7c23b5bf75682a53e7b6d163e185]]></PicMd5Sum></item><item><PicMd5Sum><![CDATA[22222222222222222222]]></PicMd5Sum></item></PicList></SendPicsInfo></xml>`)
var evtPicPhotoOrAlbum = []byte(`<xml><ToUserName><![CDATA[gh_e136c6e50636]]></ToUserName><FromUserName><![CDATA[oMgHVjngRipVsoxg6TuX3vz6glDg]]></FromUserName><CreateTime>1408090816</CreateTime><MsgType><![CDATA[event]]></MsgType><Event><![CDATA[pic_photo_or_album]]></Event><EventKey><![CDATA[6]]></EventKey><SendPicsInfo><Count>1</Count><PicList><item><PicMd5Sum><![CDATA[5a75aaca956d97be686719218f275c6b]]></PicMd5Sum></item></PicList></SendPicsInfo></xml>`)
var evtPicWeixin = []byte(`<xml><ToUserName><![CDATA[gh_e136c6e50636]]></ToUserName><FromUserName><![CDATA[oMgHVjngRipVsoxg6TuX3vz6glDg]]></FromUserName><CreateTime>1408090816</CreateTime><MsgType><![CDATA[event]]></MsgType><Event><![CDATA[pic_weixin]]></Event><EventKey><![CDATA[6]]></EventKey><SendPicsInfo><Count>1</Count><PicList><item><PicMd5Sum><![CDATA[5a75aaca956d97be686719218f275c6b]]></PicMd5Sum></item></PicList></SendPicsInfo></xml>`)
var evtLocationSelect = []byte(`<xml><ToUserName><![CDATA[gh_e136c6e50636]]></ToUserName><FromUserName><![CDATA[oMgHVjngRipVsoxg6TuX3vz6glDg]]></FromUserName><CreateTime>1408091189</CreateTime><MsgType><![CDATA[event]]></MsgType><Event><![CDATA[location_select]]></Event><EventKey><![CDATA[6]]></EventKey><SendLocationInfo><LocationX><![CDATA[23]]></LocationX><LocationY><![CDATA[6553600]]></LocationY><Scale><![CDATA[15]]></Scale><Label><![CDATA[ 广州市海珠区客村艺苑路 106号]]></Label><Poiname><![CDATA[]]></Poiname></SendLocationInfo></xml>`)

var rltText = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>12345678</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[你好]]></Content></xml>`)
var rltImage = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>12345678</CreateTime><MsgType><![CDATA[image]]></MsgType><Image><MediaID><![CDATA[media_id]]></MediaID></Image></xml>`)
var rltVoice = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>12345678</CreateTime><MsgType><![CDATA[voice]]></MsgType><Voice><MediaID><![CDATA[media_id]]></MediaID></Voice></xml>`)
var rltVideo = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>12345678</CreateTime><MsgType><![CDATA[video]]></MsgType><Video><MediaID><![CDATA[media_id]]></MediaID><Title><![CDATA[title]]></Title><Description><![CDATA[description]]></Description></Video></xml>`)
var rltMusic = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>12345678</CreateTime><MsgType><![CDATA[music]]></MsgType><Music><Title><![CDATA[TITLE]]></Title><Description><![CDATA[DESCRIPTION]]></Description><MusicUrl><![CDATA[MUSIC_Url]]></MusicUrl><HQMusicUrl><![CDATA[HQ_MUSIC_Url]]></HQMusicUrl><ThumbMediaID><![CDATA[media_id]]></ThumbMediaID></Music></xml>`)
var rltNews = []byte(`<xml><ToUserName><![CDATA[toUser]]></ToUserName><FromUserName><![CDATA[fromUser]]></FromUserName><CreateTime>12345678</CreateTime><MsgType><![CDATA[news]]></MsgType><ArticleCount>2</ArticleCount><Articles><item><Title><![CDATA[title1]]></Title><Description><![CDATA[description1]]></Description><PicURL><![CDATA[picurl]]></PicURL><URL><![CDATA[url]]></URL></item><item><Title><![CDATA[title]]></Title><Description><![CDATA[description]]></Description><PicURL><![CDATA[picurl]]></PicURL><URL><![CDATA[url]]></URL></item></Articles></xml>`)

func TestGetApp(t *testing.T) {

	server := wego.GetOfficialAccount().Server()
	//s := official_account.NewServer()
	server.RegisterCallback(func(mess *core.Message) message.Messager {
		//mess.LocationX = 123456789444
		//v, _ := xml.Marshal(*mess)
		//log.Info(string(v))
		//if mess.Event.Compare(message.EventView) == 0 {
		//	log.Info(message.EventView)
		//}
		return message.NewText(&mess.Message, "msg")
	})

	rlist := [][]byte{
		//rltText,
		//rltImage,
		//rltVoice,
		//rltVideo,
		//rltMusic,
		rltNews,
	}
	count := 0

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if count >= len(rlist) {
			count = 0
		}

		server.ServeHTTP(w, r)
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

	list := [][]byte{
		//msgText,
		//msgImage,
		//msgVoice,
		//msgVoice2,
		//msgVideo,
		//msgShortVideo,
		//msgLocaltion,
		//msgLink,
		//evtSubscribe,
		//evtQRScene,
		//evtScan,
		//evtLocation,
		//evtClick,
		//evtView,
		//evtView2,
		//evtScancodePush,
		//evtScancodeWaitmsg,
		evtPicSysphoto,
		evtPicPhotoOrAlbum,
		//evtPicWeixin,
		//evtLocationSelect,
	}

	for _, v := range list {
		resp, e := http.Post(ts.URL+"/callback", "Content-Type:application/xml", bytes.NewReader(v))
		msg := new(core.Message)
		b, e := ioutil.ReadAll(resp.Body)

		xml.Unmarshal(b, msg)
		log.Info(msg, e)
	}

}

func TestCoreUrl(t *testing.T) {
	conf := config.GetConfig("payment.default")
	url := core.NewURL(conf, core.NewClient(conf))
	l := url.ShortURL("https://y11e.com")
	log.Println(l)
}

func TestGetOfficialAccount(t *testing.T) {
	oa := wego.GetOfficialAccount()
	testBase(t, oa)
}

func testBase(t *testing.T, account wego.OfficialAccount) {
	log.Println(account.Base().GetCallbackIP())
	log.Println(account.Base().ClearQuota())
}
