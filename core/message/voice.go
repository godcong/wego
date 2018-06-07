package message

/*Voice Voice */
type Voice struct {
	Message
	MediaID     string `xml:"media_id"`    //语音消息媒体id，可以调用多媒体文件下载接口拉取数据。
	Format      string `xml:"format"`      //语音格式，如amr，speex等
	Recognition string `xml:"recognition"` //语音识别结果，UTF8编码
}
