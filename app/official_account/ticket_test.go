package official_account_test

import (
	"testing"

	"github.com/godcong/wego/official_account"
)

var ticket = official_account.NewTicket()

func TestTicket_Get(t *testing.T) {
	resp := ticket.Get("wx_card")
	t.Log(resp.ToString())
}
