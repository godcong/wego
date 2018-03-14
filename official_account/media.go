package official_account

import "io"

type Media struct {
}

//https://api.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=TYPE
//参数	是否必须	说明
//access_token	是	调用接口凭证
//type	是	媒体文件类型，分别有图片（image）、语音（voice）、视频（video）和缩略图（thumb）
//media	是	form-data中媒体文件标识，有filename、filelength、content-type等信息
func Upload(reader io.Reader, typ string) {

}
