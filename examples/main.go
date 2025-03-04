package examples

import (
	"context"
	"log"
	"os"

	"github.com/automotechnologies/doitpay-go/v2"
	"github.com/automotechnologies/doitpay-go/v2/models"
)

var (
	ClientKey = os.Getenv("CLIENT_KEY")
	PrivateKeyPath = os.Getenv("PRIVATE_KEY_PATH")
)


func main() {
	dtp, err := doitpay.NewClient(
		ClientKey,
		PrivateKeyPath,
	)
	if err != nil {
		log.Fatal(err)
	}

	dtp.Qris().GenerateQRISCode(context.Background(), &models.QrisRequestScheme{
		Amount: &models.QrisAmount{
			Value: "10000",
			Currency: "IDR",
		},
		MerchantID: "1234567890",
	})

	
}
