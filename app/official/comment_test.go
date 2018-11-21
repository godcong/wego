package official_test

import (
	"testing"

	"github.com/godcong/wego/app/official"
)

func TestComment_Open(t *testing.T) {
	c := official.NewComment(config)
	resp := c.Open(0, 1)
	t.Log(string(resp.Bytes()))
	t.Log(resp.Error())
}

func TestComment_Close(t *testing.T) {
	c := official.NewComment(config)
	resp := c.Close(0, 0)
	t.Log(string(resp.Bytes()))
	t.Log(resp.Error())
}

func TestComment_List(t *testing.T) {
	c := official.NewComment(config)
	resp := c.List(0, 0, 0, 0, 0)
	t.Log(string(resp.Bytes()))
	t.Log(resp.Error())
}

func TestComment_Marketlect(t *testing.T) {
	c := official.NewComment(config)
	resp := c.MarkElect(0, 1, 0)
	t.Log(string(resp.Bytes()))
	t.Log(resp.Error())
}

func TestComment_Unmarkelect(t *testing.T) {
	c := official.NewComment(config)
	resp := c.UnmarkElect(0, 1, 0)
	t.Log(string(resp.Bytes()))
	t.Log(resp.Error())
}

func TestComment_Delete(t *testing.T) {
	c := official.NewComment(config)
	resp := c.Delete(0, 1, 0)
	t.Log(string(resp.Bytes()))
	t.Log(resp.Error())
}

func TestComment_ReplyAdd(t *testing.T) {
	c := official.NewComment(config)
	resp := c.ReplyAdd(0, 1, 0, "content na si")
	t.Log(string(resp.Bytes()))
	t.Log(resp.Error())
}

func TestComment_ReplyDelete(t *testing.T) {
	c := official.NewComment(config)
	resp := c.ReplyDelete(0, 1, 0)
	t.Log(string(resp.Bytes()))
	t.Log(resp.Error())
}

func TestNewComment(t *testing.T) {
	// do what?
}
