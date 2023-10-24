package invoice

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/automotechnologies/doitpay-go/common"
	"github.com/automotechnologies/doitpay-go/utils"
)

type PublicInvoiceAPI interface {

	/*
		CreateInvoice Create invoice for a payment.

		Create and return invoice data.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiCreateInvoiceRequest
	*/
	CreateInvoice(ctx context.Context) ApiCreateInvoiceRequest

	// CreateInvoiceExecute executes the request
	//  @return InternalWebControllersMerchantApiv1InvoiceStandardResponse
	CreateInvoiceExecute(r ApiCreateInvoiceRequest) (*InternalWebControllersMerchantApiv1InvoiceStandardResponse, *http.Response, error)

	/*
		DownloadInvoice Download CSV file.

		Download a CSV file of invoices.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiDownloadInvoiceRequest
	*/
	DownloadInvoice(ctx context.Context) ApiDownloadInvoiceRequest

	// DownloadInvoiceExecute executes the request
	//  @return string
	DownloadInvoiceExecute(r ApiDownloadInvoiceRequest) (string, *http.Response, error)

	/*
		ExpireInvoice Expire invoice by invoice ID.

		Expire invoice by invoice ID.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param invoiceId Invoice ID to fetch payment method
		@return ApiExpireInvoiceRequest
	*/
	ExpireInvoice(ctx context.Context, invoiceId string) ApiExpireInvoiceRequest

	// ExpireInvoiceExecute executes the request
	//  @return InternalWebControllersMerchantApiv1InvoiceStandardResponse
	ExpireInvoiceExecute(r ApiExpireInvoiceRequest) (*InternalWebControllersMerchantApiv1InvoiceStandardResponse, *http.Response, error)

	/*
		GetInvoiceById Fetch invoice data by ID.

		Fetch invoice data by ID.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param invoiceId Invoice ID to fetch data
		@return ApiGetInvoiceByIdRequest
	*/
	GetInvoiceById(ctx context.Context, invoiceId string) ApiGetInvoiceByIdRequest

	// GetInvoiceByIdExecute executes the request
	//  @return InternalWebControllersMerchantApiv1InvoiceStandardResponse
	GetInvoiceByIdExecute(r ApiGetInvoiceByIdRequest) (*InternalWebControllersMerchantApiv1InvoiceStandardResponse, *http.Response, error)

	/*
		GetInvoices List all created invoice data

		List all created invoice data.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGetInvoicesRequest
	*/
	GetInvoices(ctx context.Context) ApiGetInvoicesRequest

	// GetInvoicesExecute executes the request
	//  @return InternalWebControllersMerchantApiv1InvoiceStandardResponseWithPagination
	GetInvoicesExecute(r ApiGetInvoicesRequest) (*InternalWebControllersMerchantApiv1InvoiceStandardResponseWithPagination, *http.Response, error)

	/*
		GetPaymentMethodById Fetch payment method by invoice ID.

		Fetch payment method by invoice ID.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param invoiceId Invoice ID to fetch payment method
		@return ApiGetPaymentMethodByIdRequest
	*/
	GetPaymentMethodById(ctx context.Context, invoiceId string) ApiGetPaymentMethodByIdRequest

	// GetPaymentMethodByIdExecute executes the request
	//  @return InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse
	GetPaymentMethodByIdExecute(r ApiGetPaymentMethodByIdRequest) (*InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse, *http.Response, error)

	/*
		UpdateInvoiceById Update pending invoice

		Update pending invoice

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param invoiceId invoice id to update
		@return ApiUpdateInvoiceByIdRequest
	*/
	UpdateInvoiceById(ctx context.Context, invoiceId string) ApiUpdateInvoiceByIdRequest

	// UpdateInvoiceByIdExecute executes the request
	//  @return InternalWebControllersMerchantApiv1InvoiceStandardResponse
	UpdateInvoiceByIdExecute(r ApiUpdateInvoiceByIdRequest) (*InternalWebControllersMerchantApiv1InvoiceStandardResponse, *http.Response, error)
}

type InvoiceApiService struct {
	client common.IClient
}

// NewInvoiceApi Create a new InvoiceApi service
func NewInvoiceApi(client common.IClient) PublicInvoiceAPI {
	return &InvoiceApiService{
		client: client,
	}
}

type ApiCreateInvoiceRequest struct {
	ctx        context.Context
	ApiService PublicInvoiceAPI
	request    *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest
}

// Request payload to create invoice
func (r ApiCreateInvoiceRequest) Request(request InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) ApiCreateInvoiceRequest {
	r.request = &request
	return r
}

