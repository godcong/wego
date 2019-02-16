package cipher

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"golang.org/x/xerrors"
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
	privateKey []byte
	publicKey  []byte
}

// NewRSA ...
func NewRSA(option *Option) Cipher {
	return &cryptRSA{
		privateKey: []byte(option.RSAPrivate),
		publicKey:  []byte(option.RSAPublic),
	}
}

// Encrypt ...
func (c *cryptRSA) Encrypt(data interface{}) ([]byte, error) {
	key, err := ParseRSAPublicKeyFromPEM(c.publicKey)
	if err != nil {
		return nil, xerrors.Errorf("ParseRSAPublicKeyFromPEM:%w", err)
	}
	part, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, key, parseBytes(data), nil)
	if err != nil {
		return nil, xerrors.Errorf("EncryptOAEP:%w", err)
	}

	buf := make([]byte, base64.StdEncoding.EncodedLen(len(part)))
	base64.StdEncoding.Encode(buf, part)
	return buf, nil
}

// Decrypt ...
func (c *cryptRSA) Decrypt(data interface{}) ([]byte, error) {
	key, err := ParseRSAPrivateKeyFromPEM(c.privateKey)
	if err != nil {
		return nil, err
	}

	t, err := Base64Decode(parseBytes(data))
	if err != nil {
		return nil, err
	}

	b, err := rsa.DecryptOAEP(sha1.New(), rand.Reader, key, t, nil)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Type ...
func (*cryptRSA) Type() CryptType {
	return RSA
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
