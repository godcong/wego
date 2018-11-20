package official

import (
	"encoding/json"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
	"strings"
)

// Card ...
type Card struct {
	*Account
	jssdk *JSSDK
}

func newCard(officialAccount *Account) *Card {
	return &Card{
		Account: officialAccount,
		jssdk:   NewJSSDK(officialAccount.Config),
	}
}

// NewCard ...
func NewCard(config *core.Config) *Card {
	return newCard(NewOfficialAccount(config))
}

// CardScene ...
type CardScene string

//CardSceneNearBy 附近
const CardSceneNearBy CardScene = "SCENE_NEAR_BY"

//CardSceneMenu 自定义菜单
const CardSceneMenu CardScene = "SCENE_MENU"

//CardSceneQrcode 二维码
const CardSceneQrcode CardScene = "SCENE_QRCODE"

//CardSceneArticle 公众号文章
const CardSceneArticle CardScene = "SCENE_ARTICLE"

//CardSceneH5 H5页面
const CardSceneH5 CardScene = "SCENE_H5"

//CardSceneIvr 自动回复
const CardSceneIvr CardScene = "SCENE_IVR"

//CardSceneCardCustomCell 卡券自定义cell
const CardSceneCardCustomCell CardScene = "SCENE_CARD_CUSTOM_CELL"

//CardStatus 支持开发者拉出指定状态的卡券列表
type CardStatus string

// CardStatusNotVerify ...
const (
	CardStatusNotVerify  CardStatus = "CARD_STATUS_NOT_VERIFY"  //待审核
	CardStatusVerifyFail CardStatus = "CARD_STATUS_VERIFY_FAIL" //审核失败
	CardStatusVerifyOk   CardStatus = "CARD_STATUS_VERIFY_OK"   //通过审核
	CardStatusDelete     CardStatus = "CARD_STATUS_DELETE"      //卡券被商户删除
	CardStatusDispatch   CardStatus = "CARD_STATUS_DISPATCH"    //在公众平台投放过的卡券；
)

// String ...
func (s CardScene) String() string {
	return string(s)
}

// CardList ...
type CardList struct {
	CardID   string `json:"card_id"`   // card_id	所要在页面投放的card_id	是
	ThumbURL string `json:"thumb_url"` // thumb_url	缩略图url	是
}

// CardLandingPage ...
type CardLandingPage struct {
	Banner   string     `json:"banner"`     //页面的banner图片链接，须调用，建议尺寸为640*300。	是
	Title    string     `json:"page_title"` //页面的title。	是
	CanShare bool       `json:"can_share"`  //页面是否可以分享,填入true/false	是
	Scene    CardScene  `json:"scene"`      //	投放页面的场景值； SCENE_NEAR_BY 附近 SCENE_MENU 自定义菜单 SCENE_QRCODE 二维码 SCENE_ARTICLE 公众号文章 SCENE_H5 h5页面 SCENE_IVR 自动回复 SCENE_CARD_CUSTOM_CELL 卡券自定义cell	是
	CardList []CardList `json:"card_list"`  // card_list	卡券列表，每个item有两个字段	是
}

// CardType ...
type CardType string

// String ...
func (t CardType) String() string {
	return string(t)
}

// CardTypeGroupon ...
const (
	//CardTypeGroupon GROUPON 团购券类型。
	CardTypeGroupon = "GROUPON"
	//CardTypeCash CASH	代金券类型。
	CardTypeCash = "CASH"
	//CardTypeDiscount DISCOUNT	折扣券类型。
	CardTypeDiscount = "DISCOUNT"
	//CardTypeGift GIFT 兑换券类型。
	CardTypeGift = "GIFT"
	//CardTypeGeneralCoupon GENERAL_COUPON 优惠券类型。
	CardTypeGeneralCoupon = "GENERAL_COUPON"
)

// CardDataInfo ...
type CardDataInfo struct {
	Type           string `json:"type"`             //	type	是	string	DATE_TYPE_FIX _TIME_RANGE 表示固定日期区间，DATETYPE FIX_TERM 表示固定时长 （自领取后按天算。	使用时间的类型，旧文档采用的1和2依然生效。
	BeginTimestamp int64  `json:"begin_timestamp"`  //	begin_time stamp	是	unsigned int	14300000	type为DATE_TYPE_FIX_TIME_RANGE时专用，表示起用时间。从1970年1月1日00:00:00至起用时间的秒数，最终需转换为字符串形态传入。（东八区时间,UTC+8，单位为秒）
	EndTimestamp   int64  `json:"end_timestamp"`    //	end_time stamp	是	unsigned int	15300000	表示结束时间 ， 建议设置为截止日期的23:59:59过期 。 （ 东八区时间,UTC+8，单位为秒 ）
	FixedTerm      int    `json:"fixed_term"`       //  fixed_term	是	int	15	type为DATE_TYPE_FIX_TERM时专用，表示自领取后多少天内有效，不支持填写0。
	FixedBeginTerm int    `json:"fixed_begin_term"` //  fixed_begin_term	是	int	0	type为DATE_TYPE_FIX_TERM时专用，表示自领取后多少天开始生效，领取后当天生效填写0。（单位为天）
}

