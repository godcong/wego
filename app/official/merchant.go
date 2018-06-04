package official

/*Merchant Merchant*/
type Merchant struct {
}

/*Create 增加库存
协议	https
http请求方式	POST
请求Url	https://api.weixin.qq.com/merchant/create?access_token=ACCESS_TOKEN
POST数据格式	json
参数	是否必须	说明
access_token	是	公众号的调用接口凭证，由access_token生成接口获得
POST数据	是	商品详细信息
字段	说明
product_id	商品ID
sku_info	sku信息,格式"id1:vid1;id2:vid2",如商品为统一规格，则此处赋值为空字符串即可
quantity	增加的库存数量
返回数据
字段	说明
errcode	错误码
errmsg	错误信息
*/
func (m *Merchant) Create() {

}

/*StockReduce 减少库存
 	协议	https
	http请求方式	POST
	请求Url	https://api.weixin.qq.com/merchant/stock/reduce?access_token=ACCESS_TOKEN
	POST数据格式	json
	请求参数说明
	参数	是否必须	说明
	access_token	是	调用接口凭证
	POST数据	是	商品库存信息
	POST数据
	数据示例：
	{
		"product_id": "pDF3iY5EYkMxs4-tF8xedyES5GQI",
		"sku_info": "10000983:10000995;10001007:10001010",
		"quantity": 20
	}
	字段	说明
	product_id	商品ID
	sku_info	sku信息, 格式"id1:vid1;id2:vid2"
	quantity	减少的库存数量
	返回数据说明
	数据示例：
	{
		"errcode":0,
		"errmsg":"success"
	}
	字段	说明
	errcode	错误码
	errmsg	错误信息

*/
func (m *Merchant) StockReduce() {

}

/*
增加邮费模板 ExpressAdd
	协议	https
	http请求方式	POST
	请求Url	https://api.weixin.qq.com/merchant/express/add?access_token=ACCESS_TOKEN
	POST数据格式	json
	3.1.2	请求参数说明
	参数	是否必须	说明
	access_token	是	公众号的调用接口凭证
	POST数据	是	邮费信息
	3.1.3	POST数据
	数据示例：
	{
		"delivery_template": {
			"Name": "testexpress",
			"Assumer": 0,
			"Valuation": 0,
			"TopFee": [
				{
					"Type": 10000027,
					"Normal": {
						"StartStandards": 1,
						"StartFees": 2,
						"AddStandards": 3,
						"AddFees": 1
					},
					"Custom": [{
						"StartStandards": 1,
						"StartFees": 100,
						"AddStandards": 1,
						"AddFees": 3,
						"DestCountry": "中国",
						"DestProvince": "广东省",
						"DestCity": "广州市"
					}]
				},
				{
					"Type": 10000028,
					"Normal": {
						"StartStandards": 1,
						"StartFees": 3,
						"AddStandards": 3,
						"AddFees": 2
					},
					"Custom": [{
						"StartStandards": 1,
						"StartFees": 10,
						"AddStandards": 1,
						"AddFees": 30,
						"DestCountry": "中国",
						"DestProvince": "广东省",
						"DestCity": "广州市"
					}]
				},
				{
					"Type": 10000029,
					"Normal": {
						"StartStandards": 1,
						"StartFees": 4,
						"AddStandards": 3,
						"AddFees": 3
					},
					"Custom": [{
						"StartStandards": 1,
						"StartFees": 8,
						"AddStandards": 2,
						"AddFees": 11,
						"DestCountry": "中国",
						"DestProvince": "广东省",
						"DestCity": "广州市"
					}]
				}
			]
		}
	}
	字段	说明
	Name	邮费模板名称
	Assumer	支付方式(0-买家承担运费, 1-卖家承担运费)
	Valuation	计费单位(0-按件计费, 1-按重量计费, 2-按体积计费，目前只支持按件计费，默认为0)
	TopFee		具体运费计算
		Type	快递类型ID(参见增加商品/快递列表)
		Normal		默认邮费计算方法
			StartStandards	起始计费数量(比如计费单位是按件, 填2代表起始计费为2件)
			StartFees	起始计费金额(单位: 分）
			AddStandards	递增计费数量
			AddFees	递增计费金额(单位 : 分)
		Custom		指定地区邮费计算方法
			StartStandards	起始计费数量
			StartFees	起始计费金额(单位: 分）
			AddStandards	递增计费数量
			AddFees	递增计费金额(单位 : 分)
			DestCountry	指定国家(详见《地区列表》说明)
			DestProvince	指定省份(详见《地区列表》说明)
			DestCity	指定城市(详见《地区列表》说明)
	例子解析：Type为10000027的默认邮费计算，第1件邮费2分，每增加3件邮费增加3分。Type为10000027的指定地区邮费计算，寄送到中国/广东省/广州市的商品，第一件邮费1元，每增加1件邮费增加3分。
	3.1.4	返回数据说明
	数据示例：
	{
		"errcode": 0,
	"errmsg": "success"，
	"template_id": 123456
	}
	字段	说明
	errcode	错误码
	errmsg	错误信息
	template_id	邮费模板ID
*/

