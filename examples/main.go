package examples

import (
	"log"

	"github.com/automotechnologies/doitpay-go"
)

func main() {
	_, err := doitpay.NewClient(
		"client_secret",
		"private_key_path",
	)
	if err != nil {
		log.Fatal(err)
	}


	
}
