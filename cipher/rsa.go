package cipher

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io/ioutil"

	"github.com/godcong/wego/log"
)

/* all defined errors */
var (
	ErrorKeyMustBePEMEncoded = errors.New("key must be pem encoded")
	ErrorNotECPublicKey      = errors.New("key is not a valid ECDSA public key")
	ErrorNotECPrivateKey     = errors.New("key is not a valid ECDSA private key")
	ErrorNotRSAPrivateKey    = errors.New("key is not a valid RSA private key")
	ErrorNotRSAPublicKey     = errors.New("key is not a valid RSA public key")
)

type cryptRSA struct {
	KeyPath []byte
}

// Type ...
func (*cryptRSA) Type() CryptType {
	return RSA
}

// SetParameter ...
func (c *cryptRSA) SetParameter(key string, val []byte) {
	c.KeyPath = val
}

// GetParameter ...
func (c *cryptRSA) GetParameter(key string) []byte {
	return c.KeyPath
}

// Encrypt ...
func (c *cryptRSA) Encrypt(data []byte) ([]byte, error) {

	publicKey, err := ioutil.ReadFile(string(c.KeyPath))
	if err != nil {
		log.Debug(err)
		return nil, err
	}
	key, err := ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		log.Debug(err)
		return nil, err
	}
	part, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, key, []byte(data), nil)
	if err != nil {
		log.Debug(err)
		return nil, err
	}

	buf := make([]byte, base64.StdEncoding.EncodedLen(len(part)))
	base64.StdEncoding.Encode(buf, part)
	return buf, nil
}

// Decrypt ...
func (c *cryptRSA) Decrypt(data []byte) ([]byte, error) {
	privateKey, err := ioutil.ReadFile(string(c.KeyPath))
	if err != nil {
		return nil, err
	}

	key, err := ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return nil, err
	}

	t, err := Base64Decode(data)
	if err != nil {
		return nil, err
	}

	b, err := rsa.DecryptOAEP(sha1.New(), rand.Reader, key, t, nil)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// CryptRSA ...
func CryptRSA() Cipher {
	return &cryptRSA{}
}

/*ParseRSAPrivateKeyFromPEM Parse PEM encoded PKCS1 or PKCS8 private key */
func ParseRSAPrivateKeyFromPEM(key []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, ErrorKeyMustBePEMEncoded
	}

	pkey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		if parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
			var ok bool
			if pkey, ok = parsedKey.(*rsa.PrivateKey); !ok {
				return nil, ErrorNotRSAPrivateKey
			}
		}
	}
	return pkey, nil
}

/*ParseRSAPublicKeyFromPEM Parse PEM encoded PKCS1 or PKCS8 public key */
func ParseRSAPublicKeyFromPEM(key []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, ErrorKeyMustBePEMEncoded
	}

	// Parse the key
	parsedKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		if cert, err := x509.ParseCertificate(block.Bytes); err == nil {
			parsedKey = cert.PublicKey
		} else {
			return nil, err
		}
	}

	if pkey, ok := parsedKey.(*rsa.PublicKey); ok {
		return pkey, nil
	}

	return nil, ErrorNotRSAPublicKey

}

/*RSADecrypt RSADecrypt */
func RSADecrypt(pri string, text string) string {
	privateKey, e := ioutil.ReadFile(pri)
	if e != nil {
		return ""
	}

	key, e := ParseRSAPrivateKeyFromPEM(privateKey)
	if e != nil {
		return ""
	}

	t, e := Base64Decode([]byte(text))
	if e != nil {
		return ""
	}

	b, e := rsa.DecryptOAEP(sha1.New(), rand.Reader, key, t, nil)
	if e != nil {
		return ""
	}
	return string(b)
}

/*RSAEncrypt RSAEncrypt */
func RSAEncrypt(pub string, text string) string {
	publicKey, e := ioutil.ReadFile(pub)
	if e != nil {
		log.Debug(e)
		return ""
	}
	key, e := ParseRSAPublicKeyFromPEM(publicKey)
	if e != nil {
		log.Debug(e)
		return ""
	}
	part, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, key, []byte(text), nil)
	if err != nil {
		log.Debug(e)
		return ""
	}

	return base64.StdEncoding.EncodeToString(part)
}

/*Base64Encode Base64Encode */
func Base64Encode(b []byte) []byte {
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(b)))
	base64.StdEncoding.Encode(buf, b)
	return buf
}

/*Base64Decode Base64Decode */
func Base64Decode(b []byte) ([]byte, error) {
	buf := make([]byte, base64.StdEncoding.DecodedLen(len(b)))
	n, err := base64.StdEncoding.Decode(buf, b)
	return buf[:n], err
}
