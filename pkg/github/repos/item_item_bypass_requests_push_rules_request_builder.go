package repos

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
    i47e33e59796d053e4c1aecc496d3dfe6f6063eeed71c8a93e1ea6e442a4e0c82 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/repos/item/item/bypassrequests/pushrules"
)

// ItemItemBypassRequestsPushRulesRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\bypass-requests\push-rules
type ItemItemBypassRequestsPushRulesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemBypassRequestsPushRulesRequestBuilderGetQueryParameters lists the requests made by users of a repository to bypass push protection rules
type ItemItemBypassRequestsPushRulesRequestBuilderGetQueryParameters struct {
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // The status of the bypass request to filter on. When specified, only requests with this status will be returned.
    Request_status *i47e33e59796d053e4c1aecc496d3dfe6f6063eeed71c8a93e1ea6e442a4e0c82.GetRequest_statusQueryParameterType `uriparametername:"request_status"`
    // Filter bypass requests by the handle of the GitHub user who requested the bypass.
    Requester *string `uriparametername:"requester"`
    // Filter bypass requests by the handle of the GitHub user who reviewed the bypass request.
    Reviewer *string `uriparametername:"reviewer"`
    // The time period to filter by.For example, `day` will filter for rule suites that occurred in the past 24 hours, and `week` will filter for insights that occurred in the past 7 days (168 hours).
    Time_period *i47e33e59796d053e4c1aecc496d3dfe6f6063eeed71c8a93e1ea6e442a4e0c82.GetTime_periodQueryParameterType `uriparametername:"time_period"`
}
// ByBypass_request_number gets an item from the github.com/octokit/go-sdk-enterprise-cloud/pkg/github.repos.item.item.bypassRequests.pushRules.item collection
// returns a *ItemItemBypassRequestsPushRulesWithBypass_request_numberItemRequestBuilder when successful
func (m *ItemItemBypassRequestsPushRulesRequestBuilder) ByBypass_request_number(bypass_request_number int32)(*ItemItemBypassRequestsPushRulesWithBypass_request_numberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["bypass_request_number"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(bypass_request_number), 10)
    return NewItemItemBypassRequestsPushRulesWithBypass_request_numberItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemBypassRequestsPushRulesRequestBuilderInternal instantiates a new ItemItemBypassRequestsPushRulesRequestBuilder and sets the default values.
func NewItemItemBypassRequestsPushRulesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBypassRequestsPushRulesRequestBuilder) {
    m := &ItemItemBypassRequestsPushRulesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/bypass-requests/push-rules{?page*,per_page*,request_status*,requester*,reviewer*,time_period*}", pathParameters),
    }
    return m
}
// NewItemItemBypassRequestsPushRulesRequestBuilder instantiates a new ItemItemBypassRequestsPushRulesRequestBuilder and sets the default values.
func NewItemItemBypassRequestsPushRulesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBypassRequestsPushRulesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemBypassRequestsPushRulesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get lists the requests made by users of a repository to bypass push protection rules
// returns a []PushRuleBypassRequestable when successful
// returns a BasicError error when the service returns a 404 status code
// returns a BasicError error when the service returns a 500 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/repos/bypass-requests#list-repository-push-rule-bypass-requests
func (m *ItemItemBypassRequestsPushRulesRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemBypassRequestsPushRulesRequestBuilderGetQueryParameters])([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.PushRuleBypassRequestable, error) {
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
// ToGetRequestInformation lists the requests made by users of a repository to bypass push protection rules
// returns a *RequestInformation when successful
func (m *ItemItemBypassRequestsPushRulesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemBypassRequestsPushRulesRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemBypassRequestsPushRulesRequestBuilder when successful
func (m *ItemItemBypassRequestsPushRulesRequestBuilder) WithUrl(rawUrl string)(*ItemItemBypassRequestsPushRulesRequestBuilder) {
    return NewItemItemBypassRequestsPushRulesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
