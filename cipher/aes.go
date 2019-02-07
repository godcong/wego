package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"github.com/godcong/wego/log"
	"golang.org/x/exp/xerrors"
)

// CryptAES128CBC ...
type cryptAES128CBC struct {
	iv, key []byte
}

// Encrypt ...
func (c *cryptAES128CBC) Encrypt(data interface{}) ([]byte, error) {
	block, err := aes.NewCipher(c.key) //选择加密算法
	if err != nil {
		return nil, err
	}
	plantText := PKCS7Padding(parseBytes(data), block.BlockSize())
	mode := cipher.NewCBCEncrypter(block, c.iv)
	cipherText := make([]byte, len(plantText))
	mode.CryptBlocks(cipherText, plantText)

	return Base64Encode(cipherText), nil
}

// Decrypt ...
func (c *cryptAES128CBC) Decrypt(data interface{}) ([]byte, error) {
	cipherText, e := Base64Decode(parseBytes(data))
	if e != nil {
		return nil, xerrors.Errorf("wrong data:%w", e)
	}
	block, e := aes.NewCipher(c.key)
	if e != nil {
		return nil, e
	}
	mode := cipher.NewCBCDecrypter(block, c.iv)
	plantText := make([]byte, len(cipherText))
	mode.CryptBlocks(plantText, cipherText)

	return PKCS7UnPadding(plantText), nil
}

// NewAES128CBC ...
func NewAES128CBC(option *Option) Cipher {
	key, e := Base64DecodeString(option.Key)
	if e != nil {
		log.Error("wrong key:%w", e)
		return nil
	}
	iv, e := Base64DecodeString(option.IV)
	if e != nil {
		log.Error("wrong iv:%w", e)
		return nil
	}
	return &cryptAES128CBC{
		iv:  iv,
		key: key,
	}
}

// Type ...
func (c *cryptAES128CBC) Type() CryptType {
	return AES128CBC
}

type cryptAES256ECB struct {
	key []byte
}

// NewAES256ECB ...
func NewAES256ECB(option *Option) Cipher {
	return &cryptAES256ECB{
		key: []byte(option.Key),
	}
}

// Encrypt ...
func (c *cryptAES256ECB) Encrypt(interface{}) ([]byte, error) {
	panic("aes 256 ecb encrypt was not support")
}

// Decrypt ...
func (c *cryptAES256ECB) Decrypt(data interface{}) ([]byte, error) {
	decodeData, e := Base64Decode(parseBytes(data))
	if e != nil {
		return nil, e
	}

	block, err := aes.NewCipher(c.key)
	if err != nil {
		return nil, err
	}

	mode := NewECBDecrypter(block)
	mode.CryptBlocks(decodeData, decodeData)
	return PKCS7UnPadding(decodeData), nil
}

// Type ...
func (c *cryptAES256ECB) Type() CryptType {
	return AES256ECB
}
