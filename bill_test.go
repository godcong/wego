package wego_test

import (
	"log"
	"testing"

	"github.com/godcong/wego"
)

func TestBill_GetBill(t *testing.T) {
	r := wego.GetBill().GetBill("20140603", "ALL", nil)
	log.Println(r)
}
