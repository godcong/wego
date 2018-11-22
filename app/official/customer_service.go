package official

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
)

//CustomerService CustomerService
type CustomerService struct {
	*Account
}

func newCustomerService(acc *Account) *CustomerService {
	return &CustomerService{
		Account: acc,
	}
}

//NewCustomerService 新建CustomerService
func NewCustomerService(config *core.Config) *CustomerService {
	return newCustomerService(NewOfficialAccount(config))
}

//List ...
func (c *CustomerService) List() core.Responder {
	token := c.accessToken.KeyMap()
	return core.Get(Link(customserviceGetkflist), token)
}

func (c *CustomerService) OnlineList() core.Responder {
	token := c.accessToken.KeyMap()
	return core.Get(Link(customserviceGetonlinekflist), token)
}

func (c *CustomerService) AccountAdd(account string, nickname string) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON("/customservice/kfaccount/add", token, util.Map{
		"kf_account": account,
		"nickname":   nickname,
	})
}
