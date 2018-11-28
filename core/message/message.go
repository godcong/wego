package message

import (
	"encoding/xml"
	"strings"

	"github.com/godcong/wego/util"
)

/*Messager Messager */
type Messager interface {
	ToXML() ([]byte, error)
	ToJSON() ([]byte, error)
}

/*CDATA CDATA */
type CDATA = util.CDATA

/*MsgType MsgType */
type MsgType string

/*message types */
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
	MsgTypeAll             MsgType = "_msg_type_all_"
)

/*MSGCDATA MSGCDATA */
type MSGCDATA struct {
	MsgType `xml:",cdata"`
}

/*Message Message */
type Message struct {
	XMLName      xml.Name `xml:"xml"`
	MsgType      MSGCDATA `xml:"MsgType"`
	MsgID        int64    `xml:"MsgID,omitempty"`
	ToUserName   CDATA    `xml:"to_user_name"`
	FromUserName CDATA    `xml:"from_user_name"`
	CreateTime   int64    `xml:"create_time"`
}

/*String String */
func (e MsgType) String() string {
	return string(e)
}

/*Compare compare message type*/
func (e MsgType) Compare(msgType MsgType) int {
	return strings.Compare(strings.ToLower(e.String()), msgType.String())
}

/*Compare compare message type*/
func (e *Message) Compare(msgType MsgType) int {
	return e.MsgType.Compare(msgType)
}

/*New create a new message */
func New(msgType MsgType, toUser, fromUser string, msgID, createTime int64) *Message {
	return &Message{
		MsgType: MSGCDATA{
			MsgType: msgType,
		},
		MsgID: msgID,
		ToUserName: CDATA{
			Value: toUser,
		},
		FromUserName: CDATA{
			Value: fromUser,
		},
		CreateTime: createTime,
	}
}
