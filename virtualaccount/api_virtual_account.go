package virtualaccount

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

type PublicVirtualAccountAPI interface {

	/*
		CreateVirtualAccount Create virtual account data.

		Create and return virtual account data.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiCreateVirtualAccountRequest
	*/
	CreateVirtualAccount(ctx context.Context) ApiCreateVirtualAccountRequest

	// CreateVirtualAccountExecute executes the request
	//  @return InternalWebControllersMerchantApiv1VirtualaccountStandardResponse
	CreateVirtualAccountExecute(r ApiCreateVirtualAccountRequest) (*InternalWebControllersMerchantApiv1VirtualaccountStandardResponse, *http.Response, error)

	/*
		GetVirtualAccountById Fetch virtual account data by ID.

		Fetch virtual account data by ID.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param virtualAccountId Virtual Account ID
		@return ApiGetVirtualAccountByIdRequest
	*/
	GetVirtualAccountById(ctx context.Context, virtualAccountId int32) ApiGetVirtualAccountByIdRequest

	// GetVirtualAccountByIdExecute executes the request
	//  @return InternalWebControllersMerchantApiv1VirtualaccountStandardResponse
	GetVirtualAccountByIdExecute(r ApiGetVirtualAccountByIdRequest) (*InternalWebControllersMerchantApiv1VirtualaccountStandardResponse, *http.Response, error)

	/*
		GetVirtualAccountByNumber Fetch virtual account data by virtual account number.

		Fetch virtual account data by virtual account number.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param virtualAccountNumber Virtual Account Number
		@return ApiGetVirtualAccountByNumberRequest
	*/
	GetVirtualAccountByNumber(ctx context.Context, virtualAccountNumber string) ApiGetVirtualAccountByNumberRequest

	// GetVirtualAccountByNumberExecute executes the request
	//  @return InternalWebControllersMerchantApiv1VirtualaccountStandardResponse
	GetVirtualAccountByNumberExecute(r ApiGetVirtualAccountByNumberRequest) (*InternalWebControllersMerchantApiv1VirtualaccountStandardResponse, *http.Response, error)

	/*
		GetVirtualAccounts List all created virtual account data.

		List all created virtual account data.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGetVirtualAccountsRequest
	*/
	GetVirtualAccounts(ctx context.Context) ApiGetVirtualAccountsRequest

	// GetVirtualAccountsExecute executes the request
	//  @return InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination
	GetVirtualAccountsExecute(r ApiGetVirtualAccountsRequest) (*InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination, *http.Response, error)
}

type VirtualAccountApiService struct {
	client common.IClient
}

// NewVirtualAccountApi Create a new VirtualAccountApi service
func NewVirtualAccountApi(client common.IClient) PublicVirtualAccountAPI {
	return &VirtualAccountApiService{
		client: client,
	}
}

type ApiCreateVirtualAccountRequest struct {
	ctx        context.Context
	ApiService PublicVirtualAccountAPI
	request    *InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest
}

// Request payload to create virtual account
func (r ApiCreateVirtualAccountRequest) Request(request InternalWebControllersMerchantApiv1VirtualaccountCreateVirtualAccountRequest) ApiCreateVirtualAccountRequest {
	r.request = &request
	return r
}

func (r ApiCreateVirtualAccountRequest) Execute() (*InternalWebControllersMerchantApiv1VirtualaccountStandardResponse, *http.Response, error) {
	return r.ApiService.CreateVirtualAccountExecute(r)
}

