package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemInsightsApiSummaryStatsRequestBuilder builds and executes requests for operations under \orgs\{org}\insights\api\summary-stats
type ItemInsightsApiSummaryStatsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInsightsApiSummaryStatsRequestBuilderGetQueryParameters get overall statistics of API requests made within an organization by all users and apps within a specified time frame.
type ItemInsightsApiSummaryStatsRequestBuilderGetQueryParameters struct {
    // The maximum timestamp to query for stats. Defaults to the time 30 days ago. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`.
    Max_timestamp *string `uriparametername:"max_timestamp"`
    // The minimum timestamp to query for stats. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`.
    Min_timestamp *string `uriparametername:"min_timestamp"`
}
// ByActor_type gets an item from the github.com/octokit/go-sdk-enterprise-cloud/pkg/github.orgs.item.insights.api.summaryStats.item collection
// returns a *ItemInsightsApiSummaryStatsWithActor_typeItemRequestBuilder when successful
func (m *ItemInsightsApiSummaryStatsRequestBuilder) ByActor_type(actor_type string)(*ItemInsightsApiSummaryStatsWithActor_typeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if actor_type != "" {
        urlTplParams["actor_type"] = actor_type
    }
    return NewItemInsightsApiSummaryStatsWithActor_typeItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemInsightsApiSummaryStatsRequestBuilderInternal instantiates a new ItemInsightsApiSummaryStatsRequestBuilder and sets the default values.
func NewItemInsightsApiSummaryStatsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInsightsApiSummaryStatsRequestBuilder) {
    m := &ItemInsightsApiSummaryStatsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/insights/api/summary-stats?min_timestamp={min_timestamp}{&max_timestamp*}", pathParameters),
    }
    return m
}
// NewItemInsightsApiSummaryStatsRequestBuilder instantiates a new ItemInsightsApiSummaryStatsRequestBuilder and sets the default values.
func NewItemInsightsApiSummaryStatsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInsightsApiSummaryStatsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInsightsApiSummaryStatsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get overall statistics of API requests made within an organization by all users and apps within a specified time frame.
// returns a ApiInsightsSummaryStatsable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/orgs/api-insights#get-summary-stats
func (m *ItemInsightsApiSummaryStatsRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemInsightsApiSummaryStatsRequestBuilderGetQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ApiInsightsSummaryStatsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateApiInsightsSummaryStatsFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ApiInsightsSummaryStatsable), nil
}
// ToGetRequestInformation get overall statistics of API requests made within an organization by all users and apps within a specified time frame.
// returns a *RequestInformation when successful
func (m *ItemInsightsApiSummaryStatsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemInsightsApiSummaryStatsRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// Users the users property
// returns a *ItemInsightsApiSummaryStatsUsersRequestBuilder when successful
func (m *ItemInsightsApiSummaryStatsRequestBuilder) Users()(*ItemInsightsApiSummaryStatsUsersRequestBuilder) {
    return NewItemInsightsApiSummaryStatsUsersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemInsightsApiSummaryStatsRequestBuilder when successful
func (m *ItemInsightsApiSummaryStatsRequestBuilder) WithUrl(rawUrl string)(*ItemInsightsApiSummaryStatsRequestBuilder) {
    return NewItemInsightsApiSummaryStatsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
