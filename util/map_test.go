package util_test

import (
	"encoding/xml"
	"github.com/godcong/wego/util"
	"testing"
)

// TestMap_Set ...
func TestMap_Set(t *testing.T) {
	m := make(util.Map)
	m.Set("one.two.three", "abc")

	t.Log(m.Get("one.two.three") == "abc")
}

// TestMap_Delete ...
func TestMap_Delete(t *testing.T) {
	m := make(util.Map)
	m.Set("one.two.three", "abc")
	if !m.Has("one.two") {
		t.Error("one.two")
	}
	m.Set("one.two.ab", "ddd")
	if m.GetString("one.two.three") != "abc" {
		t.Error("one.two.three")
	}

	if !m.Delete("one.two.ab") {
		t.Error("one.two.ab")
	}
}

// TestMap_Expect ...
func TestMap_Expect(t *testing.T) {
	m := make(util.Map)
	m.Set("one.two.three", "abc")
	m.Set("one.two.ab", "ddd")
	t.Log(m)
	t.Log(m.Expect([]string{"one.two.ab"}))
	t.Log(m)

}

// TestMap_UnmarshalXML ...
func TestMap_UnmarshalXML(t *testing.T) {
	json := `{
	"card": {
		"card_type": "GROUPON",
		"groupon": {
			"base_info": {
				"logo_url": "http://mmbiz.qpic.cn/mmbiz/iaL1LJM1mF9aRKPZJkmG8xXhiaHqkKSVMMWeN3hLut7X7hicFNjakmxibMLGWpXrEXB33367o7zHN0CwngnQY7zb7g/0",
				"brand_name": "微信餐厅",
				"code_type": "CODE_TYPE_TEXT",
				"title": "132元双人火锅套餐",
				"color": "Color010",
				"notice": "使用时向服务员出示此券",
				"service_phone": "020-88888888",
				"description": "不可与其他优惠同享\n如需团购券发票，请在消费时向商户提出\n店内均可使用，仅限堂食",
				"date_info": {
					"type": "DATE_TYPE_FIX_TIME_RANGE",
					"begin_timestamp": 1397577600,
					"end_timestamp": 1472724261
				},
				"sku": {
					"quantity": 500000
				},
				"use_limit": 100,
				"get_limit": 3,
				"use_custom_code": false,
				"bind_openid": false,
				"can_share": true,
				"can_give_friend": true,
				"location_id_list": [123,12321,345345],
				"center_title": "顶部居中按钮",
				"center_sub_title": "按钮下方的wording",
				"center_url": "www.qq.com",
				"custom_url_name": "立即使用",
				"custom_url": "http://www.qq.com",
				"custom_url_sub_title": "6个汉字tips",
				"promotion_url_name": "更多优惠",
				"promotion_url": "http://www.qq.com",
				"source": "大众点评"
			},
			"advanced_info": {
				"use_condition": {
					"accept_category": "鞋类",
					"reject_category": "阿迪达斯",
					"can_use_with_other_discount": true
				},
				"abstract": {
					"abstract": "微信餐厅推出多种新季菜品，期待您的光临",
					"icon_url_list": [
						"http://mmbiz.qpic.cn/mmbiz/p98FjXy8LacgHxp3sJ3vn97bGLz0ib0Sfz1bjiaoOYA027iasqSG0sj\n piby4vce3AtaPu6cIhBHkt6IjlkY9YnDsfw/0"
					]
				},
				"text_image_list": [
					{
						"image_url": "http://mmbiz.qpic.cn/mmbiz/p98FjXy8LacgHxp3sJ3vn97bGLz0ib0Sfz1bjiaoOYA027iasqSG0sjpiby4vce3AtaPu6cIhBHkt6IjlkY9YnDsfw/0",
						"text": "此菜品精选食材，以独特的烹饪方法，最大程度地刺激食 客的味蕾"
					},
					{
						"image_url": "http://mmbiz.qpic.cn/mmbiz/p98FjXy8LacgHxp3sJ3vn97bGLz0ib0Sfz1bjiaoOYA027iasqSG0sj piby4vce3AtaPu6cIhBHkt6IjlkY9YnDsfw/0",
						"text": "此菜品迎合大众口味，老少皆宜，营养均衡"
					}
				],
				"time_limit": [
					{
						"type": "MONDAY",
						"begin_hour": 0,
						"end_hour": 10,
						"begin_minute": 10,
						"end_minute": 59
					},
					{
						"type": "HOLIDAY"
					}
				],
				"business_service": [
					"BIZ_SERVICE_FREE_WIFI",
					"BIZ_SERVICE_WITH_PET",
					"BIZ_SERVICE_FREE_PARK",
					"BIZ_SERVICE_DELIVER"
				]
			},
			"deal_detail": "以下锅底2选1（有菌王锅、麻辣锅、大骨锅、番茄锅、清补 凉锅、酸菜鱼锅可选）：\n大锅1份 12元\n小锅2份 16元 "
		}
	}
}`
	m := util.JSONToMap([]byte(json))
	t.Log(m)
	x := m.ToXML()
	t.Log(string(x))
	t.Log(util.XMLToMap(x))
}

