package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"golang.org/x/exp/xerrors"
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

	var decoded []byte
	switch tmp := data.(type) {
	case []byte:
		decoded, e = Base64Decode(tmp)
	case string:
		decoded, e = Base64Decode([]byte(decoded))
	default:
		return nil, xerrors.New("wrong input data")
	}
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

// CryptAES128CBC ...
func CryptAES128CBC() Cipher {
	return &cryptAES128CBC{}
}

type cryptAES256ECB struct {
	Key []byte
}

// Encrypt ...
func (c *cryptAES256ECB) Encrypt(interface{}) ([]byte, error) {
	panic("implement me")
}

// Decrypt ...
func (c *cryptAES256ECB) Decrypt(interface{}) ([]byte, error) {
	panic("implement me")
}

// Type ...
func (c *cryptAES256ECB) Type() CryptType {
	return AES256ECB
}

// SetParameter ...
func (c *cryptAES256ECB) SetParameter(key string, val []byte) {
	c.Key = val
}

// GetParameter ...
func (c *cryptAES256ECB) GetParameter(key string) []byte {
	return c.Key
}

// Encrypt ...
func (c *cryptAES256ECB) Encrypt2([]byte) ([]byte, error) {
	panic("implement me")
}

// Decrypt ...
func (c *cryptAES256ECB) Decrypt2(data []byte) ([]byte, error) {
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

// CryptAES256ECB ...
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
