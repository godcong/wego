package official_test

import (
	"testing"

	"github.com/godcong/wego/app/official"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/core/media"
)

var material = official.NewMaterial(config)

func TestNewMaterial(t *testing.T) {

}

func TestMaterial_AddNews(t *testing.T) {

	var resp core.Response
	resp = material.AddNews([]*media.Article{
		{
			Title:            "name",
			ThumbMediaID:     "9fCk1Any5VcwmbJPzGztWMq3a1PsWv11KpgLTdM_YXgIlwdAUosdeSI_M6M7Qtwb",
			Author:           "cc",
			Digest:           "ab",
			ShowCoverPic:     "0",
			Content:          "bb",
			ContentSourceURL: "a",
		},
	})
	t.Log(string(resp.Bytes()))
}

func TestMaterial_AddMaterial(t *testing.T) {

	resp := material.AddMaterial("test.jpg", core.MediaTypeImage)
	t.Log(string(resp.Bytes()))
}

func TestMaterial_UploadVideo(t *testing.T) {

	var resp core.Response
	resp = material.UploadVideo(`D:\temp\2.mp4`, "ceshi2", "only test")
	t.Log(string(resp.Bytes()))
}

func TestMaterial_Get(t *testing.T) {

	var resp core.Response
	// resp = material.Get("HIWcj9t3AI_b8qCQSu8lrTgTis9nPHNyIkIEWaDdHzY")
	resp = material.Get("HIWcj9t3AI_b8qCQSu8lrY5DkGL1LMl8_eTrDv4aUo8")
	t.Log(string(resp.Bytes()))
}

func TestMaterial_Del(t *testing.T) {

	var resp core.Response
	// resp = material.Get("HIWcj9t3AI_b8qCQSu8lrTgTis9nPHNyIkIEWaDdHzY")
	resp = material.Del("HIWcj9t3AI_b8qCQSu8lrY5DkGL1LMl8_eTrDv4aUo8")
	t.Log(string(resp.Bytes()))
}

func TestMaterial_UpdateNews(t *testing.T) {

	var resp core.Response
	// resp = material.Get("HIWcj9t3AI_b8qCQSu8lrTgTis9nPHNyIkIEWaDdHzY")
	resp = material.UpdateNews("9fCk1Any5VcwmbJPzGztWMq3a1PsWv11KpgLTdM_YXgIlwdAUosdeSI_M6M7Qtwb", 0, []*media.Article{})
	t.Log(string(resp.Bytes()))
}

func TestMaterial_GetMaterialCount(t *testing.T) {

	var resp core.Response
	// resp = material.Get("HIWcj9t3AI_b8qCQSu8lrTgTis9nPHNyIkIEWaDdHzY")
	resp = material.GetMaterialCount()
	t.Log(string(resp.Bytes()))

}

func TestMaterial_BatchGet(t *testing.T) {

	var resp core.Response
	// resp = material.Get("HIWcj9t3AI_b8qCQSu8lrTgTis9nPHNyIkIEWaDdHzY")
	resp = material.BatchGet(core.MediaTypeVideo, 1, 1)
	t.Log(string(resp.Bytes()))

}
