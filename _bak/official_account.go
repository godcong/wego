package _bak

import (
	"github.com/godcong/wego/app/official"
)

//OfficialAccount 公众号*/
func OfficialAccount() *official.Account {
	return App().OfficialAccount("official_account.default")
}
