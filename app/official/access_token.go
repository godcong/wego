package official

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

const accessTokenKey = "access_token"
const accessTokenURLSuffix = "/cgi-bin/token"

func newAccessToken(config *core.Config) *core.AccessToken {
	//client := NewClient(config)
	return &core.AccessToken{
		URL:      accessTokenURLSuffix,
		TokenKey: accessTokenKey,
		Credentials: util.Map{
			"grant_type": "client_credential",
			"appid":      config.GetString("app_id"),
			"secret":     config.GetString("secret"),
		},
	}

}
