package mini

//type MiniProgram struct {
//	core.Config
//	acc    token.AccessTokenInterface
//	app    Application
//	client Client
//}

//func (m *MiniProgram) AccessToken() token.AccessTokenInterface {
//	if m.acc == nil {
//		m.acc = NewMiniProgramAccessToken(m.app, m.Config)
//	}
//	return m.acc
//}
//
//func (m *MiniProgram) Client() Client {
//	if m.client == nil {
//		m.client = app.Client(m.Config)
//	}
//	return m.client
//}
//
//func NewMiniProgram(application Application) MiniProgram {
//	config := application.GetConfig("mini.default")
//	return &MiniProgram{
//		Config: config,
//		app:    application,
//		client: app.Client(config),
//	}
//}
//
//type MiniProgramAccessToken struct {
//	token.AccessToken
//	core.Config
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
//func (m *MiniProgramAccessToken) getCredentials() core.Map {
//	return core.Map{
//		"grant_type": "client_credential",
//		"appid":      m.Get("app_id"),
//		"secret":     m.Get("secret"),
//	}
//}
//
//func (m *MiniProgramAccessToken) Session(code string) core.Map {
//	param := core.Map{
//		"appid":      m.Get("app_id"),
//		"secret":     m.Get("secret"),
//		"js_code":    code,
//		"grant_type": "authorization_code",
//	}
//	return m.app.Client(m.Config).Request(core.SNS_JSCODE2SESSION_URL_SUFFIX+"?"+param.ToUrl(), nil, "get", nil)
//}
