package official_account

import "github.com/godcong/wego/core"

type Base struct {
	core.Config
	core.Client
	core.AccessToken
	*OfficialAccount
}

func (b *Base) ClearQuota() core.Map {
	params := core.Map{
		"appid": b.Get("app_id"),
	}
	resp := b.HttpPostJson(b.Link(CLEAR_QUOTA_URL_SUFFIX), params, nil)
	return resp.ToMap()
}

func (b *Base) GetCallbackIp() core.Map {
	resp := b.HttpPostJson(b.Link(GETCALLBACKIP_URL_SUFFIX), nil, core.Map{core.REQUEST_TYPE_QUERY.String(): core.Map{"access_token": b.GetToken().GetKey()}})
	return resp.ToMap()
}
