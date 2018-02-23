package token

import (
	"log"
	"testing"

	"github.com/godcong/wego"
)

func TestAccessToken_GetToken(t *testing.T) {
	t0 := wego.NewAccessToken(wego.GetApplication(), wego.GetApplication().GetConfig("mini.default"))
	v := t0.GetToken()
	log.Println(v)
}
