package core_test

import (
	"github.com/godcong/wego/core"
	"testing"
)

func TestLink(t *testing.T) {
	t.Log(core.Link("/cgi-bin/customservice/getonlinekflist"))
}
