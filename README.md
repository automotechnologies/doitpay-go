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

To utilize the SDK, initialize it with your confidential API key which you can get from the [Doitpay Dashboard](https://dashboard.doitpay.co). If you haven't, you can register for a free Dashboard account [here](https://dashboard.doitpay.co/register).

```golang
dtp := doitpay.NewClient("API-KEY")
```

# Documentation

Access comprehensive API details and sample usages for each of our products by following the links provided below.

* [Balance](docs/balance/BalanceAPI.md)
* [Invoice](docs/invoice/InvoiceAPI.md)
* [Simulate](docs/simulate/SimulateAPI.md)
* [Virtual Account](docs/virtualaccount/VirtualAccountAPI.md)

Read more at:

* [Doitpay Docs](https://docs.doitpay.co/)
* [Doitpay API Reference](https://developers.doitpay.co/)

# Contact Us

support@doitpay.co