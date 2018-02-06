package wxpay

import (
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/godcong/wopay/util"
)

const DEFAULT_CONNECT_TIMEOUT_MS = 6 * 1000
const DEFAULT_READ_TIMEOUT_MS = 8 * 1000
const REPORT_URL = "http://report.mch.weixin.qq.com/wxpay/report/default"

type PayReport struct {
	reportMsgQueue sync.Map
	config         PayConfig
	//private ExecutorService executorService;

}

var payReport *PayReport

type ReportInfo struct {
	// 基本信息
	Version           string
	Sdk               string
	Uuid              string // 交易的标识
	Timestamp         int64  // 上报时的时间戳，单位秒
	ElapsedTimeMillis int64  // 耗时，单位 毫秒
	// 针对主域名
	FirstDomain               string // 第1次请求的域名
	PrimaryDomain             bool   //是否主域名
	FirstConnectTimeoutMillis int    // 第1次请求设置的连接超时时间，单位 毫秒
	FirstReadTimeoutMillis    int    // 第1次请求设置的读写超时时间，单位 毫秒
	FirstHasDnsError          int    // 第1次请求是否出现dns问题
	FirstHasConnectTimeout    int    // 第1次请求是否出现连接超时
	FirstHasReadTimeout       int    // 第1次请求是否出现连接超时
}

func newPayReport(config PayConfig) *PayReport {
	return &PayReport{
		config: config,
	}
}

func PayReportInstance(config PayConfig) *PayReport {
	if payReport == nil {
		payReport = newPayReport(config)
	}
	return payReport
}

func (report *PayReport) Report(uuid string, elapsedTimeMillis int64,
	firstDomain string, primaryDomain bool, firstConnectTimeoutMillis,
	firstReadTimeoutMillis int, firstHasDnsError, firstHasConnectTimeout,
	firstHasReadTimeout bool) {
	currentTimestamp := util.CurrentTimeStamp()
	reportInfo := NewReportInfo(uuid,
		currentTimestamp,
		elapsedTimeMillis,
		firstDomain,
		primaryDomain,
		firstConnectTimeoutMillis,
		firstReadTimeoutMillis,
		firstHasDnsError,
		firstHasConnectTimeout,
		firstHasReadTimeout)
	data := reportInfo.ToLineString(report.config.Key())
	log.Println("report {", data, "}")
}

func NewReportInfo(
	uuid string,
	timestamp int64,
	elapsedTimeMillis int64,
	firstDomain string,
	primaryDomain bool,
	firstConnectTimeoutMillis, firstReadTimeoutMillis int,
	firstHasDnsError, firstHasConnectTimeout, firstHasReadTimeout bool) *ReportInfo {
	return &ReportInfo{
		Version:                   "v0",
		Sdk:                       "wxpay go sdk v1.0",
		Uuid:                      uuid,
		Timestamp:                 timestamp,
		ElapsedTimeMillis:         elapsedTimeMillis,
		FirstDomain:               firstDomain,
		PrimaryDomain:             primaryDomain,
		FirstConnectTimeoutMillis: firstConnectTimeoutMillis,
		FirstReadTimeoutMillis:    firstReadTimeoutMillis,
		FirstHasDnsError:          ParseInt(firstHasDnsError),
		FirstHasConnectTimeout:    ParseInt(firstHasConnectTimeout),
		FirstHasReadTimeout:       ParseInt(firstHasReadTimeout),
	}
}

func (r *ReportInfo) ToString() string {
	return "ReportInfo{" +
		"version='" + r.Version + "'" +
		", sdk='" + r.Sdk + "'" +
		", uuid='" + r.Uuid + "'" +
		", timestamp=" + strconv.FormatInt(r.Timestamp, 10) +
		", elapsedTimeMillis=" + strconv.FormatInt(r.ElapsedTimeMillis, 10) +
		", firstDomain='" + r.FirstDomain + "'" +
		", primaryDomain=" + strconv.FormatBool(r.PrimaryDomain) +
		", firstConnectTimeoutMillis=" + strconv.FormatInt(int64(r.FirstConnectTimeoutMillis), 10) +
		", firstReadTimeoutMillis=" + strconv.FormatInt(int64(r.FirstReadTimeoutMillis), 10) +
		", firstHasDnsError=" + strconv.FormatInt(int64(r.FirstHasDnsError), 10) +
		", firstHasConnectTimeout=" + strconv.FormatInt(int64(r.FirstHasConnectTimeout), 10) +
		", firstHasReadTimeout=" + strconv.FormatInt(int64(r.FirstHasReadTimeout), 10) +
		"}"
}

func (r *ReportInfo) ToLineString(key string) string {
	obj := []string{
		r.Version,
		r.Sdk,
		r.Uuid,
		strconv.FormatInt(r.Timestamp, 10),
		strconv.FormatInt(r.ElapsedTimeMillis, 10),
		r.FirstDomain,
		strconv.FormatBool(r.PrimaryDomain),
		strconv.FormatInt(int64(r.FirstConnectTimeoutMillis), 10),
		strconv.FormatInt(int64(r.FirstReadTimeoutMillis), 10),
		strconv.FormatInt(int64(r.FirstHasDnsError), 10),
		strconv.FormatInt(int64(r.FirstHasConnectTimeout), 10),
		strconv.FormatInt(int64(r.FirstHasReadTimeout), 10),
	}
	s := strings.Join(obj, ",") + ","
	return s + util.MakeSignHMACSHA256(s, key)
}

func ParseInt(b bool) (i int) {
	if i = 0; b {
		i = 1
	}
	return
}
