package cipher

import (
	"crypto/aes"
	"crypto/cipher"
)

/*DataCrypt DataCrypt */
type DataCrypt struct {
	appID string
}

/*NewDataCrypt NewDataCrypt */
func NewDataCrypt(id string) *DataCrypt {
	return &DataCrypt{
		appID: id,
	}
}

// CryptAES128CBC ...
type cryptAES128CBC struct {
	Iv, Key []byte
}

// Type ...
func (c *cryptAES128CBC) Type() CryptType {
	return AES128CBC
}

// Set ...
func (c *cryptAES128CBC) SetParameter(key string, val []byte) {
	switch key {
	case "key":
		c.Key = val
	case "iv":
		c.Iv = val
	}
}

// Get ...
func (c *cryptAES128CBC) GetParameter(key string) []byte {
	switch key {
	case "key":
		return c.Key
	case "iv":
		return c.Iv
	}
	return nil
}

// Encrypt ...
func (c *cryptAES128CBC) Encrypt([]byte) ([]byte, error) {
	panic("implement me")
}

// Decrypt ...
func (c *cryptAES128CBC) Decrypt(data []byte) ([]byte, error) {
	key, e := Base64Decode(c.Key)
	if e != nil {
		return nil, e
	}

	iv, e := Base64Decode(c.Iv)
	if e != nil {
		return nil, e
	}

	decodeData, e := Base64Decode(data)
	if e != nil {
		return nil, e
	}

	block, e := aes.NewCipher(key)
	if e != nil {
		return nil, e
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(decodeData, decodeData)

	return PKCS7UnPadding(decodeData), nil
}

// CryptAES128CBC ...
func CryptAES128CBC() Cipher {
	return &cryptAES128CBC{}
}

type cryptAES256ECB struct {
	Key []byte
}

func (c *cryptAES256ECB) Type() CryptType {
	panic("implement me")
}

func (c *cryptAES256ECB) SetParameter(key string, val []byte) {
	c.Key = val
}

func (c *cryptAES256ECB) GetParameter(key string) []byte {
	return c.Key
}

func (c *cryptAES256ECB) Encrypt([]byte) ([]byte, error) {
	panic("implement me")
}

func (c *cryptAES256ECB) Decrypt(data []byte) ([]byte, error) {
	decodeData, e := Base64Decode(data)
	if e != nil {
		return nil, e
	}

	block, err := aes.NewCipher(c.Key)
	if err != nil {
		return nil, err
	}

	mode := NewECBDecrypter(block)

	mode.CryptBlocks(decodeData, decodeData)

	return PKCS7UnPadding(decodeData), nil
}

func CryptAES256ECB() Cipher {
	return &cryptAES256ECB{}
}

// Decrypt ...
//Deprecated: change to cipher
func (c *DataCrypt) Decrypt(data, iv, key string) ([]byte, error) {
	dKey, e := Base64Decode([]byte(key))
	if e != nil {
		return nil, e
	}
	dData, e := Base64Decode([]byte(data))
	if e != nil {
		return nil, e
	}

	dIv, e := Base64Decode([]byte(iv))
	if e != nil {
		return nil, e
	}

	block, e := aes.NewCipher(dKey)
	if e != nil {
		return nil, e
	}

	mode := cipher.NewCBCDecrypter(block, dIv)

	mode.CryptBlocks(dData, dData)

	return PKCS7UnPadding(dData), nil
}
