package official_account_test

import (
	"testing"
	"time"

	"github.com/godcong/wego/official_account"
)

var dc = official_account.NewDataCube()

func TestNewDataCube(t *testing.T) {
	t0, err := time.Parse("2006-01-02", "2018-03-20")
	if err != nil {
		t.Log(err)
	}
	t1, err := time.Parse("2006-01-02", "2018-03-20")
	if err != nil {
		t.Log(err)
	}
	testDataCube_GetUserCumulate(t, t0, t1)
	testDataCube_GetArticleSummary(t, t0, t1)
	testDataCube_GetArticleTotal(t, t0, t1)
	testDataCube_GetUserCumulate(t, t0, t1)
	testDataCube_GetUserRead(t, t0, t1)
	testDataCube_GetUserReadHour(t, t0, t1)
	testDataCube_GetUserShare(t, t0, t1)
	testDataCube_GetUserShareHour(t, t0, t1)
	testDataCube_GetUserSummary(t, t0, t1)
}

func testDataCube_GetUserCumulate(t *testing.T, t0, t1 time.Time) {
	resp := dc.GetUserCumulate(t0, t1)
	t.Log(resp.ToString())
}

func testDataCube_GetUserSummary(t *testing.T, t0, t1 time.Time) {
	resp := dc.GetUserSummary(t0, t1)
	t.Log(resp.ToString())
}

func testDataCube_GetArticleSummary(t *testing.T, t0, t1 time.Time) {
	resp := dc.GetArticleSummary(t0, t1)
	t.Log(resp.ToString())
}

func testDataCube_GetArticleTotal(t *testing.T, t0, t1 time.Time) {
	resp := dc.GetArticleTotal(t0, t1)
	t.Log(resp.ToString())
}

func testDataCube_GetUserRead(t *testing.T, t0, t1 time.Time) {
	resp := dc.GetUserRead(t0, t1)
	t.Log(resp.ToString())
}

func testDataCube_GetUserReadHour(t *testing.T, t0, t1 time.Time) {
	resp := dc.GetUserReadHour(t0, t1)
	t.Log(resp.ToString())
}

func testDataCube_GetUserShare(t *testing.T, t0, t1 time.Time) {
	resp := dc.GetUserShare(t0, t1)
	t.Log(resp.ToString())
}

func testDataCube_GetUserShareHour(t *testing.T, t0, t1 time.Time) {
	resp := dc.GetUserShareHour(t0, t1)
	t.Log(resp.ToString())
}
