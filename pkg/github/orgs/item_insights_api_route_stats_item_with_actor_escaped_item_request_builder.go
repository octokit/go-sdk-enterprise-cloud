package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
    i4b43c8fa5dcce52af4216c9236088f8250aada65d317285fdac04ae2bff57e70 "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/orgs/item/insights/api/routestats/item/item"
)

// ItemInsightsApiRouteStatsItemWithActor_ItemRequestBuilder builds and executes requests for operations under \orgs\{org}\insights\api\route-stats\{actor_type}\{actor_id}
type ItemInsightsApiRouteStatsItemWithActor_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInsightsApiRouteStatsItemWithActor_ItemRequestBuilderGetQueryParameters get API request count statistics for an actor broken down by route within a specified time frame.
type ItemInsightsApiRouteStatsItemWithActor_ItemRequestBuilderGetQueryParameters struct {
    // The direction to sort the results by.
    Direction *i4b43c8fa5dcce52af4216c9236088f8250aada65d317285fdac04ae2bff57e70.GetDirectionQueryParameterType `uriparametername:"direction"`
    // The maximum timestamp to query for stats
    Max_timestamp *string `uriparametername:"max_timestamp"`
    // The minimum timestamp to query for stats
    Min_timestamp *string `uriparametername:"min_timestamp"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // The property to sort the results by.
    Sort []i4b43c8fa5dcce52af4216c9236088f8250aada65d317285fdac04ae2bff57e70.GetSortQueryParameterType `uriparametername:"sort"`
}
// NewItemInsightsApiRouteStatsItemWithActor_ItemRequestBuilderInternal instantiates a new ItemInsightsApiRouteStatsItemWithActor_ItemRequestBuilder and sets the default values.
func NewItemInsightsApiRouteStatsItemWithActor_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInsightsApiRouteStatsItemWithActor_ItemRequestBuilder) {
    m := &ItemInsightsApiRouteStatsItemWithActor_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/insights/api/route-stats/{actor_type}/{actor_id}?max_timestamp={max_timestamp}&min_timestamp={min_timestamp}{&direction*,page*,per_page*,sort*}", pathParameters),
    }
    return m
}
// NewItemInsightsApiRouteStatsItemWithActor_ItemRequestBuilder instantiates a new ItemInsightsApiRouteStatsItemWithActor_ItemRequestBuilder and sets the default values.
func NewItemInsightsApiRouteStatsItemWithActor_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInsightsApiRouteStatsItemWithActor_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInsightsApiRouteStatsItemWithActor_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get API request count statistics for an actor broken down by route within a specified time frame.
// returns a []ApiInsightsRouteStatsable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/orgs/api-insights#get-route-stats-by-actor
func (m *ItemInsightsApiRouteStatsItemWithActor_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemInsightsApiRouteStatsItemWithActor_ItemRequestBuilderGetQueryParameters])([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ApiInsightsRouteStatsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateApiInsightsRouteStatsFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ApiInsightsRouteStatsable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ApiInsightsRouteStatsable)
        }
    }
    return val, nil
}
// ToGetRequestInformation get API request count statistics for an actor broken down by route within a specified time frame.
// returns a *RequestInformation when successful
func (m *ItemInsightsApiRouteStatsItemWithActor_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemInsightsApiRouteStatsItemWithActor_ItemRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemInsightsApiRouteStatsItemWithActor_ItemRequestBuilder when successful
func (m *ItemInsightsApiRouteStatsItemWithActor_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemInsightsApiRouteStatsItemWithActor_ItemRequestBuilder) {
    return NewItemInsightsApiRouteStatsItemWithActor_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}