/*
删除邮费模板 ExpressDel
	协议	https
	http请求方式	POST
	请求Url	https://api.weixin.qq.com/merchant/express/del?access_token=ACCESS_TOKEN
	POST数据格式	json
	3.2.2	请求参数说明
	参数	是否必须	说明
	access_token	是	公众号的调用接口凭证
	POST数据	是	邮费信息
	3.2.3	POST数据
	数据示例：
	{
		"template_id": 123456
	}
	字段	说明
	template_id	邮费模板ID
	3.2.4	返回数据说明
	数据示例：
	{
		"errcode": 0,
	"errmsg": "success"
	}
	字段	说明
	errcode	错误码
	errmsg	错误信息
*/

/*
3.3	修改邮费模板 express/update
	协议	https
	http请求方式	POST
	请求Url	https://api.weixin.qq.com/merchant/express/update?access_token=ACCESS_TOKEN
	POST数据格式	json
	3.3.2	请求参数说明
	参数	是否必须	说明
	access_token	是	公众号的调用接口凭证
	POST数据	是	邮费信息
	3.3.3	POST数据
	数据示例：
	{
		"template_id": 123456,
		"delivery_template": ...
	}
	字段	说明
	template_id	邮费模板ID
	delivery_template	邮费模板信息(字段说明详见增加邮费模板)
	3.3.4	返回数据说明
	数据示例：
	{
		"errcode": 0,
	"errmsg": "success"
	}
	字段	说明
	errcode	错误码
	errmsg	错误信息
*/

/*
3.4	获取指定ID的邮费模板 express/getbyid
	协议	https
	http请求方式	POST
	请求Url	https://api.weixin.qq.com/merchant/express/getbyid?access_token=ACCESS_TOKEN
	POST数据格式	json
	3.4.2	请求参数说明
	参数	是否必须	说明
	access_token	是	公众号的调用接口凭证
	POST数据	是	邮费信息
	3.4.3	POST数据
	数据示例：
	{
		"template_id": 123456
	}
	字段	说明
	template_id	邮费模板ID
	3.4.4	返回数据说明
	数据示例：
	{
	  "errcode": 0,
	  "errmsg": "success",
	  "template_info": {
		"Id": 103312916,
		"Name": "testexpress",
		"Assumer": 0,
		"Valuation": 0,
		"TopFee": [
		  {
			"Type": 10000027,
			"Normal": {
			  "StartStandards": 1,
			  "StartFees": 2,
			  "AddStandards": 3,
			  "AddFees": 1
			},
			"Custom": [
			  {
				"StartStandards": 1,
				"StartFees": 1000,
				"AddStandards": 1,
				"AddFees": 3,
				"DestCountry": "中国",
				"DestProvince": "广东省",
				"DestCity": "广州市"
			  }
			]
		  },
		  {
			"Type": 10000028,
			"Normal": {
			  "StartStandards": 1,
			  "StartFees": 3,
			  "AddStandards": 3,
			  "AddFees": 2
			},
			"Custom": [
			  {
				"StartStandards": 1,
				"StartFees": 10,
				"AddStandards": 1,
				"AddFees": 30,
				"DestCountry": "中国",
				"DestProvince": "广东省",
				"DestCity": "广州市"
			  }
			]
		  },
		  {
			"Type": 10000029,
			"Normal": {
			  "StartStandards": 1,
			  "StartFees": 4,
			  "AddStandards": 3,
			  "AddFees": 3
			},
			"Custom": [
			  {
				"StartStandards": 1,
				"StartFees": 8,
				"AddStandards": 2,
				"AddFees": 11,
				"DestCountry": "中国",
				"DestProvince": "广东省",
				"DestCity": "广州市"
			  }
			]
		  }
		]
	  }
	}
	字段	说明
	errcode	错误码
	errmsg	错误信息
	template_info	邮费模板信息(字段说明详见增加邮费模板)

*/

/*
3.5	获取所有邮费模板 express/getall
	协议	https
	http请求方式	GET
	请求Url	https://api.weixin.qq.com/merchant/express/getall?access_token=ACCESS_TOKEN
	3.5.2	请求参数说明
	参数	是否必须	说明
	access_token	是	公众号的调用接口凭证
	3.5.3	返回数据说明
	数据示例：
	{
	  "errcode": 0,
	  "errmsg": "success",
	  "templates_info": [
		  {
			"Id": 103312916,
			"Name": "testexpress1",
			"Assumer": 0,
			"Valuation": 0,
			"TopFee": [...],
		  },
		  {
			"Id": 103312917,
			"Name": "testexpress2",
			"Assumer": 0,
			"Valuation": 2,
			"TopFee": [...],
		  },
		  {
			"Id": 103312918,
			"Name": "testexpress3",
			"Assumer": 0,
			"Valuation": 1,
			"TopFee": [...],
		  }
	  ]
	}
	字段	说明
	errcode	错误码
	errmsg	错误信息
	templates_info	所有邮费模板集合(字段说明详见增加邮费模板)

*/

