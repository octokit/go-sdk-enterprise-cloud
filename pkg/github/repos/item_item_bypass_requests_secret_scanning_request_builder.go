package repos

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
    i9cc13663769aa64ef02561e63455e413491540ab13cbd39c62d793d3d9c0f769 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/repos/item/item/bypassrequests/secretscanning"
)

// ItemItemBypassRequestsSecretScanningRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\bypass-requests\secret-scanning
type ItemItemBypassRequestsSecretScanningRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemBypassRequestsSecretScanningRequestBuilderGetQueryParameters lists requests to bypass secret scanning push protection in a repository.Delegated bypass must be enabled on the repository and the user must be a bypass reviewer to access this endpoint.Personal access tokens (classic) need the `security_events` scope to use this endpoint.
type ItemItemBypassRequestsSecretScanningRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // The status of the bypass request to filter on. When specified, only requests with this status will be returned.
    Request_status *i9cc13663769aa64ef02561e63455e413491540ab13cbd39c62d793d3d9c0f769.GetRequest_statusQueryParameterType `uriparametername:"request_status"`
    // Filter bypass requests by the handle of the GitHub user who requested the bypass.
    Requester *string `uriparametername:"requester"`
    // Filter bypass requests by the handle of the GitHub user who reviewed the bypass request.
    Reviewer *string `uriparametername:"reviewer"`
    // The time period to filter by.For example, `day` will filter for rule suites that occurred in the past 24 hours, and `week` will filter for insights that occurred in the past 7 days (168 hours).
    Time_period *i9cc13663769aa64ef02561e63455e413491540ab13cbd39c62d793d3d9c0f769.GetTime_periodQueryParameterType `uriparametername:"time_period"`
}
// ByBypass_request_number gets an item from the github.com/octokit/go-sdk-enterprise-cloud/pkg/github.repos.item.item.bypassRequests.secretScanning.item collection
// returns a *ItemItemBypassRequestsSecretScanningWithBypass_request_numberItemRequestBuilder when successful
func (m *ItemItemBypassRequestsSecretScanningRequestBuilder) ByBypass_request_number(bypass_request_number int32)(*ItemItemBypassRequestsSecretScanningWithBypass_request_numberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["bypass_request_number"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(bypass_request_number), 10)
    return NewItemItemBypassRequestsSecretScanningWithBypass_request_numberItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemBypassRequestsSecretScanningRequestBuilderInternal instantiates a new ItemItemBypassRequestsSecretScanningRequestBuilder and sets the default values.
func NewItemItemBypassRequestsSecretScanningRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBypassRequestsSecretScanningRequestBuilder) {
    m := &ItemItemBypassRequestsSecretScanningRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/bypass-requests/secret-scanning{?page*,per_page*,request_status*,requester*,reviewer*,time_period*}", pathParameters),
    }
    return m
}
// NewItemItemBypassRequestsSecretScanningRequestBuilder instantiates a new ItemItemBypassRequestsSecretScanningRequestBuilder and sets the default values.
func NewItemItemBypassRequestsSecretScanningRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBypassRequestsSecretScanningRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemBypassRequestsSecretScanningRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists requests to bypass secret scanning push protection in a repository.Delegated bypass must be enabled on the repository and the user must be a bypass reviewer to access this endpoint.Personal access tokens (classic) need the `security_events` scope to use this endpoint.
// returns a []SecretScanningBypassRequestable when successful
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// returns a BasicError error when the service returns a 500 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/secret-scanning/delegated-bypass#list-bypass-requests-for-secret-scanning-for-a-repository
func (m *ItemItemBypassRequestsSecretScanningRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemBypassRequestsSecretScanningRequestBuilderGetQueryParameters])([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.SecretScanningBypassRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "403": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "404": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "500": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateSecretScanningBypassRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.SecretScanningBypassRequestable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.SecretScanningBypassRequestable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists requests to bypass secret scanning push protection in a repository.Delegated bypass must be enabled on the repository and the user must be a bypass reviewer to access this endpoint.Personal access tokens (classic) need the `security_events` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemItemBypassRequestsSecretScanningRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemBypassRequestsSecretScanningRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemBypassRequestsSecretScanningRequestBuilder when successful
func (m *ItemItemBypassRequestsSecretScanningRequestBuilder) WithUrl(rawUrl string)(*ItemItemBypassRequestsSecretScanningRequestBuilder) {
    return NewItemItemBypassRequestsSecretScanningRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
