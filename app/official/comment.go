package official

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

/*Comment Comment*/
type Comment struct {
	*Account
}

func newComment(acc *Account) *Comment {
	return &Comment{
		Account: acc,
	}
}

/*
NewComment 新建Comment
*/
func NewComment(config *core.Config) *Comment {
	return newComment(NewOfficialAccount(config))
}

/*
Open 打开文章评论

 https 请求方式: POST
 https://api.weixin.qq.com/cgi-bin/comment/open?access_token=ACCESS_TOKEN

 失败:
  {"errcode":88000,"errmsg":"without comment privilege"}
*/
func (c *Comment) Open(id, index int) core.Responder {
	p := c.accessToken.GetToken().KeyMap()
	resp := core.PostJSON(
		Link(commentOpenURLSuffix),
		p,
		util.Map{
			"msg_data_id": id,
			"index":       index,
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
func (c *Comment) Close(id, index int) core.Responder {
	p := c.accessToken.GetToken().KeyMap()
	resp := core.PostJSON(
		Link(commentCloseURLSuffix),
		p,
		util.Map{
			"msg_data_id": id,
			"index":       index,
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
func (c *Comment) List(id, index, begin, count, typ int) core.Responder {
	p := c.accessToken.GetToken().KeyMap()
	resp := core.PostJSON(
		Link(commentListURLSuffix),
		p,
		util.Map{
			"msg_data_id": id,
			"index":       index,
			"begin":       begin,
			"count":       count,
			"type":        typ,
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
func (c *Comment) Markelect(id, index, userCommentID int) core.Responder {
	p := c.accessToken.GetToken().KeyMap()
	resp := core.PostJSON(
		Link(commentMarkelectURLSuffix),
		p,
		util.Map{
			"msg_data_id":     id,
			"index":           index,
			"user_comment_id": userCommentID,
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
func (c *Comment) Unmarkelect(id, index, userCommentID int) core.Responder {
	p := c.accessToken.GetToken().KeyMap()
	resp := core.PostJSON(
		Link(commentUnmarkelectURLSuffix),
		p,
		util.Map{
			"msg_data_id":     id,
			"index":           index,
			"user_comment_id": userCommentID,
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
func (c *Comment) Delete(id, index, userCommentID int) core.Responder {
	p := c.accessToken.GetToken().KeyMap()
	resp := core.PostJSON(
		Link(commentDeleteURLSuffix),
		p,
		util.Map{
			"msg_data_id":     id,
			"index":           index,
			"user_comment_id": userCommentID,
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
func (c *Comment) ReplyAdd(id, index, userCommentID int, content string) core.Responder {
	p := c.accessToken.GetToken().KeyMap()
	resp := core.PostJSON(
		Link(commentReplyAddURLSuffix),
		p,
		util.Map{
			"msg_data_id":     id,
			"index":           index,
			"user_comment_id": userCommentID,
			"content":         content,
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
func (c *Comment) ReplyDelete(id, index, userCommentID int) core.Responder {
	p := c.accessToken.GetToken().KeyMap()
	resp := core.PostJSON(
		Link(commentReplyDeleteURLSuffix),
		p,
		util.Map{
			"msg_data_id":     id,
			"index":           index,
			"user_comment_id": userCommentID,
		})
	return resp
}
