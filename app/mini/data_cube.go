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
	return newDataCube(NewMiniProgram(config)).(*DataCube)
}

func (d *DataCube) query(api, from, to string) core.Responder {
	token := d.accessToken.GetToken()
	params := util.Map{
		"begin_date": from,
		"end_date":   to,
	}
	return core.PostJSON(api, token.KeyMap(), params)
}

//UserPortrait 用户画像
//接口地址: https://api.weixin.qq.com/datacube/getweanalysisappiduserportrait?access_token=ACCESS_TOKEN
//POST 请求参数说明:
//参数	是否必填	说明
//begin_date	是	开始日期
//end_date	是	结束日期，开始日期与结束日期相差的天数限定为0/6/29，分别表示查询最近1/7/30天数据，end_date允许设置的最大值为昨日
func (d *DataCube) UserPortrait(from, to string) core.Responder {
	return d.query(Link(datacubeGetweanalysisappiduserportrait), from, to)
}

//SummaryTrend 概况趋势
//接口地址
//https://api.weixin.qq.com/datacube/getweanalysisappiddailysummarytrend?access_token=ACCESS_TOKEN
//获取 access_token 详见文档
//POST 请求参数说明:
//参数	是否必填	说明
//begin_date	是	开始日期
//end_date	是	结束日期，限定查询1天数据，end_date允许设置的最大值为昨日
func (d *DataCube) SummaryTrend(from, to string) core.Responder {
	return d.query(Link(datacubeGetweanalysisappiddailysummarytrend), from, to)
}

//DailyVisitTrend 日趋势
//接口地址
//https://api.weixin.qq.com/datacube/getweanalysisappiddailyvisittrend?access_token=ACCESS_TOKEN
//POST 参数说明
//参数	是否必填	说明
//begin_date	是	开始日期
//end_date	是	结束日期，限定查询1天数据，end_date允许设置的最大值为昨日
func (d *DataCube) DailyVisitTrend(from, to string) core.Responder {
	return d.query(Link(datacubeGetweanalysisappiddailyvisittrend), from, to)
}

//WeeklyVisitTrend 周趋势
//接口地址
//https://api.weixin.qq.com/datacube/getweanalysisappidweeklyvisittrend?access_token=ACCESS_TOKEN
//POST 参数说明:
//参数	是否必填	说明
//begin_date	是	开始日期，为周一日期
//end_date	是	结束日期，为周日日期，限定查询一周数据
//注意:请求json和返回json与天的一致，这里限定查询一个自然周的数据，时间必须按照自然周的方式输入: 如:20170306(周一), 20170312(周日)
func (d *DataCube) WeeklyVisitTrend(from, to string) core.Responder {
	return d.query(Link(datacubeGetweanalysisappidweeklyvisittrend), from, to)
}

//MonthlyVisitTrend 月趋势
//接口地址
//https://api.weixin.qq.com/datacube/getweanalysisappidmonthlyvisittrend?access_token=ACCESS_TOKEN
//POST 参数说明:
//参数	是否必填	说明
//begin_date	是	开始日期，为自然月第一天
//end_date	是	结束日期，为自然月最后一天，限定查询一个月数据
//注意:请求json和返回json与天的一致，这里限定查询一个自然月的数据，时间必须按照自然月的方式输入: 如:20170201(月初), 20170228(月末)
func (d *DataCube) MonthlyVisitTrend(from, to string) core.Responder {
	return d.query(Link(datacubeGetweanalysisappidmonthlyvisittrend), from, to)
}

//VisitDistribution 访问分布
//接口地址
//https://api.weixin.qq.com/datacube/getweanalysisappidvisitdistribution?access_token=ACCESS_TOKEN
//POST 参数说明:
//参数	是否必填	说明
//begin_date	是	开始日期
//end_date	是	结束日期，限定查询1天数据，end_date允许设置的最大值为昨日
func (d *DataCube) VisitDistribution(from, to string) core.Responder {
	return d.query(Link(datacubeGetweanalysisappidvisitdistribution), from, to)
}

//DailyRetainInfo 日留存
//接口地址
//https://api.weixin.qq.com/datacube/getweanalysisappiddailyretaininfo?access_token=ACCESS_TOKEN
//POST 参数说明:
//参数	是否必填	说明
//begin_date	是	开始日期
//end_date	是	结束日期，限定查询1天数据，end_date允许设置的最大值为昨日
func (d *DataCube) DailyRetainInfo(from, to string) core.Responder {
	return d.query(Link(datacubeGetweanalysisappiddailyretaininfo), from, to)
}

//WeeklyRetainInfo 周留存
//接口地址
//https://api.weixin.qq.com/datacube/getweanalysisappidweeklyretaininfo?access_token=ACCESS_TOKEN
//POST 参数说明:
//参数	是否必填	说明
//begin_date	是	开始日期，为周一日期
//end_date	是	结束日期，为周日日期，限定查询一周数据
//注意:请求json和返回json与天的一致，这里限定查询一个自然周的数据，时间必须按照自然周的方式输入: 如:20170306(周一), 20170312(周日)
func (d *DataCube) WeeklyRetainInfo(from, to string) core.Responder {
	return d.query(Link(datacubeGetweanalysisappidweeklyretaininfo), from, to)
}

//MonthlyRetainInfo 月留存
//接口地址
//https://api.weixin.qq.com/datacube/getweanalysisappidmonthlyretaininfo?access_token=ACCESS_TOKEN
//POST 参数说明:
//参数	是否必填	说明
//begin_date	是	开始日期，为自然月第一天
//end_date	是	结束日期，为自然月最后一天，限定查询一个月数据
//注意:请求json和返回json与天的一致，这里限定查询一个自然月的数据，时间必须按照自然月的方式输入: 如:20170201(月初), 20170228(月末)
func (d *DataCube) MonthlyRetainInfo(from, to string) core.Responder {
	return d.query(Link(datacubeGetweanalysisappidmonthlyretaininfo), from, to)
}

//VisitPage 访问页面
//接口地址
//https://api.weixin.qq.com/datacube/getweanalysisappidvisitpage?access_token=ACCESS_TOKEN
//POST 参数说明:
//参数	是否必填	说明
//begin_date	是	开始日期
//end_date	是	结束日期，限定查询1天数据，end_date允许设置的最大值为昨日
func (d *DataCube) VisitPage(from, to string) core.Responder {
	return d.query(Link(datacubeGetweanalysisappidvisitpage), from, to)
}