// CardSku ...
type CardSku struct {
	Quantity int `json:"quantity"` // quantity	是	int	100000	卡券库存的数量，上限为100000000。
}

// CardCodeType ...
type CardCodeType string

// String ...
func (t CardCodeType) String() string {
	return string(t)
}

// CardCodeTypeText ...
const (
	//CardCodeTypeText 文 本
	CardCodeTypeText CardCodeType = "CODE_TYPE_TEXT"
	//CardCodeTypeBarcode 一维码
	CardCodeTypeBarcode CardCodeType = "CODE_TYPE_BARCODE"
	//CardCodeTypeQrcode 二维码
	CardCodeTypeQrcode CardCodeType = "CODE_TYPE_QRCODE"
	//CardCodeTypeOnlyQrcode 二维码无code显示
	CardCodeTypeOnlyQrcode CardCodeType = "CODE_TYPE_ONLY_QRCODE"
	//CardCodeTypeOnlyBarcode 一维码无code显示
	CardCodeTypeOnlyBarcode CardCodeType = "CODE_TYPE_ONLY_BARCODE"
	//CardCodeTypeNone 不显示code和条形码类型
	CardCodeTypeNone CardCodeType = "CODE_TYPE_NONE"
)

// CardBaseInfo ...
type CardBaseInfo struct {
	LogoURL                   string       `json:"logo_url"`                                //	logo_url	是	strin g(128)	http://mmbiz.qpic.cn/	卡券的商户logo，建议像素为300*300。
	BrandName                 string       `json:"brand_name"`                              //	brand_name	是	string（36）	海底捞	商户名字,字数上限为12个汉字。
	CodeType                  CardCodeType `json:"code_type"`                               //	code_type	是	string(16)	CODE_TYPE_TEXT	码型: "CODE_TYPE_TEXT"文 本 ； "CODE_TYPE_BARCODE"一维码 "CODE_TYPE_QRCODE"二维码 "CODE_TYPE_ONLY_QRCODE",二维码无code显示； "CODE_TYPE_ONLY_BARCODE",一维码无code显示；CODE_TYPE_NONE， 不显示code和条形码类型
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
	LocationIDList            []int        `json:"location_id_list,omitempty"`              //	location_id_list	否	array	1234，2312	门店位置poiid。 调用 POI门店管理接 口 获取门店位置poiid。具备线下门店 的商户为必填。
	UseAllLocations           bool         `json:"use_all_locations,omitempty"`             //  use_all_locations	否	bool	true	设置本卡券支持全部门店，与location_id_list互斥
	CenterTitle               string       `json:"center_title,omitempty"`                  //	center_title	否	string（18）	立即使用	卡券顶部居中的按钮，仅在卡券状 态正常(可以核销)时显示
	CenterSubTitle            string       `json:"center_sub_title,omitempty"`              //	center_sub_title	否	string（24）	立即享受优惠	显示在入口下方的提示语 ，仅在卡券状态正常(可以核销)时显示。
	CenterURL                 string       `json:"center_url,omitempty"`                    //	center_url	否	string（128）	www.qq.com	顶部居中的url ，仅在卡券状态正常(可以核销)时显示。
	CenterAppBrandUserName    string       `json:"center_app_brand_user_name,omitempty"`    //  center_app_brand_user_name	否	string（128）	gh_86a091e50ad4@app	卡券跳转的小程序的user_name，仅可跳转该 公众号绑定的小程序 。
	CenterAppBrandPass        string       `json:"center_app_brand_pass,omitempty"`         //  center_app_brand_pass	否	string（128）	API/cardPage	卡券跳转的小程序的path
	CustomURLName             string       `json:"custom_url_name,omitempty"`               //	custom_url_name	否	string（15）	立即使用	自定义跳转外链的入口名字。
	CustomURL                 string       `json:"custom_url,omitempty"`                    //	custom_url	否	string（128）	www.qq.com	自定义跳转的URL。
	CustomURLSubTitle         string       `json:"custom_url_sub_title,omitempty"`          //	custom_url_sub_title	否	string（18）	更多惊喜	显示在入口右侧的提示语。
	CustomAppBrandUserName    string       `json:"custom_app_brand_user_name,omitempty"`    //  custom_app_brand_user_name	否	string（128）	gh_86a091e50ad4@app	卡券跳转的小程序的user_name，仅可跳转该 公众号绑定的小程序 。
	CustomAppBrandPass        string       `json:"custom_app_brand_pass,omitempty"`         //  custom _app_brand_pass否string（128）API/cardPage卡券跳转的小程序的path
	PromotionURLName          string       `json:"promotion_url_name,omitempty"`            //	promotion_url_name	否	string（15）	产品介绍	营销场景的自定义入口名称。
	PromotionURL              string       `json:"promotion_url,omitempty"`                 //	promotion_url	否	string（128）	www.qq.com	入口跳转外链的地址链接。
	PromotionURLSubTitle      string       `json:"promotion_url_sub_title,omitempty"`       //  promotion_url_sub_title	否	string（18）	卖场大优惠。	显示在营销入口右侧的提示语。
	PromotionAppBrandUserName string       `json:"promotion_app_brand_user_name,omitempty"` //  promotion_app_brand_user_name	否	string（128）	gh_86a091e50ad4@app	卡券跳转的小程序的user_name，仅可跳转该 公众号绑定的小程序 。
	PromotionAppBrandPass     string       `json:"promotion_app_brand_pass,omitempty"`      //  promotion_app_brand_pass	否	string（128）	API/cardPage	卡券跳转的小程序的path
	Source                    string       `json:"source"`                                  //	"source": "大众点评"
}

