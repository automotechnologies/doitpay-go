package examples

import (
	"github.com/automotechnologies/doitpay-go"
	"github.com/automotechnologies/doitpay-go/client/q_r_i_s"
	"github.com/automotechnologies/doitpay-go/models"
)

func main() {
	dtp := doitpay.NewClient(
		"client_secret",
		"private_key_path",
	)


	dtp.Qris.PostQrisV10QrQrMpmGenerate(&q_r_i_s.PostQrisV10QrQrMpmGenerateParams{
		Request: &models.Apiv1QrisRequestScheme{
			Amount: &models.Apiv1QrisAmount{
				Value:    "100000",
				Currency: "IDR",
			},
		},
	})
}
