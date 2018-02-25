package message

type Card struct {
	Title        string //标题
	AppId        string //小程序appid
	PagePath     string //小程序页面路径
	ThumbUrl     string //封面图片的临时cdn链接
	ThumbMediaId string //封面图片的临时素材id
}
