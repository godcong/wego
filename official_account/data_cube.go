package official_account

import (
	"time"

	"github.com/godcong/wego/config"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/net"
	"github.com/godcong/wego/util"
)

type DataCube struct {
	config.Config
	*OfficialAccount
}

func newDataCube(officialAccount *OfficialAccount) *DataCube {
	return &DataCube{
		Config:          defaultConfig,
		OfficialAccount: officialAccount,
	}
}

func NewDataCube() *DataCube {
	return newDataCube(account)
}

// 获取用户增减数据（getusersummary）	7	https://api.weixin.qq.com/datacube/getusersummary?access_token=ACCESS_TOKEN
// 成功:
// {"list":[{"ref_date":"2018-03-19","user_source":0,"new_user":0,"cancel_user":0},{"ref_date":"2018-03-19","user_source":17,"new_user":1,"cancel_user":0}]}
// 成功:
// {"list":[{"ref_date":"2018-03-19","user_source":0,"new_user":0,"cancel_user":0},{"ref_date":"2018-03-19","user_source":17,"new_user":1,"cancel_user":0}]}
// 失败:
// {"errcode":61501,"errmsg":"date range error hint: [_muTLA05701504]"}
func (d *DataCube) GetUserSummary(beginDate, endDate time.Time) *net.Response {
	log.Debug("DataCube|GetUserSummary", beginDate, endDate)
	return d.get(
		DATACUBE_GETUSERSUMMARY_URL_SUFFIX,
		beginDate.Format(DATACUBE_TIME_LAYOUT),
		endDate.Format(DATACUBE_TIME_LAYOUT),
	)
}

// 获取累计用户数据（getusercumulate）	7	https://api.weixin.qq.com/datacube/getusercumulate?access_token=ACCESS_TOKEN
// 成功:
// {"list":[{"ref_date":"2018-03-18","user_source":0,"cumulate_user":5},{"ref_date":"2018-03-19","user_source":0,"cumulate_user":6},{"ref_date":"2018-03-20","user_source":0,"cumulate_user":6}]}
// 失败:
// {"errcode":61501,"errmsg":"date range error hint: [_muTLA05701504]"}
func (d *DataCube) GetUserCumulate(beginDate, endDate time.Time) *net.Response {
	log.Debug("DataCube|GetUserCumulate", beginDate, endDate)
	return d.get(
		DATACUBE_GETUSERCUMULATE_URL_SUFFIX,
		beginDate.Format(DATACUBE_TIME_LAYOUT),
		endDate.Format(DATACUBE_TIME_LAYOUT),
	)
}

// 获取图文群发每日数据（getarticlesummary）	1	https://api.weixin.qq.com/datacube/getarticlesummary?access_token=ACCESS_TOKEN
// 成功:
// {"list":[]}
// 失败:
// {"errcode":61501,"errmsg":"date range error hint: [_muTLA05701504]"}
func (d *DataCube) GetArticleSummary(beginDate, endDate time.Time) *net.Response {
	log.Debug("DataCube|GetArticleSummary", beginDate, endDate)
	return d.get(
		DATACUBE_GETARTICLESUMMARY_URL_SUFFIX,
		beginDate.Format(DATACUBE_TIME_LAYOUT),
		endDate.Format(DATACUBE_TIME_LAYOUT),
	)
}

// 获取图文群发总数据（getarticletotal）	1	https://api.weixin.qq.com/datacube/getarticletotal?access_token=ACCESS_TOKEN
// 成功:
// {"list":[]}
// 失败:
// {"errcode":61501,"errmsg":"date range error hint: [_muTLA05701504]"}
func (d *DataCube) GetArticleTotal(beginDate, endDate time.Time) *net.Response {
	log.Debug("DataCube|GetArticleTotal", beginDate, endDate)
	return d.get(
		DATACUBE_GETARTICLETOTAL_URL_SUFFIX,
		beginDate.Format(DATACUBE_TIME_LAYOUT),
		endDate.Format(DATACUBE_TIME_LAYOUT),
	)
}

// 获取图文统计数据（getuserread）	3	https://api.weixin.qq.com/datacube/getuserread?access_token=ACCESS_TOKEN
// 成功:
// {"list":[]}
// 失败:
// {"errcode":61501,"errmsg":"date range error hint: [_muTLA05701504]"}
func (d *DataCube) GetUserRead(beginDate, endDate time.Time) *net.Response {
	log.Debug("DataCube|GetUserRead", beginDate, endDate)
	return d.get(
		DATACUBE_GETUSERREAD_URL_SUFFIX,
		beginDate.Format(DATACUBE_TIME_LAYOUT),
		endDate.Format(DATACUBE_TIME_LAYOUT),
	)
}

