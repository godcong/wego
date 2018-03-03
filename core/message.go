package core

import "github.com/godcong/wego/core/message"

type Message struct {
	message.Message
	/*message*/
	Content      string
	PicUrl       string //图片链接（由系统生成）
	MediaId      string //图片消息媒体id，可以调用多媒体文件下载接口拉取数据。
	Title        string //标题
	AppId        string //小程序appid
	PagePath     string //小程序页面路径
	ThumbUrl     string //封面图片的临时cdn链接
	ThumbMediaId string //封面图片的临时素材id
	Items        []*message.NewItem
	Format       string //语音格式，如amr，speex等
	Recognition  string //语音识别结果，UTF8编码
	Location_X   float64
	Location_Y   float64
	Scale        int64
	Label        string
	Description  string //消息描述
	Url          string
	/*event*/
	message.Event
	EventKey  string  //事件KEY值，qrscene_为前缀，后面为二维码的参数值
	Ticket    string  //二维码的ticket，可用来换取二维码图片
	Latitude  float64 //地理位置纬度
	Longitude float64 //地理位置经度
	Precision float64 //地理位置精度
	//attributes   Map
	//properties   []string
	//aliases      Map
}

//type Article struct {
//}
//
//func NewMessage() *Message {
//	return &Message{}
//}

//
//func (m *Message) SetAttribute(key string, val interface{}) *Message {
//	m.properties = append(m.properties, key)
//	m.attributes.Set(key, val)
//	return m
//}
//
//func (m *Message) SetAttributes(m0 Map) *Message {
//	for k, v := range m0 {
//		m.SetAttribute(k, v)
//	}
//	return m
//}
//
//func (m *Message) GetAttribute(key string) interface{} {
//	return m.attributes.Get(key)
//}
//
//func (m *Message) GetAttributes(keys []string) []interface{} {
//	var m0 []interface{}
//	for _, v := range keys {
//		m0 = append(m0, m.attributes.Get(v))
//	}
//	return m0
//}

func (m *Message) SetType(msgType message.MsgType) *Message {
	m.MsgType = msgType
	return m
}

func (m *Message) GetType() message.MsgType {
	return m.MsgType
}

func (m *Message) Text() message.Text {
	var text message.Text
	text.Message = m.Message
	text.Content = m.Content
	return text
}
