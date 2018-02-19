package wego

type Sandbox interface {
	GetKey() string
}

type sandbox struct {
	Config
	app Application
}

var sandboxInst Sandbox

//func initSandbox(config Config) {
//	sandboxInst = NewSandbox(config)
//}

func NewSandbox(application Application) Sandbox {
	return &sandbox{
		Config: application.Config(),
	}
}

func (s *sandbox) GetKey() string {
	if v, e := SandboxSignKey(s.Config); e == nil {
		return string(v)
	}
	return ""
}
