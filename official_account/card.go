package official_account

import (
	"strings"

	"github.com/godcong/wego/core"
)

type Card struct {
	core.Config
	*OfficialAccount
}

func newCard(officialAccount *OfficialAccount) *Card {
	return &Card{
		Config:          defaultConfig,
		OfficialAccount: officialAccount,
	}
}

func NewCard() *Card {
	return newCard(account)
}

type CardScene string

const (
	CARD_SCENE_NEAR_BY          CardScene = "SCENE_NEAR_BY"          // 附近
	CARD_SCENE_MENU             CardScene = "SCENE_MENU"             // 自定义菜单
	CARD_SCENE_QRCODE           CardScene = "SCENE_QRCODE"           // 二维码
	CARD_SCENE_ARTICLE          CardScene = "SCENE_ARTICLE"          // 公众号文章
	CARD_SCENE_H5               CardScene = "SCENE_H5"               // h5页面
	CARD_SCENE_IVR              CardScene = "SCENE_IVR"              // 自动回复
	CARD_SCENE_CARD_CUSTOM_CELL CardScene = "SCENE_CARD_CUSTOM_CELL" // 卡券自定义cell
)

func (s CardScene) String() string {
	return string(s)
}

type CardList struct {
	CardId   string `json:"card_id"`   // card_id	所要在页面投放的card_id	是
	ThumbUrl string `json:"thumb_url"` // thumb_url	缩略图url	是
}

type CardLandingPage struct {
	Banner   string     `json:"banner"`     //页面的banner图片链接，须调用，建议尺寸为640*300。	是
	Title    string     `json:"page_title"` //页面的title。	是
	CanShare bool       `json:"can_share"`  //页面是否可以分享,填入true/false	是
	Scene    CardScene  `json:"scene"`      //	投放页面的场景值； SCENE_NEAR_BY 附近 SCENE_MENU 自定义菜单 SCENE_QRCODE 二维码 SCENE_ARTICLE 公众号文章 SCENE_H5 h5页面 SCENE_IVR 自动回复 SCENE_CARD_CUSTOM_CELL 卡券自定义cell	是
	CardList []CardList `json:"card_list"`  // card_list	卡券列表，每个item有两个字段	是
}

type CardType string

func (t CardType) String() string {
	return string(t)
}

const (
	//card_type	是	string(24)	GROUPON	团购券类型。
	//base_info	是	JSON结构	见上述示例。	基本的卡券数据 ，见下表，所有卡券类型通用。
	//deal_detail	是	string( 3072 )	双人套餐\n -进口红酒一支。\n孜然牛肉一份。	团购券专用，团购详情。
	CARD_TYPE_GROUPON = "GROUPON"

	//card_type	是	string(24)	CASH	代金券类型。
	//base_info	是	JSON结构	见上述示例。	基本的卡券数据，见下表，所有卡券通用。
	//least_cost	是	int	10000	代金券专用，表示起用金额（单位为分）,如果无起用门槛则填0。
	//reduce_cost	是	int	10000	代金券专用，表示减免金额。（单位为分）
	CARD_TYPE_CASH = "CASH"

	//card_type	是	string(24)	DISCOUNT	折扣券类型。
	//base_info	是	Json结构	见上述示例。	基本的卡券数据 ，见下表，所有卡券通用。
	//discount	是	int	30	折扣券专用，表示打折额度（百分比）。填30就是七折。
	CARD_TYPE_DISCOUNT = "DISCOUNT"

	//card_type	是	string(24)	GIFT	兑换券类型。
	//base_info	是	Json结构	见上述示例。	基本的卡券数据 ，所有卡券通用。
	//gift	是	string(3072)	可兑换音乐木盒一个。	兑换券专用，填写兑换内容的名称。
	CARD_TYPE_GIFT = "GIFT"

	//card_type	是	string(24)	GENERAL_COUPON	优惠券类型。
	//base_info	是	Json结构	见上述示例。	基本的卡券数据 ，见下表，所有卡券通用。
	//default_detail	是	string(3072)	音乐木盒。	优惠券专用，填写优惠详情。
	CARD_TYPE_GENERAL_COUPON = "GENERAL_COUPON"
)

