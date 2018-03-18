package core

type UserInfo struct {
	City           string   `json:"city"`
	Country        string   `json:"country"`
	HeadImgUrl     string   `json:"headimgurl"`
	Language       string   `json:"language"`
	Nickname       string   `json:"nickname"`
	Openid         string   `json:"openid"`
	Privilege      []string `json:"privilege"`
	Province       string   `json:"province"`
	Sex            uint     `json:"sex"`
	Subscribe      int      `json:"subscribe"`
	SubscribeTime  uint32   `json:"subscribe_time"`
	UnionId        string   `json:"unionid"`
	Remark         string   `json:"remark"`
	GroupId        int      `json:"groupid"`
	TagIdList      []int    `json:"tagid_list"`
	SubscribeScene string   `json:"subscribe_scene"`
	QrScene        int      `json:"qr_scene"`
	QrSceneStr     string   `json:"qr_scene_str"`
}
