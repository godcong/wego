package official_account

import "github.com/godcong/wego/core"

type Tag struct {
	config core.Config
	*OfficialAccount
}

func newTag(account *OfficialAccount) *Tag {
	return &Tag{
		config:          defaultConfig,
		OfficialAccount: account,
	}
}

func NewTag() *Tag {
	return newTag(account)
}

// http请求方式：POST（请使用https协议）
// https://api.weixin.qq.com/cgi-bin/tags/create?access_token=ACCESS_TOKEN
// 成功:
// {"tag":{"id":100,"name":"testtag"}}
func (t *Tag) Create(name string) *core.Response {
	core.Debug("Tag|Create", name)
	p := t.token.GetToken().KeyMap()
	resp := t.client.HttpPostJson(
		t.client.Link(TAGS_CREATE_URL_SUFFIX),
		core.Map{
			"tag": core.Map{"name": name},
		},
		core.Map{
			core.REQUEST_TYPE_QUERY.String(): p,
		})
	return resp
}

// http请求方式：GET（请使用https协议）
// https://api.weixin.qq.com/cgi-bin/tags/get?access_token=ACCESS_TOKEN
// 成功:
// {"tags":[{"id":2,"name":"星标组","count":0},{"id":100,"name":"testtag","count":0}]}
func (t *Tag) Get() *core.Response {
	core.Debug("Tag|Get")
	p := t.token.GetToken().KeyMap()
	resp := t.client.HttpGet(
		t.client.Link(TAGS_GET_URL_SUFFIX),
		core.Map{
			core.REQUEST_TYPE_QUERY.String(): p,
		})
	return resp
}

// http请求方式：POST（请使用https协议）
// https://api.weixin.qq.com/cgi-bin/tags/update?access_token=ACCESS_TOKEN
// 成功:
// {"errcode":0,"errmsg":"ok"}
func (t *Tag) Update(id int, name string) *core.Response {
	core.Debug("Tag|Update", id, name)
	p := t.token.GetToken().KeyMap()
	resp := t.client.HttpPostJson(
		t.client.Link(TAGS_UPDATE_URL_SUFFIX),
		core.Map{
			"tag": core.Map{"id": id, "name": name},
		},
		core.Map{
			core.REQUEST_TYPE_QUERY.String(): p,
		})
	return resp
}

// http请求方式：POST（请使用https协议）
// https://api.weixin.qq.com/cgi-bin/tags/delete?access_token=ACCESS_TOKEN
// 成功:
// {"errcode":0,"errmsg":"ok"}
// 失败:
// {"errcode":45058,"errmsg":"can't modify sys tag hint: [eOA5oa07591527]"}
func (t *Tag) Delete(id int) *core.Response {
	core.Debug("Tag|Update", id)
	p := t.token.GetToken().KeyMap()
	resp := t.client.HttpPostJson(
		t.client.Link(TAGS_DELETE_URL_SUFFIX),
		core.Map{
			"tag": core.Map{"id": id},
		},
		core.Map{
			core.REQUEST_TYPE_QUERY.String(): p,
		})
	return resp
}

// http请求方式：GET（请使用https协议）
// https://api.weixin.qq.com/cgi-bin/user/tag/get?access_token=ACCESS_TOKEN
// 成功:
// {"count":5,"data":{"openid":["oLyBi0tDnybg0WFkhKsn5HRetX1I","oLyBi0lCK5rQPuo0_cHJrjQ4J9XE","oLyBi0sjcrB44VQeAY7oer9st874","oLyBi0i5qhS-eO1monY34_KKTbfY","oLyBi0hSYhggnD-kOIms0IzZFqrc"]},"next_openid":"oLyBi0hSYhggnD-kOIms0IzZFqrc"}
func (t *Tag) UserTagGet(id int, nextOpenid string) *core.Response {
	core.Debug("Tag|Update", id, nextOpenid)
	params := core.Map{
		"tag": core.Map{"id": id},
	}
	if nextOpenid != "" {
		params.Set("next_openid", nextOpenid)
	}
	p := t.token.GetToken().KeyMap()
	resp := t.client.HttpPostJson(
		t.client.Link(USER_TAG_GET_URL_SUFFIX),
		params,
		core.Map{
			core.REQUEST_TYPE_QUERY.String(): p,
		})
	return resp
}

// http请求方式：POST（请使用https协议）
// https://api.weixin.qq.com/cgi-bin/tags/members/batchtagging?access_token=ACCESS_TOKEN
// 成功:
// {"errcode":0,"errmsg":"ok"}
func (t *Tag) MembersBatchTagging(id int, openids []string) *core.Response {
	core.Debug("Tag|Update", id, openids)
	params := core.Map{
		"tagid": id,
	}
	if openids != nil {
		params.Set("openid_list", openids)
	}
	p := t.token.GetToken().KeyMap()
	resp := t.client.HttpPostJson(
		t.client.Link(TAGS_MEMBERS_BATCHTAGGING_URL_SUFFIX),
		params,
		core.Map{
			core.REQUEST_TYPE_QUERY.String(): p,
		})
	return resp
}

// http请求方式：POST（请使用https协议）
// https://api.weixin.qq.com/cgi-bin/tags/members/batchuntagging?access_token=ACCESS_TOKEN
// 成功:
// {"errcode":0,"errmsg":"ok"}
func (t *Tag) MembersBatchUntagging(id int, openids []string) *core.Response {
	core.Debug("Tag|Update", id, openids)
	params := core.Map{
		"tagid": id,
	}
	if openids != nil {
		params.Set("openid_list", openids)
	}
	p := t.token.GetToken().KeyMap()
	resp := t.client.HttpPostJson(
		t.client.Link(TAGS_MEMBERS_BATCHUNTAGGING_URL_SUFFIX),
		params,
		core.Map{
			core.REQUEST_TYPE_QUERY.String(): p,
		})
	return resp
}

// http请求方式：POST（请使用https协议）
// https://api.weixin.qq.com/cgi-bin/tags/getidlist?access_token=ACCESS_TOKEN
// 成功:
// {"tagid_list":[101]}
func (t *Tag) GetIdList(openid string) *core.Response {
	core.Debug("Tag|GetIdList", openid)
	params := core.Map{
		"openid": openid,
	}

	p := t.token.GetToken().KeyMap()
	resp := t.client.HttpPostJson(
		t.client.Link(TAGS_GETIDLIST_URL_SUFFIX),
		params,
		core.Map{
			core.REQUEST_TYPE_QUERY.String(): p,
		})
	return resp
}
