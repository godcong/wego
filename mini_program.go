package wego

import (
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
)

/*Auth 授权登录*/
type Auth interface {
	Session(code string) util.Map
}

/*DataCube DataCube*/
type DataCube interface {
	UserPortrait(from, to string) util.Map
	SummaryTrend(from, to string) util.Map
	DailyVisitTrend(from, to string) util.Map
	WeeklyVisitTrend(from, to string) util.Map
	MonthlyVisitTrend(from, to string) util.Map
	VisitDistribution(from, to string) util.Map
	DailyRetainInfo(from, to string) util.Map
	WeeklyRetainInfo(from, to string) util.Map
	MonthlyRetainInfo(from, to string) util.Map
	VisitPage(from, to string) util.Map
}

/*AppCode AppCode*/
type AppCode interface {
	Get(path string, optionals util.Map) util.Map
	GetQrCode(path string, width int) util.Map
	GetUnlimit(scene string, optionals util.Map) util.Map
}

/*Program 小程序*/
type MiniProgram interface {
	Auth() Auth
	AppCode() AppCode
	//Client() core.Client
	DataCube() DataCube
	AccessToken() AccessToken
}

/*GetMiniProgram 获取小程序*/
func GetMiniProgram() MiniProgram {
	obj := GetApp().Get("mini").(MiniProgram)
	log.Debug("GetMiniProgram|obj:", obj)
	return obj
}

// func GetAuth() Auth {
// 	return GetMiniProgram().Auth()
// }

// func GetAppCode() AppCode {
// 	return GetMiniProgram().AppCode()
// }

// func GetDataCube() DataCube {
// 	return GetMiniProgram().DataCube()
// }
