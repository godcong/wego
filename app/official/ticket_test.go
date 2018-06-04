package official_test

import (
	"testing"

	"github.com/godcong/wego/app/official"
)

var ticket = official.NewTicket()

func TestTicket_Get(t *testing.T) {
	resp := ticket.Get("wx_card")
	t.Log(resp.ToString())
}
