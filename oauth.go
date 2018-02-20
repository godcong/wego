package wego

type OAuth interface {
}

type oauth struct {
	Config
	app Application
}

func NewOAuth(application Application, config Config) OAuth {
	return &oauth{
		Config: config,
		app:    application,
		//client: application.Client(),
	}
}
