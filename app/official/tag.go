package official

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/log"

	"github.com/godcong/wego/util"
)

/*Tag Tag */
type Tag struct {
	*Account
}

func newTag(acc *Account) *Tag {
	return &Tag{
		Account: acc,
	}
}

/*NewTag NewTag*/
func NewTag(config *core.Config) *Tag {
	return newTag(NewOfficialAccount(config))
}

//Create 创建标签
// http请求方式:POST（请使用https协议）
// https://api.weixin.qq.com/cgi-bin/tags/create?access_token=ACCESS_TOKEN
// 成功:
// {"tag":{"id":100,"name":"testtag"}}
func (t *Tag) Create(name string) core.Response {
	log.Debug("Tag|Create", name)
	p := t.accessToken.GetToken().KeyMap()
	resp := core.PostJSON(
		Link(tagsCreateURLSuffix),
		p,
		util.Map{
			"tag": util.Map{"name": name},
		})
	return resp
}

//Get 获取公众号已创建的标签
// http请求方式:GET（请使用https协议）
// https://api.weixin.qq.com/cgi-bin/tags/get?access_token=ACCESS_TOKEN
// 成功:
// {"tags":[{"id":2,"name":"星标组","count":0},{"id":100,"name":"testtag","count":0}]}
func (t *Tag) Get() core.Response {
	log.Debug("Tag|Get")
	p := t.accessToken.GetToken().KeyMap()
	resp := core.Get(
		Link(tagsGetURLSuffix),
		p,
	)
	return resp
}

//Update 编辑标签
// http请求方式:POST（请使用https协议）
// https://api.weixin.qq.com/cgi-bin/tags/update?access_token=ACCESS_TOKEN
// 成功:
// {"errcode":0,"errmsg":"ok"}
func (t *Tag) Update(id int, name string) core.Response {
	log.Debug("Tag|Update", id, name)
	p := t.accessToken.GetToken().KeyMap()
	resp := core.PostJSON(
		Link(tagsUpdateURLSuffix),
		p,
		util.Map{
			"tag": util.Map{"id": id, "name": name},
		})
	return resp
}

//Delete 删除标签
// http请求方式:POST（请使用https协议）
// https://api.weixin.qq.com/cgi-bin/tags/delete?access_token=ACCESS_TOKEN
// 成功:
// {"errcode":0,"errmsg":"ok"}
// 失败:
// {"errcode":45058,"errmsg":"can't modify sys tag hint: [eOA5oa07591527]"}
func (t *Tag) Delete(id int) core.Response {
	log.Debug("Tag|Update", id)
	p := t.accessToken.GetToken().KeyMap()
	resp := core.PostJSON(
		Link(tagsDeleteURLSuffix),
		p,
		util.Map{
			"tag": util.Map{"id": id},
		})
	return resp
}

//UserTagGet 获取标签下粉丝列表
// http请求方式:GET（请使用https协议）
// https://api.weixin.qq.com/cgi-bin/user/tag/get?access_token=ACCESS_TOKEN
// 成功:
// {"count":5,"data":{"openid":["oLyBi0tDnybg0WFkhKsn5HRetX1I","oLyBi0lCK5rQPuo0_cHJrjQ4J9XE","oLyBi0sjcrB44VQeAY7oer9st874","oLyBi0i5qhS-eO1monY34_KKTbfY","oLyBi0hSYhggnD-kOIms0IzZFqrc"]},"next_openid":"oLyBi0hSYhggnD-kOIms0IzZFqrc"}
func (t *Tag) UserTagGet(id int, nextOpenid string) core.Response {
	log.Debug("Tag|Update", id, nextOpenid)
	params := util.Map{
		"tag": util.Map{"id": id},
	}
	if nextOpenid != "" {
		params.Set("next_openid", nextOpenid)
	}
	p := t.accessToken.GetToken().KeyMap()
	resp := core.PostJSON(
		Link(userTagGetURLSuffix),
		p,
		params)
	return resp
}

