package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"strings"
)

type DataCrypt struct {
	id  string
	key string
	//cipher string
	//length int
}

func NewDataCrypt(id, key string) *DataCrypt {
	return &DataCrypt{
		id:  id,
		key: key,
	}
}

func (c *DataCrypt) Decrypt(data, iv string) ([]byte, error) {
	key, e := Base64Decode([]byte(c.key))
	if e != nil {
		return nil, e
	}
	decodeData, e := Base64Decode([]byte(data))
	if e != nil {
		return nil, e
	}

	decodeIv, e := Base64Decode([]byte(iv))
	if e != nil {
		return nil, e
	}

	block, e := aes.NewCipher(key)
	if e != nil {
		return nil, e
	}

	mode := cipher.NewCBCDecrypter(block, decodeIv)

	mode.CryptBlocks(decodeData, decodeData)

	//过滤所有 结尾
	idx := strings.LastIndex(string(decodeData), "}") + 1

	return decodeData[:idx], nil
}
