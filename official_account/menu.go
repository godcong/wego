package official_account

import "github.com/godcong/wego/core"

type Menu struct {
}

//个性化创建
//https://api.weixin.qq.com/cgi-bin/menu/addconditional?access_token=ACCESS_TOKEN
//自定义菜单
//https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN
func (m *Menu) Create(p core.Map, additional core.Map) {

}

func (m *Menu) List() {

}

func (m *Menu) Get() {

}

//https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=ACCESS_TOKEN
func (m *Menu) Delete() {

}
