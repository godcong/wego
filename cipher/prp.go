package cipher

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"golang.org/x/exp/xerrors"
	"sort"
	"strings"

	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
)

/*AESCrypt AESCrypt */
type AESCrypt struct {
	key []byte
}

/*NewPrp NewPrp */
func NewPrp(key []byte) *AESCrypt {
	return &AESCrypt{
		key: key,
	}
}

/*Random Random */
func (c *AESCrypt) Random() string {
	return util.GenerateRandomString(16, util.RandomAll)
}

/*LengthBytes LengthBytes */
func (c *AESCrypt) LengthBytes(s string) []byte {
	var buf = make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(len(s)))
	return buf
}

/*BytesLength BytesLength */
func (c *AESCrypt) BytesLength(b []byte) uint32 {
	return binary.BigEndian.Uint32(b)
}

// Encrypt ...
func (c *AESCrypt) Encrypt(text string, appid string) ([]byte, error) {
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

// Decrypt ...
func (c *AESCrypt) Decrypt(ciphertext []byte, appid string) ([]byte, error) {
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
		return nil, xerrors.New("ValidateAppidError")
	}
	return xmlContent, nil
}

/*SHA1 SHA1 */
func SHA1(text ...string) string {
	sort.Strings(text)
	s := strings.Join(text, "")
	log.Debug(s)
	return fmt.Sprintf("%x", sha1.Sum([]byte(s)))
}
