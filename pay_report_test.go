package wxpay_test

import (
	"bytes"
	"log"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/godcong/wopay/wxpay"
)

func BenchmarkNewReportInfo(b *testing.B) {
	sta := time.Now().UnixNano()
	s := ""
	for i := 10000; i > 0; i-- {
		s = strings.Join([]string{s, strconv.FormatInt(sta, 10)}, ",")
	}
	mid := time.Now().UnixNano()
	log.Println(mid - sta)
	buff := bytes.NewBufferString("")
	for i := 10000; i > 0; i-- {
		s = strings.Join([]string{s, strconv.FormatInt(sta, 10)}, ",")
		buff.Write([]byte(s))
	}
	end := time.Now().UnixNano()
	log.Println(end - mid)

}

func TestReportInfo_ToLineString(t *testing.T) {
	log.Println(wxpay.CurrentTimeStamp())
	info := wxpay.NewReportInfo("uuid", int64(1504178107), int64(1000), "firstDomain", true, 1000, 1000, true, true, true)
	log.Println(info.ToLineString("2ab9071b06b9f739b950ddb41db2690d"))
	//v0,wxpay go sdk v1.0,uuid,1504178107,1000,firstDomain,true,1000,1000,1,1,1,4593BAA7FEACF3DE98EEAFDB5907291966589ED0BA2327E7852B9BAF00C48913
}
