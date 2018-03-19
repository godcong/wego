package core

import "strings"

type Domain struct {
	url string
	app *Application
}

func (d *Domain) URL() string {
	if d.url == "" {
		return BASE_DOMAIN
	}
	return d.url
}

func (d *Domain) Link(s string) string {
	url := ""
	switch {
	case strings.Index(s, "/") == 0 && strings.LastIndex(d.URL(), "/") == (len(d.URL())-1):
		url = d.URL() + s[1:]
	case strings.Index(s, "/") == 0 && strings.LastIndex(d.URL(), "/") != (len(d.URL())-1):
		url = d.URL() + s
	case strings.Index(s, "/") != 0 && strings.LastIndex(d.URL(), "/") == (len(d.URL())-1):
		url = d.URL() + s
	case strings.Index(s, "/") != 0 && strings.LastIndex(d.URL(), "/") != (len(d.URL())-1):
		url = d.URL() + "/" + s
	}
	Debug("Domain|Link", url)
	return url
}

//func NewDomain(application Application) Domain {
//	return &Domain{
//		Config: application.Config(),
//		app:    application,
//	}
//}

func newDomain(s string) *Domain {
	url := GetConfig(DeployJoin("domain", s)).Get("url")
	if url == "" {
		switch s {
		case "host":
			url = "localhost"
		case "payment":
			fallthrough
		case "default":
			url = BASE_DOMAIN
		case "official_account":
			url = API_WEIXIN
		case "file":
			url = FILE_API_WEIXIN
		default:
			url = BACK_DOMAIN
		}
	}
	return &Domain{
		url: url,
	}
}

func NewDomain(prefix string) *Domain {
	return newDomain(prefix)
}

func DomainHost() *Domain {
	return newDomain("host")
}

//
//type DomainInfo struct {
//	Domain        string //域名
//	PrimaryDomain bool   //该域名是否为主域名。例如:api.mch.weixin.qq.com为主域名
//
//}
//
//func NewDomainInfo(Domain string, primary bool) *DomainInfo {
//	return &DomainInfo{
//		Domain:        Domain,
//		PrimaryDomain: primary,
//	}
//}
//
//func (info *DomainInfo) String() string {
//	return "DomainInfo{" + "Domain='" + info.Domain + "'" + ", primaryDomain=" + strconv.FormatBool(info.PrimaryDomain) + "}"
//}
//
////func (info *DomainInfo) Report(string, int64, error) {
////
////}
////func (info *DomainInfo) GetDomain(PayConfig) DomainInfo {
////	return *info
////}
//var holder PayDomain
//
//func init() {
//	holder = PayDomainSimpleInstance()
//}
//
//type payDomainSimpleImpl struct {
//	sync.Mutex
//	domainData map[string]*DomainStatics
//	domainTime int64
//}
//
//const MIN_SWITCH_PRIMARY_MSEC = 3 * 60 * 1000
//
//var switchToAlternateDomainTime = 0
//
////var domainData = map[string]DomainStatics{}
//
//type DomainStatics struct {
//	Domain              string
//	succCount           int
//	connectTimeoutCount int
//	dnsErrorCount       int
//	otherErrorCount     int
//}
//
//func NewDomainStatics(Domain string) *DomainStatics {
//	return &DomainStatics{
//		Domain:              Domain,
//		succCount:           0,
//		connectTimeoutCount: 0,
//		dnsErrorCount:       0,
//		otherErrorCount:     0,
//	}
//}
//
//func (s *DomainStatics) resetCount() {
//	s.succCount = 0
//	s.connectTimeoutCount = 0
//	s.dnsErrorCount = 0
//	s.otherErrorCount = 0
//}
//
//func (s *DomainStatics) isGood() bool {
//	return s.connectTimeoutCount <= 2 && s.dnsErrorCount <= 2
//}
//
//func (s *DomainStatics) badCount() int {
//	return s.connectTimeoutCount + s.dnsErrorCount*5 + s.otherErrorCount/4
//}
//
//func PayDomainSimpleInstance() PayDomain {
//	if holder == nil {
//		holder = NewPayDomainSimple()
//	}
//	return holder
//}
//
//func NewPayDomainSimple() *payDomainSimpleImpl {
//	return &payDomainSimpleImpl{
//		domainData: make(map[string]*DomainStatics),
//	}
//}
//
//func (Domain *payDomainSimpleImpl) Report(d string, elapsed int64, err error) {
//	Domain.Lock()
//	defer Domain.Unlock()
//	info, b := Domain.domainData[d]
//	if !b {
//		info = NewDomainStatics(d)
//		Domain.domainData[d] = info
//	}
//
//	if err == nil { //success
//		if info.succCount >= 2 { //continue succ, clear error count
//			info.connectTimeoutCount = 0
//			info.dnsErrorCount = 0
//			info.otherErrorCount = 0
//		} else {
//			info.succCount++
//		}
//	} else if err == ErrprConnectTimeout {
//		info.succCount = 0
//		info.dnsErrorCount = 0
//		info.connectTimeoutCount++
//	} else if err == ErrorUnknownHost {
//		info.succCount = 0
//		info.dnsErrorCount++
//	} else {
//		info.succCount = 0
//		info.otherErrorCount++
//	}
//
//}
//func (Domain *payDomainSimpleImpl) GetDomainInfo() *DomainInfo {
//	Domain.Lock()
//	defer Domain.Unlock()
//	if Domain == nil {
//		return nil
//	}
//	primaryDomain, b := Domain.domainData[DOMAIN_API]
//	if !b ||
//		primaryDomain.isGood() {
//		return NewDomainInfo(DOMAIN_API, true)
//	}
//
//	now := util.CurrentTimeStampMS()
//	if Domain.domainTime == 0 { //first switch
//		Domain.domainTime = now
//		return NewDomainInfo(DOMAIN_API2, false)
//	} else if now-Domain.domainTime < MIN_SWITCH_PRIMARY_MSEC {
//		alternateDomain, b := Domain.domainData[DOMAIN_API2]
//		if !b ||
//			alternateDomain.isGood() ||
//			alternateDomain.badCount() < primaryDomain.badCount() {
//			return NewDomainInfo(DOMAIN_API2, false)
//		} else {
//			return NewDomainInfo(DOMAIN_API, true)
//		}
//	} else { //force switch back
//		Domain.domainTime = 0
//		primaryDomain.resetCount()
//		alternateDomain, b := Domain.domainData[DOMAIN_API2]
//		if !b {
//			alternateDomain.resetCount()
//		}
//
//		return NewDomainInfo(DOMAIN_API, true)
//	}
//
//}