// CardUseCondition ...
type CardUseCondition struct {
	AcceptCategory          string `json:"accept_category,omitempty"`             //	accept_category	否	string（512）	指定可用的商品类目，仅用于代金券类型 ，填入后将在券面拼写适用于xxx
	RejectCategory          string `json:"reject_category,omitempty"`             //	reject_category	否	string（ 512 ）	指定不可用的商品类目，仅用于代金券类型 ，填入后将在券面拼写不适用于xxxx
	LeastCost               int    `json:"least_cost,omitempty"`                  //least_cost	否	int	满减门槛字段，可用于兑换券和代金券 ，填入后将在全面拼写消费满xx元可用。
	ObjectUseFor            string `json:"object_use_for,omitempty"`              //object_use_for	否	string（ 512 ）	购买xx可用类型门槛，仅用于兑换 ，填入后自动拼写购买xxx可用。
	CanUseWithOtherDiscount bool   `json:"can_use_with_other_discount,omitempty"` //	can_use_with_other_discount	否	bool	不可以与其他类型共享门槛 ，填写false时系统将在使用须知里 拼写“不可与其他优惠共享”， 填写true时系统将在使用须知里 拼写“可与其他优惠共享”， 默认为true
}

// CardAbstract ...
type CardAbstract struct {
	Abstract    string   `json:"abstract,omitempty"`      //	abstract	否	string（24 ）	封面摘要简介。
	IconURLList []string `json:"icon_url_list,omitempty"` //	icon_url_list	否	string（128 ）	封面图片列表，仅支持填入一 个封面图片链接， 上传图片接口 上传获取图片获得链接，填写 非CDN链接会报错，并在此填入。 建议图片尺寸像素850*350
}

// CardTextImageList ...
type CardTextImageList struct {
	ImageURL string `json:"image_url,omitempty"` //	image_url	否	string（128 ）	图片链接，必须调用 上传图片接口 上传图片获得链接，并在此填入， 否则报错
	Text     string `json:"text,omitempty"`      //	text	否	string（512 ）	图文描述
}

// CardTimeLimit ...
type CardTimeLimit struct {
	Type        string `json:"type,omitempty"`         //	type	否	string（24 ）	限制类型枚举值:支持填入 MONDAY 周一 TUESDAY 周二 WEDNESDAY 周三 THURSDAY 周四 FRIDAY 周五 SATURDAY 周六 SUNDAY 周日 此处只控制显示， 不控制实际使用逻辑，不填默认不显示
	BeginHour   int    `json:"begin_hour,omitempty"`   //	begin_hour	否	int	当前type类型下的起始时间（小时） ，如当前结构体内填写了MONDAY， 此处填写了10，则此处表示周一 10:00可用
	EndHour     int    `json:"end_hour,omitempty"`     //	end_hour	否	int	当前type类型下的结束时间（小时） ，如当前结构体内填写了MONDAY， 此处填写了20， 则此处表示周一 10:00-20:00可用
	BeginMinute int    `json:"begin_minute,omitempty"` //	begin_minute	否	int	当前type类型下的起始时间（分钟） ，如当前结构体内填写了MONDAY， begin_hour填写10，此处填写了59， 则此处表示周一 10:59可用
	EndMinute   int    `json:"end_minute,omitempty"`   //	end_minute	否	int	当前type类型下的结束时间（分钟） ，如当前结构体内填写了MONDAY， begin_hour填写10，此处填写了59， 则此处表示周一 10:59-00:59可用
}