// 获取图文统计分时数据（getuserreadhour）	1	https://api.weixin.qq.com/datacube/getuserreadhour?access_token=ACCESS_TOKEN
// 成功:
// {"list":[]}
// 失败:
// {"errcode":61501,"errmsg":"date range error hint: [_muTLA05701504]"}
func (d *DataCube) GetUserReadHour(beginDate, endDate time.Time) *net.Response {
	log.Debug("DataCube|GetUserReadHour", beginDate, endDate)
	return d.get(
		DATACUBE_GETUSERREADHOUR_URL_SUFFIX,
		beginDate.Format(DATACUBE_TIME_LAYOUT),
		endDate.Format(DATACUBE_TIME_LAYOUT),
	)
}

// 获取图文分享转发数据（getusershare）	7	https://api.weixin.qq.com/datacube/getusershare?access_token=ACCESS_TOKEN
// 成功:
// {"list":[]}
// 失败:
// {"errcode":61501,"errmsg":"date range error hint: [_muTLA05701504]"}
func (d *DataCube) GetUserShare(beginDate, endDate time.Time) *net.Response {
	log.Debug("DataCube|GetUserReadHour", beginDate, endDate)
	return d.get(
		DATACUBE_GETUSERSHARE_URL_SUFFIX,
		beginDate.Format(DATACUBE_TIME_LAYOUT),
		endDate.Format(DATACUBE_TIME_LAYOUT),
	)
}

// 获取图文分享转发分时数据（getusersharehour）	1	https://api.weixin.qq.com/datacube/getusersharehour?access_token=ACCESS_TOKEN
// 失败:
// {"errcode":61501,"errmsg":"date range error hint: [_muTLA05701504]"}
func (d *DataCube) GetUserShareHour(beginDate, endDate time.Time) *net.Response {
	log.Debug("DataCube|GetUserReadHour", beginDate, endDate)
	return d.get(
		DATACUBE_GETUSERSHAREHOUR_URL_SUFFIX,
		beginDate.Format(DATACUBE_TIME_LAYOUT),
		endDate.Format(DATACUBE_TIME_LAYOUT),
	)
}

// 获取消息发送概况数据（getupstreammsg）	7	https://api.weixin.qq.com/datacube/getupstreammsg?access_token=ACCESS_TOKEN
// 成功:
// {"list":[]}
func (d *DataCube) GetUpstreamMsg(beginDate, endDate time.Time) *net.Response {
	log.Debug("DataCube|GetUpstreamMsg", beginDate, endDate)
	return d.get(
		DATACUBE_GETUPSTREAMMSG_URL_SUFFIX,
		beginDate.Format(DATACUBE_TIME_LAYOUT),
		endDate.Format(DATACUBE_TIME_LAYOUT),
	)
}

// 获取消息分送分时数据（getupstreammsghour）	1	https://api.weixin.qq.com/datacube/getupstreammsghour?access_token=ACCESS_TOKEN
// 成功:
// {"list":[]}
func (d *DataCube) GetUpstreamMsgHour(beginDate, endDate time.Time) *net.Response {
	log.Debug("DataCube|GetUpstreamMsgHour", beginDate, endDate)
	return d.get(
		DATACUBE_GETUPSTREAMMSGHOUR_URL_SUFFIX,
		beginDate.Format(DATACUBE_TIME_LAYOUT),
		endDate.Format(DATACUBE_TIME_LAYOUT),
	)
}

// 获取消息发送周数据（getupstreammsgweek）	30	https://api.weixin.qq.com/datacube/getupstreammsgweek?access_token=ACCESS_TOKEN
// 成功:
// {"list":[]}
func (d *DataCube) GetUpstreamMsgWeek(beginDate, endDate time.Time) *net.Response {
	log.Debug("DataCube|GetUpstreamMsgWeek", beginDate, endDate)
	return d.get(
		DATACUBE_GETUPSTREAMMSGWEEK_URL_SUFFIX,
		beginDate.Format(DATACUBE_TIME_LAYOUT),
		endDate.Format(DATACUBE_TIME_LAYOUT),
	)
}

