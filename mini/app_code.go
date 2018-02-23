package mini

import "github.com/godcong/wego/core"

type AppCode struct {
	core.Config
	//mini MiniProgram
}

//
//func (a *appCode) getStream(url string, m core.Map) []byte {
//	return a.mini.Client().RequestRaw(url, nil, "post", core.Map{"json": m})
//}
