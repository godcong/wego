package wego

import (
	"encoding/xml"
	"github.com/godcong/wego/util"
	"strings"
)

/*Messager Messager */
type Messager interface {
	ToXML() ([]byte, error)
	ToJSON() ([]byte, error)
}

/*CDATA CDATA */
type CDATA = util.CDATA

/*MsgType MsgType */
type MsgType string

/*message types */
const (
	MsgTypeText            MsgType = "text"                      //表示文本消息	``
	MsgTypeImage           MsgType = "image"                     //表示图片消息
	MsgTypeVoice           MsgType = "voice"                     //表示语音消息
	MsgTypeVideo           MsgType = "video"                     //表示视频消息
	MsgTypeShortvideo      MsgType = "shortvideo"                //表示短视频消息[限接收]
	MsgTypeLocation        MsgType = "location"                  //表示坐标消息[限接收]
	MsgTypeLink            MsgType = "link"                      //表示链接消息[限接收]
	MsgTypeMusic           MsgType = "music"                     //表示音乐消息[限回复]
	MsgTypeNews            MsgType = "news"                      //表示图文消息[限回复]
	MsgTypeTransfer        MsgType = "transfer_customer_service" //表示消息消息转发到客服
	MsgTypeEvent           MsgType = "event"                     //表示事件推送消息
	MsgTypeMiniprogrampage MsgType = "miniprogrampage"
)

/*MSGCDATA MSGCDATA */
type MSGCDATA struct {
	MsgType `xml:",cdata"`
}

/*Message Message */
type Message struct {
	XMLName      xml.Name `xml:"xml"`
	MsgType      MSGCDATA `xml:"MsgType"`
	MsgID        int64    `xml:"MsgID,omitempty"`
	ToUserName   CDATA    `xml:"to_user_name"`
	FromUserName CDATA    `xml:"from_user_name"`
	CreateTime   int64    `xml:"create_time"`
}

/*String String */
func (e MsgType) String() string {
	return string(e)
}

/*Compare compare message type*/
func (e MsgType) Compare(msgType MsgType) int {
	return strings.Compare(strings.ToLower(e.String()), msgType.String())
}

/*Compare compare message type*/
func (e *Message) Compare(msgType MsgType) int {
	return e.MsgType.Compare(msgType)
}

/*NewMessage create a new message */
func NewMessage(msgType MsgType, toUser, fromUser string, msgID, createTime int64) *Message {
	return &Message{
		MsgType: MSGCDATA{
			MsgType: msgType,
		},
		MsgID: msgID,
		ToUserName: CDATA{
			Value: toUser,
		},
		FromUserName: CDATA{
			Value: fromUser,
		},
		CreateTime: createTime,
	}
}

/*EventType EventType */
type EventType string

/*event types */
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
	EventTypeMerchantOrder              EventType = "merchant_order"               //订单付款通知
)

/*EVTCDATA EVTCDATA */
type EVTCDATA struct {
	Value EventType `xml:",cdata"`
}

/*Event Event */
type Event struct {
	Event EVTCDATA
}

/*String String */
func (e EventType) String() string {
	return string(e)
}

/*Compare compare event type */
func (e EventType) Compare(evtType EventType) int {
	return strings.Compare(strings.ToLower(e.String()), evtType.String())
}

/*Compare compare event type */
func (e *Event) Compare(evtType EventType) int {
	return e.Event.Value.Compare(evtType)
}
