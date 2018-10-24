package official

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

/*Comment Comment*/
type Comment struct {
	*Account
}

func newComment(officialAccount *Account) *Comment {
	return &Comment{
		Account: officialAccount,
	}
}

/*
NewComment 新建Comment
*/
func NewComment(config *core.Config) *Comment {
	return newComment(NewAccount(config))
}

/*
Open 打开文章评论

 https 请求方式: POST
 https://api.weixin.qq.com/cgi-bin/comment/open?access_token=ACCESS_TOKEN

 失败:
  {"errcode":88000,"errmsg":"without comment privilege"}
*/
func (c *Comment) Open(id, index int) core.Response {
	p := c.token.GetToken().KeyMap()
	resp := c.client.PostJSON(
		core.Link(commentOpenURLSuffix),
		util.Map{
			"msg_data_id": id,
			"index":       index,
		},
		util.Map{
			core.DataTypeQuery: p,
		})
	return resp
}

/*
Close 关闭评论

 https 请求方式: POST
 https://api.weixin.qq.com/cgi-bin/comment/close?access_token=ACCESS_TOKEN

 失败:
 {"errcode":88000,"errmsg":"without comment privilege"}
*/
func (c *Comment) Close(id, index int) core.Response {
	p := c.token.GetToken().KeyMap()
	resp := c.client.PostJSON(
		core.Link(commentCloseURLSuffix),
		util.Map{
			"msg_data_id": id,
			"index":       index,
		},
		util.Map{
			core.DataTypeQuery: p,
		})
	return resp
}

/*
List 获取文章评论

 https 请求方式: POST
 https://api.weixin.qq.com/cgi-bin/comment/list?access_token=ACCESS_TOKEN

 失败:
 {"errcode":88000,"errmsg":"without comment privilege"}
*/
func (c *Comment) List(id, index, begin, count, typ int) core.Response {
	p := c.token.GetToken().KeyMap()
	resp := c.client.PostJSON(
		core.Link(commentListURLSuffix),
		util.Map{
			"msg_data_id": id,
			"index":       index,
			"begin":       begin,
			"count":       count,
			"type":        typ,
		},
		util.Map{
			core.DataTypeQuery: p,
		})
	return resp
}

/*
Markelect  将评论标记精选

 https 请求方式: POST
 https://api.weixin.qq.com/cgi-bin/comment/markelect?access_token=ACCESS_TOKEN

 参数	是否必须	类型	说明
 id	是	int	群发返回的msg_data_id
 index	否	int	多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文
 user_comment_id	是	int	用户评论id
*/
func (c *Comment) Markelect(id, index, userCommentID int) core.Response {
	p := c.token.GetToken().KeyMap()
	resp := c.client.PostJSON(
		core.Link(commentMarkelectURLSuffix),
		util.Map{
			"msg_data_id":     id,
			"index":           index,
			"user_comment_id": userCommentID,
		},
		util.Map{
			core.DataTypeQuery: p,
		})
	return resp
}

/*
Unmarkelect 将评论取消精选

 https 请求方式: POST
 https://api.weixin.qq.com/cgi-bin/comment/unmarkelect?access_token=ACCESS_TOKEN


参数	是否必须	类型	说明
id	是	int	群发返回的msg_data_id
index	否	int	多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文
user_comment_id	是	int	用户评论id
*/
func (c *Comment) Unmarkelect(id, index, userCommentID int) core.Response {
	p := c.token.GetToken().KeyMap()
	resp := c.client.PostJSON(
		core.Link(commentUnmarkelectURLSuffix),
		util.Map{
			"msg_data_id":     id,
			"index":           index,
			"user_comment_id": userCommentID,
		},
		util.Map{
			core.DataTypeQuery: p,
		})
	return resp
}

/*Delete 删除评论
https 请求方式: POST
https://api.weixin.qq.com/cgi-bin/comment/delete?access_token=ACCESS_TOKEN
参数	是否必须	类型	说明
id	是	int	群发返回的msg_data_id
index	否	int	多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文
user_comment_id	是	int	用户评论id
*/
func (c *Comment) Delete(id, index, userCommentID int) core.Response {
	p := c.token.GetToken().KeyMap()
	resp := c.client.PostJSON(
		core.Link(commentDeleteURLSuffix),
		util.Map{
			"msg_data_id":     id,
			"index":           index,
			"user_comment_id": userCommentID,
		},
		util.Map{
			core.DataTypeQuery: p,
		})
	return resp
}

/*
ReplyAdd 回复评论

 https 请求方式: POST
 https://api.weixin.qq.com/cgi-bin/comment/reply/add?access_token=ACCESS_TOKEN

 参数	是否必须	类型	说明
 id	是	int	群发返回的msg_data_id
 index	否	int	多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文
 user_comment_id	是	int	评论id
 content	是	string	回复内容
*/
func (c *Comment) ReplyAdd(id, index, userCommentID int, content string) core.Response {
	p := c.token.GetToken().KeyMap()
	resp := c.client.PostJSON(
		core.Link(commentReplyAddURLSuffix),
		util.Map{
			"msg_data_id":     id,
			"index":           index,
			"user_comment_id": userCommentID,
			"content":         content,
		},
		util.Map{
			core.DataTypeQuery: p,
		})
	return resp
}

/*
ReplyDelete 删除回复

 https 请求方式: POST
 https://api.weixin.qq.com/cgi-bin/comment/reply/delete?access_token=ACCESS_TOKEN

 参数	是否必须	类型	说明
 id	是	int	群发返回的msg_data_id
 index	否	int	多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文
 user_comment_id	是	int	评论id
*/
func (c *Comment) ReplyDelete(id, index, userCommentID int) core.Response {
	p := c.token.GetToken().KeyMap()
	resp := c.client.PostJSON(
		core.Link(commentReplyDeleteURLSuffix),
		util.Map{
			"msg_data_id":     id,
			"index":           index,
			"user_comment_id": userCommentID,
		},
		util.Map{
			core.DataTypeQuery: p,
		})
	return resp
}