func (r ApiCreateInvoiceRequest) Execute() (*InternalWebControllersMerchantApiv1InvoiceStandardResponse, *http.Response, error) {
	return r.ApiService.CreateInvoiceExecute(r)
}

/*
CreateInvoice Create invoice for a payment.

Create and return invoice data.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateInvoiceRequest
*/
func (a *InvoiceApiService) CreateInvoice(ctx context.Context) ApiCreateInvoiceRequest {
	return ApiCreateInvoiceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return InternalWebControllersMerchantApiv1InvoiceStandardResponse
func (a *InvoiceApiService) CreateInvoiceExecute(r ApiCreateInvoiceRequest) (*InternalWebControllersMerchantApiv1InvoiceStandardResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []common.FormFile
		localVarReturnValue *InternalWebControllersMerchantApiv1InvoiceStandardResponse
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "InvoiceApiService.CreateInvoice")
	if err != nil {
		return localVarReturnValue, nil, common.NewDoitpaySdkError(nil, "", "Error creating HTTP request: InvoiceApiService.CreateInvoiceExecute")
	}

	localVarPath := localBasePath + "/merchant/v1/invoice"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.request == nil {
		return localVarReturnValue, nil, common.NewDoitpaySdkError(nil, "", "request is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := utils.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := utils.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.request
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, common.NewDoitpaySdkError(nil, "", "Error creating HTTP request: InvoiceApiService.CreateInvoiceExecute")
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)

	localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))

	if err != nil || localVarHTTPResponse.StatusCode < 200 || localVarHTTPResponse.StatusCode >= 300 {
		doitpaySdkError := common.NewDoitpaySdkError(&localVarBody, strconv.Itoa(localVarHTTPResponse.StatusCode), localVarHTTPResponse.Status)

		return localVarReturnValue, localVarHTTPResponse, doitpaySdkError
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		doitpaySdkError := common.NewDoitpaySdkError(&localVarBody, strconv.Itoa(localVarHTTPResponse.StatusCode), localVarHTTPResponse.Status)

		return localVarReturnValue, localVarHTTPResponse, doitpaySdkError
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDownloadInvoiceRequest struct {
	ctx        context.Context
	ApiService PublicInvoiceAPI
	page       *int32
	limit      *int32
	statuses   *string
}

// Page number
func (r ApiDownloadInvoiceRequest) Page(page int32) ApiDownloadInvoiceRequest {
	r.page = &page
	return r
}

// Page limit
func (r ApiDownloadInvoiceRequest) Limit(limit int32) ApiDownloadInvoiceRequest {
	r.limit = &limit
	return r
}

// Comma-separated list of statuses to filter by. Example: ?statuses=ACTIVE,PENDING
func (r ApiDownloadInvoiceRequest) Statuses(statuses string) ApiDownloadInvoiceRequest {
	r.statuses = &statuses
	return r
}

func (r ApiDownloadInvoiceRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.DownloadInvoiceExecute(r)
}

/*
DownloadInvoice Download CSV file.

Download a CSV file of invoices.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiDownloadInvoiceRequest
*/
func (a *InvoiceApiService) DownloadInvoice(ctx context.Context) ApiDownloadInvoiceRequest {
	return ApiDownloadInvoiceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return string
func (a *InvoiceApiService) DownloadInvoiceExecute(r ApiDownloadInvoiceRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []common.FormFile
		localVarReturnValue string
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "InvoiceApiService.DownloadInvoice")
	if err != nil {
		return localVarReturnValue, nil, common.NewDoitpaySdkError(nil, "", "Error creating HTTP request: InvoiceApiService.DownloadInvoiceExecute")
	}

	localVarPath := localBasePath + "/merchant/v1/invoice/download/csv"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	} else {
		var defaultValue int32 = 1
		r.page = &defaultValue
	}
	if r.limit != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 10
		r.limit = &defaultValue
	}
	if r.statuses != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "statuses", r.statuses, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := utils.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := utils.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, common.NewDoitpaySdkError(nil, "", "Error creating HTTP request: InvoiceApiService.DownloadInvoiceExecute")
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)

	localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))

	if err != nil || localVarHTTPResponse.StatusCode < 200 || localVarHTTPResponse.StatusCode >= 300 {
		doitpaySdkError := common.NewDoitpaySdkError(&localVarBody, strconv.Itoa(localVarHTTPResponse.StatusCode), localVarHTTPResponse.Status)

		return localVarReturnValue, localVarHTTPResponse, doitpaySdkError
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		doitpaySdkError := common.NewDoitpaySdkError(&localVarBody, strconv.Itoa(localVarHTTPResponse.StatusCode), localVarHTTPResponse.Status)

		return localVarReturnValue, localVarHTTPResponse, doitpaySdkError
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiExpireInvoiceRequest struct {
	ctx        context.Context
	ApiService PublicInvoiceAPI
	invoiceId  string
}

func (r ApiExpireInvoiceRequest) Execute() (*InternalWebControllersMerchantApiv1InvoiceStandardResponse, *http.Response, error) {
	return r.ApiService.ExpireInvoiceExecute(r)
}

/*
ExpireInvoice Expire invoice by invoice ID.

Expire invoice by invoice ID.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param invoiceId Invoice ID to fetch payment method
	@return ApiExpireInvoiceRequest
*/
func (a *InvoiceApiService) ExpireInvoice(ctx context.Context, invoiceId string) ApiExpireInvoiceRequest {
	return ApiExpireInvoiceRequest{
		ApiService: a,
		ctx:        ctx,
		invoiceId:  invoiceId,
	}
}

// Execute executes the request
//
//	@return InternalWebControllersMerchantApiv1InvoiceStandardResponse
func (a *InvoiceApiService) ExpireInvoiceExecute(r ApiExpireInvoiceRequest) (*InternalWebControllersMerchantApiv1InvoiceStandardResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []common.FormFile
		localVarReturnValue *InternalWebControllersMerchantApiv1InvoiceStandardResponse
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "InvoiceApiService.ExpireInvoice")
	if err != nil {
		return localVarReturnValue, nil, common.NewDoitpaySdkError(nil, "", "Error creating HTTP request: InvoiceApiService.ExpireInvoiceExecute")
	}

	localVarPath := localBasePath + "/merchant/v1/invoice/{invoice_id}/expire"
	localVarPath = strings.Replace(localVarPath, "{"+"invoice_id"+"}", url.PathEscape(utils.ParameterValueToString(r.invoiceId, "invoiceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := utils.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := utils.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, common.NewDoitpaySdkError(nil, "", "Error creating HTTP request: InvoiceApiService.ExpireInvoiceExecute")
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)

	localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))

	if err != nil || localVarHTTPResponse.StatusCode < 200 || localVarHTTPResponse.StatusCode >= 300 {
		doitpaySdkError := common.NewDoitpaySdkError(&localVarBody, strconv.Itoa(localVarHTTPResponse.StatusCode), localVarHTTPResponse.Status)

		return localVarReturnValue, localVarHTTPResponse, doitpaySdkError
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		doitpaySdkError := common.NewDoitpaySdkError(&localVarBody, strconv.Itoa(localVarHTTPResponse.StatusCode), localVarHTTPResponse.Status)

		return localVarReturnValue, localVarHTTPResponse, doitpaySdkError
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetInvoiceByIdRequest struct {
	ctx        context.Context
	ApiService PublicInvoiceAPI
	invoiceId  string
}

func (r ApiGetInvoiceByIdRequest) Execute() (*InternalWebControllersMerchantApiv1InvoiceStandardResponse, *http.Response, error) {
	return r.ApiService.GetInvoiceByIdExecute(r)
}

/*
GetInvoiceById Fetch invoice data by ID.

Fetch invoice data by ID.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param invoiceId Invoice ID to fetch data
	@return ApiGetInvoiceByIdRequest
*/
func (a *InvoiceApiService) GetInvoiceById(ctx context.Context, invoiceId string) ApiGetInvoiceByIdRequest {
	return ApiGetInvoiceByIdRequest{
		ApiService: a,
		ctx:        ctx,
		invoiceId:  invoiceId,
	}
}

// Execute executes the request
//
//	@return InternalWebControllersMerchantApiv1InvoiceStandardResponse
func (a *InvoiceApiService) GetInvoiceByIdExecute(r ApiGetInvoiceByIdRequest) (*InternalWebControllersMerchantApiv1InvoiceStandardResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []common.FormFile
		localVarReturnValue *InternalWebControllersMerchantApiv1InvoiceStandardResponse
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "InvoiceApiService.GetInvoiceById")
	if err != nil {
		return localVarReturnValue, nil, common.NewDoitpaySdkError(nil, "", "Error creating HTTP request: InvoiceApiService.GetInvoiceByIdExecute")
	}

	localVarPath := localBasePath + "/merchant/v1/invoice/{invoice_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"invoice_id"+"}", url.PathEscape(utils.ParameterValueToString(r.invoiceId, "invoiceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := utils.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := utils.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, common.NewDoitpaySdkError(nil, "", "Error creating HTTP request: InvoiceApiService.GetInvoiceByIdExecute")
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)

	localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))

	if err != nil || localVarHTTPResponse.StatusCode < 200 || localVarHTTPResponse.StatusCode >= 300 {
		doitpaySdkError := common.NewDoitpaySdkError(&localVarBody, strconv.Itoa(localVarHTTPResponse.StatusCode), localVarHTTPResponse.Status)

		return localVarReturnValue, localVarHTTPResponse, doitpaySdkError
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		doitpaySdkError := common.NewDoitpaySdkError(&localVarBody, strconv.Itoa(localVarHTTPResponse.StatusCode), localVarHTTPResponse.Status)

		return localVarReturnValue, localVarHTTPResponse, doitpaySdkError
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetInvoicesRequest struct {
	ctx        context.Context
	ApiService PublicInvoiceAPI
	page       *int32
	limit      *int32
	statuses   *string
}

// Page number
func (r ApiGetInvoicesRequest) Page(page int32) ApiGetInvoicesRequest {
	r.page = &page
	return r
}

// Page limit
func (r ApiGetInvoicesRequest) Limit(limit int32) ApiGetInvoicesRequest {
	r.limit = &limit
	return r
}

// Comma-separated list of statuses to filter by. Example: ?statuses&#x3D;ACTIVE,PENDING
func (r ApiGetInvoicesRequest) Statuses(statuses string) ApiGetInvoicesRequest {
	r.statuses = &statuses
	return r
}

func (r ApiGetInvoicesRequest) Execute() (*InternalWebControllersMerchantApiv1InvoiceStandardResponseWithPagination, *http.Response, error) {
	return r.ApiService.GetInvoicesExecute(r)
}

/*
GetInvoices List all created invoice data

List all created invoice data.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetInvoicesRequest
*/
func (a *InvoiceApiService) GetInvoices(ctx context.Context) ApiGetInvoicesRequest {
	return ApiGetInvoicesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return InternalWebControllersMerchantApiv1InvoiceStandardResponseWithPagination
func (a *InvoiceApiService) GetInvoicesExecute(r ApiGetInvoicesRequest) (*InternalWebControllersMerchantApiv1InvoiceStandardResponseWithPagination, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []common.FormFile
		localVarReturnValue *InternalWebControllersMerchantApiv1InvoiceStandardResponseWithPagination
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "InvoiceApiService.GetInvoices")
	if err != nil {
		return localVarReturnValue, nil, common.NewDoitpaySdkError(nil, "", "Error creating HTTP request: InvoiceApiService.GetInvoicesExecute")
	}

	localVarPath := localBasePath + "/merchant/v1/invoice"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	} else {
		var defaultValue int32 = 1
		r.page = &defaultValue
	}
	if r.limit != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 10
		r.limit = &defaultValue
	}
	if r.statuses != nil {
		utils.ParameterAddToHeaderOrQuery(localVarQueryParams, "statuses", r.statuses, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := utils.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := utils.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, common.NewDoitpaySdkError(nil, "", "Error creating HTTP request: InvoiceApiService.CreateInvoiceExecute")
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)

	localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))

	if err != nil || localVarHTTPResponse.StatusCode < 200 || localVarHTTPResponse.StatusCode >= 300 {
		doitpaySdkError := common.NewDoitpaySdkError(&localVarBody, strconv.Itoa(localVarHTTPResponse.StatusCode), localVarHTTPResponse.Status)

		return localVarReturnValue, localVarHTTPResponse, doitpaySdkError
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		doitpaySdkError := common.NewDoitpaySdkError(&localVarBody, strconv.Itoa(localVarHTTPResponse.StatusCode), localVarHTTPResponse.Status)

		return localVarReturnValue, localVarHTTPResponse, doitpaySdkError
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPaymentMethodByIdRequest struct {
	ctx        context.Context
	ApiService PublicInvoiceAPI
	invoiceId  string
}

func (r ApiGetPaymentMethodByIdRequest) Execute() (*InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse, *http.Response, error) {
	return r.ApiService.GetPaymentMethodByIdExecute(r)
}

/*
GetPaymentMethodById Fetch payment method by invoice ID.

Fetch payment method by invoice ID.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param invoiceId Invoice ID to fetch payment method
	@return ApiGetPaymentMethodByIdRequest
*/
func (a *InvoiceApiService) GetPaymentMethodById(ctx context.Context, invoiceId string) ApiGetPaymentMethodByIdRequest {
	return ApiGetPaymentMethodByIdRequest{
		ApiService: a,
		ctx:        ctx,
		invoiceId:  invoiceId,
	}
}

// Execute executes the request
//
//	@return InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse
func (a *InvoiceApiService) GetPaymentMethodByIdExecute(r ApiGetPaymentMethodByIdRequest) (*InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []common.FormFile
		localVarReturnValue *InternalWebControllersMerchantApiv1InvoiceStandardPaymentMethodResponse
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "InvoiceApiService.GetPaymentMethodById")
	if err != nil {
		return localVarReturnValue, nil, common.NewDoitpaySdkError(nil, "", "Error creating HTTP request: InvoiceApiService.GetPaymentMethodByIdExecute")
	}

	localVarPath := localBasePath + "/merchant/v1/invoice/{invoice_id}/payment-method"
	localVarPath = strings.Replace(localVarPath, "{"+"invoice_id"+"}", url.PathEscape(utils.ParameterValueToString(r.invoiceId, "invoiceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := utils.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := utils.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, common.NewDoitpaySdkError(nil, "", "Error creating HTTP request: InvoiceApiService.GetPaymentMethodByIdExecute")
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)

	localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))

	if err != nil || localVarHTTPResponse.StatusCode < 200 || localVarHTTPResponse.StatusCode >= 300 {
		doitpaySdkError := common.NewDoitpaySdkError(&localVarBody, strconv.Itoa(localVarHTTPResponse.StatusCode), localVarHTTPResponse.Status)

		return localVarReturnValue, localVarHTTPResponse, doitpaySdkError
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		doitpaySdkError := common.NewDoitpaySdkError(&localVarBody, strconv.Itoa(localVarHTTPResponse.StatusCode), localVarHTTPResponse.Status)

		return localVarReturnValue, localVarHTTPResponse, doitpaySdkError
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateInvoiceByIdRequest struct {
	ctx        context.Context
	ApiService PublicInvoiceAPI
	invoiceId  string
	request    *InternalWebControllersMerchantApiv1InvoiceInvoiceRequest
}

// Invoice Request Body
func (r ApiUpdateInvoiceByIdRequest) Request(request InternalWebControllersMerchantApiv1InvoiceInvoiceRequest) ApiUpdateInvoiceByIdRequest {
	r.request = &request
	return r
}

func (r ApiUpdateInvoiceByIdRequest) Execute() (*InternalWebControllersMerchantApiv1InvoiceStandardResponse, *http.Response, error) {
	return r.ApiService.UpdateInvoiceByIdExecute(r)
}

/*
UpdateInvoiceById Update pending invoice

Update pending invoice

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param invoiceId invoice id to update
	@return ApiUpdateInvoiceByIdRequest
*/
func (a *InvoiceApiService) UpdateInvoiceById(ctx context.Context, invoiceId string) ApiUpdateInvoiceByIdRequest {
	return ApiUpdateInvoiceByIdRequest{
		ApiService: a,
		ctx:        ctx,
		invoiceId:  invoiceId,
	}
}

// Execute executes the request
//
//	@return InternalWebControllersMerchantApiv1InvoiceStandardResponse
func (a *InvoiceApiService) UpdateInvoiceByIdExecute(r ApiUpdateInvoiceByIdRequest) (*InternalWebControllersMerchantApiv1InvoiceStandardResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []common.FormFile
		localVarReturnValue *InternalWebControllersMerchantApiv1InvoiceStandardResponse
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "InvoiceApiService.UpdateInvoiceById")
	if err != nil {
		return localVarReturnValue, nil, common.NewDoitpaySdkError(nil, "", "Error creating HTTP request: InvoiceApiService.UpdateInvoiceByIdExecute")
	}

	localVarPath := localBasePath + "/merchant/v1/invoice/{invoice_id}/update"
	localVarPath = strings.Replace(localVarPath, "{"+"invoice_id"+"}", url.PathEscape(utils.ParameterValueToString(r.invoiceId, "invoiceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.request == nil {
		return localVarReturnValue, nil, utils.ReportError("request is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := utils.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := utils.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.request
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, common.NewDoitpaySdkError(nil, "", "Error creating HTTP request: InvoiceApiService.UpdateInvoiceByIdExecute")
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)

	localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))

	if err != nil || localVarHTTPResponse.StatusCode < 200 || localVarHTTPResponse.StatusCode >= 300 {
		doitpaySdkError := common.NewDoitpaySdkError(&localVarBody, strconv.Itoa(localVarHTTPResponse.StatusCode), localVarHTTPResponse.Status)

		return localVarReturnValue, localVarHTTPResponse, doitpaySdkError
	}

	err = a.client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		doitpaySdkError := common.NewDoitpaySdkError(&localVarBody, strconv.Itoa(localVarHTTPResponse.StatusCode), localVarHTTPResponse.Status)

		return localVarReturnValue, localVarHTTPResponse, doitpaySdkError
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