type CardDataInfo struct {
	Type           string `json:"type"`             //	type	是	string	DATE_TYPE_FIX _TIME_RANGE 表示固定日期区间，DATETYPE FIX_TERM 表示固定时长 （自领取后按天算。	使用时间的类型，旧文档采用的1和2依然生效。
	BeginTimestamp int64  `json:"begin_timestamp"`  //	begin_time stamp	是	unsigned int	14300000	type为DATE_TYPE_FIX_TIME_RANGE时专用，表示起用时间。从1970年1月1日00:00:00至起用时间的秒数，最终需转换为字符串形态传入。（东八区时间,UTC+8，单位为秒）
	EndTimestamp   int64  `json:"end_timestamp"`    //	end_time stamp	是	unsigned int	15300000	表示结束时间 ， 建议设置为截止日期的23:59:59过期 。 （ 东八区时间,UTC+8，单位为秒 ）
	FixedTerm      int    `json:"fixed_term"`       // fixed_term	是	int	15	type为DATE_TYPE_FIX_TERM时专用，表示自领取后多少天内有效，不支持填写0。
	FixedBeginTerm int    `json:"fixed_begin_term"` // fixed_begin_term	是	int	0	type为DATE_TYPE_FIX_TERM时专用，表示自领取后多少天开始生效，领取后当天生效填写0。（单位为天）
}

type CardSku struct {
	Quantity int `json:"quantity"` // quantity	是	int	100000	卡券库存的数量，上限为100000000。
}

type CardCodeType string

func (t CardCodeType) String() string {
	return string(t)
}

const (
	CARD_CODE_TYPE_TEXT         CardCodeType = "CODE_TYPE_TEXT"         //文 本
	CARD_CODE_TYPE_BARCODE      CardCodeType = "CODE_TYPE_BARCODE"      //一维码
	CARD_CODE_TYPE_QRCODE       CardCodeType = "CODE_TYPE_QRCODE"       //二维码
	CARD_CODE_TYPE_ONLY_QRCODE  CardCodeType = "CODE_TYPE_ONLY_QRCODE"  //二维码无code显示
	CARD_CODE_TYPE_ONLY_BARCODE CardCodeType = "CODE_TYPE_ONLY_BARCODE" //一维码无code显示
	CARD_CODE_TYPE_NONE         CardCodeType = "CODE_TYPE_NONE"         //不显示code和条形码类型
)

