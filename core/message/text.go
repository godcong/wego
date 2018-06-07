package message

import (
	"encoding/xml"

	"github.com/godcong/wego/util"
)

/*Text Text */
type Text struct {
	*Message
	Content CDATA
}

/*ToXML transfer text to xml */
func (t *Text) ToXML() ([]byte, error) {
	return xml.Marshal(*t)
}

/*ToJSON transfer text to json */
func (t *Text) ToJSON() ([]byte, error) {
	m := t.ToMap()
	return m.ToJSON(), nil
}

/*ToMap transfer text to map */
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
