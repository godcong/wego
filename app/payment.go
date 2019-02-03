package app

import (
	"crypto/md5"
	"fmt"
	"github.com/godcong/wego/cache"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

//Payment ...
type Payment struct {
	DomainHost  string
	option      *PaymentOption
	property    *PaymentProperty
	sandbox     *SandboxProperty
	accessToken *AccessToken
}

// PaymentOption ...
type PaymentOption struct {
	UsePayment bool
	Host       string
}

// NewPayment ...
func NewPayment(property *Property, opts ...*PaymentOption) *Payment {
	var opt *PaymentOption
	if opts != nil {
		opt = opts[0]
	}
	return &Payment{
		DomainHost:  BaseDomain,
		option:      opt,
		property:    property.Payment,
		sandbox:     property.Sandbox,
		accessToken: nil,
	}
}

// IsSandbox ...
func (p *Payment) IsSandbox() bool {
	if p.option != nil {
		return p.option.UsePayment
	}
	return false
}

/*GetKey 沙箱key(string类型) */
func (p *Payment) GetKey() string {
	key := cache.Get(p.getCacheKey())
	if key != nil {
		return key.(string)
	}

	response := p.SandboxSignKey().ToMap()

	if response.GetString("return_code") == "SUCCESS" {
		key := response.GetString("sandbox_signkey")
		cache.SetWithTTL(p.getCacheKey(), key, 3*24*3600)
		return key
	}
	return ""

}

func (p *Payment) getCacheKey() string {
	name := p.sandbox.AppID + "." + p.sandbox.MchID
	return "godcong.wego.payment.sandbox." + fmt.Sprintf("%x", md5.Sum([]byte(name)))
}

/*SandboxSignKey 沙箱key */
func (p *Payment) SandboxSignKey() core.Responder {
	m := make(util.Map)
	m.Set("mch_id", p.sandbox.MchID)
	m.Set("nonce_str", util.GenerateNonceStr())
	sign := util.GenerateSignature(m, p.sandbox.Key, util.MakeSignMD5)
	m.Set("sign", sign)
	resp := core.PostXML(util.URL(p.Host(), sandboxSignKeyURLSuffix), nil, m)

	return resp

}

// Host ...
func (p *Payment) Host() string {
	if p.option.Host != "" {
		return p.option.Host
	}
	return p.DomainHost
}
