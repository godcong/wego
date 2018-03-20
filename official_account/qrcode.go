package official_account

import "github.com/godcong/wego/core"

type QrCode struct {
	core.Config
	*OfficialAccount
}

//http请求方式: POST
//URL: https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=TOKEN
//POST数据格式：json
//POST数据例子：{"expire_seconds": 604800, "action_name": "QR_SCENE", "action_info": {"scene": {"scene_id": 123}}}
//
//或者也可以使用以下POST数据创建字符串形式的二维码参数：
//{"expire_seconds": 604800, "action_name": "QR_STR_SCENE", "action_info": {"scene": {"scene_str": "test"}}}
func (q *QrCode) Create() {

}
