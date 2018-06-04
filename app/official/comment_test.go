package official_test

import (
	"testing"

	"github.com/godcong/wego/app/official"
)

var c = official.NewComment()

func TestComment_Open(t *testing.T) {
	resp := c.Open(0, 1)
	t.Log(resp.ToString())
}

func TestComment_Close(t *testing.T) {
	resp := c.Close(0, 0)
	t.Log(resp.ToString())
}

func TestComment_List(t *testing.T) {
	resp := c.List(0, 0, 0, 0, 0)
	t.Log(resp.ToString())
}

func TestComment_Markelect(t *testing.T) {
	resp := c.Markelect(0, 1, 0)
	t.Log(resp.ToString())
}

func TestComment_Unmarkelect(t *testing.T) {
	resp := c.Unmarkelect(0, 1, 0)
	t.Log(resp.ToString())
}

func TestComment_Delete(t *testing.T) {
	resp := c.Delete(0, 1, 0)
	t.Log(resp.ToString())
}

func TestComment_ReplyAdd(t *testing.T) {
	resp := c.ReplyAdd(0, 1, 0, "content na si")
	t.Log(resp.ToString())
}

func TestComment_ReplyDelete(t *testing.T) {
	resp := c.ReplyDelete(0, 1, 0)
	t.Log(resp.ToString())
}

func TestNewComment(t *testing.T) {
	// do what?
}
