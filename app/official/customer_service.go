package official

import "github.com/godcong/wego/core"

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
	return core.Get(Link(getKFListURLSuffix), token)
}
