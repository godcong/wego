package mini_test

import (
	"github.com/godcong/wego"
	"github.com/godcong/wego/app/mini"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
	"testing"
)

var cfg = wego.C(util.Map{
	"app_id":  "wx3c69535993f4651d",
	"secret":  "f8c7a2cf0c6ed44e2c719964bbe13b1e",
	"key":     "aTKnSUcTkbEnhwQNdutWkQxAjnhAz2jK",
	"aes_key": "DbWPitkfiWkhLwDPA48laxJojyiNqVwtK7R1ENPvEwC",
})

func TestAppCode_Get(t *testing.T) {
	resp := mini.NewAppCode(cfg).Get("https://mp.quick58.com")
	_ = core.SaveTo(resp, "d:/Get.jpg")
}

func TestAppCode_GetQrCode(t *testing.T) {
	resp := mini.NewAppCode(cfg).GetQrCode("https://mp.quick58.com", 430)
	_ = core.SaveTo(resp, "d:/GetQrCode.jpg")
}

func TestAppCode_GetUnlimit(t *testing.T) {
	resp := mini.NewAppCode(cfg).GetUnlimit("https://mp.quick58.com")
	_ = core.SaveTo(resp, "d:/GetUnlimit.jpg")
}

// TestAuth_Session ...
func TestAuth_Session(t *testing.T) {
	auth := mini.NewAuth(cfg)
	resp := auth.Session("0022IX8c1OPfgv0tOQ6c1tGZ8c12IX8E")
	t.Log(resp.ToMap())
}

func TestDataCube_DailyRetainInfo(t *testing.T) {
	cube := mini.NewDataCube(cfg)
	resp := cube.DailyRetainInfo("20181116", "20181116")
	t.Log(resp.ToMap())
}

func TestDataCube_DailyVisitTrend(t *testing.T) {
	cube := mini.NewDataCube(cfg)
	resp := cube.DailyVisitTrend("20181109", "20181109")
	t.Log(resp.ToMap())
}

func TestDataCube_MonthlyVisitTrend(t *testing.T) {
	cube := mini.NewDataCube(cfg)
	resp := cube.MonthlyVisitTrend("20181001", "20181031")
	t.Log(resp.ToMap())
}

func TestDataCube_MonthlyRetainInfo(t *testing.T) {
	cube := mini.NewDataCube(cfg)
	resp := cube.MonthlyRetainInfo("20181001", "20181031")
	t.Log(resp.ToMap())
}

func TestDataCube_SummaryTrend(t *testing.T) {
	cube := mini.NewDataCube(cfg)
	resp := cube.SummaryTrend("20181109", "20181109")
	t.Log(resp.ToMap())
}

func TestDataCube_UserPortrait(t *testing.T) {
	cube := mini.NewDataCube(cfg)
	resp := cube.UserPortrait("20181109", "20181115")
	t.Log(resp.ToMap())
}

//TODO:not through(未取得数据)
func TestDataCube_VisitDistribution(t *testing.T) {
	cube := mini.NewDataCube(cfg)
	resp := cube.VisitDistribution("20181107", "20181107")
	t.Log(resp.ToMap())
}

//TODO:not through(未取得数据)
func TestDataCube_VisitPage(t *testing.T) {
	cube := mini.NewDataCube(cfg)
	resp := cube.VisitPage("20181109", "20181109")
	t.Log(resp.ToMap())
}

func TestDataCube_WeeklyRetainInfo(t *testing.T) {
	cube := mini.NewDataCube(cfg)
	resp := cube.WeeklyRetainInfo("20181105", "20181111")
	t.Log(resp.ToMap())
}

func TestDataCube_WeeklyVisitTrend(t *testing.T) {
	cube := mini.NewDataCube(cfg)
	resp := cube.WeeklyVisitTrend("20181105", "20181111")
	t.Log(resp.ToMap())
}

// TestAuth_UserInfo ...
func TestAuth_UserInfo(t *testing.T) {
	auth := mini.NewAuth(wego.C(util.Map{
		"app_id": "wx3c69535993f4651d",
		"secret": "f8c7a2cf0c6ed44e2c719964bbe13b1e",
	}))
	resp := auth.UserInfo("002JXxze2ilgfB0zNmAe2Amsze2JXxzJ", "rCmWuMckRqkw33i+s+NCh32iPdO+yiPS/FWJInan6XUdnXROIC8vXm7clc5NlRMFjI1hPo59eWWeLeLyfZs5lzuzOHASH2VVnwwetAjwbt9KC9v8zWGAZfvlweQWlBtKpSNS0H9dc1bhXafuA763mRq0v01Uq/LAktVAcyd1l/2JCKPhosRSov9F8FTCTt4YL1S4NeYGcjPDb+Mgb9LeRleseMZuziZbKvs66XnPw2ARtrGsiU3uyB4/WZGKERMJll3eRmgYe98F+q4ey0VAz3+Ah5x5NHDfrmxFgm4t3U78VF9q7IB706ULUgMozXJlU5cjsuaVNROXpBmWT/3fHpL3XIWl6U/m7V9o8RiLmmxSSChGCpq2zMjPqj741Z1gKe0wuQ7RpKAWrd1Ui2tG23r6TCigYCE7cb4BEI/KRJkWP0LbfTG8S/9tvuX+xuSgd78qc5nXGqEpMz+FR+b0yC2UcBBup3HO9WZ/3Ut8BjA=", "rVJM6LaFd8PboQCHvwDelQ==")
	t.Log(string(resp))
}
