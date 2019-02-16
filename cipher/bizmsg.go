package cipher

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/xml"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
	"golang.org/x/xerrors"
)

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
	"OK":                     OK,
	"ValidateSignatureError": ValidateSignatureError,
	"ParseXMLError":          ParseXMLError,
	"ComputeSignatureError":  ComputeSignatureError,
	"IllegalAesKey":          IllegalAesKey,
	"ValidateAppIDError":     ValidateAppidError,
	"EncryptAESError":        EncryptAESError,
	"DecryptAESError":        DecryptAESError,
	"IllegalBuffer":          IllegalBuffer,
	"EncodeBase64Error":      EncodeBase64Error,
	"DecodeBase64Error":      DecodeBase64Error,
	"GenReturnXMLError":      GenReturnXMLError,
}

/*BizMsg BizMsg */
type cryptBizMsg struct {
	token  string
	key    []byte
	id     string
	cipher Cipher
}

// BizMsgData ...
type BizMsgData struct {
	_            xml.Name `xml:"xml"`
	Text         string   `xml:"-"`
	RSAEncrypt   string   `xml:"RSAEncrypt"`
	TimeStamp    string   `xml:"TimeStamp"`
	Nonce        string   `xml:"Nonce"`
	MsgSignature string   `xml:"MsgSignature"`
}

// Type ...
func (obj *cryptBizMsg) Type() CryptType {
	return BizMsg
}

// Encrypt ...
func (obj *cryptBizMsg) Encrypt(data interface{}) ([]byte, error) {
	bizMsg := parseBizMsg(data)
	buf := bytes.Buffer{}
	buf.WriteString(obj.RandomString())
	buf.Write(obj.LengthBytes(bizMsg.Text))
	buf.WriteString(bizMsg.Text)
	buf.WriteString(obj.id)

	encrypt, e := obj.cipher.Encrypt(buf.Bytes())
	if e != nil {
		return nil, xerrors.Errorf("biz msg encrypt:%w", e)
	}

	r := &BizMsgData{
		Text:         "",
		RSAEncrypt:   string(encrypt),
		TimeStamp:    bizMsg.TimeStamp,
		Nonce:        bizMsg.Nonce,
		MsgSignature: util.SHA1(obj.token, bizMsg.TimeStamp, bizMsg.Nonce, string(encrypt)),
	}

	return xml.Marshal(r)
}

// Decrypt ...
func (obj *cryptBizMsg) Decrypt(data interface{}) ([]byte, error) {
	bizMsg := parseBizMsg(data)
	e := xml.Unmarshal([]byte(bizMsg.Text), bizMsg)
	if e != nil {
		log.Error(e)
		return nil, xerrors.Errorf("biz msg unmarshal:%w", e)
	}
	newSign := util.SHA1(obj.token, bizMsg.TimeStamp, bizMsg.Nonce, bizMsg.RSAEncrypt)
	if bizMsg.MsgSignature != newSign {
		log.Error(bizMsg.MsgSignature, newSign)
		return nil, xerrors.New("ValidateSignatureError")
	}
	decrypt, e := obj.cipher.Decrypt(bizMsg.RSAEncrypt)
	if e != nil {
		return nil, xerrors.Errorf("biz msg decrypt:%w", e)
	}

	buf := bytes.NewBuffer(decrypt)
	_ = buf.Next(16)                     //skip first 16 random string
	size := obj.BytesLength(buf.Next(4)) //size:4 bit
	content := buf.Next(int(size))       //content:size
	id := buf.Bytes()                    //end:id
	if string(id) != obj.id {
		return nil, xerrors.New("ValidateAppIDError")
	}
	return content, nil

}

/*NewBizMsg NewBizMsg */
func NewBizMsg(option *Option) Cipher {
	key, e := base64.RawStdEncoding.DecodeString(option.Key)
	if e != nil {
		return nil
	}
	return &cryptBizMsg{
		token: option.Token,
		key:   key,
		id:    option.ID,
		cipher: &cryptAES128CBC{
			iv:  key[:16],
			key: key,
		},
	}
}

/*RandomString RandomString */
func (obj *cryptBizMsg) RandomString() string {
	return util.GenerateRandomString(16, util.RandomAll)
}

/*LengthBytes LengthBytes */
func (obj *cryptBizMsg) LengthBytes(s string) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(len(s)))
	return buf
}

/*BytesLength BytesLength */
func (obj *cryptBizMsg) BytesLength(b []byte) uint32 {
	return binary.BigEndian.Uint32(b)
}
