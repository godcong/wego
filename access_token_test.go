package wego_test

import (
	"log"
	"testing"

	"github.com/godcong/wego"
)

func TestAccessToken_GetToken(t *testing.T) {
	v := wego.NewAccessToken(wego.GetApplication(), wego.GetApplication().GetConfig("mini_program.default")).GetToken()
	log.Println(v)
}
