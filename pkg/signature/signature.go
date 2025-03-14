package signature

import (
	"bytes"
	"crypto"
	"crypto/hmac"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
	"time"

	minify "github.com/tdewolff/minify/v2/json"
)

// GenerateSNAPAsymmetricSignature generates a signature for the given client secret and timestamp using the private key
func GenerateSNAPAsymmetricSignature(clientSecret string, timestamp time.Time, privateKey *rsa.PrivateKey) (string, error) {
	data := fmt.Sprintf("%s|%s", clientSecret, timestamp.Format("2006-01-02T15:04:05-07:00"))

	dataHash := sha256.New()
	_, err := dataHash.Write([]byte(data))
	if err != nil {
		return "", err
	}

	msgHashSum := dataHash.Sum(nil)
	signatureByte, err := rsa.SignPKCS1v15(nil, privateKey, crypto.SHA256, msgHashSum)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signatureByte), nil
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