// BenchmarkMap_ToXML ...
func BenchmarkMap_ToXML(b *testing.B) {
	json := `{
	"card": {
		"card_type": "GROUPON",
		"groupon": {
			"base_info": {
				"logo_url": "http://mmbiz.qpic.cn/mmbiz/iaL1LJM1mF9aRKPZJkmG8xXhiaHqkKSVMMWeN3hLut7X7hicFNjakmxibMLGWpXrEXB33367o7zHN0CwngnQY7zb7g/0",
				"brand_name": "微信餐厅",
				"code_type": "CODE_TYPE_TEXT",
				"title": "132元双人火锅套餐",
				"color": "Color010",
				"notice": "使用时向服务员出示此券",
				"service_phone": "020-88888888",
				"description": "不可与其他优惠同享\n如需团购券发票，请在消费时向商户提出\n店内均可使用，仅限堂食",
				"date_info": {
					"type": "DATE_TYPE_FIX_TIME_RANGE",
					"begin_timestamp": 1397577600,
					"end_timestamp": 1472724261
				},
				"sku": {
					"quantity": 500000
				},
				"use_limit": 100,
				"get_limit": 3,
				"use_custom_code": false,
				"bind_openid": false,
				"can_share": true,
				"can_give_friend": true,
				"location_id_list": [123,12321,345345],
				"center_title": "顶部居中按钮",
				"center_sub_title": "按钮下方的wording",
				"center_url": "www.qq.com",
				"custom_url_name": "立即使用",
				"custom_url": "http://www.qq.com",
				"custom_url_sub_title": "6个汉字tips",
				"promotion_url_name": "更多优惠",
				"promotion_url": "http://www.qq.com",
				"source": "大众点评"
			},
			"advanced_info": {
				"use_condition": {
					"accept_category": "鞋类",
					"reject_category": "阿迪达斯",
					"can_use_with_other_discount": true
				},
				"abstract": {
					"abstract": "微信餐厅推出多种新季菜品，期待您的光临",
					"icon_url_list": [
						"http://mmbiz.qpic.cn/mmbiz/p98FjXy8LacgHxp3sJ3vn97bGLz0ib0Sfz1bjiaoOYA027iasqSG0sj\n piby4vce3AtaPu6cIhBHkt6IjlkY9YnDsfw/0"
					]
				},
				"text_image_list": [
					{
						"image_url": "http://mmbiz.qpic.cn/mmbiz/p98FjXy8LacgHxp3sJ3vn97bGLz0ib0Sfz1bjiaoOYA027iasqSG0sjpiby4vce3AtaPu6cIhBHkt6IjlkY9YnDsfw/0",
						"text": "此菜品精选食材，以独特的烹饪方法，最大程度地刺激食 客的味蕾"
					},
					{
						"image_url": "http://mmbiz.qpic.cn/mmbiz/p98FjXy8LacgHxp3sJ3vn97bGLz0ib0Sfz1bjiaoOYA027iasqSG0sj piby4vce3AtaPu6cIhBHkt6IjlkY9YnDsfw/0",
						"text": "此菜品迎合大众口味，老少皆宜，营养均衡"
					}
				],
				"time_limit": [
					{
						"type": "MONDAY",
						"begin_hour": 0,
						"end_hour": 10,
						"begin_minute": 10,
						"end_minute": 59
					},
					{
						"type": "HOLIDAY"
					}
				],
				"business_service": [
					"BIZ_SERVICE_FREE_WIFI",
					"BIZ_SERVICE_WITH_PET",
					"BIZ_SERVICE_FREE_PARK",
					"BIZ_SERVICE_DELIVER"
				]
			},
			"deal_detail": "以下锅底2选1（有菌王锅、麻辣锅、大骨锅、番茄锅、清补 凉锅、酸菜鱼锅可选）：\n大锅1份 12元\n小锅2份 16元 "
		}
	}
}`

	m := util.JSONToMap([]byte(json))
	b.Log(m)
	x := m.ToXML()
	b.Log(string(x))
	b.ResetTimer()
	for i := 0; i < 1000; i++ {
		m2 := util.Map{}
		_ = xml.Unmarshal(x, &m2)
		b.Log(m2)
	}

}

// TestMap_MarshalXML ...
func TestMap_MarshalXML(t *testing.T) {
	json := []byte(`{"appid":"wx1ad61aeef1903b93","bank_type":"CMB_DEBIT","cash_fee":"200","fee_type":"CNY","is_subscribe":"N","mch_id":"1498009232","nonce_str":"7cda1edf536f11e88cb200163e04155d","openid":"oE_gl0bQ7iJ2g3OBMQPWRiBSoiks","out_trade_no":"8195400821515968","result_code":"SUCCESS","return_code":"SUCCESS","sign":"BE9EA07614C09FA73A683071877D9DDB","time_end":"20180509175821","total_fee":"200","trade_type":"JSAPI","transaction_id":"4200000155201805096015992498"}`)
	m := util.JSONToMap(json)
	t.Log(string(m.ToXML()))
}

// TestStructToMap ...
func TestStructToMap(t *testing.T) {
	src := struct {
		One   string
		Two   string
		Three int
		Four  uint64
	}{
		One:   "a",
		Two:   "b",
		Three: 10,
		Four:  100,
	}
	p := util.Map{}
	e := util.StructToMap(src, p)
	t.Log(p, e)
}
