package official

import (
	"github.com/godcong/wego/core"
	//"github.com/godcong/wego/config"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
)

/*PoiPhotoURL PoiPhotoURL*/
type PoiPhotoURL struct {
	PhotoURL string `json:"photo_url"`
}

/*PoiBaseInfo PoiBaseInfo*/
type PoiBaseInfo struct {
	Poi          string        `json:"poi,omitempty"`          // "poi_id ":"271864249"
	Sid          string        `json:"sid,omitempty"`          // "sid":"33788392",
	BusinessName string        `json:"business_name"`          //"business_name":"15个汉字或30个英文字符内",
	BranchName   string        `json:"branch_name"`            //"branch_name":"不超过10个字，不能含有括号和特殊字符",
	Province     string        `json:"province"`               //"province":"不超过10个字",
	City         string        `json:"city"`                   //"city":"不超过30个字",
	District     string        `json:"district"`               //"district":"不超过10个字",
	Address      string        `json:"address"`                //"address":"门店所在的详细街道地址（不要填写省市信息）:不超过80个字",
	Telephone    string        `json:"telephone"`              //"telephone":"不超53个字符（不可以出现文字）",
	Categories   []string      `json:"categories"`             //"categories":["美食,小吃快餐"],
	OffsetType   int           `json:"offset_type"`            //"offset_type":1,
	Longitude    float64       `json:"longitude"`              //"longitude":115.32375,
	Latitude     float64       `json:"latitude"`               //"latitude":25.097486,
	PhotoList    []PoiPhotoURL `json:"photo_list,omitempty"`   //"photo_list":[{"photo_url":"https:// 不超过20张.com"}，{"photo_url":"https://XXX.com"}],
	Recommend    string        `json:"recommend,omitempty"`    //"recommend":"不超过200字。麦辣鸡腿堡套餐，麦乐鸡，全家桶",
	Special      string        `json:"special,omitempty"`      //"special":"不超过200字。免费wifi，外卖服务",
	Introduction string        `json:"introduction,omitempty"` //"introduction":"不超过300字。麦当劳是全球大型跨国连锁餐厅，1940 年创立于美国，在世界上大约拥有3 万间分店。	主要售卖汉堡包，以及薯条、炸鸡、汽水、冰品、沙拉、 水果等快餐食品",
	OpenTime     string        `json:"open_time,omitempty"`    //"open_time":"8:00-20:00",
	AvgPrice     int           `json:"avg_price,omitempty"`    //"avg_price":35
}

/*Poi Poi */
type Poi struct {
	//config Config
	*Account
}

func newPoi(account *Account) *Poi {
	return &Poi{
		//config:  defaultConfig,
		Account: account,
	}
}

/*NewPoi NewPoi */
func NewPoi(config *core.Config) *Poi {
	return newPoi(NewOfficialAccount(config))
}

/*
Add 创建门店
	http请求方式	POST/FORM
	请求Url	https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=ACCESS_TOKEN
	POST数据格式	buffer
*/
func (p *Poi) Add(biz *PoiBaseInfo) core.Response {
	log.Debug("Poi|Add", *biz)
	//p.client.SetDomain(core.NewDomain("mp"))
	// base64.URLEncoding.EncodeToString([]byte(ticket))
	resp := p.client.PostJSON(
		Link(poiAddPoi),
		p.accessToken.GetToken().KeyMap(),
		util.Map{
			"business": biz,
		})
	return resp
}

/*
Get 查询门店信息
http请求方式	POST
请求Url	http://api.weixin.qq.com/cgi-bin/poi/getpoi?access_token=TOKEN
POST数据格式	json
*/
func (p *Poi) Get(poiID string) core.Response {
	log.Debug("Poi|Get", poiID)
	//p.client.SetDomain(core.NewDomain("mp"))
	// base64.URLEncoding.EncodeToString([]byte(ticket))
	resp := p.client.PostJSON(
		Link(poiGetPoi),
		p.accessToken.GetToken().KeyMap(),
		util.Map{
			"poi_id": poiID,
		})
	return resp

}

/*
Update 修改门店服务信息
http请求方式	POST/FROM
请求Url	https://api.weixin.qq.com/cgi-bin/poi/updatepoi?access_token=TOKEN
POST数据格式	buffer
字段说明:
全部字段内容同前。
特别注意:
以上8个字段，若有填写内容则为覆盖更新，若无内容则视为不修改，维持原有内容。 photo_list 字段为全列表覆盖，若需要增加图片，需将之前图片同样放入list 中，在其后增加新增图片。如:已有A、B、C 三张图片，又要增加D、E 两张图，则需要调用该接口，photo_list 传入A、B、C、D、E 五张图片的链接。
成功返回:
{
"errcode":0,
"errmsg":"ok"
}
*/
func (p *Poi) Update(biz *PoiBaseInfo) core.Response {
	log.Debug("Poi|Update", *biz)
	//p.client.SetDomain(core.NewDomain("mp"))
	// base64.URLEncoding.EncodeToString([]byte(ticket))
	resp := p.client.PostJSON(
		Link(poiUpdatePoi),
		p.accessToken.GetToken().KeyMap(),
		util.Map{
			"business": biz,
		})
	return resp
}

