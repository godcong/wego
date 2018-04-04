package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"log"

	"github.com/godcong/wego/core"
	"github.com/godcong/wego/rsa"
)

type DataCrypt struct {
	id     string
	key    string
	cipher string
	length int
}

func NewAES128CBCDataCrypt(id, key string) *DataCrypt {
	return &DataCrypt{
		id:     id,
		key:    key,
		cipher: "AES-128-CBC",
		length: 16,
	}
}

func NewAES256CBCDataCrypt(id, key string) *DataCrypt {
	return &DataCrypt{
		id:     id,
		key:    key,
		cipher: "AES-256-CBC",
		length: 32,
	}
}

func (c *DataCrypt) Decrypt(data, iv string) (core.Map, error) {
	//todo:
	// Load your secret key from a safe place and reuse it across multiple
	// NewCipher calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	key, e := rsa.Base64Decode([]byte(c.key))
	//log.Println(b)
	//key, e := hex.DecodeString(string(b))
	if e != nil {
		log.Println(e)
		return nil, e
	}
	decodeData, e := rsa.Base64Decode([]byte(data))
	//log.Println("data", string(d1))
	//dec := hex.NewDecoder(bytes.NewReader(d1))
	//decodeData, e := hex.DecodeString(string(d1))
	//var by []byte
	//dec.Read(by)
	//log.Println("data", by)
	if e != nil {
		log.Println(e)
		return nil, e
	}
	//decodeIv, e := hex.DecodeString(iv)
	decodeIv, e := rsa.Base64Decode([]byte(iv))
	if e != nil {
		log.Println(e)
		return nil, e
	}

	block, e := aes.NewCipher(key)
	if e != nil {
		log.Println(e)
		return nil, e
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	//if len(decodeData) < aes.BlockSize {
	//	panic("ciphertext too short")
	//}
	//ivs := decodeData[:aes.BlockSize]
	//decodeData = decodeData[aes.BlockSize:]
	//
	//// CBC mode always works in whole blocks.
	//if len(decodeData)%aes.BlockSize != 0 {
	//	panic("ciphertext is not a multiple of the block size")
	//}

	mode := cipher.NewCBCDecrypter(block, decodeIv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(decodeData, decodeData)

	// If the original plaintext lengths are not a multiple of the block
	// size, padding would have to be added when encrypting, which would be
	// removed at this point. For an example, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. However, it's
	// critical to note that ciphertexts must be authenticated (i.e. by
	// using crypto/hmac) before being decrypted in order to avoid creating
	// a padding oracle.

	log.Println("dec", decodeData)
	// Output: exampleplaintext
	return nil, nil
}
