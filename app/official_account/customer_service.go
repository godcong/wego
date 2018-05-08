package official_account

import (
	"github.com/godcong/wego/core"
)

type CustomerService struct {
	core.Client
	message *core.Message
}

func (c *CustomerService) List() {
	c.HttpGet(c.Link(GETKFLIST_URL_SUFFIX), nil)
}
