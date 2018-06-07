package message

import (
	"encoding/xml"

	"github.com/godcong/wego/util"
)

/*Link Link */
type Link struct {
	*Message
	Title       string //消息标题
	URL         string //消息链接
	ThumbURL    string //消息链接
	Description string //消息描述
}

/*NewLink NewLink */
func NewLink(msg *Message, title, url, thumbURL, desc string) *Link {
	return &Link{
		Message:     msg,
		Title:       title,
		URL:         url,
		ThumbURL:    thumbURL,
		Description: desc,
	}
}

/*ToXML transfer link to xml */
func (l *Link) ToXML() ([]byte, error) {
	return xml.Marshal(*l)
}

/*ToJSON transfer link to json */
func (l *Link) ToJSON() ([]byte, error) {
	m := l.ToMap()
	return m.ToJSON(), nil
}

/*ToMap transfer link to map */
func (l *Link) ToMap() util.Map {
	m := util.Map{
		"msgtype": l.MsgType.String(),
		"touser":  l.ToUserName.Value,
		l.MsgType.String(): util.Map{
			"title":       l.Title,
			"description": l.Description,
			"url":         l.URL,
			"thumb_url":   l.ThumbURL,
		},
	}
	return m
}
