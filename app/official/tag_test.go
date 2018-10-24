package official_test

import (
	"testing"

	"github.com/godcong/wego/app/official"
)

var t0 = official.NewTag(config)

func TestNewTag(t *testing.T) {
	resp := t0.Create("testtag")
	t.Log(string(resp.Bytes()))
}

func TestTag_Get(t *testing.T) {
	resp := t0.Get()
	t.Log(string(resp.Bytes()))
}

func TestTag_Update(t *testing.T) {
	resp := t0.Update(100, "changetag")
	t.Log(string(resp.Bytes()))
}

func TestTag_Delete(t *testing.T) {
	resp := t0.Delete(2)
	t.Log(string(resp.Bytes()))
}

func TestTag_UserTagGet(t *testing.T) {
	resp := t0.UserTagGet(2, "")
	t.Log(string(resp.Bytes()))
}

func TestTag_MembersBatchTagging(t *testing.T) {
	resp := t0.MembersBatchTagging(101, []string{"oLyBi0tDnybg0WFkhKsn5HRetX1I"})
	t.Log(string(resp.Bytes()))
}

func TestTag_MembersBatchUntagging(t *testing.T) {
	resp := t0.MembersBatchUntagging(101, []string{"oLyBi0tDnybg0WFkhKsn5HRetX1I"})
	t.Log(string(resp.Bytes()))
}

func TestTag_GetIdList(t *testing.T) {
	resp := t0.GetIDList("oLyBi0tDnybg0WFkhKsn5HRetX1I")
	t.Log(string(resp.Bytes()))
}

func TestTag_GetBlackList(t *testing.T) {
	resp := t0.GetBlackList("oLyBi0tDnybg0WFkhKsn5HRetX1I")
	t.Log(string(resp.Bytes()))
}

func TestTag_BatchBlackList(t *testing.T) {
	resp := t0.BatchBlackList([]string{"oLyBi0tDnybg0WFkhKsn5HRetX1I"})
	t.Log(string(resp.Bytes()))
}

func TestTag_BatchUnblackList(t *testing.T) {
	resp := t0.BatchUnblackList([]string{"oLyBi0tDnybg0WFkhKsn5HRetX1I"})
	t.Log(string(resp.Bytes()))
}
