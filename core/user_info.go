package core

type UserInfo struct {
	City       string   `jsong:"city"`
	Country    string   `jsong:"country"`
	HeadImgUrl string   `jsong:"headimgurl"`
	Language   string   `jsong:"language"`
	Nickname   string   `jsong:"nickname"`
	Openid     string   `jsong:"openid"`
	Privilege  []string `jsong:"privilege"`
	Province   string   `jsong:"province"`
	Sex        uint     `jsong:"sex"`
}
