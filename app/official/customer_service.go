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

func (c *CustomerService) AccountUpdate(account string, nickname string) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON("/customservice/kfaccount/update", token, util.Map{
		"kf_account": account,
		"nickname":   nickname,
	})
}

func (c *CustomerService) AccountDelete(account string) core.Responder {
	token := c.accessToken.KeyMap()
	token.Set("kf_account", account)
	return core.PostJSON("/customservice/kfaccount/del", token, util.Map{})
}

func (c *CustomerService) AccountInviteWorker(account, wechatID string) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON("/customservice/kfaccount/inviteworker", token, util.Map{
		"kf_account": account,
		"invite_wx":  wechatID,
	})
}

func (c *CustomerService) AccountUploadHeadImg(account, path string) core.Responder {
	token := c.accessToken.KeyMap()
	token.Set("kf_account", account)
	return core.Upload("customservice/kfaccount/uploadheadimg", token, util.Map{"media": path})
}

func (c *CustomerService) MessageSend(p util.Map) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON("cgi-bin/message/custom/send", token, p)
}

func (c *CustomerService) MessageList(startTime, endTime time.Time, msgId, number int) core.Responder {
	token := c.accessToken.KeyMap()
	p := util.Map{
		"starttime": startTime.Unix(),
		"endtime":   endTime.Unix(),
		"msgid":     msgId,
		"number":    number,
	}
	return core.PostJSON("customservice/msgrecord/getmsglist", token, p)
}

func (c *CustomerService) SessionList(account string) core.Responder {
	token := c.accessToken.KeyMap()
	token.Set("kf_account", account)
	return core.Get("customservice/kfsession/getsessionlist", token)
}

func (c *CustomerService) SessionWaitCase() core.Responder {
	token := c.accessToken.KeyMap()
	return core.Get("customservice/kfsession/getwaitcase", token)
}

func (c *CustomerService) SessionCreate(account, openID string) core.Responder {
	token := c.accessToken.KeyMap()
	return core.PostJSON("customservice/kfsession/create", token, util.Map{
		"kf_account": account,
		"openid":     openID,
	})
}
