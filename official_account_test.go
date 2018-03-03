package wego_test

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOfficialAccount(t *testing.T) {
	//o := wego.GetOfficialAccount().Base()
	//log.Println(o.GetCallbackIp())

}

var msgText = []byte(`<xml>
    <ToUserName> <![CDATA[toUser]]>
    </ToUserName>
    <FromUserName> <![CDATA[fromUser]]>
    </FromUserName>
    <CreateTime>1348831860</CreateTime>
    <MsgType> <![CDATA[text]]>
    </MsgType>
    <Content> <![CDATA[this is a test]]>
    </Content>
    <MsgId>1234567890123456</MsgId>
</xml>`)

var msgImage = []byte(`<xml>
    <ToUserName><![CDATA[toUser]]>
    </ToUserName>
    <FromUserName><![CDATA[fromUser]]>
    </FromUserName>
    <CreateTime>1348831860</CreateTime>
    <MsgType><![CDATA[image]]>
    </MsgType>
    <PicUrl><![CDATA[this is a url]]>
    </PicUrl>
    <MediaId><![CDATA[media_id]]>
    </MediaId>
    <MsgId>1234567890123456</MsgId>
</xml>`)

func TestGetApp(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		//w.Write(messageResult)
		if r.Method != "POST" {
			t.Errorf("Expected 'GET' request, got '%s'", r.Method)
		}
		if r.URL.EscapedPath() != "/callback" {
			t.Errorf("Expected request to '/callback', got '%s'", r.URL.EscapedPath())
		}
		r.ParseForm()
		body, e := ioutil.ReadAll(r.Body)
		if e != nil {
			t.Errorf(e.Error())
		}
		//t.Log(string(body))
		//var msg core.Message
		var msg map[string]interface{}
		e = xml.Unmarshal(body, &msg)
		//xml.NewDecoder(bytes.NewReader(body)).
		t.Error(msg, e)

	}))
	defer ts.Close()

	resp, e := http.Post(ts.URL+"/callback", "Content-Type:application/xml", bytes.NewReader(msgImage))
	log.Println(resp, e)
}
