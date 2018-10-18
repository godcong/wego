package core

/*InSandbox 是否沙箱环境 */
//func (a *Application) InSandbox() bool {
//	//c := a.Get("config").(Config)
//	//return a.GetBool("payment.default.sandbox")
//
//	return false
//}

/*GetKey 获取沙箱key */
//func (a *Application) GetKey(s string) string {
//	sb := Sandbox{}
//	b := a.Get("sandbox", &sb)
//	if b && a.InSandbox() {
//		sb.GetKey()
//	}
//	m := sync.Map
//	return sb.Get("aes_key")
//
//}

/*Scheme 获取微信Scheme */
//func (a *Application) Scheme(id string) string {
//	cfg := a.GetConfig().GetSubConfig("official_account.default") //TODO: get used config
//	m := make(util.Map)
//	m.Set("appid", cfg.Get("app_id"))
//	m.Set("mch_id", cfg.Get("mch_id"))
//	m.Set("time_stamp", util.Time())
//	m.Set("nonce_str", util.GenerateNonceStr())
//	m.Set("product_id", id)
//	m.Set("sign", GenerateSignature(m, cfg.Get("aes_key"), MakeSignMD5))
//	return BizPayURL + m.URLEncode()
//}

//func (a *Application) HandleNotify(typ string, f func(interface{})) {
//
//}

/*SetSubMerchant 设置子商户id */
//func (a *Application) SetSubMerchant(mchid, appid string) *Application {
//	a.Config.Set("sub_mch_id", mchid)
//	a.Config.Set("sub_appid", appid)
//	return a
//}

/*CacheOn turn on cache */
//func (s *System) CacheOn() {
//	s.UseCache = true
//}
//
///*CacheOff turn off cache */
//func (s *System) CacheOff() {
//	s.UseCache = false
//}
//
///*CacheStatus return cache status */
//func (s *System) CacheStatus() bool {
//	return s.UseCache
//}
