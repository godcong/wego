package message

/*Card Card */
type Card struct {
	Title        string `xml:"title"`          //标题
	AppID        string `xml:"app_id"`         //小程序appid
	PagePath     string `xml:"page_path"`      //小程序页面路径
	ThumbURL     string `xml:"thumb_url"`      //封面图片的临时cdn链接
	ThumbMediaID string `xml:"thumb_media_id"` //封面图片的临时素材id
}
