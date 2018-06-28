# official account 功能模块介绍 #

基本功能模块：
    base:= official.NewBase()

公众号调用或第三方平台帮公众号调用对公众号的所有api调用（包括第三方帮其调用）次数进行清零
> 接口说明：<https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1433744592>

    base.ClearQuota()

自定义菜单
    menu := official.NewMenu()

自定义菜单创建接口
> 接口说明：<https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421141013>

	    button := menu.NewButton()
        var subs0 []*menu.Button
        subs0 = append(subs0, menu.NewViewButton("sub1", "key id"))
        subs0 = append(subs0, menu.NewClickButton("sub2", "url")

        sub0 := menu.NewSubButton("button1", subs0)
        button.AddButton(sub0)

        var subs1 []*menu.Button
        subs1 = append(subs1, menu.NewViewButton("动态首页", "https://eagle.bitbili.top"))
        sub1 := menu.NewSubButton("button2", subs1)
        button.AddButton(sub1)

        button.AddButton(menu.NewViewButton("button3", "url"))

        //添加个性化菜单
        //tag_id	否	用户标签的id，可通过用户标签管理接口获取
        //sex	否	性别：男（1）女（2），不填则不做匹配
        //client_platform_type	否	客户端版本，当前只具体到系统型号：IOS(1), Android(2),Others(3)，不填则不做匹配
        //country	否	国家信息，是用户在微信中设置的地区，具体请参考地区信息表
        //province	否	省份信息，是用户在微信中设置的地区，具体请参考地区信息表
        //city	否	城市信息，是用户在微信中设置的地区，具体请参考地区信息表
        //language	否	语言信息，是用户在微信中设置的语言，具体请参考语言表：http://wximg.gtimg.com/shake_tv/mpwiki/areainfo.zip
        mr := menu.MatchRule{
                TagId:              #tag_id#,
                Sex:                #sex#,
                Country:            #country#,
                Province:           #province#,
                City:               #city#,
                ClientPlatformType: #client_platform_type#,
                Language:           #language#,
            }
        button.SetMatchRule(&mr)

        //创建Menu
        menu.Create(button)


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
