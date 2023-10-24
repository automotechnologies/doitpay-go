package balance

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/automotechnologies/doitpay-go/common"
	"github.com/automotechnologies/doitpay-go/utils"
)

type PublicBalanceAPI interface {

	/*
		GetBalance Get Balance.

		Get net & gross balance.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGetBalanceRequest
	*/
	GetBalance(ctx context.Context) ApiGetBalanceRequest

	// GetBalanceExecute executes the request
	//  @return InternalWebControllersMerchantApiv1BalanceStandardResponse
	GetBalanceExecute(r ApiGetBalanceRequest) (*InternalWebControllersMerchantApiv1BalanceStandardResponse, *http.Response, error)
}

type BalanceApiService struct {
	client common.IClient
}

// NewBalanceApi Create a new BalanceApi service
func NewBalanceApi(client common.IClient) PublicBalanceAPI {
	return &BalanceApiService{
		client: client,
	}
}

type ApiGetBalanceRequest struct {
	ctx        context.Context
	ApiService PublicBalanceAPI
}

func (r ApiGetBalanceRequest) Execute() (*InternalWebControllersMerchantApiv1BalanceStandardResponse, *http.Response, error) {
	return r.ApiService.GetBalanceExecute(r)
}

/*
GetBalance Get Balance.

Get net & gross balance.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetBalanceRequest
*/
func (a *BalanceApiService) GetBalance(ctx context.Context) ApiGetBalanceRequest {
	return ApiGetBalanceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return InternalWebControllersMerchantApiv1BalanceStandardResponse
func (a *BalanceApiService) GetBalanceExecute(r ApiGetBalanceRequest) (*InternalWebControllersMerchantApiv1BalanceStandardResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []common.FormFile
		localVarReturnValue *InternalWebControllersMerchantApiv1BalanceStandardResponse
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "BalanceApiService.GetBalance")
	if err != nil {
		return localVarReturnValue, nil, common.NewDoitpaySdkError(nil, "", "Error creating HTTP request: BalanceApiService.GetBalanceExecute")
	}

	localVarPath := localBasePath + "/merchant/v1/balance"

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
		return localVarReturnValue, nil, common.NewDoitpaySdkError(nil, "", "Error creating HTTP request: BalanceApiService.GetBalanceExecute")
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
