package core

type MessageType string
type EventType string

const (
	//表示文本消息
	MESSAGE_TYPE_TEXT MessageType = "text"
	//表示图片消息
	MESSAGE_TYPE_IMAGE = "image"
	//表示语音消息
	MESSAGE_TYPE_VOICE = "voice"
	//表示视频消息
	MESSAGE_TYPE_VIDEO = "video"
	//表示短视频消息[限接收]
	MESSAGE_TYPE_SHORTVIDEO = "shortvideo"
	//表示坐标消息[限接收]
	MESSAGE_TYPE_LOCATION = "location"
	//表示链接消息[限接收]
	MESSAGE_TYPE_LINK = "link"
	//表示音乐消息[限回复]
	MESSAGE_TYPE_MUSIC = "music"
	//表示图文消息[限回复]
	MESSAGE_TYPE_NEWS = "news"
	//表示消息消息转发到客服
	MESSAGE_TYPE_TRANSFER = "transfer_customer_service"
	//表示事件推送消息
	MESSAGE_TYPE_EVENT = "event"
)

const (
	//订阅
	EVENT_TYPE_SUBSCRIBE EventType = "subscribe"
	//取消订阅
	EVENT_TYPE_UNSUBSCRIBE = "unsubscribe"
	//用户已经关注公众号，则微信会将带场景值扫描事件推送给开发者
	EVENT_TYPE_SCAN = "SCAN"
	//上报地理位置事件
	EVENT_TYPE_LOCATION = "LOCATION"
	//点击菜单拉取消息时的事件推送
	EVENT_TYPE_CLICK = "CLICK"
	//点击菜单跳转链接时的事件推送
	EVENT_TYPE_VIEW = "VIEW"
	//扫码推事件的事件推送
	EVENT_TYPE_SCANCODEPUSH = "scancode_push"
	//扫码推事件且弹出“消息接收中”提示框的事件推送
	EVENT_TYPE_SCANCODEWAITMSG = "scancode_waitmsg"
	//弹出系统拍照发图的事件推送
	EVENT_TYPE_PICSYSPHOTO = "pic_sysphoto"
	//弹出拍照或者相册发图的事件推送
	EVENT_TYPE_PICPHOTOORALBUM = "pic_photo_or_album"
	//弹出微信相册发图器的事件推送
	EVENT_TYPE_PICWEIXIN = "pic_weixin"
	//弹出地理位置选择器的事件推送
	EVENT_TYPE_LOCATIONSELECT = "location_select"
	//发送模板消息推送通知
	EVENT_TYPE_TEMPLATESENDJOBFINISH = "TEMPLATESENDJOBFINISH"
)

type Message struct {
	url    string
	to     string
	from   string
	create int64
	typ    MessageType
	id     int64

	attributes Map
	properties []string
	aliases    Map
}

type MessageText struct {
	*Message
}

type MessageImage struct {
	*Message
}
type MessageMusic struct {
	*Message
}

func NewMessage() *Message {
	return &Message{}
}

func NewMessageText(content string) *MessageText {
	return &MessageText{
		Message: NewMessage().SetAttribute("Content", content),
	}
}

func NewMessageImage(mediaID string) *MessageImage {
	return &MessageImage{
		Message: NewMessage().SetAttribute("MediaID", mediaID),
	}
}

func NewMessageMusic(title, des, murl, hqurl, thumb string) *MessageMusic {
	return &MessageMusic{
		Message: NewMessage().SetAttributes(Map{
			"Title":        title,
			"Description":  des,
			"MusicURL":     murl,
			"HQMusicURL":   hqurl,
			"ThumbMediaID": thumb,
		}),
	}
}

func (m *Message) SetAttribute(key string, val interface{}) *Message {
	m.properties = append(m.properties, key)
	m.attributes.Set(key, val)
	return m
}

func (m *Message) SetAttributes(m0 Map) *Message {
	for k, v := range m0 {
		m.SetAttribute(k, v)
	}
	return m
}

func (m *Message) GetAttribute(key string) interface{} {
	return m.attributes.Get(key)
}

func (m *Message) GetAttributes(keys []string) []interface{} {
	var m0 []interface{}
	for _, v := range keys {
		m0 = append(m0, m.attributes.Get(v))
	}
	return m0
}

func (m *Message) SetType(typ MessageType) *Message {
	m.typ = typ
	return m
}

func (m *Message) GetType() MessageType {
	return m.typ
}
