package orgs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/models"
    if20ae0ea9c2b44f635a16d34c4ee44de4452c03dd7a90d8456f97773cba6237e "github.com/octokit/go-sdk-enterprise-cloud/pkg/github/orgs/item/insights/api/userstats/item"
)

// ItemInsightsApiUserStatsWithUser_ItemRequestBuilder builds and executes requests for operations under \orgs\{org}\insights\api\user-stats\{user_id}
type ItemInsightsApiUserStatsWithUser_ItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInsightsApiUserStatsWithUser_ItemRequestBuilderGetQueryParameters get API usage statistics within an organization for a user broken down by the type of access.
type ItemInsightsApiUserStatsWithUser_ItemRequestBuilderGetQueryParameters struct {
    // Providing a substring will filter results where the actor name contains the substring. This is a case-insensitive search.
    Actor_name_substring *string `uriparametername:"actor_name_substring"`
    // The direction to sort the results by.
    Direction *if20ae0ea9c2b44f635a16d34c4ee44de4452c03dd7a90d8456f97773cba6237e.GetDirectionQueryParameterType `uriparametername:"direction"`
    // The maximum timestamp to query for stats. Defaults to the time 30 days ago. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`.
    Max_timestamp *string `uriparametername:"max_timestamp"`
    // The minimum timestamp to query for stats. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: `YYYY-MM-DDTHH:MM:SSZ`.
    Min_timestamp *string `uriparametername:"min_timestamp"`
    // The page number of the results to fetch. For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Page *int32 `uriparametername:"page"`
    // The number of results per page (max 100). For more information, see "[Using pagination in the REST API](https://docs.github.com/enterprise-cloud@latest//rest/using-the-rest-api/using-pagination-in-the-rest-api)."
    Per_page *int32 `uriparametername:"per_page"`
    // The property to sort the results by.
    Sort []if20ae0ea9c2b44f635a16d34c4ee44de4452c03dd7a90d8456f97773cba6237e.GetSortQueryParameterType `uriparametername:"sort"`
}
// NewItemInsightsApiUserStatsWithUser_ItemRequestBuilderInternal instantiates a new ItemInsightsApiUserStatsWithUser_ItemRequestBuilder and sets the default values.
func NewItemInsightsApiUserStatsWithUser_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInsightsApiUserStatsWithUser_ItemRequestBuilder) {
    m := &ItemInsightsApiUserStatsWithUser_ItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/orgs/{org}/insights/api/user-stats/{user_id}?min_timestamp={min_timestamp}{&actor_name_substring*,direction*,max_timestamp*,page*,per_page*,sort*}", pathParameters),
    }
    return m
}
// NewItemInsightsApiUserStatsWithUser_ItemRequestBuilder instantiates a new ItemInsightsApiUserStatsWithUser_ItemRequestBuilder and sets the default values.
func NewItemInsightsApiUserStatsWithUser_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInsightsApiUserStatsWithUser_ItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInsightsApiUserStatsWithUser_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get API usage statistics within an organization for a user broken down by the type of access.
// returns a []ApiInsightsUserStatsable when successful
// [API method documentation]
// 
// [API method documentation]: https://docs.github.com/enterprise-cloud@latest//rest/orgs/api-insights#get-user-stats
func (m *ItemInsightsApiUserStatsWithUser_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemInsightsApiUserStatsWithUser_ItemRequestBuilderGetQueryParameters])([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ApiInsightsUserStatsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.CreateApiInsightsUserStatsFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    val := make([]i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ApiInsightsUserStatsable, len(res))
    for i, v := range res {
        if v != nil {
            val[i] = v.(i65c45deea5ef786561f9cd3a81f83eacee03df1f39b7b57e269c7f0477b77b5d.ApiInsightsUserStatsable)
        }
    }
    return val, nil
}
// ToGetRequestInformation get API usage statistics within an organization for a user broken down by the type of access.
// returns a *RequestInformation when successful
func (m *ItemInsightsApiUserStatsWithUser_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestConfiguration[ItemInsightsApiUserStatsWithUser_ItemRequestBuilderGetQueryParameters])(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ConfigureRequestInformation(requestInfo, requestConfiguration)
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemInsightsApiUserStatsWithUser_ItemRequestBuilder when successful
func (m *ItemInsightsApiUserStatsWithUser_ItemRequestBuilder) WithUrl(rawUrl string)(*ItemInsightsApiUserStatsWithUser_ItemRequestBuilder) {
    return NewItemInsightsApiUserStatsWithUser_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
