package mini_program

import (
	"log"

	"github.com/godcong/wego/config"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/net"
	"github.com/godcong/wego/util"
)

type AppCode struct {
	config.Config
	*MiniProgram
}

func (a *AppCode) Get(path string, optionals util.Map) util.Map {
	params := util.Map{"path": path}
	params.Join(optionals)

	j := a.getStream(a.client.Link(core.GETWXACODE_URL_SUFFIX), params)

	return util.JsonToMap(j)
}

func (a *AppCode) GetQrCode(path string, width int) util.Map {
	params := util.Map{"path": path, "width": width}

	j := a.getStream(a.client.Link(core.CREATEWXAQRCODE_URL_SUFFIX), params)
	return util.JsonToMap(j)
}

func (a *AppCode) GetUnlimit(scene string, optionals util.Map) util.Map {
	params := util.Map{"scene": scene}
	params.Join(optionals)

	j := a.getStream(a.client.Link(core.GETWXACODEUNLIMIT_URL_SUFFIX), params)
	return util.JsonToMap(j)
}

func (a *AppCode) getStream(url string, m util.Map) []byte {
	log.Println(url, m)
	token0 := a.AccessToken().GetToken()
	token := token0.KeyMap()
	//strings.Join([]string{"access_token", token0.GetKey()}, "=")

	resp := a.GetClient().RequestRaw(url+"?"+token.UrlEncode(), util.Map{net.REQUEST_TYPE_QUERY.String(): token, net.REQUEST_TYPE_JSON.String(): m}, "post")
	panic(resp)
	return nil
}