type CardBaseInfo struct {
	LogoUrl                   string       `json:"logo_url"`                                //	logo_url	是	strin g(128)	http://mmbiz.qpic.cn/	卡券的商户logo，建议像素为300*300。
	BrandName                 string       `json:"brand_name"`                              //	brand_name	是	string（36）	海底捞	商户名字,字数上限为12个汉字。
	CodeType                  CardCodeType `json:"code_type"`                               //	code_type	是	string(16)	CODE_TYPE_TEXT	码型： "CODE_TYPE_TEXT"文 本 ； "CODE_TYPE_BARCODE"一维码 "CODE_TYPE_QRCODE"二维码 "CODE_TYPE_ONLY_QRCODE",二维码无code显示； "CODE_TYPE_ONLY_BARCODE",一维码无code显示；CODE_TYPE_NONE， 不显示code和条形码类型
	Title                     string       `json:"title"`                                   //	title	是	string（27）	双人套餐100元兑换券	卡券名，字数上限为9个汉字。(建议涵盖卡券属性、服务及金额)。
	Color                     string       `json:"color"`                                   //	color	是	string（16）	Color010	券颜色。按色彩规范标注填写Color010-Color100。
	Notice                    string       `json:"notice"`                                  //	notice	是	string（48）	请出示二维码	卡券使用提醒，字数上限为16个汉字。
	ServicePhone              string       `json:"service_phone,omitempty"`                 //	service_phone	否	string（24）	40012234	客服电话。
	Description               string       `json:"description"`                             //	description	是	strin g （3072）	不可与其他优惠同享	卡券使用说明，字数上限为1024个汉字。
	DateInfo                  CardDataInfo `json:"date_info"`                               //	date_info	是	JSON结构	见上述示例。	使用日期，有效期的信息。
	Sku                       CardSku      `json:"sku"`                                     //	sku	是	JSON结构	见上述示例。	商品信息。
	UseLimit                  int          `json:"use_limit,omitempty"`                     //	use_limit否int100每人可核销的数量限制,不填写默认为50。
	GetLimit                  int          `json:"get_limit,omitempty"`                     //	get_limit	否	int	1	每人可领券的数量限制,不填写默认为50。
	UseCustomCode             bool         `json:"use_custom_code,omitempty"`               //	use_custom_code	否	bool	true	是否自定义Code码 。填写true或false，默认为false。 通常自有优惠码系统的开发者选择 自定义Code码，并在卡券投放时带入 Code码，详情见 是否自定义Code码 。
	GetCustomCodeMode         string       `json:"get_custom_code_mode,omitempty"`          // 	get_custom_code_mode	否	string(32)	GET_CUSTOM_COD E_MODE_DEPOSIT	填入 GET_CUSTOM_CODE_MODE_DEPOSIT 表示该卡券为预存code模式卡券， 须导入超过库存数目的自定义code后方可投放， 填入该字段后，quantity字段须为0,须导入code 后再增加库存
	BindOpenid                bool         `json:"bind_openid"`                             //	bind_openid	否	bool	true	是否指定用户领取，填写true或false 。默认为false。通常指定特殊用户群体 投放卡券或防止刷券时选择指定用户领取。
	CanShare                  bool         `json:"can_share,omitempty"`                     //	can_share	否	bool	false	卡券领取页面是否可分享。
	CanGiveFriend             bool         `json:"can_give_friend,omitempty"`               //	can_give_friend否boolfalse卡券是否可转赠。
	LocationIdList            []int        `json:"location_id_list,omitempty"`              //	location_id_list	否	array	1234，2312	门店位置poiid。 调用 POI门店管理接 口 获取门店位置poiid。具备线下门店 的商户为必填。
	UseAllLocations           bool         `json:"use_all_locations,omitempty"`             //    use_all_locations	否	bool	true	设置本卡券支持全部门店，与location_id_list互斥
	CenterTitle               string       `json:"center_title,omitempty"`                  //	center_title	否	string（18）	立即使用	卡券顶部居中的按钮，仅在卡券状 态正常(可以核销)时显示
	CenterSubTitle            string       `json:"center_sub_title,omitempty"`              //	center_sub_title	否	string（24）	立即享受优惠	显示在入口下方的提示语 ，仅在卡券状态正常(可以核销)时显示。
	CenterUrl                 string       `json:"center_url,omitempty"`                    //	center_url	否	string（128）	www.qq.com	顶部居中的url ，仅在卡券状态正常(可以核销)时显示。
	CenterAppBrandUserName    string       `json:"center_app_brand_user_name,omitempty"`    //    center_app_brand_user_name	否	string（128）	gh_86a091e50ad4@app	卡券跳转的小程序的user_name，仅可跳转该 公众号绑定的小程序 。
	CenterAppBrandPass        string       `json:"center_app_brand_pass,omitempty"`         //    center_app_brand_pass	否	string（128）	API/cardPage	卡券跳转的小程序的path
	CustomUrlName             string       `json:"custom_url_name,omitempty"`               //	custom_url_name	否	string（15）	立即使用	自定义跳转外链的入口名字。
	CustomUrl                 string       `json:"custom_url,omitempty"`                    //	custom_url	否	string（128）	www.qq.com	自定义跳转的URL。
	CustomUrlSubTitle         string       `json:"custom_url_sub_title,omitempty"`          //	custom_url_sub_title	否	string（18）	更多惊喜	显示在入口右侧的提示语。
	CustomAppBrandUserName    string       `json:"custom_app_brand_user_name,omitempty"`    //    custom_app_brand_user_name	否	string（128）	gh_86a091e50ad4@app	卡券跳转的小程序的user_name，仅可跳转该 公众号绑定的小程序 。
	CustomAppBrandPass        string       `json:"custom_app_brand_pass,omitempty"`         //    custom _app_brand_pass否string（128）API/cardPage卡券跳转的小程序的path
	PromotionUrlName          string       `json:"promotion_url_name,omitempty"`            //	promotion_url_name	否	string（15）	产品介绍	营销场景的自定义入口名称。
	PromotionUrl              string       `json:"promotion_url,omitempty"`                 //	promotion_url	否	string（128）	www.qq.com	入口跳转外链的地址链接。
	PromotionUrlSubTitle      string       `json:"promotion_url_sub_title,omitempty"`       //promotion_url_sub_title	否	string（18）	卖场大优惠。	显示在营销入口右侧的提示语。
	PromotionAppBrandUserName string       `json:"promotion_app_brand_user_name,omitempty"` //promotion_app_brand_user_name	否	string（128）	gh_86a091e50ad4@app	卡券跳转的小程序的user_name，仅可跳转该 公众号绑定的小程序 。
	PromotionAppBrandPass     string       `json:"promotion_app_brand_pass,omitempty"`      //promotion_app_brand_pass	否	string（128）	API/cardPage	卡券跳转的小程序的path
	Source                    string       `json:"source"`                                  //	"source": "大众点评"
}

