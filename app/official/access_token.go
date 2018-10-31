package official

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

type AccessToken struct {
	*core.AccessToken
}

func newAccessToken(config *core.Config) *core.AccessToken {
	token := core.AccessToken{
		URI:         "",
		TokenKey:    "",
		Credentials: nil,
	}
	return token.SetCredentials(util.Map{
		"grant_type": "client_credential",
		"appid":      config.GetString("app_id"),
		"secret":     config.GetString("secret"),
	})
}
