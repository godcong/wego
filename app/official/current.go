package official

import "github.com/godcong/wego/core"

type Current struct {
	*Account
}

func newCurrent(account *Account) *Current {
	return &Current{
		Account: account,
	}
}

//NewCurrent current
func NewCurrent(config *core.Config) *Current {
	return newCurrent(NewOfficialAccount(config))
}
