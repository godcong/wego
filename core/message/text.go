package message

import (
	"encoding/json"
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
	return json.Marshal(*t)
}

func (t *Text) ToMap() util.Map {
	//TODO
	txt := util.Map{}
	//	"msgtype":t.MsgType,
	//	"touser":t.ToUserName.Value,
	//	FromUserName CDATA
	//	CreateTime   int64
	//
	//}

	return txt
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
