package official_account_test

import (
	"testing"
	"time"

	"github.com/godcong/wego/app/official_account"
	"github.com/godcong/wego/util"
)

var card = official_account.NewCard()

//step 1:
//{
//"access_token": "8_lqfGbYyDDxEsIt17ShuQTe4pBgELHTv0O_-irWYo8-PMvSDi27A19-GtK1WCJrHSJwpujjrODPDclZlNTj8TQAAlirrGuEniDmYHycIiKasg3jOGZNncP8cl225DBBOmNpwaGhOEaSdHqRpVNKNjAAASPS",
//"expires_in": 7200
//}
//step 2:
//{
//"url": "http://mmbiz.qpic.cn/mmbiz_jpg/N0dLE3fL5dpH5SDtAVxq6ZjFbYkeL0PLc1q8GXETsFiaxxMSkiabshuJouPx6U1egRM3aibBXRptxWGlQ5NAvxibZw/0"
//}

func TestCard_CreateLandingPage(t *testing.T) {
	resp := card.CreateLandingPage(
		&official_account.CardLandingPage{
			Banner:   "http://mmbiz.qpic.cn/mmbiz/iaL1LJM1mF9aRKPZJkmG8xXhiaHqkKSVMMWeN3hLut7X7hicFN",
			Title:    "惠城优惠大派送",
			CanShare: true,
			Scene:    "SCENE_NEAR_BY",
			CardList: []official_account.CardList{
				{
					CardId:   "p5jo6s63aXgMQCt90t87UXA1dMJk",
					ThumbUrl: "www.qq.com/a.jpg",
				},
			},
		},
	)
	t.Log(resp.ToString())
}

func TestCard_Create(t *testing.T) {
	oc := official_account.NewOneCard(official_account.CARD_TYPE_GROUPON, nil)
	//oc.AddAdvancedInfo(&official_account.CardAdvancedInfo{
	//	UseCondition:    &official_account.CardUseCondition{},
	//	Abstract:        &official_account.CardAbstract{},
	//	TextImageList:   nil,
	//	TimeLimit:       nil,
	//	BusinessService: nil,
	//})
	oc.AddBaseInfo(&official_account.CardBaseInfo{
		LogoUrl:      "http://mmbiz.qpic.cn/mmbiz_jpg/gJHMd2C74Xq9iaaWAksvY6hd4LibvPfxoj3UYyKLt3DRqicMhEHhftPJ0PbJ0CqzzjicLBibN4nibqaa3H6CkibiaAvyYg/0",
		BrandName:    "微信餐厅",
		CodeType:     "CODE_TYPE_TEXT",
		Title:        "132元双人火锅套餐",
		Color:        "Color010",
		Notice:       "使用时向服务员出示此券",
		ServicePhone: "020-88888888",
		Description:  "不可与其他优惠同享\n如需团购券发票，请在消费时向商户提出\n店内均可使用，仅限堂食",
		DateInfo: official_account.CardDataInfo{
			Type:           "DATE_TYPE_FIX_TERM",
			BeginTimestamp: time.Now().Unix(),
			EndTimestamp:   time.Now().Unix(),
			FixedTerm:      15,
			FixedBeginTerm: 0,
		},
		//Sku:                       official_account.CardSku{},
		UseLimit:                  5,
		GetLimit:                  5,
		UseCustomCode:             true,
		GetCustomCodeMode:         "GET_CUSTOM_CODE_MODE_DEPOSIT",
		BindOpenid:                false,
		CanShare:                  false,
		CanGiveFriend:             false,
		LocationIdList:            nil,
		UseAllLocations:           false,
		CenterTitle:               "",
		CenterSubTitle:            "",
		CenterUrl:                 "",
		CenterAppBrandUserName:    "",
		CenterAppBrandPass:        "",
		CustomUrlName:             "",
		CustomUrl:                 "",
		CustomUrlSubTitle:         "",
		CustomAppBrandUserName:    "",
		CustomAppBrandPass:        "",
		PromotionUrlName:          "",
		PromotionUrl:              "",
		PromotionUrlSubTitle:      "",
		PromotionAppBrandUserName: "",
		PromotionAppBrandPass:     "",
		Source:                    "",
	})
	oc.AddDealDetail("this is a test")
	resp := card.Create(oc)
	t.Log(resp.ToString())
}

func TestCard_Deposit(t *testing.T) {
	resp := card.Deposit("p5jo6s0HmChHznHQ75T7FWfbKljw", []string{
		"11111",
		"22222",
		"33333",
		"44444",
		"55555",
	})
	t.Log(resp.ToString())
}

func TestCard_GetDepositCount(t *testing.T) {
	resp := card.GetDepositCount("pDF3iY0_dVjb_Pua96MMewA96qvA")
	t.Log(resp.ToString())
}

func TestCard_CheckCode(t *testing.T) {
	resp := card.CheckCode("pDF3iY0_dVjb_Pua96MMewA96qvA", []string{
		"11111",
		"22222",
		"33333",
		"44444",
		"55555",
	})
	t.Log(resp.ToString())
}

func TestCard_GetHtml(t *testing.T) {
	resp := card.GetHtml("oLyBi0hSYhggnD-kOIms0IzZFqrc")
	t.Log(resp.ToString())
}

func TestCard_SetTestWhiteList(t *testing.T) {
	resp := card.SetTestWhiteList("openid", []string{"o5jo6s3RZ6rxuVAW33IpTjYWQOg4"})
	t.Log(resp.ToString())
}

func TestCard_CreateQrCode(t *testing.T) {
	resp := card.CreateQrCode(&official_account.QrCodeAction{
		ExpireSeconds: 1800,
		ActionName:    "QR_CARD",
		ActionInfo: official_account.QrCodeActionInfo{
			Card: &official_account.QrCodeCard{
				CardId: "p5jo6s57F-21KqvrBXRssSEjXtWg",
				//Code:         "198374613512",
				//OpenId:       "oFS7Fjl0WsZ9AMZqrI80nbIq8xrA",
				//IsUniqueCode: false,
				//OuterStr:     "12b",
			},
			//MultipleCard: &official_account.QrCodeMultipleCard{
			//	CardList: []official_account.QrCodeCardList{
			//		{
			//			CardId:   "pFS7Fjg8kV1IdDz01r4SQwMkuCKc",
			//			Code:     "198374613512",
			//			OuterStr: "12b",
			//		},
			//	},
			//},
		},
	})
	t.Log(resp.ToString())
}

func TestCard_GetCode(t *testing.T) {
	resp := card.GetCode(util.Map{
		"card_id": "card_id_123+",
	})
	t.Log(resp.ToString())
}
