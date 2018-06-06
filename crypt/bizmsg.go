package crypt

import (
	"encoding/base64"
	"errors"

	"github.com/godcong/wego/util"
)

/*BizMsg BizMsg */
type BizMsg struct {
	token          string
	encodingAESKey []byte
	appID          string
}

/*ErrorCodeType ErrorCodeType */
type ErrorCodeType int

/*error code types */
const (
	OK                     ErrorCodeType = 0
	ValidateSignatureError               = -40001
	ParseXMLError                        = -40002
	ComputeSignatureError                = -40003
	IllegalAesKey                        = -40004
	ValidateAppidError                   = -40005
	EncryptAESError                      = -40006
	DecryptAESError                      = -40007
	IllegalBuffer                        = -40008
	EncodeBase64Error                    = -40009
	DecodeBase64Error                    = -40010
	GenReturnXMLError                    = -40011
)

/*ErrorCode ErrorCode */
var ErrorCode = map[string]ErrorCodeType{
	"OK": OK,
	"ValidateSignatureError": ValidateSignatureError,
	"ParseXMLError":          ParseXMLError,
	"ComputeSignatureError":  ComputeSignatureError,
	"IllegalAesKey":          IllegalAesKey,
	"ValidateAppidError":     ValidateAppidError,
	"EncryptAESError":        EncryptAESError,
	"DecryptAESError":        DecryptAESError,
	"IllegalBuffer":          IllegalBuffer,
	"EncodeBase64Error":      EncodeBase64Error,
	"DecodeBase64Error":      DecodeBase64Error,
	"GenReturnXMLError":      GenReturnXMLError,
}

/*NewBizMsg NewBizMsg */
func NewBizMsg(token, key, id string) *BizMsg {
	k, _ := base64.RawStdEncoding.DecodeString(key)
	return &BizMsg{
		token:          token,
		encodingAESKey: k,
		appID:          id,
	}
}

/*Encrypt Encrypt */
func (m *BizMsg) Encrypt(text, timeStamp, nonce string) (string, error) {
	prp := NewPrp(m.encodingAESKey)
	b, err := prp.Encrypt(text, m.appID)
	if err != nil {
		return "", err
	}

	p := util.Map{
		"Encrypt":      string(b),
		"MsgSignature": SHA1(m.token, timeStamp, nonce, string(b)),
		"TimeStamp":    timeStamp,
		"Nonce":        nonce,
	}
	return p.ToXML(), nil
}

/*Decrypt Decrypt */
func (m *BizMsg) Decrypt(text string, msgSignature, timeStamp, nonce string) ([]byte, error) {
	p := util.XMLToMap([]byte(text))
	enpt := p.GetString("Encrypt")
	tSign := SHA1(m.token, timeStamp, nonce, enpt)
	if msgSignature != tSign {
		return nil, errors.New("ValidateSignatureError")
	}
	prp := NewPrp(m.encodingAESKey)
	b, err := prp.Decrypt([]byte(enpt), m.appID)
	return b, err
}
