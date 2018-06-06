package core

/*UserInfo UserInfo */
type UserInfo struct {
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

/*UserID UserID */
type UserID struct {
	OpenID string `json:"openid"`
	Lang   string `json:"lang,omitempty"`
}