// CardAdvancedInfo ...
type CardAdvancedInfo struct {
	UseCondition    *CardUseCondition   `json:"use_condition,omitempty"`    //	use_condition	否	JSON结构	使用门槛（条件）字段，若不填写使用条件则在券面拼写 :无最低消费限制，全场通用，不限品类；并在使用说明显示: 可与其他优惠共享
	Abstract        *CardAbstract       `json:"abstract,omitempty"`         //	abstract	否	JSON结构	封面摘要结构体名称
	TextImageList   []CardTextImageList `json:"text_image_list,omitempty"`  //  text_image_list	否	JSON结构	图文列表，显示在详情内页 ，优惠券券开发者须至少传入 一组图文列表
	TimeLimit       []CardTimeLimit     `json:"time_limit,omitempty"`       //	time_limit否JSON结构使用时段限制，包含以下字段
	BusinessService []string            `json:"business_service,omitempty"` //	business_service	否	array	商家服务类型: BIZ_SERVICE_DELIVER 外卖服务； BIZ_SERVICE_FREE_PARK 停车位； BIZ_SERVICE_WITH_PET 可带宠物； BIZ_SERVICE_FREE_WIFI 免费wifi， 可多选
}

// OneCard ...
type OneCard struct {
	CardType CardType `json:"card_type"`
	data     util.Map
}

//CreateLandingPage 创建货架接口
//	HTTP请求方式: POST
//	URL:https://api.weixin.qq.com/card/landingpage/create?access_token=$TOKEN
//	func (c *Card) CreateLandingPage(page *CardLandingPage) core.Responder {
func (c *Card) CreateLandingPage(maps util.Map) core.Responder {
	resp := core.PostJSON(
		Link(cardLandingPageCreate),
		c.accessToken.GetToken().KeyMap(),
		maps,
	)
	return resp
}

//Deposit 导入code接口
//	HTTP请求方式: POST
//	URL:http://api.weixin.qq.com/card/code/deposit?access_token=ACCESS_TOKEN
func (c *Card) Deposit(cardID string, code []string) core.Responder {
	resp := core.PostJSON(
		Link(cardCodeDeposit),
		c.accessToken.GetToken().KeyMap(),
		util.Map{
			"card_id": cardID,
			"code":    code,
		},
	)
	return resp
}

//GetDepositCount 查询导入code数目
//
//  HTTP请求方式: POST
//  URL:http://api.weixin.qq.com/card/code/getdepositcount?access_token=ACCESS_TOKEN
func (c *Card) GetDepositCount(cardID string) core.Responder {
	resp := core.PostJSON(
		Link(cardCodeGetDepositCount),
		c.accessToken.GetToken().KeyMap(),
		util.Map{
			"card_id": cardID,
		},
	)
	return resp
}

//CheckCode 核查code接口
//	HTTP请求方式: POST
//	HTTP调用:http://api.weixin.qq.com/card/code/checkcode?access_token=ACCESS_TOKEN
func (c *Card) CheckCode(cardID string, code []string) core.Responder {
	resp := core.PostJSON(
		Link(cardCodeCheckCode),
		c.accessToken.GetToken().KeyMap(),
		util.Map{
			"card_id": cardID,
			"code":    code,
		},
	)
	return resp
}

//GetCode 查询Code接口
//	HTTP请求方式: POST
//	HTTP调用:https://api.weixin.qq.com/card/code/get?access_token=TOKEN
//	参数说明:
//	参数名	必填	类型	示例值	描述
//	code	是	string(20)	110201201245	单张卡券的唯一标准。
//	card_id	否	string(32)	pFS7Fjg8kV1I dDz01r4SQwMkuCKc	卡券ID代表一类卡券。自定义code卡券必填。
//	check_consume	否	bool	true	是否校验code核销状态，填入true和false时的code异常状态返回数据不同。
func (c *Card) GetCode(maps util.Map) core.Responder {
	resp := core.PostJSON(
		Link(cardCodeGet),
		c.accessToken.GetToken().KeyMap(),
		maps,
	)
	return resp
}

//GetHTML 图文消息群发卡券
//	HTTP请求方式: POST
//	URL:https://api.weixin.qq.com/card/mpnews/gethtml?access_token=TOKEN
func (c *Card) GetHTML(cardID string) core.Responder {
	resp := core.PostJSON(
		Link(cardMPNewsGetHTML),
		c.accessToken.GetToken().KeyMap(),
		util.Map{
			"card_id": cardID,
		},
	)
	return resp
}

