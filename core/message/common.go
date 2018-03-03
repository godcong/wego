package message

type MsgType string
type EvtType string

const (
	MESSAGE_TYPE_TEXT            MsgType = "text"                      //表示文本消息
	MESSAGE_TYPE_IMAGE           MsgType = "image"                     //表示图片消息
	MESSAGE_TYPE_VOICE           MsgType = "voice"                     //表示语音消息
	MESSAGE_TYPE_VIDEO           MsgType = "video"                     //表示视频消息
	MESSAGE_TYPE_SHORTVIDEO      MsgType = "shortvideo"                //表示短视频消息[限接收]
	MESSAGE_TYPE_LOCATION        MsgType = "location"                  //表示坐标消息[限接收]
	MESSAGE_TYPE_LINK            MsgType = "link"                      //表示链接消息[限接收]
	MESSAGE_TYPE_MUSIC           MsgType = "music"                     //表示音乐消息[限回复]
	MESSAGE_TYPE_NEWS            MsgType = "news"                      //表示图文消息[限回复]
	MESSAGE_TYPE_TRANSFER        MsgType = "transfer_customer_service" //表示消息消息转发到客服
	MESSAGE_TYPE_EVENT           MsgType = "event"                     //表示事件推送消息
	MESSAGE_TYPE_MINIPROGRAMPAGE MsgType = "miniprogrampage"
)

const (
	EVENT_TYPE_SUBSCRIBE              EvtType = "subscribe"              //订阅
	EVENT_TYPE_UNSUBSCRIBE            EvtType = "unsubscribe"            //取消订阅
	EVENT_TYPE_SCAN                   EvtType = "scan"                   //用户已经关注公众号，则微信会将带场景值扫描事件推送给开发者
	EVENT_TYPE_LOCATION               EvtType = "location"               //上报地理位置事件
	EVENT_TYPE_CLICK                  EvtType = "click"                  //点击菜单拉取消息时的事件推送
	EVENT_TYPE_VIEW                   EvtType = "view"                   //点击菜单跳转链接时的事件推送
	EVENT_TYPE_SCANCODEPUSH           EvtType = "scancode_push"          //扫码推事件的事件推送
	EVENT_TYPE_SCANCODEWAITMSG        EvtType = "scancode_waitmsg"       //扫码推事件且弹出“消息接收中”提示框的事件推送
	EVENT_TYPE_PICSYSPHOTO            EvtType = "pic_sysphoto"           //弹出系统拍照发图的事件推送
	EVENT_TYPE_PICPHOTOORALBUM        EvtType = "pic_photo_or_album"     //弹出拍照或者相册发图的事件推送
	EVENT_TYPE_PICWEIXIN              EvtType = "pic_weixin"             //弹出微信相册发图器的事件推送
	EVENT_TYPE_LOCATIONSELECT         EvtType = "location_select"        //弹出地理位置选择器的事件推送
	EVENT_TYPE_TEMPLATESENDJOBFINISH  EvtType = "templatesendjobfinish"  //发送模板消息推送通知
	EVENT_TYPE_USER_ENTER_TEMPSESSION EvtType = "user_enter_tempsession" //会话事件
)

type Message struct {
	XMLName      string `xml:"xml"`
	MsgType      MsgType
	MsgId        int64
	ToUserName   string
	FromUserName string
	CreateTime   int64
}

type Event struct {
	Event EvtType
}
