package message

import "strings"

type EventType string

const (
	EventTypeSubscribe                  EventType = "subscribe"                    // 订阅
	EventTypeUnsubscribe                EventType = "unsubscribe"                  // 取消订阅
	EventTypeScan                       EventType = "scan"                         // 用户已经关注公众号，则微信会将带场景值扫描事件推送给开发者
	EventTypeLocation                   EventType = "location"                     // 上报地理位置事件
	EventTypeClick                      EventType = "CLICK"                        // 点击菜单拉取消息时的事件推送
	EventTypeView                       EventType = "view"                         // 点击菜单跳转链接时的事件推送
	EventTypeScancodePush               EventType = "scancode_push"                // 扫码推事件的事件推送
	EventTypeScancodeWaitmsg            EventType = "scancode_waitmsg"             // 扫码推事件且弹出“消息接收中”提示框的事件推送
	EventTypePicSysphoto                EventType = "pic_sysphoto"                 // 弹出系统拍照发图的事件推送
	EventTypePicPhotoOrAlbum            EventType = "pic_photo_or_album"           // 弹出拍照或者相册发图的事件推送
	EventTypePicWeixin                  EventType = "pic_weixin"                   // 弹出微信相册发图器的事件推送
	EventTypeLocationSelect             EventType = "location_select"              // 弹出地理位置选择器的事件推送
	EventTypeTemplateSendJobFinish      EventType = "TEMPLATESENDJOBFINISH"        // 发送模板消息推送通知
	EventTypeUserEnterTempsession       EventType = "user_enter_tempsession"       // 会话事件
	EventTypeQualificationVerifySuccess EventType = "qualification_verify_success" // 资质认证成功（此时立即获得接口权限）
	EventTypeQualificationVerifyFail    EventType = "qualification_verify_fail"    // 资质认证失败
	EventTypeNamingVerifySuccess        EventType = "naming_verify_success"        // 名称认证成功（即命名成功）
	EventTypeNamingVerifyFail           EventType = "naming_verify_fail"           // 名称认证失败（这时虽然客户端不打勾，但仍有接口权限）
	EventTypeAnnualRenew                EventType = "annual_renew"                 // 年审通知
	EventTypeVerifyExpired              EventType = "verify_expired"               // 认证过期失效通知审通知
	EventTypePoiCheckNotify             EventType = "poi_check_notify"             // 审核事件推送
)

type EVTCDATA struct {
	Value EventType `xml:",cdata"`
}

type Event struct {
	Event EVTCDATA
}

func (e EventType) String() string {
	return string(e)
}

func (e EventType) Compare(evtType EventType) int {
	return strings.Compare(strings.ToLower(e.String()), evtType.String())
}

func (e *Event) Compare(evtType EventType) int {
	return e.Event.Value.Compare(evtType)
}
