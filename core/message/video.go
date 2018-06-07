package message

/*Video Video */
type Video struct {
	Message
	MediaID      string `xml:"media_id"`       //视频消息媒体id，可以调用多媒体文件下载接口拉取数据。
	ThumbMediaID string `xml:"thumb_media_id"` //	视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据。
	//MediaID      string
	//Title        string
	//Description  string
	//ThumbMediaID string
}
