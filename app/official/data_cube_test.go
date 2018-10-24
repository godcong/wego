package official_test

import (
	"testing"
	"time"

	"github.com/godcong/wego/app/official"
)

var dc = official.NewDataCube(config)

func TestNewDataCube(t *testing.T) {
	t0, err := time.Parse("2006-01-02", "2018-03-20")
	if err != nil {
		t.Log(err)
	}
	t1, err := time.Parse("2006-01-02", "2018-03-20")
	if err != nil {
		t.Log(err)
	}
	{
		// testDataCube_GetUserCumulate(t, t0, t1)
		// testDataCube_GetArticleSummary(t, t0, t1)
		// testDataCube_GetArticleTotal(t, t0, t1)
		// testDataCube_GetUserCumulate(t, t0, t1)
		// testDataCube_GetUserRead(t, t0, t1)
		// testDataCube_GetUserReadHour(t, t0, t1)
		// testDataCube_GetUserShare(t, t0, t1)
		// testDataCube_GetUserShareHour(t, t0, t1)
		// testDataCube_GetUserSummary(t, t0, t1)
	}
	{
		// testDataCube_GetUpstreamMsg(t, t0, t1)
		// testDataCube_GetUpstreamMsgDist(t, t0, t1)
		// testDataCube_GetUpstreamMsgDistMonth(t, t0, t1)
		// testDataCube_GetUpstreamMsgDistWeek(t, t0, t1)
		// testDataCube_GetUpstreamMsgHour(t, t0, t1)
		// testDataCube_GetUpstreamMsgMonth(t, t0, t1)
		// testDataCube_GetUpstreamMsgWeek(t, t0, t1)
	}
	{
		testDataCube_GetInterfaceSummary(t, t0, t1)
		testDataCube_GetInterfaceSummaryHour(t, t0, t1)
	}
}

func testDataCube_GetUpstreamMsg(t *testing.T, t0, t1 time.Time) {
	resp := dc.GetUpstreamMsg(t0, t1)
	t.Log(string(resp.Bytes()))
}

func testDataCube_GetUpstreamMsgDist(t *testing.T, t0, t1 time.Time) {
	resp := dc.GetUpstreamMsgDist(t0, t1)
	t.Log(string(resp.Bytes()))
}

func testDataCube_GetUpstreamMsgDistMonth(t *testing.T, t0, t1 time.Time) {
	resp := dc.GetUpstreamMsgDistMonth(t0, t1)
	t.Log(string(resp.Bytes()))
}

func testDataCube_GetUpstreamMsgDistWeek(t *testing.T, t0, t1 time.Time) {
	resp := dc.GetUpstreamMsgDistWeek(t0, t1)
	t.Log(string(resp.Bytes()))
}

func testDataCube_GetUpstreamMsgHour(t *testing.T, t0, t1 time.Time) {
	resp := dc.GetUpstreamMsgHour(t0, t1)
	t.Log(string(resp.Bytes()))
}

func testDataCube_GetUpstreamMsgMonth(t *testing.T, t0, t1 time.Time) {
	resp := dc.GetUpstreamMsgMonth(t0, t1)
	t.Log(string(resp.Bytes()))
}

func testDataCube_GetUpstreamMsgWeek(t *testing.T, t0, t1 time.Time) {
	resp := dc.GetUpstreamMsgWeek(t0, t1)
	t.Log(string(resp.Bytes()))
}

func testDataCube_GetInterfaceSummary(t *testing.T, t0, t1 time.Time) {
	resp := dc.GetInterfaceSummary(t0, t1)
	t.Log(string(resp.Bytes()))
}

func testDataCube_GetInterfaceSummaryHour(t *testing.T, t0, t1 time.Time) {
	resp := dc.GetInterfaceSummaryHour(t0, t1)
	t.Log(string(resp.Bytes()))
}

//
// func testDataCube_GetUserCumulate(t *testing.T, t0, t1 time.Time) {
// 	resp := dc.GetUserCumulate(t0, t1)
// 	t.Log(string(resp.Bytes()))
// }
//
// func testDataCube_GetUserSummary(t *testing.T, t0, t1 time.Time) {
// 	resp := dc.GetUserSummary(t0, t1)
// 	t.Log(string(resp.Bytes()))
// }
//
// func testDataCube_GetArticleSummary(t *testing.T, t0, t1 time.Time) {
// 	resp := dc.GetArticleSummary(t0, t1)
// 	t.Log(string(resp.Bytes()))
// }
//
// func testDataCube_GetArticleTotal(t *testing.T, t0, t1 time.Time) {
// 	resp := dc.GetArticleTotal(t0, t1)
// 	t.Log(string(resp.Bytes()))
// }
//
// func testDataCube_GetUserRead(t *testing.T, t0, t1 time.Time) {
// 	resp := dc.GetUserRead(t0, t1)
// 	t.Log(string(resp.Bytes()))
// }
//
// func testDataCube_GetUserReadHour(t *testing.T, t0, t1 time.Time) {
// 	resp := dc.GetUserReadHour(t0, t1)
// 	t.Log(string(resp.Bytes()))
// }
//
// func testDataCube_GetUserShare(t *testing.T, t0, t1 time.Time) {
// 	resp := dc.GetUserShare(t0, t1)
// 	t.Log(string(resp.Bytes()))
// }
//
// func testDataCube_GetUserShareHour(t *testing.T, t0, t1 time.Time) {
// 	resp := dc.GetUserShareHour(t0, t1)
// 	t.Log(string(resp.Bytes()))
// }
