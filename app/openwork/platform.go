package openwork

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

/*Platform Platform*/
type Platform struct {
	*core.Config
	Sub         util.Map
	client      *core.Client
	accessToken *core.AccessToken
}

func (p *Platform) AccessToken() *core.AccessToken {
	return p.accessToken
}

func (p *Platform) SetAccessToken(accessToken *core.AccessToken) {
	p.accessToken = accessToken
}

func (p *Platform) Client() *core.Client {
	return p.client
}

func (p *Platform) SetClient(client *core.Client) {
	p.client = client
}

func neOpenPlatform(config *core.Config, p util.Map) *Platform {
	return &Platform{
		Config: config,
		Sub:    p,
	}
}

//NewOfficialPlatform return a official Platform
func NewOpenPlatform(config *core.Config, v ...interface{}) *Platform {
	client := core.ClientGet(v)
	accessToken := newAccessToken(util.Map{
		"grant_type": "client_credential",
		"appid":      config.GetString("app_id"),
		"secret":     config.GetString("secret"),
	})
	accessToken.SetClient(client)

	platform := neOpenPlatform(config, util.Map{})
	platform.SetClient(client)
	platform.SetAccessToken(accessToken)
	return platform
}
