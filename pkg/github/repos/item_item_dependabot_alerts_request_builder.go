package repos

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
    ia923be11ce5ebc3a458be409140439c7fb35124f1d0ecdd2442c934d7434365e "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/repos/item/item/dependabot/alerts"
)

// ItemItemDependabotAlertsRequestBuilder builds and executes requests for operations under \repos\{owner-id}\{repo-id}\dependabot\alerts
type ItemItemDependabotAlertsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemDependabotAlertsRequestBuilderGetQueryParameters oAuth app tokens and personal access tokens (classic) need the `security_events` scope to use this endpoint. If this endpoint is only used with public repositories, the token can use the `public_repo` scope instead.
type ItemItemDependabotAlertsRequestBuilderGetQueryParameters struct {
    // A cursor, as given in the [Link header](https://docs.github.com/enterprise-cloud@latest//rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for results after this cursor. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    After *string `uriparametername:"after"`
    // A cursor, as given in the [Link header](https://docs.github.com/enterprise-cloud@latest//rest/guides/using-pagination-in-the-rest-api#using-link-headers). If specified, the query only searches for results before this cursor. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Before *string `uriparametername:"before"`
    // The direction to sort the results by.
    Direction *ia923be11ce5ebc3a458be409140439c7fb35124f1d0ecdd2442c934d7434365e.GetDirectionQueryParameterType `uriparametername:"direction"`
    // A comma-separated list of ecosystems. If specified, only alerts for these ecosystems will be returned.Can be: `composer`, `go`, `maven`, `npm`, `nuget`, `pip`, `pub`, `rubygems`, `rust`
    Ecosystem *string `uriparametername:"ecosystem"`
    // **Deprecated**. The number of results per page (max 100), starting from the first matching result.This parameter must not be used in combination with `last`.Instead, use `per_page` in combination with `after` to fetch the first page of results.
    First *int32 `uriparametername:"first"`
    // **Deprecated**. The number of results per page (max 100), starting from the last matching result.This parameter must not be used in combination with `first`.Instead, use `per_page` in combination with `before` to fetch the last page of results.
    Last *int32 `uriparametername:"last"`
    // A comma-separated list of full manifest paths. If specified, only alerts for these manifests will be returned.
    Manifest *string `uriparametername:"manifest"`
    // A comma-separated list of package names. If specified, only alerts for these packages will be returned.
    Package *string `uriparametername:"package"`
    // **Deprecated**. Page number of the results to fetch. Use cursor-based pagination with `before` or `after` instead.
    // Deprecated: 
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    // Deprecated: 
    Per_page *int32 `uriparametername:"per_page"`
    // The scope of the vulnerable dependency. If specified, only alerts with this scope will be returned.
    Scope *ia923be11ce5ebc3a458be409140439c7fb35124f1d0ecdd2442c934d7434365e.GetScopeQueryParameterType `uriparametername:"scope"`
    // A comma-separated list of severities. If specified, only alerts with these severities will be returned.Can be: `low`, `medium`, `high`, `critical`
    Severity *string `uriparametername:"severity"`
    // The property by which to sort the results.`created` means when the alert was created.`updated` means when the alert's state last changed.
    Sort *ia923be11ce5ebc3a458be409140439c7fb35124f1d0ecdd2442c934d7434365e.GetSortQueryParameterType `uriparametername:"sort"`
    // A comma-separated list of states. If specified, only alerts with these states will be returned.Can be: `auto_dismissed`, `dismissed`, `fixed`, `open`
    State *string `uriparametername:"state"`
}
// ByAlert_number gets an item from the github.com/octokit/go-sdk-enterprise-cloud/pkg/github.repos.item.item.dependabot.alerts.item collection
// returns a *ItemItemDependabotAlertsWithAlert_numberItemRequestBuilder when successful
func (m *ItemItemDependabotAlertsRequestBuilder) ByAlert_number(alert_number int32)(*ItemItemDependabotAlertsWithAlert_numberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["alert_number"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(alert_number), 10)
    return NewItemItemDependabotAlertsWithAlert_numberItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemItemDependabotAlertsRequestBuilderInternal instantiates a new ItemItemDependabotAlertsRequestBuilder and sets the default values.
func NewItemItemDependabotAlertsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDependabotAlertsRequestBuilder) {
    m := &ItemItemDependabotAlertsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/repos/{owner%2Did}/{repo%2Did}/dependabot/alerts{?after*,before*,direction*,ecosystem*,first*,last*,manifest*,package*,page*,per_page*,scope*,severity*,sort*,state*}", pathParameters),
    }
    return m
}
// NewItemItemDependabotAlertsRequestBuilder instantiates a new ItemItemDependabotAlertsRequestBuilder and sets the default values.
func NewItemItemDependabotAlertsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemDependabotAlertsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemDependabotAlertsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get oAuth app tokens and personal access tokens (classic) need the `security_events` scope to use this endpoint. If this endpoint is only used with public repositories, the token can use the `public_repo` scope instead.
// returns a []DependabotAlertable when successful
// returns a BasicError error when the service returns a 400 status code
// returns a BasicError error when the service returns a 403 status code
// returns a BasicError error when the service returns a 404 status code
// returns a ValidationErrorSimple error when the service returns a 422 status code
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/dependabot/alerts#list-dependabot-alerts-for-a-repository
func (m *ItemItemDependabotAlertsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemDependabotAlertsRequestBuilderGetQueryParameters])([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.DependabotAlertable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "403": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "404": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateBasicErrorFromDiscriminatorValue,
        "422": i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateValidationErrorSimpleFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateDependabotAlertFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    val := make([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.DependabotAlertable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.DependabotAlertable)
        }
    }
    return val, nil
}
// ToGetRequestInformation oAuth app tokens and personal access tokens (classic) need the `security_events` scope to use this endpoint. If this endpoint is only used with public repositories, the token can use the `public_repo` scope instead.
// returns a *RequestInformation when successful
func (m *ItemItemDependabotAlertsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemItemDependabotAlertsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemDependabotAlertsRequestBuilder when successful
func (m *ItemItemDependabotAlertsRequestBuilder) WithUrl(rawUrl string)(*ItemItemDependabotAlertsRequestBuilder) {
    return NewItemItemDependabotAlertsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
