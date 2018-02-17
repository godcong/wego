package wego

type Sandbox interface {
	GetKey() string
}

type sandbox struct {
	Config
}

var sandboxInst Sandbox

func initSandbox(config Config) {
	sandboxInst = NewSandbox(config)
}

func NewSandbox(config Config) Sandbox {
	return &sandbox{
		Config: config,
	}
}

func (s *sandbox) GetKey() string {
	if v, e := SandboxSignKey(s.Config); e == nil {
		return string(v)
	}
	return ""
}
