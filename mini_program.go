package wego

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/mini"
)

type Auth interface {
	Session(code string) core.Map
}

type DataCube interface {
	UserPortrait(from, to string) core.Map
	SummaryTrend(from, to string) core.Map
	DailyVisitTrend(from, to string) core.Map
	WeeklyVisitTrend(from, to string) core.Map
	MonthlyVisitTrend(from, to string) core.Map
	VisitDistribution(from, to string) core.Map
	DailyRetainInfo(from, to string) core.Map
	WeeklyRetainInfo(from, to string) core.Map
	MonthlyRetainInfo(from, to string) core.Map
	VisitPage(from, to string) core.Map
}

type AppCode interface {
	Get(path string, optionals core.Map) core.Map
	GetQrCode(path string, width int) core.Map
	GetUnlimit(scene string, optionals core.Map) core.Map
}

type MiniProgram interface {
	Auth() *mini.Auth
	AppCode() *mini.AppCode
	//Client() core.Client
	DataCube() *mini.DataCube
	//accessToken() token.AccessTokenInterface
}

func NewAppCode(application core.Application, config core.Config) AppCode {
	return &mini.AppCode{
		Config: config,
		//mini:   application.MiniProgram(),
	}
}

func GetMiniProgram() MiniProgram {
	mini := GetApp().Get("mini_program").(MiniProgram)
	core.Debug("GetMiniProgram|mini:", mini)
	return mini

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