//SetTestWhiteListByID 设置测试白名单(by openid)
func (c *Card) SetTestWhiteListByID(list []string) core.Responder {
	return c.SetTestWhiteList("openid", list)
}

//SetTestWhiteListByName 设置测试白名单(by username)
func (c *Card) SetTestWhiteListByName(list []string) core.Responder {
	return c.SetTestWhiteList("username", list)
}

//SetTestWhiteList 设置测试白名单
//	HTTP请求方式: POST
//	URL:https://api.weixin.qq.com/card/testwhitelist/set?access_token=TOKEN
func (c *Card) SetTestWhiteList(typ string, list []string) core.Responder {
	resp := core.PostJSON(
		Link(cardTestWhiteListSet),
		c.accessToken.GetToken().KeyMap(),
		util.Map{
			typ: list,
		},
	)
	return resp
}

//CreateQrCode 创建二维码
//	HTTP请求方式: POST
//	URL:https://api.weixin.qq.com/card/qrcode/create?access_token=TOKEN
func (c *Card) CreateQrCode(action *QrCodeAction) core.Responder {
	resp := core.PostJSON(
		Link(cardQrcodeCreate),
		c.accessToken.GetToken().KeyMap(),
		action,
	)
	return resp
}

//Create 创建卡券
//	HTTP请求方式: POST
//	URL: https://api.weixin.qq.com/card/create?access_token=ACCESS_TOKEN
//	func (c *Card) Create(card *OneCard) core.Responder {
func (c *Card) Create(maps util.Map) core.Responder {
	key := c.accessToken.GetToken().KeyMap()
	//_, d := maps.Get()
	resp := core.PostJSON(
		Link(cardCreate),
		key,
		util.Map{"card": maps})

	return resp
}

//Get 查看卡券详情
//	开发者可以调用该接口查询某个card_id的创建信息、审核状态以及库存数量。
//	接口调用请求说明
//	HTTP请求方式: POSTURL:https://api.weixin.qq.com/card/get?access_token=TOKEN
func (c *Card) Get(cardID string) core.Responder {
	token := c.accessToken.GetToken()
	return core.PostJSON("card/get", token.KeyMap(), util.Map{"card_id": cardID})
}

//GetApplyProtocol 卡券开放类目查询接口
//	HTTP请求方式: GET
//	URL:https://api.weixin.qq.com/card/getapplyprotocol?access_token=TOKEN
func (c *Card) GetApplyProtocol() core.Responder {
	token := c.accessToken.GetToken()
	return core.Get(Link(cardGetApplyProtocol), token.KeyMap())
}

//GetColors 卡券开放类目查询接口
//	HTTP请求方式: GET
//	URL:https://api.weixin.qq.com/card/getcolors?access_token=TOKEN
func (c *Card) GetColors() core.Responder {
	token := c.accessToken.GetToken()
	return core.Get(Link(cardGetColors), token.KeyMap())
}

//Checkin 更新飞机票信息接口
//	接口调用请求说明
//	http请求方式: POST
//	URL:https://api.weixin.qq.com/card/boardingpass/checkin?access_token=TOKEN
func (c *Card) Checkin(p util.Map) core.Responder {
	token := c.accessToken.GetToken()
	return core.PostJSON(Link(cardBoardingpassCheckin), token.KeyMap(), p)
}

//Categories 卡券开放类目查询接口
//	接口说明
//	通过调用该接口查询卡券开放的类目ID，类目会随业务发展变更，请每次用接口去查询获取实时卡券类目。
//	注意：
//	1.本接口查询的返回值还有卡券资质ID,此处的卡券资质为：已微信认证的公众号通过微信公众平台申请卡券功能时，所需的资质。
//	2.对于第三方强授权模式，子商户无论选择什么类目，均提交营业执照即可，所以不用考虑此处返回的资质字段，返回值仅参考类目ID即可。
//	接口详情
//	接口调用请求说明
//	https请求方式: GET https://api.weixin.qq.com/card/getapplyprotocol?access_token=TOKEN
func (c *Card) Categories() core.Responder {
	token := c.accessToken.GetToken()
	return core.Get(Link(cardGetapplyprotocol), token.KeyMap())

}

//BatchGet 批量查询卡券列表
//	接口调用请求说明
//	HTTP请求方式: POST URL:https://api.weixin.qq.com/card/batchget?access_token=TOKEN
func (c *Card) BatchGet(offset, count int, statusList []CardStatus) core.Responder {
	token := c.accessToken.GetToken()
	maps := util.Map{
		"offset":      offset,
		"count":       count,
		"status_list": statusList,
	}
	return core.PostJSON(Link(cardBatchget), token.KeyMap(), maps)
}

