package mini

import "github.com/godcong/wego/core"

type AppCode struct {
	core.Config
	*MiniProgram
}

func (a *AppCode) getStream(url string, m core.Map) []byte {
	return a.GetClient().RequestRaw(url, nil, "post", core.Map{"json": m})
}
