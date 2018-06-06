package mini

import (
	"github.com/godcong/wego"
	"github.com/godcong/wego/config"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/log"
)

/*Program Program */
type Program struct {
	config.Config
	client   *core.Client
	app      *core.Application
	token    *core.AccessToken
	auth     *Auth
	appCode  *AppCode
	dataCube *DataCube
	template *Template
}

var defaultConfig config.Config
var program *Program

func init() {
	defaultConfig = config.GetConfig("mini_program.default")
	app := core.App()
	program = newMiniProgram(defaultConfig, app)
	app.Register("mini_program", program)
	//app.Register(newMiniProgram())
}

func newMiniProgram(config config.Config, application *core.Application) *Program {
	mini0 := &Program{
		Config: config,
		app:    application,
		client: core.NewClient(config),
	}
	mini0.token = core.NewAccessToken(config, mini0.client)
	mini0.client.SetDataType(core.DataTypeJSON)
	mini0.client.SetDomain(core.NewDomain("mini"))
	return mini0
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

/*Auth Auth */
func (m *Program) Auth() wego.Auth {
	if m.auth == nil {
		m.auth = &Auth{
			Config:  m.Config,
			Program: m,
		}
	}
	return m.auth
}

/*AppCode AppCode */
func (m *Program) AppCode() wego.AppCode {
	if m.appCode == nil {
		m.appCode = &AppCode{
			Config:  m.Config,
			Program: m,
		}
	}
	return m.appCode
}

/*DataCube DataCube */
func (m *Program) DataCube() wego.DataCube {
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
		m.token = core.NewAccessToken(m.Config, m.client)
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
//	config.Config
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
