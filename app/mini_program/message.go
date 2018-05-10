package mini_program

import (
	"github.com/godcong/wego/config"
	"github.com/godcong/wego/core/message"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/net"
	"github.com/godcong/wego/util"
)

func MessageSend() {
	//https://api.weixin.qq.com/cgi-bin/message/custom/send
}

type Message struct {
	config.Config
	*MiniProgram
	//client *core.Client
}

func newMessage(program *MiniProgram) *Message {
	Message := Message{
		Config:      defaultConfig,
		MiniProgram: program,
		//client:      program.GetClient(),
	}
	//Message.client.SetDomain(core.NewDomain(""))
	return &Message
}

func NewMessage() *Message {
	return newMessage(program)
}

//
// 接口调用请求说明
//
//http请求方式: POST
//https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN
//各消息类型所需的JSON数据包如下：
//
//发送文本消息
//
//{
//    "touser":"OPENID",
//    "msgtype":"text",
//    "text":
//    {
//         "content":"Hello World"
//    }
//}
//参数说明
//
//参数	是否必须	说明
//access_token	是	调用接口凭证
//touser	是	普通用户(openid)
//msgtype	是	消息类型，文本为text，图文链接为link
//content	是	文本消息内容
//media_id	是	发送的图片的媒体ID，通过新增素材接口上传图片文件获得。
//title	是	消息标题
//description	是	图文链接消息
//url	是	图文链接消息被点击后跳转的链接
//picurl	是	图文链接消息的图片链接，支持 JPG、PNG 格式，较好的效果为大图 640 X 320，小图 80 X 80
//pagepath	是	小程序的页面路径，跟app.json对齐，支持参数，比如pages/index/index?foo=bar
//thumb_media_id	是	小程序消息卡片的封面， image类型的media_id，通过新增素材接口上传图片文件获得，建议大小为520*416
func (m *Message) Send(messager message.Messager) *net.Response {
	log.Debug("Message|Send", messager)

	p := util.Map{
		"touser":  "oE_gl0Yr54fUjBhU5nBlP4hS2efo",
		"msgtype": "link",
		"link": util.Map{
			"title":       "Happy Day",
			"description": "Is Really A Happy Day",
			"url":         "https://mp.weixin.qq.com/mp/profile_ext?action=home&__biz=MzI1MzE3OTI2NA==&from=timeline&isappinstalled=0#wechat_redirect",
			//"content":   `<a href="https://mp.weixin.qq.com/s/JU0wJpo2HYUzIGsPcfbIuA" data-miniprogram-appid="appid" data-miniprogram-path="pages/index/index">点击跳小程序</a>`,
			"thumb_url": "THUMB_URL",
		},
	}

	key := m.token.GetToken().KeyMap()
	resp := m.client.HttpPostJson(
		m.client.Link(CUSTOM_SEND_URL_SUFFIX),
		key,
		p)
	return resp
}
