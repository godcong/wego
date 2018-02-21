package wego

type MiniProgram interface {
	Client() Client
	AccessToken() AccessTokenInterface
}

type miniProgram struct {
	Config
	acc    AccessTokenInterface
	app    Application
	client Client
}

func (m *miniProgram) AccessToken() AccessTokenInterface {
	if m.acc == nil {
		m.acc = NewMiniProgramAccessToken(m.app, m.Config)
	}
	return m.acc
}

func (m *miniProgram) Client() Client {
	if m.client == nil {
		m.client = app.Client(m.Config)
	}
	return m.client
}

func NewMiniProgram(application Application) MiniProgram {
	config := application.GetConfig("mini_program.default")
	return &miniProgram{
		Config: config,
		app:    application,
		client: app.Client(config),
	}
}

type mpAccessToken struct {
	AccessToken
	Config
	app Application
}

func NewMiniProgramAccessToken(application Application, config Config) AccessTokenInterface {
	return &mpAccessToken{
		Config: config,
		app:    application,
	}
}

func (m *mpAccessToken) getCredentials() Map {
	return Map{
		"grant_type": "client_credential",
		"appid":      m.Get("app_id"),
		"secret":     m.Get("secret"),
	}
}