// 获取消息发送月数据（getupstreammsgmonth）	30	https://api.weixin.qq.com/datacube/getupstreammsgmonth?access_token=ACCESS_TOKEN
// 成功:
// {"list":[]}
func (d *DataCube) GetUpstreamMsgMonth(beginDate, endDate time.Time) *net.Response {
	log.Debug("DataCube|GetUpstreamMsgMonth", beginDate, endDate)
	return d.get(
		DATACUBE_GETUPSTREAMMSGMONTH_URL_SUFFIX,
		beginDate.Format(DATACUBE_TIME_LAYOUT),
		endDate.Format(DATACUBE_TIME_LAYOUT),
	)
}

// 获取消息发送分布数据（getupstreammsgdist）	15	https://api.weixin.qq.com/datacube/getupstreammsgdist?access_token=ACCESS_TOKEN
// 成功:
// {"list":[]}
func (d *DataCube) GetUpstreamMsgDist(beginDate, endDate time.Time) *net.Response {
	log.Debug("DataCube|GetUpstreamMsgDist", beginDate, endDate)
	return d.get(
		DATACUBE_GETUPSTREAMMSGDIST_URL_SUFFIX,
		beginDate.Format(DATACUBE_TIME_LAYOUT),
		endDate.Format(DATACUBE_TIME_LAYOUT),
	)
}

// 获取消息发送分布周数据（getupstreammsgdistweek）	30	https://api.weixin.qq.com/datacube/getupstreammsgdistweek?access_token=ACCESS_TOKEN
// 成功:
// {"list":[]}
func (d *DataCube) GetUpstreamMsgDistWeek(beginDate, endDate time.Time) *net.Response {
	log.Debug("DataCube|GetUpstreamMsgDistWeek", beginDate, endDate)
	return d.get(
		DATACUBE_GETUPSTREAMMSGDISTWEEK_URL_SUFFIX,
		beginDate.Format(DATACUBE_TIME_LAYOUT),
		endDate.Format(DATACUBE_TIME_LAYOUT),
	)
}

// 获取消息发送分布月数据（getupstreammsgdistmonth）	30	https://api.weixin.qq.com/datacube/getupstreammsgdistmonth?access_token=ACCESS_TOKEN
// 成功:
// {"list":[]}
func (d *DataCube) GetUpstreamMsgDistMonth(beginDate, endDate time.Time) *net.Response {
	log.Debug("DataCube|GetUpstreamMsgDistMonth", beginDate, endDate)
	return d.get(
		DATACUBE_GETUPSTREAMMSGDISTMONTH_URL_SUFFIX,
		beginDate.Format(DATACUBE_TIME_LAYOUT),
		endDate.Format(DATACUBE_TIME_LAYOUT),
	)
}

// 获取接口分析数据（getinterfacesummary）	30	https://api.weixin.qq.com/datacube/getinterfacesummary?access_token=ACCESS_TOKEN
// 成功:
// {"list":[{"ref_date":"2018-03-20","callback_count":24,"fail_count":0,"total_time_cost":5965,"max_time_cost":1290}]}
func (d *DataCube) GetInterfaceSummary(beginDate, endDate time.Time) *net.Response {
	log.Debug("DataCube|GetInterfaceSummary", beginDate, endDate)
	return d.get(
		DATACUBE_GETINTERFACESUMMARY_URL_SUFFIX,
		beginDate.Format(DATACUBE_TIME_LAYOUT),
		endDate.Format(DATACUBE_TIME_LAYOUT),
	)
}

// 获取接口分析分时数据（getinterfacesummaryhour）	1	https://api.weixin.qq.com/datacube/getinterfacesummaryhour?access_token=ACCESS_TOKEN
// 成功:
// {"list":[{"ref_date":"2018-03-20","ref_hour":1800,"callback_count":24,"fail_count":0,"total_time_cost":5965,"max_time_cost":1290}]}
func (d *DataCube) GetInterfaceSummaryHour(beginDate, endDate time.Time) *net.Response {
	log.Debug("DataCube|GetInterfaceSummaryHour", beginDate, endDate)
	return d.get(
		DATACUBE_GETINTERFACESUMMARYHOUR_URL_SUFFIX,
		beginDate.Format(DATACUBE_TIME_LAYOUT),
		endDate.Format(DATACUBE_TIME_LAYOUT),
	)
}

func (d *DataCube) get(url, beginDate, endDate string) *net.Response {
	key := d.token.GetToken().KeyMap()
	resp := d.client.HttpPostJson(
		d.client.Link(url),
		util.Map{"begin_date": beginDate, "end_date": endDate},
		util.Map{net.REQUEST_TYPE_QUERY.String(): key})
	return resp
}
