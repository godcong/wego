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
	Iv, Key string
}

// Type ...
func (c *cryptAES128CBC) Type() CryptType {
	return AES128CBC
}

// Set ...
func (c *cryptAES128CBC) Set(key, val string) {
	switch key {
	case "key":
	case "iv":
	}
}

// Get ...
func (c *cryptAES128CBC) Get(key string) string {
	panic("implement me")
}

// Encrypt ...
func (c *cryptAES128CBC) Encrypt([]byte) ([]byte, error) {
	panic("implement me")
}

// Decrypt ...
func (c *cryptAES128CBC) Decrypt(data []byte) ([]byte, error) {
	panic("implement me")
}

// CryptAES128CBC ...
func CryptAES128CBC() Cipher {
	panic("implement me")
}

// DecryptAES128CBC ...
func decryptAES128CBC(data, iv, key string) ([]byte, error) {
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
