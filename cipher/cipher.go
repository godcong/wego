package cipher

import (
	"encoding/base64"
	"golang.org/x/xerrors"
)

// CryptType ...
type CryptType int

// AES128CBC ...
const (
	AES128CBC CryptType = iota
	AES256ECB           = iota
	BizMsg              = iota
	RSA
)

// InstanceFunc ...
type InstanceFunc func(option *Option) Cipher

// Cipher ...
type Cipher interface {
	Type() CryptType
	Encrypt(interface{}) ([]byte, error)
	Decrypt(interface{}) ([]byte, error)
}

var cipherList = []InstanceFunc{
	AES128CBC: NewAES128CBC,
	AES256ECB: NewAES256ECB,
	BizMsg:    NewBizMsg,
	RSA:       NewRSA,
}

// Option ...
type Option struct {
	IV         string
	Key        string
	RSAPrivate string
	RSAPublic  string
	Token      string
	ID         string
}

// New create a new cipher
func New(cryptType CryptType, option *Option) Cipher {
	return cipherList[cryptType](option)
}

func parseBytes(data interface{}) []byte {
	switch tmp := data.(type) {
	case []byte:
		return tmp
	case string:
		return []byte(tmp)
	default:
		return nil
	}
}

func parseBizMsg(data interface{}) (d *BizMsgData, e error) {
	switch tmp := data.(type) {
	case *BizMsgData:
		d = tmp
	case BizMsgData:
		d = &tmp
	default:
		e = xerrors.New("wrong type inputed")
	}
	return
}

/*Base64Encode Base64Encode */
func Base64Encode(b []byte) []byte {
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(b)))
	base64.StdEncoding.Encode(buf, b)
	return buf
}

/*Base64Decode Base64Decode */
func Base64Decode(b []byte) ([]byte, error) {
	buf := make([]byte, base64.StdEncoding.DecodedLen(len(b)))
	n, err := base64.StdEncoding.Decode(buf, b)
	return buf[:n], err
}

/*Base64DecodeString Base64DecodeString */
func Base64DecodeString(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}
