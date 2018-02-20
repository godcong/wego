package wego

import "strconv"

type RedPack interface {
	Info(Map) Map
	SendNormal(Map) Map
	SendGroup(Map) Map
}

type redPack struct {
	Config
	Payment
}

func NewRedPack(application Application, config Config) RedPack {
	return &redPack{
		Config:  config,
		Payment: application.Payment(),
	}
}

func (r *redPack) Info(m Map) Map {
	m.Set("appid", r.Get("app_id"))
	m.Set("bill_type", "MCHT")
	return r.SafeRequest(GETHBINFO_URL_SUFFIX, m)

}

func (r *redPack) SendNormal(m Map) Map {
	m.Set("total_num", strconv.Itoa(1))
	m.Set("client_ip", GetServerIp())
	m.Set("wxappid", r.Get("app_id"))
	return r.SafeRequest(SENDREDPACK_URL_SUFFIX, m)
}

func (r *redPack) SendGroup(m Map) Map {
	m.Set("amt_type", "ALL_RAND")
	m.Set("wxappid", r.Get("app_id"))
	return r.SafeRequest(SENDGROUPREDPACK_URL_SUFFIX, m)
}
