package message

import (
	"encoding/xml"

	"github.com/godcong/wego/util"
)

type Link struct {
	*Message
	Title       string //消息标题
	Url         string //消息链接
	ThumbUrl    string //消息链接
	Description string //消息描述
}

func NewLink(msg *Message, title, url, thumbUrl, desc string) *Link {
	return &Link{
		Message:     msg,
		Title:       title,
		Url:         url,
		ThumbUrl:    thumbUrl,
		Description: desc,
	}
}

func (l *Link) ToXml() ([]byte, error) {
	return xml.Marshal(*l)
}

func (l *Link) ToJson() ([]byte, error) {
	m := l.ToMap()
	return m.ToJSON(), nil
}

func (l *Link) ToMap() util.Map {
	m := util.Map{
		"msgtype": l.MsgType.String(),
		"touser":  l.ToUserName.Value,
		l.MsgType.String(): util.Map{
			"title":       l.Title,
			"description": l.Description,
			"url":         l.Url,
			"thumb_url":   l.ThumbUrl,
		},
	}
	return m
}
