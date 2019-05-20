package wego

import (
	"context"
	"github.com/godcong/wego/util"
	jsoniter "github.com/json-iterator/go"
	"log"
)

//UserUpdateRemark 设置用户备注名
// http请求方式: POST（请使用https协议）
// https://api.weixin.qq.com/cgi-bin/user/info/updateremark?access_token=ACCESS_TOKEN
// POST数据格式:JSON
// POST数据例子:
// {
// "openid":"oDF3iY9ffA-hqb2vVvbr7qxf6A0Q",
// "remark":"pangzi"
// }
// 成功:
// {"errcode":0,"errmsg":"ok"}
// 失败:
// {"errcode":40013,"errmsg":"invalid appid"}
func (obj *OfficialAccount) UserUpdateRemark(openid, remark string) Responder {
	log.Debug("OfficialAccount|UserUpdateRemark", openid, remark)
	u := util.URL(userInfoUpdateRemark)
	return obj.Client().Post(context.Background(), u, nil, util.Map{"openid": openid, "remark": remark})
}

//UserInfo 获取用户信息
// 接口调用请求说明
// http请求方式: GET
// https://api.weixin.qq.com/cgi-bin/user/info?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN
// 成功:
// {"subscribe":1,"openid":"o6_bmjrPTlm6_2sgVt7hMZOPfL2M","nickname":"Band","sex":1,"language":"zh_CN","city":"广州","province":"广东","country":"中国","headimgurl":"http://thirdwx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/0","subscribe_time":1382694957,"unionid":"o6_bmasdasdsad6_2sgVt7hMZOPfL""remark":"","groupid":0,"tagid_list":[128,2],"subscribe_scene":"ADD_SCENE_QR_CODE","qr_scene":98765,"qr_scene_str":""}
func (obj *OfficialAccount) UserInfo(openid, lang string) (info *UserInfo, e error) {
	log.Debug("OfficialAccount|UserInfo", openid, lang)
	p := util.Map{"openid": openid}
	if lang != "" {
		p.Set("lang", lang)
	}
	u := util.URL(userInfo)
	resp := obj.Client().Get(context.Background(), u, p)
	if e = resp.Error(); e != nil {
		return nil, e
	}
	info = new(UserInfo)
	e = jsoniter.Unmarshal(resp.Bytes(), info)
	if e != nil {
		return nil, e
	}
	return info, nil
}

//UserBatchGet 批量获取用户基本信息
// http请求方式: POST
// https://api.weixin.qq.com/cgi-bin/user/info/batchget?access_token=ACCESS_TOKEN
// 成功:
// {"user_info_list":[{"subscribe":1,"openid":"oLyBi0tDnybg0WFkhKsn5HRetX1I","nickname":"sean","sex":1,"language":"zh_CN","city":"浦东新区","province":"上海","country":"中国","headimgurl":"http:\/\/thirdwx.qlogo.cn\/mmopen\/anblvjPKYbMGjBnTVxw5gEZiasF6LiaMHheNxN4vWJcfCLRl8gEX0L6M7sNjtMkFYx8PJRCS1lr9RGxadkFlBibpA\/132","subscribe_time":1521022410,"remark":"nishi123","groupid":101,"tagid_list":[101],"subscribe_scene":"ADD_SCENE_PROFILE_CARD","qr_scene":0,"qr_scene_str":""}]}
// 失败:
// {"errcode":40013,"errmsg":"invalid appid"}
func (obj *OfficialAccount) UserBatchGet(openids []string, lang string) (infos []*UserInfo, e error) {
	log.Debug("User|BatchGet", openids, lang)
	u := util.URL(userInfoBatchGet)
	var list []*UserID
	for _, v := range openids {
		if lang != "" {
			list = append(list, &UserID{OpenID: v, Lang: lang})
		} else {
			list = append(list, &UserID{OpenID: v})
		}

	}
	resp := obj.Client().Post(context.Background(), u, nil, util.Map{"user_list": list})
	if e = resp.Error(); e != nil {
		return nil, e
	}
	infoList := UserInfoList{}
	e = jsoniter.Unmarshal(resp.Bytes(), &infoList)
	if e != nil {
		return nil, e
	}
	return infoList.UserInfoList, nil
}

//UserGet 获取用户列表
// http请求方式: GET（请使用https协议）
// https://api.weixin.qq.com/cgi-bin/user/get?access_token=ACCESS_TOKEN&next_openid=NEXT_OPENID
func (obj *OfficialAccount) UserGet(nextOpenid string) Responder {
	log.Debug("OfficialAccount|UserGet", nextOpenid)
	u := util.URL(userGet)
	if nextOpenid == "" {
		return obj.Client().Get(context.Background(), u, nil)
	}
	return obj.Client().Get(context.Background(), u, util.Map{"next_openid": nextOpenid})
}
