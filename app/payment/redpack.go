package payment

import (
	"strconv"

	"github.com/godcong/wego/config"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/net"
	"github.com/godcong/wego/util"
)

type RedPack struct {
	config.Config
	*Payment
}

func (r *RedPack) Info(m util.Map) *net.Response {
	m.Set("appid", r.Config.Get("app_id"))
	m.Set("bill_type", "MCHT")
	return r.SafeRequest(GETHBINFO_URL_SUFFIX, m)

}

func (r *RedPack) SendNormal(m util.Map) *net.Response {
	m.Set("total_num", strconv.Itoa(1))
	m.Set("client_ip", core.GetServerIp())
	m.Set("wxappid", r.Config.Get("app_id"))
	return r.SafeRequest(SENDREDPACK_URL_SUFFIX, m)
}

func (r *RedPack) SendGroup(m util.Map) *net.Response {
	m.Set("amt_type", "ALL_RAND")
	m.Set("wxappid", r.Config.Get("app_id"))
	return r.SafeRequest(SENDGROUPREDPACK_URL_SUFFIX, m)
}