//Update 更改卡券信息接口
//	接口说明
//	支持更新所有卡券类型的部分通用字段及特殊卡券（会员卡、飞机票、电影票、会议门票）中特定字段的信息。
//	接口调用请求说明
//	HTTP请求方式: POST URL:https://api.weixin.qq.com/card/update?access_token=TOKEN
func (c *Card) Update(cardID string, p util.Map) core.Responder {
	token := c.accessToken.GetToken()
	maps := util.Map{
		"card_id": cardID,
	}
	maps.Join(p)
	return core.PostJSON(Link(cardUpdate), token.KeyMap(), maps)
}

// UpdateCode ...
func (c *Card) UpdateCode(code string, newCode string, cardID string) core.Responder {
	token := c.accessToken.GetToken()
	maps := util.Map{
		"code":     code,
		"new_code": newCode,
		"card_id":  cardID,
	}
	return core.PostJSON(Link(cardCodeUpdate), token.KeyMap(), maps)

}

// CodeUnavailable ...
func (c *Card) CodeUnavailable(code string, cardID string) core.Responder {
	token := c.accessToken.GetToken()
	maps := util.Map{
		"code":    code,
		"card_id": cardID,
	}
	return core.PostJSON(Link(cardCodeUnavailable), token.KeyMap(), maps)
}

//CodeConsume 核销用户礼品卡接口
//	接口说明
//	当礼品卡被使用完毕或者发生转存、绑定等操作后，开发者可以通过该接口核销用户的礼品卡，使礼品卡在列表中沉底并不再被使用。
//	接口调用请求说明
//	协议	HTTPS
//	http请求方式	POST
//	请求Url	https://api.weixin.qq.com/card/code/consume?access_token=TOKEN
//	POST数据格式	JSON
func (c *Card) CodeConsume(code string, cardID string) core.Responder {
	token := c.accessToken.GetToken()
	maps := util.Map{
		"code": code,
	}
	if cardID != "" {
		maps.Set("card_id", cardID)
	}
	return core.PostJSON(Link(cardCodeConsume), token.KeyMap(), maps)
}

// CodeDecrypt ...
func (c *Card) CodeDecrypt(encryptCode string) core.Responder {
	token := c.accessToken.GetToken()
	maps := util.Map{
		"encrypt_code": encryptCode,
	}

	return core.PostJSON(Link(cardCodeDecrypt), token.KeyMap(), maps)
}

//Delete 删除卡券接口
//删除卡券接口允许商户删除任意一类卡券。删除卡券后，该卡券对应已生成的领取用二维码、添加到卡包JS API均会失效。 注意：如用户在商家删除卡券前已领取一张或多张该卡券依旧有效。即删除卡券不能删除已被用户领取，保存在微信客户端中的卡券。
//接口调用请求说明
//HTTP请求方式: POST URL:https://api.weixin.qq.com/card/delete?access_token=TOKEN
func (c *Card) Delete(cardID string) core.Responder {
	token := c.accessToken.GetToken()
	maps := util.Map{
		"card_id": cardID,
	}
	return core.PostJSON(Link(cardDelete), token.KeyMap(), maps)
}

// GetUserCards ...
func (c *Card) GetUserCards(openID, cardID string) core.Responder {
	token := c.accessToken.GetToken()
	maps := util.Map{
		"openid":  openID,
		"card_id": cardID,
	}
	return core.PostJSON(Link(cardUserGetcardlist), token.KeyMap(), maps)
}

// SetPayCell ...
func (c *Card) SetPayCell(cardID string, isOpen bool) core.Responder {
	token := c.accessToken.GetToken()
	maps := util.Map{
		"is_open": isOpen,
		"card_id": cardID,
	}
	return core.PostJSON(Link(cardPaycellSet), token.KeyMap(), maps)
}

// ModifyStock ...
func (c *Card) ModifyStock(cardID string, option util.Map) core.Responder {
	token := c.accessToken.GetToken()
	maps := util.Map{
		"card_id": cardID,
	}
	maps.Join(option)
	return core.PostJSON(Link(cardModifystock), token.KeyMap(), maps)
}

// PayActivate ...
func (c *Card) PayActivate() core.Responder {
	token := c.accessToken.GetToken()
	return core.Get(Link(cardPayActivate), token.KeyMap())
}

// PayGetPrice ...
func (c *Card) PayGetPrice(cardID string, quantity int) core.Responder {
	token := c.accessToken.GetToken()
	return core.PostJSON(Link(cardPayGetpayprice), token.KeyMap(), util.Map{
		"card_id":  cardID,
		"quantity": quantity,
	})
}

// PayGetCoinsInfo ...
func (c *Card) PayGetCoinsInfo() core.Responder {
	token := c.accessToken.GetToken()
	return core.Get(Link(cardPayGetcoinsinfo), token.KeyMap())
}

