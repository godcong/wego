package cipher

// CryptType ...
type CryptType string

// AES128CBC ...
const AES128CBC = "AES-128-CBC"

// Cipher ...
type Cipher interface {
	Type() CryptType
	SetParam(key, val string)
	GetParam(key string) string
	Encrypt([]byte) ([]byte, error)
	Decrypt(data []byte) ([]byte, error)
}

var cipherList []Cipher

func init() {
	cipherList = []Cipher{
		CryptAES128CBC(),
	}
}

// NewCipher ...
func NewCipher(cryptType CryptType) Cipher {

}

//
//// Type ...
//func (c *cipher) Type() CryptType {
//	return c.cryptType
//}
