package mini_program

import (
	"log"

	"github.com/godcong/wego/core"
)

type AppCode struct {
	core.Config
	*MiniProgram
}

func (a *AppCode) Get(path string, optionals core.Map) core.Map {
	params := core.Map{"path": path}
	params.Join(optionals)

	j := a.getStream(a.client.Link(core.GETWXACODE_URL_SUFFIX), params)

	return core.JsonToMap(j)
}

func (a *AppCode) GetQrCode(path string, width int) core.Map {
	params := core.Map{"path": path, "width": width}

	j := a.getStream(a.client.Link(core.CREATEWXAQRCODE_URL_SUFFIX), params)
	return core.JsonToMap(j)
}

func (a *AppCode) GetUnlimit(scene string, optionals core.Map) core.Map {
	params := core.Map{"scene": scene}
	params.Join(optionals)

	j := a.getStream(a.client.Link(core.GETWXACODEUNLIMIT_URL_SUFFIX), params)
	return core.JsonToMap(j)
}

func (a *AppCode) getStream(url string, m core.Map) []byte {
	log.Println(url, m)
	token0 := a.AccessToken().GetToken()
	token := token0.KeyMap()
	//strings.Join([]string{"access_token", token0.GetKey()}, "=")

	resp := a.GetClient().RequestRaw(url+"?"+token.UrlEncode(), core.Map{core.REQUEST_TYPE_QUERY.String(): token, core.REQUEST_TYPE_JSON.String(): m}, "post")
	panic(resp)
	return nil
}
