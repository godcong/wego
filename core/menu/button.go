package menu

import (
	"github.com/godcong/wego/core/message"
	"github.com/godcong/wego/util"
)

/*Button Button */
type Button struct {
	util.Map
}

//{
//	Type     string //是	菜单的响应动作类型，view表示网页类型，click表示点击类型，miniprogram表示小程序类型
//	Name     string //是	菜单标题，不超过16个字节，子菜单不超过60个字节
//	Key      string //click等点击类型必须	菜单KEY值，用于消息接口推送，不超过128字节
//	URL      string //view、miniprogram类型必须	网页 链接，用户点击菜单可打开链接，不超过1024字节。 type为miniprogram时，不支持小程序的老版本客户端将打开本url。
//	MediaID  string //media_id类型和view_limited类型必须	调用新增永久素材接口返回的合法media_id
//	AppID    string //miniprogram类型必须	小程序的appid（仅认证公众号可配置）
//	Pagepath string //string miniprogram类型必须	小程序的页面路径
//}

/*NewClickButton NewClickButton*/
func NewClickButton(name, key string) *Button {
	return newButton(message.EventTypeClick, util.Map{"name": name, "key": key})

}

/*NewViewButton NewViewButton*/
func NewViewButton(name, url string) *Button {
	return newButton(message.EventTypeView, util.Map{"name": name, "url": url})
}

//func NewScanCodeWaitMsgButton(name, key string) *Button {
//	return newButton(typeToString(message.EventTypeScancodeWaitmsg), util.Map{"name": name, "key": key})
//}
//
//func NewScanCodePushButton(name, key string) *Button {
//	return newButton(typeToString(message.EventTypeScancodePush), util.Map{"name": name, "key": key})
//}
//
//func NewPicSysPhotoButton(name, key string) *Button {
//	return newButton(typeToString(message.EventTypePicSysphoto), util.Map{"name": name, "key": key})
//}
//
//func NewPicPhotoOrAlbumButton(name, key string) *Button {
//	return newButton(typeToString(message.EventTypePicPhotoOrAlbum), util.Map{"name": name, "key": key})
//}
//
//func NewPicWeixinButton(name, key string) *Button {
//	return newButton(typeToString(message.EventTypePicWeixin), util.Map{"name": name, "key": key})
//}
//
//func NewMediaIDButton(name, mediaId string) *Button {
//	return newButton("media_id", util.Map{"name": name, "media_id": mediaId})
//}
//
//func NewViewLimitedButton(name, mediaId string) *Button {
//	return newButton("view_limited", util.Map{"name": name, "media_id": mediaId})
//}
//
//func NewMiniProgramButton(name, url, pagepath string) *Button {
//	return newButton("miniprogram", util.Map{"name": name, "url": url, "pagepath": pagepath})
//}

/*NewSubButton NewSubButton*/
func NewSubButton(name string, sub []*Button) *Button {
	return newButton("", util.Map{"name": name, "key": "testkey", "sub_button": sub})
}

func newButton(typ message.EventType, val util.Map) *Button {
	button := NewBaseButton()
	if typ != "" {
		button.Set("type", typ)
	}
	button.Join(val)
	return button
}

/*SetSub SetSub*/
func (b *Button) SetSub(name string, sub []*Button) *Button {
	b.Map = util.Map{}
	b.Set("name", name)
	b.Set("sub_button", sub)
	return b
}

/*NewBaseButton NewBaseButton */
func NewBaseButton() *Button {
	return &Button{
		Map: make(util.Map),
	}
}

/*SetButtons SetButtons*/
func (b *Button) SetButtons(buttons []*Button) *Button {
	b.Set("button", buttons)
	return b
}

/*GetButtons GetButtons */
func (b *Button) GetButtons() []*Button {
	buttons := b.Get("button")
	if v0, b := buttons.([]*Button); b {
		return v0
	}
	return nil
}

/*GetMatchRule GetMatchRule*/
func (b *Button) GetMatchRule() *MatchRule {
	if mr := b.Get("matchrule"); mr != nil {
		return mr.(*MatchRule)
	}
	return nil
}

/*SetMatchRule SetMatchRule*/
func (b *Button) SetMatchRule(rule *MatchRule) *Button {
	b.Set("matchrule", rule)
	return b
}

func (b *Button) mapGet(name string) interface{} {
	return b.Get(name)
}

func (b *Button) mapSet(name string, v interface{}) *util.Map {
	return b.Set(name, v)
}

/*AddButton AddButton*/
func (b *Button) AddButton(buttons *Button) *Button {
	if v := b.GetButtons(); v != nil {
		b.SetButtons(append(v, buttons))
	} else {
		b.SetButtons([]*Button{buttons})
	}
	return b
}
