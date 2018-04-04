package mini_program_test

import (
	"testing"

	"github.com/godcong/wego/mini_program"
)

var auth = mini_program.NewAuth()

func TestAuth_Session(t *testing.T) {
	resp := auth.Session("002R4kYz1aGBfe0HWgYz1O4mYz1R4kY4")
	t.Log(resp.String())
}
