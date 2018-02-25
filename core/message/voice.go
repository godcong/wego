package message

type Voice struct {
	MediaId     string //语音消息媒体id，可以调用多媒体文件下载接口拉取数据。
	Format      string //语音格式，如amr，speex等
	Recognition string //语音识别结果，UTF8编码
}