type CardUseCondition struct {
	AcceptCategory          string `json:"accept_category,omitempty"`             //	accept_category	否	string（512）	指定可用的商品类目，仅用于代金券类型 ，填入后将在券面拼写适用于xxx
	RejectCategory          string `json:"reject_category,omitempty"`             //	reject_category	否	string（ 512 ）	指定不可用的商品类目，仅用于代金券类型 ，填入后将在券面拼写不适用于xxxx
	LeastCost               int    `json:"least_cost,omitempty"`                  //least_cost	否	int	满减门槛字段，可用于兑换券和代金券 ，填入后将在全面拼写消费满xx元可用。
	ObjectUseFor            string `json:"object_use_for,omitempty"`              //object_use_for	否	string（ 512 ）	购买xx可用类型门槛，仅用于兑换 ，填入后自动拼写购买xxx可用。
	CanUseWithOtherDiscount bool   `json:"can_use_with_other_discount,omitempty"` //	can_use_with_other_discount	否	bool	不可以与其他类型共享门槛 ，填写false时系统将在使用须知里 拼写“不可与其他优惠共享”， 填写true时系统将在使用须知里 拼写“可与其他优惠共享”， 默认为true
}

type CardAbstract struct {
	Abstract    string   `json:"abstract,omitempty"`      //	abstract	否	string（24 ）	封面摘要简介。
	IconUrlList []string `json:"icon_url_list,omitempty"` //	icon_url_list	否	string（128 ）	封面图片列表，仅支持填入一 个封面图片链接， 上传图片接口 上传获取图片获得链接，填写 非CDN链接会报错，并在此填入。 建议图片尺寸像素850*350
}

type CardTextImageList struct {
	ImageUrl string `json:"image_url,omitempty"` //	image_url	否	string（128 ）	图片链接，必须调用 上传图片接口 上传图片获得链接，并在此填入， 否则报错
	Text     string `json:"text,omitempty"`      //	text	否	string（512 ）	图文描述
}

