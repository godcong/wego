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
type NotifyFunc func(p util.Map) (util.Map, error)

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
		rlt = Fail(err.Error())
	} else {
		//TODO: decode req_info
		_, err = n.NotifyFunc(maps)
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
	var err error
	rlt := SUCCESS()
	var p util.Map
	maps, err := core.RequestToMap(req)

	if err != nil {
		log.Error(err)
		rlt = FailDes(err.Error())
	} else {
		if ValidateSign(maps, n.GetKey()) {
			p, err = n.NotifyFunc(maps)
			if err != nil {
				rlt = FailDes(err.Error())
			}
			if !p.Has("prepay_id") {
				log.Error("nil prepay_id")
				rlt = FailDes("nil prepay_id")
			} else {
				//公众账号ID	appid	String(32)	是	wx8888888888888888	微信分配的公众账号ID
				//商户号	mch_id	String(32)	是	1900000109	微信支付分配的商户号
				//随机字符串	nonce_str	String(32)	是	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	微信返回的随机字符串
				//预支付ID	prepay_id	String(64)	是	wx201410272009395522657a690389285100	调用统一下单接口生成的预支付ID
				//业务结果	result_code	String(16)	是	SUCCESS	SUCCESS/FAIL
				//错误描述	err_code_des	String(128)	否		当result_code为FAIL时，商户展示给用户的错误提
				//签名	sign	String(32)	是	C380BEC2BFD727A4B6845133519F3AD6	返回数据签名，签名生成算法
				rlt.Set("appid", n.Get("app_id"))
				rlt.Set("mch_id", n.Get("mch_id"))
				rlt.Set("nonce_str", util.GenerateNonceStr())
				rlt.Set("prepay_id", p.Get("prepay_id"))
				switch maps.GetString("sign_type") {
				case HMACSHA256:
					rlt.Set("sign", GenerateSignature(maps, n.GetKey(), MakeSignHMACSHA256))
				default:
					rlt.Set("sign", GenerateSignature(maps, n.GetKey(), MakeSignMD5))
				}

			}

		}
	}

	err = NotifyResponseXML(w, rlt.ToXML())
	if err != nil {
		log.Error(err)
	}
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

	if err != nil {
		log.Error(err)
		rlt = Fail(err.Error())
	} else {
		if ValidateSign(maps, n.GetKey()) {
			_, err = n.NotifyFunc(maps)
			if err != nil {
				rlt = Fail(err.Error())
			}
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

// FailDes ...
func FailDes(msg string) util.Map {
	return util.Map{
		"return_code":  "FAIL",
		"err_code_des": msg,
	}
}
