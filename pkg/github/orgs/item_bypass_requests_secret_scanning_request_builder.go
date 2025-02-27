package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
    i15defc734b13284be446bb31e32c2f6ec266db17953cad62ec3cca62e58723d1 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/orgs/item/bypassrequests/secretscanning"
)

// ItemBypassRequestsSecretScanningRequestBuilder builds and executes requests for operations under \orgs\{org}\bypass-requests\secret-scanning
type ItemBypassRequestsSecretScanningRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemBypassRequestsSecretScanningRequestBuilderGetQueryParameters list requests to bypass secret scanning push protection in an org.Delegated bypass must be enabled on repositories in the org and the user must be a bypass reviewer to access this endpoint.Personal access tokens (classic) need the `security_events` scope to use this endpoint.
type ItemBypassRequestsSecretScanningRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // The name of the repository to filter on.
    Repository_name *string `uriparametername:"repository_name"`
    // The status of the bypass request to filter on. When specified, only requests with this status will be returned.
    Request_status *i15defc734b13284be446bb31e32c2f6ec266db17953cad62ec3cca62e58723d1.GetRequest_statusQueryParameterType `uriparametername:"request_status"`
    // Filter bypass requests by the handle of the GitHub user who requested the bypass.
    Requester *string `uriparametername:"requester"`
    // Filter bypass requests by the handle of the GitHub user who reviewed the bypass request.
    Reviewer *string `uriparametername:"reviewer"`
    // The time period to filter by.For example, `day` will filter for rule suites that occurred in the past 24 hours, and `week` will filter for insights that occurred in the past 7 days (168 hours).
    Time_period *i15defc734b13284be446bb31e32c2f6ec266db17953cad62ec3cca62e58723d1.GetTime_periodQueryParameterType `uriparametername:"time_period"`
}
// NewItemBypassRequestsSecretScanningRequestBuilderInternal instantiates a new ItemBypassRequestsSecretScanningRequestBuilder and sets the default values.
func NewItemBypassRequestsSecretScanningRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemBypassRequestsSecretScanningRequestBuilder) {
    m := &ItemBypassRequestsSecretScanningRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/bypass-requests/secret-scanning{?page*,per_page*,repository_name*,request_status*,requester*,reviewer*,time_period*}", pathParameters),
    }
    return m
}
// NewItemBypassRequestsSecretScanningRequestBuilder instantiates a new ItemBypassRequestsSecretScanningRequestBuilder and sets the default values.
func NewItemBypassRequestsSecretScanningRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemBypassRequestsSecretScanningRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemBypassRequestsSecretScanningRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list requests to bypass secret scanning push protection in an org.Delegated bypass must be enabled on repositories in the org and the user must be a bypass reviewer to access this endpoint.Personal access tokens (classic) need the `security_events` scope to use this endpoint.
// returns a []SecretScanningBypassRequestable when successful
// returns a BasicError error when the service returns a 404 status code
// returns a BasicError error when the service returns a 500 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/secret-scanning/delegated-bypass#list-bypass-requests-for-secret-scanning-for-an-org
func (m *ItemBypassRequestsSecretScanningRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemBypassRequestsSecretScanningRequestBuilderGetQueryParameters])([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.SecretScanningBypassRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
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
// ToGetRequestInformation list requests to bypass secret scanning push protection in an org.Delegated bypass must be enabled on repositories in the org and the user must be a bypass reviewer to access this endpoint.Personal access tokens (classic) need the `security_events` scope to use this endpoint.
// returns a *RequestInformation when successful
func (m *ItemBypassRequestsSecretScanningRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemBypassRequestsSecretScanningRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemBypassRequestsSecretScanningRequestBuilder when successful
func (m *ItemBypassRequestsSecretScanningRequestBuilder) WithUrl(rawUrl string)(*ItemBypassRequestsSecretScanningRequestBuilder) {
    return NewItemBypassRequestsSecretScanningRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
