package mini

import (
	"log"
	"strings"

	"github.com/godcong/wego/core"
)

type AppCode struct {
	core.Config
	*MiniProgram
}

func (a *AppCode) Get(path string, optionals core.Map) core.Map {
	params := core.Map{"path": path}
	params.Join(optionals)

	j := a.getStream(a.prefix(core.GETWXACODE_URL_SUFFIX), params)

	return core.JsonToMap(j)
}

func (a *AppCode) GetQrCode(path string, width int) core.Map {
	params := core.Map{"path": path, "width": width}

	j := a.getStream(a.prefix(core.CREATEWXAQRCODE_URL_SUFFIX), params)
	return core.JsonToMap(j)
}

func (a *AppCode) GetUnlimit(scene string, optionals core.Map) core.Map {
	params := core.Map{"scene": scene}
	params.Join(optionals)

	j := a.getStream(a.prefix(core.GETWXACODEUNLIMIT_URL_SUFFIX), params)
	return core.JsonToMap(j)
}

func (a *AppCode) getStream(url string, m core.Map) []byte {
	log.Println(url, m)
	token0 := a.AccessToken().GetToken()
	token := strings.Join([]string{"access_token", token0.GetKey()}, "=")

	return a.GetClient().RequestRaw(url+"?"+token, nil, "post", core.Map{"json": m})
}