/*
4.1	增加分组 group/add
	协议	https
	http请求方式	POST
	请求Url	https://api.weixin.qq.com/merchant/group/add?access_token=ACCESS_TOKEN
	POST数据格式	json
	4.1.2	请求参数说明
	参数	是否必须	说明
	access_token	是	调用接口凭证
	POST数据	是	商品分组信息
	4.1.3	POST数据
	数据示例：
	{
		"group_detail" : {
			"group_name": "测试分组",
			"product_list" : [
				"pDF3iY9cEWyMimNlKbik_NYJTzYU",
				"pDF3iY4kpZagQfwJ_LVQBaOC-LsM"
			]
		}
	}
	字段	说明
	group_name	分组名称
	product_list	商品ID集合
	4.1.4	返回数据说明
	数据示例：
	{
		"errcode":0,
	"errmsg":"success",
	"group_id": 19
	}
	字段	说明
	errcode	错误码
	errmsg	错误信息
	group_id	分组ID

*/

/*
4.2	删除分组 group/del
	协议	https
	http请求方式	POST
	请求Url	https://api.weixin.qq.com/merchant/group/del?access_token=ACCESS_TOKEN
	POST数据格式	json
	4.2.2	请求参数说明
	参数	是否必须	说明
	access_token	是	调用接口凭证
	POST数据	是	商品分组信息
	4.2.3	POST数据
	数据示例：
	{
		"group_id": 19
	}
	字段	说明
	group_id	分组ID
	4.2.4	返回数据说明
	数据示例：
	{
		"errcode":0,
	"errmsg":"success"
	}
	字段	说明
	errcode	错误码
	errmsg	错误信息

*/

/*
4.3	修改分组属性 group/propertymod
	协议	https
	http请求方式	POST
	请求Url	https://api.weixin.qq.com/merchant/group/propertymod?access_token=ACCESS_TOKEN
	POST数据格式	json
	4.3.2	请求参数说明
	参数	是否必须	说明
	access_token	是	调用接口凭证
	POST数据	是	分组属性
	4.3.3	POST数据
	数据示例：
	{
		"group_id": 28,
		"group_name":"特惠专场"
	}
	4.3.4	返回数据说明
	数据示例：
	{
		"errcode":0,
	"errmsg":"success"
	}
	字段	说明
	errcode	错误码
	errmsg	错误信息

*/

/*
4.4	修改分组商品 group/productmod
	协议	https
	http请求方式	POST
	请求Url	https://api.weixin.qq.com/merchant/group/productmod?access_token=ACCESS_TOKEN
	POST数据格式	json
	4.4.2	请求参数说明
	参数	是否必须	说明
	access_token	是	调用接口凭证
	POST数据	是	分组商品信息
	4.4.3	POST数据
	数据示例：
	{
		"group_id": 28,
		"product": [
			{
				"product_id": "pDF3iY-CgqlAL3k8Ilz-6sj0UYpk",
				"mod_action": 1
			},
			{
				"product_id": "pDF3iY-RewlAL3k8Ilz-6sjsepp9",
				"mod_action": 0
			},
		]
	}
	字段	说明
	group_id	分组ID
	product		分组的商品集合
		product_id	商品ID
		mod_action	修改操作(0-删除, 1-增加)
	4.4.4	返回数据说明
	数据示例：
	{
		"errcode":0,
	"errmsg":"success"
	}
	字段	说明
	errcode	错误码
	errmsg	错误信息

*/

/*
4.5	获取所有分组 group/getall
	协议	https
	http请求方式	GET
	请求Url	https://api.weixin.qq.com/merchant/group/getall?access_token=ACCESS_TOKEN
	4.5.2	请求参数说明
	参数	是否必须	说明
	access_token	是	调用接口凭证
	4.5.3	返回数据说明
	数据示例：
	{
		"errcode": 0,
		"errmsg": "success",
		"groups_detail": [
			{
			  "group_id": 200077549,
			  "group_name": "最新上架"
			},
			{
			  "group_id": 200079772,
			  "group_name": "全球热卖"
			}
		]
	}
	字段	说明
	errcode	错误码
	errmsg	错误信息
	groups_detail		分组集合
		group_id	分组ID
		group_name	分组名称

*/

/*
4.6	根据分组ID获取分组信息 group/getbyid
	协议	https
	http请求方式	POST
	请求Url	https://api.weixin.qq.com/merchant/group/getbyid?access_token=ACCESS_TOKEN
	POST数据格式	json
	4.6.2	请求参数说明
	参数	是否必须	说明
	access_token	是	调用接口凭证
	POST数据	是	分组信息
	4.6.3	POST数据
	数据示例：
	{
		"group_id": 29
	}
	字段	说明
	group_id	分组ID
	4.6.4	返回数据说明
	数据示例：
	{
		"errcode": 0,
		"errmsg": "success",
		"group_detail": {
			"group_id": 200077549,
			"group_name": "最新上架",
			"product_list": [
			  "pDF3iYzZoY-Budrzt8O6IxrwIJAA",
			  "pDF3iY3pnWSGJcO2MpS2Nxy3HWx8",
			  "pDF3iY33jNt0Dj3M3UqiGlUxGrio"
			]
		}
	}
	字段	说明
	errcode	错误码
	errmsg	错误信息
	groups_detail		分组信息
		group_id	分组ID
		group_name	分组名称
		product_list	商品ID集合

*/

