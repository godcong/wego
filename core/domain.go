package core

import (
	"strings"
)

//Link link domain address
func Link(uri string, host ...string) string {
	domain := "default"
	if host != nil {
		domain = host[0]
	}
	domainAddr := DomainHost(domain)
	return Splice(domainAddr, uri)
}

/*Splice 拼接地址 */
func Splice(domain string, uri string) string {
	switch {
	case strings.Index(uri, "/") == 0 && strings.LastIndex(domain, "/") == (len(domain)-1):
		domain = domain[:len(domain)-1]
		//uri = uri[1:]
	case strings.Index(uri, "/") == 0 && strings.LastIndex(domain, "/") != (len(domain)-1):
		//uri = uri[1:]
	case strings.Index(uri, "/") != 0 && strings.LastIndex(domain, "/") == (len(domain)-1):
		domain = domain[:len(domain)-1]
		uri = "/" + uri
	case strings.Index(uri, "/") != 0 && strings.LastIndex(domain, "/") != (len(domain)-1):
		uri = "/" + uri
	}
	return domain + uri
}

//DomainHost get host domain
func DomainHost(domain string) string {
	url := DefaultConfig().GetString("domain." + domain + ".url")
	if url == "" {
		switch domain {
		case "host":
			url = "http://localhost"
		case "payment", "default":
			url = BaseDomain
		case "official_account", "mini_program":
			url = APIWeixin
		case "file":
			url = FileAPIWeixin
		case "mp":
			url = MPDomain
		case "api2":
			url = API2Domain
		default:
			url = BaseDomain
		}
	}
	return url
}
