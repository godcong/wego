package message

type PicList struct {
	PicMd5Sum CDATA `xml:"PicMd5Sum"`
}

type SendPicsInfo struct {
	Count   int       //发送的图片数量
	PicList []PicList `xml:"PicList>item"`
}
