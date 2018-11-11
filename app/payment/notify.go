package payment

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
	"net/http"
)

// Notify ...
type Notify interface {
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

// NotifyFunc ...
type NotifyFunc func(p util.Map) error

/*Notify 监听 */
type refundedNotify struct {
	*Payment
	NotifyFunc
}

// ServeHTTP ...
func (n *refundedNotify) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var err error
	rlt := SUCCESS()
	maps, err := core.RequestToMap(req)
	//wrong request will do nothing
	if err != nil {
		log.Error(err)
		return
	}

	if ValidateSign(maps, n.GetKey()) {
		err = n.NotifyFunc(maps)
		if err != nil {
			rlt = Fail(err.Error())
		}
	}

	err = NotifyResponseXML(w, rlt.ToXML())
	if err != nil {
		log.Error(err)
	}
}

/*Notify 监听 */
type scannedNotify struct {
	*Payment
	NotifyFunc
}

// ServeHTTP ...
func (n *scannedNotify) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	panic("implement me")
}

/*Notify 监听 */
type paidNotify struct {
	*Payment
	NotifyFunc
}

// ServerHttp ...
func (n *paidNotify) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var err error
	rlt := SUCCESS()
	maps, err := core.RequestToMap(req)
	//wrong request will do nothing
	if err != nil {
		log.Error(err)
		return
	}

	if ValidateSign(maps, n.GetKey()) {
		err = n.NotifyFunc(maps)
		if err != nil {
			rlt = Fail(err.Error())
		}
	}

	err = NotifyResponseXML(w, rlt.ToXML())
	if err != nil {
		log.Error(err)
	}
}

// NotifyResponseXML ...
func NotifyResponseXML(w http.ResponseWriter, data []byte) error {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{"application/json; charset=utf-8"}
	}
	_, err := w.Write(data)
	if err != nil {
		return err
	}
	return nil
}

// SUCCESS ...
func SUCCESS() util.Map {
	return util.Map{
		"return_code": "SUCCESS",
		"return_msg":  "OK",
	}
}

// Fail ...
func Fail(msg string) util.Map {
	return util.Map{
		"return_code": "FAIL",
		"return_msg":  msg,
	}
}
