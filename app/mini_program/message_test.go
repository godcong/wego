package mini_program

import (
	"testing"

	"github.com/godcong/wego/util"
)

var textMessage = util.Map{
	"touser":  "oE_gl0Yr54fUjBhU5nBlP4hS2efo",
	"msgtype": "text",
	"text":    util.Map{"content": "Hello World"},
}

var linkMessage = util.Map{
	"touser":  "oE_gl0Yr54fUjBhU5nBlP4hS2efo",
	"msgtype": "link",
	"link": util.Map{
		"title":       "最新文章",
		"description": "关注公众号获取最新文章更新",
		"url":         "https://mp.weixin.qq.com/s/qiaIpcQW9UB9Qe09cWc02w",
		"thumb_url":   "http://mmbiz.qpic.cn/mmbiz_png/Z68q4LGWAW3JFr2w1Lk8tWch9o8C4BUWwKoqRibXT5yR4UUW8FZmHia3TOPuhiaxCTJQLtXGYibSVLaTCFFGk1jgpw/640?tp=webp&wxfrom=5&wx_lazy=1",
	},
}

var mpMessage = util.Map{
	"touser":  "oE_gl0Yr54fUjBhU5nBlP4hS2efo",
	"msgtype": "miniprogrampage",
	"miniprogrampage": util.Map{
		"title":          "最新文章",
		"pagepath":       "https://mp.weixin.qq.com/s/qiaIpcQW9UB9Qe09cWc02w",
		"thumb_media_id": "LWOqgv64HBvdT_fjOzJLfsGydEGz6eRq2T6tZA2D2T2V9pGFOu8x_BF2xEXfWCmI",
	},
}

var imageMessage = util.Map{
	"touser":  "oE_gl0Yr54fUjBhU5nBlP4hS2efo",
	"msgtype": "image",
	"image":   util.Map{"media_id": "LWOqgv64HBvdT_fjOzJLfsGydEGz6eRq2T6tZA2D2T2V9pGFOu8x_BF2xEXfWCmI"},
}

func TestMessage_Send(t *testing.T) {
	msg := NewMessage()
	t.Log(msg.Send(linkMessage).ToString())
	t.Log(msg.Send(linkMessage).ToString())
}
