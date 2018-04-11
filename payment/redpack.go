package payment

import (
	"strconv"

	"github.com/godcong/wego/core"
)

type RedPack struct {
	core.Config
	*Payment
}

func (r *RedPack) Info(m core.Map) *core.Response {
	m.Set("appid", r.Config.Get("app_id"))
	m.Set("bill_type", "MCHT")
	return r.SafeRequest(GETHBINFO_URL_SUFFIX, m)

}

func (r *RedPack) SendNormal(m core.Map) *core.Response {
	m.Set("total_num", strconv.Itoa(1))
	m.Set("client_ip", core.GetServerIp())
	m.Set("wxappid", r.Config.Get("app_id"))
	return r.SafeRequest(SENDREDPACK_URL_SUFFIX, m)
}

func (r *RedPack) SendGroup(m core.Map) *core.Response {
	m.Set("amt_type", "ALL_RAND")
	m.Set("wxappid", r.Config.Get("app_id"))
	return r.SafeRequest(SENDGROUPREDPACK_URL_SUFFIX, m)
}
