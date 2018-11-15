package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"strings"
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

	decodeData, e := Base64Decode([]byte(data))
	if e != nil {
		return nil, e
	}

	block, e := aes.NewCipher(key)
	if e != nil {
		return nil, e
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(decodeData, decodeData)

	//过滤所有 非正常字符结尾
	idx := strings.LastIndex(string(decodeData), "}") + 1

	return decodeData[:idx], nil
}

// CryptAES128CBC ...
func CryptAES128CBC() Cipher {
	return &cryptAES128CBC{}
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

	//过滤所有 非正常字符结尾
	idx := strings.LastIndex(string(dData), "}") + 1

	return dData[:idx], nil
}
