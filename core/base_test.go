package core_test

import (
	"log"
	"testing"

	"github.com/godcong/wego/core"
)

func TestLocalAddress(t *testing.T) {
	core.GetServerIp()
}

func TestXmlToMap(t *testing.T) {
	m := core.XmlToMap([]byte(`<xml>
<return_code><![CDATA[FAIL]]></return_code>
<return_msg><![CDATA[CERT_ERR]]></return_msg>
</xml>`))
	log.Println(m)
}
