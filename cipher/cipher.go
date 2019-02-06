package cipher

// CryptType ...
type CryptType int

// AES128CBC ...
const (
	AES128CBC CryptType = iota
	AES256ECB           = iota
	RSA                 = iota
)

// InstanceFunc ...
type InstanceFunc func(option Option) Cipher

// Cipher ...
type Cipher interface {
	Type() CryptType
	Encrypt(interface{}) ([]byte, error)
	Decrypt(interface{}) ([]byte, error)
}

var cipherList []InstanceFunc

func init() {
	cipherList = []InstanceFunc{
		AES128CBC: NewAES128CBC,
		//CryptAES128CBC,
		//CryptRSA,
	}
}

// Option ...
type Option struct {
	IV    string
	Key   string
	Token string
	AppID string
}

// New create a new cipher
func New(cryptType CryptType, option Option) Cipher {
	return cipherList[cryptType](option)
}
