package payment

import (
	"strconv"

	"github.com/godcong/wego/core"
)

type RedPack struct {
	core.Config
	Payment
}

func (r *RedPack) Info(m core.Map) core.Map {
	m.Set("appid", r.Get("app_id"))
	m.Set("bill_type", "MCHT")
	return r.SafeRequest(core.GETHBINFO_URL_SUFFIX, m)

}

func (r *RedPack) SendNormal(m core.Map) core.Map {
	m.Set("total_num", strconv.Itoa(1))
	m.Set("client_ip", core.GetServerIp())
	m.Set("wxappid", r.Get("app_id"))
	return r.SafeRequest(core.SENDREDPACK_URL_SUFFIX, m)
}

func (r *RedPack) SendGroup(m core.Map) core.Map {
	m.Set("amt_type", "ALL_RAND")
	m.Set("wxappid", r.Get("app_id"))
	return r.SafeRequest(core.SENDGROUPREDPACK_URL_SUFFIX, m)
}
