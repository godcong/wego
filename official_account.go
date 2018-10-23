package wego

//
///*Base 基础*/
//type Base interface {
//	GetCallbackIP() util.Map
//	ClearQuota() util.Map
//}
//
///*Menu 菜单*/
//type Menu interface {
//}
//
///*Server 服务器*/
//type Server interface {
//	RegisterCallback(sc core.MessageCallback, types ...message.MsgType)
//	ServeHTTP(w http.ResponseWriter, r *http.Request)
//}
//
///*OfficialAccount 公众号*/
//type OfficialAccount interface {
//	Base() Base
//	Menu() Menu
//	Server() Server
//}
//
///*GetOfficialAccount 获取公众号*/
//func GetOfficialAccount() OfficialAccount {
//	obj := GetApp().Get("official").(OfficialAccount)
//	log.Debug("GetOfficialAccount|obj:", obj)
//	return obj
//}
