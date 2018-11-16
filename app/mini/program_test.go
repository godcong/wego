package mini_test

import (
	"github.com/godcong/wego"
	"github.com/godcong/wego/app/mini"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
	"testing"
)

var cfg = wego.C(util.Map{
	"app_id":  "wx1ad61aeef1903b93",
	"secret":  "c96956c2fd5ce7bfd7a0db1f7679ff6d",
	"key":     "O9aVVkSTmgJK4qSibhSYpGQzRbZ2NQSJ",
	"aes_key": "DbWPitkfiWkhLwDPA48laxJojyiNqVwtK7R1ENPvEwC",
})

func TestAppCode_Get(t *testing.T) {
	obj := mini.NewAppCode(cfg)
	resp := obj.Get("https://mp.quick58.com")
	_ = core.SaveTo(resp, "d:/Get.jpg")
}

func TestAppCode_GetQrCode(t *testing.T) {
	obj := mini.NewAppCode(cfg)
	resp := obj.GetQrCode("https://mp.quick58.com", 430)
	_ = core.SaveTo(resp, "d:/GetQrCode.jpg")
}

func TestAppCode_GetUnlimit(t *testing.T) {
	obj := mini.NewAppCode(cfg)
	resp := obj.GetUnlimit("https://mp.quick58.com")
	_ = core.SaveTo(resp, "d:/GetUnlimit.jpg")
}

// TestAuth_Session ...
func TestAuth_Session(t *testing.T) {
	obj := mini.NewAuth(cfg)
	resp := obj.Session("0022IX8c1OPfgv0tOQ6c1tGZ8c12IX8E")
	t.Log(resp.ToMap())
}

func TestDataCube_DailyRetainInfo(t *testing.T) {
	obj := mini.NewDataCube(cfg)
	resp := obj.DailyRetainInfo("20181116", "20181116")
	t.Log(resp.ToMap())
}

func TestDataCube_DailyVisitTrend(t *testing.T) {
	obj := mini.NewDataCube(cfg)
	resp := obj.DailyVisitTrend("20181109", "20181109")
	t.Log(resp.ToMap())
}

func TestDataCube_MonthlyVisitTrend(t *testing.T) {
	obj := mini.NewDataCube(cfg)
	resp := obj.MonthlyVisitTrend("20181001", "20181031")
	t.Log(resp.ToMap())
}

func TestDataCube_MonthlyRetainInfo(t *testing.T) {
	obj := mini.NewDataCube(cfg)
	resp := obj.MonthlyRetainInfo("20181001", "20181031")
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

//TODO:not through(未取得正常返回数据)
func TestDataCube_VisitDistribution(t *testing.T) {
	obj := mini.NewDataCube(cfg)
	resp := obj.VisitDistribution("20181107", "20181107")
	t.Log(resp.ToMap())
}

//TODO:not through(未取得正常返回数据)
func TestDataCube_VisitPage(t *testing.T) {
	obj := mini.NewDataCube(cfg)
	resp := obj.VisitPage("20181109", "20181109")
	t.Log(resp.ToMap())
}

func TestDataCube_WeeklyRetainInfo(t *testing.T) {
	obj := mini.NewDataCube(cfg)
	resp := obj.WeeklyRetainInfo("20181105", "20181111")
	t.Log(resp.ToMap())
}

func TestDataCube_WeeklyVisitTrend(t *testing.T) {
	obj := mini.NewDataCube(cfg)
	resp := obj.WeeklyVisitTrend("20181105", "20181111")
	t.Log(resp.ToMap())
}

func TestPlugin_List(t *testing.T) {
	obj := mini.NewPlugin(cfg)
	resp := obj.List()
	t.Log(resp.ToMap())
}

//TODO:not through(未取得正常返回数据)
func TestPlugin_Apply(t *testing.T) {
	obj := mini.NewPlugin(cfg)
	resp := obj.Apply("wx1ad61aeef1903b93")
	t.Log(resp.ToMap())
}

//TODO:not through(未取得正常返回数据)
func TestPlugin_DevAgree(t *testing.T) {
	obj := mini.NewPlugin(cfg)
	resp := obj.DevAgree("wx1ad61aeef1903b93")
	t.Log(resp.ToMap())
}

//TODO:not through(未取得正常返回数据)
func TestPlugin_Unbind(t *testing.T) {
	obj := mini.NewPlugin(cfg)
	resp := obj.Unbind("wx1ad61aeef1903b93")
	t.Log(resp.ToMap())
}

func TestPlugin_DevApplyList(t *testing.T) {
	obj := mini.NewPlugin(cfg)
	resp := obj.DevApplyList("wx1ad61aeef1903b93", 0, 0)
	t.Log(resp.ToMap())
}

func TestPlugin_DevDelete(t *testing.T) {
	obj := mini.NewPlugin(cfg)
	resp := obj.DevApplyList("wx1ad61aeef1903b93", 0, 0)
	t.Log(resp.ToMap())
}

func TestPlugin_DevRefuse(t *testing.T) {
	obj := mini.NewPlugin(cfg)
	resp := obj.DevApplyList("wx1ad61aeef1903b93", 0, 0)
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
