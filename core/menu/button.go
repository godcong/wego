package menu

import (
	"github.com/godcong/wego/core/message"
	"github.com/godcong/wego/core/util"
)

type Button util.Map

//{

//	Type     string //是	菜单的响应动作类型，view表示网页类型，click表示点击类型，miniprogram表示小程序类型
//	Name     string //是	菜单标题，不超过16个字节，子菜单不超过60个字节
//	Key      string //click等点击类型必须	菜单KEY值，用于消息接口推送，不超过128字节
//	Url      string //view、miniprogram类型必须	网页 链接，用户点击菜单可打开链接，不超过1024字节。 type为miniprogram时，不支持小程序的老版本客户端将打开本url。
//	MediaId  string //media_id类型和view_limited类型必须	调用新增永久素材接口返回的合法media_id
//	AppId    string //miniprogram类型必须	小程序的appid（仅认证公众号可配置）
//	Pagepath string //string miniprogram类型必须	小程序的页面路径
//}

func NewClickButton(name, key string) *Button {
	return newButton(message.EventTypeClick.String(), util.Map{"name": name, "key": key})

}

func NewViewButton(name, url string) *Button {
	return newButton(message.EventTypeView.String(), util.Map{"name": name, "url": url})
}

func NewScanCodeWaitMsgButton(name, key string) *Button {
	return newButton(message.EventTypeScancodeWaitmsg.String(), util.Map{"name": name, "key": key})
}

func NewScanCodePushButton(name, key string) *Button {
	return newButton(message.EventTypeScancodePush.String(), util.Map{"name": name, "key": key})
}

func NewPicSysPhotoButton(name, key string) *Button {
	return newButton(message.EventTypePicSysphoto.String(), util.Map{"name": name, "key": key})
}

func NewPicPhotoOrAlbumButton(name, key string) *Button {
	return newButton(message.EventTypePicPhotoOrAlbum.String(), util.Map{"name": name, "key": key})
}

func NewPicWeixinButton(name, key string) *Button {
	return newButton(message.EventTypePicWeixin.String(), util.Map{"name": name, "key": key})
}

func NewMediaIDButton(name, mediaId string) *Button {
	return newButton("media_id", util.Map{"name": name, "media_id": mediaId})
}

func NewViewLimitedButton(name, mediaId string) *Button {
	return newButton("view_limited", util.Map{"name": name, "media_id": mediaId})
}

func NewMiniProgramButton(name, url, pagepath string) *Button {
	return newButton("miniprogram", util.Map{"name": name, "url": url, "pagepath": pagepath})
}

func NewSubButton(name string, sub []*Button) *Button {
	return newButton("", util.Map{"name": name, "key": "testkey", "sub_button": sub})
}

func newButton(typ string, val util.Map) *Button {
	v := make(Button)
	if typ != "" {
		(*util.Map)(&v).Set("type", typ)
	}
	(*util.Map)(&v).Join(val)
	return &v
}

func (b *Button) SetSub(name string, sub []*Button) *Button {
	*b = make(Button)
	(*util.Map)(b).Set("name", name)
	(*util.Map)(b).Set("sub_button", sub)
	return b
}
