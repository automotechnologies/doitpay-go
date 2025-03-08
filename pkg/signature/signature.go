package signature

import (
	"bytes"
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
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

	if err := minify.Minify(nil, &w, r, nil); err != nil {
		return "", err
	}

	datahash := sha256.Sum256(w.Bytes())

	datahashString := strings.ToLower(hex.EncodeToString(datahash[:]))

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

