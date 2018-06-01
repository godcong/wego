# official account 功能模块介绍 #

基本功能模块
base:= official_account.NewBase()

公众号调用或第三方平台帮公众号调用对公众号的所有api调用（包括第三方帮其调用）次数进行清零
接口说明：<https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1433744592>
>
base.ClearQuota()

自定义菜单
menu := official_account.NewMenu()

自定义菜单创建接口
接口说明：<https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421141013>
>
    menu.Create()





消息管理  
NewMessage()  

客服帐号管理  
NewCustomerService()  

微信网页授权  
NewOAuth()  

素材管理  
NewMedia()  

图文消息留言管理  
NewMaterial()  

标签管理  
NewTag()  

生成带参数的二维码  
NewQrCode()  

数据统计
NewDataCube()

微信卡券
NewCard()
