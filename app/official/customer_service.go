package official

import "github.com/godcong/wego/core"

/*CustomerService CustomerService*/
type CustomerService struct {
	message *core.Message
}

/*List List */
func (c *CustomerService) List() {
	core.Get(Link(getKFListURLSuffix), nil)
}
