package message

import "strings"

type EvtType string

const (
	EventSubscribe                  EvtType = "subscribe"                    // 订阅
	EventUnsubscribe                EvtType = "unsubscribe"                  // 取消订阅
	EventScan                       EvtType = "scan"                         // 用户已经关注公众号，则微信会将带场景值扫描事件推送给开发者
	EventLocation                   EvtType = "location"                     // 上报地理位置事件
	EventClick                      EvtType = "click"                        // 点击菜单拉取消息时的事件推送
	EventView                       EvtType = "view"                         // 点击菜单跳转链接时的事件推送
	EventScancodePush               EvtType = "scancode_push"                // 扫码推事件的事件推送
	EventScancodeWaitmsg            EvtType = "scancode_waitmsg"             // 扫码推事件且弹出“消息接收中”提示框的事件推送
	EventPicSysphoto                EvtType = "pic_sysphoto"                 // 弹出系统拍照发图的事件推送
	EventPicPhotoOrAlbum            EvtType = "pic_photo_or_album"           // 弹出拍照或者相册发图的事件推送
	EventPicWeixin                  EvtType = "pic_weixin"                   // 弹出微信相册发图器的事件推送
	EventLocationSelect             EvtType = "location_select"              // 弹出地理位置选择器的事件推送
	EventTemplatesendjobfinish      EvtType = "templatesendjobfinish"        // 发送模板消息推送通知
	EventUserEnterTempsession       EvtType = "user_enter_tempsession"       // 会话事件
	EventQualificationVerifySuccess EvtType = "qualification_verify_success" // 资质认证成功（此时立即获得接口权限）
	EventQualificationVerifyFail    EvtType = "qualification_verify_fail"    // 资质认证失败
	EventNamingVerifySuccess        EvtType = "naming_verify_success"        // 名称认证成功（即命名成功）
	EventNamingVerifyFail           EvtType = "naming_verify_fail"           // 名称认证失败（这时虽然客户端不打勾，但仍有接口权限）
	EventAnnualRenew                EvtType = "annual_renew"                 // 年审通知
	EventVerifyExpired              EvtType = "verify_expired"               // 认证过期失效通知审通知

)

type EVTCDATA struct {
	Value EvtType `xml:",cdata"`
}

type Event struct {
	Event EVTCDATA
}

func (e EvtType) String() string {
	return string(e)
}

func (e EvtType) Compare(evtType EvtType) int {
	return strings.Compare(strings.ToLower(e.String()), evtType.String())
}

func (e *Event) Compare(evtType EvtType) int {
	return e.Event.Value.Compare(evtType)
}
