package core

import (
	"github.com/godcong/wego/core/message"
)

/*Message 回调消息结构 */
type Message struct {
	message.Message
	/*message*/
	Content      message.CDATA      `xml:"content"`
	PicURL       message.CDATA      `xml:"pic_url"`        // 图片链接（由系统生成）
	MediaID      message.CDATA      `xml:"media_id"`       // 图片消息媒体id，可以调用多媒体文件下载接口拉取数据。
	Title        message.CDATA      `xml:"title"`          // 标题
	AppID        message.CDATA      `xml:"app_id"`         // 小程序appid
	PagePath     message.CDATA      `xml:"page_path"`      // 小程序页面路径
	ThumbURL     message.CDATA      `xml:"thumb_url"`      // 封面图片的临时cdn链接
	ThumbMediaID message.CDATA      `xml:"thumb_media_id"` // 封面图片的临时素材id
	Items        []*message.NewItem `xml:"items"`
	Format       message.CDATA      `xml:"format"`      // 语音格式，如amr，speex等
	Recognition  message.CDATA      `xml:"recognition"` // 语音识别结果，UTF8编码
	LocationX    float64            `xml:"location_x"`
	LocationY    float64            `xml:"location_y"`
	Scale        int64              `xml:"scale"`
	Label        message.CDATA      `xml:"label"`
	Description  message.CDATA      `xml:"description"` // 消息描述
	URL          message.CDATA      `xml:"url"`
	/*event*/
	message.Event
	EventKey  message.CDATA `xml:"event_key"` // 事件KEY值，qrscene_为前缀，后面为二维码的参数值
	Ticket    message.CDATA `xml:"ticket"`    // 二维码的ticket，可用来换取二维码图片
	Latitude  float64       `xml:"latitude"`  // 地理位置纬度
	Longitude float64       `xml:"longitude"` // 地理位置经度
	Precision float64       `xml:"precision"` // 地理位置精度

	MenuID message.CDATA `xml:"menu_id"` // 指菜单ID，如果是个性化菜单，则可以通过这个字段，知道是哪个规则的菜单被点击了。

	ScanCodeInfo     message.ScanCodeInfo     `xml:"scan_code_info"`     // 扫描信息
	SendPicsInfo     message.SendPicsInfo     `xml:"send_pics_info"`     // 发送的图片信息
	SendLocationInfo message.SendLocationInfo `xml:"send_location_info"` // 发送的位置信息

	Status      message.CDATA `xml:"status"`       // 	发送状态为成功
	ExpiredTime int64         `xml:"expired_time"` // 有效期 (整形)，指的是时间戳，将于该时间戳认证过期
	FailTime    int64         `xml:"fail_time"`    // 失败发生时间 (整形)，时间戳
	FailReason  message.CDATA `xml:"fail_reason"`  // 认证失败的原因
	// 名称认证成功（即命名成功）
	UniqID      message.CDATA `xml:"uniq_id"`
	PoiID       message.CDATA `xml:"poi_id"`
	Result      message.CDATA `xml:"result"`
	Msg         message.CDATA `xml:"msg"`
	SessionFrom message.CDATA `xml:"session_from"`

	OrderID     message.CDATA `xml:"order_id"`
	OrderStatus int64         `xml:"order_status"`
	ProductID   message.CDATA `xml:"product_id"`
	SkuInfo     message.CDATA `xml:"sku_info"`
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

/*SetType 设置消息类型 */
func (m *Message) SetType(msgType message.MsgType) *Message {
	m.MsgType = message.MSGCDATA{MsgType: msgType}
	return m
}

/*GetType 获取消息类型 */
func (m *Message) GetType() message.MsgType {
	return m.MsgType.MsgType
}

//func (m *Message) Text() message.Text {
//	var text message.Text
//	text.Message = m.Message
//	text.Content = m.Content
//	return text
//}
