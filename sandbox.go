package wego

type Sandbox interface {
	GetKey() string
	GetCacheKey() string
}

type sandbox struct {
	Config
	app Application
}

func NewSandbox(application Application, config Config) Sandbox {
	return &sandbox{
		Config: config,
		app:    application,
	}
}

func (s *sandbox) GetKey() string {
	return string(SandboxSignKey(s.Config))
}

func (s *sandbox) GetCacheKey() string {
	return ""
}
