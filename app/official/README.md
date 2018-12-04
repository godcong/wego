# official account 功能模块介绍 #
    obj := wego.OfficialAccount()
    obj := App().OfficialAccount("official_account.default") //选择加载对应的toml配置
##基础接口:
    
### 公众号的所有api调用（包括第三方帮其调用）次数进行清零

    obj.Base().ClearQuota() //公众号的所有api调用（包括第三方帮其调用）次数进行清零 PS:此接口官方有每月调用限制
    
### 请求微信的服务器IP列表

    obj.Base().GetCallbackIP()
    
## JSSDK    

    生成js配置文件：
    obj.JSSDK().BuildConfig(util.Map{
        "jsApiList":[]string{"onMenuShareQQ","onMenuShareWeibo"},
    }) 

    结果如下：
    <script src="http://res.wx.qq.com/open/js/jweixin-1.2.0.js" type="text/javascript" charset="utf-8"></script>
    <script type="text/javascript" charset="utf-8">
    wx.config({
        debug: true,
        appId: 'wx3cf0f39249eb0e60',
        timestamp: 1430009304,
        nonceStr: 'qey94m021ik',
        signature: '4F76593A4245644FAE4E1BC940F6422A0C3EC03E',
        jsApiList: ['onMenuShareQQ', 'onMenuShareWeibo']
    });
    </script>
    
## 获取用户信息

    user := obj.User().Get(#openId#);