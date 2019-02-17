package core

/*WechatUserInfo WechatUserInfo */
type WechatUserInfo struct {
	City           string   `json:"city"`
	Country        string   `json:"country"`
	HeadImgURL     string   `json:"headimgurl"`
	Language       string   `json:"language"`
	Nickname       string   `json:"nickname"`
	Openid         string   `json:"openid"`
	Privilege      []string `json:"privilege"`
	Province       string   `json:"province"`
	Sex            uint     `json:"sex"`
	Subscribe      int      `json:"subscribe"`
	SubscribeTime  uint32   `json:"subscribe_time"`
	UnionID        string   `json:"unionid"`
	Remark         string   `json:"remark"`
	GroupID        int      `json:"groupid"`
	TagIDList      []int    `json:"tagid_list"`
	SubscribeScene string   `json:"subscribe_scene"`
	QrScene        int      `json:"qr_scene"`
	QrSceneStr     string   `json:"qr_scene_str"`
}

/*WechatUserID WechatUserID */
type WechatUserID struct {
	OpenID string `json:"openid"`
	Lang   string `json:"lang,omitempty"`
}