/*
5.1	增加货架 shelf/add
	协议	https
	http请求方式	POST
	请求Url	https://api.weixin.qq.com/merchant/shelf/add?access_token=ACCESS_TOKEN
	POST数据格式	json
	5.1.2	请求参数说明
	参数	是否必须	说明
	access_token	是	调用接口凭证
	POST数据	是	货架详情信息
	5.1.3	POST数据
	{
		"shelf_data": {
		  "module_infos": [
			{
			  "group_info": {
				"filter": {
				  "count": 2
				},
				"group_id": 50
			  },
			  "eid": 1
			},
			{
				"group_infos": {
					"groups": [
					  {
						"group_id": 49
					  },
					  {
						"group_id": 50
					  },
					  {
						"group_id": 51
					  }
					]
			  },
			  "eid": 2
			},
			{
			  "group_info": {
				"group_id": 52,
				"img": "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl29nqqObBwFwnIX3licVPnFV5Jm64z4I0TTicv0TjN7Vl9bykUUibYKIOjicAwIt6Oy0Y6a1Rjp5Tos8tg/0"
			  },
			  "eid": 3
			},
			{
			  "group_infos": {
				"groups": [
				  {
					"group_id": 49,
					"img": "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl29nqqObBwFwnIX3licVPnFV5uUQx7TLx4tB9qZfbe3JmqR4NkkEmpb5LUWoXF1ek9nga0IkeSSFZ8g/0"
				  },
				  {
					"group_id": 50,
					"img": "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl29nqqObBwFwnIX3licVPnFV5G1kdy3ViblHrR54gbCmbiaMnl5HpLGm5JFeENyO9FEZAy6mPypEpLibLA/0"
				  },
				  {
					"group_id": 52,
					"img": "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl29nqqObBwFwnIX3licVPnFV5uUQx7TLx4tB9qZfbe3JmqR4NkkEmpb5LUWoXF1ek9nga0IkeSSFZ8g/0"
				  }
				]
			  },
			  "eid": 4
			},
			{
			  "group_infos": {
				"groups": [
				  {
					"group_id": 43
				  },
				  {
					"group_id": 44
				  },
				  {
					"group_id": 45
				  },
				  {
					"group_id": 46
				  }
				],
				"img_background": "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl29nqqObBwFwnIX3licVPnFV5uUQx7TLx4tB9qZfbe3JmqR4NkkEmpb5LUWoXF1ek9nga0IkeSSFZ8g/0"
			  },
			  "eid": 5
			}
		  ]
		},
		"shelf_banner": "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl2ibrWQn8zWFUh1YznsMV0XEiavFfLzDWYyvQOBBszXlMaiabGWzz5B2KhNn2IDemHa3iarmCyribYlZYyw/0",
		"shelf_name": "测试货架"
	}
	字段	说明
	shelf_data	货架信息(数据说明详见《货架控件说明》)
	shelf_banner	货架招牌图片Url(图片需调用图片上传接口获得图片Url填写至此，否则添加货架失败，建议尺寸为640*120，仅控件1-4有banner，控件5没有banner)
	shelf_name	货架名称
	5.1.4	货架控件说明
	5.1.4.1	控件1
	控件1是由一个分组组成，展示该分组指定数量的商品列表，可与控件2、控件3、控件4联合使用。
	A. POST数据/控件UI


	{
	  "group_info": {
		"filter": {
		  "count": 4
		},
		"group_id": 50
	  },
	  "eid": 1
	}



	B. 数据说明
	字段	说明
	group_info		分组信息
		count	该控件展示商品个数
		group_id	分组ID
	eid	控件1的ID
	5.1.4.2	控件2
	控件2是由多个分组组成（最多有4个分组），展示指定分组的名称，可与控件1、控件3、控件4联合使用。
	A. POST数据/控件UI

	{
		"group_infos": {
			"groups": [
			  {
				"group_id": 49
			  },
			  {
				"group_id": 50
			  },
			  {
				"group_id": 51
			  }
			]
	  },
	  "eid": 2
	}


	B. 数据说明
	字段	说明
	groups		分组数组
		group_id	分组ID
	eid	控件2的ID
	5.1.4.3	控件3
	控件3是由一个分组组成，展示指定分组的分组图片，可与控件1、控件2、控件4联合使用。
	A. POST数据/控件UI


	{
	  "group_info": {
		"group_id": 52,
		"img": "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl29nqqObBwFwnIX3licVPnFV5Jm64z4I0TTicv0TjN7Vl9bykUUibYKIOjicAwIt6Oy0Y6a1Rjp5Tos8tg/0"
	  },
	  "eid": 3
	}


	B. 数据说明
	字段	说明
	group_info		分组信息
		group_id	分组ID
	img	分组照片(图片需调用图片上传接口获得图片Url填写至此，否则添加货架失败，建议分辨率600*208)
	eid	控件3的ID
	5.1.4.4	控件4
	控件4是由多个分组组成（最多3个分组），展示指定分组的分组图片，可与控件1、控件2、控件4联合使用。
	A. POST数据/控件UI
	{
	  "group_infos": {
		"groups": [
		  {
			"group_id": 49,
			"img": "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl29nqqObBwFwnIX3licVPnFV5uUQx7TLx4tB9qZfbe3JmqR4NkkEmpb5LUWoXF1ek9nga0IkeSSFZ8g/0"
		  },
		  {
			"group_id": 50,
			"img": "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl29nqqObBwFwnIX3licVPnFV5G1kdy3ViblHrR54gbCmbiaMnl5HpLGm5JFeENyO9FEZAy6mPypEpLibLA/0"
		  },
		  {
			"group_id": 52,
			"img": "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl29nqqObBwFwnIX3licVPnFV5uUQx7TLx4tB9qZfbe3JmqR4NkkEmpb5LUWoXF1ek9nga0IkeSSFZ8g/0"
		  }
		]
	  },
	  "eid": 4
	}
	B. 数据说明
	字段	说明
	groups		分组列表
		group_id	分组ID
		img	分组照片(图片需调用图片上传接口获得图片Url填写至此，否则添加货架失败，3个分组建议分辨率分别为: 350*350, 244*172, 244*172)
	eid	控件4的ID
	5.1.4.5	控件5
	控件5是由多个分组组成，展示指定分组的名称，不可与其他控件联合使用。
	A. POST数据/UI展示
	{
	  "group_infos": {
		"groups": [
		  {
			"group_id": 43
		  },
		  {
			"group_id": 44
		  },
		  {
			"group_id": 45
		  },
		  {
			"group_id": 46
		  }
		],
		"img_background": "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl29nqqObBwFwnIX3licVPnFV5uUQx7TLx4tB9qZfbe3JmqR4NkkEmpb5LUWoXF1ek9nga0IkeSSFZ8g/0"
	  },
	  "eid": 5
	}



	B. 数据说明
	字段	说明
	groups		分组列表
		group_id	分组ID
		img	分组照片(图片需调用图片上传接口获得图片Url填写至此，否则添加货架失败，建议分辨率640*1008)
	eid	控件5的ID
	5.1.5	返回数据说明
	数据示例：
	{
		"errcode":0,
	"errmsg":"success",
	"shelf_id": 12
	}
	字段	说明
	errcode	错误码
	errmsg	错误信息
	shelf_id	货架ID

*/

