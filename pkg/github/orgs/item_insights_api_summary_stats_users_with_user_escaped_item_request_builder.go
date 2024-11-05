package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
)

// ItemInsightsApiSummaryStatsUsersWithUser_ItemRequestBuilder builds and executes requests for operations under \orgs\{org}\insights\api\summary-stats\users\{user_id}
type ItemInsightsApiSummaryStatsUsersWithUser_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInsightsApiSummaryStatsUsersWithUser_ItemRequestBuilderGetQueryParameters get overall statistics of API requests within the organization for a user.
type ItemInsightsApiSummaryStatsUsersWithUser_ItemRequestBuilderGetQueryParameters struct {
    // The maximum timestamp to query for stats
    Max_timestamp *string `uriparametername:"max_timestamp"`
    // The minimum timestamp to query for stats
    Min_timestamp *string `uriparametername:"min_timestamp"`
}
// NewItemInsightsApiSummaryStatsUsersWithUser_ItemRequestBuilderInternal instantiates a new ItemInsightsApiSummaryStatsUsersWithUser_ItemRequestBuilder and sets the default values.
func NewItemInsightsApiSummaryStatsUsersWithUser_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInsightsApiSummaryStatsUsersWithUser_ItemRequestBuilder) {
    m := &ItemInsightsApiSummaryStatsUsersWithUser_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/insights/api/summary-stats/users/{user_id}?max_timestamp={max_timestamp}&min_timestamp={min_timestamp}", pathParameters),
    }
    return m
}
// NewItemInsightsApiSummaryStatsUsersWithUser_ItemRequestBuilder instantiates a new ItemInsightsApiSummaryStatsUsersWithUser_ItemRequestBuilder and sets the default values.
func NewItemInsightsApiSummaryStatsUsersWithUser_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInsightsApiSummaryStatsUsersWithUser_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInsightsApiSummaryStatsUsersWithUser_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get overall statistics of API requests within the organization for a user.
// returns a ApiInsightsSummaryStatsable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/orgs/api-insights#get-summary-stats-by-user
func (m *ItemInsightsApiSummaryStatsUsersWithUser_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemInsightsApiSummaryStatsUsersWithUser_ItemRequestBuilderGetQueryParameters])(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ApiInsightsSummaryStatsable, error) {
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
// ToGetRequestInformation get overall statistics of API requests within the organization for a user.
// returns a *RequestInformation when successful
func (m *ItemInsightsApiSummaryStatsUsersWithUser_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemInsightsApiSummaryStatsUsersWithUser_ItemRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemInsightsApiSummaryStatsUsersWithUser_ItemRequestBuilder when successful
func (m *ItemInsightsApiSummaryStatsUsersWithUser_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemInsightsApiSummaryStatsUsersWithUser_ItemRequestBuilder) {
    return NewItemInsightsApiSummaryStatsUsersWithUser_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}