/*
CreateVirtualAccount Create virtual account data.

Create and return virtual account data.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateVirtualAccountRequest
*/
func (a *VirtualAccountApiService) CreateVirtualAccount(ctx context.Context) ApiCreateVirtualAccountRequest {
	return ApiCreateVirtualAccountRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return InternalWebControllersMerchantApiv1VirtualaccountStandardResponse
func (a *VirtualAccountApiService) CreateVirtualAccountExecute(r ApiCreateVirtualAccountRequest) (*InternalWebControllersMerchantApiv1VirtualaccountStandardResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []common.FormFile
		localVarReturnValue *InternalWebControllersMerchantApiv1VirtualaccountStandardResponse
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "VirtualAccountApiService.CreateVirtualAccount")
	if err != nil {
		return localVarReturnValue, nil, common.NewDoitpaySdkError(nil, "", "Error creating HTTP request: VirtualAccountApiService.CreateVirtualAccountExecute")
	}

	localVarPath := localBasePath + "/merchant/v1/virtual-account"

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
		return localVarReturnValue, nil, common.NewDoitpaySdkError(nil, "", "Error creating HTTP request: VirtualAccountApiService.CreateVirtualAccountExecute")
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

type ApiGetVirtualAccountByIdRequest struct {
	ctx              context.Context
	ApiService       PublicVirtualAccountAPI
	virtualAccountId int32
}

func (r ApiGetVirtualAccountByIdRequest) Execute() (*InternalWebControllersMerchantApiv1VirtualaccountStandardResponse, *http.Response, error) {
	return r.ApiService.GetVirtualAccountByIdExecute(r)
}

/*
GetVirtualAccountById Fetch virtual account data by ID.

Fetch virtual account data by ID.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param virtualAccountId Virtual Account ID
	@return ApiGetVirtualAccountByIdRequest
*/
func (a *VirtualAccountApiService) GetVirtualAccountById(ctx context.Context, virtualAccountId int32) ApiGetVirtualAccountByIdRequest {
	return ApiGetVirtualAccountByIdRequest{
		ApiService:       a,
		ctx:              ctx,
		virtualAccountId: virtualAccountId,
	}
}

// Execute executes the request
//
//	@return InternalWebControllersMerchantApiv1VirtualaccountStandardResponse
func (a *VirtualAccountApiService) GetVirtualAccountByIdExecute(r ApiGetVirtualAccountByIdRequest) (*InternalWebControllersMerchantApiv1VirtualaccountStandardResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []common.FormFile
		localVarReturnValue *InternalWebControllersMerchantApiv1VirtualaccountStandardResponse
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "VirtualAccountApiService.GetVirtualAccountById")
	if err != nil {
		return localVarReturnValue, nil, common.NewDoitpaySdkError(nil, "", "Error creating HTTP request: VirtualAccountApiService.GetVirtualAccountByIdExecute")
	}

	localVarPath := localBasePath + "/merchant/v1/virtual-account/{virtualAccountId}"
	localVarPath = strings.Replace(localVarPath, "{"+"virtualAccountId"+"}", url.PathEscape(utils.ParameterValueToString(r.virtualAccountId, "virtualAccountId")), -1)

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
		return localVarReturnValue, nil, common.NewDoitpaySdkError(nil, "", "Error creating HTTP request: VirtualAccountApiService.GetVirtualAccountByIdExecute")
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

type ApiGetVirtualAccountByNumberRequest struct {
	ctx                  context.Context
	ApiService           PublicVirtualAccountAPI
	virtualAccountNumber string
}

func (r ApiGetVirtualAccountByNumberRequest) Execute() (*InternalWebControllersMerchantApiv1VirtualaccountStandardResponse, *http.Response, error) {
	return r.ApiService.GetVirtualAccountByNumberExecute(r)
}

/*
GetVirtualAccountByNumber Fetch virtual account data by virtual account number.

Fetch virtual account data by virtual account number.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param virtualAccountNumber Virtual Account Number
	@return ApiGetVirtualAccountByNumberRequest
*/
func (a *VirtualAccountApiService) GetVirtualAccountByNumber(ctx context.Context, virtualAccountNumber string) ApiGetVirtualAccountByNumberRequest {
	return ApiGetVirtualAccountByNumberRequest{
		ApiService:           a,
		ctx:                  ctx,
		virtualAccountNumber: virtualAccountNumber,
	}
}

// Execute executes the request
//
//	@return InternalWebControllersMerchantApiv1VirtualaccountStandardResponse
func (a *VirtualAccountApiService) GetVirtualAccountByNumberExecute(r ApiGetVirtualAccountByNumberRequest) (*InternalWebControllersMerchantApiv1VirtualaccountStandardResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []common.FormFile
		localVarReturnValue *InternalWebControllersMerchantApiv1VirtualaccountStandardResponse
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "VirtualAccountApiService.GetVirtualAccountByNumber")
	if err != nil {
		return localVarReturnValue, nil, common.NewDoitpaySdkError(nil, "", "Error creating HTTP request: VirtualAccountApiService.GetVirtualAccountByNumberExecute")
	}

	localVarPath := localBasePath + "/merchant/v1/virtual-account/number/{virtualAccountNumber}"
	localVarPath = strings.Replace(localVarPath, "{"+"virtualAccountNumber"+"}", url.PathEscape(utils.ParameterValueToString(r.virtualAccountNumber, "virtualAccountNumber")), -1)

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
		return localVarReturnValue, nil, common.NewDoitpaySdkError(nil, "", "Error creating HTTP request: VirtualAccountApiService.GetVirtualAccountByNumberExecute")
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

type ApiGetVirtualAccountsRequest struct {
	ctx        context.Context
	ApiService PublicVirtualAccountAPI
	page       *int32
	limit      *int32
	statuses   *string
}

// Page number
func (r ApiGetVirtualAccountsRequest) Page(page int32) ApiGetVirtualAccountsRequest {
	r.page = &page
	return r
}

// Page limit
func (r ApiGetVirtualAccountsRequest) Limit(limit int32) ApiGetVirtualAccountsRequest {
	r.limit = &limit
	return r
}

// Comma-separated list of statuses to filter by. Example: ?statuses&#x3D;ACTIVE,PENDING
func (r ApiGetVirtualAccountsRequest) Statuses(statuses string) ApiGetVirtualAccountsRequest {
	r.statuses = &statuses
	return r
}

func (r ApiGetVirtualAccountsRequest) Execute() (*InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination, *http.Response, error) {
	return r.ApiService.GetVirtualAccountsExecute(r)
}

/*
GetVirtualAccounts List all created virtual account data.

List all created virtual account data.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetVirtualAccountsRequest
*/
func (a *VirtualAccountApiService) GetVirtualAccounts(ctx context.Context) ApiGetVirtualAccountsRequest {
	return ApiGetVirtualAccountsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination
func (a *VirtualAccountApiService) GetVirtualAccountsExecute(r ApiGetVirtualAccountsRequest) (*InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []common.FormFile
		localVarReturnValue *InternalWebControllersMerchantApiv1VirtualaccountStandardResponseWithPagination
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "VirtualAccountApiService.GetVirtualAccounts")
	if err != nil {
		return localVarReturnValue, nil, common.NewDoitpaySdkError(nil, "", "Error creating HTTP request: VirtualAccountApiService.GetVirtualAccountsExecute")
	}

	localVarPath := localBasePath + "/merchant/v1/virtual-account"

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
		return localVarReturnValue, nil, common.NewDoitpaySdkError(nil, "", "Error creating HTTP request: VirtualAccountApiService.GetVirtualAccountsExecute")
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
