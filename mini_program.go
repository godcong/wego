package wego

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/mini"
)

type Auth interface {
	Session(code string) core.Map
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
	//AccessToken() token.AccessTokenInterface
}

func NewAppCode(application core.Application, config core.Config) AppCode {
	return &mini.AppCode{
		Config: config,
		//mini:   application.MiniProgram(),
	}
}

func GetMiniProgram() MiniProgram {
	obj := GetApplication().Get("mini_program").(MiniProgram)
	return obj
}

func GetAuth() Auth {
	return GetMiniProgram().Auth()
}

func GetAppCode() AppCode {
	return GetMiniProgram().AppCode()
}
