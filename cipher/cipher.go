package cipher

// CryptType ...
type CryptType string

// AES128CBC ...
const AES128CBC = "AES-128-CBC"

// Cipher ...
type Cipher interface {
	Type() CryptType
}

type cipher struct {
	cryptType CryptType
}

// NewCipher ...
func NewCipher(cryptType CryptType) {

}

// Type ...
func (c *cipher) Type() CryptType {
	return c.cryptType
}
