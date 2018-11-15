package mini

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

/*DataCube DataCube */
type DataCube struct {
	*Program
}

func newDataCube(program *Program) interface{} {
	return &DataCube{
		Program: program,
	}
}

// NewDataCube ...
func NewDataCube(config *core.Config) *DataCube {
	return newAppcode(NewMiniProgram(config)).(*DataCube)
}

func (d *DataCube) query(api, from, to string) core.Responder {
	token := d.accessToken.GetToken()
	params := util.Map{
		"begin_date": from,
		"end_date":   to,
	}
	return core.PostJSON(api, token.KeyMap(), params)
}

/*UserPortrait 用户画像
接口地址: https://api.weixin.qq.com/datacube/getweanalysisappiduserportrait?access_token=ACCESS_TOKEN
POST 请求参数说明:
参数	是否必填	说明
begin_date	是	开始日期
end_date	是	结束日期，开始日期与结束日期相差的天数限定为0/6/29，分别表示查询最近1/7/30天数据，end_date允许设置的最大值为昨日
返回参数说明:
每次请求返回选定的时间范围及以下指标项:
参数	说明
ref_date	时间范围,如: "20170611-20170617"
visit_uv_new	新用户
visit_uv	活跃用户
每个指标项下包括的属性:
参数	说明
province	省份，如北京、广东等
city	城市，如北京、广州等
genders	性别，包括男、女、未知
platforms	终端类型，包括iPhone, android,其他
devices	机型，如苹果iPhone6, OPPO R9等
ages	年龄，包括17岁以下、18-24岁等区间
每个属性下包括的数据项:
参数	说明
id	属性值id
name	属性值名称，与id一一对应。如属性为province时，返回的属性值名称包括“广东”等
value	属性值对应的指标值，如指标为visit_uv,属性为province,属性值为"广东省”，value对应广东地区的活跃用户数
*/
func (d *DataCube) UserPortrait(from, to string) core.Responder {
	return d.query(Link(datacubeGetweanalysisappiduserportrait), from, to)
}

/*SummaryTrend 概况趋势
接口地址
https://api.weixin.qq.com/datacube/getweanalysisappiddailysummarytrend?access_token=ACCESS_TOKEN
获取 access_token 详见文档
POST 请求参数说明:
参数	是否必填	说明
begin_date	是	开始日期
end_date	是	结束日期，限定查询1天数据，end_date允许设置的最大值为昨日
返回参数说明:
参数	说明
visit_total	累计用户数
share_pv	转发次数
share_uv	转发人数
*/
func (d *DataCube) SummaryTrend(from, to string) core.Responder {
	return d.query(Link(datacubeGetweanalysisappiddailysummarytrend), from, to)
}

/*DailyVisitTrend 日趋势
接口地址
https://api.weixin.qq.com/datacube/getweanalysisappiddailyvisittrend?access_token=ACCESS_TOKEN
POST 参数说明
参数	是否必填	说明
begin_date	是	开始日期
end_date	是	结束日期，限定查询1天数据，end_date允许设置的最大值为昨日
返回参数说明:
参数	说明
ref_date	时间: 如: "20170313"
session_cnt	打开次数
visit_pv	访问次数
visit_uv	访问人数
visit_uv_new	新用户数
stay_time_uv	人均停留时长 (浮点型，单位:秒)
stay_time_session	次均停留时长 (浮点型，单位:秒)
visit_depth	平均访问深度 (浮点型)
*/
func (d *DataCube) DailyVisitTrend(from, to string) core.Responder {
	return d.query(Link(datacubeGetweanalysisappiddailyvisittrend), from, to)
}

