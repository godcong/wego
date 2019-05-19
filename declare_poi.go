package wego

// PhotoList ...
type PhotoList []struct {
	PhotoURL string `json:"photo_url"`
}

/*PoiBaseInfo PoiBaseInfo*/
type PoiBaseInfo struct {
	Poi          string    `json:"poi,omitempty"`          // "poi_id ":"271864249"
	Sid          string    `json:"sid,omitempty"`          // "sid":"33788392",
	BusinessName string    `json:"business_name"`          //"business_name":"15个汉字或30个英文字符内",
	BranchName   string    `json:"branch_name"`            //"branch_name":"不超过10个字，不能含有括号和特殊字符",
	Province     string    `json:"province"`               //"province":"不超过10个字",
	City         string    `json:"city"`                   //"city":"不超过30个字",
	District     string    `json:"district"`               //"district":"不超过10个字",
	Address      string    `json:"address"`                //"address":"门店所在的详细街道地址（不要填写省市信息）:不超过80个字",
	Telephone    string    `json:"telephone"`              //"telephone":"不超53个字符（不可以出现文字）",
	Categories   []string  `json:"categories"`             //"categories":["美食,小吃快餐"],
	OffsetType   int       `json:"offset_type"`            //"offset_type":1,
	Longitude    float64   `json:"longitude"`              //"longitude":115.32375,
	Latitude     float64   `json:"latitude"`               //"latitude":25.097486,
	PhotoList    PhotoList `json:"photo_list,omitempty"`   //"photo_list":[{"photo_url":"https:// 不超过20张.com"}，{"photo_url":"https://XXX.com"}],
	Recommend    string    `json:"recommend,omitempty"`    //"recommend":"不超过200字。麦辣鸡腿堡套餐，麦乐鸡，全家桶",
	Special      string    `json:"special,omitempty"`      //"special":"不超过200字。免费wifi，外卖服务",
	Introduction string    `json:"introduction,omitempty"` //"introduction":"不超过300字。麦当劳是全球大型跨国连锁餐厅，1940 年创立于美国，在世界上大约拥有3 万间分店。	主要售卖汉堡包，以及薯条、炸鸡、汽水、冰品、沙拉、 水果等快餐食品",
	OpenTime     string    `json:"open_time,omitempty"`    //"open_time":"8:00-20:00",
	AvgPrice     int       `json:"avg_price,omitempty"`    //"avg_price":35
}
