package menu_test

import (
	"encoding/json"
	"testing"

	"github.com/godcong/wego/core/menu"
)

func TestNewButton(t *testing.T) {
	button := menu.NewBaseButton()

	var subs0 []*menu.Button
	subs0 = append(subs0, menu.NewViewButton("设置首页", "https://eagle.bitbili.top"))
	subs0 = append(subs0, menu.NewClickButton("地址监控", "更多内容，敬请期待"))
	subs0 = append(subs0, menu.NewClickButton("大单买卖", "更多内容，敬请期待"))

	sub0 := menu.NewSubButton("设置提醒", subs0)
	button.AddButton(sub0)

	var subs1 []*menu.Button
	subs1 = append(subs1, menu.NewViewButton("动态首页", "https://eagle.bitbili.top"))
	sub1 := menu.NewSubButton("市场动态", subs1)
	button.AddButton(sub1)

	button.AddButton(menu.NewViewButton("联系我们", "https://eagle.bitbili.top"))

	//matchrule	是	菜单匹配规则
	//tag_id	否	用户标签的id，可通过用户标签管理接口获取
	//sex	否	性别：男（1）女（2），不填则不做匹配
	//client_platform_type	否	客户端版本，当前只具体到系统型号：IOS(1), Android(2),Others(3)，不填则不做匹配
	//country	否	国家信息，是用户在微信中设置的地区，具体请参考地区信息表
	//province	否	省份信息，是用户在微信中设置的地区，具体请参考地区信息表
	//city	否	城市信息，是用户在微信中设置的地区，具体请参考地区信息表
	//language	否	语言信息，是用户在微信中设置的语言，具体请参考语言表： 1、简体中文 "zh_CN" 2、繁体中文TW "zh_TW" 3、繁体中文HK "zh_HK" 4、英文 "en" 5、印尼 "id" 6、马来 "ms" 7、西班牙 "es" 8、韩国 "ko" 9、意大利 "it" 10、日本 "ja" 11、波兰 "pl" 12、葡萄牙 "pt" 13、俄国 "ru" 14、泰文 "th" 15、越南 "vi" 16、阿拉伯语 "ar" 17、北印度 "hi" 18、希伯来 "he" 19、土耳其 "tr" 20、德语 "de" 21、法语 "fr"
	//matchrule共六个字段，均可为空，但不能全部为空，至少要有一个匹配信息是不为空的。 country、province、city组成地区信息，将按照country、province、city的顺序进行验证，要符合地区信息表的内容。地区信息从大到小验证，小的可以不填，即若填写了省份信息，则国家信息也必填并且匹配，城市信息可以不填。 例如 “中国 广东省 广州市”、“中国 广东省”都是合法的地域信息，而“中国 广州市”则不合法，因为填写了城市信息但没有填写省份信息。 地区信息表请点击下载。
	//http://wximg.gtimg.com/shake_tv/mpwiki/areainfo.zip
	mr := menu.MatchRule{
		TagID:              "",
		Sex:                "",
		Country:            "",
		Province:           "",
		City:               "",
		ClientPlatformType: "",
		Language:           "",
	}

	button.SetMatchRule(&mr)

	v, _ := json.Marshal(button)
	t.Log(string(v))
}
