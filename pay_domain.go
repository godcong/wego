package wxpay

import (
	"errors"
	"strconv"
	"sync"

	"github.com/godcong/wopay/util"
)

var (
	ErrprConnectTimeout = errors.New("ErrorConnectTimeout")
	ErrorUnknownHost    = errors.New("ErrorUnknownHostException")
)

type PayDomain interface {
	Report(string, int64, error)
	GetDomainInfo() *DomainInfo
}

type DomainInfo struct {
	Domain        string //域名
	PrimaryDomain bool   //该域名是否为主域名。例如:api.mch.weixin.qq.com为主域名

}

func NewDomainInfo(domain string, primary bool) *DomainInfo {
	return &DomainInfo{
		Domain:        domain,
		PrimaryDomain: primary,
	}
}

func (info *DomainInfo) String() string {
	return "DomainInfo{" + "domain='" + info.Domain + "'" + ", primaryDomain=" + strconv.FormatBool(info.PrimaryDomain) + "}"
}

//func (info *DomainInfo) Report(string, int64, error) {
//
//}
//func (info *DomainInfo) GetDomain(PayConfig) DomainInfo {
//	return *info
//}
var holder PayDomain

func init() {
	holder = PayDomainSimpleInstance()
}

type payDomainSimpleImpl struct {
	sync.Mutex
	domainData map[string]*DomainStatics
	domainTime int64
}

const MIN_SWITCH_PRIMARY_MSEC = 3 * 60 * 1000

var switchToAlternateDomainTime = 0

//var domainData = map[string]DomainStatics{}

type DomainStatics struct {
	domain              string
	succCount           int
	connectTimeoutCount int
	dnsErrorCount       int
	otherErrorCount     int
}

func NewDomainStatics(domain string) *DomainStatics {
	return &DomainStatics{
		domain:              domain,
		succCount:           0,
		connectTimeoutCount: 0,
		dnsErrorCount:       0,
		otherErrorCount:     0,
	}
}

func (s *DomainStatics) resetCount() {
	s.succCount = 0
	s.connectTimeoutCount = 0
	s.dnsErrorCount = 0
	s.otherErrorCount = 0
}

func (s *DomainStatics) isGood() bool {
	return s.connectTimeoutCount <= 2 && s.dnsErrorCount <= 2
}

func (s *DomainStatics) badCount() int {
	return s.connectTimeoutCount + s.dnsErrorCount*5 + s.otherErrorCount/4
}

func PayDomainSimpleInstance() PayDomain {
	if holder == nil {
		holder = NewPayDomainSimple()
	}
	return holder
}

func NewPayDomainSimple() *payDomainSimpleImpl {
	return &payDomainSimpleImpl{
		domainData: make(map[string]*DomainStatics),
	}
}

func (domain *payDomainSimpleImpl) Report(d string, elapsed int64, err error) {
	domain.Lock()
	defer domain.Unlock()
	info, b := domain.domainData[d]
	if !b {
		info = NewDomainStatics(d)
		domain.domainData[d] = info
	}

	if err == nil { //success
		if info.succCount >= 2 { //continue succ, clear error count
			info.connectTimeoutCount = 0
			info.dnsErrorCount = 0
			info.otherErrorCount = 0
		} else {
			info.succCount++
		}
	} else if err == ErrprConnectTimeout {
		info.succCount = 0
		info.dnsErrorCount = 0
		info.connectTimeoutCount++
	} else if err == ErrorUnknownHost {
		info.succCount = 0
		info.dnsErrorCount++
	} else {
		info.succCount = 0
		info.otherErrorCount++
	}

}
func (domain *payDomainSimpleImpl) GetDomainInfo() *DomainInfo {
	domain.Lock()
	defer domain.Unlock()
	if domain == nil {
		return nil
	}
	primaryDomain, b := domain.domainData[DOMAIN_API]
	if !b ||
		primaryDomain.isGood() {
		return NewDomainInfo(DOMAIN_API, true)
	}

	now := util.CurrentTimeStampMS()
	if domain.domainTime == 0 { //first switch
		domain.domainTime = now
		return NewDomainInfo(DOMAIN_API2, false)
	} else if now-domain.domainTime < MIN_SWITCH_PRIMARY_MSEC {
		alternateDomain, b := domain.domainData[DOMAIN_API2]
		if !b ||
			alternateDomain.isGood() ||
			alternateDomain.badCount() < primaryDomain.badCount() {
			return NewDomainInfo(DOMAIN_API2, false)
		} else {
			return NewDomainInfo(DOMAIN_API, true)
		}
	} else { //force switch back
		domain.domainTime = 0
		primaryDomain.resetCount()
		alternateDomain, b := domain.domainData[DOMAIN_API2]
		if !b {
			alternateDomain.resetCount()
		}

		return NewDomainInfo(DOMAIN_API, true)
	}

}
