package mini_program

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/core/util"
)

type DataCube struct {
	core.Config
	*MiniProgram
}

func (d *DataCube) query(api, from, to string) []byte {
	params := util.Map{
		"begin_date": from,
		"end_date":   to,
	}
	m := d.GetClient().HttpPostJson(api, params, nil)
	return m.ToBytes()
}

func (d *DataCube) UserPortrait(from, to string) util.Map {
	json := d.query(d.client.Link(core.DATACUBE_USERPORTRAIT_URL_SUFFIX), from, to)
	return core.JsonToMap(json)
}

func (d *DataCube) SummaryTrend(from, to string) util.Map {
	json := d.query(d.client.Link(core.DATACUBE_DAILYSUMMARYTREND_URL_SUFFIX), from, to)
	return core.JsonToMap(json)
}

func (d *DataCube) DailyVisitTrend(from, to string) util.Map {
	json := d.query(d.client.Link(core.DATACUBE_DAILYVISITTREND_URL_SUFFIX), from, to)
	return core.JsonToMap(json)
}

func (d *DataCube) WeeklyVisitTrend(from, to string) util.Map {
	json := d.query(d.client.Link(core.DATACUBE_WEEKLYVISITTREND_URL_SUFFIX), from, to)
	return core.JsonToMap(json)
}
func (d *DataCube) MonthlyVisitTrend(from, to string) util.Map {
	json := d.query(d.client.Link(core.DATACUBE_MONTHLYVISITTREND_URL_SUFFIX), from, to)
	return core.JsonToMap(json)
}
func (d *DataCube) VisitDistribution(from, to string) util.Map {
	json := d.query(d.client.Link(core.DATACUBE_VISITDISTRIBUTION_URL_SUFFIX), from, to)
	return core.JsonToMap(json)
}

func (d *DataCube) DailyRetainInfo(from, to string) util.Map {
	json := d.query(d.client.Link(core.DATACUBE_DAILYRETAININFO_URL_SUFFIX), from, to)
	return core.JsonToMap(json)
}

func (d *DataCube) WeeklyRetainInfo(from, to string) util.Map {
	json := d.query(d.client.Link(core.DATACUBE_WEEKLYRETAININFO_URL_SUFFIX), from, to)
	return core.JsonToMap(json)
}

func (d *DataCube) MonthlyRetainInfo(from, to string) util.Map {
	json := d.query(d.client.Link(core.DATACUBE_MONTHLYRETAININFO_URL_SUFFIX), from, to)
	return core.JsonToMap(json)
}

func (d *DataCube) VisitPage(from, to string) util.Map {
	json := d.query(d.client.Link(core.DATACUBE_VISITPAGE_URL_SUFFIX), from, to)
	return core.JsonToMap(json)
}
