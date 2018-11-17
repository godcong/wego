package official

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

/*Broadcasting 消息 */
type Broadcasting struct {
	*Account
}

func newBroadcasting(account *Account) *Broadcasting {
	return &Broadcasting{
		Account: account,
	}
}

// NewBroadcasting 消息
func NewBroadcasting(config *core.Config) *Broadcasting {
	return newBroadcasting(NewOfficialAccount(config))
}

//Send 根据OpenID列表群发【订阅号不可用，服务号认证后可用】
//接口调用请求说明
//http请求方式: POST
//https://api.weixin.qq.com/cgi-bin/message/mass/send?access_token=ACCESS_TOKEN
func (b *Broadcasting) Send(msg util.Map) core.Responder {
	token := b.accessToken.GetToken()
	return core.PostJSON(Link(messageMassSend), token.KeyMap(), msg)
}

//SendAll 根据标签进行群发【订阅号与服务号认证后均可用】
//接口调用请求说明
//http请求方式: POST
//https://api.weixin.qq.com/cgi-bin/message/mass/sendall?access_token=ACCESS_TOKEN
func (b *Broadcasting) SendAll(msg util.Map) core.Responder {
	token := b.accessToken.GetToken()
	return core.PostJSON(Link(messageMassSendall), token.KeyMap(), msg)
}

//Preview 预览接口【订阅号与服务号认证后均可用】
//开发者可通过该接口发送消息给指定用户，在手机端查看消息的样式和排版。为了满足第三方平台开发者的需求，在保留对openID预览能力的同时，增加了对指定微信号发送预览的能力，但该能力每日调用次数有限制（100次），请勿滥用。
//接口调用请求说明
//http请求方式: POST
//https://api.weixin.qq.com/cgi-bin/message/mass/preview?access_token=ACCESS_TOKEN
func (b *Broadcasting) Preview(msg util.Map) core.Responder {
	token := b.accessToken.GetToken()
	return core.PostJSON(Link(messageMassPreview), token.KeyMap(), msg)

}

//Delete 删除群发【订阅号与服务号认证后均可用】
//群发之后，随时可以通过该接口删除群发。
//接口调用请求说明
//http请求方式: POST
//https://api.weixin.qq.com/cgi-bin/message/mass/delete?access_token=ACCESS_TOKEN
func (b *Broadcasting) Delete(msgID string) core.Responder {
	token := b.accessToken.GetToken()
	return core.PostJSON(Link(messageMassDelete), token.KeyMap(), util.Map{"msg_id": msgID})

}

//Status 查询群发消息发送状态【订阅号与服务号认证后均可用】
//接口调用请求说明
//http请求方式: POST
//https://api.weixin.qq.com/cgi-bin/message/mass/get?access_token=ACCESS_TOKEN
func (b *Broadcasting) Status(msgID string) core.Responder {
	token := b.accessToken.GetToken()
	return core.PostJSON(Link(messageMassGet), token.KeyMap(), util.Map{"msg_id": msgID})

}

func (b *Broadcasting) SendText() {

}
