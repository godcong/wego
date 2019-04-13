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
type InstanceFunc func(opts *Options) Cipher

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

// Options ...
type Options struct {
	IV         string
	Key        string
	RSAPrivate string
	RSAPublic  string
	Token      string
	ID         string
}

// OptionKey ...
func OptionKey(s string) func(opts *Options) {
	return func(opts *Options) {
		opts.Key = s
	}
}

// OptionIV ...
func OptionIV(s string) func(opts *Options) {
	return func(opts *Options) {
		opts.IV = s
	}
}

// OptionToken ...
func OptionToken(s string) func(opts *Options) {
	return func(opts *Options) {
		opts.Token = s
	}
}

// OptionID ...
func OptionID(s string) func(opts *Options) {
	return func(opts *Options) {
		opts.ID = s
	}
}

// OptionPublic ...
func OptionPublic(s string) func(opts *Options) {
	return func(opts *Options) {
		opts.RSAPublic = s
	}
}

// OptionPrivate ...
func OptionPrivate(s string) func(opts *Options) {
	return func(opts *Options) {
		opts.RSAPrivate = s
	}
}

// Option ...
type Option func(opts *Options)

// New create a new cipher
func New(cryptType CryptType, opts ...Option) Cipher {
	var options Options
	for _, o := range opts {
		o(&options)
	}
	return cipherList[cryptType](&options)
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
	case string:
		d = &BizMsgData{
			Text: tmp,
		}
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
