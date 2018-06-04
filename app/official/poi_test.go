package official_test

import (
	"testing"

	"github.com/godcong/wego/app/official"
)

var poi = official.NewPoi()

func TestPoi_Add(t *testing.T) {

	resp := poi.Add(&official.PoiBaseInfo{
		Sid:          "",
		BusinessName: "15个汉字或30个英文字符内",
		BranchName:   "不超过10个字，不能含有括号和特殊字符",
		Province:     "",
		City:         "",
		District:     "",
		Address:      "",
		Telephone:    "",
		Categories:   []string{},
		OffsetType:   0,
		Longitude:    0,
		Latitude:     0,
		PhotoList: []official.PoiPhotoURL{
			{
				PhotoURL: "url://",
			},
		},
		Recommend:    "",
		Special:      "",
		Introduction: "",
		OpenTime:     "",
		AvgPrice:     0,
	})
	t.Log(resp.ToString())
}

func TestPoi_Get(t *testing.T) {
	resp := poi.Get("12121321")
	t.Log(resp.ToString())
}
