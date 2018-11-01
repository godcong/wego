# official account 功能模块介绍 #
    obj := wego.OfficialAccount()
    obj := App().OfficialAccount("official_account.default") //选择加载对应的toml配置
#基础接口:
    
## 公众号的所有api调用（包括第三方帮其调用）次数进行清零

    obj.Base().ClearQuota() //公众号的所有api调用（包括第三方帮其调用）次数进行清零 PS:此接口官方有每月调用限制
    
## 请求微信的服务器IP列表

    obj.Base().GetCallbackIP