package wego_test

import (
	"log"
	"testing"

	"github.com/godcong/wego"
)

func TestGetMiniProgram(t *testing.T) {
	log.Println(wego.GetMiniProgram().AppCode().AccessToken().GetToken())
}

func TestGetAuth(t *testing.T) {
	log.Println(wego.GetMiniProgram().Auth().Session("1234"))
}

func TestGetAppCode(t *testing.T) {
	log.Println(wego.GetAppCode().Get("path", nil))
}

func TestNewAppCode(t *testing.T) {
	var v []string
	log.Println(v == nil)
}
