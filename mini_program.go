package wego

import (
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
)

type Auth interface {
	Session(code string) util.Map
}

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

type AppCode interface {
	Get(path string, optionals util.Map) util.Map
	GetQrCode(path string, width int) util.Map
	GetUnlimit(scene string, optionals util.Map) util.Map
}

type MiniProgram interface {
	Auth() Auth
	AppCode() AppCode
	//Client() core.Client
	DataCube() DataCube
	AccessToken() AccessToken
}

//func NewAppCode(application core.Application, config config.Config) AppCode {
//	return &mini_program.AppCode{
//		Config: config,
//		//mini_program:   application.MiniProgram(),
//	}
//}

func GetMiniProgram() MiniProgram {
	obj := GetApp().Get("mini_program").(MiniProgram)
	log.Debug("GetMiniProgram|obj:", obj)
	return obj
}

func GetAuth() Auth {
	return GetMiniProgram().Auth()
}

func GetAppCode() AppCode {
	return GetMiniProgram().AppCode()
}

func GetDataCube() DataCube {
	return GetMiniProgram().DataCube()
}
