package message

import (
	"encoding/xml"

	"github.com/godcong/wego/util"
)

type Image struct {
	*Message
	PicUrl  string //图片链接（由系统生成）
	MediaId string //图片消息媒体id，可以调用多媒体文件下载接口拉取数据。
}

//NewImage 初始化图像消息
func NewImage(msg *Message, picUrl, mediaId string) *Image {
	return &Image{
		Message: msg,
		PicUrl:  picUrl,
		MediaId: mediaId,
	}
}

func (t *Image) ToXml() ([]byte, error) {
	return xml.Marshal(*t)
}
func (t *Image) ToJson() ([]byte, error) {
	m := t.ToMap()
	return m.ToJSON(), nil
}

func (t *Image) ToMap() util.Map {
	m := util.Map{
		"msgtype": t.MsgType.String(),
		"touser":  t.ToUserName.Value,
		t.MsgType.String(): util.Map{
			"media_id": t.MediaId,
		},
	}
	if t.PicUrl != "" {
		m.Set("picurl", t.PicUrl)
	}

	return m
}