/*
5.2	删除货架 shelf/del
	协议	https
	http请求方式	POST
	请求Url	https://api.weixin.qq.com/merchant/shelf/del?access_token=ACCESS_TOKEN
	POST数据格式	json
	5.2.2	请求参数说明
	参数	是否必须	说明
	access_token	是	调用接口凭证
	POST数据	是	货架信息
	5.2.3	POST数据
	数据示例：
	{
		"shelf_id": 19
	}
	字段	说明
	shelf_id	货架ID
	5.2.4	返回数据说明
	数据示例：
	{
		"errcode":0,
	"errmsg":"success"
	}
	字段	说明
	errcode	错误码
	errmsg	错误信息

*/

/*
5.3	修改货架 shelf/mod
协议	https
http请求方式	POST
请求Url	https://api.weixin.qq.com/merchant/shelf/mod?access_token=ACCESS_TOKEN
POST数据格式	json
5.3.2	请求参数说明
参数	是否必须	说明
access_token	是	调用接口凭证
POST数据	是	货架详情信息
5.3.3	POST数据
数据示例：
{
	"shelf_id": 12345,
	"shelf_data": ...,
	"shelf_banner": "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl2ibrWQn8zWFUh1YznsMV0XEiavFfLzDWYyvQOBBszXlMaiabGWzz5B2KhNn2IDemHa3iarmCyribYlZYyw/0",
	"shelf_name": "测试货架"
}
字段	说明
shelf_id	货架ID
shelf_data	货架详情(字段说明详见增加货架)
shelf_banner	货架banner(图片需调用图片上传接口获得图片Url填写至此，否则修改货架失败)
shelf_name	货架名称
5.3.4	返回数据说明
数据示例：
{
    "errcode":0,
"errmsg":"success"
}
字段	说明
errcode	错误码
errmsg	错误信息

*/

/*
5.4	获取所有货架 shelf/getall
	协议	https
	http请求方式	GET
	请求Url	https://api.weixin.qq.com/merchant/shelf/getall?access_token=ACCESS_TOKEN
	5.4.2	请求参数说明
	参数	是否必须	说明
	access_token	是	调用接口凭证
	5.4.3	返回数据说明
	数据示例：
	{
		"errcode": 0,
		"errmsg": "success",
		"shelves": [
			{
			  "shelf_info": {
				"module_infos": [
				  {
					"group_infos": {
					  "groups": [
						{
						  "group_id": 200080093
						},
						{
						  "group_id": 200080118
						},
						{
						  "group_id": 200080119
						},
						{
						  "group_id": 200080135
						}
					  ],
					  "img_background": "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl294FzPwnf9dAcaN7ButStztAZyy2yHY8pW6sTQKicIhAy5F0a2CqmrvDBjMFLtc2aEhAQ7uHsPow9A/0"
					},
					"eid": 5
				  }
				]
			  },
			  "shelf_banner": "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl294FzPwnf9dAcaN7ButStztAZyy2yHY8pW6sTQKicIhAy5F0a2CqmrvDBjMFLtc2aEhAQ7uHsPow9A/0",
			  "shelf_name": "新新人类",
			  "shelf_id": 22
			},
			{
			  "shelf_info": {
				"module_infos": [
				  {
					"group_info": {
					  "group_id": 200080119,
					  "filter": {
						"count": 4
					  }
					},
					"eid": 1
				  }
				]
			  },
			  "shelf_banner": "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl294FzPwnf9dAcaN7ButStztAZyy2yHY8pW6sTQKicIhAy5F0a2CqmrvDBjMFLtc2aEhAQ7uHsPow9A/0",
			  "shelf_name": "店铺",
			  "shelf_id": 23
			}
		]
	}
	字段	说明
	errcode	错误码
	errmsg	错误信息
	shelves	所有货架集合(字段说明详见增加货架)

*/

