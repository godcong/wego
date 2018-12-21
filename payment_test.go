package wego_test

import (
	"github.com/godcong/wego/util"
	"testing"
)

// TestPayment ...
func TestPayment(t *testing.T) {
	t.Log(util.CurrentTimeStampMS())
	t.Log(util.CurrentTimeStampNS())
	t.Log(util.CurrentTimeStamp())
	t.Log(util.CurrentTimeStampString())
}
