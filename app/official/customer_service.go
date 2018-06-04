package official

import "github.com/godcong/wego/core"

/*CustomerService CustomerService*/
type CustomerService struct {
	core.Client
	message *core.Message
}

func (c *CustomerService) List() {
	c.HttpGet(c.Link(GetKFListUrlSuffix), nil)
}
