package core

import (
	"bytes"
	"github.com/godcong/wego/core/message"
	"github.com/godcong/wego/util"
	"io/ioutil"
	"net/http"
)

/*MessageCallback 消息回调函数定义 */
type MessageCallback func(message *Message) message.Messager

/*PaymentCallback 支付回调函数定义 */
type PaymentCallback func(p util.Map) bool

/*WriteAble WriteAble*/
type WriteAble interface {
	ToBytes() []byte
}

// Server ...
type Server struct {
}

// ServeHTTP ...
func (s *Server) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	var respData []byte
	var reqBody []byte
	var err error
	defer writer.Write(respData)
	if req.Body == nil {
		return
	}
	reqBody, err = ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}

	req.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
}