// PayRecharge ...
func (c *Card) PayRecharge(count int) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardPayGetpayprice), token, util.Map{
		"coin_count": count,
	})
}

// PayOrder ...
func (c *Card) PayOrder(orderID string) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardPayGetorder), token, util.Map{
		"order_id": orderID,
	})

}

// PayGetOrderList ...
func (c *Card) PayGetOrderList(p util.Map) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardPayGetorderlist), token, util.MapNilMake(p))
}

// PayConfirm ...
func (c *Card) PayConfirm(cardID, orderID string, quantity int) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardPayConfirm), token, util.Map{
		"card_id":  cardID,
		"order_id": orderID,
		"quantity": quantity,
	})
}

// GeneralActivate ...
func (c *Card) GeneralActivate(p util.Map) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardGeneralcardActivate), token, util.MapNilMake(p))
}

// GeneralDeactivate ...
func (c *Card) GeneralDeactivate(cardID, code string) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardGeneralcardUnactivate), token, util.Map{
		"card_id": cardID,
		"code":    code,
	})
}

//GeneralUpdateUser 更新用户礼品卡信息接口
//	接口说明
//	当礼品卡被使用后，开发者可以通过该接口变更某个礼品卡的余额信息。
//	接口调用请求说明
//	协议	HTTPS
//	http请求方式	POST
//	请求Url	https://api.weixin.qq.com/card/generalcard/updateuser?access_token=TOKEN
//	POST数据格式	JSON
func (c *Card) GeneralUpdateUser(p util.Map) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardGeneralcardUpdateuser), token, util.MapNilMake(p))
}

// MeetingUpdateUser ...
func (c *Card) MeetingUpdateUser(p util.Map) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardMeetingticketUpdateuser), token, util.MapNilMake(p))
}

//GiftAdd 创建-礼品卡货架接口
//	接口说明
//	开发者可以通过该接口创建一个礼品卡货架并且用于公众号、门店的礼品卡售卖。
//	接口调用请求说明
//	协议	HTTPS
//	http请求方式	POST
//	请求Url	https://api.weixin.qq.com/card/giftcard/page/add?access_token=ACCESS_TOKEN
//	POST数据格式	JSON
func (c *Card) GiftAdd(p util.Map) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardGiftcardPageAdd), token, util.MapNilMake(p))
}

//GiftGet 查询-礼品卡货架信息接口
//	接口说明
//	开发者可以查询某个礼品卡货架信息。
//	接口调用请求说明
//	协议	HTTPS
//	http请求方式	POST
//	请求Url	https://api.weixin.qq.com/card/giftcard/page/get?access_token=ACCESS_TOKEN
//	POST数据格式	JSON
func (c *Card) GiftGet(pageID string) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardGiftcardPageGet), token, util.Map{
		"page_id": pageID,
	})
}

//GiftUpdate 修改-礼品卡货架信息接口
//	接口说明
//	开发者可以通过该接口更新礼品卡货架信息。
//	接口调用请求说明
//	协议	HTTPS
//	http请求方式	POST
//	请求Url	https://api.weixin.qq.com/card/giftcard/page/update?access_token=ACCESS_TOKEN
//	POST数据格式	JSON
func (c *Card) GiftUpdate(pageID string, p util.Map) core.Responder {
	token := c.accessToken.KeyMap()
	p = util.MapNilMake(p)
	p.Set("page_id", pageID)
	return core.PostJSON(Link(cardGiftcardPageUpdate), token, p)
}

//GiftBatchGet 查询-礼品卡货架列表接口
//接口说明
//开发者可以通过该接口查询当前商户下所有的礼品卡货架id。
//接口调用请求说明
//协议	HTTPS
//http请求方式	POST
//请求Url	https://api.weixin.qq.com/card/giftcard/page/batchget?access_token=ACCESS_TOKEN
func (c *Card) GiftBatchGet() core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardGiftcardPageBatchget), token, util.Map{})
}

//GiftSet 下架-礼品卡货架接口
//接口说明
//开发者可以通过该接口查询当前商户下所有的礼品卡货架id。
//接口调用请求说明
//协议	HTTPS
//http请求方式	POST
//请求Url	https://api.weixin.qq.com/card/giftcard/maintain/set?access_token=ACCESS_TOKEN
func (c *Card) giftSet(pageID string, all, maintain bool) core.Responder {
	token := c.accessToken.KeyMap()
	maps := util.Map{
		"maintain": true,
	}
	if pageID == "" {
		maps.Set("all", all)
	} else {
		maps.Set("page_id", pageID)
	}
	return core.PostJSON(Link(cardGiftcardMaintainSet), token, maps)
}

