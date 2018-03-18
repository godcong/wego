package official_account

import "github.com/godcong/wego/core"

type Comment struct {
	config core.Config
	*OfficialAccount
}

func newComment(officialAccount *OfficialAccount) *Comment {
	return &Comment{
		config:          defaultConfig,
		OfficialAccount: officialAccount,
	}
}

func NewComment() *Comment {
	return newComment(account)
}

// https 请求方式: POST
// https://api.weixin.qq.com/cgi-bin/comment/open?access_token=ACCESS_TOKEN
// 失败:
//  {"errcode":88000,"errmsg":"without comment privilege"}
func (c *Comment) Open(id, index int) *core.Response {
	p := c.token.GetToken().KeyMap()
	resp := c.client.HttpPostJson(
		c.client.Link(COMMENT_OPEN_URL_SUFFIX),
		core.Map{
			"msg_data_id": id,
			"index":       index,
		},
		core.Map{
			core.REQUEST_TYPE_QUERY.String(): p,
		})
	return resp
}

// https 请求方式: POST
// https://api.weixin.qq.com/cgi-bin/comment/close?access_token=ACCESS_TOKEN
// 失败:
//  {"errcode":88000,"errmsg":"without comment privilege"}
func (c *Comment) Close(id, index int) *core.Response {
	p := c.token.GetToken().KeyMap()
	resp := c.client.HttpPostJson(
		c.client.Link(COMMENT_CLOSE_URL_SUFFIX),
		core.Map{
			"msg_data_id": id,
			"index":       index,
		},
		core.Map{
			core.REQUEST_TYPE_QUERY.String(): p,
		})
	return resp
}

// https 请求方式: POST
// https://api.weixin.qq.com/cgi-bin/comment/list?access_token=ACCESS_TOKEN
// 失败:
//  {"errcode":88000,"errmsg":"without comment privilege"}
func (c *Comment) List(id, index, begin, count, typ int) *core.Response {
	p := c.token.GetToken().KeyMap()
	resp := c.client.HttpPostJson(
		c.client.Link(COMMENT_LIST_URL_SUFFIX),
		core.Map{
			"msg_data_id": id,
			"index":       index,
			"begin":       begin,
			"count":       count,
			"type":        typ,
		},
		core.Map{
			core.REQUEST_TYPE_QUERY.String(): p,
		})
	return resp
}

func (c *Comment) Markelect(id, index, userCommentId int) *core.Response {
	p := c.token.GetToken().KeyMap()
	resp := c.client.HttpPostJson(
		c.client.Link(COMMENT_MARKELECT_URL_SUFFIX),
		core.Map{
			"msg_data_id":     id,
			"index":           index,
			"user_comment_id": userCommentId,
		},
		core.Map{
			core.REQUEST_TYPE_QUERY.String(): p,
		})
	return resp
}

func (c *Comment) Unmarkelect(id, index, userCommentId int) *core.Response {
	p := c.token.GetToken().KeyMap()
	resp := c.client.HttpPostJson(
		c.client.Link(COMMENT_UNMARKELECT_URL_SUFFIX),
		core.Map{
			"msg_data_id":     id,
			"index":           index,
			"user_comment_id": userCommentId,
		},
		core.Map{
			core.REQUEST_TYPE_QUERY.String(): p,
		})
	return resp
}

// https 请求方式: POST
// https://api.weixin.qq.com/cgi-bin/comment/delete?access_token=ACCESS_TOKEN
func (c *Comment) Delete(id, index, userCommentId int) *core.Response {
	p := c.token.GetToken().KeyMap()
	resp := c.client.HttpPostJson(
		c.client.Link(COMMENT_DELETE_URL_SUFFIX),
		core.Map{
			"msg_data_id":     id,
			"index":           index,
			"user_comment_id": userCommentId,
		},
		core.Map{
			core.REQUEST_TYPE_QUERY.String(): p,
		})
	return resp
}

// https 请求方式: POST
// https://api.weixin.qq.com/cgi-bin/comment/reply/add?access_token=ACCESS_TOKEN
func (c *Comment) ReplyAdd(id, index, userCommentId int, content string) *core.Response {
	p := c.token.GetToken().KeyMap()
	resp := c.client.HttpPostJson(
		c.client.Link(COMMENT_REPLY_ADD_URL_SUFFIX),
		core.Map{
			"msg_data_id":     id,
			"index":           index,
			"user_comment_id": userCommentId,
			"content":         content,
		},
		core.Map{
			core.REQUEST_TYPE_QUERY.String(): p,
		})
	return resp
}

// https 请求方式: POST
// https://api.weixin.qq.com/cgi-bin/comment/reply/delete?access_token=ACCESS_TOKEN
func (c *Comment) ReplyDelete(id, index, userCommentId int) *core.Response {
	p := c.token.GetToken().KeyMap()
	resp := c.client.HttpPostJson(
		c.client.Link(COMMENT_REPLY_DELETE_URL_SUFFIX),
		core.Map{
			"msg_data_id":     id,
			"index":           index,
			"user_comment_id": userCommentId,
		},
		core.Map{
			core.REQUEST_TYPE_QUERY.String(): p,
		})
	return resp
}
