package official_account

import (
	"github.com/godcong/wego/config"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/net"
	"github.com/godcong/wego/util"
)

type Tag struct {
	config.Config
	*OfficialAccount
}

func newTag(account *OfficialAccount) *Tag {
	return &Tag{
		Config:          defaultConfig,
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
func (t *Tag) Create(name string) *net.Response {
	log.Debug("Tag|Create", name)
	p := t.token.GetToken().KeyMap()
	resp := t.client.HttpPostJson(
		t.client.Link(TAGS_CREATE_URL_SUFFIX),
		util.Map{
			"tag": util.Map{"name": name},
		},
		util.Map{
			net.REQUEST_TYPE_QUERY.String(): p,
		})
	return resp
}

// http请求方式：GET（请使用https协议）
// https://api.weixin.qq.com/cgi-bin/tags/get?access_token=ACCESS_TOKEN
// 成功:
// {"tags":[{"id":2,"name":"星标组","count":0},{"id":100,"name":"testtag","count":0}]}
func (t *Tag) Get() *net.Response {
	log.Debug("Tag|Get")
	p := t.token.GetToken().KeyMap()
	resp := t.client.HttpGet(
		t.client.Link(TAGS_GET_URL_SUFFIX),
		util.Map{
			net.REQUEST_TYPE_QUERY.String(): p,
		})
	return resp
}

// http请求方式：POST（请使用https协议）
// https://api.weixin.qq.com/cgi-bin/tags/update?access_token=ACCESS_TOKEN
// 成功:
// {"errcode":0,"errmsg":"ok"}
func (t *Tag) Update(id int, name string) *net.Response {
	log.Debug("Tag|Update", id, name)
	p := t.token.GetToken().KeyMap()
	resp := t.client.HttpPostJson(
		t.client.Link(TAGS_UPDATE_URL_SUFFIX),
		util.Map{
			"tag": util.Map{"id": id, "name": name},
		},
		util.Map{
			net.REQUEST_TYPE_QUERY.String(): p,
		})
	return resp
}

// http请求方式：POST（请使用https协议）
// https://api.weixin.qq.com/cgi-bin/tags/delete?access_token=ACCESS_TOKEN
// 成功:
// {"errcode":0,"errmsg":"ok"}
// 失败:
// {"errcode":45058,"errmsg":"can't modify sys tag hint: [eOA5oa07591527]"}
func (t *Tag) Delete(id int) *net.Response {
	log.Debug("Tag|Update", id)
	p := t.token.GetToken().KeyMap()
	resp := t.client.HttpPostJson(
		t.client.Link(TAGS_DELETE_URL_SUFFIX),
		util.Map{
			"tag": util.Map{"id": id},
		},
		util.Map{
			net.REQUEST_TYPE_QUERY.String(): p,
		})
	return resp
}

// http请求方式：GET（请使用https协议）
// https://api.weixin.qq.com/cgi-bin/user/tag/get?access_token=ACCESS_TOKEN
// 成功:
// {"count":5,"data":{"openid":["oLyBi0tDnybg0WFkhKsn5HRetX1I","oLyBi0lCK5rQPuo0_cHJrjQ4J9XE","oLyBi0sjcrB44VQeAY7oer9st874","oLyBi0i5qhS-eO1monY34_KKTbfY","oLyBi0hSYhggnD-kOIms0IzZFqrc"]},"next_openid":"oLyBi0hSYhggnD-kOIms0IzZFqrc"}
func (t *Tag) UserTagGet(id int, nextOpenid string) *net.Response {
	log.Debug("Tag|Update", id, nextOpenid)
	params := util.Map{
		"tag": util.Map{"id": id},
	}
	if nextOpenid != "" {
		params.Set("next_openid", nextOpenid)
	}
	p := t.token.GetToken().KeyMap()
	resp := t.client.HttpPostJson(
		t.client.Link(USER_TAG_GET_URL_SUFFIX),
		params,
		util.Map{
			net.REQUEST_TYPE_QUERY.String(): p,
		})
	return resp
}

// http请求方式：POST（请使用https协议）
// https://api.weixin.qq.com/cgi-bin/tags/members/batchtagging?access_token=ACCESS_TOKEN
// 成功:
// {"errcode":0,"errmsg":"ok"}
func (t *Tag) MembersBatchTagging(id int, openids []string) *net.Response {
	log.Debug("Tag|Update", id, openids)
	params := util.Map{
		"tagid": id,
	}
	if openids != nil {
		params.Set("openid_list", openids)
	}
	p := t.token.GetToken().KeyMap()
	resp := t.client.HttpPostJson(
		t.client.Link(TAGS_MEMBERS_BATCHTAGGING_URL_SUFFIX),
		params,
		util.Map{
			net.REQUEST_TYPE_QUERY.String(): p,
		})
	return resp
}

// http请求方式：POST（请使用https协议）
// https://api.weixin.qq.com/cgi-bin/tags/members/batchuntagging?access_token=ACCESS_TOKEN
// 成功:
// {"errcode":0,"errmsg":"ok"}
func (t *Tag) MembersBatchUntagging(id int, openids []string) *net.Response {
	log.Debug("Tag|Update", id, openids)
	params := util.Map{
		"tagid": id,
	}
	if openids != nil {
		params.Set("openid_list", openids)
	}
	p := t.token.GetToken().KeyMap()
	resp := t.client.HttpPostJson(
		t.client.Link(TAGS_MEMBERS_BATCHUNTAGGING_URL_SUFFIX),
		params,
		util.Map{
			net.REQUEST_TYPE_QUERY.String(): p,
		})
	return resp
}

// http请求方式：POST（请使用https协议）
// https://api.weixin.qq.com/cgi-bin/tags/getidlist?access_token=ACCESS_TOKEN
// 成功:
// {"tagid_list":[101]}
func (t *Tag) GetIdList(openid string) *net.Response {
	log.Debug("Tag|GetIdList", openid)
	params := util.Map{
		"openid": openid,
	}

	p := t.token.GetToken().KeyMap()
	resp := t.client.HttpPostJson(
		t.client.Link(TAGS_GETIDLIST_URL_SUFFIX),
		params,
		util.Map{
			net.REQUEST_TYPE_QUERY.String(): p,
		})
	return resp
}

//http请求方式：POST（请使用https协议）
//https://api.weixin.qq.com/cgi-bin/tags/members/getblacklist?access_token=ACCESS_TOKEN
func (t *Tag) GetBlackList(beginOpenid string) *net.Response {
	log.Debug("Tag|GetBlackList", beginOpenid)
	var param util.Map
	if beginOpenid != "" {
		param = util.Map{"begin_openid": beginOpenid}
	}

	query := t.token.GetToken().KeyMap()
	resp := t.client.HttpPostJson(
		t.client.Link(TAGS_MEMBERS_GETBLACKLIST_URL_SUFFIX),
		param,
		util.Map{
			net.REQUEST_TYPE_QUERY.String(): query,
		})
	return resp
}

// http请求方式：POST（请使用https协议）
// https://api.weixin.qq.com/cgi-bin/tags/members/batchblacklist?access_token=ACCESS_TOKEN
// 成功:
// {"errcode":0,"errmsg":"ok"}
func (t *Tag) BatchBlackList(openidList []string) *net.Response {
	log.Debug("Tag|BatchBlackList", openidList)
	var param util.Map
	if l := len(openidList); l > 0 && l <= 20 {
		param = util.Map{"openid_list": openidList}
	}

	query := t.token.GetToken().KeyMap()
	resp := t.client.HttpPostJson(
		t.client.Link(TAGS_MEMBERS_BATCHBLACKLIST_URL_SUFFIX),
		param,
		util.Map{
			net.REQUEST_TYPE_QUERY.String(): query,
		})
	return resp
}

// http请求方式：POST（请使用https协议）
// https://api.weixin.qq.com/cgi-bin/tags/members/batchunblacklist?access_token=ACCESS_TOKEN
// 成功:
// {"errcode":0,"errmsg":"ok"}
func (t *Tag) BatchUnblackList(openidList []string) *net.Response {
	log.Debug("Tag|BatchUnblackList", openidList)
	var param util.Map
	if l := len(openidList); l > 0 && l <= 20 {
		param = util.Map{"openid_list": openidList}
	}

	query := t.token.GetToken().KeyMap()
	resp := t.client.HttpPostJson(
		t.client.Link(TAGS_MEMBERS_BATCHUNBLACKLIST_URL_SUFFIX),
		param,
		util.Map{
			net.REQUEST_TYPE_QUERY.String(): query,
		})
	return resp
}
