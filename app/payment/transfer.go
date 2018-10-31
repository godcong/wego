package payment

import (
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/crypt"
	"github.com/godcong/wego/log"

	"github.com/godcong/wego/util"
)

/*Transfer Transfer */
type Transfer struct {
	*Payment
}

func newTransfer(pay *Payment) *Transfer {
	return &Transfer{
		Payment: pay,
	}
}

/*NewTransfer NewTransfer */
func NewTransfer(config *core.Config) *Transfer {
	return newTransfer(NewPayment(config))
}

/*QueryBalanceOrder 查询企业付款
接口说明
请求Url	https://api.mch.weixin.qq.com/mmpaymkttransfers/gettransferinfo
是否需要证书	请求需要双向证书。 详见证书使用
请求方式	POST
请求参数
字段名	字段	必填	示例值	类型	说明
随机字符串	nonce_str	是	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	String(32)	随机字符串，不长于32位
签名	sign	是	C380BEC2BFD727A4B6845133519F3AD6	String(32)	生成签名方式查看3.2.1节
商户订单号	partner_trade_no	是	10000098201411111234567890	String(28)	商户调用企业付款API时使用的商户订单号
商户号	mch_id	是	10000098	String(32)	微信支付分配的商户号
Appid	appid	是	wxe062425f740d30d8	String(32)	商户号的appid
*/
func (t *Transfer) QueryBalanceOrder(s string) core.Response {
	m := util.Map{
		"appid":            t.Get("app_id"),
		"mch_id":           t.Get("mch_id"),
		"partner_trade_no": s,
	}
	return t.SafeRequest(getTransferInfoURLSuffix, m)
}

/*ToBalance 企业付款
接口地址
接口链接：https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers
是否需要证书
请求需要双向证书。 详见证书使用
请求参数
字段名	变量名	必填	示例值	类型	描述
商户账号appid	mch_appid	是	wx8888888888888888	String	申请商户号的appid或商户号绑定的appid
商户号	mchid	是	1900000109	String(32)	微信支付分配的商户号
设备号	device_info	否	013467007045764	String(32)	微信支付分配的终端设备号
随机字符串	nonce_str	是	5K8264ILTKCH16CQ2502SI8ZNMTM67VS	String(32)	随机字符串，不长于32位
签名	sign	是	C380BEC2BFD727A4B6845133519F3AD6	String(32)	签名，详见签名算法
商户订单号	partner_trade_no	是	10000098201411111234567890	String	商户订单号，需保持唯一性
(只能是字母或者数字，不能包含有符号)
用户openid	openid	是	oxTWIuGaIt6gTKsQRLau2M0yL16E	String	商户appid下，某用户的openid
校验用户姓名选项	check_name	是	FORCE_CHECK	String	NO_CHECK：不校验真实姓名
FORCE_CHECK：强校验真实姓名
收款用户姓名	re_user_name	可选	王小王	String	收款用户真实姓名。
如果check_name设置为FORCE_CHECK，则必填用户真实姓名
金额	amount	是	10099	int	企业付款金额，单位为分
企业付款描述信息	desc	是	理赔	String	企业付款操作说明信息。必填。
Ip地址	spbill_create_ip	是	192.168.0.1	String(32)	该IP同在商户平台设置的IP白名单中的IP没有关联，该IP可传用户端或者服务端的IP。
*/
func (t *Transfer) ToBalance(m util.Map) core.Response {
	m.Delete("mch_id")
	m.Set("mchid", t.Get("mch_id"))
	m.Set("mch_appid", t.Get("app_id"))

	if !m.Has("spbill_create_ip") {
		m.Set("spbill_create_ip", core.GetServerIP())
	}
	return t.SafeRequest(promotionTransfersURLSuffix, m)
}

/*QueryBankCardOrder 查询企业付款银行卡API
接口说明
请求Url	https://api.mch.weixin.qq.com/mmpaysptrans/query_bank
是否需要证书	请求需要双向证书。 详见证书使用
请求方式	POST
请求参数
字段名	字段	必填	示例值	类型	说明
商户号	mch_id	是	1900000109	string(32)	微信支付分配的商户号
商户企业付款单号	partner_trade_no	是	1212121221227	string(32)	商户订单号，需保持唯一（只允许数字[0~9]或字母[A~Z]和[a~z]最短8位，最长32位）
随机字符串	nonce_str	是	5K8264ILTKCH16CQ2502SI8ZNMTM67Vs	string(32)	随机字符串，长度小于32位
签名	sign	是	C380BEC2BFD727A4B6845133519F3AD6	string(32)	商户自带签名
返回参数
字段名	变量名	必填	类型	说明
返回状态码	return_code	是	String(16)
SUCCESS/FAIL
此字段是通信标识，非付款标识，付款是否成功需要查看result_code来判断
返回信息	return_msg	否	String(128)
返回信息，如非空，为错误原因
签名失败
参数格式校验错误
以下字段在return_code为SUCCESS的时候有返回
业务结果	result_code	是	String(16)	SUCCESS/FAIL，非付款标识，付款是否成功需要查看status字段来判断
错误代码	err_code	否	String(32)	错误码信息
错误代码描述	err_code_des	否	String(128)	结果信息描述
以下字段在return_code 和result_code都为SUCCESS的时候有返回
商户号	mch_id	是	string(32)	商户号
商户企业付款单号	partner_trade_no	是	string(32)	商户单号
微信企业付款单号	payment_no	是	string(64)	即为微信内部业务单号
银行卡号	bank_no_md5	是	string(32)	收款用户银行卡号(MD5加密)
用户真实姓名	true_name_md5	是	string(32)	收款人真实姓名（MD5加密）
代付金额	amount	是	int	代付订单金额RMB：分
代付单状态	status	是	string	代付订单状态：
PROCESSING（处理中，如有明确失败，则返回额外失败原因；否则没有错误原因）
SUCCESS（付款成功）
FAILED（付款失败,需要替换付款单号重新发起付款）
BANK_FAIL（银行退票，订单状态由付款成功流转至退票,退票时付款金额和手续费会自动退还）
手续费金额	cmms_amt	是	int	手续费订单金额 RMB：分
商户下单时间	create_time	是	String	微信侧订单创建时间
成功付款时间	pay_succ_time	否	String	微信侧付款成功时间（但无法保证银行不会退票）
失败原因	reason	否	String	订单失败原因（如：余额不足）
*/
func (t *Transfer) QueryBankCardOrder(s string) core.Response {
	m := util.Map{
		"mch_id":           t.Get("mch_id"),
		"partner_trade_no": s,
	}
	return t.SafeRequest(mmPaySpTransQueryBankURLSuffix, m)
}

