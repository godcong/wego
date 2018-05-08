package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"strings"
)

type DataCrypt struct {
	appId string
}

func NewDataCrypt(id string) *DataCrypt {
	return &DataCrypt{
		appId: id,
	}
}

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

	//过滤所有 结尾
	idx := strings.LastIndex(string(dData), "}") + 1

	return dData[:idx], nil
}
