package open

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

/*Platform Platform */
type Platform struct {
	config *core.Config
	client *core.Client
	token  *core.AccessToken
	sub    util.Map
}

func newPlatform(config *core.Config) *Platform {
	client := core.NewClient()
	token := core.NewAccessToken(config, client)
	return &Platform{
		config: config,
		client: client,
		token:  token,
		sub:    util.Map{},
	}
}