/*WeeklyVisitTrend 周趋势
接口地址
https://api.weixin.qq.com/datacube/getweanalysisappidweeklyvisittrend?access_token=ACCESS_TOKEN
POST 参数说明:
参数	是否必填	说明
begin_date	是	开始日期，为周一日期
end_date	是	结束日期，为周日日期，限定查询一周数据
注意:请求json和返回json与天的一致，这里限定查询一个自然周的数据，时间必须按照自然周的方式输入: 如:20170306(周一), 20170312(周日)
返回参数说明:
参数	说明
ref_date	时间，如:"20170306-20170312"
session_cnt	打开次数（自然周内汇总）
visit_pv	访问次数（自然周内汇总）
visit_uv	访问人数（自然周内去重）
visit_uv_new	新用户数（自然周内去重）
stay_time_uv	人均停留时长 (浮点型，单位:秒)
stay_time_session	次均停留时长 (浮点型，单位:秒)
visit_depth	平均访问深度 (浮点型)
*/
func (d *DataCube) WeeklyVisitTrend(from, to string) core.Responder {
	return d.query(Link(datacubeGetweanalysisappidweeklyvisittrend), from, to)
}

/*MonthlyVisitTrend 月趋势
接口地址
https://api.weixin.qq.com/datacube/getweanalysisappidmonthlyvisittrend?access_token=ACCESS_TOKEN
POST 参数说明:
参数	是否必填	说明
begin_date	是	开始日期，为自然月第一天
end_date	是	结束日期，为自然月最后一天，限定查询一个月数据
注意:请求json和返回json与天的一致，这里限定查询一个自然月的数据，时间必须按照自然月的方式输入: 如:20170201(月初), 20170228(月末)
返回参数说明:
参数	说明
ref_date	时间，如:"201702"
session_cnt	打开次数（自然月内汇总）
visit_pv	访问次数（自然月内汇总）
visit_uv	访问人数（自然月内去重）
visit_uv_new	新用户数（自然月内去重）
stay_time_uv	人均停留时长 (浮点型，单位:秒)
stay_time_session	次均停留时长 (浮点型，单位:秒)
visit_depth	平均访问深度 (浮点型)
*/
func (d *DataCube) MonthlyVisitTrend(from, to string) core.Responder {
	return d.query(Link(datacubeGetweanalysisappidmonthlyvisittrend), from, to)
}

/*VisitDistribution 访问分布
接口地址
https://api.weixin.qq.com/datacube/getweanalysisappidvisitdistribution?access_token=ACCESS_TOKEN
POST 参数说明:
参数	是否必填	说明
begin_date	是	开始日期
end_date	是	结束日期，限定查询1天数据，end_date允许设置的最大值为昨日
返回参数说明:
参数	说明
ref_date	时间: 如: "20170313"
list	存入所有类型的指标情况
list 的每一项包括:
参数	说明
index	分布类型
item_list	分布数据列表
分布类型（index）的取值范围:
值	说明
access_source_session_cnt	访问来源分布
access_staytime_info	访问时长分布
access_depth_info	访问深度的分布
每个数据项包括:
参数	说明
key	场景 id
value 场景下的值（均为整数型）
key对应关系如下:
访问来源:(index="access_source_session_cnt")
1:小程序历史列表
2:搜索
3:会话
4:二维码
5:公众号主页
6:聊天顶部
7:系统桌面
8:小程序主页
9:附近的小程序
10:其他
11:模板消息
12:客服消息
13: 公众号菜单
14: APP分享
15: 支付完成页
16: 长按识别二维码
17: 相册选取二维码
18: 公众号文章
19:钱包
20:卡包
21:小程序内卡券
22:其他小程序
23:其他小程序返回
24:卡券适用门店列表
25:搜索框快捷入口
26:小程序客服消息
27:公众号下发
28: 会话左下角菜单
29: 小程序任务栏
30: 长按小程序菜单圆点
31: 连wifi成功页
32: 城市服务
访问时长:(index="access_staytime_info")
1: 0-2s
2: 3-5s
3: 6-10s
4: 11-20s
5: 20-30s
6: 30-50s
7: 50-100s
8: > 100s
平均访问深度:(index="access_depth_info")
1: 1页
2: 2页
3: 3页
4: 4页
5: 5页
6: 6-10页
7: >10页
*/
func (d *DataCube) VisitDistribution(from, to string) core.Responder {
	return d.query(Link(datacubeGetweanalysisappidvisitdistribution), from, to)
}

