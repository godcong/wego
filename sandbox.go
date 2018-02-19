package wego

type Sandbox interface {
	GetKey() string
	GetCacheKey() string
}

type sandbox struct {
	app Application
}

func NewSandbox(application Application) Sandbox {
	return &sandbox{
		app: application,
	}
}

func (s *sandbox) GetKey() string {
	if v, e := SandboxSignKey(s.app.Config().GetConfig("payment.default")); e == nil {
		return string(v)
	}
	return ""
}

func (s *sandbox) GetCacheKey() string {
	return ""
}
