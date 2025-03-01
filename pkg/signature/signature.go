package signature

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"
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