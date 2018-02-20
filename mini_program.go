package wego

type MiniProgram interface {
	Client() Client
}

type miniProgram struct {
	Config
	app    Application
	client Client
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
