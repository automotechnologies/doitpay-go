package main

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/automotechnologies/doitpay-go/v2"
	"github.com/automotechnologies/doitpay-go/v2/models"
	"github.com/oklog/ulid"
)

var (
	ClientKey = os.Getenv("CLIENT_KEY")
	PrivateKeyPath = os.Getenv("PRIVATE_KEY_PATH")
)


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

func main() {
	dtp, err := doitpay.NewClient("01J20VM8TNRK9VBMGJAQDBSF8X", "01J20VM8W11EP9M6KBHHJZMHPW", PrivateKeyPath, doitpay.WithHost("localhost:8000"))
	if err != nil {
		log.Fatalf("Failed to create DTP client: %v", err)
	}

	PartnerReferenceNo := ulid.MustNew(ulid.Timestamp(time.Now()), nil).String()

	qrisParams := doitpay.GenerateQRISParams{
		Request: &models.QrisRequestScheme{
			PartnerReferenceNo: PartnerReferenceNo,
			Amount:     &models.QrisAmount{
				Value:    "100000.00",
			Currency: "IDR",
		},
		MerchantID:         "",
		SubMerchantID:      "",
		StoreID:            "ID2020081412725",
		ValidityPeriod:     time.Now().Add(30 * time.Minute).Format(time.RFC3339),
		AdditionalInfo:     map[string]string{
			"validTime": "9000",
			"tip":       "false",
				"qrType":    "03",
			},
		},
		ExternalID: fmt.Sprintf("%06d", rand.Intn(1000000)),
		ChannelID:  "BS",
	}

	responseTransaction, err := dtp.Qris().GenerateQRISCode(context.Background(), &qrisParams)

	if err != nil {
		log.Fatalf("Failed to generate QRIS code: %v", err)
	}

	fmt.Printf("%+v", responseTransaction)

	
}
