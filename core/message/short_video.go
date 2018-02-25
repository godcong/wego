package message

type ShortVideo struct {
	MediaId      string //视频消息媒体id，可以调用多媒体文件下载接口拉取数据。
	ThumbMediaId string //视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据。
}
