package open

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

const accessTokenKey = "component_access_token"
const accessTokenURLSuffix = "/cgi-bin/component/api_component_token"

func newAccessToken(p util.Map) *core.AccessToken {
	token := &core.AccessToken{
		URL:      accessTokenURLSuffix,
		TokenKey: accessTokenKey,
	}
	token.SetCredentials(p)

	return token
}
