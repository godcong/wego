package mini

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
)

/*Program Program */
type Program struct {
	*core.Config
	client   *core.Client
	token    *core.AccessToken
	auth     *Auth
	appCode  *AppCode
	dataCube *DataCube
	template *Template
}

func initAccessToken(config *core.Config) *core.AccessToken {
	token := core.NewAccessToken()
	return token.SetCredentials(util.Map{
		"grant_type": "client_credential",
		"appid":      config.GetString("app_id"),
		"secret":     config.GetString("secret"),
	})
}

func newMiniProgram(config *core.Config) *Program {
	mini := &Program{
		Config: config,
		token:  initAccessToken(config),
	}

	return mini
}

/*SetClient SetClient */
func (m *Program) SetClient(c *core.Client) *Program {
	m.client = c
	return m
}

/*GetClient GetClient */
func (m *Program) GetClient() *core.Client {
	return m.client
}

/*OAuth OAuth */
func (m *Program) Auth() *Auth {
	if m.auth == nil {
		m.auth = &Auth{
			Config:  m.Config,
			Program: m,
		}
	}
	return m.auth
}

/*AppCode AppCode */
func (m *Program) AppCode() *AppCode {
	if m.appCode == nil {
		m.appCode = &AppCode{
			Config:  m.Config,
			Program: m,
		}
	}
	return m.appCode
}

/*DataCube DataCube */
func (m *Program) DataCube() *DataCube {
	if m.dataCube == nil {
		m.dataCube = &DataCube{
			Config:  m.Config,
			Program: m,
		}
	}
	return m.dataCube
}

/*Template Template */
func (m *Program) Template() *Template {
	if m.template == nil {
		m.template = &Template{
			Config:  m.Config,
			Program: m,
		}
	}
	return m.template
}

/*AccessToken AccessToken */
func (m *Program) AccessToken() *core.AccessToken {
	log.Debug("Program|AccessToken")
	if m.token == nil {
		m.token = core.NewAccessToken()
	}
	return m.token
}

//func (m *Program) accessToken() token.AccessTokenInterface {
//	if m.acc == nil {
//		m.acc = NewMiniProgramAccessToken(m.app, m.Config)
//	}
//	return m.acc
//}
//
//func (m *Program) Client() Client {
//	if m.client == nil {
//		m.client = app.Client(m.Config)
//	}
//	return m.client
//}
//
//func NewMiniProgram(application Application) Program {
//	config := application.GetConfig("mini.default")
//	return &Program{
//		Config: config,
//		app:    application,
//		client: app.Client(config),
//	}
//}
//
//type MiniProgramAccessToken struct {
//	token.accessToken
//	Config
//	app core.Application
//}
//
//func NewMiniProgramAccessToken(application Application, config Config) *MiniProgramAccessToken {
//	return &MiniProgramAccessToken{
//		Config: config,
//		app:    application,
//	}
//}
//
//func (m *MiniProgramAccessToken) getCredentials() util.Map {
//	return util.Map{
//		"grant_type": "client_credential",
//		"appid":      m.Get("app_id"),
//		"secret":     m.Get("secret"),
//	}
//}
