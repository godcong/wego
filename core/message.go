package core

type MessageType string
type EventType string

const (
	MESSAGE_TYPE_TEXT            MessageType = "text"                      //表示文本消息
	MESSAGE_TYPE_IMAGE           MessageType = "image"                     //表示图片消息
	MESSAGE_TYPE_VOICE           MessageType = "voice"                     //表示语音消息
	MESSAGE_TYPE_VIDEO           MessageType = "video"                     //表示视频消息
	MESSAGE_TYPE_SHORTVIDEO      MessageType = "shortvideo"                //表示短视频消息[限接收]
	MESSAGE_TYPE_LOCATION        MessageType = "location"                  //表示坐标消息[限接收]
	MESSAGE_TYPE_LINK            MessageType = "link"                      //表示链接消息[限接收]
	MESSAGE_TYPE_MUSIC           MessageType = "music"                     //表示音乐消息[限回复]
	MESSAGE_TYPE_NEWS            MessageType = "news"                      //表示图文消息[限回复]
	MESSAGE_TYPE_TRANSFER        MessageType = "transfer_customer_service" //表示消息消息转发到客服
	MESSAGE_TYPE_EVENT           MessageType = "event"                     //表示事件推送消息
	MESSAGE_TYPE_MINIPROGRAMPAGE MessageType = "miniprogrampage"
)

const (
	EVENT_TYPE_SUBSCRIBE              EventType = "subscribe"              //订阅
	EVENT_TYPE_UNSUBSCRIBE            EventType = "unsubscribe"            //取消订阅
	EVENT_TYPE_SCAN                   EventType = "scan"                   //用户已经关注公众号，则微信会将带场景值扫描事件推送给开发者
	EVENT_TYPE_LOCATION               EventType = "location"               //上报地理位置事件
	EVENT_TYPE_CLICK                  EventType = "click"                  //点击菜单拉取消息时的事件推送
	EVENT_TYPE_VIEW                   EventType = "view"                   //点击菜单跳转链接时的事件推送
	EVENT_TYPE_SCANCODEPUSH           EventType = "scancode_push"          //扫码推事件的事件推送
	EVENT_TYPE_SCANCODEWAITMSG        EventType = "scancode_waitmsg"       //扫码推事件且弹出“消息接收中”提示框的事件推送
	EVENT_TYPE_PICSYSPHOTO            EventType = "pic_sysphoto"           //弹出系统拍照发图的事件推送
	EVENT_TYPE_PICPHOTOORALBUM        EventType = "pic_photo_or_album"     //弹出拍照或者相册发图的事件推送
	EVENT_TYPE_PICWEIXIN              EventType = "pic_weixin"             //弹出微信相册发图器的事件推送
	EVENT_TYPE_LOCATIONSELECT         EventType = "location_select"        //弹出地理位置选择器的事件推送
	EVENT_TYPE_TEMPLATESENDJOBFINISH  EventType = "templatesendjobfinish"  //发送模板消息推送通知
	EVENT_TYPE_USER_ENTER_TEMPSESSION EventType = "user_enter_tempsession" //会话事件
)

type Message struct {
	messageType MessageType
	url         string
	to          string
	from        string
	create      int64

	id int64

	attributes Map
	properties []string
	aliases    Map
}

type Article struct {
}

func NewMessage() *Message {
	return &Message{}
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

func (m *Message) SetType(messageType MessageType) *Message {
	m.messageType = messageType
	return m
}

func (m *Message) GetType() MessageType {
	return m.messageType
}
