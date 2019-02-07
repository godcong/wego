package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"golang.org/x/exp/xerrors"
)

// CryptAES128CBC ...
type cryptAES128CBC struct {
	iv, key []byte
}

// Encrypt ...
func (c *cryptAES128CBC) Encrypt(interface{}) ([]byte, error) {
	panic("aes 128 cbc encrypt was not support")
}

// Decrypt ...
func (c *cryptAES128CBC) Decrypt(data interface{}) ([]byte, error) {
	key, e := Base64Decode(c.key)
	if e != nil {
		return nil, xerrors.Errorf("wrong key:%w", e)
	}

	iv, e := Base64Decode(c.iv)
	if e != nil {
		return nil, xerrors.Errorf("wrong iv:%w", e)
	}

	decoded, e := Base64Decode(parseBytes(data))
	if e != nil {
		return nil, xerrors.Errorf("wrong data:%w", e)
	}

	block, e := aes.NewCipher(key)
	if e != nil {
		return nil, e
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decoded, decoded)
	return PKCS7UnPadding(decoded), nil
}

// NewAES128CBC ...
func NewAES128CBC(option Option) Cipher {
	return &cryptAES128CBC{
		iv:  []byte(option.IV),
		key: []byte(option.Key),
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
func NewAES256ECB(option Option) Cipher {
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
