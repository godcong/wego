package payment_test

import (
	"github.com/godcong/wego"
	"github.com/godcong/wego/util"
	"net/http"
	"testing"
)

func TestScannedNotify_ServeHTTP(t *testing.T) {
	ScannedCallbackFunction := func(p util.Map) (maps util.Map, e error) {
		return util.Map{"prepay_id": "wxxxxxxxxxxxxx"}, nil
	}

	serve1 := wego.Payment().HandleScannedNotify(ScannedCallbackFunction).ServeHTTP

	serve2 := wego.Payment().HandleScanned(ScannedCallbackFunction)

	http.HandleFunc("/scanned/callback/address", serve1)
	http.HandleFunc("/scanned/callback/address2", serve2)

	t.Fatal(http.ListenAndServe(":8080", nil))
}
