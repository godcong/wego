package official

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
	"time"
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

// OnlineList ...
func (c *CustomerService) OnlineList() core.Responder {
	token := c.accessToken.KeyMap()
	return core.Get(Link(customserviceGetonlinekflist), token)
}

// AccountAdd ...
func (c *CustomerService) AccountAdd(account string, nickname string) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(customserviceKfaccountAdd, token, util.Map{
		"kf_account": account,
		"nickname":   nickname,
	})
}

// AccountUpdate ...
func (c *CustomerService) AccountUpdate(account string, nickname string) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(customserviceKfaccountUpdate, token, util.Map{
		"kf_account": account,
		"nickname":   nickname,
	})
}

// AccountDelete ...
func (c *CustomerService) AccountDelete(account string) core.Responder {
	token := c.accessToken.KeyMap()
	token.Set("kf_account", account)
	return core.PostJSON(customserviceKfaccountDel, token, util.Map{})
}

// AccountInviteWorker ...
func (c *CustomerService) AccountInviteWorker(account, wechatID string) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(customserviceKfaccountInviteworker, token, util.Map{
		"kf_account": account,
		"invite_wx":  wechatID,
	})
}

// AccountUploadHeadImg ...
func (c *CustomerService) AccountUploadHeadImg(account, path string) core.Responder {
	token := c.accessToken.KeyMap()
	token.Set("kf_account", account)
	return core.Upload(customserviceKfaccountUploadheadimg, token, util.Map{"media": path})
}

// MessageSend ...
func (c *CustomerService) MessageSend(p util.Map) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(messageCustomSend, token, p)
}

// MessageList ...
func (c *CustomerService) MessageList(startTime, endTime time.Time, msgID, number int) core.Responder {
	token := c.accessToken.KeyMap()
	p := util.Map{
		"starttime": startTime.Unix(),
		"endtime":   endTime.Unix(),
		"msgid":     msgID,
		"number":    number,
	}
	return core.PostJSON(customserviceMsgrecordGetmsglist, token, p)
}

// SessionList ...
func (c *CustomerService) SessionList(account string) core.Responder {
	token := c.accessToken.KeyMap()
	token.Set("kf_account", account)
	return core.Get(customserviceKfsessionGetsessionlist, token)
}

// SessionWaitCase ...
func (c *CustomerService) SessionWaitCase() core.Responder {
	token := c.accessToken.KeyMap()
	return core.Get(customserviceKfsessionGetwaitcase, token)
}

// SessionCreate ...
func (c *CustomerService) SessionCreate(account, openID string) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(customserviceKfsessionCreate, token, util.Map{
		"kf_account": account,
		"openid":     openID,
	})
}

// SessionClose ...
func (c *CustomerService) SessionClose(account, openID string) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON(customserviceKfsessionClose, token, util.Map{
		"kf_account": account,
		"openid":     openID,
	})
}

// SessionGet ...
func (c *CustomerService) SessionGet(openID string) core.Responder {
	token := c.accessToken.KeyMap()
	token.Set("openid", openID)
	return core.Get(customserviceKfsessionGetsession, token)
}
