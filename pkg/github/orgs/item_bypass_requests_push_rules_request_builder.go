package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
    i811b95e605f5f099e4ac868d225c22ef396cff018613f70cf8a20f250c7f3cec "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/orgs/item/bypassrequests/pushrules"
)

// ItemBypassRequestsPushRulesRequestBuilder builds and executes requests for operations under \orgs\{org}\bypass-requests\push-rules
type ItemBypassRequestsPushRulesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemBypassRequestsPushRulesRequestBuilderGetQueryParameters lists the requests made by users of a repository to bypass push protection rules within an organization.
type ItemBypassRequestsPushRulesRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // The name of the repository to filter on.
    Repository_name *string `uriparametername:"repository_name"`
    // The status of the bypass request to filter on. When specified, only requests with this status will be returned.
    Request_status *i811b95e605f5f099e4ac868d225c22ef396cff018613f70cf8a20f250c7f3cec.GetRequest_statusQueryParameterType `uriparametername:"request_status"`
    // Filter bypass requests by the handle of the GitHub user who requested the bypass.
    Requester *string `uriparametername:"requester"`
    // Filter bypass requests by the handle of the GitHub user who reviewed the bypass request.
    Reviewer *string `uriparametername:"reviewer"`
    // The time period to filter by.For example, `day` will filter for rule suites that occurred in the past 24 hours, and `week` will filter for insights that occurred in the past 7 days (168 hours).
    Time_period *i811b95e605f5f099e4ac868d225c22ef396cff018613f70cf8a20f250c7f3cec.GetTime_periodQueryParameterType `uriparametername:"time_period"`
}
// NewItemBypassRequestsPushRulesRequestBuilderInternal instantiates a new ItemBypassRequestsPushRulesRequestBuilder and sets the default values.
func NewItemBypassRequestsPushRulesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemBypassRequestsPushRulesRequestBuilder) {
    m := &ItemBypassRequestsPushRulesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/bypass-requests/push-rules{?page*,per_page*,repository_name*,request_status*,requester*,reviewer*,time_period*}", pathParameters),
    }
    return m
}
// NewItemBypassRequestsPushRulesRequestBuilder instantiates a new ItemBypassRequestsPushRulesRequestBuilder and sets the default values.
func NewItemBypassRequestsPushRulesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemBypassRequestsPushRulesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemBypassRequestsPushRulesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the requests made by users of a repository to bypass push protection rules within an organization.
// returns a []PushRuleBypassRequestable when successful
// returns a BasicError error when the service returns a 404 status code
// returns a BasicError error when the service returns a 500 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/orgs/bypass-requests#list-push-rule-bypass-requests-within-an-organization
func (m *ItemBypassRequestsPushRulesRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemBypassRequestsPushRulesRequestBuilderGetQueryParameters])([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.PushRuleBypassRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "500": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreatePushRuleBypassRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.PushRuleBypassRequestable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.PushRuleBypassRequestable)
        }
    }
    return val, nil
}
// ToGetRequestInformation lists the requests made by users of a repository to bypass push protection rules within an organization.
// returns a *RequestInformation when successful
func (m *ItemBypassRequestsPushRulesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemBypassRequestsPushRulesRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemBypassRequestsPushRulesRequestBuilder when successful
func (m *ItemBypassRequestsPushRulesRequestBuilder) WithUrl(rawUrl string)(*ItemBypassRequestsPushRulesRequestBuilder) {
    return NewItemBypassRequestsPushRulesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
