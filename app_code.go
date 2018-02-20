package wego

type AppCode interface {
}

type appCode struct {
	Config
	mini MiniProgram
}

func NewAppCode(application Application, config Config) AppCode {
	return &appCode{
		Config: config,
		mini:   application.MiniProgram(),
	}
}

func (a *appCode) getStream(url string, m Map) []byte {
	return a.mini.Client().RequestRaw(url, nil, "post", map[string]Map{"json": m})
}
