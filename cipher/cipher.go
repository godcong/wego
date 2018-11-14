package cipher

// CryptType ...
type CryptType int

// AES128CBC ...
const (
	AES128CBC CryptType = iota
	RSA                 = iota
)

// InstanceFunc ...
type InstanceFunc func() Cipher

// Cipher ...
type Cipher interface {
	Type() CryptType
	Set(key, val string)
	Get(key string) string
	Encrypt([]byte) ([]byte, error)
	Decrypt(data []byte) ([]byte, error)
}

var cipherList []InstanceFunc

func init() {
	cipherList = []InstanceFunc{
		CryptAES128CBC,
		CryptRSA,
	}
}

// New ...
func New(cryptType CryptType) Cipher {
	return cipherList[cryptType]()
}
