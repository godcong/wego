package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/binary"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/godcong/wego/core"
	"github.com/godcong/wego/core/tool"
)

type PrpCrypt struct {
	key []byte
}

func NewPrp(key []byte) *PrpCrypt {
	//k, err := base64.RawStdEncoding.DecodeString(key)
	//if err != nil {
	//	return &PrpCrypt{}
	//}
	return &PrpCrypt{
		key: key,
	}
}

func (c *PrpCrypt) Random() string {
	return tool.GenerateRandomString(16, tool.T_RAND_ALL)
}

func (c *PrpCrypt) LengthBytes(s string) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(len(s)))
	return buf
}

func (c *PrpCrypt) BytesLength(b []byte) uint32 {
	return binary.BigEndian.Uint32(b)
}

func (c *PrpCrypt) Encrypt(text string, appId string) ([]byte, error) {
	buf := bytes.Buffer{}

	buf.WriteString(c.Random())
	buf.Write(c.LengthBytes(text))
	buf.WriteString(text)
	buf.WriteString(appId)
	iv := c.key[:16]
	block, err := aes.NewCipher([]byte(c.key)) //选择加密算法
	if err != nil {
		return nil, err
	}
	plantText := PKCS7Padding(buf.Bytes(), block.BlockSize())

	blockModel := cipher.NewCBCEncrypter(block, []byte(iv))

	ciphertext := make([]byte, len(plantText))

	blockModel.CryptBlocks(ciphertext, plantText)
	return Base64Encode(ciphertext), nil
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func (c *PrpCrypt) Decrypt(ciphertext []byte, appId string) ([]byte, error) {
	ciphertext, err := Base64Decode(ciphertext)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(c.key) //选择加密算法
	if err != nil {
		return nil, err
	}
	iv := c.key[:16]
	blockModel := cipher.NewCBCDecrypter(block, iv)
	plantText := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plantText, ciphertext)
	plantText = PKCS7UnPadding(plantText)
	content := plantText[16:]
	xmlLen := c.BytesLength(content[0:4])
	xmlContent := content[4 : xmlLen+4]
	fromAppId := content[xmlLen+4:]
	if string(fromAppId) != appId {
		return nil, errors.New("ValidateAppidError")
	}

	return xmlContent, nil
}

func PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	if unpadding < 1 || unpadding > 32 {
		unpadding = 0
	}
	return plantText[:(length - unpadding)]
}

func SHA1(text ...string) string {
	sort.Strings(text)
	s := strings.Join(text, "")
	core.Debug(s)
	return fmt.Sprintf("%x", sha1.Sum([]byte(s)))
}
