package message

/*PicList PicList */
type PicList struct {
	PicMd5Sum CDATA `xml:"PicMd5Sum"`
}

/*SendPicsInfo SendPicsInfo */
type SendPicsInfo struct {
	Count   int       `xml:"count"` //发送的图片数量
	PicList []PicList `xml:"PicList>item"`
}
