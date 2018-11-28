package official

import (
	"bytes"
	"encoding/xml"
	"github.com/godcong/wego/cipher"
	"github.com/godcong/wego/core/message"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Notify ...
type Notify interface {
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

// NotifyCallback ...
type NotifyCallback func(p util.Map) (util.Map, error)

// NotifyFunc ...
type NotifyFunc func(w http.ResponseWriter, req *http.Request)

/*Notify 监听 */
type messageNotify struct {
	*Account
	NotifyCallback
	msgType message.MsgType
	bizMsg  *cipher.BizMsg
}

// ServeHTTP ...
func (n *messageNotify) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var err error
	//rlt := SUCCESS()
	//defer func() {
	//	err = NotifyResponseXML(w, rlt.ToXML())
	//	log.Error(err)
	//}()

	maps, err := n.decodeReqInfo(req)
	if err != nil {
		log.Error(err)
		return
	}
	r, err := n.NotifyCallback(maps)
	if err != nil {
		log.Error(err)
		return
	}
	_, err = w.Write(r.ToXML())

	if err != nil {
		log.Error(err)
		return
	}
}

// DecodeReqInfo ...
func (n *messageNotify) decodeReqInfo(req *http.Request) (util.Map, error) {
	maps := util.Map{}
	var bodyBytes []byte
	var err error
	if req.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(req.Body)
	}
	// Restore the io.ReadCloser to its original state
	req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	query, err := url.ParseQuery(req.URL.RawQuery)
	//错误返回,并记录log
	if err != nil {
		log.Error(err)
		return nil, err
	}
	encryptType := query.Get("encrypt_type")
	timeStamp := query.Get("timestamp")
	nonce := query.Get("nonce")
	msgSignature := query.Get("msg_signature")
	if encryptType == "aes" {
		bodyBytes, err = n.bizMsg.Decrypt(string(bodyBytes), msgSignature, timeStamp, nonce)
		//错误返回,并记录log
		if err != nil {
			log.Error(err)
			return nil, err
		}
	}

	err = xml.Unmarshal(bodyBytes, &maps)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return maps, err
}