// GiftSetByID ...
func (c *Card) GiftSetByID(pageID string, maintain bool) core.Responder {
	return c.giftSet(pageID, false, maintain)
}

// GiftSetAll ...
func (c *Card) GiftSetAll(all, maintain bool) core.Responder {
	return c.giftSet("", all, maintain)
}

//GiftRefund 退款接口
//	接口说明
//	开发者可以通过该接口对某一笔订单操作退款，注意该接口比较隐私，请开发者提高操作该功能的权限等级。
//	接口调用请求说明
//	协议	HTTPS
//	http请求方式	POST
//	请求Url	https://api.weixin.qq.com/card/giftcard/order/refund?access_token=ACCESS_TOKEN
//	POST数据格式	JSON
//	请求数据说明：
//	参数	说明	是否必填
//	order_id	须退款的订单id	是
func (c *Card) GiftRefund(orderID string) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(Link(cardGiftcardMaintainSet), token, util.Map{
		"order_id": orderID,
	})
}

//InvoiceSetPayMch 设置支付后开票信息
//	接口说明
//	商户可以通过该接口设置某个商户号发生收款后在支付消息上出现开票授权按钮。
//	请求url：
//	https://api.weixin.qq.com/card/invoice/setbizattr?action=set_pay_mch&access_token={access_token}
//	请求方法：POST
//	参数	类型	是否必填	描述
//	paymch_info	Object	是	授权页字段
//	paymch_info包含以下字段：
//	参数	类型	是否必填	描述
//	mchid	String	是	微信支付商户号
//	s_pappid	String	是	开票平台id，需要找开票平台提供
func (c *Card) InvoiceSetPayMch(mchID string, appID string) core.Responder {
	token := c.accessToken.KeyMap()
	token.Set("action", "set_pay_mch")
	maps := util.Map{}
	maps.Set("paymch_info.mchid", mchID)
	maps.Set("paymch_info.s_pappid", appID)
	return core.PostJSON(Link(cardGiftcardMaintainSet), token, maps)
}

//InvoiceGetPayMch 查询支付后开票信息接口
//	请求url：
//	https://api.weixin.qq.com/card/invoice/setbizattr?action=get_pay_mch&access_token={access_token}
//	请求方法：POST
func (c *Card) InvoiceGetPayMch() core.Responder {
	token := c.accessToken.KeyMap()
	token.Set("action", "get_pay_mch")
	maps := util.Map{}
	return core.PostJSON(Link(cardGiftcardMaintainSet), token, maps)
}

//GetCardAPITicket get ticket
func (c *Card) GetCardAPITicket(refresh bool) {
	c.jssdk.GetTicket("wx_card", refresh)
}

//NewOneCard 创建卡券信息
//	参数:
//	cardType 卡券类型
//	data	卡券信息 (可传nil)
func NewOneCard(cardType CardType, data util.Map) *OneCard {
	ct := strings.ToLower(cardType.String())
	return &OneCard{
		CardType: cardType,
		data: util.Map{
			"card_type": cardType,
			ct:          data,
		},
	}
}

//AddAdvancedInfo 添加卡券advanced_info
func (c *OneCard) AddAdvancedInfo(info *CardAdvancedInfo) *OneCard {
	return c.add("advanced_info", info)
}

//AddBaseInfo 添加卡券base_info
func (c *OneCard) AddBaseInfo(info *CardBaseInfo) *OneCard {
	return c.add("base_info", info)
}

//AddDealDetail 添加卡券deal_detail
func (c *OneCard) AddDealDetail(d string) *OneCard {
	return c.add("deal_detail", d)
}

func (c *OneCard) add(name string, info interface{}) *OneCard {
	ct := strings.ToLower(c.CardType.String())
	if c.data != nil {
		if v, b := c.data[ct].(util.Map); b {
			if v != nil {
				v[name] = info
			} else {
				v = util.Map{
					name: info,
				}
			}
			c.data[ct] = v
		}
	} else {
		c.data = util.Map{
			"card_type": c.CardType,
			ct: util.Map{
				name: info,
			},
		}
	}
	return c
}

//Set 设置卡券信息(包含base_info,advanced_info,deal_detail)
func (c *OneCard) Set(cardType CardType, data util.Map) {
	ct := strings.ToLower(cardType.String())
	c.CardType = cardType
	c.data = util.Map{
		"card_type": ct,
		ct:          data,
	}
}

//Get 获取卡券类型,卡券信息
func (c *OneCard) Get() (CardType, util.Map) {
	return c.CardType, c.data
}

// ToMap ...
func (c *OneCard) ToMap() util.Map {
	maps := util.Map{}
	v, err := json.Marshal(c)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(v, maps)
	if err != nil {
		return nil
	}
	return maps
}