/*
GetList 查询门店列表
	http请求方式	POST
	请求Url	https://api.weixin.qq.com/cgi-bin/poi/getpoilist?access_token=TOKEN
	POST数据格式	json
	字段	说明	是否必填
	begin	开始位置，0 即为从第一条开始查询	是
	limit	返回数据条数，最大允许50，默认为20	是
成功返回:
{
    "errcode":0,
    "errmsg":"ok"
    "business_list":[
    {"base_info":{
    "sid":"101",
    "business_name":"麦当劳",
    "branch_name":"艺苑路店",
    "address":"艺苑路11号",
    "telephone":"020-12345678",
    "categories":["美食,快餐小吃"],
    "city":"广州市",
    "province":"广东省",
    "offset_type":1,
    "longitude":115.32375,
    "latitude":25.097486,
    "photo_list":[{"photo_url":"http: ...."}],
    "introduction":"麦当劳是全球大型跨国连锁餐厅，1940 年创立于美国，在世界上大约拥有3 万间分店。主要售卖汉堡包，以及薯条、炸鸡、汽水、冰品、沙拉、水果等快餐食品",
    "recommend":"麦辣鸡腿堡套餐，麦乐鸡，全家桶",
    "special":"免费wifi，外卖服务",
    "open_time":"8:00-20:00",
    "avg_price":35,
    "poi_id":"285633617",
    "available_state":3,
    "district":"海珠区",
    "update_status":0
}},
{"base_info":{
    "sid":"101",
    "business_name":"麦当劳",
    "branch_name":"北京路店",
    "address":"北京路12号",
    "telephone":"020-12345689",
    "categories":["美食,快餐小吃"],
    "city":"广州市",
    "province":"广东省",
    "offset_type":1,
    "longitude":115.3235,
    "latitude":25.092386,
    "photo_list":[{"photo_url":"http: ...."}],
    "introduction":"麦当劳是全球大型跨国连锁餐厅，1940 年创立于美国，在世界上大约拥有3 万间分店。主要售卖汉堡包，以及薯条、炸鸡、汽水、冰品、沙拉、水果等快餐食品",
    "recommend":"麦辣鸡腿堡套餐，麦乐鸡，全家桶",
    "special":"免费wifi，外卖服务",
    "open_time":"8:00-20:00",
    "avg_price":35,
    "poi_id":"285633618",
    "available_state":4,
    "district":"越秀区",
    "update_status":0
}},
{"base_info":{
    "sid":"101",
    "business_name":"麦当劳",
    "branch_name":"龙洞店",
    "address":"迎龙路122号",
    "telephone":"020-12345659",
    "categories":["美食,快餐小吃"],
    "city":"广州市",
    "province":"广东省",
    "offset_type":1,
    "longitude":115.32345,
    "latitude":25.056686,
    "photo_list":[{"photo_url":"http: ...."}],
    "introduction":"麦当劳是全球大型跨国连锁餐厅，1940 年创立于美国，在世界上大约拥有3 万间分店。主要售卖汉堡包，以及薯条、炸鸡、汽水、冰品、沙拉、水果等快餐食品",
    "recommend":"麦辣鸡腿堡套餐，麦乐鸡，全家桶",
    "special":"免费wifi，外卖服务",
    "open_time":"8:00-20:00",
    "avg_price":35,
    "poi_id":"285633619",
    "available_state":2,
    "district":"天河区",
    "update_status":0
}},
],
"total_count":"3",
}
失败返回:
*/
func (p *Poi) GetList(begin int, limit int) core.Response {
	log.Debug("Poi|GetList", begin, limit)
	//p.client.SetDomain(core.NewDomain("mp"))
	// base64.URLEncoding.EncodeToString([]byte(ticket))
	resp := p.client.PostJSON(
		Link(poiGetListPoi),
		p.accessToken.GetToken().KeyMap(),
		util.Map{
			"begin": begin,
			"limit": limit,
		})
	return resp
}

/*
Del 删除门店
协议	https
http请求方式	POST/FROM
请求Url	https://api.weixin.qq.com/cgi-bin/poi/delpoi?access_token=TOKEN
POST数据格式	buffer

*/
func (p *Poi) Del(poiID string) core.Response {
	log.Debug("Poi|Del", poiID)
	//p.client.SetDomain(core.NewDomain("mp"))
	// base64.URLEncoding.EncodeToString([]byte(ticket))
	resp := p.client.PostJSON(
		Link(poiDelPoi),
		p.accessToken.GetToken().KeyMap(),
		util.Map{
			"poi_id": poiID,
		})
	return resp
}

/*
GetCategory 门店类目表
http请求方式	GET
请求Url	http://api.weixin.qq.com/cgi-bin/poi/getwxcategory?access_token=TOKEN
成功返回:
{
"category_list":
["美食,江浙菜,上海菜","美食,江浙菜,淮扬菜","美食,江浙菜,浙江菜","美食,江浙菜,南京菜 ","美食,江浙菜,苏帮菜…"]
}
*/
func (p *Poi) GetCategory() core.Response {
	log.Debug("Poi|GetCategory")
	//p.client.SetDomain(core.NewDomain("mp"))
	// base64.URLEncoding.EncodeToString([]byte(ticket))
	resp := p.client.Get(
		Link(poiGetWXCategory),
		p.accessToken.GetToken().KeyMap())
	return resp
}
