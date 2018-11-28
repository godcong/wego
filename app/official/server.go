package official

import (
	"bytes"
	"encoding/xml"
	"github.com/godcong/wego/cipher"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/core/message"
	"github.com/godcong/wego/log"
	"io/ioutil"
	"net/http"
	"net/url"
)

/*Server Server */
type Server struct {
	*Account
	CryptResponse bool
	//message         *core.Message
	msgType         string
	bizMsg          *cipher.BizMsg
	defaultCallback []core.MessageCallback
	callback        map[message.MsgType][]core.MessageCallback
}

/*RegisterCallback RegisterCallback */
func (s *Server) RegisterCallback(sc core.MessageCallback, types ...message.MsgType) {
	size := len(types)
	if size == 0 {
		s.defaultCallback = append(s.defaultCallback, sc)
		return
	}
	for _, t := range types {
		if callback, b := s.callback[t]; b {
			s.callback[t] = append(callback, sc)
		} else {
			s.callback[t] = []core.MessageCallback{sc}
		}
	}
}

// ServeHTTP ...
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var bodyBytes []byte
	var rltXML []byte
	var err error
	if req.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(req.Body)
	}
	// Restore the io.ReadCloser to its original state
	req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	//无数据直接返回
	w.WriteHeader(http.StatusOK)
	if len(bodyBytes) == 0 {
		return
	}
	query, err := url.ParseQuery(req.URL.RawQuery)
	//错误返回,并记录log
	if err != nil {
		log.Error(err)
		return
	}
	encryptType := query.Get("encrypt_type")
	ts := query.Get("timestamp")
	nonce := query.Get("nonce")
	msgSignature := query.Get("msg_signature")

	if encryptType == "aes" {
		log.Debug(ts, nonce, msgSignature, string(bodyBytes))
		bodyBytes, err = s.bizMsg.Decrypt(string(bodyBytes), msgSignature, ts, nonce)
		//错误返回,并记录log
		if err != nil {
			log.Error(err)
			return
		}
	}

	message := new(core.Message)
	log.Debug(string(bodyBytes))
	err = xml.Unmarshal(bodyBytes, message)
	//错误返回,并记录log
	if err != nil {
		log.Error(err)
		return
	}
	result := s.CallbackFunc(message)

	rltXML, err = result.ToXML()
	//错误返回,并记录log
	if err != nil {
		log.Error(err)
		return
	}

	//if encryptType == "aes" {
	//	tmpStr, err := s.bizMsg.RSAEncrypt(string(rltXML), ts, nonce)
	//	if err != nil {
	//		log.Error(err)
	//		return
	//	}
	//	rltXML = []byte(tmpStr)
	//}
	if s.msgType == "xml" {
		header := w.Header()
		if val := header["Content-Type"]; len(val) == 0 {
			header["Content-Type"] = []string{"application/xml; charset=utf-8"}
		}
	} else {
		header := w.Header()
		if val := header["Content-Type"]; len(val) == 0 {
			header["Content-Type"] = []string{"application/json; charset=utf-8"}
		}
	}
	log.Debug(string(rltXML))
	w.Write(rltXML)
	return
}

/*CallbackFunc message回调函数*/
func (s *Server) CallbackFunc(msg *core.Message) message.Messager {
	var result message.Messager
	for _, v := range s.defaultCallback {
		if r := v(msg); r != nil {
			result = r
		}
	}

	if v0, b := s.callback[msg.GetType()]; b {
		for _, v := range v0 {
			if r := v(msg); r != nil {
				result = r
			}
		}
	}
	return result
}

func newServer(account *Account) *Server {
	token := account.GetString("token")
	key := account.GetString("aes_key")
	id := account.GetString("app_id")

	return &Server{
		msgType:         "xml",
		bizMsg:          cipher.NewBizMsg(token, key, id),
		defaultCallback: []core.MessageCallback{},
		callback:        map[message.MsgType][]core.MessageCallback{},
	}
}

/*NewServer NewServer*/
func NewServer(config *core.Config) *Server {
	return newServer(NewOfficialAccount(config))
}