type CardTimeLimit struct {
	Type        string `json:"type,omitempty"`         //	type	否	string（24 ）	限制类型枚举值：支持填入 MONDAY 周一 TUESDAY 周二 WEDNESDAY 周三 THURSDAY 周四 FRIDAY 周五 SATURDAY 周六 SUNDAY 周日 此处只控制显示， 不控制实际使用逻辑，不填默认不显示
	BeginHour   int    `json:"begin_hour,omitempty"`   //	begin_hour	否	int	当前type类型下的起始时间（小时） ，如当前结构体内填写了MONDAY， 此处填写了10，则此处表示周一 10:00可用
	EndHour     int    `json:"end_hour,omitempty"`     //	end_hour	否	int	当前type类型下的结束时间（小时） ，如当前结构体内填写了MONDAY， 此处填写了20， 则此处表示周一 10:00-20:00可用
	BeginMinute int    `json:"begin_minute,omitempty"` //	begin_minute	否	int	当前type类型下的起始时间（分钟） ，如当前结构体内填写了MONDAY， begin_hour填写10，此处填写了59， 则此处表示周一 10:59可用
	EndMinute   int    `json:"end_minute,omitempty"`   //	end_minute	否	int	当前type类型下的结束时间（分钟） ，如当前结构体内填写了MONDAY， begin_hour填写10，此处填写了59， 则此处表示周一 10:59-00:59可用
}

type CardAdvancedInfo struct {
	UseCondition    *CardUseCondition   `json:"use_condition,omitempty"`    //	use_condition	否	JSON结构	使用门槛（条件）字段，若不填写使用条件则在券面拼写 ：无最低消费限制，全场通用，不限品类；并在使用说明显示： 可与其他优惠共享
	Abstract        *CardAbstract       `json:"abstract,omitempty"`         //	abstract	否	JSON结构	封面摘要结构体名称
	TextImageList   []CardTextImageList `json:"text_image_list,omitempty"`  //  text_image_list	否	JSON结构	图文列表，显示在详情内页 ，优惠券券开发者须至少传入 一组图文列表
	TimeLimit       []CardTimeLimit     `json:"time_limit,omitempty"`       //	time_limit否JSON结构使用时段限制，包含以下字段
	BusinessService []string            `json:"business_service,omitempty"` //	business_service	否	array	商家服务类型： BIZ_SERVICE_DELIVER 外卖服务； BIZ_SERVICE_FREE_PARK 停车位； BIZ_SERVICE_WITH_PET 可带宠物； BIZ_SERVICE_FREE_WIFI 免费wifi， 可多选
}

type OneCard struct {
	CardType CardType `json:"card_type"`
	data     core.Map
}

// HTTP请求方式: POST
// URL:https://api.weixin.qq.com/card/landingpage/create?access_token=$TOKEN
// 失败:
// {"errcode":48001,"errmsg":"api unauthorized hint: [9JA27a06332890]"}
// 成功:
// {"errcode":0,"errmsg":"ok","url":"https:\/\/open.weixin.qq.com\/connect\/oauth2\/authorize?appid=wx3be6367203f983ac&redirect_uri=https%3A%2F%2Fmp.weixin.qq.com%2Fbizmall%2Fcardlandingpage%3Fbiz%3DMzAwNjkzNzgzMA%3D%3D%26page_id%3D1%26scene%3D1&response_type=code&scope=snsapi_base#wechat_redirect","page_id":1}
func (c *Card) CreateLandingPage(page *CardLandingPage) *core.Response {
	key := c.token.GetToken().KeyMap()
	resp := c.client.HttpPost(
		c.client.Link(CARD_LANDINGPAGE_CREATE_URL_SUFFIX),
		core.Map{
			core.REQUEST_TYPE_QUERY.String(): key,
			core.REQUEST_TYPE_JSON.String():  page,
		})
	return resp
}

// HTTP请求方式: POST
// URL:http://api.weixin.qq.com/card/code/deposit?access_token=ACCESS_TOKEN
// 失败:
// {"errcode":48001,"errmsg":"api unauthorized hint: [gM8quA01752947]"}
func (c *Card) Deposit(cid string, code []string) *core.Response {
	key := c.token.GetToken().KeyMap()
	resp := c.client.HttpPost(
		c.client.Link(CARD_CODE_DEPOSIT_URL_SUFFIX),
		core.Map{
			core.REQUEST_TYPE_QUERY.String(): key,
			core.REQUEST_TYPE_JSON.String(): core.Map{
				"card_id": cid,
				"code":    code,
			},
		})
	return resp
}

