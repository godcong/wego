package crypt

import (
	"encoding/base64"
	"errors"

	"github.com/godcong/wego/util"
)

type BizMsg struct {
	token          string
	encodingAESKey []byte
	appId          string
}
type ErrorCodeType int

const OK ErrorCodeType = 0
const ValidateSignatureError = -40001
const ParseXmlError = -40002
const ComputeSignatureError = -40003
const IllegalAesKey = -40004
const ValidateAppidError = -40005
const EncryptAESError = -40006
const DecryptAESError = -40007
const IllegalBuffer = -40008
const EncodeBase64Error = -40009
const DecodeBase64Error = -40010
const GenReturnXmlError = -40011

var ErrorCode = map[string]ErrorCodeType{
	"OK": OK,
	"ValidateSignatureError": ValidateSignatureError,
	"ParseXmlError":          ParseXmlError,
	"ComputeSignatureError":  ComputeSignatureError,
	"IllegalAesKey":          IllegalAesKey,
	"ValidateAppidError":     ValidateAppidError,
	"EncryptAESError":        EncryptAESError,
	"DecryptAESError":        DecryptAESError,
	"IllegalBuffer":          IllegalBuffer,
	"EncodeBase64Error":      EncodeBase64Error,
	"DecodeBase64Error":      DecodeBase64Error,
	"GenReturnXmlError":      GenReturnXmlError,
}

func NewBizMsg(token, key, id string) *BizMsg {
	k, _ := base64.RawStdEncoding.DecodeString(key)
	return &BizMsg{
		token:          token,
		encodingAESKey: k,
		appId:          id,
	}
}

func (m *BizMsg) Encrypt(text, timeStamp, nonce string) (string, error) {
	prp := NewPrp(m.encodingAESKey)
	b, err := prp.Encrypt(text, m.appId)
	if err != nil {
		return "", err
	}

	p := util.Map{
		"Encrypt":      string(b),
		"MsgSignature": SHA1(m.token, timeStamp, nonce, string(b)),
		"TimeStamp":    timeStamp,
		"Nonce":        nonce,
	}
	return p.ToXml(), nil
}

func (m *BizMsg) Decrypt(text string, msgSignature, timeStamp, nonce string) ([]byte, error) {
	p := util.XmlToMap([]byte(text))
	enpt := p.GetString("Encrypt")
	tSign := SHA1(m.token, timeStamp, nonce, enpt)
	if msgSignature != tSign {
		return nil, errors.New("ValidateSignatureError")
	}
	prp := NewPrp(m.encodingAESKey)
	b, err := prp.Decrypt([]byte(enpt), m.appId)
	return b, err
}