/*DailyRetainInfo 日留存
接口地址
https://api.weixin.qq.com/datacube/getweanalysisappiddailyretaininfo?access_token=ACCESS_TOKEN
POST 参数说明:
参数	是否必填	说明
begin_date	是	开始日期
end_date	是	结束日期，限定查询1天数据，end_date允许设置的最大值为昨日
返回参数说明:
参数	说明
visit_uv_new	新增用户留存
visit_uv	活跃用户留存
visit_uv、visit_uv_new 的每一项包括:
参数	说明
key	标识，0开始，0表示当天，1表示1天后，依此类推，key取值分别是:0,1,2,3,4,5,6,7,14,30
value	key对应日期的新增用户数/活跃用户数（key=0时）或留存用户数（k>0时）
*/
func (d *DataCube) DailyRetainInfo(from, to string) core.Responder {
	return d.query(Link(datacubeGetweanalysisappiddailyretaininfo), from, to)
}

/*WeeklyRetainInfo 周留存
接口地址
https://api.weixin.qq.com/datacube/getweanalysisappidweeklyretaininfo?access_token=ACCESS_TOKEN
POST 参数说明:
参数	是否必填	说明
begin_date	是	开始日期，为周一日期
end_date	是	结束日期，为周日日期，限定查询一周数据
注意:请求json和返回json与天的一致，这里限定查询一个自然周的数据，时间必须按照自然周的方式输入: 如:20170306(周一), 20170312(周日)
返回参数说明:
参数	说明
ref_date	时间，如:"20170306-20170312"
visit_uv_new	新增用户留存
visit_uv	活跃用户留存
visit_uv、visit_uv_new 的每一项包括:
参数	说明
key	标识，0开始，0表示当周，1表示1周后，依此类推，key取值分别是:0,1,2,3,4
value	key对应日期的新增用户数/活跃用户数（key=0时）或留存用户数（k>0时）
*/
func (d *DataCube) WeeklyRetainInfo(from, to string) core.Responder {
	return d.query(Link(datacubeGetweanalysisappidweeklyretaininfo), from, to)
}

/*MonthlyRetainInfo 月留存
接口地址
https://api.weixin.qq.com/datacube/getweanalysisappidmonthlyretaininfo?access_token=ACCESS_TOKEN
POST 参数说明:
参数	是否必填	说明
begin_date	是	开始日期，为自然月第一天
end_date	是	结束日期，为自然月最后一天，限定查询一个月数据
注意:请求json和返回json与天的一致，这里限定查询一个自然月的数据，时间必须按照自然月的方式输入: 如:20170201(月初), 20170228(月末)
返回参数说明:
参数	说明
ref_date	时间，如:"201702"
visit_uv_new	新增用户留存
visit_uv	活跃用户留存
visit_uv、visit_uv_new 的每一项包括:
参数	说明
key	标识，0开始，0表示当月，1表示1月后，key取值分别是:0,1
value	key对应日期的新增用户数/活跃用户数（key=0时）或留存用户数（k>0时）
*/
func (d *DataCube) MonthlyRetainInfo(from, to string) core.Responder {
	return d.query(Link(datacubeGetweanalysisappidmonthlyretaininfo), from, to)
}

/*VisitPage 访问页面
接口地址
https://api.weixin.qq.com/datacube/getweanalysisappidvisitpage?access_token=ACCESS_TOKEN
POST 参数说明:
参数	是否必填	说明
begin_date	是	开始日期
end_date	是	结束日期，限定查询1天数据，end_date允许设置的最大值为昨日
返回参数说明:
参数	说明
page_path	页面路径
page_visit_pv	访问次数
page_visit_uv	访问人数
page_staytime_pv	次均停留时长
entrypage_pv	进入页次数
exitpage_pv	退出页次数
page_share_pv	转发次数
page_share_uv	转发人数
*/
func (d *DataCube) VisitPage(from, to string) core.Responder {
	return d.query(Link(datacubeGetweanalysisappidvisitpage), from, to)
}
