package message

import (
	"encoding/xml"

	"github.com/godcong/wego/util"
)

/*Image Image */
type Image struct {
	*Message
	PicURL  string `xml:"pic_url"`  //图片链接（由系统生成）
	MediaID string `xml:"media_id"` //图片消息媒体id，可以调用多媒体文件下载接口拉取数据。
}

//NewImage 初始化图像消息
func NewImage(msg *Message, picURL, mediaID string) *Image {
	return &Image{
		Message: msg,
		PicURL:  picURL,
		MediaID: mediaID,
	}
}

/*ToXML transfer image to xml */
func (t *Image) ToXML() ([]byte, error) {
	return xml.Marshal(*t)
}

/*ToJSON transfer image to json */
func (t *Image) ToJSON() ([]byte, error) {
	m := t.ToMap()
	return m.ToJSON(), nil
}

/*ToMap transfer image to map */
func (t *Image) ToMap() util.Map {
	m := util.Map{
		"msgtype": t.MsgType.String(),
		"touser":  t.ToUserName.Value,
		t.MsgType.String(): util.Map{
			"media_id": t.MediaID,
		},
	}
	if t.PicURL != "" {
		m.Set("picurl", t.PicURL)
	}

	return m
}
