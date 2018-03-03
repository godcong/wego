package message

type Image struct {
	Message
	PicUrl  string //图片链接（由系统生成）
	MediaId string //图片消息媒体id，可以调用多媒体文件下载接口拉取数据。
}
