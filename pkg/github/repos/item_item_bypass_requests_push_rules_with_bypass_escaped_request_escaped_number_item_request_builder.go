package repos

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemItemBypassRequestsPushRulesWithBypass_request_numberItemRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\bypass-requests\push-rules\{bypass_request_number}
type ItemItemBypassRequestsPushRulesWithBypass_request_numberItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemItemBypassRequestsPushRulesWithBypass_request_numberItemRequestBuilderInternal instantiates a new ItemItemBypassRequestsPushRulesWithBypass_request_numberItemRequestBuilder and sets the default values.
func NewItemItemBypassRequestsPushRulesWithBypass_request_numberItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBypassRequestsPushRulesWithBypass_request_numberItemRequestBuilder) {
    m := &ItemItemBypassRequestsPushRulesWithBypass_request_numberItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/bypass-requests/push-rules/{bypass_request_number}", pathParameters),
    }
    return m
}
// NewItemItemBypassRequestsPushRulesWithBypass_request_numberItemRequestBuilder instantiates a new ItemItemBypassRequestsPushRulesWithBypass_request_numberItemRequestBuilder and sets the default values.
func NewItemItemBypassRequestsPushRulesWithBypass_request_numberItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemBypassRequestsPushRulesWithBypass_request_numberItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemBypassRequestsPushRulesWithBypass_request_numberItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get information about a request to bypass push protection rules for a repository.
// returns a PushRuleBypassRequestable when successful
// returns a BasicError error when the service returns a 404 status code
// returns a BasicError error when the service returns a 500 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/repos/bypass-requests#get-a-repository-push-bypass-request
func (m *ItemItemBypassRequestsPushRulesWithBypass_request_numberItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.PushRuleBypassRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "404": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "500": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreatePushRuleBypassRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.PushRuleBypassRequestable), nil
}
// ToGetRequestInformation get information about a request to bypass push protection rules for a repository.
// returns a *RequestInformation when successful
func (m *ItemItemBypassRequestsPushRulesWithBypass_request_numberItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DefaultQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemBypassRequestsPushRulesWithBypass_request_numberItemRequestBuilder when successful
func (m *ItemItemBypassRequestsPushRulesWithBypass_request_numberItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemBypassRequestsPushRulesWithBypass_request_numberItemRequestBuilder) {
    return NewItemItemBypassRequestsPushRulesWithBypass_request_numberItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
