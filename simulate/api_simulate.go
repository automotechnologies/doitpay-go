package simulate

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

type PublicSimulateAPI interface {

	/*
		SimulatePayment SimulatePayment is handler for simulate payment system.

		Return status code 200 if the payment is completed.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiSimulatePaymentRequest
	*/
	SimulatePayment(ctx context.Context) ApiSimulatePaymentRequest

	// SimulatePaymentExecute executes the request
	SimulatePaymentExecute(r ApiSimulatePaymentRequest) (*http.Response, error)
}

type SimulateApiService struct {
	client common.IClient
}

// NewSimulateApi Create a new SimulateApi service
func NewSimulateApi(client common.IClient) PublicSimulateAPI {
	return &SimulateApiService{
		client: client,
	}
}

type ApiSimulatePaymentRequest struct {
	ctx        context.Context
	ApiService PublicSimulateAPI
	request    *InternalWebControllersMerchantApiv1SimulateSimulateRequest
}

// Request payload to simulate payment
func (r ApiSimulatePaymentRequest) Request(request InternalWebControllersMerchantApiv1SimulateSimulateRequest) ApiSimulatePaymentRequest {
	r.request = &request
	return r
}

func (r ApiSimulatePaymentRequest) Execute() (*http.Response, error) {
	return r.ApiService.SimulatePaymentExecute(r)
}

/*
SimulatePayment SimulatePayment is handler for simulate payment system.

Return status code 200 if the payment is completed.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSimulatePaymentRequest
*/
func (a *SimulateApiService) SimulatePayment(ctx context.Context) ApiSimulatePaymentRequest {
	return ApiSimulatePaymentRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *SimulateApiService) SimulatePaymentExecute(r ApiSimulatePaymentRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []common.FormFile
	)

	localBasePath, err := a.client.GetConfig().ServerURLWithContext(r.ctx, "SimulateApiService.SimulatePayment")
	if err != nil {
		return nil, common.NewDoitpaySdkError(nil, "", "Error creating HTTP request: SimulateApiService.SimulatePaymentExecute")
	}

	localVarPath := localBasePath + "/merchant/v1/simulate-payment"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.request == nil {
		return nil, utils.ReportError("request is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := utils.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := utils.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.request
	req, err := a.client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, common.NewDoitpaySdkError(nil, "", "Error creating HTTP request: SimulateApiService.SimulatePaymentExecute")
	}

	localVarHTTPResponse, err := a.client.CallAPI(req)

	localVarBody, _ := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))

	if err != nil || localVarHTTPResponse.StatusCode < 200 || localVarHTTPResponse.StatusCode >= 300 {
		doitpaySdkError := common.NewDoitpaySdkError(&localVarBody, strconv.Itoa(localVarHTTPResponse.StatusCode), localVarHTTPResponse.Status)

		return localVarHTTPResponse, doitpaySdkError
	}

	return localVarHTTPResponse, nil
}