// HTTP请求方式: POST
// URL:http://api.weixin.qq.com/card/code/getdepositcount?access_token=ACCESS_TOKEN
// 失败:
// {"errcode":48001,"errmsg":"api unauthorized hint: [TQ2Iga07222944]"}
func (c *Card) GetDepositCount(cid string) *core.Response {
	key := c.token.GetToken().KeyMap()
	resp := c.client.HttpPost(
		c.client.Link(CARD_CODE_GETDEPOSITCOUNT_URL_SUFFIX),
		core.Map{
			core.REQUEST_TYPE_QUERY.String(): key,
			core.REQUEST_TYPE_JSON.String(): core.Map{
				"card_id": cid,
			},
		})
	return resp
}

// HTTP请求方式: POST
// URL:http://api.weixin.qq.com/card/code/checkcode?access_token=ACCESS_TOKEN
// 失败:
// {"errcode":48001,"errmsg":"api unauthorized hint: [xxvWsa08782921]"}
func (c *Card) CheckCode(cid string, code []string) *core.Response {
	key := c.token.GetToken().KeyMap()
	resp := c.client.HttpPost(
		c.client.Link(CARD_CODE_CHECKCODE_URL_SUFFIX),
		core.Map{
			core.REQUEST_TYPE_QUERY.String(): key,
			core.REQUEST_TYPE_JSON.String(): core.Map{
				"card_id": cid,
				"code":    code,
			},
		})
	return resp
}

//HTTP请求方式: POST
//URL:https://api.weixin.qq.com/card/code/get?access_token=TOKEN
//code	是	string(20)	110201201245	单张卡券的唯一标准。
//card_id	否	string(32)	pFS7Fjg8kV1I dDz01r4SQwMkuCKc	卡券ID代表一类卡券。自定义code卡券必填。
//check_consume	否	bool	true	是否校验code核销状态，填入true和false时的code异常状态返回数据不同。
func (c *Card) GetCode(p core.Map) *core.Response {
	key := c.token.GetToken().KeyMap()
	resp := c.client.HttpPost(
		c.client.Link(CARD_CODE_GET_URL_SUFFIX),
		core.Map{
			core.REQUEST_TYPE_QUERY.String(): key,
			core.REQUEST_TYPE_JSON.String():  p,
		})
	return resp
}

// HTTP请求方式: POST
// URL:https://api.weixin.qq.com/card/mpnews/gethtml?access_token=TOKEN
func (c *Card) GetHtml(cid string) *core.Response {
	key := c.token.GetToken().KeyMap()
	resp := c.client.HttpPost(
		c.client.Link(CARD_MPNEWS_GETHTML_URL_SUFFIX),
		core.Map{
			core.REQUEST_TYPE_QUERY.String(): key,
			core.REQUEST_TYPE_JSON.String(): core.Map{
				"card_id": cid,
			},
		})
	return resp
}

// HTTP请求方式: POST
// URL:https://api.weixin.qq.com/card/testwhitelist/set?access_token=TOKEN
func (c *Card) SetTestWhiteListById(list []string) *core.Response {
	return c.SetTestWhiteList("openid", list)
}

// HTTP请求方式: POST
// URL:https://api.weixin.qq.com/card/testwhitelist/set?access_token=TOKEN
func (c *Card) SetTestWhiteListByName(list []string) *core.Response {
	return c.SetTestWhiteList("username", list)
}

// HTTP请求方式: POST
// URL:https://api.weixin.qq.com/card/testwhitelist/set?access_token=TOKEN
// 成功:
// {"errcode":0,"errmsg":"ok","white_list_size":1,"success_openid":["o5jo6s3RZ6rxuVAW33IpTjYWQOg4"],"success_username":[]}
func (c *Card) SetTestWhiteList(typ string, list []string) *core.Response {
	key := c.token.GetToken().KeyMap()
	resp := c.client.HttpPost(
		c.client.Link(CARD_TESTWHITELIST_SET_URL_SUFFIX),
		core.Map{
			core.REQUEST_TYPE_QUERY.String(): key,
			core.REQUEST_TYPE_JSON.String(): core.Map{
				typ: list,
			},
		})
	return resp
}

