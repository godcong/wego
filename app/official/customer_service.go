package official

import "github.com/godcong/wego/core"

/*CustomerService CustomerService*/
type CustomerService struct {
	core.Client
	message *core.Message
}

/*List List */
func (c *CustomerService) List() {
	c.HttpGet(c.Link(GetKFListUrlSuffix), nil)
}
