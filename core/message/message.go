package message

import (
	"encoding/xml"
	"strings"
)

type Messager interface {
	ToXml() ([]byte, error)
	ToJson() ([]byte, error)
}

type CDATA struct {
	Value string `xml:",cdata"`
}

type MsgType string

const (
	MsgTypeText            MsgType = "text"                      //表示文本消息	``
	MsgTypeImage           MsgType = "image"                     //表示图片消息
	MsgTypeVoice           MsgType = "voice"                     //表示语音消息
	MsgTypeVideo           MsgType = "video"                     //表示视频消息
	MsgTypeShortvideo      MsgType = "shortvideo"                //表示短视频消息[限接收]
	MsgTypeLocation        MsgType = "location"                  //表示坐标消息[限接收]
	MsgTypeLink            MsgType = "link"                      //表示链接消息[限接收]
	MsgTypeMusic           MsgType = "music"                     //表示音乐消息[限回复]
	MsgTypeNews            MsgType = "news"                      //表示图文消息[限回复]
	MsgTypeTransfer        MsgType = "transfer_customer_service" //表示消息消息转发到客服
	MsgTypeEvent           MsgType = "event"                     //表示事件推送消息
	MsgTypeMiniprogrampage MsgType = "miniprogrampage"
)

type MSGCDATA struct {
	MsgType `xml:",cdata"`
}

type Message struct {
	XMLName      xml.Name `xml:"xml"`
	MsgType      MSGCDATA `xml:"MsgType"`
	MsgId        int64    `xml:"MsgId,omitempty"`
	ToUserName   CDATA
	FromUserName CDATA
	CreateTime   int64
}

func (e MsgType) String() string {
	return string(e)
}

func (e MsgType) Compare(msgType MsgType) int {
	return strings.Compare(strings.ToLower(e.String()), msgType.String())
}

func (e *Message) Compare(msgType MsgType) int {
	return e.MsgType.Compare(msgType)
}
