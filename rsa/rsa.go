package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io/ioutil"
)

var (
	ErrorKeyMustBePEMEncoded       = errors.New("key must be pem encoded")
	ErrorNotECPublicKey      error = errors.New("Key is not a valid ECDSA public key")
	ErrorNotECPrivateKey           = errors.New("Key is not a valid ECDSA private key")
	ErrorNotRSAPrivateKey          = errors.New("Key is not a valid RSA private key")
	ErrorNotRSAPublicKey           = errors.New("Key is not a valid RSA public key")
)

// Parse PEM encoded PKCS1 or PKCS8 private key
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

// Parse PEM encoded PKCS1 or PKCS8 public key
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

func Decrypt(pri string, text string) string {
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

	b, e := rsa.DecryptPKCS1v15(rand.Reader, key, t)
	if e != nil {
		return ""
	}
	return string(b)
}

func Encrypt(pub string, text string) string {
	publicKey, e := ioutil.ReadFile(pub)
	if e != nil {
		return ""
	}

	key, e := ParseRSAPublicKeyFromPEM(publicKey)
	if e != nil {
		return ""
	}

	b, e := rsa.EncryptPKCS1v15(rand.Reader, key, []byte(text))
	if e != nil {
		return ""
	}
	return base64.RawURLEncoding.EncodeToString(b)
}

func Base64Encode(b []byte) []byte {
	buf := make([]byte, base64.RawURLEncoding.EncodedLen(len(b)))
	base64.RawURLEncoding.Encode(buf, b)
	return buf
}

// Base64Decode
func Base64Decode(b []byte) ([]byte, error) {
	buf := make([]byte, base64.RawURLEncoding.DecodedLen(len(b)))
	n, err := base64.RawURLEncoding.Decode(buf, b)
	return buf[:n], err
}
