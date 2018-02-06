package wxpay

import (
	"io/ioutil"
)

type PayConfigImpl struct {
	connectTimeoutMs   int
	readTimeoutMs      int
	autoReport         bool
	reportWorkerNum    int
	reportQueueMaxSize int
	reportBatchSize    int
	payDomain          PayDomain
	cert               []byte
}

type PayConfig interface {
	AppID() string
	MchID() string
	Key() string
	Cert() []byte
	ConnectTimeoutMs() int
	ReadTimeoutMs() int
	PayDomainInstance() PayDomain
	AutoReport() bool
	ReportWorkNum() int
	ReportQueueMaxSize() int
	ReportBatchSize() int
}

var config PayConfig

func init() {
	PayConfigInstance()
}

func PayConfigInstance() PayConfig {
	if config == nil {
		config = NewPayConfig()
	}
	return config
}

func SetPayConfig(conf PayConfig) PayConfig {
	config = conf
	return config
}

func NewPayConfig() PayConfig {
	return &PayConfigImpl{
		connectTimeoutMs:   6000,
		readTimeoutMs:      8000,
		autoReport:         true,
		reportWorkerNum:    6,
		reportQueueMaxSize: 10000,
		reportBatchSize:    10,
	}
}

func (impl *PayConfigImpl) AppID() string {
	return "wxa8550ae44b713f1e"
}
func (impl *PayConfigImpl) MchID() string {
	return "1300910101"
}
func (impl *PayConfigImpl) Key() string {
	return "MvIJ6ZlAIcrgY7OGI5Z9PIU3RcVPfKZs"
}
func (impl *PayConfigImpl) Cert() []byte {
	cert, _ := ioutil.ReadFile(`\apiclient_cert.p12`)
	return cert
}
func (impl *PayConfigImpl) ConnectTimeoutMs() int {
	return impl.connectTimeoutMs
}
func (impl *PayConfigImpl) ReadTimeoutMs() int {
	return impl.readTimeoutMs
}
func (impl *PayConfigImpl) PayDomainInstance() PayDomain {
	return PayDomainSimpleInstance()
}
func (impl *PayConfigImpl) AutoReport() bool {
	return impl.autoReport
}
func (impl *PayConfigImpl) ReportWorkNum() int {
	return impl.reportWorkerNum
}
func (impl *PayConfigImpl) ReportQueueMaxSize() int {
	return impl.reportQueueMaxSize
}
func (impl *PayConfigImpl) ReportBatchSize() int {
	return impl.reportBatchSize
}