/*
5.5	根据货架ID获取货架信息 shelf/getbyid
	协议	https
	http请求方式	POST
	请求Url	https://api.weixin.qq.com/merchant/shelf/getbyid?access_token=ACCESS_TOKEN
	POST数据格式	json
	5.5.2	请求参数说明
	参数	是否必须	说明
	access_token	是	调用接口凭证
	POST数据	是	货架信息
	5.5.3	POST数据
	数据示例：
	{
		"shelf_id": 19
	}
	字段	说明
	shelf_id	货架ID
	5.5.4	返回数据说明
	数据示例：
	{
		"errcode": 0,
		"errmsg": "success",
		"shelf_info": {
			"module_infos": [...]
		},
		"shelf_banner": "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl2ibp2DgDXiaic6WdflMpNdInS8qUia2BztlPu1gPlCDLZXEjia2qBdjoLiaCGUno9zbs1UyoqnaTJJGeEew/0",
		"shelf_name": "新建货架",
		"shelf_id": 97
	}
	字段	说明
	errcode	错误码
	errmsg	错误信息
	shelf_info	货架详情(字段说明详见增加货架)
	shelf_banner	货架banner
	shelf_name	货架名称
	shelf_id	货架ID

*/

/*
5.6	开发者将自己的页面作为货架
微信小店的货架支持开放给开发者使用，即开发者可以将自己的页面作为货架，通过JavaScript API来调起微信客户端原生的商品详情页。

请注意：
1、	开发者需要预先通过1.1中的增加商品API，预先上传商品，得到product_id后，才能在自己的页面通过JavaScript API来调起商品详情页
2、	即使开发者将自己的页面作为货架，但由于商品存储仍在微信服务器，所以用户下单后，订单、库存管理等事务，开发者仍需要根据微信小店系列接口来完成。

具体JavaScript API调用示例代码如下：
function openProductView(){
if (typeof WeixinJSBridge == "undefined")
return false;

var pid = "pDF3iY_G88cM_d-wuImym3tkVfG5";//只需要传递
WeixinJSBridge.invoke('openProductViewWithPid',{
"pid":pid
},function(res){
// 返回res.err_msg,取值
// open_product_view_with_id:ok 打开成功
alert(res.err_msg);
if (res.err_msg != "open_product_view_with_id:ok"){
WeixinJSBridge.invoke('openProductView',{
"productInfo":"{\"product_id\":\""+pid+"\",\"product_type\":0}"
},function(res){
alert(res.err_msg);
});
}
});
}
说明：
1、	变量pid即为在增加商品接口中获得的product_id。

*/

/*
6.	订单管理接口
6.1	订单付款通知
在用户在微信中付款成功后，微信服务器会将订单付款通知推送到开发者在公众平台网站中设置的回调URL（在开发模式中设置）中，如未设置回调URL，则获取不到该事件推送。

事件推送的内容如下：
<xml>
<ToUserName><![CDATA[weixin_media1]]></ToUserName>
<FromUserName><![CDATA[oDF3iYyVlek46AyTBbMRVV8VZVlI]]></FromUserName>
<CreateTime>1398144192</CreateTime>
<MsgType><![CDATA[event]]></MsgType>
<Event><![CDATA[merchant_order]]></Event>
<OrderId><![CDATA[test_order_id]]></OrderId>
<OrderStatus>2</OrderStatus>
<ProductId><![CDATA[test_product_id]]></ProductId>
<SkuInfo><![CDATA[10001:1000012;10002:100021]]></SkuInfo>
</xml>
字段说明请见订单详情。

*/

