package wxpay

import (
	"crypto/tls"

	"net/http"

	"bytes"

	"crypto/x509"
	"errors"
	"io/ioutil"
	"log"
	"time"

	"github.com/godcong/wopay/util"
)

type PayRequest struct {
	config PayConfig
}

var (
	ErrorNilDomain       = errors.New("PayConfig.PayDomain().getDomain() is empty or null")
	ErrorLoadX509KeyPair = errors.New("LoadX509KeyPair() is empty to load")
	ErrorReadRootCAFile  = errors.New("read rootca.pem file error")
)

func NewPayRequest(config PayConfig) *PayRequest {
	return &PayRequest{config: config}
}

/**
 * 请求，只请求一次，不做重试
 * @param domain
 * @param urlSuffix
 * @param uuid
 * @param data
 * @param connectTimeoutMs
 * @param readTimeoutMs
 * @param useCert 是否使用证书，针对退款、撤销等操作
 * @return
 * @throws Exception
 */
func (request *PayRequest) RequestOnce(domain, urlSuffix, uuid, data string, connectTimeoutMs, readTimeoutMs int, useCert bool) (string, error) {
	return request.requestOnce(domain, urlSuffix, uuid, data, connectTimeoutMs, readTimeoutMs, useCert)
}

func (request *PayRequest) requestOnce(domain, urlSuffix, uuid, data string, connectTimeoutMs, readTimeoutMs int, useCert bool) (string, error) {
	var tr *http.Transport

	if useCert {
		cert, err := tls.LoadX509KeyPair(SSLCERT_PATH, SSLKEY_PATH)
		if err != nil {
			return "", ErrorLoadX509KeyPair
		}

		caCert, err := ioutil.ReadFile("./cert/rootca.pem")
		if err != nil {
			return "", ErrorReadRootCAFile
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)
		tlsConfig := &tls.Config{
			Certificates:       []tls.Certificate{cert},
			RootCAs:            caCertPool,
			InsecureSkipVerify: false,
		}
		tlsConfig.BuildNameToCertificate()
		tr = &http.Transport{
			//Dial: (&net.Dialer{
			//	Timeout:   30 * time.Second,
			//	KeepAlive: 30 * time.Second,
			//}).Dial,
			TLSClientConfig: tlsConfig,
			//TLSHandshakeTimeout:   10 * time.Second,
			//ResponseHeaderTimeout: 10 * time.Second,
			//ExpectContinueTimeout: 1 * time.Second,
		}
	} else {
		tr = &http.Transport{
			//Dial: (&net.Dialer{
			//	Timeout:   30 * time.Second,
			//	KeepAlive: 30 * time.Second,
			//}).Dial,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			//TLSHandshakeTimeout:   10 * time.Second,
			//ResponseHeaderTimeout: 10 * time.Second,
			//ExpectContinueTimeout: 1 * time.Second,
		}
	}
	url := "https://" + domain + urlSuffix

	client := &http.Client{
		Transport: tr,
		Timeout:   time.Duration((connectTimeoutMs + readTimeoutMs) * 1000000),
	}
	log.Println(data)
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "text/xml")
	req.Header.Set("User-Agent", "wxpay sdk go v1.0 "+request.config.MchID())

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}

func (request *PayRequest) RequestWithCert(urlSuffix, uuid, data string, autoReport bool) (string, error) {
	return request.request(urlSuffix, uuid, data, request.config.ConnectTimeoutMs(), request.config.ReadTimeoutMs(), true, autoReport)
}

func (request *PayRequest) RequestWithCertTimeout(urlSuffix, uuid, data string, connectTimeoutMs, readTimeoutMs int, autoReport bool) (string, error) {
	return request.request(urlSuffix, uuid, data, connectTimeoutMs, readTimeoutMs, true, autoReport)
}

func (request *PayRequest) RequestWithoutCert(urlSuffix, uuid, data string, autoReport bool) (string, error) {
	return request.request(urlSuffix, uuid, data, request.config.ConnectTimeoutMs(), request.config.ReadTimeoutMs(), false, autoReport)
}

func (request *PayRequest) RequestWithoutCertTimeout(urlSuffix, uuid, data string, connectTimeoutMs, readTimeoutMs int, autoReport bool) (string, error) {
	return request.request(urlSuffix, uuid, data, connectTimeoutMs, readTimeoutMs, false, autoReport)
}

func (request *PayRequest) request(urlSuffix, uuid, data string, connectTimeoutMs, readTimeoutMs int, useCert, autoReport bool) (string, error) {
	startTimestampMs := util.CurrentTimeStampMS()
	firstHasDnsErr, firstHasConnectTimeout, firstHasReadTimeout := false, false, false
	domainInfo := request.config.PayDomainInstance().GetDomainInfo()
	if domainInfo == nil {
		return "", ErrorNilDomain
	}
	result, err := request.requestOnce(domainInfo.Domain, urlSuffix, uuid, data, connectTimeoutMs, readTimeoutMs, useCert)
	elapsedTimeMillis := util.CurrentTimeStampMS() - startTimestampMs
	request.config.PayDomainInstance().Report(domainInfo.Domain, elapsedTimeMillis, nil)

	PayReportInstance(request.config).Report(uuid,
		elapsedTimeMillis,
		domainInfo.Domain,
		domainInfo.PrimaryDomain,
		connectTimeoutMs,
		readTimeoutMs,
		firstHasDnsErr,
		firstHasConnectTimeout,
		firstHasReadTimeout)

	return result, err
}
