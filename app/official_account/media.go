package official_account

import (
	"github.com/godcong/wego/config"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/net"
	"github.com/godcong/wego/util"
)

type Media struct {
	config config.Config
	*OfficialAccount
}

func newMedia(account *OfficialAccount) *Media {
	return &Media{
		config:          defaultConfig,
		OfficialAccount: account,
	}
}

func NewMedia() *Media {
	return newMedia(account)
}

// https://api.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=TYPE
// 参数	是否必须	说明
// access_token	是	调用接口凭证
// type	是	媒体文件类型，分别有图片（image）、语音（voice）、视频（video）和缩略图（thumb）
// media	是	form-data中媒体文件标识，有filename、filelength、content-type等信息
// 成功:
// {"type":"image","media_id":"w6fY9-444YS7Dmgt7_CaOApjbqBPyOSt-BbgQcbt0Pc_4t31u5JXQE8OGs6iqdqv","created_at":1521343152}
// {"type":"video","media_id":"9fCk1Any5VcwmbJPzGztWMq3a1PsWv11KpgLTdM_YXgIlwdAUosdeSI_M6M7Qtwb","created_at":1521346725}
// 失败:
// {"errcode":41005,"errmsg":"media data missing hint: [1HqFUa09681538]"}
func (m *Media) Upload(filePath string, mediaType core.MediaType) *net.Response {
	log.Debug("Media|Upload", filePath, mediaType)
	p := m.token.GetToken().KeyMap()
	p.Set("type", mediaType.String())
	resp := m.client.HttpUpload(
		m.client.Link(MEDIA_UPLOAD_URL_SUFFIX),
		util.Map{
			"media": filePath,
		},
		util.Map{
			net.REQUEST_TYPE_QUERY.String(): p,
		})
	return resp
}

func (m *Media) UploadThumb(filePath string) *net.Response {
	return m.Upload(filePath, core.MediaTypeThumb)
}

func (m *Media) UploadVoice(filePath string) *net.Response {
	return m.Upload(filePath, core.MediaTypeVoice)
}

func (m *Media) UploadVideo(filePath string) *net.Response {
	return m.Upload(filePath, core.MediaTypeVideo)
}

func (m *Media) UploadImage(filePath string) *net.Response {
	return m.Upload(filePath, core.MediaTypeImage)
}

// http请求方式: GET,https调用
// https://api.weixin.qq.com/cgi-bin/media/get?access_token=ACCESS_TOKEN&media_id=MEDIA_ID
// 请求示例（示例为通过curl命令获取多媒体文件）
// curl -I -G "https://api.weixin.qq.com/cgi-bin/media/get?access_token=ACCESS_TOKEN&media_id=MEDIA_ID"
// 成功:
// {"video_url":"http://203.205.158.73/vweixinp.tc.qq.com/1007_49fe0f8b21124a8a93339e23789356cc.f10.mp4?vkey=9966EBE6CA73990B37A5A8F05AB8FC9906A2A96CCA3D2F7730FFA56696A978B984C4DC5A7633D24F3A98A3C3CF91A2391CFBB0290410BC07DFDC84662BC2CD97256A6B988B0F56CDD95EAA617CE634B8E26ABAD5974025F4&sha=0&save=1"}
// 失败:
// {"errcode":40007,"errmsg":"invalid media_id"}
func (m *Media) Get(mediaId string) *net.Response {
	log.Debug("Media|Get", mediaId)
	p := m.token.GetToken().KeyMap()
	p.Set("media_id", mediaId)
	resp := m.client.HttpGet(
		m.client.Link(MEDIA_GET_URL_SUFFIX),
		util.Map{
			net.REQUEST_TYPE_QUERY.String(): p,
		})
	return resp
}

// http请求方式: GET,https调用
// https://api.weixin.qq.com/cgi-bin/media/get/jssdk?access_token=ACCESS_TOKEN&media_id=MEDIA_ID
// 请求示例（示例为通过curl命令获取多媒体文件）
// curl -I -G "https://api.weixin.qq.com/cgi-bin/media/get/jssdk?access_token=ACCESS_TOKEN&media_id=MEDIA_ID"
// 失败:
// {"errcode":40007,"errmsg":"invalid media_id"}
func (m *Media) GetJssdk(mediaId string) *net.Response {
	p := m.token.GetToken().KeyMap()
	p.Set("media_id", mediaId)
	resp := m.client.HttpGet(
		m.client.Link(MEDIA_GET_JSSDK_URL_SUFFIX),
		util.Map{
			net.REQUEST_TYPE_QUERY.String(): p,
		})
	return resp
}

// http请求方式: POST，https协议
// https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=ACCESS_TOKEN
// 调用示例（使用curl命令，用FORM表单方式上传一个图片）:
// curl -F media=@test.jpg "https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=ACCESS_TOKEN"
// 成功:
// {"url":"http:\/\/mmbiz.qpic.cn\/mmbiz_jpg\/gJHMd2C74XpfUBCTPocUe1Dd8cXnAlDmRqdPoFWq1DvJZjdW5BCaYyu7NfHusicU50nRs8Vb1oiaNrwMbTtNcFtQ\/0"}
func (m *Media) UploadMediaImg(filePath string) *net.Response {
	return m.uploadImg("media", filePath)
}

// HTTP请求方式: POST/FROMURL:https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=ACCESS_TOKEN
// 调用示例（使用curl命令，用FORM表单方式上传一个图片）：curl –Fbuffer=@test.jpg
func (m *Media) UploadBufferImg(filePath string) *net.Response {
	return m.uploadImg("buffer", filePath)
}

func (m *Media) uploadImg(name string, filePath string) *net.Response {
	p := m.token.GetToken().KeyMap()
	resp := m.client.HttpUpload(
		m.client.Link(MEDIA_UPLOADIMG_URL_SUFFIX),
		util.Map{
			name: filePath,
		},
		util.Map{
			net.REQUEST_TYPE_QUERY.String(): p,
		})
	return resp
}
