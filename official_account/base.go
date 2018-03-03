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

//成功：
//{
//"ip_list": [
//"127.0.0.1",
//"127.0.0.2",
//"101.226.103.0/25"
//]
//}
//失败:
//{"errcode":40013,"errmsg":"invalid appid"}
func (b *Base) GetCallbackIp() core.Map {
	token := b.GetToken()

	resp := b.HttpPostJson(b.Link(GETCALLBACKIP_URL_SUFFIX), nil, core.Map{core.REQUEST_TYPE_QUERY.String(): core.Map{"access_token": token.GetKey()}})
	return resp.ToMap()
}
