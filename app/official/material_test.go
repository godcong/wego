package official_test

import (
	"testing"

	"github.com/godcong/wego/app/official"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/core/media"
	"github.com/godcong/wego/net"
)

func TestNewMaterial(t *testing.T) {

}

func TestMaterial_AddNews(t *testing.T) {
	material := official.NewMaterial()
	var resp *net.Response
	resp = material.AddNews([]*media.Article{
		{
			Title:            "name",
			ThumbMediaId:     "9fCk1Any5VcwmbJPzGztWMq3a1PsWv11KpgLTdM_YXgIlwdAUosdeSI_M6M7Qtwb",
			Author:           "cc",
			Digest:           "ab",
			ShowCoverPic:     "0",
			Content:          "bb",
			ContentSourceUrl: "a",
		},
	})
	t.Log(resp.ToString())
}

func TestMaterial_AddMaterial(t *testing.T) {
	material := official.NewMaterial()
	resp := material.AddMaterial("test.jpg", core.MediaTypeImage)
	t.Log(resp.ToString())
}

func TestMaterial_UploadVideo(t *testing.T) {
	material := official.NewMaterial()
	var resp *net.Response
	resp = material.UploadVideo(`D:\temp\2.mp4`, "ceshi2", "only test")
	t.Log(resp.ToString())
}

func TestMaterial_Get(t *testing.T) {
	material := official.NewMaterial()
	var resp *net.Response
	// resp = material.Get("HIWcj9t3AI_b8qCQSu8lrTgTis9nPHNyIkIEWaDdHzY")
	resp = material.Get("HIWcj9t3AI_b8qCQSu8lrY5DkGL1LMl8_eTrDv4aUo8")
	t.Log(resp.ToString())
}

func TestMaterial_Del(t *testing.T) {
	material := official.NewMaterial()
	var resp *net.Response
	// resp = material.Get("HIWcj9t3AI_b8qCQSu8lrTgTis9nPHNyIkIEWaDdHzY")
	resp = material.Del("HIWcj9t3AI_b8qCQSu8lrY5DkGL1LMl8_eTrDv4aUo8")
	t.Log(resp.ToString())
}

func TestMaterial_UpdateNews(t *testing.T) {
	material := official.NewMaterial()
	var resp *net.Response
	// resp = material.Get("HIWcj9t3AI_b8qCQSu8lrTgTis9nPHNyIkIEWaDdHzY")
	resp = material.UpdateNews("9fCk1Any5VcwmbJPzGztWMq3a1PsWv11KpgLTdM_YXgIlwdAUosdeSI_M6M7Qtwb", 0, []*media.Article{})
	t.Log(resp.ToString())
}

func TestMaterial_GetMaterialCount(t *testing.T) {
	material := official.NewMaterial()
	var resp *net.Response
	// resp = material.Get("HIWcj9t3AI_b8qCQSu8lrTgTis9nPHNyIkIEWaDdHzY")
	resp = material.GetMaterialCount()
	t.Log(resp.ToString())

}

func TestMaterial_BatchGet(t *testing.T) {
	material := official.NewMaterial()
	var resp *net.Response
	// resp = material.Get("HIWcj9t3AI_b8qCQSu8lrTgTis9nPHNyIkIEWaDdHzY")
	resp = material.BatchGet(core.MediaTypeVideo, 1, 1)
	t.Log(resp.ToString())

}
