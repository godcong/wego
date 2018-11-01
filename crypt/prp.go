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

	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
)

/*PrpCrypt PrpCrypt */
type PrpCrypt struct {
	key []byte
}

/*NewPrp NewPrp */
func NewPrp(key []byte) *PrpCrypt {
	return &PrpCrypt{
		key: key,
	}
}

/*Random Random */
func (c *PrpCrypt) Random() string {
	return util.GenerateRandomString(16, util.RandomAll)
}

/*LengthBytes LengthBytes */
func (c *PrpCrypt) LengthBytes(s string) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(len(s)))
	return buf
}

/*BytesLength BytesLength */
func (c *PrpCrypt) BytesLength(b []byte) uint32 {
	return binary.BigEndian.Uint32(b)
}

// Encrypt ...
func (c *PrpCrypt) Encrypt(text string, appid string) ([]byte, error) {
	buf := bytes.Buffer{}

	buf.WriteString(c.Random())
	buf.Write(c.LengthBytes(text))
	buf.WriteString(text)
	buf.WriteString(appid)
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

/*PKCS7Padding PKCS7Padding */
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// Decrypt ...
func (c *PrpCrypt) Decrypt(ciphertext []byte, appid string) ([]byte, error) {
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
	fromAppID := content[xmlLen+4:]
	if string(fromAppID) != appid {
		return nil, errors.New("ValidateAppidError")
	}

	return xmlContent, nil
}

/*PKCS7UnPadding PKCS7UnPadding */
func PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	if unpadding < 1 || unpadding > 32 {
		unpadding = 0
	}
	return plantText[:(length - unpadding)]
}

/*SHA1 SHA1 */
func SHA1(text ...string) string {
	sort.Strings(text)
	s := strings.Join(text, "")
	log.Debug(s)
	return fmt.Sprintf("%x", sha1.Sum([]byte(s)))
}
