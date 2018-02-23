package core

type Sandbox struct {
	Config
	//app Application
}

func (s *Sandbox) GetKey() string {
	return string(SandboxSignKey(s.Config))
}

func (s *Sandbox) GetCacheKey() string {
	return ""
}