/*ToBankCard 转账至银行卡
接口介绍
业务流程	接口	简介
付款	企业付款到银行卡	用于企业向微信用户银行卡付款
目前支持接口API的方式向指定微信用户的银行卡付款。
接口调用规则：
◆ 单商户日限额——单日100w
◆ 单次限额——单次5w
◆ 单商户给同一银行卡单日限额——单日5w
接口地址
接口链接：https://api.mch.weixin.qq.com/mmpaysptrans/pay_bank
是否需要证书
请求需要双向证书。 详见证书使用
请求参数
字段名	变量名	必填	示例值	类型	描述
商户号	mch_id	是	1900000109	string(32)	微信支付分配的商户号
商户企业付款单号	partner_trade_no	是	1212121221227	string(32)	商户订单号，需保持唯一（只允许数字[0~9]或字母[A~Z]和[a~z]，最短8位，最长32位）
随机字符串	nonce_str	是	5K8264ILTKCH16CQ2502SI8ZNMTM67Vs	string(32)	随机字符串，不长于32位
签名	sign	是	C380BEC2BFD727A4B6845133519F3AD6	string(32)	通过MD5签名算法计算得出的签名值，详见MD5签名生成算法
收款方银行卡号	enc_bank_no	是	8609cb22e1774a50a930e414cc71eca06121bcd266335cda230d24a7886a8d9f	string(64)	收款方银行卡号（采用标准RSA算法，公钥由微信侧提供）,详见获取RSA加密公钥API
收款方用户名	enc_true_name	是	ca775af5f841bdf424b2e6eb86a6e21e	string(64)	收款方用户名（采用标准RSA算法，公钥由微信侧提供）详见获取RSA加密公钥API
收款方开户行	bank_code	是	1001	string(64)	银行卡所在开户行编号,详见银行编号列表
付款金额	amount	是	100000	int	付款金额：RMB分（支付总额，不含手续费）
注：大于0的整数
付款说明	desc	否	理财	string	企业付款到银行卡付款说明,即订单备注（UTF8编码，允许100个字符以内）
返回参数
字段名	变量名	必填	类型	描述
返回状态码	return_code	是	String(16)	SUCCESS/FAIL
此字段是通信标识，非付款标识，付款是否成功需要查看result_code来判断
返回信息	return_msg	否	String(128)	返回信息，如非空，为错误原因
签名失败
参数格式校验错误
以下字段在return_code为SUCCESS的时候有返回
字段名	变量名	必填	类型	描述
业务结果	result_code	是	string(32)	SUCCESS/FAIL，注意：当状态为FAIL时，存在业务结果未明确的情况，所以如果状态为FAIL，请务必通过查询接口确认此次付款的结果（关注错误码err_code字段）。如果要继续进行这笔付款，请务必用原商户订单号和原参数来重入此接口。
错误代码	err_code
否
string(32)
错误码信息，注意：出现未明确的错误码时，如（SYSTEMERROR）等，请务必用原商户订单号重试，或通过查询接口确认此次付款的结果
错误代码描述	err_code_des	否	string(128)	错误信息描述
商户号	mch_id	是	string(32)	微信支付分配的商户号
商户企业付款单号	partner_trade_no	是	string(32)	商户订单号，需要保持唯一
代付金额	amount	是	int	代付金额RMB:分
随机字符串	nonce_str	是	string(32)	随机字符串，长度小于32位
签名	sign	是	string(32)	返回包携带签名给商户
以下字段在return_code 和result_code都为SUCCESS的时候有返回
字段名	变量名	必填	类型	描述
微信企业付款单号	payment_no	是	string(64)	代付成功后，返回的内部业务单号
手续费金额	cmms_amt	是	int	手续费金额 RMB：分
*/
func (t *Transfer) ToBankCard(m util.Map) core.Response {
	keys := []string{"bank_code", "partner_trade_no", "enc_bank_no", "enc_true_name", "amount"}
	for _, v := range keys {
		if !m.Has(v) {
			log.Error(v + " is required.")
			return nil
		}
	}
	m.Set("mch_id", t.Get("mch_id"))
	m.Set("nonce_str", util.GenerateUUID())

	m.Set("enc_bank_no", crypt.Encrypt(t.GetString("pubkey_path"), m.GetString("enc_bank_no")))
	m.Set("enc_true_name", crypt.Encrypt(t.GetString("pubkey_path"), m.GetString("enc_true_name")))
	m.Set("sign", GenerateSignature(m, t.GetString("key"), MakeSignMD5))
	return t.client.SafeRequest(core.Link(mmPaySpTransPayBankURLSuffix), "post", util.Map{
		core.DataTypeXML: m,
	})
}
