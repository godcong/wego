package official

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/core/media"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
)

/*Material Material */
type Material struct {
	*Media
}

func newMaterial(media *Media) *Material {
	return &Material{
		Media: media,
	}
}

/*NewMaterial NewMaterial */
func NewMaterial(config *core.Config) *Material {
	return newMaterial(NewMedia(config))
}

//AddNews 新增永久素材
// http请求方式: POST，https协议
// https://api.weixin.qq.com/cgi-bin/material/add_news?access_token=ACCESS_TOKEN
//func (m *Material) AddNews(articles []*media.Article) core.Responder {
func (m *Material) AddNews(maps util.Map) core.Responder {
	log.Debug("Material|AddNews", maps)
	key := m.accessToken.GetToken().KeyMap()
	resp := core.PostJSON(
		Link(materialAddNewsURLSuffix),
		key,
		maps)
	return resp
}

//AddMaterial 新增其他类型永久素材
// http请求方式: POST，需使用https
// https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=ACCESS_TOKEN&type=TYPE
func (m *Material) AddMaterial(filePath string, mediaType core.MediaType) core.Responder {
	log.Debug("Material|AddMaterial", filePath, mediaType)
	if mediaType == core.MediaTypeVideo {
		log.Error("please use Material.UploadVideo() function")
	}

	p := m.accessToken.GetToken().KeyMap()
	p.Set("type", mediaType.String())
	resp := core.Upload(
		Link(materialAddMaterialURLSuffix),
		p,
		util.Map{
			"media": filePath,
		})
	return resp
}

//UploadVideo 新增其他类型永久素材
// http请求方式: POST，需使用https
// https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=ACCESS_TOKEN&type=TYPE
func (m *Material) UploadVideo(filePath string, title, introduction string) core.Responder {
	log.Debug("Media|UploadVideo", filePath, title, introduction)
	p := m.accessToken.GetToken().KeyMap()
	p.Set("type", core.MediaTypeVideo.String())
	resp := core.Upload(
		Link(materialAddMaterialURLSuffix),
		p,
		util.Map{
			"media": filePath,
			"description": util.Map{
				"title":        title,
				"introduction": introduction,
			},
		})
	return resp
}

//Get 获取永久素材
// http请求方式: POST,https协议
// https://api.weixin.qq.com/cgi-bin/material/get_material?access_token=ACCESS_TOKEN
func (m *Material) Get(mediaID string) core.Responder {
	log.Debug("Material|Get", mediaID)
	p := m.accessToken.GetToken().KeyMap()
	resp := core.PostJSON(
		Link(materialGetMaterialURLSuffix),
		p,
		util.Map{
			"media_id": mediaID,
		})
	return resp
}

//Del 删除永久素材
// http请求方式: POST
// https://api.weixin.qq.com/cgi-bin/material/del_material?access_token=ACCESS_TOKEN
func (m *Material) Del(mediaID string) core.Responder {
	log.Debug("Material|Del", mediaID)
	p := m.accessToken.GetToken().KeyMap()
	resp := core.PostJSON(
		Link(materialDelMaterialURLSuffix),
		p,
		util.Map{
			"media_id": mediaID,
		})
	return resp

}

//UpdateNews 修改永久图文素材
// http请求方式: POST
// https://api.weixin.qq.com/cgi-bin/material/update_news?access_token=ACCESS_TOKEN
func (m *Material) UpdateNews(mediaID string, index int, articles []*media.Article) core.Responder {
	log.Debug("Material|UpdateNews", mediaID)
	p := m.accessToken.GetToken().KeyMap()
	resp := core.PostJSON(
		Link(materialUpdateNewsURLSuffix),
		p,
		util.Map{
			"media_id": mediaID,
			"index":    index,
			"articles": articles,
		})
	return resp

}

//GetMaterialCount 获取素材总数
// http请求方式: GET
// https://api.weixin.qq.com/cgi-bin/material/get_materialcount?access_token=ACCESS_TOKEN
func (m *Material) GetCount() core.Responder {
	log.Debug("Material|GetMaterialCount")
	p := m.accessToken.GetToken().KeyMap()
	resp := core.Get(
		Link(materialGetMaterialcountURLSuffix),
		p)
	return resp
}

//BatchGet 获取素材列表
// http请求方式: POST
// https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token=ACCESS_TOKEN
//参数说明
//参数	是否必须	说明
//type	是	素材的类型，图片（image）、视频（video）、语音 （voice）、图文（news）
//offset	是	从全部素材的该偏移位置开始返回，0表示从第一个素材 返回
//count	是	返回素材的数量，取值在1到20之间
func (m *Material) BatchGet(mediaType core.MediaType, offset, count int) core.Responder {
	log.Debug("Material|BatchGet", mediaType, offset, count)
	p := m.accessToken.GetToken().KeyMap()
	resp := core.PostJSON(
		Link(materialBatchgetMaterialURLSuffix),
		p,
		util.Map{
			"type":   mediaType.String(),
			"offset": offset,
			"count":  count,
		})
	return resp

}
