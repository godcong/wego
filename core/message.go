package core

import (
	"github.com/godcong/wego/core/message"
)

type Message struct {
	message.Message
	/*message*/
	Content      CDATA
	PicUrl       CDATA // 图片链接（由系统生成）
	MediaId      CDATA // 图片消息媒体id，可以调用多媒体文件下载接口拉取数据。
	Title        CDATA // 标题
	AppId        CDATA // 小程序appid
	PagePath     CDATA // 小程序页面路径
	ThumbUrl     CDATA // 封面图片的临时cdn链接
	ThumbMediaId CDATA // 封面图片的临时素材id
	Items        []*message.NewItem
	Format       CDATA // 语音格式，如amr，speex等
	Recognition  CDATA // 语音识别结果，UTF8编码
	Location_X   float64
	Location_Y   float64
	Scale        int64
	Label        CDATA
	Description  CDATA // 消息描述
	Url          CDATA
	/*event*/
	message.Event
	EventKey  CDATA   // 事件KEY值，qrscene_为前缀，后面为二维码的参数值
	Ticket    CDATA   // 二维码的ticket，可用来换取二维码图片
	Latitude  float64 // 地理位置纬度
	Longitude float64 // 地理位置经度
	Precision float64 // 地理位置精度

	MenuID CDATA // 指菜单ID，如果是个性化菜单，则可以通过这个字段，知道是哪个规则的菜单被点击了。

	ScanCodeInfo     message.ScanCodeInfo     // 扫描信息
	SendPicsInfo     message.SendPicsInfo     // 发送的图片信息
	SendLocationInfo message.SendLocationInfo // 发送的位置信息

	Status CDATA // 	发送状态为成功

	ExpiredTime int64 // 有效期 (整形)，指的是时间戳，将于该时间戳认证过期
	FailTime    int64 // 失败发生时间 (整形)，时间戳
	FailReason  CDATA // 认证失败的原因
	// 名称认证成功（即命名成功）
	UniqId CDATA
	PoiId  CDATA
	Result CDATA
	Msg    CDATA `xml:"msg"`
}

// type Article struct {
// }
//
// func NewMessage() *Message {
// 	return &Message{}
// }

//
// func (m *Message) SetAttribute(key string, val interface{}) *Message {
// 	m.properties = append(m.properties, key)
// 	m.attributes.Set(key, val)
// 	return m
// }
//
// func (m *Message) SetAttributes(m0 Map) *Message {
// 	for k, v := range m0 {
// 		m.SetAttribute(k, v)
// 	}
// 	return m
// }
//
// func (m *Message) GetAttribute(key string) interface{} {
// 	return m.attributes.Get(key)
// }
//
// func (m *Message) GetAttributes(keys []string) []interface{} {
// 	var m0 []interface{}
// 	for _, v := range keys {
// 		m0 = append(m0, m.attributes.Get(v))
// 	}
// 	return m0
// }

func (m *Message) SetType(msgType message.MsgType) *Message {
	m.MsgType = message.MSGCDATA{MsgType: msgType}
	return m
}

func (m *Message) GetType() message.MsgType {
	return m.MsgType.MsgType
}

//func (m *Message) Text() message.Text {
//	var text message.Text
//	text.Message = m.Message
//	text.Content = m.Content
//	return text
//}