//HTTP请求方式: POST
//URL:https://api.weixin.qq.com/card/qrcode/create?access_token=TOKEN
//成功:
//{"errcode":0,"errmsg":"ok","ticket":"gQHz8DwAAAAAAAAAAS5odHRwOi8vd2VpeGluLnFxLmNvbS9xLzAydWNTSGNqQ1c4V1AxXzJ2ME5xNGcAAgS62MBaAwQIBwAA","expire_seconds":1800,"url":"http:\/\/weixin.qq.com\/q\/02ucSHcjCW8WP1_2v0Nq4g","show_qrcode_url":"https:\/\/mp.weixin.qq.com\/cgi-bin\/showqrcode?ticket=gQHz8DwAAAAAAAAAAS5odHRwOi8vd2VpeGluLnFxLmNvbS9xLzAydWNTSGNqQ1c4V1AxXzJ2ME5xNGcAAgS62MBaAwQIBwAA"}
func (c *Card) CreateQrCode(action *QrCodeAction) *core.Response {
	key := c.token.GetToken().KeyMap()
	resp := c.client.HttpPost(
		c.client.Link(CARD_QRCODE_CREATE_URL_SUFFIX),
		core.Map{
			core.REQUEST_TYPE_QUERY.String(): key,
			core.REQUEST_TYPE_JSON.String():  action,
		})
	return resp
}

// HTTP请求方式: POST
// URL: https://api.weixin.qq.com/card/create?access_token=ACCESS_TOKEN
// 失败:
// {"errcode":48001,"errmsg":"api unauthorized hint: [wk0hOa0525e193]"}
// 成功:
// {"errcode":0,"errmsg":"ok","card_id":"p5jo6s-f42donkk9mmZ0VILQXOME"}
func (c *Card) Create(card *OneCard) *core.Response {
	key := c.token.GetToken().KeyMap()
	_, d := card.Get()
	resp := c.client.HttpPostJson(
		c.client.Link(CARD_CREATE_URL_SUFFIX),
		core.Map{"card": d},
		core.Map{
			core.REQUEST_TYPE_QUERY.String(): key,
		})
	return resp
}

func NewOneCard(cardType CardType, data core.Map) *OneCard {
	ct := strings.ToLower(cardType.String())
	return &OneCard{
		CardType: cardType,
		data: core.Map{
			"card_type": cardType,
			ct:          data,
		},
	}
}
func (c *OneCard) AddAdvancedInfo(info *CardAdvancedInfo) *OneCard {
	return c.add("advanced_info", info)
}

func (c *OneCard) AddBaseInfo(info *CardBaseInfo) *OneCard {
	return c.add("base_info", info)
}

func (c *OneCard) AddDealDetail(d string) *OneCard {
	return c.add("deal_detail", d)
}

func (c *OneCard) add(name string, info interface{}) *OneCard {
	ct := strings.ToLower(c.CardType.String())
	if c.data != nil {
		if v, b := c.data[ct].(core.Map); b {
			if v != nil {
				v[name] = info
			} else {
				v = core.Map{
					name: info,
				}
			}
			c.data[ct] = v
		}
	} else {
		c.data = core.Map{
			"card_type": c.CardType,
			ct: core.Map{
				name: info,
			},
		}
	}
	return c
}

func (c *OneCard) Set(cardType CardType, data core.Map) {
	ct := strings.ToLower(cardType.String())
	c.CardType = cardType
	c.data = core.Map{
		"card_type": ct,
		ct:          data,
	}
}

func (c *OneCard) Get() (CardType, core.Map) {
	return c.CardType, c.data
}
