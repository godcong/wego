package mini_program_test

import (
	"testing"

	"github.com/godcong/wego/mini_program"
)

var auth = mini_program.NewAuth()

func TestAuth_Session(t *testing.T) {
	resp := auth.Session("0022IX8c1OPfgv0tOQ6c1tGZ8c12IX8E")
	t.Log(resp.String())
}
