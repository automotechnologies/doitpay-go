package signature

import (
	"bytes"
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strings"
	"time"

	minify "github.com/tdewolff/minify/v2/json"
)

// GenerateSNAPAsymmetricSignature generates a signature for the given client secret and timestamp using the private key
func GenerateSNAPAsymmetricSignature(clientSecret string, timestamp time.Time, privateKey *rsa.PrivateKey) (string, error) {
	data := fmt.Sprintf("%s|%s", clientSecret, timestamp.Format(time.RFC3339))

	// generate signature using PKCS1 private key
	dataHash := sha256.Sum256([]byte(data))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, dataHash[:])
	if err != nil {
		return "", fmt.Errorf("failed to populate signature, err: %s", err)
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}

func GenerateSNAPSymmetricSignature(httpMethod, endpointUrl, accessToken, clientSecret string, timestamp time.Time, body []byte) (string, error) {
	r := bytes.NewBuffer(body)
	var w bytes.Buffer

	var datahashString string
	if strings.EqualFold(httpMethod, http.MethodGet) {
		datahashString = ""
	} else {
		if err := minify.Minify(nil, &w, r, nil); err != nil {
			return "", err
		}
		datahash := sha256.Sum256(w.Bytes())
		datahashString = strings.ToLower(hex.EncodeToString(datahash[:]))
	}

	stringToSign := fmt.Sprintf("%s:%s:%s:%s:%s",
		strings.ToTitle(httpMethod),
		endpointUrl,
		accessToken,
		datahashString,
		timestamp.Format(time.RFC3339),
	)

	hm := hmac.New(sha512.New, []byte(clientSecret))
	if _, err := hm.Write([]byte(stringToSign)); err != nil {
		return "", fmt.Errorf("failed to populate signature, err: %s", err)
	}

	return base64.StdEncoding.EncodeToString(hm.Sum(nil)), nil
}

func ParsePrivateKey(private []byte) (*rsa.PrivateKey, error) {
	privatePem, _ := pem.Decode(private)
	if privatePem == nil {
		return nil, errors.New("failed parse private key")
	}
	var privatePemBytes []byte
	if privatePem.Type != "RSA PRIVATE KEY" && privatePem.Type != "ENCRYPTED PRIVATE KEY" && privatePem.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("RSA private key is of the wrong type: %s", privatePem.Type)
	}
	privatePemBytes = privatePem.Bytes

	var parsedKey interface{}
	parsedKey, err := x509.ParsePKCS1PrivateKey(privatePemBytes)
	if err != nil {
		if parsedKey, err = x509.ParsePKCS8PrivateKey(privatePemBytes); err != nil {
			return nil, fmt.Errorf("unable to parse RSA private key, generating a temp one: %s", err.Error())
		}
	}

	var privateKey *rsa.PrivateKey
	var ok bool
	privateKey, ok = parsedKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("unable to parse RSA private key, generating a temp one : %s", err.Error())
	}

	return privateKey, nil
}

func ParsePublicKey(pub []byte) (*rsa.PublicKey, error) {
	pubPem, _ := pem.Decode(pub)
	if pubPem == nil {
		return nil, errors.New("failed decode public key")
	}
	if pubPem.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("RSA public key is of the wrong type, Pem Type :%s", pubPem.Type)
	}

	parsedKey, err := x509.ParsePKIXPublicKey(pubPem.Bytes)
	if err != nil {
		return nil, fmt.Errorf("unable to parse RSA public key, generating a temp one: %s", err.Error())
	}

	pubKey, ok := parsedKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("unable to parse RSA public key")
	}

	return pubKey, nil
}