//MembersBatchTagging  批量为用户打标签
// http请求方式:POST（请使用https协议）
// https://api.weixin.qq.com/cgi-bin/tags/members/batchtagging?access_token=ACCESS_TOKEN
// 成功:
// {"errcode":0,"errmsg":"ok"}
func (t *Tag) MembersBatchTagging(id int, openids []string) core.Response {
	log.Debug("Tag|Update", id, openids)
	params := util.Map{
		"tagid": id,
	}
	if openids != nil {
		params.Set("openid_list", openids)
	}
	p := t.accessToken.GetToken().KeyMap()
	resp := core.PostJSON(
		Link(tagsMembersBatchTaggingURLSuffix),
		p,
		params)
	return resp
}

//MembersBatchUntagging 批量为用户取消标签
// http请求方式:POST（请使用https协议）
// https://api.weixin.qq.com/cgi-bin/tags/members/batchuntagging?access_token=ACCESS_TOKEN
// 成功:
// {"errcode":0,"errmsg":"ok"}
func (t *Tag) MembersBatchUntagging(id int, openids []string) core.Response {
	log.Debug("Tag|Update", id, openids)
	params := util.Map{
		"tagid": id,
	}
	if openids != nil {
		params.Set("openid_list", openids)
	}
	p := t.accessToken.GetToken().KeyMap()
	resp := core.PostJSON(
		Link(tagsMembersBatchUntaggingURLSuffix),
		p,
		params)
	return resp
}

//GetIDList 获取用户身上的标签列表
// http请求方式:POST（请使用https协议）
// https://api.weixin.qq.com/cgi-bin/tags/getidlist?access_token=ACCESS_TOKEN
// 成功:
// {"tagid_list":[101]}
func (t *Tag) GetIDList(openid string) core.Response {
	log.Debug("Tag|GetIDList", openid)
	params := util.Map{
		"openid": openid,
	}

	p := t.accessToken.GetToken().KeyMap()
	resp := core.PostJSON(
		Link(tagsGetIDListURLSuffix),
		p,
		params)
	return resp
}

//GetBlackList 获取公众号的黑名单列表
// http请求方式:POST（请使用https协议）
//https://api.weixin.qq.com/cgi-bin/tags/members/getblacklist?access_token=ACCESS_TOKEN
func (t *Tag) GetBlackList(beginOpenid string) core.Response {
	log.Debug("Tag|GetBlackList", beginOpenid)
	var params util.Map
	if beginOpenid != "" {
		params = util.Map{"begin_openid": beginOpenid}
	}

	p := t.accessToken.GetToken().KeyMap()
	resp := core.PostJSON(
		Link(tagsMembersGetBlackListURLSuffix),
		p,
		params)
	return resp
}

//BatchBlackList 拉黑用户
// http请求方式:POST（请使用https协议）
// https://api.weixin.qq.com/cgi-bin/tags/members/batchblacklist?access_token=ACCESS_TOKEN
func (t *Tag) BatchBlackList(openidList []string) core.Response {
	log.Debug("Tag|BatchBlackList", openidList)
	var params util.Map
	if l := len(openidList); l > 0 && l <= 20 {
		params = util.Map{"openid_list": openidList}
	}

	p := t.accessToken.GetToken().KeyMap()
	resp := core.PostJSON(
		Link(tagsMembersBatchBlackListURLSuffix),
		p,
		params)
	return resp
}

//BatchUnblackList 取消拉黑用户
// http请求方式:POST（请使用https协议）
// https://api.weixin.qq.com/cgi-bin/tags/members/batchunblacklist?access_token=ACCESS_TOKEN
func (t *Tag) BatchUnblackList(openidList []string) core.Response {
	log.Debug("Tag|BatchUnblackList", openidList)
	var params util.Map
	if l := len(openidList); l > 0 && l <= 20 {
		params = util.Map{"openid_list": openidList}
	}

	p := t.accessToken.GetToken().KeyMap()
	return core.PostJSON(
		Link(tagsMembersBatchUnblackListURLSuffix),
		p,
		params)
}
