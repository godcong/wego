package core_test

import (
	"log"
	"testing"

	"github.com/godcong/wego/core"
)

func TestSandbox_GetKey(t *testing.T) {
	box := core.NewSandbox(core.GetConfig("official_account.default"))

	log.Println(box.GetKey())
}
