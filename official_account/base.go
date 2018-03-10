package official_account

import "github.com/godcong/wego/core"

type Base struct {
	core.Config
	client *core.Client
	token  *core.AccessToken
	*OfficialAccount
}

func (b *Base) ClearQuota() core.Map {
	params := core.Map{
		"appid": b.Get("app_id"),
	}
	resp := b.client.HttpPostJson(b.client.Link(CLEAR_QUOTA_URL_SUFFIX), params, nil)
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
	token := b.token.GetToken()

	resp := b.client.HttpGet(b.client.Link(GETCALLBACKIP_URL_SUFFIX), core.Map{"access_token": token.GetKey()})

	return resp.ToMap()
}
