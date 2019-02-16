package wego

import (
	"encoding/xml"
	"github.com/godcong/wego/cipher"
	"github.com/godcong/wego/util"
	"github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"net/http"
)

// Notify ...
type Notify struct {
	payment *Payment
	Key     string
}

// NotifyResult ...
type NotifyResult struct {
	ReturnCode string `json:"return_code" xml:"return_code"`
	ReturnMsg  string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	AppID      string `json:"appid,omitempty" xml:"appid,omitempty"`
	MchID      string `json:"mch_id,omitempty" xml:"mch_id,omitempty"`
	NonceStr   string `json:"nonce_str,omitempty" xml:"nonce_str,omitempty"`
	PrepayID   string `json:"prepay_id,omitempty" xml:"prepay_id,omitempty"`
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	ErrCodeDes string `json:"err_code_des,omitempty" xml:"err_code_des,omitempty"`
	Sign       string `json:"sign,omitempty" xml:"sign,omitempty"`
}

// Notifier ...
type Notifier interface {
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

// ServeNotify ...
type ServeNotify func(req Requester) (util.Map, error)

// ServeHTTPFunc ...
type ServeHTTPFunc func(w http.ResponseWriter, req *http.Request)

// NewNotify ...
func NewNotify(payment *Payment, key string) *Notify {
	return &Notify{
		payment: payment,
		Key:     key,
	}
}

// HandleRefunded ...
func (n *Notify) HandleRefunded(f ServeNotify) Notifier {
	return &paymentRefunded{
		cipher: cipher.New(cipher.AES256ECB, &cipher.Option{
			Key: n.Key,
		}),
		ServeNotify: f,
	}
}

// HandleRefundedNotify ...
func (n *Notify) HandleRefundedNotify(f ServeNotify) ServeHTTPFunc {
	return n.HandleRefunded(f).ServeHTTP
}

// HandleScannedNotify ...
func (n *Notify) HandleScannedNotify(f ServeNotify) Notifier {
	return &paymentScanned{
		Notify:      n,
		ServeNotify: f,
	}
}

// HandleScanned ...
func (n *Notify) HandleScanned(f ServeNotify) ServeHTTPFunc {
	return n.HandleScannedNotify(f).ServeHTTP
}

// HandlePaidNotify ...
func (n *Notify) HandlePaidNotify(f ServeNotify) Notifier {
	return &paymentPaid{
		Notify:      n,
		ServeNotify: f,
	}
}

// ServerHttp ...
func (n *paymentPaid) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var e error
	requester := BuildRequester(req)
	resp := NotifyTypeResponder(requester.Type(), NotifySuccess())
	defer func() {
		e = resp.Write(w)
		log.Error(e)
	}()

	if e = requester.Error(); e != nil {
		log.Error(e.Error())
		resp.SetNotifyResult(NotifyFail(e.Error()))
		return
	}
	reqData := requester.ToMap()
	if util.ValidateSign(reqData, n.payment.GetKey()) {
		if n.ServeNotify == nil {
			log.Error(xerrors.New("null notify callback "))
			return
		}
		_, e = n.ServeNotify(requester)
		if e != nil {
			log.Error(e.Error())
			resp.SetNotifyResult(NotifyFail(e.Error()))
		}
	}

}

/*Notifier 监听 */
type paymentRefunded struct {
	cipher cipher.Cipher
	ServeNotify
}

// ServeHTTP ...
func (obj *paymentRefunded) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var e error
	requester := BuildRequester(req)
	resp := NotifyTypeResponder(requester.Type(), NotifySuccess())
	defer func() {
		e = resp.Write(w)
		log.Error(e)
	}()

	if e = requester.Error(); e != nil {
		log.Error(e.Error())
		resp.SetNotifyResult(NotifyFail(e.Error()))
		return
	}
	reqData := requester.ToMap()
	reqInfo := reqData.GetString("req_info")
	reqData.Set("reqInfo", obj.DecodeReqInfo(reqInfo))
	if obj.ServeNotify == nil {
		log.Error(xerrors.New("null notify callback"))
		return
	}
	_, e = obj.ServeNotify(requester)
	if e != nil {
		log.Error(e.Error())
		resp.SetNotifyResult(NotifyFail(e.Error()))
	}
}

// DecodeReqInfo ...
func (obj *paymentRefunded) DecodeReqInfo(info string) util.Map {
	maps := util.Map{}
	dec, _ := obj.cipher.Decrypt(info)
	e := xml.Unmarshal(dec, &maps)
	if e != nil {
		log.Error(e)
	}
	return maps
}

/*Notifier 监听 */
type paymentScanned struct {
	*Notify
	ServeNotify
}

