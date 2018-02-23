package wego

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/mini"
	"github.com/godcong/wego/token"
)

type AppCode interface {
}
type MiniProgram interface {
	Client() core.Client
	AccessToken() token.AccessTokenInterface
}

func NewAppCode(application core.Application, config core.Config) AppCode {
	return &mini.AppCode{
		Config: config,
		//mini:   application.MiniProgram(),
	}
}