/*
6.2	根据订单ID获取订单详情 order/getbyid
	协议	https
	http请求方式	POST
	请求Url	https://api.weixin.qq.com/merchant/order/getbyid?access_token=ACCESS_TOKEN
	POST数据格式	json
	6.2.2	请求参数说明
	参数	是否必须	说明
	access_token	是	公众号的调用接口凭证
	POST数据	是	商品订单信息
	6.2.3	POST数据
	数据示例：
	{
		"order_id": "7197417460812584720"
	}
	字段	说明
	order_id	订单ID
	6.2.4	返回数据说明
	数据示例：
	{
		"errcode": 0,
		"errmsg": "success",
		"order": {
			"order_id": "7197417460812533543",
			"order_status": 6,
			"order_total_price": 6,
			"order_create_time": 1394635817,
			"order_express_price": 5,
			"buyer_openid": "oDF3iY17NsDAW4UP2qzJXPsz1S9Q",
			"buyer_nick": "likeacat",
			"receiver_name": "张小猫",
			"receiver_province": "广东省",
			"receiver_city": "广州市",
	"receiver_zone": "天河区",
			"receiver_address": "华景路一号南方通信大厦5楼",
			"receiver_mobile": "123456789",
			"receiver_phone": "123456789",
			"product_id": "pDF3iYx7KDQVGzB7kDg6Tge5OKFo",
			"product_name": "安莉芳E-BRA专柜女士舒适内衣蕾丝3/4薄杯聚拢上托性感文胸KB0716",
			"product_price": 1,
			"product_sku": "10000983:10000995;10001007:10001010",
			"product_count": 1,
			"product_img": "http://img2.paipaiimg.com/00000000/item-52B87243-63CCF66C00000000040100003565C1EA.0.300x300.jpg",
			"delivery_id": "1900659372473",
			"delivery_company": "059Yunda",
			"trans_id": "1900000109201404103172199813",
	"products": [
			{
				"product_id": "p8BCTv77lY4io_q00F9qsaniimFc",
				"product_name": "product_name",
				"product_price": 1,
				"product_sku": "",
				"product_count": 1,
				"product_img": "http://mmbiz.qpic.cn/mmbiz_gif/KfrZwACMrmxj8XRiaTUzFNsTkWdTEJySicGKMHxuG0ibDfjTtb6ZIjNgakbnKq569TbBjvicSnWdnt46gEKjWe6Vcg/0?wx_fmt=gif"
			}
		]
		}
	}
	字段	说明
	errcode	错误码
	errmsg	错误信息
	order		订单详情
		order_id	订单ID
		order_status	订单状态
		order_total_price	订单总价格(单位 : 分)
		order_create_time	订单创建时间
		order_express_price	订单运费价格(单位 : 分)
		buyer_openid	买家微信OPENID
		buyer_nick	买家微信昵称
		receiver_name	收货人姓名
		receiver_province	收货地址省份
		receiver_city	收货地址城市
		receiver_zone	收货地址区/县
		receiver_address	收货详细地址
		receiver_mobile	收货人移动电话
		receiver_phone	收货人固定电话
		product_id	商品ID
		product_name	商品名称
		product_price	商品价格(单位 : 分)
		product_sku	商品SKU
		product_count	商品个数
		product_img	商品图片
		delivery_id	运单ID
		delivery_company	物流公司编码
		trans_id	交易ID
		products		多商品订单
			product_id	商品ID
			product_name	商品名称
			product_price	商品价格(单位 : 分)
			product_sku	商品SKU
			product_count	商品个数
			product_img	商品图片

*/

/*
6.3	根据订单状态/创建时间获取订单详情 order/getbyfilter
	协议	https
	http请求方式	GET
	请求Url	https://api.weixin.qq.com/merchant/order/getbyfilter?access_token=ACCESS_TOKEN
	POST数据格式	json
	6.3.2	请求参数说明
	参数	是否必须	说明
	access_token	是	公众号的调用接口凭证
	6.3.3	POST数据
	数据示例：
	{
	"status": 2,
	"begintime": 1397130460,
	"endtime": 1397130470
	}
	字段	说明
	status	订单状态(不带该字段-全部状态, 2-待发货, 3-已发货, 5-已完成, 8-维权中, )
	begintime	订单创建时间起始时间(不带该字段则不按照时间做筛选)
	endtime	订单创建时间终止时间(不带该字段则不按照时间做筛选)
	6.3.4	返回数据说明
	数据示例：
	{
		"errcode": 0,
		"errmsg": "success",
		"order_list": [
			{
			  "order_id": "7197417460812533543",
			  "order_status": 6,
			  "order_total_price": 6,
			  "order_create_time": 1394635817,
			  "order_express_price": 5,
			  "buyer_openid": "oDF3iY17NsDAW4UP2qzJXPsz1S9Q",
			  "buyer_nick": "likeacat",
			  "receiver_name": "张小猫",
			  "receiver_province": "广东省",
			  "receiver_city": "广州市",
			  "receiver_address": "华景路一号南方通信大厦5楼",
			  "receiver_mobile": "123456",
			  "receiver_phone": "123456",
			  "product_id": "pDF3iYx7KDQVGzB7kDg6Tge5OKFo",
			  "product_name": "安莉芳E-BRA专柜女士舒适内衣蕾丝3/4薄杯聚拢上托性感文胸KB0716",
			  "product_price": 1,
			  "product_sku": "10000983:10000995;10001007:10001010",
			  "product_count": 1,
			  "product_img": "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl2icND8WwMThBEcehjhDv2icY4GrDSG5RLM3B2qd9kOicWGVJcsAhvXfibhWRNoGOvCfMC33G9z5yQr2Qw/0",
			  "delivery_id": "1900659372473",
			  "delivery_company": "059Yunda",
			  "trans_id": "1900000109201404103172199813",
			  "products": [
			  {
					  "product_id": "p8BCTv77lY4io_q00F9qsaniimFc",
					  "product_name": "product_name",
					  "product_price": 1,
					  "product_sku": "",
					  "product_count": 1,
					  "product_img": "http://mmbiz.qpic.cn/mmbiz_gif/KfrZwACMrmxj8XRiaTUzFNsTkWdTEJySicGKMHxuG0ibDfjTtb6ZIjNgakbnKq569TbBjvicSnWdnt46gEKjWe6Vcg/0?wx_fmt=gif"
			   }
	   ]
			},
			{
			  "order_id": "7197417460812533569",
			  "order_status": 8,
			  "order_total_price": 1,
			  "order_create_time": 1394636235,
			  "order_express_price": 0,
			  "buyer_openid": "oDF3iY17NsDAW4UP2qzJXPsz1S9Q",
			  "buyer_nick": "likeacat",
			  "receiver_name": "张小猫",
			  "receiver_province": "广东省",
			  "receiver_city": "广州市",
			  "receiver_address": "华景路一号南方通信大厦5楼",
			  "receiver_mobile": "123456",
			  "receiver_phone": "123456",
			  "product_id": "pDF3iYx7KDQVGzB7kDg6Tge5OKFo",
			  "product_name": "项坠333",
			  "product_price": 1,
			  "product_sku": "1075741873:1079742377",
			  "product_count": 1,
			  "product_img": "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl2icND8WwMThBEcehjhDv2icY4GrDSG5RLM3B2qd9kOicWGVJcsAhvXfibhWRNoGOvCfMC33G9z5yQr2Qw/0",
			  "delivery_id": "1900659372473",
			  "delivery_company": "059Yunda",
			  "trans_id": "1900000109201404103172199813",
			  "products": [
			  {
					  "product_id": "p8BCTv77lY4io_q00F9qsaniimFc",
					  "product_name": "product_name",
					  "product_price": 1,
					  "product_sku": "",
					  "product_count": 1,
					  "product_img": "http://mmbiz.qpic.cn/mmbiz_gif/KfrZwACMrmxj8XRiaTUzFNsTkWdTEJySicGKMHxuG0ibDfjTtb6ZIjNgakbnKq569TbBjvicSnWdnt46gEKjWe6Vcg/0?wx_fmt=gif"
			   }
	   ]

			}
		]
	}
	字段	说明
	errcode	错误码
	errmsg	错误信息
	order_list	所有订单集合(字段说明详见根据订单ID获取订单详情)

*/