// ServeHTTP ...
func (obj *paymentScanned) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var e error
	var p util.Map
	requester := BuildRequester(req)
	resp := NotifyTypeResponder(requester.Type(), NotifySuccess())
	defer func() {
		e = resp.Write(w)
		log.Error(e)
	}()

	if e = requester.Error(); e != nil {
		log.Error(e.Error())
		resp.SetNotifyResult(NotifyFail(e.Error()))
		return
	}
	reqData := requester.ToMap()
	if util.ValidateSign(reqData, obj.payment.GetKey()) {
		if obj.ServeNotify == nil {
			log.Error(xerrors.New("null notify callback"))
			return
		}
		p, e = obj.ServeNotify(requester)
		if e != nil {
			log.Error(e.Error())
			resp.SetNotifyResult(NotifyFailDes(resp.NotifyResult(), e.Error()))
		}
		if !p.Has("prepay_id") {
			log.Error("null prepay_id")
			resp.SetNotifyResult(NotifyFailDes(resp.NotifyResult(), "null prepay_id"))
		} else {
			//公众账号ID	appid	String(32)	是	wx8888888888888888	微信分配的公众账号ID
			//商户号	mch_id	String(32)	是	1900000109	微信支付分配的商户号
			//随机字符串	nonce_str	String(32)	是	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	微信返回的随机字符串
			//预支付ID	prepay_id	String(64)	是	wx201410272009395522657a690389285100	调用统一下单接口生成的预支付ID
			//业务结果	result_code	String(16)	是	SUCCESS	SUCCESS/FAIL
			//错误描述	err_code_des	String(128)	否		当result_code为FAIL时，商户展示给用户的错误提
			//签名	sign	String(32)	是	C380BEC2BFD727A4B6845133519F3AD6	返回数据签名，签名生成算法
			res := resp.NotifyResult()
			res.AppID = obj.payment.AppID
			res.MchID = obj.payment.MchID
			res.NonceStr = util.GenerateNonceStr()
			res.PrepayID = p.GetString("prepay_id")
			res.Sign = util.GenSign(reqData, obj.payment.GetKey())
		}

	}

}

/*Notifier 监听 */
type paymentPaid struct {
	*Notify
	ServeNotify
}

// NotifyResponder ...
type NotifyResponder interface {
	SetNotifyResult(result *NotifyResult)
	NotifyResult() *NotifyResult
	Write(w http.ResponseWriter) error
}

type xmlNotify struct {
	notifyResult *NotifyResult
}

// NotifyResult ...
func (obj *xmlNotify) NotifyResult() *NotifyResult {
	return obj.notifyResult
}

// SetNotifyResult ...
func (obj *xmlNotify) SetNotifyResult(notifyResult *NotifyResult) {
	obj.notifyResult = notifyResult
}

// Write ...
func (obj *xmlNotify) Write(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusOK)
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{"application/xml; charset=utf-8"}
	}
	if obj.notifyResult == nil {
		return xerrors.New("null notify result")
	}
	_, err := w.Write(obj.notifyResult.ToXML())
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

type jsonNotify struct {
	notifyResult *NotifyResult
}

// NotifyResult ...
func (obj *jsonNotify) NotifyResult() *NotifyResult {
	return obj.notifyResult
}

// SetNotifyResult ...
func (obj *jsonNotify) SetNotifyResult(notifyResult *NotifyResult) {
	obj.notifyResult = notifyResult
}

// Write ...
func (obj *jsonNotify) Write(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusOK)
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{"application/json; charset=utf-8"}
	}
	if obj.notifyResult == nil {
		return xerrors.New("null notify result")
	}
	_, err := w.Write(obj.notifyResult.ToJSON())
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// NotifyTypeResponder ...
func NotifyTypeResponder(bodyType BodyType, notifyResult *NotifyResult) NotifyResponder {
	switch bodyType {
	case BodyTypeJSON:
		return &jsonNotify{
			notifyResult: notifyResult,
		}
	case BodyTypeXML:
		return &xmlNotify{
			notifyResult: notifyResult,
		}
	}
	return nil
}

// ToJSON ...
func (obj *NotifyResult) ToJSON() []byte {
	bytes, e := jsoniter.Marshal(obj)
	if e != nil {
		log.Error(e)
		return nil
	}
	return bytes
}

// ToXML ...
func (obj *NotifyResult) ToXML() []byte {
	bytes, e := xml.Marshal(obj)
	if e != nil {
		log.Error(e)
		return nil
	}
	return bytes
}

// NotifySuccess ...
func NotifySuccess() *NotifyResult {
	return &NotifyResult{
		ReturnCode: "SUCCESS",
		ReturnMsg:  "OK",
	}
}

// NotifyFail ...
func NotifyFail(msg string) *NotifyResult {
	return &NotifyResult{
		ReturnCode: "FAIL",
		ReturnMsg:  msg,
	}
}

// NotifyFailDes ...
func NotifyFailDes(r *NotifyResult, msg string) *NotifyResult {
	r.ResultCode = "FAIL"
	r.ErrCodeDes = msg
	return r
}
