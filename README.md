# Doitpay Go SDK

The official Doitpay Go SDK offers an easy and user-friendly method to access Doitpay's REST API in applications coded in Go.

* Package version: 1.0.0

# Getting Started

## Installation

Install doitpay-go in your Go application:

```shell
go get github.com/automotechnologies/doitpay-go
```

Place the package in your project directory and include the following in the import:

```golang
import doitpay "github.com/automotechnologies/doitpay-go"
```

## Authorization

To utilize the SDK, initialize it with your client key and partner ID. You can get these credentials from the [Doitpay Dashboard](https://dashboard.doitpay.co). If you haven't, you can register for a free Dashboard account [here](https://dashboard.doitpay.co/register).

```golang
dtp, err := doitpay.NewClient(
    "YOUR_PARTNER_ID",
    "YOUR_CLIENT_SECRET",
    doitpay.WithPrivateKeyPath("path/to/private/key.pem"), //  specify private key path or private key bytes
    doitpay.WithHost("api.doitpay.co"), // Optional: specify host
)
if err != nil {
    log.Fatalf("Failed to create DTP client: %v", err)
}
```

## Usage Examples

### Creating a Payment

```golang
import (
    "context"
    "github.com/automotechnologies/doitpay-go"
    "github.com/automotechnologies/doitpay-go/models"
    "github.com/oklog/ulid"
    "time"
)

func main() {
    // Initialize client as shown above
    
    // Generate a unique reference number
    referenceNo := ulid.MustNew(ulid.Timestamp(time.Now()), nil).String()
    
    // Create payment parameters
    params := &doitpay.CreatePaymentParams{
        Request: &models.CreateEwalletRequest{
            // Fill in the request details
            PartnerReferenceNo: referenceNo,
            // Add other required fields
        },
        ExternalID: "your-external-id",
        ChannelID: "your-channel-id",
        RequestID: "your-request-id",
    }

    // Create payment
    response, err := dtp.Payment().Create(context.Background(), params)
    if err != nil {
        log.Fatalf("Failed to create payment: %v", err)
    }
}
```

### Checking Payment Status

```golang
params := &doitpay.CheckPaymentStatusParams{
    Request: &models.CheckPaymentStatusEwalletRequest{
        // Fill in the request details
    },
    ExternalID: "your-external-id",
    ChannelID: "your-channel-id",
    RequestID: "your-request-id",
}

response, err := dtp.Payment().GetStatus(context.Background(), params)
if err != nil {
    log.Fatalf("Failed to check payment status: %v", err)
}
```

### Canceling a Payment

```golang
params := &doitpay.CancelPaymentParams{
    Request: &models.CancelHostToHostPaymentRequest{
        // Fill in the request details
    },
    ExternalID: "your-external-id",
    ChannelID: "your-channel-id",
}

response, err := dtp.Payment().Cancel(context.Background(), params)
if err != nil {
    log.Fatalf("Failed to cancel payment: %v", err)
}
```

# Documentation

Access comprehensive API details and sample usages for each of our products by following the links provided below.

* [Balance](docs/balance/BalanceAPI.md)
* [Invoice](docs/invoice/InvoiceAPI.md)
* [Payment](docs/payment/PaymentAPI.md)
* [Simulate](docs/simulate/SimulateAPI.md)
* [Virtual Account](docs/virtualaccount/VirtualAccountAPI.md)

Read more at:

* [Doitpay Docs](https://docs.doitpay.co/)
* [Doitpay API Reference](https://developers.doitpay.co/)

# Contact Us

support@doitpay.co
