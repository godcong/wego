package message

import (
	"encoding/xml"

	"github.com/godcong/wego/util"
)

type Text struct {
	*Message
	Content CDATA
}

func (t *Text) ToXml() ([]byte, error) {
	return xml.Marshal(*t)
}
func (t *Text) ToJson() ([]byte, error) {
	m := t.ToMap()
	return m.ToJson(), nil
}

func (t *Text) ToMap() util.Map {
	m := util.Map{
		"msgtype": t.MsgType.String(),
		"touser":  t.ToUserName.Value,
		t.MsgType.String(): util.Map{
			"content": t.Content.Value,
		},
	}
	return m
}

//NewText 初始化文本消息
func NewText(msg *Message, content string) *Text {
	return &Text{
		Message: msg,
		Content: CDATA{
			Value: content,
		},
	}

}
