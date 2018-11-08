package wego_test

import (
	"github.com/godcong/wego"
	"github.com/godcong/wego/util"
	"testing"
)

// TestOfficialAccount ...
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

// TestGetApp ...
func TestGetApp(t *testing.T) {

}

// TestCoreUrl ...
func TestCoreUrl(t *testing.T) {

}

// TestGetOfficialAccount ...
func TestGetOfficialAccount(t *testing.T) {
	base := wego.OfficialAccount().Base()
	resp := base.GetCallbackIP()
	t.Log(resp.ToMap())
	//testBase(t, oa)
}

// TestXml ...
func TestXml(t *testing.T) {
	json := `{
	"card": {
		"card_type": "GROUPON",
		"groupon": {
			"base_info": {
				"logo_url": "http://mmbiz.qpic.cn/mmbiz/iaL1LJM1mF9aRKPZJkmG8xXhiaHqkKSVMMWeN3hLut7X7hicFNjakmxibMLGWpXrEXB33367o7zHN0CwngnQY7zb7g/0",
				"brand_name": "微信餐厅",
				"code_type": "CODE_TYPE_TEXT",
				"title": "132元双人火锅套餐",
				"color": "Color010",
				"notice": "使用时向服务员出示此券",
				"service_phone": "020-88888888",
				"description": "不可与其他优惠同享\n如需团购券发票，请在消费时向商户提出\n店内均可使用，仅限堂食",
				"date_info": {
					"type": "DATE_TYPE_FIX_TIME_RANGE",
					"begin_timestamp": 1397577600,
					"end_timestamp": 1472724261
				},
				"sku": {
					"quantity": 500000
				},
				"use_limit": 100,
				"get_limit": 3,
				"use_custom_code": false,
				"bind_openid": false,
				"can_share": true,
				"can_give_friend": true,
				"location_id_list": [123,12321,345345],
				"center_title": "顶部居中按钮",
				"center_sub_title": "按钮下方的wording",
				"center_url": "www.qq.com",
				"custom_url_name": "立即使用",
				"custom_url": "http://www.qq.com",
				"custom_url_sub_title": "6个汉字tips",
				"promotion_url_name": "更多优惠",
				"promotion_url": "http://www.qq.com",
				"source": "大众点评"
			},
			"advanced_info": {
				"use_condition": {
					"accept_category": "鞋类",
					"reject_category": "阿迪达斯",
					"can_use_with_other_discount": true
				},
				"abstract": {
					"abstract": "微信餐厅推出多种新季菜品，期待您的光临",
					"icon_url_list": [
						"http://mmbiz.qpic.cn/mmbiz/p98FjXy8LacgHxp3sJ3vn97bGLz0ib0Sfz1bjiaoOYA027iasqSG0sj\n piby4vce3AtaPu6cIhBHkt6IjlkY9YnDsfw/0"
					]
				},
				"text_image_list": [
					{
						"image_url": "http://mmbiz.qpic.cn/mmbiz/p98FjXy8LacgHxp3sJ3vn97bGLz0ib0Sfz1bjiaoOYA027iasqSG0sjpiby4vce3AtaPu6cIhBHkt6IjlkY9YnDsfw/0",
						"text": "此菜品精选食材，以独特的烹饪方法，最大程度地刺激食 客的味蕾"
					},
					{
						"image_url": "http://mmbiz.qpic.cn/mmbiz/p98FjXy8LacgHxp3sJ3vn97bGLz0ib0Sfz1bjiaoOYA027iasqSG0sj piby4vce3AtaPu6cIhBHkt6IjlkY9YnDsfw/0",
						"text": "此菜品迎合大众口味，老少皆宜，营养均衡"
					}
				],
				"time_limit": [
					{
						"type": "MONDAY",
						"begin_hour": 0,
						"end_hour": 10,
						"begin_minute": 10,
						"end_minute": 59
					},
					{
						"type": "HOLIDAY"
					}
				],
				"business_service": [
					"BIZ_SERVICE_FREE_WIFI",
					"BIZ_SERVICE_WITH_PET",
					"BIZ_SERVICE_FREE_PARK",
					"BIZ_SERVICE_DELIVER"
				]
			},
			"deal_detail": "以下锅底2选1（有菌王锅、麻辣锅、大骨锅、番茄锅、清补 凉锅、酸菜鱼锅可选）：\n大锅1份 12元\n小锅2份 16元 "
		}
	}
}`
	m := util.JSONToMap([]byte(json))
	t.Log(m)
	x := m.ToXML()
	t.Log(string(x))
	t.Log(util.XMLToMap(x))
}

//func testBase(t *testing.T, account wego.OfficialAccount) {
//log.Println(account.Base().GetCallbackIP())
//log.Println(account.Base().ClearQuota())
//}
