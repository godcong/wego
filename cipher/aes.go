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
	iv, Key string
}

// SetParam ...
func (*cryptAES128CBC) SetParam(key, val string) {
	panic("implement me")
}

// GetParam ...
func (*cryptAES128CBC) GetParam(key string) string {
	panic("implement me")
}

// CryptAES128CBC ...
func CryptAES128CBC() Cipher {
	panic("implement me")
}

// Type ...
func (*cryptAES128CBC) Type() CryptType {
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
