package payment

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"github.com/godcong/wego/cache"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
	"io"
	"strings"
)

/*SignType SignType */
type SignType int

/*sign types */
const (
	SignTypeMd5        SignType = iota
	SignTypeHmacSha256 SignType = iota
)

func (t SignType) String() string {
	if t == SignTypeHmacSha256 {
		return HMACSHA256
	}
	return MD5
}

/*Sandbox 沙箱 */
type Sandbox struct {
	*Payment
}

func newSandbox(payment *Payment) *Sandbox {
	return &Sandbox{
		Payment: payment,
	}
}

/*NewSandbox NewSandbox */
func NewSandbox(config *core.Config) *Sandbox {
	return newSandbox(NewPayment(config))
}

/*GetKey 沙箱key(string类型) */
func (s *Sandbox) GetKey() string {
	key := cache.Get(s.cacheName())
	if key != nil {
		return key.(string)
	}

	//TODO

	return string(s.SandboxSignKey())
}

func (s *Sandbox) cacheName() string {
	name := s.GetString("app_id") + s.GetString("mch_id")
	return "godcong.wego.payment.sandbox." + fmt.Sprintf("%x", md5.Sum([]byte(name)))
}

/*SandboxSignKey 沙箱key */
func (s *Sandbox) SandboxSignKey() []byte {
	m := make(util.Map)
	m.Set("mch_id", s.Get("mch_id"))
	m.Set("nonce_str", util.GenerateNonceStr())
	sign := GenerateSignature(m, s.GetString("aes_key"), MakeSignMD5)
	m.Set("sign", sign)
	resp := s.client.RequestRaw(Link(sandboxSignKeyURLSuffix), "post", m)

	return resp

}

/*SignFunc sign函数定义 */
type SignFunc func(data, key string) string

// MakeSignHMACSHA256 make sign with hmac-sha256
func MakeSignHMACSHA256(data, key string) string {
	m := hmac.New(sha256.New, []byte(key))
	m.Write([]byte(data))
	return strings.ToUpper(fmt.Sprintf("%x", m.Sum(nil)))
}

// MakeSignMD5 make sign with md5
func MakeSignMD5(data, key string) string {
	m := md5.New()
	_, _ = io.WriteString(m, data)

	return strings.ToUpper(fmt.Sprintf("%x", m.Sum(nil)))
}

// GenerateSignature make sign from map data
func GenerateSignature(m util.Map, key string, fn SignFunc) string {
	keys := m.SortKeys()
	var sign []string

	for _, k := range keys {
		if k == FieldSign {
			continue
		}
		v := strings.TrimSpace(m.GetString(k))

		if len(v) > 0 {
			log.Debug(k, v)
			sign = append(sign, strings.Join([]string{k, v}, "="))
		}
	}

	sign = append(sign, strings.Join([]string{"key", key}, "="))
	sb := strings.Join(sign, "&")
	return fn(sb, key)
}