/*
6.4	设置订单发货信息 order/setdelivery
	协议	https
	http请求方式	POST
	请求Url	https://api.weixin.qq.com/merchant/order/setdelivery?access_token=ACCESS_TOKEN
	POST数据格式	json
	6.4.2	请求参数说明
	参数	是否必须	说明
	access_token	是	公众号的调用接口凭证
	POST数据	是	商品物流信息
	6.4.3	POST数据
	数据示例：
	{
		"order_id": "7197417460812533543",
		"delivery_company": "059Yunda",
		"delivery_track_no": "1900659372473"，
	"need_delivery": 1，
	"is_others": 0
	}
	字段	说明
	order_id	订单ID
	delivery_company	物流公司ID(参考《物流公司ID》；
	当need_delivery为0时，可不填本字段；
	当need_delivery为1时，该字段不能为空；
	当need_delivery为1且is_others为1时，本字段填写其它物流公司名称)
	delivery_track_no	运单ID(
	当need_delivery为0时，可不填本字段；
	当need_delivery为1时，该字段不能为空；
	)
	need_delivery	商品是否需要物流(0-不需要，1-需要，无该字段默认为需要物流)
	is_others	是否为6.4.5表之外的其它物流公司(0-否，1-是，无该字段默认为不是其它物流公司)
	6.4.4	返回数据说明
	数据示例：
	{
		"errcode": 0,
		"errmsg": "success"
	}
	字段	说明
	errcode	错误码
	errmsg	错误信息
	6.4.5	附：物流公司ID
	物流公司	ID
	邮政EMS	Fsearch_code
	申通快递	002shentong
	中通速递	066zhongtong
	圆通速递	056yuantong
	天天快递	042tiantian
	顺丰速运	003shunfeng
	韵达快运	059Yunda
	宅急送	064zhaijisong
	汇通快运	020huitong
	易迅快递	zj001yixun

*/

/*
6.5	关闭订单 order/close
	协议	https
	http请求方式	POST
	请求Url	https://api.weixin.qq.com/merchant/order/close?access_token=ACCESS_TOKEN
	POST数据格式	json
	6.5.2	请求参数说明
	参数	是否必须	说明
	access_token	是	公众号的调用接口凭证
	POST数据	是	商品订单信息
	6.5.3	POST数据
	数据示例：
	{
		"order_id": "7197417460812584720"
	}
	字段	说明
	order_id	订单ID
	6.5.4	返回数据说明
	数据示例：
	{
		"errcode": 0,
		"errmsg": "success"
	}
	字段	说明
	errcode	错误码
	errmsg	错误信息

*/

/*
7.1	上传图片 common/upload_img
协议	https
http请求方式	POST
请求Url	https://api.weixin.qq.com/merchant/common/upload_img?access_token=ACCESS_TOKEN&filename=test.png
POST数据	图片
7.1.2	请求参数说明
参数	是否必须	说明
access_token	是	调用接口凭证
filename	是	图片文件名(带文件类型后缀)
POST数据	是	图片数据
7.1.3	返回数据说明
数据示例：
{
    "errcode":0,
"errmsg":"success",
"image_url": 	"http://mmbiz.qpic.cn/mmbiz/4whpV1VZl2ibl4JWwwnW3icSJGqecVtRiaPxwWEIr99eYYL6AAAp1YBo12CpQTXFH6InyQWXITLvU4CU7kic4PcoXA/0"
}
字段	说明
errcode	错误码
errmsg	错误信息
image_url	图片Url
*/
