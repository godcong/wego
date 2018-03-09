package core_test

import (
	"log"
	"testing"

	"github.com/godcong/wego/core"
)

func TestURL_ShortUrl(t *testing.T) {
	c := core.NewClient(core.GetConfig("official_account.default"))
	url := core.NewURL(core.GetConfig("official_account.default"), c)
	log.Println(url.ShortUrl("https://y11e.com"))
}
